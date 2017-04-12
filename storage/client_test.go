package storage

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	chk "gopkg.in/check.v1"
)

// Hook up gocheck to testing
func Test(t *testing.T) { chk.TestingT(t) }

type StorageClientSuite struct{}

var _ = chk.Suite(&StorageClientSuite{})

// getBasicClient returns a test client from storage credentials in the env
func getBasicClient(c *chk.C) *Client {
	name := os.Getenv("ACCOUNT_NAME")
	if name == "" {
		name = dummyStorageAccount
	}
	key := os.Getenv("ACCOUNT_KEY")
	if key == "" {
		key = dummyMiniStorageKey
	}
	cli, err := NewBasicClient(name, key)
	c.Assert(err, chk.IsNil)

	return &cli
}

func (client *Client) appendRecorder(c *chk.C) *recorder.Recorder {
	tests := strings.Split(c.TestName(), ".")
	path := filepath.Join(recordingsFolder, tests[0], tests[1])
	rec, err := recorder.New(path)
	c.Assert(err, chk.IsNil)
	client.HTTPClient = &http.Client{
		Transport: rec,
	}
	rec.SetMatcher(func(r *http.Request, i cassette.Request) bool {
		return compareMethods(r, i) &&
			compareURLs(r, i) &&
			compareHeaders(r, i) &&
			compareBodies(r, i)
	})
	return rec
}

func (client *Client) usesDummies() bool {
	key, err := base64.StdEncoding.DecodeString(dummyMiniStorageKey)
	if err != nil {
		return false
	}
	if string(client.accountKey) == string(key) &&
		client.accountName == dummyStorageAccount {
		return true
	}
	return false
}

func compareMethods(r *http.Request, i cassette.Request) bool {
	return r.Method == i.Method
}

func compareURLs(r *http.Request, i cassette.Request) bool {
	newURL := modifyURL(r.URL)
	return newURL.String() == i.URL
}

func modifyURL(url *url.URL) *url.URL {
	// The URL host looks like this...
	// accountname.service.storageEndpointSuffix
	parts := strings.Split(url.Host, ".")
	// parts[0] corresponds to the storage account name, so it can be (almost) any string
	// parts[1] corresponds to the service name (table, blob, etc.).
	if !(parts[1] == blobServiceName ||
		parts[1] == tableServiceName ||
		parts[1] == queueServiceName ||
		parts[1] == fileServiceName) {
		return nil
	}
	// The rest of the host depends on which Azure cloud is used
	storageEndpointSuffix := strings.Join(parts[2:], ".")
	if !(storageEndpointSuffix == azure.PublicCloud.StorageEndpointSuffix ||
		storageEndpointSuffix == azure.USGovernmentCloud.StorageEndpointSuffix ||
		storageEndpointSuffix == azure.ChinaCloud.StorageEndpointSuffix ||
		storageEndpointSuffix == azure.GermanCloud.StorageEndpointSuffix) {
		return nil
	}

	host := dummyStorageAccount + "." + parts[1] + "." + azure.PublicCloud.StorageEndpointSuffix
	newURL := url
	newURL.Host = host
	return newURL
}

func compareHeaders(r *http.Request, i cassette.Request) bool {
	requestHeaders := r.Header
	cassetteHeaders := i.Headers
	// Some headers shall not be compared...
	requestHeaders.Del("User-Agent")
	requestHeaders.Del("Authorization")
	requestHeaders.Del("X-Ms-Date")

	cassetteHeaders.Del("User-Agent")
	cassetteHeaders.Del("Authorization")
	cassetteHeaders.Del("X-Ms-Date")

	srcURLstr := requestHeaders.Get("X-Ms-Copy-Source")
	if srcURLstr != "" {
		srcURL, err := url.Parse(srcURLstr)
		if err != nil {
			return false
		}
		modifiedURL := modifyURL(srcURL)
		requestHeaders.Set("X-Ms-Copy-Source", modifiedURL.String())
	}

	// Do not compare the complete Content-Type header in table batch requests
	if isBatchOp(r.URL.String()) {
		// They all start like this, but then they have a UUID...
		ctPrefixBatch := "multipart/mixed; boundary=batch_"
		contentTypeRequest := requestHeaders.Get("Content-Type")
		contentTypeCassette := cassetteHeaders.Get("Content-Type")
		if !(strings.HasPrefix(contentTypeRequest, ctPrefixBatch) &&
			strings.HasPrefix(contentTypeCassette, ctPrefixBatch)) {
			return false
		}
		requestHeaders.Del("Content-Type")
		cassetteHeaders.Del("Content-Type")
	}

	return reflect.DeepEqual(requestHeaders, cassetteHeaders)
}

func compareBodies(r *http.Request, i cassette.Request) bool {
	body := bytes.Buffer{}
	if r.Body != nil {
		_, err := body.ReadFrom(r.Body)
		if err != nil {
			return false
		}
		r.Body = ioutil.NopCloser(&body)
	}
	// Comparing bodies in table batch operations is trickier, because the bodies include UUIDs
	if isBatchOp(r.URL.String()) {
		return compareBatchBodies(body.String(), i.Body)
	}
	return body.String() == i.Body
}

func compareBatchBodies(rBody, cBody string) bool {
	// UUIDs in the batch body look like this...
	// 2d7f2323-1e42-11e7-8c6c-6451064d81e8
	exp, err := regexp.Compile("[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}")
	if err != nil {
		return false
	}
	rBody = replaceStorageAccount(replaceUUIDs(rBody, exp))
	cBody = replaceUUIDs(cBody, exp)
	return rBody == cBody
}

func replaceUUIDs(body string, exp *regexp.Regexp) string {
	indexes := exp.FindAllStringIndex(body, -1)
	for _, pair := range indexes {
		body = strings.Replace(body, body[pair[0]:pair[1]], "00000000-0000-0000-0000-000000000000", -1)
	}
	return body
}

func isBatchOp(url string) bool {
	return url == "https://golangrocksonazure.table.core.windows.net/$batch"
}

//getEmulatorClient returns a test client for Azure Storeage Emulator
func getEmulatorClient(c *chk.C) Client {
	cli, err := NewBasicClient(StorageEmulatorAccountName, "")
	c.Assert(err, chk.IsNil)
	return cli
}

func (s *StorageClientSuite) TestNewEmulatorClient(c *chk.C) {
	cli, err := NewBasicClient(StorageEmulatorAccountName, "")
	c.Assert(err, chk.IsNil)
	c.Assert(cli.accountName, chk.Equals, StorageEmulatorAccountName)
	expectedKey, err := base64.StdEncoding.DecodeString(StorageEmulatorAccountKey)
	c.Assert(err, chk.IsNil)
	c.Assert(cli.accountKey, chk.DeepEquals, expectedKey)
}

func (s *StorageClientSuite) TestMalformedKeyError(c *chk.C) {
	_, err := NewBasicClient(dummyStorageAccount, "malformed")
	c.Assert(err, chk.ErrorMatches, "azure: malformed storage account key: .*")
}

func (s *StorageClientSuite) TestGetBaseURL_Basic_Https(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, dummyMiniStorageKey)
	c.Assert(err, chk.IsNil)
	c.Assert(cli.apiVersion, chk.Equals, DefaultAPIVersion)
	c.Assert(err, chk.IsNil)
	c.Assert(cli.getBaseURL("table"), chk.Equals, "https://golangrocksonazure.table.core.windows.net")
}

func (s *StorageClientSuite) TestGetBaseURL_Custom_NoHttps(c *chk.C) {
	apiVersion := "2015-01-01" // a non existing one
	cli, err := NewClient(dummyStorageAccount, dummyMiniStorageKey, "core.chinacloudapi.cn", apiVersion, false)
	c.Assert(err, chk.IsNil)
	c.Assert(cli.apiVersion, chk.Equals, apiVersion)
	c.Assert(cli.getBaseURL("table"), chk.Equals, "http://golangrocksonazure.table.core.chinacloudapi.cn")
}

func (s *StorageClientSuite) TestGetBaseURL_StorageEmulator(c *chk.C) {
	cli, err := NewBasicClient(StorageEmulatorAccountName, StorageEmulatorAccountKey)
	c.Assert(err, chk.IsNil)

	type test struct{ service, expected string }
	tests := []test{
		{blobServiceName, "http://127.0.0.1:10000"},
		{tableServiceName, "http://127.0.0.1:10002"},
		{queueServiceName, "http://127.0.0.1:10001"},
	}
	for _, i := range tests {
		baseURL := cli.getBaseURL(i.service)
		c.Assert(baseURL, chk.Equals, i.expected)
	}
}

func (s *StorageClientSuite) TestGetEndpoint_None(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)
	output := cli.getEndpoint(blobServiceName, "", url.Values{})
	c.Assert(output, chk.Equals, "https://golangrocksonazure.blob.core.windows.net/")
}

func (s *StorageClientSuite) TestGetEndpoint_PathOnly(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)
	output := cli.getEndpoint(blobServiceName, "path", url.Values{})
	c.Assert(output, chk.Equals, "https://golangrocksonazure.blob.core.windows.net/path")
}

func (s *StorageClientSuite) TestGetEndpoint_ParamsOnly(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)
	params := url.Values{}
	params.Set("a", "b")
	params.Set("c", "d")
	output := cli.getEndpoint(blobServiceName, "", params)
	c.Assert(output, chk.Equals, "https://golangrocksonazure.blob.core.windows.net/?a=b&c=d")
}

func (s *StorageClientSuite) TestGetEndpoint_Mixed(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)
	params := url.Values{}
	params.Set("a", "b")
	params.Set("c", "d")
	output := cli.getEndpoint(blobServiceName, "path", params)
	c.Assert(output, chk.Equals, "https://golangrocksonazure.blob.core.windows.net/path?a=b&c=d")
}

func (s *StorageClientSuite) TestGetEndpoint_StorageEmulator(c *chk.C) {
	cli, err := NewBasicClient(StorageEmulatorAccountName, StorageEmulatorAccountKey)
	c.Assert(err, chk.IsNil)

	type test struct{ service, expected string }
	tests := []test{
		{blobServiceName, "http://127.0.0.1:10000/devstoreaccount1/"},
		{tableServiceName, "http://127.0.0.1:10002/devstoreaccount1/"},
		{queueServiceName, "http://127.0.0.1:10001/devstoreaccount1/"},
	}
	for _, i := range tests {
		endpoint := cli.getEndpoint(i.service, "", url.Values{})
		c.Assert(endpoint, chk.Equals, i.expected)
	}
}

func (s *StorageClientSuite) Test_getStandardHeaders(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)

	headers := cli.getStandardHeaders()
	c.Assert(len(headers), chk.Equals, 3)
	c.Assert(headers["x-ms-version"], chk.Equals, cli.apiVersion)
	if _, ok := headers["x-ms-date"]; !ok {
		c.Fatal("Missing date header")
	}
	c.Assert(headers[userAgentHeader], chk.Equals, cli.getDefaultUserAgent())
}

func (s *StorageClientSuite) TestReturnsStorageServiceError(c *chk.C) {
	// attempt to delete a nonexisting container
	cli := getBlobClient(c)
	rec := cli.client.appendRecorder(c)
	defer rec.Stop()

	cnt := cli.GetContainerReference(containerName(c))
	_, err := cnt.delete(nil)
	c.Assert(err, chk.NotNil)

	v, ok := err.(AzureStorageServiceError)
	c.Check(ok, chk.Equals, true)
	c.Assert(v.StatusCode, chk.Equals, 404)
	c.Assert(v.Code, chk.Equals, "ContainerNotFound")
	c.Assert(v.Code, chk.Not(chk.Equals), "")
	c.Assert(v.RequestID, chk.Not(chk.Equals), "")
}

func (s *StorageClientSuite) TestReturnsStorageServiceError_withoutResponseBody(c *chk.C) {
	// HEAD on non-existing blob
	cli := getBlobClient(c)
	rec := cli.client.appendRecorder(c)
	defer rec.Stop()

	cnt := cli.GetContainerReference("non-existing-container")
	b := cnt.GetBlobReference("non-existing-blob")
	err := b.GetProperties(nil)

	c.Assert(err, chk.NotNil)
	c.Assert(err, chk.FitsTypeOf, AzureStorageServiceError{})

	v, ok := err.(AzureStorageServiceError)
	c.Check(ok, chk.Equals, true)
	c.Assert(v.StatusCode, chk.Equals, http.StatusNotFound)
	c.Assert(v.Code, chk.Equals, "404 The specified container does not exist.")
	c.Assert(v.RequestID, chk.Not(chk.Equals), "")
	c.Assert(v.Message, chk.Equals, "no response body was available for error status code")
}

func (s *StorageClientSuite) Test_createServiceClients(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)

	ua := cli.getDefaultUserAgent()

	headers := cli.getStandardHeaders()
	c.Assert(headers[userAgentHeader], chk.Equals, ua)
	c.Assert(cli.userAgent, chk.Equals, ua)

	b := cli.GetBlobService()
	c.Assert(b.client.userAgent, chk.Equals, ua+" "+blobServiceName)
	c.Assert(cli.userAgent, chk.Equals, ua)

	t := cli.GetTableService()
	c.Assert(t.client.userAgent, chk.Equals, ua+" "+tableServiceName)
	c.Assert(cli.userAgent, chk.Equals, ua)

	q := cli.GetQueueService()
	c.Assert(q.client.userAgent, chk.Equals, ua+" "+queueServiceName)
	c.Assert(cli.userAgent, chk.Equals, ua)

	f := cli.GetFileService()
	c.Assert(f.client.userAgent, chk.Equals, ua+" "+fileServiceName)
	c.Assert(cli.userAgent, chk.Equals, ua)
}

func (s *StorageClientSuite) TestAddToUserAgent(c *chk.C) {
	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)

	ua := cli.getDefaultUserAgent()

	err = cli.AddToUserAgent("rofl")
	c.Assert(err, chk.IsNil)
	c.Assert(cli.userAgent, chk.Equals, ua+" rofl")

	err = cli.AddToUserAgent("")
	c.Assert(err, chk.NotNil)
}

func (s *StorageClientSuite) Test_protectUserAgent(c *chk.C) {
	extraheaders := map[string]string{
		"1":             "one",
		"2":             "two",
		"3":             "three",
		userAgentHeader: "four",
	}

	cli, err := NewBasicClient(dummyStorageAccount, "YmFy")
	c.Assert(err, chk.IsNil)

	ua := cli.getDefaultUserAgent()

	got := cli.protectUserAgent(extraheaders)
	c.Assert(cli.userAgent, chk.Equals, ua+" four")
	c.Assert(got, chk.HasLen, 3)
	c.Assert(got, chk.DeepEquals, map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	})
}
