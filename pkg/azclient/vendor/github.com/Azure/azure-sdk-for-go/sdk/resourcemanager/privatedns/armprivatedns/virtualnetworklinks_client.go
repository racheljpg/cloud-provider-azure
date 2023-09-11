//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armprivatedns

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualNetworkLinksClient contains the methods for the VirtualNetworkLinks group.
// Don't use this type directly, use NewVirtualNetworkLinksClient() instead.
type VirtualNetworkLinksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworkLinksClient creates a new instance of VirtualNetworkLinksClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworkLinksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworkLinksClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualNetworkLinksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworkLinksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - privateZoneName - The name of the Private DNS zone (without a terminating dot).
//   - virtualNetworkLinkName - The name of the virtual network link.
//   - parameters - Parameters supplied to the CreateOrUpdate operation.
//   - options - VirtualNetworkLinksClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginCreateOrUpdate
//     method.
func (client *VirtualNetworkLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworkLinksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualNetworkLinksClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkLinksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
func (client *VirtualNetworkLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualNetworkLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a virtual network link to the specified Private DNS zone. WARNING: In case of a registration virtual
// network, all auto-registered DNS records in the zone for the virtual network will also be
// deleted. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - privateZoneName - The name of the Private DNS zone (without a terminating dot).
//   - virtualNetworkLinkName - The name of the virtual network link.
//   - options - VirtualNetworkLinksClientBeginDeleteOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginDelete
//     method.
func (client *VirtualNetworkLinksClient) BeginDelete(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworkLinksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualNetworkLinksClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkLinksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a virtual network link to the specified Private DNS zone. WARNING: In case of a registration virtual network,
// all auto-registered DNS records in the zone for the virtual network will also be
// deleted. This operation cannot be undone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
func (client *VirtualNetworkLinksClient) deleteOperation(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualNetworkLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a virtual network link to the specified Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - privateZoneName - The name of the Private DNS zone (without a terminating dot).
//   - virtualNetworkLinkName - The name of the virtual network link.
//   - options - VirtualNetworkLinksClientGetOptions contains the optional parameters for the VirtualNetworkLinksClient.Get method.
func (client *VirtualNetworkLinksClient) Get(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientGetOptions) (VirtualNetworkLinksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, options)
	if err != nil {
		return VirtualNetworkLinksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworkLinksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualNetworkLinksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworkLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, options *VirtualNetworkLinksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworkLinksClient) getHandleResponse(resp *http.Response) (VirtualNetworkLinksClientGetResponse, error) {
	result := VirtualNetworkLinksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkLink); err != nil {
		return VirtualNetworkLinksClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the virtual network links to the specified Private DNS zone.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - privateZoneName - The name of the Private DNS zone (without a terminating dot).
//   - options - VirtualNetworkLinksClientListOptions contains the optional parameters for the VirtualNetworkLinksClient.NewListPager
//     method.
func (client *VirtualNetworkLinksClient) NewListPager(resourceGroupName string, privateZoneName string, options *VirtualNetworkLinksClientListOptions) *runtime.Pager[VirtualNetworkLinksClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworkLinksClientListResponse]{
		More: func(page VirtualNetworkLinksClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworkLinksClientListResponse) (VirtualNetworkLinksClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, privateZoneName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworkLinksClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworkLinksClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworkLinksClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualNetworkLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, options *VirtualNetworkLinksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualNetworkLinksClient) listHandleResponse(resp *http.Response) (VirtualNetworkLinksClientListResponse, error) {
	result := VirtualNetworkLinksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkLinkListResult); err != nil {
		return VirtualNetworkLinksClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
//   - resourceGroupName - The name of the resource group.
//   - privateZoneName - The name of the Private DNS zone (without a terminating dot).
//   - virtualNetworkLinkName - The name of the virtual network link.
//   - parameters - Parameters supplied to the Update operation.
//   - options - VirtualNetworkLinksClientBeginUpdateOptions contains the optional parameters for the VirtualNetworkLinksClient.BeginUpdate
//     method.
func (client *VirtualNetworkLinksClient) BeginUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginUpdateOptions) (*runtime.Poller[VirtualNetworkLinksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[VirtualNetworkLinksClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworkLinksClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a virtual network link to the specified Private DNS zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-06-01
func (client *VirtualNetworkLinksClient) update(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualNetworkLinksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters VirtualNetworkLink, options *VirtualNetworkLinksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateZoneName == "" {
		return nil, errors.New("parameter privateZoneName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateZoneName}", url.PathEscape(privateZoneName))
	if virtualNetworkLinkName == "" {
		return nil, errors.New("parameter virtualNetworkLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkLinkName}", url.PathEscape(virtualNetworkLinkName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
