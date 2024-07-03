# \UtilitiesCommonAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArpFromDevice**](UtilitiesCommonAPI.md#ArpFromDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/arp | arpFromDevice
[**ClearSiteDeviceMacTable**](UtilitiesCommonAPI.md#ClearSiteDeviceMacTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_mac_table | clearSiteDeviceMacTable
[**CreateSiteDeviceShellSession**](UtilitiesCommonAPI.md#CreateSiteDeviceShellSession) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/shell | createSiteDeviceShellSession
[**GetSiteDeviceArpTable**](UtilitiesCommonAPI.md#GetSiteDeviceArpTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_arp | getSiteDeviceArpTable
[**GetSiteDeviceBgpSummary**](UtilitiesCommonAPI.md#GetSiteDeviceBgpSummary) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_bgp_rummary | getSiteDeviceBgpSummary
[**GetSiteDeviceConfigCmd**](UtilitiesCommonAPI.md#GetSiteDeviceConfigCmd) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/config_cmd | getSiteDeviceConfigCmd
[**GetSiteDeviceEvpnDatabase**](UtilitiesCommonAPI.md#GetSiteDeviceEvpnDatabase) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_evpn_database | getSiteDeviceEvpnDatabase
[**GetSiteDeviceForwardingTable**](UtilitiesCommonAPI.md#GetSiteDeviceForwardingTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_forwarding_table | getSiteDeviceForwardingTable
[**GetSiteDeviceMacTable**](UtilitiesCommonAPI.md#GetSiteDeviceMacTable) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_mac_table | getSiteDeviceMacTable
[**GetSiteDeviceZtpPassword**](UtilitiesCommonAPI.md#GetSiteDeviceZtpPassword) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/request_ztp_password | getSiteDeviceZtpPassword
[**MonitorSiteDeviceTraffic**](UtilitiesCommonAPI.md#MonitorSiteDeviceTraffic) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/monitor_traffic | monitorSiteDeviceTraffic
[**PingFromDevice**](UtilitiesCommonAPI.md#PingFromDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/ping | pingFromDevice
[**ReadoptSiteOctermDevice**](UtilitiesCommonAPI.md#ReadoptSiteOctermDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/readopt | readoptSiteOctermDevice
[**ReprovisionSiteOctermDevice**](UtilitiesCommonAPI.md#ReprovisionSiteOctermDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/reprovision | readoptSiteOctermDevice
[**RestartSiteDevice**](UtilitiesCommonAPI.md#RestartSiteDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/restart | restartSiteDevice
[**StartSiteLocateDevice**](UtilitiesCommonAPI.md#StartSiteLocateDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/locate | startSiteLocateDevice
[**StopSiteLocateDevice**](UtilitiesCommonAPI.md#StopSiteLocateDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/unlocate | stopSiteLocateDevice
[**TracerouteFromDevice**](UtilitiesCommonAPI.md#TracerouteFromDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/traceroute | tracerouteFromDevice
[**UploadSiteDeviceSupportFile**](UtilitiesCommonAPI.md#UploadSiteDeviceSupportFile) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/support | uploadSiteDeviceSupportFile



## ArpFromDevice

> WebsocketSession ArpFromDevice(ctx, siteId, deviceId).HaClusterNode(haClusterNode).Execute()

arpFromDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	haClusterNode := *openapiclient.NewHaClusterNode() // HaClusterNode |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.ArpFromDevice(context.Background(), siteId, deviceId).HaClusterNode(haClusterNode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.ArpFromDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArpFromDevice`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.ArpFromDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArpFromDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **haClusterNode** | [**HaClusterNode**](HaClusterNode.md) |  | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearSiteDeviceMacTable

> WebsocketSession ClearSiteDeviceMacTable(ctx, siteId, deviceId).UtilsMacTable(utilsMacTable).Execute()

clearSiteDeviceMacTable



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsMacTable := *openapiclient.NewUtilsMacTable() // UtilsMacTable | all attributes are optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.ClearSiteDeviceMacTable(context.Background(), siteId, deviceId).UtilsMacTable(utilsMacTable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.ClearSiteDeviceMacTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearSiteDeviceMacTable`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.ClearSiteDeviceMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearSiteDeviceMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsMacTable** | [**UtilsMacTable**](UtilsMacTable.md) | all attributes are optional | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteDeviceShellSession

> WebsocketSessionWithUrl CreateSiteDeviceShellSession(ctx, siteId, deviceId).Execute()

createSiteDeviceShellSession



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.CreateSiteDeviceShellSession(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.CreateSiteDeviceShellSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteDeviceShellSession`: WebsocketSessionWithUrl
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.CreateSiteDeviceShellSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteDeviceShellSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebsocketSessionWithUrl**](WebsocketSessionWithUrl.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceArpTable

> WebsocketSession GetSiteDeviceArpTable(ctx, siteId, deviceId).UtilsShowArp(utilsShowArp).Execute()

getSiteDeviceArpTable



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsShowArp := *openapiclient.NewUtilsShowArp() // UtilsShowArp | all attributes are optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.GetSiteDeviceArpTable(context.Background(), siteId, deviceId).UtilsShowArp(utilsShowArp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.GetSiteDeviceArpTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceArpTable`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.GetSiteDeviceArpTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceArpTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsShowArp** | [**UtilsShowArp**](UtilsShowArp.md) | all attributes are optional | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceBgpSummary

> WebsocketSession GetSiteDeviceBgpSummary(ctx, siteId, deviceId).UtilsShowBgpRummary(utilsShowBgpRummary).Execute()

getSiteDeviceBgpSummary



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsShowBgpRummary := *openapiclient.NewUtilsShowBgpRummary() // UtilsShowBgpRummary | all attributes are optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.GetSiteDeviceBgpSummary(context.Background(), siteId, deviceId).UtilsShowBgpRummary(utilsShowBgpRummary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.GetSiteDeviceBgpSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceBgpSummary`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.GetSiteDeviceBgpSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceBgpSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsShowBgpRummary** | [**UtilsShowBgpRummary**](UtilsShowBgpRummary.md) | all attributes are optional | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceConfigCmd

> ResponseDeviceConfigCli GetSiteDeviceConfigCmd(ctx, siteId, deviceId).Sort(sort).Execute()

getSiteDeviceConfigCmd



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sort := true // bool | Make output cmds sorted (for better readability) or not. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.GetSiteDeviceConfigCmd(context.Background(), siteId, deviceId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.GetSiteDeviceConfigCmd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceConfigCmd`: ResponseDeviceConfigCli
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.GetSiteDeviceConfigCmd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceConfigCmdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **bool** | Make output cmds sorted (for better readability) or not. | [default to false]

### Return type

[**ResponseDeviceConfigCli**](ResponseDeviceConfigCli.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceEvpnDatabase

> WebsocketSession GetSiteDeviceEvpnDatabase(ctx, siteId, deviceId).UtilsShowEvpnDatabase(utilsShowEvpnDatabase).Execute()

getSiteDeviceEvpnDatabase



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsShowEvpnDatabase := *openapiclient.NewUtilsShowEvpnDatabase() // UtilsShowEvpnDatabase | all attributes are optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.GetSiteDeviceEvpnDatabase(context.Background(), siteId, deviceId).UtilsShowEvpnDatabase(utilsShowEvpnDatabase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.GetSiteDeviceEvpnDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceEvpnDatabase`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.GetSiteDeviceEvpnDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceEvpnDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsShowEvpnDatabase** | [**UtilsShowEvpnDatabase**](UtilsShowEvpnDatabase.md) | all attributes are optional | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceForwardingTable

> WebsocketSession GetSiteDeviceForwardingTable(ctx, siteId, deviceId).UtilsShowForwardingTable(utilsShowForwardingTable).Execute()

getSiteDeviceForwardingTable



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsShowForwardingTable := *openapiclient.NewUtilsShowForwardingTable() // UtilsShowForwardingTable | all attributes are optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.GetSiteDeviceForwardingTable(context.Background(), siteId, deviceId).UtilsShowForwardingTable(utilsShowForwardingTable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.GetSiteDeviceForwardingTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceForwardingTable`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.GetSiteDeviceForwardingTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceForwardingTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsShowForwardingTable** | [**UtilsShowForwardingTable**](UtilsShowForwardingTable.md) | all attributes are optional | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceMacTable

> WebsocketSession GetSiteDeviceMacTable(ctx, siteId, deviceId).UtilsMacTable(utilsMacTable).Execute()

getSiteDeviceMacTable



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsMacTable := *openapiclient.NewUtilsMacTable() // UtilsMacTable | all attributes are optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.GetSiteDeviceMacTable(context.Background(), siteId, deviceId).UtilsMacTable(utilsMacTable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.GetSiteDeviceMacTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceMacTable`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.GetSiteDeviceMacTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceMacTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsMacTable** | [**UtilsMacTable**](UtilsMacTable.md) | all attributes are optional | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceZtpPassword

> RootPasswordString GetSiteDeviceZtpPassword(ctx, siteId, deviceId).Execute()

getSiteDeviceZtpPassword



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.GetSiteDeviceZtpPassword(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.GetSiteDeviceZtpPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceZtpPassword`: RootPasswordString
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.GetSiteDeviceZtpPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceZtpPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RootPasswordString**](RootPasswordString.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitorSiteDeviceTraffic

> WebsocketSessionWithUrl MonitorSiteDeviceTraffic(ctx, siteId, deviceId).UtilsMonitorTraffic(utilsMonitorTraffic).Execute()

monitorSiteDeviceTraffic



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsMonitorTraffic := *openapiclient.NewUtilsMonitorTraffic() // UtilsMonitorTraffic |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.MonitorSiteDeviceTraffic(context.Background(), siteId, deviceId).UtilsMonitorTraffic(utilsMonitorTraffic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.MonitorSiteDeviceTraffic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MonitorSiteDeviceTraffic`: WebsocketSessionWithUrl
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.MonitorSiteDeviceTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitorSiteDeviceTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsMonitorTraffic** | [**UtilsMonitorTraffic**](UtilsMonitorTraffic.md) |  | 

### Return type

[**WebsocketSessionWithUrl**](WebsocketSessionWithUrl.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingFromDevice

> WebsocketSession PingFromDevice(ctx, siteId, deviceId).UtilsPing(utilsPing).Execute()

pingFromDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsPing := *openapiclient.NewUtilsPing("1.1.1.1") // UtilsPing | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.PingFromDevice(context.Background(), siteId, deviceId).UtilsPing(utilsPing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.PingFromDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PingFromDevice`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.PingFromDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingFromDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsPing** | [**UtilsPing**](UtilsPing.md) | Request Body | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadoptSiteOctermDevice

> ReadoptSiteOctermDevice(ctx, siteId, deviceId).Execute()

readoptSiteOctermDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesCommonAPI.ReadoptSiteOctermDevice(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.ReadoptSiteOctermDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadoptSiteOctermDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReprovisionSiteOctermDevice

> ReprovisionSiteOctermDevice(ctx, siteId, deviceId).Execute()

readoptSiteOctermDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesCommonAPI.ReprovisionSiteOctermDevice(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.ReprovisionSiteOctermDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReprovisionSiteOctermDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartSiteDevice

> RestartSiteDevice(ctx, siteId, deviceId).UtilsDevicesRestart(utilsDevicesRestart).Execute()

restartSiteDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsDevicesRestart := *openapiclient.NewUtilsDevicesRestart() // UtilsDevicesRestart |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesCommonAPI.RestartSiteDevice(context.Background(), siteId, deviceId).UtilsDevicesRestart(utilsDevicesRestart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.RestartSiteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartSiteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsDevicesRestart** | [**UtilsDevicesRestart**](UtilsDevicesRestart.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSiteLocateDevice

> StartSiteLocateDevice(ctx, siteId, deviceId).LocateSwitch(locateSwitch).Execute()

startSiteLocateDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	locateSwitch := *openapiclient.NewLocateSwitch() // LocateSwitch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesCommonAPI.StartSiteLocateDevice(context.Background(), siteId, deviceId).LocateSwitch(locateSwitch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.StartSiteLocateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSiteLocateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locateSwitch** | [**LocateSwitch**](LocateSwitch.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopSiteLocateDevice

> StopSiteLocateDevice(ctx, siteId, deviceId).Execute()

stopSiteLocateDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesCommonAPI.StopSiteLocateDevice(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.StopSiteLocateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSiteLocateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TracerouteFromDevice

> WebsocketSession TracerouteFromDevice(ctx, siteId, deviceId).UtilsTraceroute(utilsTraceroute).Execute()

tracerouteFromDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsTraceroute := *openapiclient.NewUtilsTraceroute() // UtilsTraceroute | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesCommonAPI.TracerouteFromDevice(context.Background(), siteId, deviceId).UtilsTraceroute(utilsTraceroute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.TracerouteFromDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TracerouteFromDevice`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesCommonAPI.TracerouteFromDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTracerouteFromDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsTraceroute** | [**UtilsTraceroute**](UtilsTraceroute.md) | Request Body | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSiteDeviceSupportFile

> UploadSiteDeviceSupportFile(ctx, siteId, deviceId).UtilsSendSupportLogs(utilsSendSupportLogs).Execute()

uploadSiteDeviceSupportFile



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsSendSupportLogs := *openapiclient.NewUtilsSendSupportLogs() // UtilsSendSupportLogs | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesCommonAPI.UploadSiteDeviceSupportFile(context.Background(), siteId, deviceId).UtilsSendSupportLogs(utilsSendSupportLogs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesCommonAPI.UploadSiteDeviceSupportFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSiteDeviceSupportFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsSendSupportLogs** | [**UtilsSendSupportLogs**](UtilsSendSupportLogs.md) | Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

