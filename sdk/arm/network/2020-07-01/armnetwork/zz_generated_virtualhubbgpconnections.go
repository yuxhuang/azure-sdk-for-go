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

// VirtualHubBgpConnectionsClient contains the methods for the VirtualHubBgpConnections group.
// Don't use this type directly, use NewVirtualHubBgpConnectionsClient() instead.
type VirtualHubBgpConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualHubBgpConnectionsClient creates a new instance of VirtualHubBgpConnectionsClient with the specified values.
func NewVirtualHubBgpConnectionsClient(con *armcore.Connection, subscriptionID string) VirtualHubBgpConnectionsClient {
	return VirtualHubBgpConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualHubBgpConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// List - Retrieves the details of all VirtualHubBgpConnections.
func (client VirtualHubBgpConnectionsClient) List(resourceGroupName string, virtualHubName string, options *VirtualHubBgpConnectionsListOptions) ListVirtualHubBgpConnectionResultsPager {
	return &listVirtualHubBgpConnectionResultsPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp *ListVirtualHubBgpConnectionResultsResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVirtualHubBgpConnectionResults.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client VirtualHubBgpConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubBgpConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
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
func (client VirtualHubBgpConnectionsClient) listHandleResponse(resp *azcore.Response) (*ListVirtualHubBgpConnectionResultsResponse, error) {
	result := ListVirtualHubBgpConnectionResultsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVirtualHubBgpConnectionResults)
}

// listHandleError handles the List error response.
func (client VirtualHubBgpConnectionsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListAdvertisedRoutes - Retrieves a list of routes the virtual hub bgp connection is advertising to the specified peer.
func (client VirtualHubBgpConnectionsClient) BeginListAdvertisedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsListAdvertisedRoutesOptions) (*PeerRouteListPollerResponse, error) {
	resp, err := client.ListAdvertisedRoutes(ctx, resourceGroupName, hubName, connectionName, options)
	if err != nil {
		return nil, err
	}
	result := &PeerRouteListPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualHubBgpConnectionsClient.ListAdvertisedRoutes", "location", resp, client.listAdvertisedRoutesHandleError)
	if err != nil {
		return nil, err
	}
	poller := &peerRouteListPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PeerRouteListResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListAdvertisedRoutes creates a new PeerRouteListPoller from the specified resume token.
// token - The value must come from a previous call to PeerRouteListPoller.ResumeToken().
func (client VirtualHubBgpConnectionsClient) ResumeListAdvertisedRoutes(token string) (PeerRouteListPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualHubBgpConnectionsClient.ListAdvertisedRoutes", token, client.listAdvertisedRoutesHandleError)
	if err != nil {
		return nil, err
	}
	return &peerRouteListPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListAdvertisedRoutes - Retrieves a list of routes the virtual hub bgp connection is advertising to the specified peer.
func (client VirtualHubBgpConnectionsClient) ListAdvertisedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsListAdvertisedRoutesOptions) (*azcore.Response, error) {
	req, err := client.listAdvertisedRoutesCreateRequest(ctx, resourceGroupName, hubName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listAdvertisedRoutesHandleError(resp)
	}
	return resp, nil
}

// listAdvertisedRoutesCreateRequest creates the ListAdvertisedRoutes request.
func (client VirtualHubBgpConnectionsClient) listAdvertisedRoutesCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsListAdvertisedRoutesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{hubName}/bgpConnections/{connectionName}/advertisedRoutes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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

// listAdvertisedRoutesHandleResponse handles the ListAdvertisedRoutes response.
func (client VirtualHubBgpConnectionsClient) listAdvertisedRoutesHandleResponse(resp *azcore.Response) (*PeerRouteListResponse, error) {
	result := PeerRouteListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PeerRouteList)
}

// listAdvertisedRoutesHandleError handles the ListAdvertisedRoutes error response.
func (client VirtualHubBgpConnectionsClient) listAdvertisedRoutesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListLearnedRoutes - Retrieves a list of routes the virtual hub bgp connection has learned.
func (client VirtualHubBgpConnectionsClient) BeginListLearnedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsListLearnedRoutesOptions) (*PeerRouteListPollerResponse, error) {
	resp, err := client.ListLearnedRoutes(ctx, resourceGroupName, hubName, connectionName, options)
	if err != nil {
		return nil, err
	}
	result := &PeerRouteListPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualHubBgpConnectionsClient.ListLearnedRoutes", "location", resp, client.listLearnedRoutesHandleError)
	if err != nil {
		return nil, err
	}
	poller := &peerRouteListPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PeerRouteListResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListLearnedRoutes creates a new PeerRouteListPoller from the specified resume token.
// token - The value must come from a previous call to PeerRouteListPoller.ResumeToken().
func (client VirtualHubBgpConnectionsClient) ResumeListLearnedRoutes(token string) (PeerRouteListPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualHubBgpConnectionsClient.ListLearnedRoutes", token, client.listLearnedRoutesHandleError)
	if err != nil {
		return nil, err
	}
	return &peerRouteListPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListLearnedRoutes - Retrieves a list of routes the virtual hub bgp connection has learned.
func (client VirtualHubBgpConnectionsClient) ListLearnedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsListLearnedRoutesOptions) (*azcore.Response, error) {
	req, err := client.listLearnedRoutesCreateRequest(ctx, resourceGroupName, hubName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listLearnedRoutesHandleError(resp)
	}
	return resp, nil
}

// listLearnedRoutesCreateRequest creates the ListLearnedRoutes request.
func (client VirtualHubBgpConnectionsClient) listLearnedRoutesCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsListLearnedRoutesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{hubName}/bgpConnections/{connectionName}/learnedRoutes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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

// listLearnedRoutesHandleResponse handles the ListLearnedRoutes response.
func (client VirtualHubBgpConnectionsClient) listLearnedRoutesHandleResponse(resp *azcore.Response) (*PeerRouteListResponse, error) {
	result := PeerRouteListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PeerRouteList)
}

// listLearnedRoutesHandleError handles the ListLearnedRoutes error response.
func (client VirtualHubBgpConnectionsClient) listLearnedRoutesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
