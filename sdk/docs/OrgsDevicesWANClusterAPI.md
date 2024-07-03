# \OrgsDevicesWANClusterAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgGatewayHaCluster**](OrgsDevicesWANClusterAPI.md#CreateOrgGatewayHaCluster) | **Post** /api/v1/orgs/{org_id}/inventory/create_ha_cluster | createOrgGatewayHaCluster
[**DeleteOrgGatewayHaCluster**](OrgsDevicesWANClusterAPI.md#DeleteOrgGatewayHaCluster) | **Post** /api/v1/orgs/{org_id}/inventory/delete_ha_cluster | deleteOrgGatewayHaCluster



## CreateOrgGatewayHaCluster

> CreateOrgGatewayHaCluster(ctx, orgId).HaClusterConfig(haClusterConfig).Execute()

createOrgGatewayHaCluster



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	haClusterConfig := *openapiclient.NewHaClusterConfig() // HaClusterConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDevicesWANClusterAPI.CreateOrgGatewayHaCluster(context.Background(), orgId).HaClusterConfig(haClusterConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesWANClusterAPI.CreateOrgGatewayHaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgGatewayHaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **haClusterConfig** | [**HaClusterConfig**](HaClusterConfig.md) |  | 

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


## DeleteOrgGatewayHaCluster

> DeleteOrgGatewayHaCluster(ctx, orgId).HaClusterDelete(haClusterDelete).Execute()

deleteOrgGatewayHaCluster



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	haClusterDelete := *openapiclient.NewHaClusterDelete() // HaClusterDelete |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDevicesWANClusterAPI.DeleteOrgGatewayHaCluster(context.Background(), orgId).HaClusterDelete(haClusterDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesWANClusterAPI.DeleteOrgGatewayHaCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgGatewayHaClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **haClusterDelete** | [**HaClusterDelete**](HaClusterDelete.md) |  | 

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

