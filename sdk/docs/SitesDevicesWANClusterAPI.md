# \SitesDevicesWANClusterAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteDeviceHaCluster**](SitesDevicesWANClusterAPI.md#CreateSiteDeviceHaCluster) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/ha | createSiteDeviceHaCluster
[**DeleteSiteDeviceHaCluster**](SitesDevicesWANClusterAPI.md#DeleteSiteDeviceHaCluster) | **Delete** /api/v1/sites/{site_id}/devices/{device_id}/ha | deleteSiteDeviceHaCluster
[**SwapSiteDeviceHaClusterNode**](SitesDevicesWANClusterAPI.md#SwapSiteDeviceHaClusterNode) | **Put** /api/v1/sites/{site_id}/devices/{device_id}/ha | swapSiteDeviceHaClusterNode



## CreateSiteDeviceHaCluster

> CreateSiteDeviceHaCluster(ctx, siteId, deviceId).GatewayClusterForm(gatewayClusterForm).Execute()

createSiteDeviceHaCluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/mistapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	gatewayClusterForm := *openapiclient.NewGatewayClusterForm([]openapiclient.GatewayClusterFormNode{*openapiclient.NewGatewayClusterFormNode("Mac_example")}) // GatewayClusterForm |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesWANClusterAPI.CreateSiteDeviceHaCluster(context.Background(), siteId, deviceId).GatewayClusterForm(gatewayClusterForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWANClusterAPI.CreateSiteDeviceHaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteDeviceHaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayClusterForm** | [**GatewayClusterForm**](GatewayClusterForm.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteDeviceHaCluster

> DeleteSiteDeviceHaCluster(ctx, siteId, deviceId).Execute()

deleteSiteDeviceHaCluster



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/mistapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesWANClusterAPI.DeleteSiteDeviceHaCluster(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWANClusterAPI.DeleteSiteDeviceHaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteDeviceHaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapSiteDeviceHaClusterNode

> SwapSiteDeviceHaClusterNode(ctx, siteId, deviceId).GatewayClusterSwap(gatewayClusterSwap).Execute()

swapSiteDeviceHaClusterNode



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/mistapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	gatewayClusterSwap := *openapiclient.NewGatewayClusterSwap(openapiclient.gateway_cluster_swap_op("")) // GatewayClusterSwap |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesWANClusterAPI.SwapSiteDeviceHaClusterNode(context.Background(), siteId, deviceId).GatewayClusterSwap(gatewayClusterSwap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWANClusterAPI.SwapSiteDeviceHaClusterNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapSiteDeviceHaClusterNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayClusterSwap** | [**GatewayClusterSwap**](GatewayClusterSwap.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

