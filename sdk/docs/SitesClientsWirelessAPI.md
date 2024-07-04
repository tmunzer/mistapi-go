# \SitesClientsWirelessAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWirelessClientEvents**](SitesClientsWirelessAPI.md#CountSiteWirelessClientEvents) | **Get** /api/v1/sites/{site_id}/clients/events/count | countSiteWirelessClientEvents
[**CountSiteWirelessClientSessions**](SitesClientsWirelessAPI.md#CountSiteWirelessClientSessions) | **Get** /api/v1/sites/{site_id}/clients/sessions/count | countSiteWirelessClientSessions
[**CountSiteWirelessClients**](SitesClientsWirelessAPI.md#CountSiteWirelessClients) | **Get** /api/v1/sites/{site_id}/clients/count | countSiteWirelessClients
[**GetSiteEventsForClient**](SitesClientsWirelessAPI.md#GetSiteEventsForClient) | **Get** /api/v1/sites/{site_id}/clients/{client_mac}/events | getSiteEventsForClient
[**GetSiteWirelessClientStats**](SitesClientsWirelessAPI.md#GetSiteWirelessClientStats) | **Get** /api/v1/sites/{site_id}/stats/clients/{client_mac} | getSiteWirelessClientStats
[**GetSiteWirelessClientsStatsByMap**](SitesClientsWirelessAPI.md#GetSiteWirelessClientsStatsByMap) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/clients | getSiteWirelessClientsStatsByMap
[**ListSiteUnconnectedClientStats**](SitesClientsWirelessAPI.md#ListSiteUnconnectedClientStats) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/unconnected_clients | listSiteUnconnectedClientStats
[**ListSiteWirelessClientsStats**](SitesClientsWirelessAPI.md#ListSiteWirelessClientsStats) | **Get** /api/v1/sites/{site_id}/stats/clients | listSiteWirelessClientsStats
[**SearchSiteWirelessClientEvents**](SitesClientsWirelessAPI.md#SearchSiteWirelessClientEvents) | **Get** /api/v1/sites/{site_id}/clients/events/search | searchSiteWirelessClientEvents
[**SearchSiteWirelessClientSessions**](SitesClientsWirelessAPI.md#SearchSiteWirelessClientSessions) | **Get** /api/v1/sites/{site_id}/clients/sessions/search | searchSiteWirelessClientSessions
[**SearchSiteWirelessClients**](SitesClientsWirelessAPI.md#SearchSiteWirelessClients) | **Get** /api/v1/sites/{site_id}/clients/search | searchSiteWirelessClients



## CountSiteWirelessClientEvents

> RepsonseCount CountSiteWirelessClientEvents(ctx, siteId).Distinct(distinct).Type_(type_).ReasonCode(reasonCode).Ssid(ssid).Ap(ap).Proto(proto).Band(band).WlanId(wlanId).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteWirelessClientEvents



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
	distinct := openapiclient.site_client_events_count_distinct("") // SiteClientEventsCountDistinct |  (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	reasonCode := int32(56) // int32 | for assoc/disassoc events (optional)
	ssid := "ssid_example" // string | SSID Name (optional)
	ap := "ap_example" // string | AP MAC (optional)
	proto := openapiclient.dot11_proto("") // Dot11Proto | a / b / g / n / ac / ax (optional)
	band := openapiclient.dot11_band("") // Dot11Band | 802.11 Band (optional)
	wlanId := "wlanId_example" // string | wlan_id (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.CountSiteWirelessClientEvents(context.Background(), siteId).Distinct(distinct).Type_(type_).ReasonCode(reasonCode).Ssid(ssid).Ap(ap).Proto(proto).Band(band).WlanId(wlanId).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.CountSiteWirelessClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWirelessClientEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.CountSiteWirelessClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWirelessClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteClientEventsCountDistinct**](SiteClientEventsCountDistinct.md) |  | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **reasonCode** | **int32** | for assoc/disassoc events | 
 **ssid** | **string** | SSID Name | 
 **ap** | **string** | AP MAC | 
 **proto** | [**Dot11Proto**](Dot11Proto.md) | a / b / g / n / ac / ax | 
 **band** | [**Dot11Band**](Dot11Band.md) | 802.11 Band | 
 **wlanId** | **string** | wlan_id | 
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


## CountSiteWirelessClientSessions

> RepsonseCount CountSiteWirelessClientSessions(ctx, siteId).Distinct(distinct).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteWirelessClientSessions



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
	distinct := openapiclient.site_client_sessions_count_distinct("") // SiteClientSessionsCountDistinct |  (optional) (default to "mac")
	ap := "ap_example" // string | AP MAC (optional)
	band := openapiclient.dot11_band("") // Dot11Band | 802.11 Band (optional)
	clientFamily := "clientFamily_example" // string | E.g. “Mac”, “iPhone”, “Apple watch” (optional)
	clientManufacture := "clientManufacture_example" // string | E.g. “Apple” (optional)
	clientModel := "clientModel_example" // string | E.g. “8+”, “XS” (optional)
	clientOs := "clientOs_example" // string | E.g. “Mojave”, “Windows 10”, “Linux” (optional)
	ssid := "ssid_example" // string | SSID (optional)
	wlanId := "wlanId_example" // string | wlan_id (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.CountSiteWirelessClientSessions(context.Background(), siteId).Distinct(distinct).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.CountSiteWirelessClientSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWirelessClientSessions`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.CountSiteWirelessClientSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWirelessClientSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteClientSessionsCountDistinct**](SiteClientSessionsCountDistinct.md) |  | [default to &quot;mac&quot;]
 **ap** | **string** | AP MAC | 
 **band** | [**Dot11Band**](Dot11Band.md) | 802.11 Band | 
 **clientFamily** | **string** | E.g. “Mac”, “iPhone”, “Apple watch” | 
 **clientManufacture** | **string** | E.g. “Apple” | 
 **clientModel** | **string** | E.g. “8+”, “XS” | 
 **clientOs** | **string** | E.g. “Mojave”, “Windows 10”, “Linux” | 
 **ssid** | **string** | SSID | 
 **wlanId** | **string** | wlan_id | 
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


## CountSiteWirelessClients

> RepsonseCount CountSiteWirelessClients(ctx, siteId).Distinct(distinct).Ssid(ssid).Ap(ap).IpAddress(ipAddress).Vlan(vlan).Hostname(hostname).Os(os).Model(model).Device(device).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteWirelessClients



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
	distinct := openapiclient.site_clients_count_distinct("") // SiteClientsCountDistinct |  (optional) (default to "device")
	ssid := "ssid_example" // string |  (optional)
	ap := "ap_example" // string |  (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	vlan := "vlan_example" // string |  (optional)
	hostname := "hostname_example" // string |  (optional)
	os := "os_example" // string |  (optional)
	model := "model_example" // string |  (optional)
	device := "device_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.CountSiteWirelessClients(context.Background(), siteId).Distinct(distinct).Ssid(ssid).Ap(ap).IpAddress(ipAddress).Vlan(vlan).Hostname(hostname).Os(os).Model(model).Device(device).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.CountSiteWirelessClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWirelessClients`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.CountSiteWirelessClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWirelessClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteClientsCountDistinct**](SiteClientsCountDistinct.md) |  | [default to &quot;device&quot;]
 **ssid** | **string** |  | 
 **ap** | **string** |  | 
 **ipAddress** | **string** |  | 
 **vlan** | **string** |  | 
 **hostname** | **string** |  | 
 **os** | **string** |  | 
 **model** | **string** |  | 
 **device** | **string** |  | 
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


## GetSiteEventsForClient

> ResponseClientEventsSearch GetSiteEventsForClient(ctx, siteId, clientMac).Type_(type_).Proto(proto).Band(band).Channel(channel).WlanId(wlanId).Ssid(ssid).Start(start).End(end).Page(page).Limit(limit).Duration(duration).Execute()

getSiteEventsForClient



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
	clientMac := "0000000000ab" // string | 
	type_ := "type__example" // string | e.g. MARVIS_EVENT_CLIENT_DHCP_STUCK (optional)
	proto := openapiclient.dot11_proto("") // Dot11Proto | a / b / g / n / ac / ax (optional)
	band := openapiclient.dot11_band("") // Dot11Band | 802.11 Band (optional)
	channel := "channel_example" // string |  (optional)
	wlanId := "wlanId_example" // string |  (optional)
	ssid := "ssid_example" // string |  (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.GetSiteEventsForClient(context.Background(), siteId, clientMac).Type_(type_).Proto(proto).Band(band).Channel(channel).WlanId(wlanId).Ssid(ssid).Start(start).End(end).Page(page).Limit(limit).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.GetSiteEventsForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteEventsForClient`: ResponseClientEventsSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.GetSiteEventsForClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteEventsForClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | **string** | e.g. MARVIS_EVENT_CLIENT_DHCP_STUCK | 
 **proto** | [**Dot11Proto**](Dot11Proto.md) | a / b / g / n / ac / ax | 
 **band** | [**Dot11Band**](Dot11Band.md) | 802.11 Band | 
 **channel** | **string** |  | 
 **wlanId** | **string** |  | 
 **ssid** | **string** |  | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseClientEventsSearch**](ResponseClientEventsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteWirelessClientStats

> []ClientStats GetSiteWirelessClientStats(ctx, siteId, clientMac).Wired(wired).Execute()

getSiteWirelessClientStats



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
	clientMac := "0000000000ab" // string | 
	wired := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.GetSiteWirelessClientStats(context.Background(), siteId, clientMac).Wired(wired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.GetSiteWirelessClientStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWirelessClientStats`: []ClientStats
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.GetSiteWirelessClientStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWirelessClientStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wired** | **bool** |  | [default to false]

### Return type

[**[]ClientStats**](ClientStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteWirelessClientsStatsByMap

> [][]ClientWirelessStats GetSiteWirelessClientsStatsByMap(ctx, siteId, mapId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

getSiteWirelessClientsStatsByMap



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.GetSiteWirelessClientsStatsByMap(context.Background(), siteId, mapId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.GetSiteWirelessClientsStatsByMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWirelessClientsStatsByMap`: [][]ClientWirelessStats
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.GetSiteWirelessClientsStatsByMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWirelessClientsStatsByMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[][]ClientWirelessStats**](set.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteUnconnectedClientStats

> []UnconnectedClientStat ListSiteUnconnectedClientStats(ctx, siteId, mapId).Execute()

listSiteUnconnectedClientStats



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.ListSiteUnconnectedClientStats(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.ListSiteUnconnectedClientStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteUnconnectedClientStats`: []UnconnectedClientStat
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.ListSiteUnconnectedClientStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteUnconnectedClientStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]UnconnectedClientStat**](UnconnectedClientStat.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteWirelessClientsStats

> []ClientStats ListSiteWirelessClientsStats(ctx, siteId).Wired(wired).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listSiteWirelessClientsStats



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
	wired := true // bool |  (optional) (default to false)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.ListSiteWirelessClientsStats(context.Background(), siteId).Wired(wired).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.ListSiteWirelessClientsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteWirelessClientsStats`: []ClientStats
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.ListSiteWirelessClientsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteWirelessClientsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wired** | **bool** |  | [default to false]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]ClientStats**](ClientStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteWirelessClientEvents

> ResponseEventsSearch SearchSiteWirelessClientEvents(ctx, siteId).Type_(type_).ReasonCode(reasonCode).Ssid(ssid).Ap(ap).Proto(proto).Band(band).WlanId(wlanId).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteWirelessClientEvents



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
	reasonCode := int32(56) // int32 | for assoc/disassoc events (optional)
	ssid := "ssid_example" // string | SSID Name (optional)
	ap := "ap_example" // string | AP MAC (optional)
	proto := openapiclient.dot11_proto("") // Dot11Proto | a / b / g / n / ac / ax (optional)
	band := openapiclient.dot11_band("") // Dot11Band | 802.11 Band (optional)
	wlanId := "wlanId_example" // string | wlan_id (optional)
	nacruleId := "nacruleId_example" // string | nacrule_id (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.SearchSiteWirelessClientEvents(context.Background(), siteId).Type_(type_).ReasonCode(reasonCode).Ssid(ssid).Ap(ap).Proto(proto).Band(band).WlanId(wlanId).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.SearchSiteWirelessClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWirelessClientEvents`: ResponseEventsSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.SearchSiteWirelessClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWirelessClientEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **reasonCode** | **int32** | for assoc/disassoc events | 
 **ssid** | **string** | SSID Name | 
 **ap** | **string** | AP MAC | 
 **proto** | [**Dot11Proto**](Dot11Proto.md) | a / b / g / n / ac / ax | 
 **band** | [**Dot11Band**](Dot11Band.md) | 802.11 Band | 
 **wlanId** | **string** | wlan_id | 
 **nacruleId** | **string** | nacrule_id | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseEventsSearch**](ResponseEventsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteWirelessClientSessions

> ResponseClientSessionsSearch SearchSiteWirelessClientSessions(ctx, siteId).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientUsername(clientUsername).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).PskId(pskId).PskName(pskName).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteWirelessClientSessions



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
	ap := "ap_example" // string | AP MAC (optional)
	band := openapiclient.dot11_band("") // Dot11Band | 802.11 Band (optional)
	clientFamily := "clientFamily_example" // string | E.g. “Mac”, “iPhone”, “Apple watch” (optional)
	clientManufacture := "clientManufacture_example" // string | E.g. “Apple” (optional)
	clientModel := "clientModel_example" // string | E.g. “8+”, “XS” (optional)
	clientUsername := "clientUsername_example" // string | Username (optional)
	clientOs := "clientOs_example" // string | E.g. “Mojave”, “Windows 10”, “Linux” (optional)
	ssid := "ssid_example" // string | SSID (optional)
	wlanId := "wlanId_example" // string | wlan_id (optional)
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID (optional)
	pskName := "pskName_example" // string | PSK Name (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.SearchSiteWirelessClientSessions(context.Background(), siteId).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientUsername(clientUsername).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).PskId(pskId).PskName(pskName).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.SearchSiteWirelessClientSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWirelessClientSessions`: ResponseClientSessionsSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.SearchSiteWirelessClientSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWirelessClientSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ap** | **string** | AP MAC | 
 **band** | [**Dot11Band**](Dot11Band.md) | 802.11 Band | 
 **clientFamily** | **string** | E.g. “Mac”, “iPhone”, “Apple watch” | 
 **clientManufacture** | **string** | E.g. “Apple” | 
 **clientModel** | **string** | E.g. “8+”, “XS” | 
 **clientUsername** | **string** | Username | 
 **clientOs** | **string** | E.g. “Mojave”, “Windows 10”, “Linux” | 
 **ssid** | **string** | SSID | 
 **wlanId** | **string** | wlan_id | 
 **pskId** | **string** | PSK ID | 
 **pskName** | **string** | PSK Name | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseClientSessionsSearch**](ResponseClientSessionsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteWirelessClients

> ResponseClientSearch SearchSiteWirelessClients(ctx, siteId).Mac(mac).IpAddress(ipAddress).Hostname(hostname).Device(device).Os(os).Model(model).Ap(ap).Ssid(ssid).Text(text).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteWirelessClients



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
	ipAddress := "192.168.1.1" // string |  (optional)
	hostname := "hostname_example" // string | partial / full hostname (optional)
	device := "device_example" // string | device type, e.g. Mac, Nvidia, iPhone (optional)
	os := "os_example" // string | os, e.g. Sierra, Yosemite, Windows 10 (optional)
	model := "model_example" // string | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” (optional)
	ap := "ap_example" // string | AP mac where the client has connected to (optional)
	ssid := "ssid_example" // string |  (optional)
	text := "text_example" // string | partial / full MAC address, hostname, username, psk_name or ip (optional)
	nacruleId := "nacruleId_example" // string | nacrule_id (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsWirelessAPI.SearchSiteWirelessClients(context.Background(), siteId).Mac(mac).IpAddress(ipAddress).Hostname(hostname).Device(device).Os(os).Model(model).Ap(ap).Ssid(ssid).Text(text).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsWirelessAPI.SearchSiteWirelessClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWirelessClients`: ResponseClientSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsWirelessAPI.SearchSiteWirelessClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWirelessClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | partial / full MAC address | 
 **ipAddress** | **string** |  | 
 **hostname** | **string** | partial / full hostname | 
 **device** | **string** | device type, e.g. Mac, Nvidia, iPhone | 
 **os** | **string** | os, e.g. Sierra, Yosemite, Windows 10 | 
 **model** | **string** | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” | 
 **ap** | **string** | AP mac where the client has connected to | 
 **ssid** | **string** |  | 
 **text** | **string** | partial / full MAC address, hostname, username, psk_name or ip | 
 **nacruleId** | **string** | nacrule_id | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseClientSearch**](ResponseClientSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

