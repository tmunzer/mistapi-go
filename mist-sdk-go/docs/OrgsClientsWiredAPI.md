# \OrgsClientsWiredAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgWiredClients**](OrgsClientsWiredAPI.md#CountOrgWiredClients) | **Get** /api/v1/orgs/{org_id}/wired_clients/count | countOrgWiredClients
[**SearchOrgWiredClients**](OrgsClientsWiredAPI.md#SearchOrgWiredClients) | **Get** /api/v1/orgs/{org_id}/wired_clients/search | searchOrgWiredClients



## CountOrgWiredClients

> RepsonseCount CountOrgWiredClients(ctx, orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgWiredClients



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
	distinct := openapiclient.org_wired_clients_count_distinct("") // OrgWiredClientsCountDistinct |  (optional) (default to "mac")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWiredAPI.CountOrgWiredClients(context.Background(), orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWiredAPI.CountOrgWiredClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgWiredClients`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWiredAPI.CountOrgWiredClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgWiredClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgWiredClientsCountDistinct**](OrgWiredClientsCountDistinct.md) |  | [default to &quot;mac&quot;]
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


## SearchOrgWiredClients

> SearchWiredClient SearchOrgWiredClients(ctx, orgId).SiteId(siteId).DeviceMac(deviceMac).Mac(mac).PortId(portId).Vlan(vlan).Ip(ip).Manufacture(manufacture).Text(text).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgWiredClients



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
	siteId := "siteId_example" // string | Site ID (optional)
	deviceMac := "deviceMac_example" // string | device mac (optional)
	mac := "mac_example" // string | client mac (optional)
	portId := "portId_example" // string | port id (optional)
	vlan := int32(56) // int32 | vlan (optional)
	ip := "ip_example" // string | ip (optional)
	manufacture := "manufacture_example" // string | client manufacturer (optional)
	text := "text_example" // string | single entry of hostname/mac (optional)
	nacruleId := "nacruleId_example" // string | nacrule_id (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWiredAPI.SearchOrgWiredClients(context.Background(), orgId).SiteId(siteId).DeviceMac(deviceMac).Mac(mac).PortId(portId).Vlan(vlan).Ip(ip).Manufacture(manufacture).Text(text).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWiredAPI.SearchOrgWiredClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgWiredClients`: SearchWiredClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWiredAPI.SearchOrgWiredClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgWiredClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | Site ID | 
 **deviceMac** | **string** | device mac | 
 **mac** | **string** | client mac | 
 **portId** | **string** | port id | 
 **vlan** | **int32** | vlan | 
 **ip** | **string** | ip | 
 **manufacture** | **string** | client manufacturer | 
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

