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
)

// WebCategoriesOperations contains the methods for the WebCategories group.
type WebCategoriesOperations interface {
	// Get - Gets the specified Azure Web Category.
	Get(ctx context.Context, name string, options *WebCategoriesGetOptions) (*AzureWebCategoryResponse, error)
	// ListBySubscription - Gets all the Azure Web Categories in a subscription.
	ListBySubscription(options *WebCategoriesListBySubscriptionOptions) AzureWebCategoryListResultPager
}

// WebCategoriesClient implements the WebCategoriesOperations interface.
// Don't use this type directly, use NewWebCategoriesClient() instead.
type WebCategoriesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewWebCategoriesClient creates a new instance of WebCategoriesClient with the specified values.
func NewWebCategoriesClient(con *armcore.Connection, subscriptionID string) WebCategoriesOperations {
	return &WebCategoriesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *WebCategoriesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets the specified Azure Web Category.
func (client *WebCategoriesClient) Get(ctx context.Context, name string, options *WebCategoriesGetOptions) (*AzureWebCategoryResponse, error) {
	req, err := client.GetCreateRequest(ctx, name, options)
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
func (client *WebCategoriesClient) GetCreateRequest(ctx context.Context, name string, options *WebCategoriesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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

// GetHandleResponse handles the Get response.
func (client *WebCategoriesClient) GetHandleResponse(resp *azcore.Response) (*AzureWebCategoryResponse, error) {
	result := AzureWebCategoryResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureWebCategory)
}

// GetHandleError handles the Get error response.
func (client *WebCategoriesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListBySubscription - Gets all the Azure Web Categories in a subscription.
func (client *WebCategoriesClient) ListBySubscription(options *WebCategoriesListBySubscriptionOptions) AzureWebCategoryListResultPager {
	return &azureWebCategoryListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.ListBySubscriptionHandleResponse,
		errorer:   client.ListBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp *AzureWebCategoryListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureWebCategoryListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *WebCategoriesClient) ListBySubscriptionCreateRequest(ctx context.Context, options *WebCategoriesListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories"
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

// ListBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *WebCategoriesClient) ListBySubscriptionHandleResponse(resp *azcore.Response) (*AzureWebCategoryListResultResponse, error) {
	result := AzureWebCategoryListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AzureWebCategoryListResult)
}

// ListBySubscriptionHandleError handles the ListBySubscription error response.
func (client *WebCategoriesClient) ListBySubscriptionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}