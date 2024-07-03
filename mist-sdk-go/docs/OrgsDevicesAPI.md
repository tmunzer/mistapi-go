# \OrgsDevicesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgDeviceEvents**](OrgsDevicesAPI.md#CountOrgDeviceEvents) | **Get** /api/v1/orgs/{org_id}/devices/events/count | countOrgDeviceEvents
[**CountOrgDeviceLastConfigs**](OrgsDevicesAPI.md#CountOrgDeviceLastConfigs) | **Get** /api/v1/orgs/{org_id}/devices/last_config/count | countOrgDeviceLastConfigs
[**CountOrgDevices**](OrgsDevicesAPI.md#CountOrgDevices) | **Get** /api/v1/orgs/{org_id}/devices/count | countOrgDevices
[**GetOrgJuniperDevicesCommand**](OrgsDevicesAPI.md#GetOrgJuniperDevicesCommand) | **Get** /api/v1/orgs/{org_id}/ocdevices/outbound_ssh_cmd | getOrgJuniperDevicesCommand
[**ListOrgApsMacs**](OrgsDevicesAPI.md#ListOrgApsMacs) | **Get** /api/v1/orgs/{org_id}/devices/radio_macs | listOrgApsMacs
[**ListOrgDevices**](OrgsDevicesAPI.md#ListOrgDevices) | **Get** /api/v1/orgs/{org_id}/devices | listOrgDevices
[**ListOrgDevicesStats**](OrgsDevicesAPI.md#ListOrgDevicesStats) | **Get** /api/v1/orgs/{org_id}/stats/devices | listOrgDevicesStats
[**SearchOrgDeviceEvents**](OrgsDevicesAPI.md#SearchOrgDeviceEvents) | **Get** /api/v1/orgs/{org_id}/devices/events/search | searchOrgDeviceEvents
[**SearchOrgDeviceLastConfigs**](OrgsDevicesAPI.md#SearchOrgDeviceLastConfigs) | **Get** /api/v1/orgs/{org_id}/devices/last_config/search | searchOrgDeviceLastConfigs
[**SearchOrgDevices**](OrgsDevicesAPI.md#SearchOrgDevices) | **Get** /api/v1/orgs/{org_id}/devices/search | searchOrgDevices



## CountOrgDeviceEvents

> RepsonseCount CountOrgDeviceEvents(ctx, orgId).Distinct(distinct).SiteId(siteId).Ap(ap).Apfw(apfw).Model(model).Text(text).Timestamp(timestamp).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgDeviceEvents



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	distinct := openapiclient.org_devices_events_count_distinct("") // OrgDevicesEventsCountDistinct |  (optional) (default to "model")
	siteId := "siteId_example" // string | site id (optional)
	ap := "ap_example" // string | AP mac (optional)
	apfw := "apfw_example" // string | AP Firmware (optional)
	model := "model_example" // string | device model (optional)
	text := "text_example" // string | event message (optional)
	timestamp := "timestamp_example" // string | event time (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.CountOrgDeviceEvents(context.Background(), orgId).Distinct(distinct).SiteId(siteId).Ap(ap).Apfw(apfw).Model(model).Text(text).Timestamp(timestamp).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.CountOrgDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgDeviceEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.CountOrgDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgDevicesEventsCountDistinct**](OrgDevicesEventsCountDistinct.md) |  | [default to &quot;model&quot;]
 **siteId** | **string** | site id | 
 **ap** | **string** | AP mac | 
 **apfw** | **string** | AP Firmware | 
 **model** | **string** | device model | 
 **text** | **string** | event message | 
 **timestamp** | **string** | event time | 
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


## CountOrgDeviceLastConfigs

> RepsonseCount CountOrgDeviceLastConfigs(ctx, orgId).Type_(type_).Distinct(distinct).Start(start).End(end).Limit(limit).Execute()

countOrgDeviceLastConfigs



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	distinct := openapiclient.org_devices_last_configs_count_distinct("") // OrgDevicesLastConfigsCountDistinct |  (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.CountOrgDeviceLastConfigs(context.Background(), orgId).Type_(type_).Distinct(distinct).Start(start).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.CountOrgDeviceLastConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgDeviceLastConfigs`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.CountOrgDeviceLastConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgDeviceLastConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **distinct** | [**OrgDevicesLastConfigsCountDistinct**](OrgDevicesLastConfigsCountDistinct.md) |  | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
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


## CountOrgDevices

> RepsonseCount CountOrgDevices(ctx, orgId).Distinct(distinct).Hostname(hostname).SiteId(siteId).Model(model).Mac(mac).Version(version).IpAddress(ipAddress).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgDevices



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	distinct := openapiclient.org_devices_count_distinct("") // OrgDevicesCountDistinct |  (optional) (default to "model")
	hostname := "hostname_example" // string | partial / full hostname (optional)
	siteId := "siteId_example" // string | site id (optional)
	model := "model_example" // string | device model (optional)
	mac := "mac_example" // string | AP mac (optional)
	version := "version_example" // string | version (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	mxtunnelStatus := openapiclient.count_org_devices_mxtunnel_status("") // CountOrgDevicesMxtunnelStatus | MxTunnel status, up / down (optional)
	mxedgeId := "mxedgeId_example" // string | Mist Edge id, if AP is connecting to a Mist Edge (optional)
	lldpSystemName := "lldpSystemName_example" // string | LLDP system name (optional)
	lldpSystemDesc := "lldpSystemDesc_example" // string | LLDP system description (optional)
	lldpPortId := "lldpPortId_example" // string | LLDP port id (optional)
	lldpMgmtAddr := "lldpMgmtAddr_example" // string | LLDP management ip address (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.CountOrgDevices(context.Background(), orgId).Distinct(distinct).Hostname(hostname).SiteId(siteId).Model(model).Mac(mac).Version(version).IpAddress(ipAddress).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.CountOrgDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgDevices`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.CountOrgDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgDevicesCountDistinct**](OrgDevicesCountDistinct.md) |  | [default to &quot;model&quot;]
 **hostname** | **string** | partial / full hostname | 
 **siteId** | **string** | site id | 
 **model** | **string** | device model | 
 **mac** | **string** | AP mac | 
 **version** | **string** | version | 
 **ipAddress** | **string** |  | 
 **mxtunnelStatus** | [**CountOrgDevicesMxtunnelStatus**](CountOrgDevicesMxtunnelStatus.md) | MxTunnel status, up / down | 
 **mxedgeId** | **string** | Mist Edge id, if AP is connecting to a Mist Edge | 
 **lldpSystemName** | **string** | LLDP system name | 
 **lldpSystemDesc** | **string** | LLDP system description | 
 **lldpPortId** | **string** | LLDP port id | 
 **lldpMgmtAddr** | **string** | LLDP management ip address | 
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


## GetOrgJuniperDevicesCommand

> ResponseDeviceConfigCmd GetOrgJuniperDevicesCommand(ctx, orgId).SiteId(siteId).Execute()

getOrgJuniperDevicesCommand



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	siteId := "siteId_example" // string | site_id would be used for proxy config check of the site and automatic site assignment (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.GetOrgJuniperDevicesCommand(context.Background(), orgId).SiteId(siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.GetOrgJuniperDevicesCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgJuniperDevicesCommand`: ResponseDeviceConfigCmd
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.GetOrgJuniperDevicesCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgJuniperDevicesCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | site_id would be used for proxy config check of the site and automatic site assignment | 

### Return type

[**ResponseDeviceConfigCmd**](ResponseDeviceConfigCmd.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgApsMacs

> []ApRadioMac ListOrgApsMacs(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgApsMacs



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.ListOrgApsMacs(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.ListOrgApsMacs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgApsMacs`: []ApRadioMac
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.ListOrgApsMacs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgApsMacsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]ApRadioMac**](ApRadioMac.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgDevices

> ResponseOrgDevices ListOrgDevices(ctx, orgId).Execute()

listOrgDevices



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.ListOrgDevices(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.ListOrgDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgDevices`: ResponseOrgDevices
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.ListOrgDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseOrgDevices**](ResponseOrgDevices.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgDevicesStats

> []StatsDevice ListOrgDevicesStats(ctx, orgId).Type_(type_).Status(status).SiteId(siteId).Mac(mac).EvpntopoId(evpntopoId).EvpnUnused(evpnUnused).Fields(fields).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listOrgDevicesStats



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	type_ := openapiclient.device_type_with_all("") // DeviceTypeWithAll |  (optional) (default to "ap")
	status := openapiclient.device_status("") // DeviceStatus |  (optional) (default to "all")
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	mac := "mac_example" // string |  (optional)
	evpntopoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | EVPN Topology ID (optional)
	evpnUnused := "evpnUnused_example" // string | if `evpn_unused`==`true`, find EVPN eligible switches which don’t belong to any EVPN Topology yet (optional)
	fields := "field1,field2" // string | list of additional fields requests, comma separeted, or `fields=*` for all of them (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.ListOrgDevicesStats(context.Background(), orgId).Type_(type_).Status(status).SiteId(siteId).Mac(mac).EvpntopoId(evpntopoId).EvpnUnused(evpnUnused).Fields(fields).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.ListOrgDevicesStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgDevicesStats`: []StatsDevice
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.ListOrgDevicesStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgDevicesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceTypeWithAll**](DeviceTypeWithAll.md) |  | [default to &quot;ap&quot;]
 **status** | [**DeviceStatus**](DeviceStatus.md) |  | [default to &quot;all&quot;]
 **siteId** | **string** |  | 
 **mac** | **string** |  | 
 **evpntopoId** | **string** | EVPN Topology ID | 
 **evpnUnused** | **string** | if &#x60;evpn_unused&#x60;&#x3D;&#x3D;&#x60;true&#x60;, find EVPN eligible switches which don’t belong to any EVPN Topology yet | 
 **fields** | **string** | list of additional fields requests, comma separeted, or &#x60;fields&#x3D;*&#x60; for all of them | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]StatsDevice**](StatsDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgDeviceEvents

> ResponseDeviceEventsSearch SearchOrgDeviceEvents(ctx, orgId).Mac(mac).Model(model).DeviceType(deviceType).Text(text).Timestamp(timestamp).Type_(type_).LastBy(lastBy).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgDeviceEvents



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mac := "mac_example" // string | device mac (optional)
	model := "model_example" // string | device model (optional)
	deviceType := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
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
	resp, r, err := apiClient.OrgsDevicesAPI.SearchOrgDeviceEvents(context.Background(), orgId).Mac(mac).Model(model).DeviceType(deviceType).Text(text).Timestamp(timestamp).Type_(type_).LastBy(lastBy).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.SearchOrgDeviceEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgDeviceEvents`: ResponseDeviceEventsSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.SearchOrgDeviceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgDeviceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | device mac | 
 **model** | **string** | device model | 
 **deviceType** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **text** | **string** | event message | 
 **timestamp** | **string** | event time | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **lastBy** | **string** | Return last/recent event for passed in field | 
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


## SearchOrgDeviceLastConfigs

> ResponseConfigHistorySearch SearchOrgDeviceLastConfigs(ctx, orgId).Type_(type_).Mac(mac).Name(name).Version(version).Start(start).End(end).Limit(limit).Duration(duration).Execute()

searchOrgDeviceLastConfigs



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	mac := "mac_example" // string | Device MAC address (optional)
	name := "name_example" // string | Devices Name (optional)
	version := "version_example" // string | Device Version (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.SearchOrgDeviceLastConfigs(context.Background(), orgId).Type_(type_).Mac(mac).Name(name).Version(version).Start(start).End(end).Limit(limit).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.SearchOrgDeviceLastConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgDeviceLastConfigs`: ResponseConfigHistorySearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.SearchOrgDeviceLastConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgDeviceLastConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **mac** | **string** | Device MAC address | 
 **name** | **string** | Devices Name | 
 **version** | **string** | Device Version | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **limit** | **int32** |  | [default to 100]
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


## SearchOrgDevices

> ResponseDeviceSearch SearchOrgDevices(ctx, orgId).Hostname(hostname).SiteId(siteId).Model(model).Mac(mac).Version(version).PowerConstrained(powerConstrained).IpAddress(ipAddress).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).Band24Bandwith(band24Bandwith).Band5Bandwith(band5Bandwith).Band6Bandwith(band6Bandwith).Band24Channel(band24Channel).Band5Channel(band5Channel).Band6Channel(band6Channel).Eth0PortSpeed(eth0PortSpeed).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgDevices



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	hostname := "hostname_example" // string | partial / full hostname (optional)
	siteId := "siteId_example" // string | site id (optional)
	model := "model_example" // string | device model (optional)
	mac := "mac_example" // string | AP mac (optional)
	version := "version_example" // string | version (optional)
	powerConstrained := true // bool | power_constrained (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	mxtunnelStatus := openapiclient.search_org_devices_mxtunnel_status("") // SearchOrgDevicesMxtunnelStatus | MxTunnel status, up / down (optional)
	mxedgeId := "mxedgeId_example" // string | Mist Edge id, if AP is connecting to a Mist Edge (optional)
	lldpSystemName := "lldpSystemName_example" // string | LLDP system name (optional)
	lldpSystemDesc := "lldpSystemDesc_example" // string | LLDP system description (optional)
	lldpPortId := "lldpPortId_example" // string | LLDP port id (optional)
	lldpMgmtAddr := "lldpMgmtAddr_example" // string | LLDP management ip address (optional)
	band24Bandwith := int32(56) // int32 | Bandwith of band_24 (optional)
	band5Bandwith := int32(56) // int32 | Bandwith of band_5 (optional)
	band6Bandwith := int32(56) // int32 | Bandwith of band_6 (optional)
	band24Channel := int32(56) // int32 | Channel of band_24 (optional)
	band5Channel := int32(56) // int32 | Channel of band_5 (optional)
	band6Channel := int32(56) // int32 | Channel of band_6 (optional)
	eth0PortSpeed := int32(56) // int32 | Port speed of eth0 (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesAPI.SearchOrgDevices(context.Background(), orgId).Hostname(hostname).SiteId(siteId).Model(model).Mac(mac).Version(version).PowerConstrained(powerConstrained).IpAddress(ipAddress).MxtunnelStatus(mxtunnelStatus).MxedgeId(mxedgeId).LldpSystemName(lldpSystemName).LldpSystemDesc(lldpSystemDesc).LldpPortId(lldpPortId).LldpMgmtAddr(lldpMgmtAddr).Band24Bandwith(band24Bandwith).Band5Bandwith(band5Bandwith).Band6Bandwith(band6Bandwith).Band24Channel(band24Channel).Band5Channel(band5Channel).Band6Channel(band6Channel).Eth0PortSpeed(eth0PortSpeed).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesAPI.SearchOrgDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgDevices`: ResponseDeviceSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesAPI.SearchOrgDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostname** | **string** | partial / full hostname | 
 **siteId** | **string** | site id | 
 **model** | **string** | device model | 
 **mac** | **string** | AP mac | 
 **version** | **string** | version | 
 **powerConstrained** | **bool** | power_constrained | 
 **ipAddress** | **string** |  | 
 **mxtunnelStatus** | [**SearchOrgDevicesMxtunnelStatus**](SearchOrgDevicesMxtunnelStatus.md) | MxTunnel status, up / down | 
 **mxedgeId** | **string** | Mist Edge id, if AP is connecting to a Mist Edge | 
 **lldpSystemName** | **string** | LLDP system name | 
 **lldpSystemDesc** | **string** | LLDP system description | 
 **lldpPortId** | **string** | LLDP port id | 
 **lldpMgmtAddr** | **string** | LLDP management ip address | 
 **band24Bandwith** | **int32** | Bandwith of band_24 | 
 **band5Bandwith** | **int32** | Bandwith of band_5 | 
 **band6Bandwith** | **int32** | Bandwith of band_6 | 
 **band24Channel** | **int32** | Channel of band_24 | 
 **band5Channel** | **int32** | Channel of band_5 | 
 **band6Channel** | **int32** | Channel of band_6 | 
 **eth0PortSpeed** | **int32** | Port speed of eth0 | 
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

