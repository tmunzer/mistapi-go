# \SitesCallsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteCalls**](SitesCallsAPI.md#CountSiteCalls) | **Get** /api/v1/sites/{site_id}/stats/calls/count | countSiteCalls
[**ListSiteTroubleshootCalls**](SitesCallsAPI.md#ListSiteTroubleshootCalls) | **Get** /api/v1/sites/{site_id}/stats/calls/troubleshoot | listSiteTroubleshootCalls
[**SearchSiteCalls**](SitesCallsAPI.md#SearchSiteCalls) | **Get** /api/v1/sites/{site_id}/stats/calls/search | searchSiteCalls
[**TroubleshootSiteCall**](SitesCallsAPI.md#TroubleshootSiteCall) | **Get** /api/v1/sites/{site_id}/stats/calls/client/{client_mac}/troubleshoot | troubleshootSiteCall



## CountSiteCalls

> RepsonseCount CountSiteCalls(ctx, siteId).Distrinct(distrinct).Rating(rating).App(app).Start(start).End(end).Execute()

countSiteCalls



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
	distrinct := openapiclient.count_site_calls_distrinct("") // CountSiteCallsDistrinct |  (optional) (default to "mac")
	rating := int32(56) // int32 | feedback rating (e.g. \"rating=1\" or \"rating=1,2\") (optional)
	app := "app_example" // string |  (optional)
	start := "start_example" // string |  (optional)
	end := "end_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesCallsAPI.CountSiteCalls(context.Background(), siteId).Distrinct(distrinct).Rating(rating).App(app).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesCallsAPI.CountSiteCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteCalls`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesCallsAPI.CountSiteCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distrinct** | [**CountSiteCallsDistrinct**](CountSiteCallsDistrinct.md) |  | [default to &quot;mac&quot;]
 **rating** | **int32** | feedback rating (e.g. \&quot;rating&#x3D;1\&quot; or \&quot;rating&#x3D;1,2\&quot;) | 
 **app** | **string** |  | 
 **start** | **string** |  | 
 **end** | **string** |  | 

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


## ListSiteTroubleshootCalls

> CallTroubleshootSummary ListSiteTroubleshootCalls(ctx, siteId, clientMac).Ap(ap).MeetingId(meetingId).Mac(mac).App(app).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

listSiteTroubleshootCalls



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
	clientMac := "0000000000ab" // string | 
	ap := "ap_example" // string | AP MAC (optional)
	meetingId := "meetingId_example" // string | meeting_id (optional)
	mac := "mac_example" // string | device identifier (optional)
	app := "zoom" // string | Third party app name (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesCallsAPI.ListSiteTroubleshootCalls(context.Background(), siteId, clientMac).Ap(ap).MeetingId(meetingId).Mac(mac).App(app).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesCallsAPI.ListSiteTroubleshootCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteTroubleshootCalls`: CallTroubleshootSummary
	fmt.Fprintf(os.Stdout, "Response from `SitesCallsAPI.ListSiteTroubleshootCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteTroubleshootCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ap** | **string** | AP MAC | 
 **meetingId** | **string** | meeting_id | 
 **mac** | **string** | device identifier | 
 **app** | **string** | Third party app name | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**CallTroubleshootSummary**](CallTroubleshootSummary.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteCalls

> ResponseStatsCalls SearchSiteCalls(ctx, siteId).Mac(mac).App(app).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteCalls



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
	mac := "mac_example" // string | device identifier (optional)
	app := "zoom" // string | Third party app name (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesCallsAPI.SearchSiteCalls(context.Background(), siteId).Mac(mac).App(app).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesCallsAPI.SearchSiteCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteCalls`: ResponseStatsCalls
	fmt.Fprintf(os.Stdout, "Response from `SitesCallsAPI.SearchSiteCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | device identifier | 
 **app** | **string** | Third party app name | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseStatsCalls**](ResponseStatsCalls.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TroubleshootSiteCall

> CallTroubleshoot TroubleshootSiteCall(ctx, siteId, clientMac).MeetingId(meetingId).Mac(mac).App(app).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

troubleshootSiteCall



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
	clientMac := "0000000000ab" // string | 
	meetingId := "meetingId_example" // string | meeting_id
	mac := "mac_example" // string | device identifier (optional)
	app := "zoom" // string | Third party app name (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesCallsAPI.TroubleshootSiteCall(context.Background(), siteId, clientMac).MeetingId(meetingId).Mac(mac).App(app).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesCallsAPI.TroubleshootSiteCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TroubleshootSiteCall`: CallTroubleshoot
	fmt.Fprintf(os.Stdout, "Response from `SitesCallsAPI.TroubleshootSiteCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTroubleshootSiteCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **meetingId** | **string** | meeting_id | 
 **mac** | **string** | device identifier | 
 **app** | **string** | Third party app name | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**CallTroubleshoot**](CallTroubleshoot.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

