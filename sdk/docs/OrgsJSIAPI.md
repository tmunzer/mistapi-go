# \OrgsJSIAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdoptOrgJsiDevice**](OrgsJSIAPI.md#AdoptOrgJsiDevice) | **Get** /api/v1/orgs/{org_id}/jsi/devices/outbound_ssh_cmd | adoptOrgJsiDevice
[**CreateOrgJsiDeviceShellSession**](OrgsJSIAPI.md#CreateOrgJsiDeviceShellSession) | **Post** /api/v1/orgs/{org_id}/jsi/devices/{device_mac}/shell | createOrgJsiDeviceShellSession
[**ListOrgJsiDevices**](OrgsJSIAPI.md#ListOrgJsiDevices) | **Get** /api/v1/orgs/{org_id}/jsi/devices | listOrgJsiDevices
[**ListOrgJsiPastPurchases**](OrgsJSIAPI.md#ListOrgJsiPastPurchases) | **Get** /api/v1/orgs/{org_id}/jsi/inventory | listOrgJsiPastPurchases



## AdoptOrgJsiDevice

> ResponseDeviceConfigCmd AdoptOrgJsiDevice(ctx, orgId).Execute()

adoptOrgJsiDevice



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
	resp, r, err := apiClient.OrgsJSIAPI.AdoptOrgJsiDevice(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSIAPI.AdoptOrgJsiDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdoptOrgJsiDevice`: ResponseDeviceConfigCmd
	fmt.Fprintf(os.Stdout, "Response from `OrgsJSIAPI.AdoptOrgJsiDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdoptOrgJsiDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseDeviceConfigCmd**](ResponseDeviceConfigCmd.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgJsiDeviceShellSession

> WebsocketSessionWithUrl CreateOrgJsiDeviceShellSession(ctx, orgId, deviceMac).Execute()

createOrgJsiDeviceShellSession



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsJSIAPI.CreateOrgJsiDeviceShellSession(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSIAPI.CreateOrgJsiDeviceShellSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgJsiDeviceShellSession`: WebsocketSessionWithUrl
	fmt.Fprintf(os.Stdout, "Response from `OrgsJSIAPI.CreateOrgJsiDeviceShellSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgJsiDeviceShellSessionRequest struct via the builder pattern


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


## ListOrgJsiDevices

> []JseDevice ListOrgJsiDevices(ctx, orgId).Limit(limit).Page(page).Model(model).Serial(serial).Mac(mac).Execute()

listOrgJsiDevices



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
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)
	model := "model_example" // string | Device model (optional)
	serial := "serial_example" // string | Device serial (optional)
	mac := "mac_example" // string | Device MAC Address (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsJSIAPI.ListOrgJsiDevices(context.Background(), orgId).Limit(limit).Page(page).Model(model).Serial(serial).Mac(mac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSIAPI.ListOrgJsiDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgJsiDevices`: []JseDevice
	fmt.Fprintf(os.Stdout, "Response from `OrgsJSIAPI.ListOrgJsiDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgJsiDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]
 **model** | **string** | Device model | 
 **serial** | **string** | Device serial | 
 **mac** | **string** | Device MAC Address | 

### Return type

[**[]JseDevice**](JseDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgJsiPastPurchases

> []JseInventoryItem ListOrgJsiPastPurchases(ctx, orgId).Limit(limit).Page(page).Model(model).Serial(serial).Execute()

listOrgJsiPastPurchases



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
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)
	model := "model_example" // string |  (optional)
	serial := "serial_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsJSIAPI.ListOrgJsiPastPurchases(context.Background(), orgId).Limit(limit).Page(page).Model(model).Serial(serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSIAPI.ListOrgJsiPastPurchases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgJsiPastPurchases`: []JseInventoryItem
	fmt.Fprintf(os.Stdout, "Response from `OrgsJSIAPI.ListOrgJsiPastPurchases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgJsiPastPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]
 **model** | **string** |  | 
 **serial** | **string** |  | 

### Return type

[**[]JseInventoryItem**](JseInventoryItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

