# \SitesEventsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteSystemEvents**](SitesEventsAPI.md#CountSiteSystemEvents) | **Get** /api/v1/sites/{site_id}/events/system/count | countSiteSystemEvents
[**GetSiteRoamingEvents**](SitesEventsAPI.md#GetSiteRoamingEvents) | **Get** /api/v1/sites/{site_id}/events/fast_roam | getSiteRoamingEvents
[**SearchSiteSystemEvents**](SitesEventsAPI.md#SearchSiteSystemEvents) | **Get** /api/v1/sites/{site_id}/events/system/search | searchSiteSystemEvents



## CountSiteSystemEvents

> RepsonseCount CountSiteSystemEvents(ctx, siteId).Distinct(distinct).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteSystemEvents



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
	distinct := openapiclient.site_system_events_count_distinct("") // SiteSystemEventsCountDistinct |  (optional) (default to "type")
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesEventsAPI.CountSiteSystemEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEventsAPI.CountSiteSystemEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteSystemEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesEventsAPI.CountSiteSystemEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteSystemEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteSystemEventsCountDistinct**](SiteSystemEventsCountDistinct.md) |  | [default to &quot;type&quot;]
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**RepsonseCount**](RepsonseCount.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteRoamingEvents

> ResponseEventsFastroam GetSiteRoamingEvents(ctx, siteId).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()

getSiteRoamingEvents



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
	type_ := openapiclient.fast_roam_result("") // FastRoamResult | event type (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesEventsAPI.GetSiteRoamingEvents(context.Background(), siteId).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEventsAPI.GetSiteRoamingEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteRoamingEvents`: ResponseEventsFastroam
	fmt.Fprintf(os.Stdout, "Response from `SitesEventsAPI.GetSiteRoamingEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRoamingEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**FastRoamResult**](FastRoamResult.md) | event type | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseEventsFastroam**](ResponseEventsFastroam.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteSystemEvents

> ResponseDeviceEventsSearch SearchSiteSystemEvents(ctx, siteId).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteSystemEvents



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
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesEventsAPI.SearchSiteSystemEvents(context.Background(), siteId).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEventsAPI.SearchSiteSystemEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteSystemEvents`: ResponseDeviceEventsSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesEventsAPI.SearchSiteSystemEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteSystemEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseDeviceEventsSearch**](ResponseDeviceEventsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

