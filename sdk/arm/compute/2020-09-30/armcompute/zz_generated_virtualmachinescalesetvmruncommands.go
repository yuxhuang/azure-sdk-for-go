// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualMachineScaleSetVMRunCommandsOperations contains the methods for the VirtualMachineScaleSetVMRunCommands group.
type VirtualMachineScaleSetVMRunCommandsOperations interface {
	// BeginCreateOrUpdate - The operation to create or update the VMSS VM run command.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsCreateOrUpdateOptions) (*VirtualMachineRunCommandPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VirtualMachineRunCommandPoller, error)
	// BeginDelete - The operation to delete the VMSS VM run command.
	BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - The operation to get the VMSS VM run command.
	Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsGetOptions) (*VirtualMachineRunCommandResponse, error)
	// List - The operation to get all run commands of an instance in Virtual Machine Scaleset.
	List(resourceGroupName string, vmScaleSetName string, instanceId string, options *VirtualMachineScaleSetVMRunCommandsListOptions) VirtualMachineRunCommandsListResultPager
	// BeginUpdate - The operation to update the VMSS VM run command.
	BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsUpdateOptions) (*VirtualMachineRunCommandPollerResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (VirtualMachineRunCommandPoller, error)
}

// VirtualMachineScaleSetVMRunCommandsClient implements the VirtualMachineScaleSetVMRunCommandsOperations interface.
// Don't use this type directly, use NewVirtualMachineScaleSetVMRunCommandsClient() instead.
type VirtualMachineScaleSetVMRunCommandsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineScaleSetVMRunCommandsClient creates a new instance of VirtualMachineScaleSetVMRunCommandsClient with the specified values.
func NewVirtualMachineScaleSetVMRunCommandsClient(con *armcore.Connection, subscriptionID string) VirtualMachineScaleSetVMRunCommandsOperations {
	return &VirtualMachineScaleSetVMRunCommandsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualMachineScaleSetVMRunCommandsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsCreateOrUpdateOptions) (*VirtualMachineRunCommandPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, vmScaleSetName, instanceId, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineRunCommandPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.CreateOrUpdate", "", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineRunCommandPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineRunCommandResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineScaleSetVMRunCommandsClient) ResumeCreateOrUpdate(token string) (VirtualMachineRunCommandPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMRunCommandsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineRunCommandPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - The operation to create or update the VMSS VM run command.
func (client *VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, runCommandName, runCommand, options)
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
func (client *VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommand, options *VirtualMachineScaleSetVMRunCommandsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, req.MarshalAsJSON(runCommand)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*VirtualMachineRunCommandResponse, error) {
	result := VirtualMachineRunCommandResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineRunCommand)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginDelete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, vmScaleSetName, instanceId, runCommandName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.Delete", "", resp, client.DeleteHandleError)
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

func (client *VirtualMachineScaleSetVMRunCommandsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMRunCommandsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - The operation to delete the VMSS VM run command.
func (client *VirtualMachineScaleSetVMRunCommandsClient) Delete(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, runCommandName, options)
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
func (client *VirtualMachineScaleSetVMRunCommandsClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - The operation to get the VMSS VM run command.
func (client *VirtualMachineScaleSetVMRunCommandsClient) Get(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsGetOptions) (*VirtualMachineRunCommandResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, runCommandName, options)
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
func (client *VirtualMachineScaleSetVMRunCommandsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, options *VirtualMachineScaleSetVMRunCommandsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) GetHandleResponse(resp *azcore.Response) (*VirtualMachineRunCommandResponse, error) {
	result := VirtualMachineRunCommandResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineRunCommand)
}

// GetHandleError handles the Get error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - The operation to get all run commands of an instance in Virtual Machine Scaleset.
func (client *VirtualMachineScaleSetVMRunCommandsClient) List(resourceGroupName string, vmScaleSetName string, instanceId string, options *VirtualMachineScaleSetVMRunCommandsListOptions) VirtualMachineRunCommandsListResultPager {
	return &virtualMachineRunCommandsListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *VirtualMachineRunCommandsListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualMachineRunCommandsListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// ListCreateRequest creates the List request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) ListCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, options *VirtualMachineScaleSetVMRunCommandsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) ListHandleResponse(resp *azcore.Response) (*VirtualMachineRunCommandsListResultResponse, error) {
	result := VirtualMachineRunCommandsListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineRunCommandsListResult)
}

// ListHandleError handles the List error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *VirtualMachineScaleSetVMRunCommandsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsUpdateOptions) (*VirtualMachineRunCommandPollerResponse, error) {
	resp, err := client.Update(ctx, resourceGroupName, vmScaleSetName, instanceId, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualMachineRunCommandPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetVMRunCommandsClient.Update", "", resp, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualMachineRunCommandPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualMachineRunCommandResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *VirtualMachineScaleSetVMRunCommandsClient) ResumeUpdate(token string) (VirtualMachineRunCommandPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetVMRunCommandsClient.Update", token, client.UpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualMachineRunCommandPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Update - The operation to update the VMSS VM run command.
func (client *VirtualMachineScaleSetVMRunCommandsClient) Update(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsUpdateOptions) (*azcore.Response, error) {
	req, err := client.UpdateCreateRequest(ctx, resourceGroupName, vmScaleSetName, instanceId, runCommandName, runCommand, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateHandleError(resp)
	}
	return resp, nil
}

// UpdateCreateRequest creates the Update request.
func (client *VirtualMachineScaleSetVMRunCommandsClient) UpdateCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, instanceId string, runCommandName string, runCommand VirtualMachineRunCommandUpdate, options *VirtualMachineScaleSetVMRunCommandsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/virtualMachines/{instanceId}/runCommands/{runCommandName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceId))
	urlPath = strings.ReplaceAll(urlPath, "{runCommandName}", url.PathEscape(runCommandName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json, text/json")
	return req, req.MarshalAsJSON(runCommand)
}

// UpdateHandleResponse handles the Update response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) UpdateHandleResponse(resp *azcore.Response) (*VirtualMachineRunCommandResponse, error) {
	result := VirtualMachineRunCommandResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualMachineRunCommand)
}

// UpdateHandleError handles the Update error response.
func (client *VirtualMachineScaleSetVMRunCommandsClient) UpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
