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

// NetworkManagementClient contains the methods for the NetworkManagementClient group.
// Don't use this type directly, use NewNetworkManagementClient() instead.
type NetworkManagementClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkManagementClient creates a new instance of NetworkManagementClient with the specified values.
func NewNetworkManagementClient(con *armcore.Connection, subscriptionID string) NetworkManagementClient {
	return NetworkManagementClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client NetworkManagementClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CheckDNSNameAvailability - Checks whether a domain name in the cloudapp.azure.com zone is available for use.
func (client NetworkManagementClient) CheckDNSNameAvailability(ctx context.Context, location string, domainNameLabel string, options *NetworkManagementClientCheckDNSNameAvailabilityOptions) (*DNSNameAvailabilityResultResponse, error) {
	req, err := client.checkDnsNameAvailabilityCreateRequest(ctx, location, domainNameLabel, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.checkDnsNameAvailabilityHandleError(resp)
	}
	result, err := client.checkDnsNameAvailabilityHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// checkDnsNameAvailabilityCreateRequest creates the CheckDNSNameAvailability request.
func (client NetworkManagementClient) checkDnsNameAvailabilityCreateRequest(ctx context.Context, location string, domainNameLabel string, options *NetworkManagementClientCheckDNSNameAvailabilityOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/CheckDnsNameAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("domainNameLabel", domainNameLabel)
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// checkDnsNameAvailabilityHandleResponse handles the CheckDNSNameAvailability response.
func (client NetworkManagementClient) checkDnsNameAvailabilityHandleResponse(resp *azcore.Response) (*DNSNameAvailabilityResultResponse, error) {
	result := DNSNameAvailabilityResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DNSNameAvailabilityResult)
}

// checkDnsNameAvailabilityHandleError handles the CheckDNSNameAvailability error response.
func (client NetworkManagementClient) checkDnsNameAvailabilityHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDeleteBastionShareableLink - Deletes the Bastion Shareable Links for all the VMs specified in the request.
func (client NetworkManagementClient) BeginDeleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientDeleteBastionShareableLinkOptions) (*HTTPPollerResponse, error) {
	resp, err := client.DeleteBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.DeleteBastionShareableLink", "location", resp, client.deleteBastionShareableLinkHandleError)
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

// ResumeDeleteBastionShareableLink creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client NetworkManagementClient) ResumeDeleteBastionShareableLink(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.DeleteBastionShareableLink", token, client.deleteBastionShareableLinkHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// DeleteBastionShareableLink - Deletes the Bastion Shareable Links for all the VMs specified in the request.
func (client NetworkManagementClient) DeleteBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientDeleteBastionShareableLinkOptions) (*azcore.Response, error) {
	req, err := client.deleteBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.deleteBastionShareableLinkHandleError(resp)
	}
	return resp, nil
}

// deleteBastionShareableLinkCreateRequest creates the DeleteBastionShareableLink request.
func (client NetworkManagementClient) deleteBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientDeleteBastionShareableLinkOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/deleteShareableLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, req.MarshalAsJSON(bslRequest)
}

// deleteBastionShareableLinkHandleError handles the DeleteBastionShareableLink error response.
func (client NetworkManagementClient) deleteBastionShareableLinkHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DisconnectActiveSessions - Returns the list of currently active sessions on the Bastion.
func (client NetworkManagementClient) DisconnectActiveSessions(resourceGroupName string, bastionHostName string, sessionIds SessionIDs, options *NetworkManagementClientDisconnectActiveSessionsOptions) BastionSessionDeleteResultPager {
	return &bastionSessionDeleteResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.disconnectActiveSessionsCreateRequest(ctx, resourceGroupName, bastionHostName, sessionIds, options)
		},
		responder: client.disconnectActiveSessionsHandleResponse,
		errorer:   client.disconnectActiveSessionsHandleError,
		advancer: func(ctx context.Context, resp *BastionSessionDeleteResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.BastionSessionDeleteResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// disconnectActiveSessionsCreateRequest creates the DisconnectActiveSessions request.
func (client NetworkManagementClient) disconnectActiveSessionsCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, sessionIds SessionIDs, options *NetworkManagementClientDisconnectActiveSessionsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/disconnectActiveSessions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, req.MarshalAsJSON(sessionIds)
}

// disconnectActiveSessionsHandleResponse handles the DisconnectActiveSessions response.
func (client NetworkManagementClient) disconnectActiveSessionsHandleResponse(resp *azcore.Response) (*BastionSessionDeleteResultResponse, error) {
	result := BastionSessionDeleteResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionSessionDeleteResult)
}

// disconnectActiveSessionsHandleError handles the DisconnectActiveSessions error response.
func (client NetworkManagementClient) disconnectActiveSessionsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginGeneratevirtualwanvpnserverconfigurationvpnprofile - Generates a unique VPN profile for P2S clients for VirtualWan and associated VpnServerConfiguration
// combination in the specified resource group.
func (client NetworkManagementClient) BeginGeneratevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWanName string, vpnClientParams VirtualWanVpnProfileParameters, options *NetworkManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*VpnProfileResponsePollerResponse, error) {
	resp, err := client.Generatevirtualwanvpnserverconfigurationvpnprofile(ctx, resourceGroupName, virtualWanName, vpnClientParams, options)
	if err != nil {
		return nil, err
	}
	result := &VpnProfileResponsePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.Generatevirtualwanvpnserverconfigurationvpnprofile", "location", resp, client.generatevirtualwanvpnserverconfigurationvpnprofileHandleError)
	if err != nil {
		return nil, err
	}
	poller := &vpnProfileResponsePoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VpnProfileResponseResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeGeneratevirtualwanvpnserverconfigurationvpnprofile creates a new VpnProfileResponsePoller from the specified resume token.
// token - The value must come from a previous call to VpnProfileResponsePoller.ResumeToken().
func (client NetworkManagementClient) ResumeGeneratevirtualwanvpnserverconfigurationvpnprofile(token string) (VpnProfileResponsePoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.Generatevirtualwanvpnserverconfigurationvpnprofile", token, client.generatevirtualwanvpnserverconfigurationvpnprofileHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnProfileResponsePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Generatevirtualwanvpnserverconfigurationvpnprofile - Generates a unique VPN profile for P2S clients for VirtualWan and associated VpnServerConfiguration
// combination in the specified resource group.
func (client NetworkManagementClient) Generatevirtualwanvpnserverconfigurationvpnprofile(ctx context.Context, resourceGroupName string, virtualWanName string, vpnClientParams VirtualWanVpnProfileParameters, options *NetworkManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*azcore.Response, error) {
	req, err := client.generatevirtualwanvpnserverconfigurationvpnprofileCreateRequest(ctx, resourceGroupName, virtualWanName, vpnClientParams, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.generatevirtualwanvpnserverconfigurationvpnprofileHandleError(resp)
	}
	return resp, nil
}

// generatevirtualwanvpnserverconfigurationvpnprofileCreateRequest creates the Generatevirtualwanvpnserverconfigurationvpnprofile request.
func (client NetworkManagementClient) generatevirtualwanvpnserverconfigurationvpnprofileCreateRequest(ctx context.Context, resourceGroupName string, virtualWanName string, vpnClientParams VirtualWanVpnProfileParameters, options *NetworkManagementClientGeneratevirtualwanvpnserverconfigurationvpnprofileOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/GenerateVpnProfile"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWanName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnClientParams)
}

// generatevirtualwanvpnserverconfigurationvpnprofileHandleResponse handles the Generatevirtualwanvpnserverconfigurationvpnprofile response.
func (client NetworkManagementClient) generatevirtualwanvpnserverconfigurationvpnprofileHandleResponse(resp *azcore.Response) (*VpnProfileResponseResponse, error) {
	result := VpnProfileResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnProfileResponse)
}

// generatevirtualwanvpnserverconfigurationvpnprofileHandleError handles the Generatevirtualwanvpnserverconfigurationvpnprofile error response.
func (client NetworkManagementClient) generatevirtualwanvpnserverconfigurationvpnprofileHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginGetActiveSessions - Returns the list of currently active sessions on the Bastion.
func (client NetworkManagementClient) BeginGetActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, options *NetworkManagementClientGetActiveSessionsOptions) (*BastionActiveSessionListResultPagerPollerResponse, error) {
	resp, err := client.GetActiveSessions(ctx, resourceGroupName, bastionHostName, options)
	if err != nil {
		return nil, err
	}
	result := &BastionActiveSessionListResultPagerPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.GetActiveSessions", "location", resp, client.getActiveSessionsHandleError)
	if err != nil {
		return nil, err
	}
	poller := &bastionActiveSessionListResultPagerPoller{
		pt: pt,
		errHandler: func(resp *azcore.Response) error {
			if resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
				return nil
			}
			return client.getActiveSessionsHandleError(resp)
		},
		respHandler: func(resp *azcore.Response) (*BastionActiveSessionListResultResponse, error) {
			result := BastionActiveSessionListResultResponse{RawResponse: resp.Response}
			return &result, resp.UnmarshalAsJSON(&result.BastionActiveSessionListResult)
		},
		statusCodes: []int{http.StatusOK, http.StatusAccepted, http.StatusNoContent},
		pipeline:    client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (BastionActiveSessionListResultPager, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeGetActiveSessions creates a new BastionActiveSessionListResultPagerPoller from the specified resume token.
// token - The value must come from a previous call to BastionActiveSessionListResultPagerPoller.ResumeToken().
func (client NetworkManagementClient) ResumeGetActiveSessions(token string) (BastionActiveSessionListResultPagerPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.GetActiveSessions", token, client.getActiveSessionsHandleError)
	if err != nil {
		return nil, err
	}
	return &bastionActiveSessionListResultPagerPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// GetActiveSessions - Returns the list of currently active sessions on the Bastion.
func (client NetworkManagementClient) GetActiveSessions(ctx context.Context, resourceGroupName string, bastionHostName string, options *NetworkManagementClientGetActiveSessionsOptions) (*azcore.Response, error) {
	req, err := client.getActiveSessionsCreateRequest(ctx, resourceGroupName, bastionHostName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.getActiveSessionsHandleError(resp)
	}
	return resp, nil
}

// getActiveSessionsCreateRequest creates the GetActiveSessions request.
func (client NetworkManagementClient) getActiveSessionsCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, options *NetworkManagementClientGetActiveSessionsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getActiveSessions"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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

// getActiveSessionsHandleResponse handles the GetActiveSessions response.
func (client NetworkManagementClient) getActiveSessionsHandleResponse(resp *azcore.Response) (*BastionActiveSessionListResultResponse, error) {
	result := BastionActiveSessionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionActiveSessionListResult)
}

// getActiveSessionsHandleError handles the GetActiveSessions error response.
func (client NetworkManagementClient) getActiveSessionsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetBastionShareableLink - Return the Bastion Shareable Links for all the VMs specified in the request.
func (client NetworkManagementClient) GetBastionShareableLink(resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientGetBastionShareableLinkOptions) BastionShareableLinkListResultPager {
	return &bastionShareableLinkListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
		},
		responder: client.getBastionShareableLinkHandleResponse,
		errorer:   client.getBastionShareableLinkHandleError,
		advancer: func(ctx context.Context, resp *BastionShareableLinkListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.BastionShareableLinkListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getBastionShareableLinkCreateRequest creates the GetBastionShareableLink request.
func (client NetworkManagementClient) getBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientGetBastionShareableLinkOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/getShareableLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, req.MarshalAsJSON(bslRequest)
}

// getBastionShareableLinkHandleResponse handles the GetBastionShareableLink response.
func (client NetworkManagementClient) getBastionShareableLinkHandleResponse(resp *azcore.Response) (*BastionShareableLinkListResultResponse, error) {
	result := BastionShareableLinkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionShareableLinkListResult)
}

// getBastionShareableLinkHandleError handles the GetBastionShareableLink error response.
func (client NetworkManagementClient) getBastionShareableLinkHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginPutBastionShareableLink - Creates a Bastion Shareable Links for all the VMs specified in the request.
func (client NetworkManagementClient) BeginPutBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientPutBastionShareableLinkOptions) (*BastionShareableLinkListResultPagerPollerResponse, error) {
	resp, err := client.PutBastionShareableLink(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	result := &BastionShareableLinkListResultPagerPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("NetworkManagementClient.PutBastionShareableLink", "location", resp, client.putBastionShareableLinkHandleError)
	if err != nil {
		return nil, err
	}
	poller := &bastionShareableLinkListResultPagerPoller{
		pt: pt,
		errHandler: func(resp *azcore.Response) error {
			if resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
				return nil
			}
			return client.putBastionShareableLinkHandleError(resp)
		},
		respHandler: func(resp *azcore.Response) (*BastionShareableLinkListResultResponse, error) {
			result := BastionShareableLinkListResultResponse{RawResponse: resp.Response}
			return &result, resp.UnmarshalAsJSON(&result.BastionShareableLinkListResult)
		},
		statusCodes: []int{http.StatusOK, http.StatusAccepted, http.StatusNoContent},
		pipeline:    client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (BastionShareableLinkListResultPager, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePutBastionShareableLink creates a new BastionShareableLinkListResultPagerPoller from the specified resume token.
// token - The value must come from a previous call to BastionShareableLinkListResultPagerPoller.ResumeToken().
func (client NetworkManagementClient) ResumePutBastionShareableLink(token string) (BastionShareableLinkListResultPagerPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("NetworkManagementClient.PutBastionShareableLink", token, client.putBastionShareableLinkHandleError)
	if err != nil {
		return nil, err
	}
	return &bastionShareableLinkListResultPagerPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// PutBastionShareableLink - Creates a Bastion Shareable Links for all the VMs specified in the request.
func (client NetworkManagementClient) PutBastionShareableLink(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientPutBastionShareableLinkOptions) (*azcore.Response, error) {
	req, err := client.putBastionShareableLinkCreateRequest(ctx, resourceGroupName, bastionHostName, bslRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.putBastionShareableLinkHandleError(resp)
	}
	return resp, nil
}

// putBastionShareableLinkCreateRequest creates the PutBastionShareableLink request.
func (client NetworkManagementClient) putBastionShareableLinkCreateRequest(ctx context.Context, resourceGroupName string, bastionHostName string, bslRequest BastionShareableLinkListRequest, options *NetworkManagementClientPutBastionShareableLinkOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/bastionHosts/{bastionHostName}/createShareableLinks"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{bastionHostName}", url.PathEscape(bastionHostName))
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
	return req, req.MarshalAsJSON(bslRequest)
}

// putBastionShareableLinkHandleResponse handles the PutBastionShareableLink response.
func (client NetworkManagementClient) putBastionShareableLinkHandleResponse(resp *azcore.Response) (*BastionShareableLinkListResultResponse, error) {
	result := BastionShareableLinkListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BastionShareableLinkListResult)
}

// putBastionShareableLinkHandleError handles the PutBastionShareableLink error response.
func (client NetworkManagementClient) putBastionShareableLinkHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SupportedSecurityProviders - Gives the supported security providers for the virtual wan.
func (client NetworkManagementClient) SupportedSecurityProviders(ctx context.Context, resourceGroupName string, virtualWanName string, options *NetworkManagementClientSupportedSecurityProvidersOptions) (*VirtualWanSecurityProvidersResponse, error) {
	req, err := client.supportedSecurityProvidersCreateRequest(ctx, resourceGroupName, virtualWanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.supportedSecurityProvidersHandleError(resp)
	}
	result, err := client.supportedSecurityProvidersHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// supportedSecurityProvidersCreateRequest creates the SupportedSecurityProviders request.
func (client NetworkManagementClient) supportedSecurityProvidersCreateRequest(ctx context.Context, resourceGroupName string, virtualWanName string, options *NetworkManagementClientSupportedSecurityProvidersOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/supportedSecurityProviders"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWanName))
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

// supportedSecurityProvidersHandleResponse handles the SupportedSecurityProviders response.
func (client NetworkManagementClient) supportedSecurityProvidersHandleResponse(resp *azcore.Response) (*VirtualWanSecurityProvidersResponse, error) {
	result := VirtualWanSecurityProvidersResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualWanSecurityProviders)
}

// supportedSecurityProvidersHandleError handles the SupportedSecurityProviders error response.
func (client NetworkManagementClient) supportedSecurityProvidersHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
