/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// +azure:enableclientgen:=true
package interfaceclient

import (
	"context"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"

	"sigs.k8s.io/cloud-provider-azure/pkg/azclient/utils"
)

// +azure:client:verbs=get;createorupdate;delete;list,resource=Interface,packageName=github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3,packageAlias=armnetwork,clientName=InterfacesClient,expand=true,rateLimitKey=interfaceRateLimit
type Interface interface {
	// GetVirtualMachineScaleSetNetworkInterface gets a network.Interface of VMSS VM.
	GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, options *armnetwork.InterfacesClientGetVirtualMachineScaleSetNetworkInterfaceOptions) (armnetwork.InterfacesClientGetVirtualMachineScaleSetNetworkInterfaceResponse, error)
	utils.GetWithExpandFunc[armnetwork.Interface]
	utils.CreateOrUpdateFunc[armnetwork.Interface]
	utils.DeleteFunc[armnetwork.Interface]
	utils.ListFunc[armnetwork.Interface]
}
