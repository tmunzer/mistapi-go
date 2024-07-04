# \OrgsDevicesOthersAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgOtherDeviceEvents**](OrgsDevicesOthersAPI.md#CountOrgOtherDeviceEvents) | **Get** /api/v1/orgs/{org_id}/otherdevices/events/count | countOrgOtherDeviceEvents
[**DeleteOrgOtherDevice**](OrgsDevicesOthersAPI.md#DeleteOrgOtherDevice) | **Delete** /api/v1/orgs/{org_id}/otherdevices/{device_mac} | deleteOrgOtherDevice
[**GetOrgOtherDevice**](OrgsDevicesOthersAPI.md#GetOrgOtherDevice) | **Get** /api/v1/orgs/{org_id}/otherdevices/{device_mac} | getOrgOtherDevice
[**GetOrgOtherDeviceStats**](OrgsDevicesOthersAPI.md#GetOrgOtherDeviceStats) | **Get** /api/v1/orgs/{org_id}/stats/otherdevices/{device_mac} | getOrgOtherDeviceStats
[**ListOrgOtherDevices**](OrgsDevicesOthersAPI.md#ListOrgOtherDevices) | **Get** /api/v1/orgs/{org_id}/otherdevices | listOrgOtherDevices
[**RebootOrgOtherDevice**](OrgsDevicesOthersAPI.md#RebootOrgOtherDevice) | **Post** /api/v1/orgs/{org_id}/otherdevices/{device_mac}/reboot | rebootOrgOtherDevice
[**SearchOrgOtherDeviceEvents**](OrgsDevicesOthersAPI.md#SearchOrgOtherDeviceEvents) | **Get** /api/v1/orgs/{org_id}/otherdevices/events/search | searchOrgOtherDeviceEvents
[**UpdateOrgOtherDevice**](OrgsDevicesOthersAPI.md#UpdateOrgOtherDevice) | **Put** /api/v1/orgs/{org_id}/otherdevices/{device_mac} | updateOrgOtherDevice
[**UpdateOrgOtherDevices**](OrgsDevicesOthersAPI.md#UpdateOrgOtherDevices) | **Put** /api/v1/orgs/{org_id}/otherdevices | updateOrgOtherDevices



## CountOrgOtherDeviceEvents

> RepsonseCount CountOrgOtherDeviceEvents(ctx, orgId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countOrgOtherDeviceEvents



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
	distinct := openapiclient.org_otherdevices_events_count_distinct("") // OrgOtherdevicesEventsCountDistinct |  (optional) (default to "mac")
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesOthersAPI.CountOrgOtherDeviceEvents(context.Background(), orgId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.CountOrgOtherDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgOtherDeviceEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesOthersAPI.CountOrgOtherDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgOtherDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgOtherdevicesEventsCountDistinct**](OrgOtherdevicesEventsCountDistinct.md) |  | [default to &quot;mac&quot;]
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


## DeleteOrgOtherDevice

> DeleteOrgOtherDevice(ctx, orgId, deviceMac).Execute()

deleteOrgOtherDevice



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDevicesOthersAPI.DeleteOrgOtherDevice(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.DeleteOrgOtherDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgOtherDeviceRequest struct via the builder pattern


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


## GetOrgOtherDevice

> DeviceOther GetOrgOtherDevice(ctx, orgId, deviceMac).Execute()

getOrgOtherDevice



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesOthersAPI.GetOrgOtherDevice(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.GetOrgOtherDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgOtherDevice`: DeviceOther
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesOthersAPI.GetOrgOtherDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgOtherDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceOther**](DeviceOther.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgOtherDeviceStats

> DeviceOtherStats GetOrgOtherDeviceStats(ctx, orgId, deviceMac).Execute()

getOrgOtherDeviceStats



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesOthersAPI.GetOrgOtherDeviceStats(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.GetOrgOtherDeviceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgOtherDeviceStats`: DeviceOtherStats
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesOthersAPI.GetOrgOtherDeviceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgOtherDeviceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceOtherStats**](DeviceOtherStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgOtherDevices

> []DeviceOther ListOrgOtherDevices(ctx, orgId).Vendor(vendor).Mac(mac).Serial(serial).Model(model).Name(name).Page(page).Limit(limit).Execute()

listOrgOtherDevices



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
	vendor := "vendor_example" // string |  (optional)
	mac := "mac_example" // string |  (optional)
	serial := "serial_example" // string |  (optional)
	model := "model_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesOthersAPI.ListOrgOtherDevices(context.Background(), orgId).Vendor(vendor).Mac(mac).Serial(serial).Model(model).Name(name).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.ListOrgOtherDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgOtherDevices`: []DeviceOther
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesOthersAPI.ListOrgOtherDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgOtherDevicesRequest struct via the builder pattern


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


## RebootOrgOtherDevice

> RebootOrgOtherDevice(ctx, orgId, deviceMac).Execute()

rebootOrgOtherDevice



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDevicesOthersAPI.RebootOrgOtherDevice(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.RebootOrgOtherDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootOrgOtherDeviceRequest struct via the builder pattern


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


## SearchOrgOtherDeviceEvents

> ResponseEventsOtherDevicesSearch SearchOrgOtherDeviceEvents(ctx, orgId).SiteId(siteId).Mac(mac).DeviceMac(deviceMac).Model(model).Vendor(vendor).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchOrgOtherDeviceEvents



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
	siteId := "siteId_example" // string | site id (optional)
	mac := "mac_example" // string | mac (optional)
	deviceMac := "deviceMac_example" // string | mac of attached device (optional)
	model := "model_example" // string | device model (optional)
	vendor := "vendor_example" // string | vendor name (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listOtherDeviceEventsDefinitions) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesOthersAPI.SearchOrgOtherDeviceEvents(context.Background(), orgId).SiteId(siteId).Mac(mac).DeviceMac(deviceMac).Model(model).Vendor(vendor).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.SearchOrgOtherDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgOtherDeviceEvents`: ResponseEventsOtherDevicesSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesOthersAPI.SearchOrgOtherDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgOtherDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | site id | 
 **mac** | **string** | mac | 
 **deviceMac** | **string** | mac of attached device | 
 **model** | **string** | device model | 
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


## UpdateOrgOtherDevice

> UpdateOrgOtherDevice(ctx, orgId, deviceMac).OtherDeviceUpdate(otherDeviceUpdate).Execute()

updateOrgOtherDevice



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
	deviceMac := "0000000000ab" // string | 
	otherDeviceUpdate := *openapiclient.NewOtherDeviceUpdate() // OtherDeviceUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDevicesOthersAPI.UpdateOrgOtherDevice(context.Background(), orgId, deviceMac).OtherDeviceUpdate(otherDeviceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.UpdateOrgOtherDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgOtherDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **otherDeviceUpdate** | [**OtherDeviceUpdate**](OtherDeviceUpdate.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgOtherDevices

> UpdateOrgOtherDevices(ctx, orgId).OtherDeviceUpdateMulti(otherDeviceUpdateMulti).Execute()

updateOrgOtherDevices



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
	otherDeviceUpdateMulti := *openapiclient.NewOtherDeviceUpdateMulti(openapiclient.other_device_update_operation("")) // OtherDeviceUpdateMulti |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDevicesOthersAPI.UpdateOrgOtherDevices(context.Background(), orgId).OtherDeviceUpdateMulti(otherDeviceUpdateMulti).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesOthersAPI.UpdateOrgOtherDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgOtherDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **otherDeviceUpdateMulti** | [**OtherDeviceUpdateMulti**](OtherDeviceUpdateMulti.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

