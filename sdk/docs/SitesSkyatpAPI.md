# \SitesSkyatpAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteSkyatpEvents**](SitesSkyatpAPI.md#CountSiteSkyatpEvents) | **Get** /api/v1/sites/{site_id}/skyatp/events/count | countSiteSkyatpEvents
[**SearchSiteSkyatpEvents**](SitesSkyatpAPI.md#SearchSiteSkyatpEvents) | **Get** /api/v1/sites/{site_id}/skyatp/events/search | searchSiteSkyatpEvents



## CountSiteSkyatpEvents

> RepsonseCount CountSiteSkyatpEvents(ctx, siteId).Distinct(distinct).Type_(type_).Mac(mac).DeviceMac(deviceMac).ThreatLevel(threatLevel).IpAddress(ipAddress).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteSkyatpEvents



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
	distinct := openapiclient.site_sky_atp_events_count_distinct("") // SiteSkyAtpEventsCountDistinct |  (optional) (default to "type")
	type_ := "type__example" // string | event type, e.g. cc, fs, mw (optional)
	mac := "mac_example" // string | client MAC (optional)
	deviceMac := "deviceMac_example" // string | device MAC (optional)
	threatLevel := int32(56) // int32 | threat level (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSkyatpAPI.CountSiteSkyatpEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).Mac(mac).DeviceMac(deviceMac).ThreatLevel(threatLevel).IpAddress(ipAddress).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSkyatpAPI.CountSiteSkyatpEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteSkyatpEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesSkyatpAPI.CountSiteSkyatpEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteSkyatpEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteSkyAtpEventsCountDistinct**](SiteSkyAtpEventsCountDistinct.md) |  | [default to &quot;type&quot;]
 **type_** | **string** | event type, e.g. cc, fs, mw | 
 **mac** | **string** | client MAC | 
 **deviceMac** | **string** | device MAC | 
 **threatLevel** | **int32** | threat level | 
 **ipAddress** | **string** |  | 
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


## SearchSiteSkyatpEvents

> ResponseEventsSkyAtpSearch SearchSiteSkyatpEvents(ctx, siteId).Type_(type_).Mac(mac).DeviceMac(deviceMac).ThreatLevel(threatLevel).IpAddress(ipAddress).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteSkyatpEvents



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
	type_ := "type__example" // string | event type, e.g. cc, fs, mw (optional)
	mac := "mac_example" // string | client MAC (optional)
	deviceMac := "deviceMac_example" // string | device MAC (optional)
	threatLevel := int32(56) // int32 | threat level (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSkyatpAPI.SearchSiteSkyatpEvents(context.Background(), siteId).Type_(type_).Mac(mac).DeviceMac(deviceMac).ThreatLevel(threatLevel).IpAddress(ipAddress).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSkyatpAPI.SearchSiteSkyatpEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteSkyatpEvents`: ResponseEventsSkyAtpSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesSkyatpAPI.SearchSiteSkyatpEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteSkyatpEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | event type, e.g. cc, fs, mw | 
 **mac** | **string** | client MAC | 
 **deviceMac** | **string** | device MAC | 
 **threatLevel** | **int32** | threat level | 
 **ipAddress** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseEventsSkyAtpSearch**](ResponseEventsSkyAtpSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

