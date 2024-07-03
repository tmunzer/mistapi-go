# \OrgsNetworksAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNetwork**](OrgsNetworksAPI.md#CreateOrgNetwork) | **Post** /api/v1/orgs/{org_id}/networks | createOrgNetwork
[**DeleteOrgNetwork**](OrgsNetworksAPI.md#DeleteOrgNetwork) | **Delete** /api/v1/orgs/{org_id}/networks/{network_id} | deleteOrgNetwork
[**GetOrgNetwork**](OrgsNetworksAPI.md#GetOrgNetwork) | **Get** /api/v1/orgs/{org_id}/networks/{network_id} | getOrgNetwork
[**ListOrgNetworks**](OrgsNetworksAPI.md#ListOrgNetworks) | **Get** /api/v1/orgs/{org_id}/networks | listOrgNetworks
[**UpdateOrgNetwork**](OrgsNetworksAPI.md#UpdateOrgNetwork) | **Put** /api/v1/orgs/{org_id}/networks/{network_id} | updateOrgNetwork



## CreateOrgNetwork

> Network CreateOrgNetwork(ctx, orgId).Network(network).Execute()

createOrgNetwork



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
	network := *openapiclient.NewNetwork() // Network |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNetworksAPI.CreateOrgNetwork(context.Background(), orgId).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworksAPI.CreateOrgNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgNetwork`: Network
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworksAPI.CreateOrgNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **network** | [**Network**](Network.md) |  | 

### Return type

[**Network**](Network.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgNetwork

> DeleteOrgNetwork(ctx, orgId, networkId).Execute()

deleteOrgNetwork



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
	networkId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsNetworksAPI.DeleteOrgNetwork(context.Background(), orgId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworksAPI.DeleteOrgNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgNetworkRequest struct via the builder pattern


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


## GetOrgNetwork

> Network GetOrgNetwork(ctx, orgId, networkId).Execute()

getOrgNetwork



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
	networkId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNetworksAPI.GetOrgNetwork(context.Background(), orgId, networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworksAPI.GetOrgNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgNetwork`: Network
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworksAPI.GetOrgNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Network**](Network.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgNetworks

> []Network ListOrgNetworks(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgNetworks



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNetworksAPI.ListOrgNetworks(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworksAPI.ListOrgNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgNetworks`: []Network
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworksAPI.ListOrgNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Network**](Network.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgNetwork

> Network UpdateOrgNetwork(ctx, orgId, networkId).Network(network).Execute()

updateOrgNetwork



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
	networkId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	network := *openapiclient.NewNetwork() // Network |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNetworksAPI.UpdateOrgNetwork(context.Background(), orgId, networkId).Network(network).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworksAPI.UpdateOrgNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgNetwork`: Network
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworksAPI.UpdateOrgNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **network** | [**Network**](Network.md) |  | 

### Return type

[**Network**](Network.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

