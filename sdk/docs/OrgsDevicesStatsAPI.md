# \OrgsDevicesStatsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgBgpStats**](OrgsDevicesStatsAPI.md#CountOrgBgpStats) | **Get** /api/v1/orgs/{org_id}/stats/bgp_peers/count | countOrgBgpStats
[**CountOrgSwitchPorts**](OrgsDevicesStatsAPI.md#CountOrgSwitchPorts) | **Get** /api/v1/orgs/{org_id}/stats/switch_ports/count | countOrgSwitchPorts
[**SearchOrgBgpStats**](OrgsDevicesStatsAPI.md#SearchOrgBgpStats) | **Get** /api/v1/orgs/{org_id}/stats/bgp_peers/search | searchOrgBgpStats
[**SearchOrgSwOrGwPorts**](OrgsDevicesStatsAPI.md#SearchOrgSwOrGwPorts) | **Get** /api/v1/orgs/{org_id}/stats/ports/search | searchOrgSwOrGwPorts



## CountOrgBgpStats

> RepsonseCount CountOrgBgpStats(ctx, orgId).Execute()

countOrgBgpStats



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesStatsAPI.CountOrgBgpStats(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesStatsAPI.CountOrgBgpStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgBgpStats`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesStatsAPI.CountOrgBgpStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgBgpStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CountOrgSwitchPorts

> RepsonseCount CountOrgSwitchPorts(ctx, orgId).Distinct(distinct).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgSwitchPorts



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
	distinct := openapiclient.org_switch_port_count_distinct("") // OrgSwitchPortCountDistinct |  (optional) (default to "mac")
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
	stpState := openapiclient.count_org_switch_ports_stp_state("") // CountOrgSwitchPortsStpState | if `up`==`true` (optional)
	stpRole := openapiclient.count_org_switch_ports_stp_role("") // CountOrgSwitchPortsStpRole | if `up`==`true` (optional)
	authState := openapiclient.count_org_switch_ports_auth_state("") // CountOrgSwitchPortsAuthState | if `up`==`true` (optional)
	up := true // bool | indicates if interface is up (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesStatsAPI.CountOrgSwitchPorts(context.Background(), orgId).Distinct(distinct).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).StpState(stpState).StpRole(stpRole).AuthState(authState).Up(up).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesStatsAPI.CountOrgSwitchPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgSwitchPorts`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesStatsAPI.CountOrgSwitchPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgSwitchPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgSwitchPortCountDistinct**](OrgSwitchPortCountDistinct.md) |  | [default to &quot;mac&quot;]
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
 **stpState** | [**CountOrgSwitchPortsStpState**](CountOrgSwitchPortsStpState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**CountOrgSwitchPortsStpRole**](CountOrgSwitchPortsStpRole.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**CountOrgSwitchPortsAuthState**](CountOrgSwitchPortsAuthState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
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


## SearchOrgBgpStats

> ResponseSearchBgps SearchOrgBgpStats(ctx, orgId).Execute()

searchOrgBgpStats



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesStatsAPI.SearchOrgBgpStats(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesStatsAPI.SearchOrgBgpStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgBgpStats`: ResponseSearchBgps
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesStatsAPI.SearchOrgBgpStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgBgpStatsRequest struct via the builder pattern


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


## SearchOrgSwOrGwPorts

> ResponsePortStatsSearch SearchOrgSwOrGwPorts(ctx, orgId).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxErrors(txErrors).RxErrors(rxErrors).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).MacLimit(macLimit).MacCount(macCount).Up(up).StpState(stpState).StpRole(stpRole).AuthState(authState).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgSwOrGwPorts



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
	stpState := openapiclient.search_org_sw_or_gw_ports_stp_state("") // SearchOrgSwOrGwPortsStpState | if `up`==`true` (optional)
	stpRole := openapiclient.search_org_sw_or_gw_ports_stp_role("") // SearchOrgSwOrGwPortsStpRole | if `up`==`true` (optional)
	authState := openapiclient.search_org_sw_or_gw_ports_auth_state("") // SearchOrgSwOrGwPortsAuthState | if `up`==`true` && has Authenticator role (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDevicesStatsAPI.SearchOrgSwOrGwPorts(context.Background(), orgId).FullDuplex(fullDuplex).Mac(mac).NeighborMac(neighborMac).NeighborPortDesc(neighborPortDesc).NeighborSystemName(neighborSystemName).PoeDisabled(poeDisabled).PoeMode(poeMode).PoeOn(poeOn).PortId(portId).PortMac(portMac).PowerDraw(powerDraw).TxPkts(txPkts).RxPkts(rxPkts).RxBytes(rxBytes).TxBps(txBps).RxBps(rxBps).TxErrors(txErrors).RxErrors(rxErrors).TxMcastPkts(txMcastPkts).TxBcastPkts(txBcastPkts).RxMcastPkts(rxMcastPkts).RxBcastPkts(rxBcastPkts).Speed(speed).MacLimit(macLimit).MacCount(macCount).Up(up).StpState(stpState).StpRole(stpRole).AuthState(authState).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesStatsAPI.SearchOrgSwOrGwPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgSwOrGwPorts`: ResponsePortStatsSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesStatsAPI.SearchOrgSwOrGwPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgSwOrGwPortsRequest struct via the builder pattern


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
 **stpState** | [**SearchOrgSwOrGwPortsStpState**](SearchOrgSwOrGwPortsStpState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **stpRole** | [**SearchOrgSwOrGwPortsStpRole**](SearchOrgSwOrGwPortsStpRole.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; | 
 **authState** | [**SearchOrgSwOrGwPortsAuthState**](SearchOrgSwOrGwPortsAuthState.md) | if &#x60;up&#x60;&#x3D;&#x3D;&#x60;true&#x60; &amp;&amp; has Authenticator role | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponsePortStatsSearch**](ResponsePortStatsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

