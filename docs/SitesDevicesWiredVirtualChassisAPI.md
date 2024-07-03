# \SitesDevicesWiredVirtualChassisAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteVirtualChassis**](SitesDevicesWiredVirtualChassisAPI.md#CreateSiteVirtualChassis) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/vc | createSiteVirtualChassis
[**DeleteSiteVirtualChassis**](SitesDevicesWiredVirtualChassisAPI.md#DeleteSiteVirtualChassis) | **Delete** /api/v1/sites/{site_id}/devices/{device_id}/vc | deleteSiteVirtualChassis
[**GetSiteDeviceVirtualChassis**](SitesDevicesWiredVirtualChassisAPI.md#GetSiteDeviceVirtualChassis) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/vc | getSiteDeviceVirtualChassis
[**SetSiteVcPort**](SitesDevicesWiredVirtualChassisAPI.md#SetSiteVcPort) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/vc/vc_port | setSiteVcPort
[**UpdateSiteVirtualChassisMember**](SitesDevicesWiredVirtualChassisAPI.md#UpdateSiteVirtualChassisMember) | **Put** /api/v1/sites/{site_id}/devices/{device_id}/vc | updateSiteVirtualChassisMember



## CreateSiteVirtualChassis

> ResponseVirtualChassisConfig CreateSiteVirtualChassis(ctx, siteId, deviceId).VirtualChassisConfig(virtualChassisConfig).Execute()

createSiteVirtualChassis



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
	virtualChassisConfig := *openapiclient.NewVirtualChassisConfig() // VirtualChassisConfig | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesWiredVirtualChassisAPI.CreateSiteVirtualChassis(context.Background(), siteId, deviceId).VirtualChassisConfig(virtualChassisConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredVirtualChassisAPI.CreateSiteVirtualChassis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteVirtualChassis`: ResponseVirtualChassisConfig
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesWiredVirtualChassisAPI.CreateSiteVirtualChassis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteVirtualChassisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualChassisConfig** | [**VirtualChassisConfig**](VirtualChassisConfig.md) | Request Body | 

### Return type

[**ResponseVirtualChassisConfig**](ResponseVirtualChassisConfig.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteVirtualChassis

> DeleteSiteVirtualChassis(ctx, siteId, deviceId).Execute()

deleteSiteVirtualChassis



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
	r, err := apiClient.SitesDevicesWiredVirtualChassisAPI.DeleteSiteVirtualChassis(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredVirtualChassisAPI.DeleteSiteVirtualChassis``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteVirtualChassisRequest struct via the builder pattern


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


## GetSiteDeviceVirtualChassis

> ResponseVirtualChassisConfig GetSiteDeviceVirtualChassis(ctx, siteId, deviceId).Execute()

getSiteDeviceVirtualChassis



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
	resp, r, err := apiClient.SitesDevicesWiredVirtualChassisAPI.GetSiteDeviceVirtualChassis(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredVirtualChassisAPI.GetSiteDeviceVirtualChassis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceVirtualChassis`: ResponseVirtualChassisConfig
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesWiredVirtualChassisAPI.GetSiteDeviceVirtualChassis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceVirtualChassisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseVirtualChassisConfig**](ResponseVirtualChassisConfig.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSiteVcPort

> SetSiteVcPort(ctx, siteId, deviceId).VirtualChassisPort(virtualChassisPort).Execute()

setSiteVcPort



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
	virtualChassisPort := *openapiclient.NewVirtualChassisPort([]openapiclient.ConfigVcPortMember{*openapiclient.NewConfigVcPortMember(float32(123))}, openapiclient.virtual_chassis_port_operation("")) // VirtualChassisPort |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesWiredVirtualChassisAPI.SetSiteVcPort(context.Background(), siteId, deviceId).VirtualChassisPort(virtualChassisPort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredVirtualChassisAPI.SetSiteVcPort``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetSiteVcPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualChassisPort** | [**VirtualChassisPort**](VirtualChassisPort.md) |  | 

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


## UpdateSiteVirtualChassisMember

> ResponseVirtualChassisConfig UpdateSiteVirtualChassisMember(ctx, siteId, deviceId).VirtualChassisUpdate(virtualChassisUpdate).Execute()

updateSiteVirtualChassisMember



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
	virtualChassisUpdate := *openapiclient.NewVirtualChassisUpdate() // VirtualChassisUpdate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesWiredVirtualChassisAPI.UpdateSiteVirtualChassisMember(context.Background(), siteId, deviceId).VirtualChassisUpdate(virtualChassisUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredVirtualChassisAPI.UpdateSiteVirtualChassisMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteVirtualChassisMember`: ResponseVirtualChassisConfig
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesWiredVirtualChassisAPI.UpdateSiteVirtualChassisMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteVirtualChassisMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualChassisUpdate** | [**VirtualChassisUpdate**](VirtualChassisUpdate.md) | Request Body | 

### Return type

[**ResponseVirtualChassisConfig**](ResponseVirtualChassisConfig.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

