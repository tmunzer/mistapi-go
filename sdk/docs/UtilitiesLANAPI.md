# \UtilitiesLANAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CableTestFromSwitch**](UtilitiesLANAPI.md#CableTestFromSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/cable_test | cableTestFromSwitch
[**ClearAllLearnedMacsFromPortOnSwitch**](UtilitiesLANAPI.md#ClearAllLearnedMacsFromPortOnSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_macs | clearAllLearnedMacsFromPortOnSwitch
[**ClearBpduErrosFromPortsOnSwitch**](UtilitiesLANAPI.md#ClearBpduErrosFromPortsOnSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_bpdu_error | clearBpduErrosFromPortsOnSwitch
[**CreateSiteDeviceSnapshot**](UtilitiesLANAPI.md#CreateSiteDeviceSnapshot) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/snapshot | createSiteDeviceSnapshot
[**PollSiteSwitchStats**](UtilitiesLANAPI.md#PollSiteSwitchStats) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/poll_stats | pollSiteSwitchStats
[**PortsBounceFromSwitch**](UtilitiesLANAPI.md#PortsBounceFromSwitch) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/bounce_port | portsBounceFromSwitch
[**UpgradeDeviceBios**](UtilitiesLANAPI.md#UpgradeDeviceBios) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/upgrade_bios | upgradeDeviceBios
[**UpgradeDeviceFPGA**](UtilitiesLANAPI.md#UpgradeDeviceFPGA) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/upgrade_fpga | upgradeDeviceFPGA
[**UpgradeSiteDevicesBios**](UtilitiesLANAPI.md#UpgradeSiteDevicesBios) | **Post** /api/v1/sites/{site_id}/devices/upgrade_bios | upgradeSiteDevicesBios
[**UpgradeSiteDevicesFpga**](UtilitiesLANAPI.md#UpgradeSiteDevicesFpga) | **Post** /api/v1/sites/{site_id}/devices/upgrade_fpga | upgradeSiteDevicesFpga



## CableTestFromSwitch

> WebsocketSession CableTestFromSwitch(ctx, siteId, deviceId).UtilsCableTest(utilsCableTest).Execute()

cableTestFromSwitch



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
	utilsCableTest := *openapiclient.NewUtilsCableTest("Port_example") // UtilsCableTest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesLANAPI.CableTestFromSwitch(context.Background(), siteId, deviceId).UtilsCableTest(utilsCableTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.CableTestFromSwitch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CableTestFromSwitch`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesLANAPI.CableTestFromSwitch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCableTestFromSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsCableTest** | [**UtilsCableTest**](UtilsCableTest.md) |  | 

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


## ClearAllLearnedMacsFromPortOnSwitch

> ClearAllLearnedMacsFromPortOnSwitch(ctx, siteId, deviceId).UtilsClearMacs(utilsClearMacs).Execute()

clearAllLearnedMacsFromPortOnSwitch



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
	utilsClearMacs := *openapiclient.NewUtilsClearMacs() // UtilsClearMacs |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesLANAPI.ClearAllLearnedMacsFromPortOnSwitch(context.Background(), siteId, deviceId).UtilsClearMacs(utilsClearMacs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.ClearAllLearnedMacsFromPortOnSwitch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClearAllLearnedMacsFromPortOnSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsClearMacs** | [**UtilsClearMacs**](UtilsClearMacs.md) |  | 

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


## ClearBpduErrosFromPortsOnSwitch

> ClearBpduErrosFromPortsOnSwitch(ctx, siteId, deviceId).UtilsClearBpdu(utilsClearBpdu).Execute()

clearBpduErrosFromPortsOnSwitch



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
	utilsClearBpdu := *openapiclient.NewUtilsClearBpdu() // UtilsClearBpdu |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesLANAPI.ClearBpduErrosFromPortsOnSwitch(context.Background(), siteId, deviceId).UtilsClearBpdu(utilsClearBpdu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.ClearBpduErrosFromPortsOnSwitch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClearBpduErrosFromPortsOnSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsClearBpdu** | [**UtilsClearBpdu**](UtilsClearBpdu.md) |  | 

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


## CreateSiteDeviceSnapshot

> ResponseDeviceSnapshot CreateSiteDeviceSnapshot(ctx, siteId, deviceId).Execute()

createSiteDeviceSnapshot



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
	resp, r, err := apiClient.UtilitiesLANAPI.CreateSiteDeviceSnapshot(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.CreateSiteDeviceSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteDeviceSnapshot`: ResponseDeviceSnapshot
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesLANAPI.CreateSiteDeviceSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteDeviceSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeviceSnapshot**](ResponseDeviceSnapshot.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PollSiteSwitchStats

> PollSiteSwitchStats(ctx, siteId, deviceId).Execute()

pollSiteSwitchStats



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
	r, err := apiClient.UtilitiesLANAPI.PollSiteSwitchStats(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.PollSiteSwitchStats``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPollSiteSwitchStatsRequest struct via the builder pattern


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


## PortsBounceFromSwitch

> PortsBounceFromSwitch(ctx, siteId, deviceId).UtilsBouncePort(utilsBouncePort).Execute()

portsBounceFromSwitch



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
	utilsBouncePort := *openapiclient.NewUtilsBouncePort() // UtilsBouncePort | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesLANAPI.PortsBounceFromSwitch(context.Background(), siteId, deviceId).UtilsBouncePort(utilsBouncePort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.PortsBounceFromSwitch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPortsBounceFromSwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsBouncePort** | [**UtilsBouncePort**](UtilsBouncePort.md) | Request Body | 

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


## UpgradeDeviceBios

> ResponseDeviceBiosUpgrade UpgradeDeviceBios(ctx, siteId, deviceId).UpgradeBios(upgradeBios).Execute()

upgradeDeviceBios



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
	upgradeBios := *openapiclient.NewUpgradeBios() // UpgradeBios |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesLANAPI.UpgradeDeviceBios(context.Background(), siteId, deviceId).UpgradeBios(upgradeBios).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.UpgradeDeviceBios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeDeviceBios`: ResponseDeviceBiosUpgrade
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesLANAPI.UpgradeDeviceBios`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeDeviceBiosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upgradeBios** | [**UpgradeBios**](UpgradeBios.md) |  | 

### Return type

[**ResponseDeviceBiosUpgrade**](ResponseDeviceBiosUpgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeDeviceFPGA

> ResponseDeviceBiosUpgrade UpgradeDeviceFPGA(ctx, siteId, deviceId).UpgradeFpga(upgradeFpga).Execute()

upgradeDeviceFPGA



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
	upgradeFpga := *openapiclient.NewUpgradeFpga() // UpgradeFpga |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesLANAPI.UpgradeDeviceFPGA(context.Background(), siteId, deviceId).UpgradeFpga(upgradeFpga).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.UpgradeDeviceFPGA``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeDeviceFPGA`: ResponseDeviceBiosUpgrade
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesLANAPI.UpgradeDeviceFPGA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeDeviceFPGARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upgradeFpga** | [**UpgradeFpga**](UpgradeFpga.md) |  | 

### Return type

[**ResponseDeviceBiosUpgrade**](ResponseDeviceBiosUpgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSiteDevicesBios

> UpgradeSiteDevicesBios(ctx, siteId).UpgradeBiosMulti(upgradeBiosMulti).Execute()

upgradeSiteDevicesBios



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
	upgradeBiosMulti := *openapiclient.NewUpgradeBiosMulti() // UpgradeBiosMulti |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesLANAPI.UpgradeSiteDevicesBios(context.Background(), siteId).UpgradeBiosMulti(upgradeBiosMulti).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.UpgradeSiteDevicesBios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSiteDevicesBiosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeBiosMulti** | [**UpgradeBiosMulti**](UpgradeBiosMulti.md) |  | 

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


## UpgradeSiteDevicesFpga

> UpgradeSiteDevicesFpga(ctx, siteId).UpgradeFpgaMulti(upgradeFpgaMulti).Execute()

upgradeSiteDevicesFpga



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
	upgradeFpgaMulti := *openapiclient.NewUpgradeFpgaMulti() // UpgradeFpgaMulti |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesLANAPI.UpgradeSiteDevicesFpga(context.Background(), siteId).UpgradeFpgaMulti(upgradeFpgaMulti).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLANAPI.UpgradeSiteDevicesFpga``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSiteDevicesFpgaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeFpgaMulti** | [**UpgradeFpgaMulti**](UpgradeFpgaMulti.md) |  | 

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

