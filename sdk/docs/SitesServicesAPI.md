# \SitesServicesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteServicePathEvents**](SitesServicesAPI.md#CountSiteServicePathEvents) | **Get** /api/v1/sites/{site_id}/services/events/count | countSiteServicePathEvents
[**ListSiteServicesDerived**](SitesServicesAPI.md#ListSiteServicesDerived) | **Get** /api/v1/sites/{site_id}/services/derived | listSiteServicesDerived
[**SearchSiteServicePathEvents**](SitesServicesAPI.md#SearchSiteServicePathEvents) | **Get** /api/v1/sites/{site_id}/services/events/search | searchSiteServicePathEvents



## CountSiteServicePathEvents

> RepsonseCount CountSiteServicePathEvents(ctx, siteId).Distinct(distinct).Type_(type_).Text(text).VpnName(vpnName).VpnPath(vpnPath).Policy(policy).PortId(portId).Model(model).Version(version).Timestamp(timestamp).Mac(mac).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countSiteServicePathEvents



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
	distinct := openapiclient.site_service_events_count_distinct("") // SiteServiceEventsCountDistinct |  (optional) (default to "type")
	type_ := "type__example" // string | Event type, e.g. GW_SERVICE_PATH_DOWN (optional)
	text := "text_example" // string | Description of the event including the reason it is triggered (optional)
	vpnName := "vpnName_example" // string | Peer name (optional)
	vpnPath := "vpnPath_example" // string | Peer path name (optional)
	policy := "policy_example" // string | Service policy associated with that specific path (optional)
	portId := "portId_example" // string | Network interface (optional)
	model := "model_example" // string | Device model (optional)
	version := "version_example" // string | Device firmware version (optional)
	timestamp := float32(8.14) // float32 | Start time, in epoch (optional)
	mac := "mac_example" // string | MAC address (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesServicesAPI.CountSiteServicePathEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).Text(text).VpnName(vpnName).VpnPath(vpnPath).Policy(policy).PortId(portId).Model(model).Version(version).Timestamp(timestamp).Mac(mac).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesServicesAPI.CountSiteServicePathEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteServicePathEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesServicesAPI.CountSiteServicePathEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteServicePathEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteServiceEventsCountDistinct**](SiteServiceEventsCountDistinct.md) |  | [default to &quot;type&quot;]
 **type_** | **string** | Event type, e.g. GW_SERVICE_PATH_DOWN | 
 **text** | **string** | Description of the event including the reason it is triggered | 
 **vpnName** | **string** | Peer name | 
 **vpnPath** | **string** | Peer path name | 
 **policy** | **string** | Service policy associated with that specific path | 
 **portId** | **string** | Network interface | 
 **model** | **string** | Device model | 
 **version** | **string** | Device firmware version | 
 **timestamp** | **float32** | Start time, in epoch | 
 **mac** | **string** | MAC address | 
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


## ListSiteServicesDerived

> []Service ListSiteServicesDerived(ctx, siteId).Resolve(resolve).Execute()

listSiteServicesDerived



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
	resolve := true // bool | whether resolve the site variables (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesServicesAPI.ListSiteServicesDerived(context.Background(), siteId).Resolve(resolve).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesServicesAPI.ListSiteServicesDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteServicesDerived`: []Service
	fmt.Fprintf(os.Stdout, "Response from `SitesServicesAPI.ListSiteServicesDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteServicesDerivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **bool** | whether resolve the site variables | [default to false]

### Return type

[**[]Service**](Service.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteServicePathEvents

> ResponseEventsPathSearch SearchSiteServicePathEvents(ctx, siteId).Type_(type_).Text(text).VpnName(vpnName).VpnPath(vpnPath).Policy(policy).PortId(portId).Model(model).Version(version).Timestamp(timestamp).Mac(mac).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchSiteServicePathEvents



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
	type_ := "type__example" // string | Event type, e.g. GW_SERVICE_PATH_DOWN (optional)
	text := "text_example" // string | Description of the event including the reason it is triggered (optional)
	vpnName := "vpnName_example" // string | Peer name (optional)
	vpnPath := "vpnPath_example" // string | Peer path name (optional)
	policy := "policy_example" // string | Service policy associated with that specific path (optional)
	portId := "portId_example" // string | Network interface (optional)
	model := "model_example" // string | Device model (optional)
	version := "version_example" // string | Device firmware version (optional)
	timestamp := float32(8.14) // float32 | Start time, in epoch (optional)
	mac := "mac_example" // string | MAC address (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesServicesAPI.SearchSiteServicePathEvents(context.Background(), siteId).Type_(type_).Text(text).VpnName(vpnName).VpnPath(vpnPath).Policy(policy).PortId(portId).Model(model).Version(version).Timestamp(timestamp).Mac(mac).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesServicesAPI.SearchSiteServicePathEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteServicePathEvents`: ResponseEventsPathSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesServicesAPI.SearchSiteServicePathEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteServicePathEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Event type, e.g. GW_SERVICE_PATH_DOWN | 
 **text** | **string** | Description of the event including the reason it is triggered | 
 **vpnName** | **string** | Peer name | 
 **vpnPath** | **string** | Peer path name | 
 **policy** | **string** | Service policy associated with that specific path | 
 **portId** | **string** | Network interface | 
 **model** | **string** | Device model | 
 **version** | **string** | Device firmware version | 
 **timestamp** | **float32** | Start time, in epoch | 
 **mac** | **string** | MAC address | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**ResponseEventsPathSearch**](ResponseEventsPathSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

