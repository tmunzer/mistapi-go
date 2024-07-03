# \OrgsClientsWanAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgWanClientEvents**](OrgsClientsWanAPI.md#CountOrgWanClientEvents) | **Get** /api/v1/orgs/{org_id}/wan_client/events/count | countOrgWanClientEvents
[**CountOrgWanClients**](OrgsClientsWanAPI.md#CountOrgWanClients) | **Get** /api/v1/orgs/{org_id}/wan_clients/count | countOrgWanClients
[**SearchOrgWanClientEvents**](OrgsClientsWanAPI.md#SearchOrgWanClientEvents) | **Get** /api/v1/orgs/{org_id}/wan_clients/events/search | searchOrgWanClientEvents
[**SearchOrgWanClients**](OrgsClientsWanAPI.md#SearchOrgWanClients) | **Get** /api/v1/orgs/{org_id}/wan_clients/search | searchOrgWanClients



## CountOrgWanClientEvents

> RepsonseCount CountOrgWanClientEvents(ctx, orgId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countOrgWanClientEvents



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
	distinct := openapiclient.org_wan_clients_events_count_distinct("") // OrgWanClientsEventsCountDistinct |  (optional) (default to "type")
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWanAPI.CountOrgWanClientEvents(context.Background(), orgId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWanAPI.CountOrgWanClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgWanClientEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWanAPI.CountOrgWanClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgWanClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgWanClientsEventsCountDistinct**](OrgWanClientsEventsCountDistinct.md) |  | [default to &quot;type&quot;]
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
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


## CountOrgWanClients

> RepsonseCount CountOrgWanClients(ctx, orgId).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

countOrgWanClients



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
	distinct := openapiclient.org_wan_clients_count_distinct("") // OrgWanClientsCountDistinct |  (optional) (default to "mac")
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWanAPI.CountOrgWanClients(context.Background(), orgId).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWanAPI.CountOrgWanClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgWanClients`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWanAPI.CountOrgWanClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgWanClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgWanClientsCountDistinct**](OrgWanClientsCountDistinct.md) |  | [default to &quot;mac&quot;]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

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


## SearchOrgWanClientEvents

> SearchEventsWanClient SearchOrgWanClientEvents(ctx, orgId).Type_(type_).Mac(mac).Hostname(hostname).Ip(ip).Mfg(mfg).NacruleId(nacruleId).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchOrgWanClientEvents



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
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	mac := "mac_example" // string | partial / full MAC address (optional)
	hostname := "hostname_example" // string | partial / full hostname (optional)
	ip := "ip_example" // string | client IP (optional)
	mfg := "mfg_example" // string | Manufacture (optional)
	nacruleId := "nacruleId_example" // string | nacrule_id (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWanAPI.SearchOrgWanClientEvents(context.Background(), orgId).Type_(type_).Mac(mac).Hostname(hostname).Ip(ip).Mfg(mfg).NacruleId(nacruleId).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWanAPI.SearchOrgWanClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgWanClientEvents`: SearchEventsWanClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWanAPI.SearchOrgWanClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgWanClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **mac** | **string** | partial / full MAC address | 
 **hostname** | **string** | partial / full hostname | 
 **ip** | **string** | client IP | 
 **mfg** | **string** | Manufacture | 
 **nacruleId** | **string** | nacrule_id | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**SearchEventsWanClient**](SearchEventsWanClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgWanClients

> SearchWanClient SearchOrgWanClients(ctx, orgId).Mac(mac).Hostname(hostname).Ip(ip).Network(network).Mfg(mfg).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

searchOrgWanClients



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
	mac := "mac_example" // string | partial / full MAC address (optional)
	hostname := "hostname_example" // string | partial / full hostname (optional)
	ip := "ip_example" // string | client IP (optional)
	network := "network_example" // string | network (optional)
	mfg := "mfg_example" // string | Manufacture (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWanAPI.SearchOrgWanClients(context.Background(), orgId).Mac(mac).Hostname(hostname).Ip(ip).Network(network).Mfg(mfg).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWanAPI.SearchOrgWanClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgWanClients`: SearchWanClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWanAPI.SearchOrgWanClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgWanClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | partial / full MAC address | 
 **hostname** | **string** | partial / full hostname | 
 **ip** | **string** | client IP | 
 **network** | **string** | network | 
 **mfg** | **string** | Manufacture | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**SearchWanClient**](SearchWanClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

