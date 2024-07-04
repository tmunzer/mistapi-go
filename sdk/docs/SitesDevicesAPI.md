# \SitesDevicesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSiteDeviceImage**](SitesDevicesAPI.md#AddSiteDeviceImage) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/image{image_number} | addSiteDeviceImage
[**CountSiteDeviceConfigHistory**](SitesDevicesAPI.md#CountSiteDeviceConfigHistory) | **Get** /api/v1/sites/{site_id}/devices/config_history/count | countSiteDeviceConfigHistory
[**CountSiteDeviceEvents**](SitesDevicesAPI.md#CountSiteDeviceEvents) | **Get** /api/v1/sites/{site_id}/devices/events/count | countSiteDeviceEvents
[**CountSiteDeviceLastConfig**](SitesDevicesAPI.md#CountSiteDeviceLastConfig) | **Get** /api/v1/sites/{site_id}/devices/last_config/count | countSiteDeviceLastConfig
[**CountSiteDevices**](SitesDevicesAPI.md#CountSiteDevices) | **Get** /api/v1/sites/{site_id}/devices/count | countSiteDevices
[**DeleteSiteDevice**](SitesDevicesAPI.md#DeleteSiteDevice) | **Delete** /api/v1/sites/{site_id}/devices/{device_id} | deleteSiteDevice
[**DeleteSiteDeviceImage**](SitesDevicesAPI.md#DeleteSiteDeviceImage) | **Delete** /api/v1/sites/{site_id}/devices/{device_id}/image{image_number} | deleteSiteDeviceImage
[**ExportSiteDevices**](SitesDevicesAPI.md#ExportSiteDevices) | **Get** /api/v1/sites/{site_id}/devices/export | exportSiteDevices
[**GetSiteDevice**](SitesDevicesAPI.md#GetSiteDevice) | **Get** /api/v1/sites/{site_id}/devices/{device_id} | getSiteDevice
[**ImportSiteDevices**](SitesDevicesAPI.md#ImportSiteDevices) | **Post** /api/v1/sites/{site_id}/devices/import | importSiteDevices
[**ListSiteDevices**](SitesDevicesAPI.md#ListSiteDevices) | **Get** /api/v1/sites/{site_id}/devices | listSiteDevices
[**RestartSiteMultipleDevices**](SitesDevicesAPI.md#RestartSiteMultipleDevices) | **Post** /api/v1/sites/{site_id}/devices/restart | restartSiteMultipleDevices
[**SearchSiteDeviceConfigHistory**](SitesDevicesAPI.md#SearchSiteDeviceConfigHistory) | **Get** /api/v1/sites/{site_id}/devices/config_history/search | searchSiteDeviceConfigHistory
[**SearchSiteDeviceEvents**](SitesDevicesAPI.md#SearchSiteDeviceEvents) | **Get** /api/v1/sites/{site_id}/devices/events/search | searchSiteDeviceEvents
[**SearchSiteDeviceLastConfigs**](SitesDevicesAPI.md#SearchSiteDeviceLastConfigs) | **Get** /api/v1/sites/{site_id}/devices/last_config/search | searchSiteDeviceLastConfigs
[**SearchSiteDevices**](SitesDevicesAPI.md#SearchSiteDevices) | **Get** /api/v1/sites/{site_id}/devices/search | searchSiteDevices
[**UpdateSiteDevice**](SitesDevicesAPI.md#UpdateSiteDevice) | **Put** /api/v1/sites/{site_id}/devices/{device_id} | updateSiteDevice



## AddSiteDeviceImage

> AddSiteDeviceImage(ctx, siteId, deviceId, imageNumber).File(file).Json(json).Execute()

addSiteDeviceImage



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	imageNumber := int32(56) // int32 | 
	file := os.NewFile(1234, "some_file") // *os.File | binary file
	json := "json_example" // string | JSON string describing your upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesAPI.AddSiteDeviceImage(context.Background(), siteId, deviceId, imageNumber).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.AddSiteDeviceImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 
**imageNumber** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSiteDeviceImageRequest struct via the builder pattern


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


## CountSiteDeviceConfigHistory

> RepsonseCount CountSiteDeviceConfigHistory(ctx, siteId).Distinct(distinct).Mac(mac).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteDeviceConfigHistory



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
	distinct := "distinct_example" // string |  (optional)
	mac := "mac_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.CountSiteDeviceConfigHistory(context.Background(), siteId).Distinct(distinct).Mac(mac).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.CountSiteDeviceConfigHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteDeviceConfigHistory`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.CountSiteDeviceConfigHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteDeviceConfigHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | **string** |  | 
 **mac** | **string** |  | 
 **page** | **int32** |  | [default to 1]
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


## CountSiteDeviceEvents

> RepsonseCount CountSiteDeviceEvents(ctx, siteId).Distinct(distinct).Model(model).Type_(type_).TypeCode(typeCode).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteDeviceEvents



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
	distinct := openapiclient.site_device_events_count_distinct("") // SiteDeviceEventsCountDistinct |  (optional) (default to "model")
	model := "model_example" // string |  (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	typeCode := "typeCode_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.CountSiteDeviceEvents(context.Background(), siteId).Distinct(distinct).Model(model).Type_(type_).TypeCode(typeCode).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.CountSiteDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteDeviceEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.CountSiteDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteDeviceEventsCountDistinct**](SiteDeviceEventsCountDistinct.md) |  | [default to &quot;model&quot;]
 **model** | **string** |  | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **typeCode** | **string** |  | 
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


## CountSiteDeviceLastConfig

> RepsonseCount CountSiteDeviceLastConfig(ctx, siteId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteDeviceLastConfig



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
	distinct := openapiclient.site_device_last_config_count_distinct("") // SiteDeviceLastConfigCountDistinct |  (optional) (default to "mac")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.CountSiteDeviceLastConfig(context.Background(), siteId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.CountSiteDeviceLastConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteDeviceLastConfig`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.CountSiteDeviceLastConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteDeviceLastConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteDeviceLastConfigCountDistinct**](SiteDeviceLastConfigCountDistinct.md) |  | [default to &quot;mac&quot;]
 **page** | **int32** |  | [default to 1]
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


## CountSiteDevices

> RepsonseCount CountSiteDevices(ctx, siteId).Distinct(distinct).Hostname(hostname).Model(model).Mac(mac).Version(version).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).MapId(mapId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteDevices



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
	distinct := openapiclient.site_devices_count_distinct("") // SiteDevicesCountDistinct |  (optional) (default to "model")
	hostname := "hostname_example" // string |  (optional)
	model := "model_example" // string |  (optional)
	mac := "mac_example" // string |  (optional)
	version := "version_example" // string |  (optional)
	mxtunnelStatus := "mxtunnelStatus_example" // string |  (optional)
	mxedgeId := "mxedgeId_example" // string |  (optional)
	lldpSystemName := "lldpSystemName_example" // string |  (optional)
	lldpSystemDesc := "lldpSystemDesc_example" // string |  (optional)
	lldpPortId := "lldpPortId_example" // string |  (optional)
	lldpMgmtAddr := "lldpMgmtAddr_example" // string |  (optional)
	mapId := "mapId_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.CountSiteDevices(context.Background(), siteId).Distinct(distinct).Hostname(hostname).Model(model).Mac(mac).Version(version).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).MapId(mapId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.CountSiteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteDevices`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.CountSiteDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteDevicesCountDistinct**](SiteDevicesCountDistinct.md) |  | [default to &quot;model&quot;]
 **hostname** | **string** |  | 
 **model** | **string** |  | 
 **mac** | **string** |  | 
 **version** | **string** |  | 
 **mxtunnelStatus** | **string** |  | 
 **mxedgeId** | **string** |  | 
 **lldpSystemName** | **string** |  | 
 **lldpSystemDesc** | **string** |  | 
 **lldpPortId** | **string** |  | 
 **lldpMgmtAddr** | **string** |  | 
 **mapId** | **string** |  | 
 **page** | **int32** |  | [default to 1]
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


## DeleteSiteDevice

> DeleteSiteDevice(ctx, siteId, deviceId).Execute()

deleteSiteDevice



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesAPI.DeleteSiteDevice(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.DeleteSiteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteDeviceRequest struct via the builder pattern


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


## DeleteSiteDeviceImage

> DeleteSiteDeviceImage(ctx, siteId, deviceId, imageNumber).Execute()

deleteSiteDeviceImage



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	imageNumber := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesAPI.DeleteSiteDeviceImage(context.Background(), siteId, deviceId, imageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.DeleteSiteDeviceImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 
**imageNumber** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteDeviceImageRequest struct via the builder pattern


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


## ExportSiteDevices

> *os.File ExportSiteDevices(ctx, siteId).Execute()

exportSiteDevices



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.ExportSiteDevices(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.ExportSiteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportSiteDevices`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.ExportSiteDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSiteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDevice

> MistDevice GetSiteDevice(ctx, siteId, deviceId).Execute()

getSiteDevice



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.GetSiteDevice(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.GetSiteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDevice`: MistDevice
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.GetSiteDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MistDevice**](MistDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSiteDevices

> []ConfigDevice ImportSiteDevices(ctx, siteId).Ap(ap).Execute()

importSiteDevices



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
	ap := []openapiclient.Ap{*openapiclient.NewAp()} // []Ap |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.ImportSiteDevices(context.Background(), siteId).Ap(ap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.ImportSiteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportSiteDevices`: []ConfigDevice
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.ImportSiteDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSiteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ap** | [**[]Ap**](Ap.md) |  | 

### Return type

[**[]ConfigDevice**](ConfigDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteDevices

> []ConfigDevice ListSiteDevices(ctx, siteId).Type_(type_).Name(name).Page(page).Limit(limit).Execute()

listSiteDevices



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
	type_ := openapiclient.device_type_with_all("") // DeviceTypeWithAll |  (optional) (default to "ap")
	name := "name_example" // string |  (optional) (default to "")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.ListSiteDevices(context.Background(), siteId).Type_(type_).Name(name).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.ListSiteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteDevices`: []ConfigDevice
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.ListSiteDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceTypeWithAll**](DeviceTypeWithAll.md) |  | [default to &quot;ap&quot;]
 **name** | **string** |  | [default to &quot;&quot;]
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]ConfigDevice**](ConfigDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartSiteMultipleDevices

> RestartSiteMultipleDevices(ctx, siteId).UtilsDevicesRestartMulti(utilsDevicesRestartMulti).Execute()

restartSiteMultipleDevices



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
	utilsDevicesRestartMulti := *openapiclient.NewUtilsDevicesRestartMulti([]string{"DeviceIds_example"}) // UtilsDevicesRestartMulti | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesAPI.RestartSiteMultipleDevices(context.Background(), siteId).UtilsDevicesRestartMulti(utilsDevicesRestartMulti).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.RestartSiteMultipleDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartSiteMultipleDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **utilsDevicesRestartMulti** | [**UtilsDevicesRestartMulti**](UtilsDevicesRestartMulti.md) | Request Body | 

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


## SearchSiteDeviceConfigHistory

> ResponseConfigHistorySearch SearchSiteDeviceConfigHistory(ctx, siteId).Type_(type_).Mac(mac).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteDeviceConfigHistory



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
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	mac := "mac_example" // string | Device MAC Address (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.SearchSiteDeviceConfigHistory(context.Background(), siteId).Type_(type_).Mac(mac).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.SearchSiteDeviceConfigHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteDeviceConfigHistory`: ResponseConfigHistorySearch
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.SearchSiteDeviceConfigHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteDeviceConfigHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **mac** | **string** | Device MAC Address | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseConfigHistorySearch**](ResponseConfigHistorySearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteDeviceEvents

> ResponseEventsDevices SearchSiteDeviceEvents(ctx, siteId).Mac(mac).Model(model).Text(text).Timestamp(timestamp).Type_(type_).LastBy(lastBy).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteDeviceEvents



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
	mac := "mac_example" // string | device mac (optional)
	model := "model_example" // string | device model (optional)
	text := "text_example" // string | event message (optional)
	timestamp := "timestamp_example" // string | event time (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	lastBy := "port_id" // string | Return last/recent event for passed in field (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.SearchSiteDeviceEvents(context.Background(), siteId).Mac(mac).Model(model).Text(text).Timestamp(timestamp).Type_(type_).LastBy(lastBy).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.SearchSiteDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteDeviceEvents`: ResponseEventsDevices
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.SearchSiteDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | device mac | 
 **model** | **string** | device model | 
 **text** | **string** | event message | 
 **timestamp** | **string** | event time | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **lastBy** | **string** | Return last/recent event for passed in field | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseEventsDevices**](ResponseEventsDevices.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteDeviceLastConfigs

> ResponseConfigHistorySearch SearchSiteDeviceLastConfigs(ctx, siteId).Type_(type_).Mac(mac).Version(version).Name(name).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteDeviceLastConfigs



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
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	mac := "mac_example" // string |  (optional)
	version := "version_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.SearchSiteDeviceLastConfigs(context.Background(), siteId).Type_(type_).Mac(mac).Version(version).Name(name).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.SearchSiteDeviceLastConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteDeviceLastConfigs`: ResponseConfigHistorySearch
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.SearchSiteDeviceLastConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteDeviceLastConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **mac** | **string** |  | 
 **version** | **string** |  | 
 **name** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseConfigHistorySearch**](ResponseConfigHistorySearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteDevices

> ResponseDeviceSearch SearchSiteDevices(ctx, siteId).Hostname(hostname).Type_(type_).Model(model).Mac(mac).Version(version).PowerConstrained(powerConstrained).IpAddress(ipAddress).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).Band24Channel(band24Channel).Band5Channel(band5Channel).Band6Channel(band6Channel).Band24Bandwith(band24Bandwith).Band5Bandwith(band5Bandwith).Band6Bandwith(band6Bandwith).Eth0PortSpeed(eth0PortSpeed).Sort(sort).DescSort(descSort).Stats(stats).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteDevices



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
	hostname := "hostname_example" // string | partial / full hostname (optional)
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	model := "model_example" // string | device model (optional)
	mac := "mac_example" // string | device MAC (optional)
	version := "version_example" // string | version (optional)
	powerConstrained := true // bool | power_constrained (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	mxtunnelStatus := openapiclient.search_site_devices_mxtunnel_status("") // SearchSiteDevicesMxtunnelStatus | MxTunnel status, up / down (optional)
	mxedgeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Mist Edge id, if AP is connecting to a Mist Edge (optional)
	lldpSystemName := "lldpSystemName_example" // string | LLDP system name (optional)
	lldpSystemDesc := "lldpSystemDesc_example" // string | LLDP system description (optional)
	lldpPortId := "lldpPortId_example" // string | LLDP port id (optional)
	lldpMgmtAddr := "lldpMgmtAddr_example" // string | LLDP management ip address (optional)
	band24Channel := int32(56) // int32 | Channel of band_24 (optional)
	band5Channel := int32(56) // int32 | Channel of band_5 (optional)
	band6Channel := int32(56) // int32 | Channel of band_6 (optional)
	band24Bandwith := int32(56) // int32 | Bandwidth of band_24 (optional)
	band5Bandwith := int32(56) // int32 | Bandwidth of band_5 (optional)
	band6Bandwith := int32(56) // int32 | Bandwidth of band_6 (optional)
	eth0PortSpeed := int32(56) // int32 | Port speed of eth0 (optional)
	sort := openapiclient.search_site_devices_sort("") // SearchSiteDevicesSort | sort options (optional) (default to "timestamp")
	descSort := openapiclient.search_site_devices_desc_sort("") // SearchSiteDevicesDescSort | sort options in reverse order (optional)
	stats := true // bool | whether to return device stats (optional) (default to false)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.SearchSiteDevices(context.Background(), siteId).Hostname(hostname).Type_(type_).Model(model).Mac(mac).Version(version).PowerConstrained(powerConstrained).IpAddress(ipAddress).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).Band24Channel(band24Channel).Band5Channel(band5Channel).Band6Channel(band6Channel).Band24Bandwith(band24Bandwith).Band5Bandwith(band5Bandwith).Band6Bandwith(band6Bandwith).Eth0PortSpeed(eth0PortSpeed).Sort(sort).DescSort(descSort).Stats(stats).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.SearchSiteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteDevices`: ResponseDeviceSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.SearchSiteDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostname** | **string** | partial / full hostname | 
 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **model** | **string** | device model | 
 **mac** | **string** | device MAC | 
 **version** | **string** | version | 
 **powerConstrained** | **bool** | power_constrained | 
 **ipAddress** | **string** |  | 
 **mxtunnelStatus** | [**SearchSiteDevicesMxtunnelStatus**](SearchSiteDevicesMxtunnelStatus.md) | MxTunnel status, up / down | 
 **mxedgeId** | **string** | Mist Edge id, if AP is connecting to a Mist Edge | 
 **lldpSystemName** | **string** | LLDP system name | 
 **lldpSystemDesc** | **string** | LLDP system description | 
 **lldpPortId** | **string** | LLDP port id | 
 **lldpMgmtAddr** | **string** | LLDP management ip address | 
 **band24Channel** | **int32** | Channel of band_24 | 
 **band5Channel** | **int32** | Channel of band_5 | 
 **band6Channel** | **int32** | Channel of band_6 | 
 **band24Bandwith** | **int32** | Bandwidth of band_24 | 
 **band5Bandwith** | **int32** | Bandwidth of band_5 | 
 **band6Bandwith** | **int32** | Bandwidth of band_6 | 
 **eth0PortSpeed** | **int32** | Port speed of eth0 | 
 **sort** | [**SearchSiteDevicesSort**](SearchSiteDevicesSort.md) | sort options | [default to &quot;timestamp&quot;]
 **descSort** | [**SearchSiteDevicesDescSort**](SearchSiteDevicesDescSort.md) | sort options in reverse order | 
 **stats** | **bool** | whether to return device stats | [default to false]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseDeviceSearch**](ResponseDeviceSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteDevice

> MistDevice UpdateSiteDevice(ctx, siteId, deviceId).MistDevice(mistDevice).Execute()

updateSiteDevice



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mistDevice := openapiclient.mist_device{Ap: openapiclient.NewAp()} // MistDevice | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesAPI.UpdateSiteDevice(context.Background(), siteId, deviceId).MistDevice(mistDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesAPI.UpdateSiteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteDevice`: MistDevice
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesAPI.UpdateSiteDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mistDevice** | [**MistDevice**](MistDevice.md) | Request Body | 

### Return type

[**MistDevice**](MistDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

