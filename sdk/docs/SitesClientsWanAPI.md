# \SitesClientsWanAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWanClientEvents**](SitesClientsWanAPI.md#CountSiteWanClientEvents) | **Get** /api/v1/sites/{site_id}/wan_client/events/count | countSiteWanClientEvents
[**CountSiteWanClients**](SitesClientsWanAPI.md#CountSiteWanClients) | **Get** /api/v1/sites/{site_id}/wan_clients/count | countSiteWanClients
[**SearchSiteWanClientEvents**](SitesClientsWanAPI.md#SearchSiteWanClientEvents) | **Get** /api/v1/sites/{site_id}/wan_clients/events/search | searchSiteWanClientEvents
[**SearchSiteWanClients**](SitesClientsWanAPI.md#SearchSiteWanClients) | **Get** /api/v1/sites/{site_id}/wan_clients/search | searchSiteWanClients



## CountSiteWanClientEvents

> RepsonseCount CountSiteWanClientEvents(ctx, siteId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countSiteWanClientEvents



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
	distinct := openapiclient.site_wan_client_events_distinct("") // SiteWanClientEventsDistinct |  (optional) (default to "type")
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWanAPI.CountSiteWanClientEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWanAPI.CountSiteWanClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWanClientEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWanAPI.CountSiteWanClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWanClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteWanClientEventsDistinct**](SiteWanClientEventsDistinct.md) |  | [default to &quot;type&quot;]
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


## CountSiteWanClients

> RepsonseCount CountSiteWanClients(ctx, siteId).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

countSiteWanClients



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
	distinct := openapiclient.site_wan_clients_count_distinct("") // SiteWanClientsCountDistinct |  (optional) (default to "mac")
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWanAPI.CountSiteWanClients(context.Background(), siteId).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWanAPI.CountSiteWanClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWanClients`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWanAPI.CountSiteWanClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWanClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteWanClientsCountDistinct**](SiteWanClientsCountDistinct.md) |  | [default to &quot;mac&quot;]
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


## SearchSiteWanClientEvents

> SearchEventsWanClient SearchSiteWanClientEvents(ctx, siteId).Type_(type_).Mac(mac).Hostname(hostname).Ip(ip).Mfg(mfg).NacruleId(nacruleId).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchSiteWanClientEvents



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
	resp, r, err := apiClient.SitesClientsWanAPI.SearchSiteWanClientEvents(context.Background(), siteId).Type_(type_).Mac(mac).Hostname(hostname).Ip(ip).Mfg(mfg).NacruleId(nacruleId).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWanAPI.SearchSiteWanClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWanClientEvents`: SearchEventsWanClient
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWanAPI.SearchSiteWanClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWanClientEventsRequest struct via the builder pattern


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


## SearchSiteWanClients

> SearchWanClient SearchSiteWanClients(ctx, siteId).Mac(mac).Hostname(hostname).Ip(ip).Mfg(mfg).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

searchSiteWanClients



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
	mac := "mac_example" // string | partial / full MAC address (optional)
	hostname := "hostname_example" // string | partial / full hostname (optional)
	ip := "ip_example" // string | client IP (optional)
	mfg := "mfg_example" // string | Manufacture (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWanAPI.SearchSiteWanClients(context.Background(), siteId).Mac(mac).Hostname(hostname).Ip(ip).Mfg(mfg).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWanAPI.SearchSiteWanClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWanClients`: SearchWanClient
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWanAPI.SearchSiteWanClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWanClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | partial / full MAC address | 
 **hostname** | **string** | partial / full hostname | 
 **ip** | **string** | client IP | 
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

