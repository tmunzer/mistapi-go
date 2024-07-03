# \SitesDevicesWirelessAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteDeviceIotPort**](SitesDevicesWirelessAPI.md#GetSiteDeviceIotPort) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/iot | getSiteDeviceIotPort
[**GetSiteDeviceRadioChannels**](SitesDevicesWirelessAPI.md#GetSiteDeviceRadioChannels) | **Get** /api/v1/sites/{site_id}/devices/ap_channels | getSiteDeviceRadioChannels
[**SetSiteDeviceIotPort**](SitesDevicesWirelessAPI.md#SetSiteDeviceIotPort) | **Put** /api/v1/sites/{site_id}/devices/{device_id}/iot | setSiteDeviceIotPort



## GetSiteDeviceIotPort

> map[string]int32 GetSiteDeviceIotPort(ctx, siteId, deviceId).Execute()

getSiteDeviceIotPort



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
	resp, r, err := apiClient.SitesDevicesWirelessAPI.GetSiteDeviceIotPort(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWirelessAPI.GetSiteDeviceIotPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceIotPort`: map[string]int32
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesWirelessAPI.GetSiteDeviceIotPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceIotPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]int32**

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceRadioChannels

> ResponseDeviceRadioChannels GetSiteDeviceRadioChannels(ctx, siteId).CountryCode(countryCode).Execute()

getSiteDeviceRadioChannels



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
	countryCode := "US" // string | country code for the site (for AP config generation), in [two-character](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesWirelessAPI.GetSiteDeviceRadioChannels(context.Background(), siteId).CountryCode(countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWirelessAPI.GetSiteDeviceRadioChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceRadioChannels`: ResponseDeviceRadioChannels
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesWirelessAPI.GetSiteDeviceRadioChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceRadioChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **countryCode** | **string** | country code for the site (for AP config generation), in [two-character](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | 

### Return type

[**ResponseDeviceRadioChannels**](ResponseDeviceRadioChannels.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSiteDeviceIotPort

> map[string]int32 SetSiteDeviceIotPort(ctx, siteId, deviceId).RequestBody(requestBody).Execute()

setSiteDeviceIotPort



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
	requestBody := map[string]int32{"key": int32(123)} // map[string]int32 | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesWirelessAPI.SetSiteDeviceIotPort(context.Background(), siteId, deviceId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWirelessAPI.SetSiteDeviceIotPort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSiteDeviceIotPort`: map[string]int32
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesWirelessAPI.SetSiteDeviceIotPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSiteDeviceIotPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]int32** | Request Body | 

### Return type

**map[string]int32**

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

