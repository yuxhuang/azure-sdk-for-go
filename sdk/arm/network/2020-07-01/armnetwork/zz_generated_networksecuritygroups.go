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

// NetworkSecurityGroupsClient contains the methods for the NetworkSecurityGroups group.
// Don't use this type directly, use NewNetworkSecurityGroupsClient() instead.
type NetworkSecurityGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkSecurityGroupsClient creates a new instance of NetworkSecurityGroupsClient with the specified values.
func NewNetworkSecurityGroupsClient(con *armcore.Connection, subscriptionID string) NetworkSecurityGroupsClient {
	return NetworkSecurityGroupsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client NetworkSecurityGroupsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates a network security group in the specified resource group.
func (client NetworkSecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters NetworkSecurityGroup, options *NetworkSecurityGroupsCreateOrUpdateOptions) (*NetworkSecurityGroupPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &NetworkSecurityGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkSecurityGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &networkSecurityGroupPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*NetworkSecurityGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new NetworkSecurityGroupPoller from the specified resume token.
// token - The value must come from a previous call to NetworkSecurityGroupPoller.ResumeToken().
func (client NetworkSecurityGroupsClient) ResumeCreateOrUpdate(token string) (NetworkSecurityGroupPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkSecurityGroupsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &networkSecurityGroupPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a network security group in the specified resource group.
func (client NetworkSecurityGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters NetworkSecurityGroup, options *NetworkSecurityGroupsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client NetworkSecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters NetworkSecurityGroup, options *NetworkSecurityGroupsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client NetworkSecurityGroupsClient) createOrUpdateHandleResponse(resp *azcore.Response) (*NetworkSecurityGroupResponse, error) {
	result := NetworkSecurityGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkSecurityGroup)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client NetworkSecurityGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified network security group.
func (client NetworkSecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkSecurityGroupsClient.Delete", "location", resp, client.deleteHandleError)
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

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client NetworkSecurityGroupsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkSecurityGroupsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified network security group.
func (client NetworkSecurityGroupsClient) Delete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client NetworkSecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
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

// deleteHandleError handles the Delete error response.
func (client NetworkSecurityGroupsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified network security group.
func (client NetworkSecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsGetOptions) (*NetworkSecurityGroupResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client NetworkSecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client NetworkSecurityGroupsClient) getHandleResponse(resp *azcore.Response) (*NetworkSecurityGroupResponse, error) {
	result := NetworkSecurityGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkSecurityGroup)
}

// getHandleError handles the Get error response.
func (client NetworkSecurityGroupsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all network security groups in a resource group.
func (client NetworkSecurityGroupsClient) List(resourceGroupName string, options *NetworkSecurityGroupsListOptions) NetworkSecurityGroupListResultPager {
	return &networkSecurityGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp *NetworkSecurityGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkSecurityGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client NetworkSecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkSecurityGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listHandleResponse handles the List response.
func (client NetworkSecurityGroupsClient) listHandleResponse(resp *azcore.Response) (*NetworkSecurityGroupListResultResponse, error) {
	result := NetworkSecurityGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkSecurityGroupListResult)
}

// listHandleError handles the List error response.
func (client NetworkSecurityGroupsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all network security groups in a subscription.
func (client NetworkSecurityGroupsClient) ListAll(options *NetworkSecurityGroupsListAllOptions) NetworkSecurityGroupListResultPager {
	return &networkSecurityGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp *NetworkSecurityGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkSecurityGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client NetworkSecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *NetworkSecurityGroupsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups"
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

// listAllHandleResponse handles the ListAll response.
func (client NetworkSecurityGroupsClient) listAllHandleResponse(resp *azcore.Response) (*NetworkSecurityGroupListResultResponse, error) {
	result := NetworkSecurityGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkSecurityGroupListResult)
}

// listAllHandleError handles the ListAll error response.
func (client NetworkSecurityGroupsClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates a network security group tags.
func (client NetworkSecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *NetworkSecurityGroupsUpdateTagsOptions) (*NetworkSecurityGroupResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client NetworkSecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *NetworkSecurityGroupsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// updateTagsHandleResponse handles the UpdateTags response.
func (client NetworkSecurityGroupsClient) updateTagsHandleResponse(resp *azcore.Response) (*NetworkSecurityGroupResponse, error) {
	result := NetworkSecurityGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkSecurityGroup)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client NetworkSecurityGroupsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
