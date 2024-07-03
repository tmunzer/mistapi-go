# \SitesRoguesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteRogueEvents**](SitesRoguesAPI.md#CountSiteRogueEvents) | **Get** /api/v1/sites/{site_id}/rogues/events/count | countSiteRogueEvents
[**GetSiteRogueAP**](SitesRoguesAPI.md#GetSiteRogueAP) | **Get** /api/v1/sites/{site_id}/rogues/{rogue_bssid} | getSiteRogueAP
[**ListSiteRogueAPs**](SitesRoguesAPI.md#ListSiteRogueAPs) | **Get** /api/v1/sites/{site_id}/insights/rogues | listSiteRogueAPs
[**ListSiteRogueClients**](SitesRoguesAPI.md#ListSiteRogueClients) | **Get** /api/v1/sites/{site_id}/insights/rogues/clients | listSiteRogueClients
[**SearchSiteRogueEvents**](SitesRoguesAPI.md#SearchSiteRogueEvents) | **Get** /api/v1/sites/{site_id}/rogues/events/search | searchSiteRogueEvents



## CountSiteRogueEvents

> RepsonseCount CountSiteRogueEvents(ctx, siteId).Distinct(distinct).Type_(type_).Ssid(ssid).Bssid(bssid).ApMac(apMac).Channel(channel).SeenOnLan(seenOnLan).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteRogueEvents



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
	distinct := openapiclient.site_rogue_events_count_distinct("") // SiteRogueEventsCountDistinct |  (optional) (default to "bssid")
	type_ := openapiclient.rogue_type("") // RogueType |  (optional)
	ssid := "ssid_example" // string | ssid of the network detected as threat (optional)
	bssid := "bssid_example" // string | bssid of the network detected as threat (optional)
	apMac := "apMac_example" // string | mac of the device that had strongest signal strength for ssid/bssid pair (optional)
	channel := "channel_example" // string | channel over which ap_mac heard ssid/bssid pair (optional)
	seenOnLan := true // bool | whether the reporting AP see a wireless client (on LAN) connecting to it (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRoguesAPI.CountSiteRogueEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).Ssid(ssid).Bssid(bssid).ApMac(apMac).Channel(channel).SeenOnLan(seenOnLan).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRoguesAPI.CountSiteRogueEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteRogueEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesRoguesAPI.CountSiteRogueEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteRogueEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteRogueEventsCountDistinct**](SiteRogueEventsCountDistinct.md) |  | [default to &quot;bssid&quot;]
 **type_** | [**RogueType**](RogueType.md) |  | 
 **ssid** | **string** | ssid of the network detected as threat | 
 **bssid** | **string** | bssid of the network detected as threat | 
 **apMac** | **string** | mac of the device that had strongest signal strength for ssid/bssid pair | 
 **channel** | **string** | channel over which ap_mac heard ssid/bssid pair | 
 **seenOnLan** | **bool** | whether the reporting AP see a wireless client (on LAN) connecting to it | 
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


## GetSiteRogueAP

> RogueDetails GetSiteRogueAP(ctx, siteId, rogueBssid).Execute()

getSiteRogueAP



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
	rogueBssid := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRoguesAPI.GetSiteRogueAP(context.Background(), siteId, rogueBssid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRoguesAPI.GetSiteRogueAP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteRogueAP`: RogueDetails
	fmt.Fprintf(os.Stdout, "Response from `SitesRoguesAPI.GetSiteRogueAP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rogueBssid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRogueAPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RogueDetails**](RogueDetails.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteRogueAPs

> ResponseInsightRogue ListSiteRogueAPs(ctx, siteId).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()

listSiteRogueAPs



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
	type_ := openapiclient.rogue_type("") // RogueType |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	interval := "10m" // string | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRoguesAPI.ListSiteRogueAPs(context.Background(), siteId).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRoguesAPI.ListSiteRogueAPs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteRogueAPs`: ResponseInsightRogue
	fmt.Fprintf(os.Stdout, "Response from `SitesRoguesAPI.ListSiteRogueAPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteRogueAPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**RogueType**](RogueType.md) |  | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **interval** | **string** | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 

### Return type

[**ResponseInsightRogue**](ResponseInsightRogue.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteRogueClients

> ResponseInsightRogueClient ListSiteRogueClients(ctx, siteId).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()

listSiteRogueClients



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
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	interval := "10m" // string | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRoguesAPI.ListSiteRogueClients(context.Background(), siteId).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRoguesAPI.ListSiteRogueClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteRogueClients`: ResponseInsightRogueClient
	fmt.Fprintf(os.Stdout, "Response from `SitesRoguesAPI.ListSiteRogueClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteRogueClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **interval** | **string** | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 

### Return type

[**ResponseInsightRogueClient**](ResponseInsightRogueClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteRogueEvents

> ResponseEventsRogueSearch SearchSiteRogueEvents(ctx, siteId).Type_(type_).Ssid(ssid).Bssid(bssid).ApMac(apMac).Channel(channel).SeenOnLan(seenOnLan).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteRogueEvents



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
	type_ := openapiclient.rogue_type("") // RogueType |  (optional)
	ssid := "ssid_example" // string | ssid of the network detected as threat (optional)
	bssid := "bssid_example" // string | bssid of the network detected as threat (optional)
	apMac := "apMac_example" // string | mac of the device that had strongest signal strength for ssid/bssid pair (optional)
	channel := int32(56) // int32 | channel over which ap_mac heard ssid/bssid pair (optional)
	seenOnLan := true // bool | whether the reporting AP see a wireless client (on LAN) connecting to it (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRoguesAPI.SearchSiteRogueEvents(context.Background(), siteId).Type_(type_).Ssid(ssid).Bssid(bssid).ApMac(apMac).Channel(channel).SeenOnLan(seenOnLan).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRoguesAPI.SearchSiteRogueEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteRogueEvents`: ResponseEventsRogueSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesRoguesAPI.SearchSiteRogueEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteRogueEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**RogueType**](RogueType.md) |  | 
 **ssid** | **string** | ssid of the network detected as threat | 
 **bssid** | **string** | bssid of the network detected as threat | 
 **apMac** | **string** | mac of the device that had strongest signal strength for ssid/bssid pair | 
 **channel** | **int32** | channel over which ap_mac heard ssid/bssid pair | 
 **seenOnLan** | **bool** | whether the reporting AP see a wireless client (on LAN) connecting to it | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseEventsRogueSearch**](ResponseEventsRogueSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

