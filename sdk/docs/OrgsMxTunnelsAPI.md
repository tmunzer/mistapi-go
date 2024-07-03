# \OrgsMxTunnelsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgMxTunnel**](OrgsMxTunnelsAPI.md#CreateOrgMxTunnel) | **Post** /api/v1/orgs/{org_id}/mxtunnels | createOrgMxTunnel
[**DeleteOrgMxTunnel**](OrgsMxTunnelsAPI.md#DeleteOrgMxTunnel) | **Delete** /api/v1/orgs/{org_id}/mxtunnels/{mxtunnel_id} | deleteOrgMxTunnel
[**GetOrgMxTunnel**](OrgsMxTunnelsAPI.md#GetOrgMxTunnel) | **Get** /api/v1/orgs/{org_id}/mxtunnels/{mxtunnel_id} | getOrgMxTunnel
[**ListOrgMxTunnels**](OrgsMxTunnelsAPI.md#ListOrgMxTunnels) | **Get** /api/v1/orgs/{org_id}/mxtunnels | listOrgMxTunnels
[**UpdateOrgMxTunnel**](OrgsMxTunnelsAPI.md#UpdateOrgMxTunnel) | **Put** /api/v1/orgs/{org_id}/mxtunnels/{mxtunnel_id} | updateOrgMxTunnel



## CreateOrgMxTunnel

> Mxtunnel CreateOrgMxTunnel(ctx, orgId).Mxtunnel(mxtunnel).Execute()

createOrgMxTunnel



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
	mxtunnel := *openapiclient.NewMxtunnel() // Mxtunnel | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxTunnelsAPI.CreateOrgMxTunnel(context.Background(), orgId).Mxtunnel(mxtunnel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxTunnelsAPI.CreateOrgMxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgMxTunnel`: Mxtunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxTunnelsAPI.CreateOrgMxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgMxTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxtunnel** | [**Mxtunnel**](Mxtunnel.md) | Request Body | 

### Return type

[**Mxtunnel**](Mxtunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgMxTunnel

> DeleteOrgMxTunnel(ctx, orgId, mxtunnelId).Execute()

deleteOrgMxTunnel



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
	mxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxTunnelsAPI.DeleteOrgMxTunnel(context.Background(), orgId, mxtunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxTunnelsAPI.DeleteOrgMxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMxTunnelRequest struct via the builder pattern


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


## GetOrgMxTunnel

> Mxtunnel GetOrgMxTunnel(ctx, orgId, mxtunnelId).Execute()

getOrgMxTunnel



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
	mxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxTunnelsAPI.GetOrgMxTunnel(context.Background(), orgId, mxtunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxTunnelsAPI.GetOrgMxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMxTunnel`: Mxtunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxTunnelsAPI.GetOrgMxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMxTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Mxtunnel**](Mxtunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMxTunnels

> []Mxtunnel ListOrgMxTunnels(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgMxTunnels



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
	resp, r, err := apiClient.OrgsMxTunnelsAPI.ListOrgMxTunnels(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxTunnelsAPI.ListOrgMxTunnels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgMxTunnels`: []Mxtunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxTunnelsAPI.ListOrgMxTunnels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMxTunnelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Mxtunnel**](Mxtunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMxTunnel

> Mxtunnel UpdateOrgMxTunnel(ctx, orgId, mxtunnelId).Mxtunnel(mxtunnel).Execute()

updateOrgMxTunnel



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
	mxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mxtunnel := *openapiclient.NewMxtunnel() // Mxtunnel | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxTunnelsAPI.UpdateOrgMxTunnel(context.Background(), orgId, mxtunnelId).Mxtunnel(mxtunnel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxTunnelsAPI.UpdateOrgMxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgMxTunnel`: Mxtunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxTunnelsAPI.UpdateOrgMxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMxTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mxtunnel** | [**Mxtunnel**](Mxtunnel.md) | Request Body | 

### Return type

[**Mxtunnel**](Mxtunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

