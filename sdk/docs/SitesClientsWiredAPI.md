# \SitesClientsWiredAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWiredClients**](SitesClientsWiredAPI.md#CountSiteWiredClients) | **Get** /api/v1/sites/{site_id}/wired_clients/count | countSiteWiredClients
[**SearchSiteWiredClients**](SitesClientsWiredAPI.md#SearchSiteWiredClients) | **Get** /api/v1/sites/{site_id}/wired_clients/search | searchSiteWiredClients



## CountSiteWiredClients

> RepsonseCount CountSiteWiredClients(ctx, siteId).Distinct(distinct).Mac(mac).DeviceMac(deviceMac).PortId(portId).Vlan(vlan).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteWiredClients



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
	distinct := openapiclient.site_wired_clients_count_distinct("") // SiteWiredClientsCountDistinct |  (optional) (default to "mac")
	mac := "mac_example" // string | client mac (optional)
	deviceMac := "deviceMac_example" // string | device mac (optional)
	portId := "portId_example" // string | port id (optional)
	vlan := "vlan_example" // string | vlan (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWiredAPI.CountSiteWiredClients(context.Background(), siteId).Distinct(distinct).Mac(mac).DeviceMac(deviceMac).PortId(portId).Vlan(vlan).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWiredAPI.CountSiteWiredClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWiredClients`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWiredAPI.CountSiteWiredClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWiredClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteWiredClientsCountDistinct**](SiteWiredClientsCountDistinct.md) |  | [default to &quot;mac&quot;]
 **mac** | **string** | client mac | 
 **deviceMac** | **string** | device mac | 
 **portId** | **string** | port id | 
 **vlan** | **string** | vlan | 
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


## SearchSiteWiredClients

> SearchWiredClient SearchSiteWiredClients(ctx, siteId).DeviceMac(deviceMac).Mac(mac).Ip(ip).PortId(portId).Vlan(vlan).Manufacture(manufacture).Text(text).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteWiredClients



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
	deviceMac := "deviceMac_example" // string | device mac (optional)
	mac := "mac_example" // string | client mac (optional)
	ip := "ip_example" // string | client ip (optional)
	portId := "portId_example" // string | port id (optional)
	vlan := "vlan_example" // string | vlan (optional)
	manufacture := "manufacture_example" // string | manufacture (optional)
	text := "text_example" // string | single entry of hostname/mac (optional)
	nacruleId := "nacruleId_example" // string | nacrule_id (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWiredAPI.SearchSiteWiredClients(context.Background(), siteId).DeviceMac(deviceMac).Mac(mac).Ip(ip).PortId(portId).Vlan(vlan).Manufacture(manufacture).Text(text).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWiredAPI.SearchSiteWiredClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWiredClients`: SearchWiredClient
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWiredAPI.SearchSiteWiredClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWiredClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceMac** | **string** | device mac | 
 **mac** | **string** | client mac | 
 **ip** | **string** | client ip | 
 **portId** | **string** | port id | 
 **vlan** | **string** | vlan | 
 **manufacture** | **string** | manufacture | 
 **text** | **string** | single entry of hostname/mac | 
 **nacruleId** | **string** | nacrule_id | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**SearchWiredClient**](SearchWiredClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

