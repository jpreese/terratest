package azure

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-04-01/network"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func GetVirtualNetwork(t *testing.T, subscription string, resourceGroupName string, virtualNetworkName string) (network.VirtualNetwork, error) {
	vnetClient := NewVirtualNetworksClient(t, subscription)

	response, err := vnetClient.Get(context.Background(), resourceGroupName, virtualNetworkName, "")
	if err != nil {
		return network.VirtualNetwork{}, err
	}

	return response, nil
}

func GetSubnetsForVirtualNetwork(t *testing.T, subscription string, resourceGroupName string, virtualNetworkName string) (*[]network.Subnet, error) {
	virtualNetwork, err := GetVirtualNetwork(t, subscription, resourceGroupName, virtualNetworkName)
	if err != nil {
		return nil, err
	}

	if properties := virtualNetwork.VirtualNetworkPropertiesFormat; properties != nil {
		return properties.Subnets, nil
	}

	return &[]network.Subnet{}, nil
}

func NewVirtualNetworksClient(t *testing.T, subscription string) network.VirtualNetworksClient {
	virtualNetworksClient := network.NewVirtualNetworksClient(subscription)
	authorizer, err := auth.NewAuthorizerFromEnvironment()

	if err == nil {
		virtualNetworksClient.Authorizer = authorizer
	}

	return virtualNetworksClient
}
