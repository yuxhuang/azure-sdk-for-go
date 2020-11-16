// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// UsageOperations contains the methods for the Usage group.
type UsageOperations interface {
	// List - Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
	List(location string, options *UsageListOptions) ListUsagesResultPager
}

// UsageClient implements the UsageOperations interface.
// Don't use this type directly, use NewUsageClient() instead.
type UsageClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewUsageClient creates a new instance of UsageClient with the specified values.
func NewUsageClient(con *armcore.Connection, subscriptionID string) UsageOperations {
	return &UsageClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *UsageClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// List - Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
func (client *UsageClient) List(location string, options *UsageListOptions) ListUsagesResultPager {
	return &listUsagesResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, location, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *ListUsagesResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListUsagesResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *UsageClient) ListCreateRequest(ctx context.Context, location string, options *UsageListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *UsageClient) ListHandleResponse(resp *azcore.Response) (*ListUsagesResultResponse, error) {
	result := ListUsagesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListUsagesResult)
}

// ListHandleError handles the List error response.
func (client *UsageClient) ListHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
