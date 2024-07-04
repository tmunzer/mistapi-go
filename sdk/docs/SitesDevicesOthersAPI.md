# \SitesDevicesOthersAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteOtherDeviceEvents**](SitesDevicesOthersAPI.md#CountSiteOtherDeviceEvents) | **Get** /api/v1/sites/{site_id}/otherdevices/events/count | countSiteOtherDeviceEvents
[**ListSiteOtherDevices**](SitesDevicesOthersAPI.md#ListSiteOtherDevices) | **Get** /api/v1/sites/{site_id}/otherdevices | listSiteOtherDevices
[**SearchSiteOtherDeviceEvents**](SitesDevicesOthersAPI.md#SearchSiteOtherDeviceEvents) | **Get** /api/v1/sites/{site_id}/otherdevices/events/search | searchSiteOtherDeviceEvents



## CountSiteOtherDeviceEvents

> RepsonseCount CountSiteOtherDeviceEvents(ctx, siteId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countSiteOtherDeviceEvents



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
	distinct := openapiclient.site_other_device_events_count_distinct("") // SiteOtherDeviceEventsCountDistinct |  (optional) (default to "mac")
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesOthersAPI.CountSiteOtherDeviceEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesOthersAPI.CountSiteOtherDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteOtherDeviceEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesOthersAPI.CountSiteOtherDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteOtherDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteOtherDeviceEventsCountDistinct**](SiteOtherDeviceEventsCountDistinct.md) |  | [default to &quot;mac&quot;]
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

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


## ListSiteOtherDevices

> []DeviceOther ListSiteOtherDevices(ctx, siteId).Vendor(vendor).Mac(mac).Serial(serial).Model(model).Name(name).Page(page).Limit(limit).Execute()

listSiteOtherDevices



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
	vendor := "vendor_example" // string |  (optional)
	mac := "mac_example" // string |  (optional)
	serial := "serial_example" // string |  (optional)
	model := "model_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesOthersAPI.ListSiteOtherDevices(context.Background(), siteId).Vendor(vendor).Mac(mac).Serial(serial).Model(model).Name(name).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesOthersAPI.ListSiteOtherDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteOtherDevices`: []DeviceOther
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesOthersAPI.ListSiteOtherDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteOtherDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vendor** | **string** |  | 
 **mac** | **string** |  | 
 **serial** | **string** |  | 
 **model** | **string** |  | 
 **name** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]DeviceOther**](DeviceOther.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteOtherDeviceEvents

> ResponseEventsOtherDevicesSearch SearchSiteOtherDeviceEvents(ctx, siteId).Mac(mac).DeviceMac(deviceMac).Vendor(vendor).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchSiteOtherDeviceEvents



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
	mac := "mac_example" // string | mac (optional)
	deviceMac := "deviceMac_example" // string | mac of attached device (optional)
	vendor := "vendor_example" // string | vendor name (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesOthersAPI.SearchSiteOtherDeviceEvents(context.Background(), siteId).Mac(mac).DeviceMac(deviceMac).Vendor(vendor).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesOthersAPI.SearchSiteOtherDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteOtherDeviceEvents`: ResponseEventsOtherDevicesSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesOthersAPI.SearchSiteOtherDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteOtherDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | mac | 
 **deviceMac** | **string** | mac of attached device | 
 **vendor** | **string** | vendor name | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**ResponseEventsOtherDevicesSearch**](ResponseEventsOtherDevicesSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

