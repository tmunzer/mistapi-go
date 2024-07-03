# \SitesWlansAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWlan**](SitesWlansAPI.md#CreateSiteWlan) | **Post** /api/v1/sites/{site_id}/wlans | createSiteWlan
[**DeleteSiteWlan**](SitesWlansAPI.md#DeleteSiteWlan) | **Delete** /api/v1/sites/{site_id}/wlans/{wlan_id} | deleteSiteWlan
[**GetSiteWlan**](SitesWlansAPI.md#GetSiteWlan) | **Get** /api/v1/sites/{site_id}/wlans/{wlan_id} | getSiteWlan
[**ListSiteWlanDerived**](SitesWlansAPI.md#ListSiteWlanDerived) | **Get** /api/v1/sites/{site_id}/wlans/derived | listSiteWlanDerived
[**ListSiteWlans**](SitesWlansAPI.md#ListSiteWlans) | **Get** /api/v1/sites/{site_id}/wlans | listSiteWlans
[**UpdateSiteWlan**](SitesWlansAPI.md#UpdateSiteWlan) | **Put** /api/v1/sites/{site_id}/wlans/{wlan_id} | updateSiteWlan
[**UpdateSiteWlanPortalTemplate**](SitesWlansAPI.md#UpdateSiteWlanPortalTemplate) | **Put** /api/v1/sites/{site_id}/wlans/{wlan_id}/portal_template | updateSiteWlanPortalTemplate
[**UploadSiteWlanPortalImage**](SitesWlansAPI.md#UploadSiteWlanPortalImage) | **Post** /api/v1/sites/{site_id}/wlans/{wlan_id}/portal_image | uploadSiteWlanPortalImage



## CreateSiteWlan

> Wlan CreateSiteWlan(ctx, siteId).Wlan(wlan).Execute()

createSiteWlan



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
	wlan := *openapiclient.NewWlan("corporate") // Wlan | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWlansAPI.CreateSiteWlan(context.Background(), siteId).Wlan(wlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.CreateSiteWlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWlan`: Wlan
	fmt.Fprintf(os.Stdout, "Response from `SitesWlansAPI.CreateSiteWlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlan** | [**Wlan**](Wlan.md) | Request Body | 

### Return type

[**Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteWlan

> DeleteSiteWlan(ctx, siteId, wlanId).Execute()

deleteSiteWlan



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesWlansAPI.DeleteSiteWlan(context.Background(), siteId, wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.DeleteSiteWlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteWlanRequest struct via the builder pattern


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


## GetSiteWlan

> Wlan GetSiteWlan(ctx, siteId, wlanId).Execute()

getSiteWlan



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWlansAPI.GetSiteWlan(context.Background(), siteId, wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.GetSiteWlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWlan`: Wlan
	fmt.Fprintf(os.Stdout, "Response from `SitesWlansAPI.GetSiteWlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteWlanDerived

> []Wlan ListSiteWlanDerived(ctx, siteId).Resolve(resolve).WlanId(wlanId).Execute()

listSiteWlanDerived



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
	resolve := true // bool | whether to resolve SITE_VARS (optional) (default to false)
	wlanId := "wlanId_example" // string | filter by WLAN ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWlansAPI.ListSiteWlanDerived(context.Background(), siteId).Resolve(resolve).WlanId(wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.ListSiteWlanDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteWlanDerived`: []Wlan
	fmt.Fprintf(os.Stdout, "Response from `SitesWlansAPI.ListSiteWlanDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteWlanDerivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **bool** | whether to resolve SITE_VARS | [default to false]
 **wlanId** | **string** | filter by WLAN ID | 

### Return type

[**[]Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteWlans

> []Wlan ListSiteWlans(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteWlans



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
	resp, r, err := apiClient.SitesWlansAPI.ListSiteWlans(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.ListSiteWlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteWlans`: []Wlan
	fmt.Fprintf(os.Stdout, "Response from `SitesWlansAPI.ListSiteWlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteWlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteWlan

> Wlan UpdateSiteWlan(ctx, siteId, wlanId).Wlan(wlan).Execute()

updateSiteWlan



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wlan := *openapiclient.NewWlan("corporate") // Wlan | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWlansAPI.UpdateSiteWlan(context.Background(), siteId, wlanId).Wlan(wlan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.UpdateSiteWlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteWlan`: Wlan
	fmt.Fprintf(os.Stdout, "Response from `SitesWlansAPI.UpdateSiteWlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteWlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wlan** | [**Wlan**](Wlan.md) | Request Body | 

### Return type

[**Wlan**](Wlan.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteWlanPortalTemplate

> PortalTemplate UpdateSiteWlanPortalTemplate(ctx, siteId, wlanId).WlanPortalTemplate(wlanPortalTemplate).Execute()

updateSiteWlanPortalTemplate



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wlanPortalTemplate := *openapiclient.NewWlanPortalTemplate() // WlanPortalTemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWlansAPI.UpdateSiteWlanPortalTemplate(context.Background(), siteId, wlanId).WlanPortalTemplate(wlanPortalTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.UpdateSiteWlanPortalTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteWlanPortalTemplate`: PortalTemplate
	fmt.Fprintf(os.Stdout, "Response from `SitesWlansAPI.UpdateSiteWlanPortalTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteWlanPortalTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wlanPortalTemplate** | [**WlanPortalTemplate**](WlanPortalTemplate.md) | Request Body | 

### Return type

[**PortalTemplate**](PortalTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSiteWlanPortalImage

> UploadSiteWlanPortalImage(ctx, siteId, wlanId).File(file).Json(json).Execute()

uploadSiteWlanPortalImage



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
	wlanId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | binary file
	json := "json_example" // string | JSON string describing your upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesWlansAPI.UploadSiteWlanPortalImage(context.Background(), siteId, wlanId).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWlansAPI.UploadSiteWlanPortalImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSiteWlanPortalImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | binary file | 
 **json** | **string** | JSON string describing your upload | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

