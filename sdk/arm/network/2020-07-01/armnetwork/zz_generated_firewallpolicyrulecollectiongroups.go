// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// FirewallPolicyRuleCollectionGroupsOperations contains the methods for the FirewallPolicyRuleCollectionGroups group.
type FirewallPolicyRuleCollectionGroupsOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsCreateOrUpdateOptions) (*FirewallPolicyRuleCollectionGroupPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FirewallPolicyRuleCollectionGroupPoller, error)
	// BeginDelete - Deletes the specified FirewallPolicyRuleCollectionGroup.
	BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified FirewallPolicyRuleCollectionGroup.
	Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsGetOptions) (*FirewallPolicyRuleCollectionGroupResponse, error)
	// List - Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
	List(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsListOptions) FirewallPolicyRuleCollectionGroupListResultPager
}

// FirewallPolicyRuleCollectionGroupsClient implements the FirewallPolicyRuleCollectionGroupsOperations interface.
// Don't use this type directly, use NewFirewallPolicyRuleCollectionGroupsClient() instead.
type FirewallPolicyRuleCollectionGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFirewallPolicyRuleCollectionGroupsClient creates a new instance of FirewallPolicyRuleCollectionGroupsClient with the specified values.
func NewFirewallPolicyRuleCollectionGroupsClient(con *armcore.Connection, subscriptionID string) FirewallPolicyRuleCollectionGroupsOperations {
	return &FirewallPolicyRuleCollectionGroupsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *FirewallPolicyRuleCollectionGroupsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *FirewallPolicyRuleCollectionGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsCreateOrUpdateOptions) (*FirewallPolicyRuleCollectionGroupPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &FirewallPolicyRuleCollectionGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPolicyRuleCollectionGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &firewallPolicyRuleCollectionGroupPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FirewallPolicyRuleCollectionGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FirewallPolicyRuleCollectionGroupsClient) ResumeCreateOrUpdate(token string) (FirewallPolicyRuleCollectionGroupPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPolicyRuleCollectionGroupsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &firewallPolicyRuleCollectionGroupPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates the specified FirewallPolicyRuleCollectionGroup.
func (client *FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup, options *FirewallPolicyRuleCollectionGroupsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleCollectionGroupResponse, error) {
	result := FirewallPolicyRuleCollectionGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleCollectionGroup)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *FirewallPolicyRuleCollectionGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("FirewallPolicyRuleCollectionGroupsClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FirewallPolicyRuleCollectionGroupsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FirewallPolicyRuleCollectionGroupsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified FirewallPolicyRuleCollectionGroup.
func (client *FirewallPolicyRuleCollectionGroupsClient) Delete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *FirewallPolicyRuleCollectionGroupsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *FirewallPolicyRuleCollectionGroupsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified FirewallPolicyRuleCollectionGroup.
func (client *FirewallPolicyRuleCollectionGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsGetOptions) (*FirewallPolicyRuleCollectionGroupResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *FirewallPolicyRuleCollectionGroupsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *FirewallPolicyRuleCollectionGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{ruleCollectionGroupName}", url.PathEscape(ruleCollectionGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *FirewallPolicyRuleCollectionGroupsClient) GetHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleCollectionGroupResponse, error) {
	result := FirewallPolicyRuleCollectionGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleCollectionGroup)
}

// GetHandleError handles the Get error response.
func (client *FirewallPolicyRuleCollectionGroupsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
func (client *FirewallPolicyRuleCollectionGroupsClient) List(resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsListOptions) FirewallPolicyRuleCollectionGroupListResultPager {
	return &firewallPolicyRuleCollectionGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *FirewallPolicyRuleCollectionGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FirewallPolicyRuleCollectionGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *FirewallPolicyRuleCollectionGroupsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPolicyRuleCollectionGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *FirewallPolicyRuleCollectionGroupsClient) ListHandleResponse(resp *azcore.Response) (*FirewallPolicyRuleCollectionGroupListResultResponse, error) {
	result := FirewallPolicyRuleCollectionGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FirewallPolicyRuleCollectionGroupListResult)
}

// ListHandleError handles the List error response.
func (client *FirewallPolicyRuleCollectionGroupsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}