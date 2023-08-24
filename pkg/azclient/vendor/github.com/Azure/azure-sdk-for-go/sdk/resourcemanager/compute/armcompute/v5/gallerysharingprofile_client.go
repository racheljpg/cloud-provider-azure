//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// GallerySharingProfileClient contains the methods for the GallerySharingProfile group.
// Don't use this type directly, use NewGallerySharingProfileClient() instead.
type GallerySharingProfileClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGallerySharingProfileClient creates a new instance of GallerySharingProfileClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGallerySharingProfileClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GallerySharingProfileClient, error) {
	cl, err := arm.NewClient(moduleName+".GallerySharingProfileClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GallerySharingProfileClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginUpdate - Update sharing profile of a gallery.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-03
//   - resourceGroupName - The name of the resource group.
//   - galleryName - The name of the Shared Image Gallery.
//   - sharingUpdate - Parameters supplied to the update gallery sharing profile.
//   - options - GallerySharingProfileClientBeginUpdateOptions contains the optional parameters for the GallerySharingProfileClient.BeginUpdate
//     method.
func (client *GallerySharingProfileClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileClientBeginUpdateOptions) (*runtime.Poller[GallerySharingProfileClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, galleryName, sharingUpdate, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[GallerySharingProfileClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[GallerySharingProfileClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update sharing profile of a gallery.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-03
func (client *GallerySharingProfileClient) update(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, sharingUpdate, options)
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
func (client *GallerySharingProfileClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, sharingUpdate SharingUpdate, options *GallerySharingProfileClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/share"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sharingUpdate)
}
