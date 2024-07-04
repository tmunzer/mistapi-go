# \SitesDevicesStatsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteBgpStats**](SitesDevicesStatsAPI.md#CountSiteBgpStats) | **Get** /api/v1/sites/{site_id}/stats/bgp_peers/count | countSiteBgpStats
[**CountSiteSwOrGwPorts**](SitesDevicesStatsAPI.md#CountSiteSwOrGwPorts) | **Get** /api/v1/sites/{site_id}/stats/ports/count | countSiteSwOrGwPorts
[**CountSiteSwitchPorts**](SitesDevicesStatsAPI.md#CountSiteSwitchPorts) | **Get** /api/v1/sites/{site_id}/stats/switch_ports/count | countSiteSwitchPorts
[**GetSiteAllClientsStatsByDevice**](SitesDevicesStatsAPI.md#GetSiteAllClientsStatsByDevice) | **Get** /api/v1/sites/{site_id}/stats/devices/{device_id}/clients | getSiteAllClientsStatsByDevice
[**GetSiteDeviceStats**](SitesDevicesStatsAPI.md#GetSiteDeviceStats) | **Get** /api/v1/sites/{site_id}/stats/devices/{device_id} | getSiteDeviceStats
[**GetSiteGatewayMetrics**](SitesDevicesStatsAPI.md#GetSiteGatewayMetrics) | **Get** /api/v1/sites/{site_id}/stats/gateways/metrics | getSiteGatewayMetrics
[**GetSiteMxEdgeStats**](SitesDevicesStatsAPI.md#GetSiteMxEdgeStats) | **Get** /api/v1/sites/{site_id}/stats/mxedges/{mxedge_id} | getSiteMxEdgeStats
[**ListSiteDevicesStats**](SitesDevicesStatsAPI.md#ListSiteDevicesStats) | **Get** /api/v1/sites/{site_id}/stats/devices | listSiteDevicesStats
[**ListSiteMxEdgesStats**](SitesDevicesStatsAPI.md#ListSiteMxEdgesStats) | **Get** /api/v1/sites/{site_id}/stats/mxedges | listSiteMxEdgesStats
[**SearchSiteBgpStats**](SitesDevicesStatsAPI.md#SearchSiteBgpStats) | **Get** /api/v1/sites/{site_id}/stats/bgp_peers/search | searchSiteBgpStats
[**SearchSiteSwOrGwPorts**](SitesDevicesStatsAPI.md#SearchSiteSwOrGwPorts) | **Get** /api/v1/sites/{site_id}/stats/ports/search | searchSiteSwOrGwPorts
[**SearchSiteSwitchPorts**](SitesDevicesStatsAPI.md#SearchSiteSwitchPorts) | **Get** /api/v1/sites/{site_id}/stats/switch_ports/search | searchSiteSwitchPorts



## CountSiteBgpStats

> RepsonseCount CountSiteBgpStats(ctx, siteId).State(state).Distinct(distinct).Execute()

countSiteBgpStats



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
	state := "state_example" // string |  (optional)
	distinct := "distinct_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.CountSiteBgpStats(context.Background(), siteId).State(state).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.CountSiteBgpStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteBgpStats`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.CountSiteBgpStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteBgpStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** |  | 
 **distinct** | **string** |  | 

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


## CountSiteSwOrGwPorts

> RepsonseCount CountSiteSwOrGwPorts(ctx, siteId).Distinct(distinct).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteSwOrGwPorts



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
	distinct := openapiclient.site_ports_count_distinct("") // SitePortsCountDistinct |  (optional) (default to "mac")
	fullDuplex := true // bool | indicates full or half duplex (optional)
	mac := "mac_example" // string | device identifier (optional)
	neighborMac := "neighborMac_example" // string | Chassis identifier of the chassis type listed (optional)
	neighborPortDesc := "neighborPortDesc_example" // string | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” (optional)
	neighborSystemName := "neighborSystemName_example" // string | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” (optional)
	poeDisabled := true // bool | is the POE configured not be disabled. (optional)
	poeMode := "poeMode_example" // string | poe mode depending on class E.g. “802.3at” (optional)
	poeOn := true // bool | is the device attached to POE (optional)
	portId := "portId_example" // string | interface name (optional)
	portMac := "portMac_example" // string | interface mac address (optional)
	powerDraw := float32(8.14) // float32 | Amount of power being used by the interface at the time the command is executed. Unit in watts. (optional)
	txPkts := int32(56) // int32 | Output packets (optional)
	rxPkts := int32(56) // int32 | Input packets (optional)
	rxBytes := int32(56) // int32 | Input bytes (optional)
	txBps := int32(56) // int32 | Output rate (optional)
	rxBps := int32(56) // int32 | Input rate (optional)
	txMcastPkts := int32(56) // int32 | Multicast output packets (optional)
	txBcastPkts := int32(56) // int32 | Broadcast output packets (optional)
	rxMcastPkts := int32(56) // int32 | Multicast input packets (optional)
	rxBcastPkts := int32(56) // int32 | Broadcast input packets (optional)
	speed := int32(56) // int32 | port speed (optional)
	stpState := openapiclient.count_site_sw_or_gw_ports_stp_state("") // CountSiteSwOrGwPortsStpState | if `up`==`true` (optional)
	stpRole := openapiclient.count_site_sw_or_gw_ports_stp_role("") // CountSiteSwOrGwPortsStpRole | if `up`==`true` (optional)
	authState := openapiclient.count_site_sw_or_gw_ports_auth_state("") // CountSiteSwOrGwPortsAuthState | if `up`==`true` && has Authenticator role (optional)
	up := true // bool | indicates if interface is up (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.CountSiteSwOrGwPorts(context.Background(), siteId).Distinct(distinct).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.CountSiteSwOrGwPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteSwOrGwPorts`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.CountSiteSwOrGwPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteSwOrGwPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SitePortsCountDistinct**](SitePortsCountDistinct.md) |  | [default to &quot;mac&quot;]
 **fullDuplex** | **bool** | indicates full or half duplex | 
 **mac** | **string** | device identifier | 
 **neighborMac** | **string** | Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **string** | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **string** | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **bool** | is the POE configured not be disabled. | 
 **poeMode** | **string** | poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **bool** | is the device attached to POE | 
 **portId** | **string** | interface name | 
 **portMac** | **string** | interface mac address | 
 **powerDraw** | **float32** | Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **int32** | Output packets | 
 **rxPkts** | **int32** | Input packets | 
 **rxBytes** | **int32** | Input bytes | 
 **txBps** | **int32** | Output rate | 
 **rxBps** | **int32** | Input rate | 
 **txMcastPkts** | **int32** | Multicast output packets | 
 **txBcastPkts** | **int32** | Broadcast output packets | 
 **rxMcastPkts** | **int32** | Multicast input packets | 
 **rxBcastPkts** | **int32** | Broadcast input packets | 
 **speed** | **int32** | port speed | 
 **stpState** | [**CountSiteSwOrGwPortsStpState**](CountSiteSwOrGwPortsStpState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**CountSiteSwOrGwPortsStpRole**](CountSiteSwOrGwPortsStpRole.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**CountSiteSwOrGwPortsAuthState**](CountSiteSwOrGwPortsAuthState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **up** | **bool** | indicates if interface is up | 
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


## CountSiteSwitchPorts

> RepsonseCount CountSiteSwitchPorts(ctx, siteId).Distinct(distinct).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteSwitchPorts



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
	distinct := openapiclient.site_switch_ports_count_distinct("") // SiteSwitchPortsCountDistinct |  (optional) (default to "mac")
	fullDuplex := true // bool | indicates full or half duplex (optional)
	mac := "mac_example" // string | device identifier (optional)
	neighborMac := "neighborMac_example" // string | Chassis identifier of the chassis type listed (optional)
	neighborPortDesc := "neighborPortDesc_example" // string | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” (optional)
	neighborSystemName := "neighborSystemName_example" // string | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” (optional)
	poeDisabled := true // bool | is the POE configured not be disabled. (optional)
	poeMode := "poeMode_example" // string | poe mode depending on class E.g. “802.3at” (optional)
	poeOn := true // bool | is the device attached to POE (optional)
	portId := "portId_example" // string | interface name (optional)
	portMac := "portMac_example" // string | interface mac address (optional)
	powerDraw := float32(8.14) // float32 | Amount of power being used by the interface at the time the command is executed. Unit in watts. (optional)
	txPkts := int32(56) // int32 | Output packets (optional)
	rxPkts := int32(56) // int32 | Input packets (optional)
	rxBytes := int32(56) // int32 | Input bytes (optional)
	txBps := int32(56) // int32 | Output rate (optional)
	rxBps := int32(56) // int32 | Input rate (optional)
	txMcastPkts := int32(56) // int32 | Multicast output packets (optional)
	txBcastPkts := int32(56) // int32 | Broadcast output packets (optional)
	rxMcastPkts := int32(56) // int32 | Multicast input packets (optional)
	rxBcastPkts := int32(56) // int32 | Broadcast input packets (optional)
	speed := int32(56) // int32 | port speed (optional)
	stpState := openapiclient.count_site_switch_ports_stp_state("") // CountSiteSwitchPortsStpState | if `up`==`true` (optional)
	stpRole := openapiclient.count_site_switch_ports_stp_role("") // CountSiteSwitchPortsStpRole | if `up`==`true` (optional)
	authState := openapiclient.count_site_switch_ports_auth_state("") // CountSiteSwitchPortsAuthState | if `up`==`true` (optional)
	up := true // bool | indicates if interface is up (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.CountSiteSwitchPorts(context.Background(), siteId).Distinct(distinct).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.CountSiteSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteSwitchPorts`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.CountSiteSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteSwitchPortsCountDistinct**](SiteSwitchPortsCountDistinct.md) |  | [default to &quot;mac&quot;]
 **fullDuplex** | **bool** | indicates full or half duplex | 
 **mac** | **string** | device identifier | 
 **neighborMac** | **string** | Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **string** | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **string** | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **bool** | is the POE configured not be disabled. | 
 **poeMode** | **string** | poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **bool** | is the device attached to POE | 
 **portId** | **string** | interface name | 
 **portMac** | **string** | interface mac address | 
 **powerDraw** | **float32** | Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **int32** | Output packets | 
 **rxPkts** | **int32** | Input packets | 
 **rxBytes** | **int32** | Input bytes | 
 **txBps** | **int32** | Output rate | 
 **rxBps** | **int32** | Input rate | 
 **txMcastPkts** | **int32** | Multicast output packets | 
 **txBcastPkts** | **int32** | Broadcast output packets | 
 **rxMcastPkts** | **int32** | Multicast input packets | 
 **rxBcastPkts** | **int32** | Broadcast input packets | 
 **speed** | **int32** | port speed | 
 **stpState** | [**CountSiteSwitchPortsStpState**](CountSiteSwitchPortsStpState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**CountSiteSwitchPortsStpRole**](CountSiteSwitchPortsStpRole.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**CountSiteSwitchPortsAuthState**](CountSiteSwitchPortsAuthState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **up** | **bool** | indicates if interface is up | 
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


## GetSiteAllClientsStatsByDevice

> [][]ClientWirelessStats GetSiteAllClientsStatsByDevice(ctx, siteId, deviceId).Execute()

getSiteAllClientsStatsByDevice



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
	resp, r, err := apiClient.SitesDevicesStatsAPI.GetSiteAllClientsStatsByDevice(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.GetSiteAllClientsStatsByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAllClientsStatsByDevice`: [][]ClientWirelessStats
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.GetSiteAllClientsStatsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAllClientsStatsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetSiteDeviceStats

> StatsDevice GetSiteDeviceStats(ctx, siteId, deviceId).Execute()

getSiteDeviceStats



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
	resp, r, err := apiClient.SitesDevicesStatsAPI.GetSiteDeviceStats(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.GetSiteDeviceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceStats`: StatsDevice
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.GetSiteDeviceStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StatsDevice**](StatsDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteGatewayMetrics

> GatewayMetrics GetSiteGatewayMetrics(ctx, siteId).Execute()

getSiteGatewayMetrics



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
	resp, r, err := apiClient.SitesDevicesStatsAPI.GetSiteGatewayMetrics(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.GetSiteGatewayMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteGatewayMetrics`: GatewayMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.GetSiteGatewayMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteGatewayMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayMetrics**](GatewayMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteMxEdgeStats

> MxedgeStats GetSiteMxEdgeStats(ctx, siteId, mxedgeId).Start(start).End(end).Duration(duration).Execute()

getSiteMxEdgeStats



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.GetSiteMxEdgeStats(context.Background(), siteId, mxedgeId).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.GetSiteMxEdgeStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteMxEdgeStats`: MxedgeStats
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.GetSiteMxEdgeStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMxEdgeStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**MxedgeStats**](MxedgeStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteDevicesStats

> []ResponseStatsDevice ListSiteDevicesStats(ctx, siteId).Type_(type_).Status(status).Page(page).Limit(limit).Execute()

listSiteDevicesStats



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
	status := openapiclient.stat_device_status_filter("") // StatDeviceStatusFilter |  (optional) (default to "all")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.ListSiteDevicesStats(context.Background(), siteId).Type_(type_).Status(status).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.ListSiteDevicesStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteDevicesStats`: []ResponseStatsDevice
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.ListSiteDevicesStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteDevicesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceTypeWithAll**](DeviceTypeWithAll.md) |  | [default to &quot;ap&quot;]
 **status** | [**StatDeviceStatusFilter**](StatDeviceStatusFilter.md) |  | [default to &quot;all&quot;]
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]ResponseStatsDevice**](ResponseStatsDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteMxEdgesStats

> []MxedgeStats ListSiteMxEdgesStats(ctx, siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listSiteMxEdgesStats



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.ListSiteMxEdgesStats(context.Background(), siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.ListSiteMxEdgesStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteMxEdgesStats`: []MxedgeStats
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.ListSiteMxEdgesStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteMxEdgesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]MxedgeStats**](MxedgeStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteBgpStats

> ResponseSearchBgps SearchSiteBgpStats(ctx, siteId).Execute()

searchSiteBgpStats



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
	resp, r, err := apiClient.SitesDevicesStatsAPI.SearchSiteBgpStats(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.SearchSiteBgpStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteBgpStats`: ResponseSearchBgps
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.SearchSiteBgpStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteBgpStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseSearchBgps**](ResponseSearchBgps.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteSwOrGwPorts

> ResponseSwitchPortSearch SearchSiteSwOrGwPorts(ctx, siteId).FullDuplex(fullDuplex).Mac(mac).DeviceType(deviceType).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxErrors(txErrors).RxErrors(rxErrors).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).MacLimit(macLimit).MacCount(macCount).Up(up).Active(active).Jitter(jitter).Loss(loss).Latency(latency).StpState(stpState).StpRole(stpRole).XcvrPartNumber(xcvrPartNumber).AuthState(authState).LteImsi(lteImsi).LteIccid(lteIccid).LteImei(lteImei).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteSwOrGwPorts



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
	fullDuplex := true // bool | indicates full or half duplex (optional)
	mac := "mac_example" // string | device identifier (optional)
	deviceType := openapiclient.search_site_sw_or_gw_ports_device_type("") // SearchSiteSwOrGwPortsDeviceType | device type (optional)
	neighborMac := "neighborMac_example" // string | Chassis identifier of the chassis type listed (optional)
	neighborPortDesc := "neighborPortDesc_example" // string | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” (optional)
	neighborSystemName := "neighborSystemName_example" // string | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” (optional)
	poeDisabled := true // bool | is the POE configured not be disabled. (optional)
	poeMode := "poeMode_example" // string | poe mode depending on class E.g. “802.3at” (optional)
	poeOn := true // bool | is the device attached to POE (optional)
	portId := "portId_example" // string | interface name (optional)
	portMac := "portMac_example" // string | interface mac address (optional)
	powerDraw := float32(8.14) // float32 | Amount of power being used by the interface at the time the command is executed. Unit in watts. (optional)
	txPkts := int32(56) // int32 | Output packets (optional)
	rxPkts := int32(56) // int32 | Input packets (optional)
	rxBytes := int32(56) // int32 | Input bytes (optional)
	txBps := int32(56) // int32 | Output rate (optional)
	rxBps := int32(56) // int32 | Input rate (optional)
	txErrors := int32(56) // int32 | Output errors (optional)
	rxErrors := int32(56) // int32 | Input errors (optional)
	txMcastPkts := int32(56) // int32 | Multicast output packets (optional)
	txBcastPkts := int32(56) // int32 | Broadcast output packets (optional)
	rxMcastPkts := int32(56) // int32 | Multicast input packets (optional)
	rxBcastPkts := int32(56) // int32 | Broadcast input packets (optional)
	speed := int32(56) // int32 | port speed (optional)
	macLimit := int32(56) // int32 | Limit on number of dynamically learned macs (optional)
	macCount := int32(56) // int32 | Number of mac addresses in the forwarding table (optional)
	up := true // bool | indicates if interface is up (optional)
	active := true // bool | indicates if interface is active/inactive (optional)
	jitter := float32(8.14) // float32 | Last sampled jitter of the interface (optional)
	loss := float32(8.14) // float32 | Last sampled loss of the interface (optional)
	latency := float32(8.14) // float32 | Last sampled latency of the interface (optional)
	stpState := openapiclient.search_site_sw_or_gw_ports_stp_state("") // SearchSiteSwOrGwPortsStpState | if `up`==`true` (optional)
	stpRole := openapiclient.search_site_sw_or_gw_ports_stp_role("") // SearchSiteSwOrGwPortsStpRole | if `up`==`true` (optional)
	xcvrPartNumber := "xcvrPartNumber_example" // string | Optic Slot Partnumber, Check for null/empty (optional)
	authState := openapiclient.search_site_sw_or_gw_ports_auth_state("") // SearchSiteSwOrGwPortsAuthState | if `up`==`true` && has Authenticator role (optional)
	lteImsi := "lteImsi_example" // string | LTE IMSI value, Check for null/empty (optional)
	lteIccid := "lteIccid_example" // string | LTE ICCID value, Check for null/empty (optional)
	lteImei := "lteImei_example" // string | LTE IMEI value, Check for null/empty (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.SearchSiteSwOrGwPorts(context.Background(), siteId).FullDuplex(fullDuplex).Mac(mac).DeviceType(deviceType).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxErrors(txErrors).RxErrors(rxErrors).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).MacLimit(macLimit).MacCount(macCount).Up(up).Active(active).Jitter(jitter).Loss(loss).Latency(latency).StpState(stpState).StpRole(stpRole).XcvrPartNumber(xcvrPartNumber).AuthState(authState).LteImsi(lteImsi).LteIccid(lteIccid).LteImei(lteImei).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.SearchSiteSwOrGwPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteSwOrGwPorts`: ResponseSwitchPortSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.SearchSiteSwOrGwPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteSwOrGwPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullDuplex** | **bool** | indicates full or half duplex | 
 **mac** | **string** | device identifier | 
 **deviceType** | [**SearchSiteSwOrGwPortsDeviceType**](SearchSiteSwOrGwPortsDeviceType.md) | device type | 
 **neighborMac** | **string** | Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **string** | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **string** | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **bool** | is the POE configured not be disabled. | 
 **poeMode** | **string** | poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **bool** | is the device attached to POE | 
 **portId** | **string** | interface name | 
 **portMac** | **string** | interface mac address | 
 **powerDraw** | **float32** | Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **int32** | Output packets | 
 **rxPkts** | **int32** | Input packets | 
 **rxBytes** | **int32** | Input bytes | 
 **txBps** | **int32** | Output rate | 
 **rxBps** | **int32** | Input rate | 
 **txErrors** | **int32** | Output errors | 
 **rxErrors** | **int32** | Input errors | 
 **txMcastPkts** | **int32** | Multicast output packets | 
 **txBcastPkts** | **int32** | Broadcast output packets | 
 **rxMcastPkts** | **int32** | Multicast input packets | 
 **rxBcastPkts** | **int32** | Broadcast input packets | 
 **speed** | **int32** | port speed | 
 **macLimit** | **int32** | Limit on number of dynamically learned macs | 
 **macCount** | **int32** | Number of mac addresses in the forwarding table | 
 **up** | **bool** | indicates if interface is up | 
 **active** | **bool** | indicates if interface is active/inactive | 
 **jitter** | **float32** | Last sampled jitter of the interface | 
 **loss** | **float32** | Last sampled loss of the interface | 
 **latency** | **float32** | Last sampled latency of the interface | 
 **stpState** | [**SearchSiteSwOrGwPortsStpState**](SearchSiteSwOrGwPortsStpState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**SearchSiteSwOrGwPortsStpRole**](SearchSiteSwOrGwPortsStpRole.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **xcvrPartNumber** | **string** | Optic Slot Partnumber, Check for null/empty | 
 **authState** | [**SearchSiteSwOrGwPortsAuthState**](SearchSiteSwOrGwPortsAuthState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **lteImsi** | **string** | LTE IMSI value, Check for null/empty | 
 **lteIccid** | **string** | LTE ICCID value, Check for null/empty | 
 **lteImei** | **string** | LTE IMEI value, Check for null/empty | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseSwitchPortSearch**](ResponseSwitchPortSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteSwitchPorts

> ResponseSwitchPortSearch SearchSiteSwitchPorts(ctx, siteId).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteSwitchPorts



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
	fullDuplex := true // bool | indicates full or half duplex (optional)
	mac := "mac_example" // string | device identifier (optional)
	neighborMac := "neighborMac_example" // string | Chassis identifier of the chassis type listed (optional)
	neighborPortDesc := "neighborPortDesc_example" // string | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” (optional)
	neighborSystemName := "neighborSystemName_example" // string | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” (optional)
	poeDisabled := true // bool | is the POE configured not be disabled. (optional)
	poeMode := "poeMode_example" // string | poe mode depending on class E.g. “802.3at” (optional)
	poeOn := true // bool | is the device attached to POE (optional)
	portId := "portId_example" // string | interface name (optional)
	portMac := "portMac_example" // string | interface mac address (optional)
	powerDraw := float32(8.14) // float32 | Amount of power being used by the interface at the time the command is executed. Unit in watts. (optional)
	txPkts := int32(56) // int32 | Output packets (optional)
	rxPkts := int32(56) // int32 | Input packets (optional)
	rxBytes := int32(56) // int32 | Input bytes (optional)
	txBps := int32(56) // int32 | Output rate (optional)
	rxBps := int32(56) // int32 | Input rate (optional)
	txMcastPkts := int32(56) // int32 | Multicast output packets (optional)
	txBcastPkts := int32(56) // int32 | Broadcast output packets (optional)
	rxMcastPkts := int32(56) // int32 | Multicast input packets (optional)
	rxBcastPkts := int32(56) // int32 | Broadcast input packets (optional)
	speed := int32(56) // int32 | port speed (optional)
	stpState := openapiclient.search_site_switch_ports_stp_state("") // SearchSiteSwitchPortsStpState | if `up`==`true` (optional)
	stpRole := openapiclient.search_site_switch_ports_stp_role("") // SearchSiteSwitchPortsStpRole | if `up`==`true` (optional)
	authState := openapiclient.search_site_switch_ports_auth_state("") // SearchSiteSwitchPortsAuthState | if `up`==`true` && has Authenticator role (optional)
	up := true // bool | indicates if interface is up (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesStatsAPI.SearchSiteSwitchPorts(context.Background(), siteId).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesStatsAPI.SearchSiteSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteSwitchPorts`: ResponseSwitchPortSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesStatsAPI.SearchSiteSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullDuplex** | **bool** | indicates full or half duplex | 
 **mac** | **string** | device identifier | 
 **neighborMac** | **string** | Chassis identifier of the chassis type listed | 
 **neighborPortDesc** | **string** | Description supplied by the system on the interface E.g. “GigabitEthernet2/0/39” | 
 **neighborSystemName** | **string** | Name supplied by the system on the interface E.g. neighbor system name E.g. “Kumar-Acc-SW.mist.local” | 
 **poeDisabled** | **bool** | is the POE configured not be disabled. | 
 **poeMode** | **string** | poe mode depending on class E.g. “802.3at” | 
 **poeOn** | **bool** | is the device attached to POE | 
 **portId** | **string** | interface name | 
 **portMac** | **string** | interface mac address | 
 **powerDraw** | **float32** | Amount of power being used by the interface at the time the command is executed. Unit in watts. | 
 **txPkts** | **int32** | Output packets | 
 **rxPkts** | **int32** | Input packets | 
 **rxBytes** | **int32** | Input bytes | 
 **txBps** | **int32** | Output rate | 
 **rxBps** | **int32** | Input rate | 
 **txMcastPkts** | **int32** | Multicast output packets | 
 **txBcastPkts** | **int32** | Broadcast output packets | 
 **rxMcastPkts** | **int32** | Multicast input packets | 
 **rxBcastPkts** | **int32** | Broadcast input packets | 
 **speed** | **int32** | port speed | 
 **stpState** | [**SearchSiteSwitchPortsStpState**](SearchSiteSwitchPortsStpState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**SearchSiteSwitchPortsStpRole**](SearchSiteSwitchPortsStpRole.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**SearchSiteSwitchPortsAuthState**](SearchSiteSwitchPortsAuthState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **up** | **bool** | indicates if interface is up | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseSwitchPortSearch**](ResponseSwitchPortSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

