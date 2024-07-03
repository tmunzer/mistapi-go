# \SitesWxTunnelsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWxTunnel**](SitesWxTunnelsAPI.md#CreateSiteWxTunnel) | **Post** /api/v1/sites/{site_id}/wxtunnels | createSiteWxTunnel
[**DeleteSiteWxTunnel**](SitesWxTunnelsAPI.md#DeleteSiteWxTunnel) | **Delete** /api/v1/sites/{site_id}/wxtunnels/{wxtunnel_id} | deleteSiteWxTunnel
[**GetSiteWxTunnel**](SitesWxTunnelsAPI.md#GetSiteWxTunnel) | **Get** /api/v1/sites/{site_id}/wxtunnels/{wxtunnel_id} | getSiteWxTunnel
[**ListSiteWxTunnels**](SitesWxTunnelsAPI.md#ListSiteWxTunnels) | **Get** /api/v1/sites/{site_id}/wxtunnels | listSiteWxTunnels
[**UpdateSiteWxTunnel**](SitesWxTunnelsAPI.md#UpdateSiteWxTunnel) | **Put** /api/v1/sites/{site_id}/wxtunnels/{wxtunnel_id} | updateSiteWxTunnel



## CreateSiteWxTunnel

> WxlanTunnel CreateSiteWxTunnel(ctx, siteId).WxlanTunnel(wxlanTunnel).Execute()

createSiteWxTunnel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanTunnel := *openapiclient.NewWxlanTunnel("Name_example") // WxlanTunnel | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTunnelsAPI.CreateSiteWxTunnel(context.Background(), siteId).WxlanTunnel(wxlanTunnel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTunnelsAPI.CreateSiteWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWxTunnel`: WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTunnelsAPI.CreateSiteWxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWxTunnelRequest struct via the builder pattern


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


## DeleteSiteWxTunnel

> DeleteSiteWxTunnel(ctx, siteId, wxtunnelId).Execute()

deleteSiteWxTunnel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesWxTunnelsAPI.DeleteSiteWxTunnel(context.Background(), siteId, wxtunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTunnelsAPI.DeleteSiteWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteWxTunnelRequest struct via the builder pattern


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


## GetSiteWxTunnel

> WxlanTunnel GetSiteWxTunnel(ctx, siteId, wxtunnelId).Execute()

getSiteWxTunnel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTunnelsAPI.GetSiteWxTunnel(context.Background(), siteId, wxtunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTunnelsAPI.GetSiteWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWxTunnel`: WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTunnelsAPI.GetSiteWxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWxTunnelRequest struct via the builder pattern


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


## ListSiteWxTunnels

> []WxlanTunnel ListSiteWxTunnels(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteWxTunnels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTunnelsAPI.ListSiteWxTunnels(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTunnelsAPI.ListSiteWxTunnels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteWxTunnels`: []WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTunnelsAPI.ListSiteWxTunnels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteWxTunnelsRequest struct via the builder pattern


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


## UpdateSiteWxTunnel

> WxlanTunnel UpdateSiteWxTunnel(ctx, siteId, wxtunnelId).WxlanTunnel(wxlanTunnel).Execute()

updateSiteWxTunnel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func main() {
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanTunnel := *openapiclient.NewWxlanTunnel("Name_example") // WxlanTunnel | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTunnelsAPI.UpdateSiteWxTunnel(context.Background(), siteId, wxtunnelId).WxlanTunnel(wxlanTunnel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTunnelsAPI.UpdateSiteWxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteWxTunnel`: WxlanTunnel
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTunnelsAPI.UpdateSiteWxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteWxTunnelRequest struct via the builder pattern


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

