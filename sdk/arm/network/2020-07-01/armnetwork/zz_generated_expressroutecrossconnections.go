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

// ExpressRouteCrossConnectionsClient contains the methods for the ExpressRouteCrossConnections group.
// Don't use this type directly, use NewExpressRouteCrossConnectionsClient() instead.
type ExpressRouteCrossConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCrossConnectionsClient creates a new instance of ExpressRouteCrossConnectionsClient with the specified values.
func NewExpressRouteCrossConnectionsClient(con *armcore.Connection, subscriptionID string) ExpressRouteCrossConnectionsClient {
	return ExpressRouteCrossConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client ExpressRouteCrossConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Update the specified ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsCreateOrUpdateOptions) (*ExpressRouteCrossConnectionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCrossConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteCrossConnectionPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCrossConnectionPoller.ResumeToken().
func (client ExpressRouteCrossConnectionsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCrossConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Update the specified ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, crossConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client ExpressRouteCrossConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, parameters ExpressRouteCrossConnection, options *ExpressRouteCrossConnectionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
func (client ExpressRouteCrossConnectionsClient) createOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client ExpressRouteCrossConnectionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets details about the specified ExpressRouteCrossConnection.
func (client ExpressRouteCrossConnectionsClient) Get(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsGetOptions) (*ExpressRouteCrossConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, crossConnectionName, options)
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
func (client ExpressRouteCrossConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, options *ExpressRouteCrossConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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

// getHandleResponse handles the Get response.
func (client ExpressRouteCrossConnectionsClient) getHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// getHandleError handles the Get error response.
func (client ExpressRouteCrossConnectionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Retrieves all the ExpressRouteCrossConnections in a subscription.
func (client ExpressRouteCrossConnectionsClient) List(options *ExpressRouteCrossConnectionsListOptions) ExpressRouteCrossConnectionListResultPager {
	return &expressRouteCrossConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client ExpressRouteCrossConnectionsClient) listCreateRequest(ctx context.Context, options *ExpressRouteCrossConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCrossConnections"
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
func (client ExpressRouteCrossConnectionsClient) listHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionListResultResponse, error) {
	result := ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionListResult)
}

// listHandleError handles the List error response.
func (client ExpressRouteCrossConnectionsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
func (client ExpressRouteCrossConnectionsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListArpTableOptions) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	resp, err := client.ListArpTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsArpTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListArpTable", "location", resp, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListArpTable creates a new ExpressRouteCircuitsArpTableListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitsArpTableListResultPoller.ResumeToken().
func (client ExpressRouteCrossConnectionsClient) ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListArpTable", token, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route cross connection in a resource group.
func (client ExpressRouteCrossConnectionsClient) ListArpTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListArpTableOptions) (*azcore.Response, error) {
	req, err := client.listArpTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listArpTableHandleError(resp)
	}
	return resp, nil
}

// listArpTableCreateRequest creates the ListArpTable request.
func (client ExpressRouteCrossConnectionsClient) listArpTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListArpTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/arpTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// listArpTableHandleResponse handles the ListArpTable response.
func (client ExpressRouteCrossConnectionsClient) listArpTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
	result := ExpressRouteCircuitsArpTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsArpTableListResult)
}

// listArpTableHandleError handles the ListArpTable error response.
func (client ExpressRouteCrossConnectionsClient) listArpTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Retrieves all the ExpressRouteCrossConnections in a resource group.
func (client ExpressRouteCrossConnectionsClient) ListByResourceGroup(resourceGroupName string, options *ExpressRouteCrossConnectionsListByResourceGroupOptions) ExpressRouteCrossConnectionListResultPager {
	return &expressRouteCrossConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCrossConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCrossConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client ExpressRouteCrossConnectionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCrossConnectionsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections"
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client ExpressRouteCrossConnectionsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionListResultResponse, error) {
	result := ExpressRouteCrossConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionListResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client ExpressRouteCrossConnectionsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
func (client ExpressRouteCrossConnectionsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableOptions) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	resp, err := client.ListRoutesTable(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTable", "location", resp, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListRoutesTable creates a new ExpressRouteCircuitsRoutesTableListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitsRoutesTableListResultPoller.ResumeToken().
func (client ExpressRouteCrossConnectionsClient) ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListRoutesTable", token, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route cross connection in a resource group.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableOptions) (*azcore.Response, error) {
	req, err := client.listRoutesTableCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableHandleError(resp)
	}
	return resp, nil
}

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client ExpressRouteCrossConnectionsClient) listRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// listRoutesTableHandleResponse handles the ListRoutesTable response.
func (client ExpressRouteCrossConnectionsClient) listRoutesTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
	result := ExpressRouteCircuitsRoutesTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsRoutesTableListResult)
}

// listRoutesTableHandleError handles the ListRoutesTable error response.
func (client ExpressRouteCrossConnectionsClient) listRoutesTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
func (client ExpressRouteCrossConnectionsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableSummaryOptions) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse, error) {
	resp, err := client.ListRoutesTableSummary(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCrossConnectionsRoutesTableSummaryListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCrossConnectionsClient.ListRoutesTableSummary", "location", resp, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListRoutesTableSummary creates a new ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller.ResumeToken().
func (client ExpressRouteCrossConnectionsClient) ResumeListRoutesTableSummary(token string) (ExpressRouteCrossConnectionsRoutesTableSummaryListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCrossConnectionsClient.ListRoutesTableSummary", token, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCrossConnectionsRoutesTableSummaryListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListRoutesTableSummary - Gets the route table summary associated with the express route cross connection in a resource group.
func (client ExpressRouteCrossConnectionsClient) ListRoutesTableSummary(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableSummaryOptions) (*azcore.Response, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(ctx, resourceGroupName, crossConnectionName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableSummaryHandleError(resp)
	}
	return resp, nil
}

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client ExpressRouteCrossConnectionsClient) listRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, peeringName string, devicePath string, options *ExpressRouteCrossConnectionsListRoutesTableSummaryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// listRoutesTableSummaryHandleResponse handles the ListRoutesTableSummary response.
func (client ExpressRouteCrossConnectionsClient) listRoutesTableSummaryHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse, error) {
	result := ExpressRouteCrossConnectionsRoutesTableSummaryListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnectionsRoutesTableSummaryListResult)
}

// listRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client ExpressRouteCrossConnectionsClient) listRoutesTableSummaryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates an express route cross connection tags.
func (client ExpressRouteCrossConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsUpdateTagsOptions) (*ExpressRouteCrossConnectionResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, crossConnectionName, crossConnectionParameters, options)
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
func (client ExpressRouteCrossConnectionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, crossConnectionName string, crossConnectionParameters TagsObject, options *ExpressRouteCrossConnectionsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCrossConnections/{crossConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{crossConnectionName}", url.PathEscape(crossConnectionName))
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
	return req, req.MarshalAsJSON(crossConnectionParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client ExpressRouteCrossConnectionsClient) updateTagsHandleResponse(resp *azcore.Response) (*ExpressRouteCrossConnectionResponse, error) {
	result := ExpressRouteCrossConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCrossConnection)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client ExpressRouteCrossConnectionsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
