# \SitesRSSIZonesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteRssiZone**](SitesRSSIZonesAPI.md#CreateSiteRssiZone) | **Post** /api/v1/sites/{site_id}/rssizones | createSiteRssiZone
[**DeleteSiteRssiZone**](SitesRSSIZonesAPI.md#DeleteSiteRssiZone) | **Delete** /api/v1/sites/{site_id}/rssizones/{rssizone_id} | deleteSiteRssiZone
[**GetSiteRssiZone**](SitesRSSIZonesAPI.md#GetSiteRssiZone) | **Get** /api/v1/sites/{site_id}/rssizones/{rssizone_id} | getSiteRssiZone
[**ListSiteRssiZones**](SitesRSSIZonesAPI.md#ListSiteRssiZones) | **Get** /api/v1/sites/{site_id}/rssizones | listSiteRssiZones
[**UpdateSiteRssiZone**](SitesRSSIZonesAPI.md#UpdateSiteRssiZone) | **Put** /api/v1/sites/{site_id}/rssizones/{rssizone_id} | updateSiteRssiZone



## CreateSiteRssiZone

> RssiZone CreateSiteRssiZone(ctx, siteId).RssiZone(rssiZone).Execute()

createSiteRssiZone



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
	rssiZone := *openapiclient.NewRssiZone([]openapiclient.RssiZoneDevice{*openapiclient.NewRssiZoneDevice("b069b358-4c97-5319-1f8c-7c5ca64d6ab1", int32(0))}) // RssiZone | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRSSIZonesAPI.CreateSiteRssiZone(context.Background(), siteId).RssiZone(rssiZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRSSIZonesAPI.CreateSiteRssiZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteRssiZone`: RssiZone
	fmt.Fprintf(os.Stdout, "Response from `SitesRSSIZonesAPI.CreateSiteRssiZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteRssiZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rssiZone** | [**RssiZone**](RssiZone.md) | Request Body | 

### Return type

[**RssiZone**](RssiZone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteRssiZone

> DeleteSiteRssiZone(ctx, siteId, rssizoneId).Execute()

deleteSiteRssiZone



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
	rssizoneId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesRSSIZonesAPI.DeleteSiteRssiZone(context.Background(), siteId, rssizoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRSSIZonesAPI.DeleteSiteRssiZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rssizoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteRssiZoneRequest struct via the builder pattern


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


## GetSiteRssiZone

> []RssiZone GetSiteRssiZone(ctx, siteId, rssizoneId).Execute()

getSiteRssiZone



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
	rssizoneId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRSSIZonesAPI.GetSiteRssiZone(context.Background(), siteId, rssizoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRSSIZonesAPI.GetSiteRssiZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteRssiZone`: []RssiZone
	fmt.Fprintf(os.Stdout, "Response from `SitesRSSIZonesAPI.GetSiteRssiZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rssizoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRssiZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RssiZone**](RssiZone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteRssiZones

> []RssiZone ListSiteRssiZones(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteRssiZones



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
	resp, r, err := apiClient.SitesRSSIZonesAPI.ListSiteRssiZones(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRSSIZonesAPI.ListSiteRssiZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteRssiZones`: []RssiZone
	fmt.Fprintf(os.Stdout, "Response from `SitesRSSIZonesAPI.ListSiteRssiZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteRssiZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]RssiZone**](RssiZone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteRssiZone

> RssiZone UpdateSiteRssiZone(ctx, siteId, rssizoneId).RssiZone(rssiZone).Execute()

updateSiteRssiZone



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
	rssizoneId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rssiZone := *openapiclient.NewRssiZone([]openapiclient.RssiZoneDevice{*openapiclient.NewRssiZoneDevice("b069b358-4c97-5319-1f8c-7c5ca64d6ab1", int32(0))}) // RssiZone | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRSSIZonesAPI.UpdateSiteRssiZone(context.Background(), siteId, rssizoneId).RssiZone(rssiZone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRSSIZonesAPI.UpdateSiteRssiZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteRssiZone`: RssiZone
	fmt.Fprintf(os.Stdout, "Response from `SitesRSSIZonesAPI.UpdateSiteRssiZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rssizoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRssiZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rssiZone** | [**RssiZone**](RssiZone.md) | Request Body | 

### Return type

[**RssiZone**](RssiZone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

