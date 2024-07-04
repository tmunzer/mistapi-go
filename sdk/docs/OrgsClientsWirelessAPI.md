# \OrgsClientsWirelessAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgWirelessClients**](OrgsClientsWirelessAPI.md#CountOrgWirelessClients) | **Get** /api/v1/orgs/{org_id}/clients/count | countOrgWirelessClients
[**CountOrgWirelessClientsSessions**](OrgsClientsWirelessAPI.md#CountOrgWirelessClientsSessions) | **Get** /api/v1/orgs/{org_id}/clients/sessions/count | countOrgWirelessClientsSessions
[**SearchOrgWirelessClientEvents**](OrgsClientsWirelessAPI.md#SearchOrgWirelessClientEvents) | **Get** /api/v1/orgs/{org_id}/clients/events/search | searchOrgWirelessClientEvents
[**SearchOrgWirelessClientSessions**](OrgsClientsWirelessAPI.md#SearchOrgWirelessClientSessions) | **Get** /api/v1/orgs/{org_id}/clients/sessions/search | searchOrgWirelessClientSessions
[**SearchOrgWirelessClients**](OrgsClientsWirelessAPI.md#SearchOrgWirelessClients) | **Get** /api/v1/orgs/{org_id}/clients/search | searchOrgWirelessClients



## CountOrgWirelessClients

> RepsonseCount CountOrgWirelessClients(ctx, orgId).Distinct(distinct).Mac(mac).Hostname(hostname).Device(device).Os(os).Model(model).Ap(ap).Vlan(vlan).Ssid(ssid).IpAddress(ipAddress).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgWirelessClients



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
	distinct := openapiclient.org_clients_count_distinct("") // OrgClientsCountDistinct |  (optional) (default to "device")
	mac := "mac_example" // string | partial / full MAC address (optional)
	hostname := "hostname_example" // string | partial / full hostname (optional)
	device := "device_example" // string | device type, e.g. Mac, Nvidia, iPhone (optional)
	os := "os_example" // string | os, e.g. Sierra, Yosemite, Windows 10 (optional)
	model := "model_example" // string | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” (optional)
	ap := "ap_example" // string | AP mac where the client has connected to (optional)
	vlan := "vlan_example" // string | vlan (optional)
	ssid := "ssid_example" // string | SSID (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWirelessAPI.CountOrgWirelessClients(context.Background(), orgId).Distinct(distinct).Mac(mac).Hostname(hostname).Device(device).Os(os).Model(model).Ap(ap).Vlan(vlan).Ssid(ssid).IpAddress(ipAddress).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWirelessAPI.CountOrgWirelessClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgWirelessClients`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWirelessAPI.CountOrgWirelessClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgWirelessClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgClientsCountDistinct**](OrgClientsCountDistinct.md) |  | [default to &quot;device&quot;]
 **mac** | **string** | partial / full MAC address | 
 **hostname** | **string** | partial / full hostname | 
 **device** | **string** | device type, e.g. Mac, Nvidia, iPhone | 
 **os** | **string** | os, e.g. Sierra, Yosemite, Windows 10 | 
 **model** | **string** | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” | 
 **ap** | **string** | AP mac where the client has connected to | 
 **vlan** | **string** | vlan | 
 **ssid** | **string** | SSID | 
 **ipAddress** | **string** |  | 
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


## CountOrgWirelessClientsSessions

> RepsonseCount CountOrgWirelessClientsSessions(ctx, orgId).Distinct(distinct).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgWirelessClientsSessions



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
	distinct := openapiclient.org_client_sessions_count_distinct("") // OrgClientSessionsCountDistinct |  (optional) (default to "device")
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
	resp, r, err := apiClient.OrgsClientsWirelessAPI.CountOrgWirelessClientsSessions(context.Background(), orgId).Distinct(distinct).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWirelessAPI.CountOrgWirelessClientsSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgWirelessClientsSessions`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWirelessAPI.CountOrgWirelessClientsSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgWirelessClientsSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgClientSessionsCountDistinct**](OrgClientSessionsCountDistinct.md) |  | [default to &quot;device&quot;]
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


## SearchOrgWirelessClientEvents

> ResponseEventsSearch SearchOrgWirelessClientEvents(ctx, orgId).Type_(type_).ReasonCode(reasonCode).Ssid(ssid).Ap(ap).Proto(proto).Band(band).WlanId(wlanId).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgWirelessClientEvents



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
	resp, r, err := apiClient.OrgsClientsWirelessAPI.SearchOrgWirelessClientEvents(context.Background(), orgId).Type_(type_).ReasonCode(reasonCode).Ssid(ssid).Ap(ap).Proto(proto).Band(band).WlanId(wlanId).NacruleId(nacruleId).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWirelessAPI.SearchOrgWirelessClientEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgWirelessClientEvents`: ResponseEventsSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWirelessAPI.SearchOrgWirelessClientEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgWirelessClientEventsRequest struct via the builder pattern


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


## SearchOrgWirelessClientSessions

> SearchWirelssClientSession SearchOrgWirelessClientSessions(ctx, orgId).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientUsername(clientUsername).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).PskId(pskId).PskName(pskName).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgWirelessClientSessions



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
	resp, r, err := apiClient.OrgsClientsWirelessAPI.SearchOrgWirelessClientSessions(context.Background(), orgId).Ap(ap).Band(band).ClientFamily(clientFamily).ClientManufacture(clientManufacture).ClientModel(clientModel).ClientUsername(clientUsername).ClientOs(clientOs).Ssid(ssid).WlanId(wlanId).PskId(pskId).PskName(pskName).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWirelessAPI.SearchOrgWirelessClientSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgWirelessClientSessions`: SearchWirelssClientSession
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWirelessAPI.SearchOrgWirelessClientSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgWirelessClientSessionsRequest struct via the builder pattern


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

[**SearchWirelssClientSession**](SearchWirelssClientSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgWirelessClients

> ResponseClientSearch SearchOrgWirelessClients(ctx, orgId).SiteId(siteId).Mac(mac).IpAddress(ipAddress).Hostname(hostname).Device(device).Os(os).Model(model).Ap(ap).PskId(pskId).PskName(pskName).Vlan(vlan).Ssid(ssid).Text(text).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgWirelessClients



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
	siteId := "siteId_example" // string | Site ID (optional)
	mac := "mac_example" // string | partial / full MAC address (optional)
	ipAddress := "192.168.1.1" // string |  (optional)
	hostname := "hostname_example" // string | partial / full hostname (optional)
	device := "device_example" // string | device type, e.g. Mac, Nvidia, iPhone (optional)
	os := "os_example" // string | os, e.g. Sierra, Yosemite, Windows 10 (optional)
	model := "model_example" // string | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” (optional)
	ap := "ap_example" // string | AP mac where the client has connected to (optional)
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID (optional)
	pskName := "pskName_example" // string | PSK Name (optional)
	vlan := "vlan_example" // string | vlan (optional)
	ssid := "ssid_example" // string | SSID (optional)
	text := "text_example" // string | partial / full MAC address, hostname, username, psk_name or ip (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsWirelessAPI.SearchOrgWirelessClients(context.Background(), orgId).SiteId(siteId).Mac(mac).IpAddress(ipAddress).Hostname(hostname).Device(device).Os(os).Model(model).Ap(ap).PskId(pskId).PskName(pskName).Vlan(vlan).Ssid(ssid).Text(text).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsWirelessAPI.SearchOrgWirelessClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgWirelessClients`: ResponseClientSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsWirelessAPI.SearchOrgWirelessClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgWirelessClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | Site ID | 
 **mac** | **string** | partial / full MAC address | 
 **ipAddress** | **string** |  | 
 **hostname** | **string** | partial / full hostname | 
 **device** | **string** | device type, e.g. Mac, Nvidia, iPhone | 
 **os** | **string** | os, e.g. Sierra, Yosemite, Windows 10 | 
 **model** | **string** | model, e.g. “MBP 15 late 2013”, 6, 6s, “8+ GSM” | 
 **ap** | **string** | AP mac where the client has connected to | 
 **pskId** | **string** | PSK ID | 
 **pskName** | **string** | PSK Name | 
 **vlan** | **string** | vlan | 
 **ssid** | **string** | SSID | 
 **text** | **string** | partial / full MAC address, hostname, username, psk_name or ip | 
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

