# \OrgsWxTunnelsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWxTunnel**](OrgsWxTunnelsAPI.md#CreateOrgWxTunnel) | **Post** /api/v1/orgs/{org_id}/wxtunnels | createOrgWxTunnel
[**DeleteOrgWxTunnel**](OrgsWxTunnelsAPI.md#DeleteOrgWxTunnel) | **Delete** /api/v1/orgs/{org_id}/wxtunnels/{wxtunnel_id} | deleteOrgWxTunnel
[**GetOrgWxTunnel**](OrgsWxTunnelsAPI.md#GetOrgWxTunnel) | **Get** /api/v1/orgs/{org_id}/wxtunnels/{wxtunnel_id} | getOrgWxTunnel
[**ListOrgWxTunnels**](OrgsWxTunnelsAPI.md#ListOrgWxTunnels) | **Get** /api/v1/orgs/{org_id}/wxtunnels | listOrgWxTunnels
[**UpdateOrgWxTunnel**](OrgsWxTunnelsAPI.md#UpdateOrgWxTunnel) | **Put** /api/v1/orgs/{org_id}/wxtunnels/{wxtunnel_id} | updateOrgWxTunnel



## CreateOrgWxTunnel

> WxlanTunnel CreateOrgWxTunnel(ctx, orgId).WxlanTunnel(wxlanTunnel).Execute()

createOrgWxTunnel



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
	wxlanTunnel := *openapiclient.NewWxlanTunnel("Name_example") // WxlanTunnel | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxTunnelsAPI.CreateOrgWxTunnel(context.Background(), orgId).WxlanTunnel(wxlanTunnel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTunnelsAPI.CreateOrgWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgWxTunnel`: WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTunnelsAPI.CreateOrgWxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWxTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wxlanTunnel** | [**WxlanTunnel**](WxlanTunnel.md) | Request Body | 

### Return type

[**WxlanTunnel**](WxlanTunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgWxTunnel

> DeleteOrgWxTunnel(ctx, orgId, wxtunnelId).Execute()

deleteOrgWxTunnel



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
	wxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsWxTunnelsAPI.DeleteOrgWxTunnel(context.Background(), orgId, wxtunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTunnelsAPI.DeleteOrgWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWxTunnelRequest struct via the builder pattern


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


## GetOrgWxTunnel

> WxlanTunnel GetOrgWxTunnel(ctx, orgId, wxtunnelId).Execute()

getOrgWxTunnel



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
	wxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxTunnelsAPI.GetOrgWxTunnel(context.Background(), orgId, wxtunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTunnelsAPI.GetOrgWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgWxTunnel`: WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTunnelsAPI.GetOrgWxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWxTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WxlanTunnel**](WxlanTunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgWxTunnels

> []WxlanTunnel ListOrgWxTunnels(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgWxTunnels



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
	resp, r, err := apiClient.OrgsWxTunnelsAPI.ListOrgWxTunnels(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTunnelsAPI.ListOrgWxTunnels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgWxTunnels`: []WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTunnelsAPI.ListOrgWxTunnels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWxTunnelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]WxlanTunnel**](WxlanTunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgWxTunnel

> WxlanTunnel UpdateOrgWxTunnel(ctx, orgId, wxtunnelId).WxlanTunnel(wxlanTunnel).Execute()

updateOrgWxTunnel



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
	wxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanTunnel := *openapiclient.NewWxlanTunnel("Name_example") // WxlanTunnel | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxTunnelsAPI.UpdateOrgWxTunnel(context.Background(), orgId, wxtunnelId).WxlanTunnel(wxlanTunnel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTunnelsAPI.UpdateOrgWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgWxTunnel`: WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTunnelsAPI.UpdateOrgWxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWxTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wxlanTunnel** | [**WxlanTunnel**](WxlanTunnel.md) | Request Body | 

### Return type

[**WxlanTunnel**](WxlanTunnel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

