# \SitesDevicesWiredAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSiteLocalSwitchPortConfig**](SitesDevicesWiredAPI.md#DeleteSiteLocalSwitchPortConfig) | **Delete** /api/v1/sites/{site_id}/devices/{device_id}/local_port_config | deleteSiteLocalSwitchPortConfig
[**GetSiteSwitchesMetrics**](SitesDevicesWiredAPI.md#GetSiteSwitchesMetrics) | **Get** /api/v1/sites/{site_id}/stats/switches/metrics | getSiteSwitchesMetrics
[**UpdateSiteLocalSwitchPortConfig**](SitesDevicesWiredAPI.md#UpdateSiteLocalSwitchPortConfig) | **Put** /api/v1/sites/{site_id}/devices/{device_id}/local_port_config | updateSiteLocalSwitchPortConfig



## DeleteSiteLocalSwitchPortConfig

> DeleteSiteLocalSwitchPortConfig(ctx, siteId, deviceId).Execute()

deleteSiteLocalSwitchPortConfig



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
	r, err := apiClient.SitesDevicesWiredAPI.DeleteSiteLocalSwitchPortConfig(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredAPI.DeleteSiteLocalSwitchPortConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteLocalSwitchPortConfigRequest struct via the builder pattern


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


## GetSiteSwitchesMetrics

> ResponseSwitchMetrics GetSiteSwitchesMetrics(ctx, siteId).Type_(type_).Scope(scope).SwitchMac(switchMac).Execute()

getSiteSwitchesMetrics



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
	type_ := openapiclient.switch_metric_type("") // SwitchMetricType |  (optional)
	scope := openapiclient.switch_metric_scope("") // SwitchMetricScope |  (optional)
	switchMac := "switchMac_example" // string | switch mac, used only with metric `type`==`active_ports_summary` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesWiredAPI.GetSiteSwitchesMetrics(context.Background(), siteId).Type_(type_).Scope(scope).SwitchMac(switchMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredAPI.GetSiteSwitchesMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSwitchesMetrics`: ResponseSwitchMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesWiredAPI.GetSiteSwitchesMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSwitchesMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**SwitchMetricType**](SwitchMetricType.md) |  | 
 **scope** | [**SwitchMetricScope**](SwitchMetricScope.md) |  | 
 **switchMac** | **string** | switch mac, used only with metric &#x60;type&#x60;&#x3D;&#x3D;&#x60;active_ports_summary&#x60; | 

### Return type

[**ResponseSwitchMetrics**](ResponseSwitchMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteLocalSwitchPortConfig

> UpdateSiteLocalSwitchPortConfig(ctx, siteId, deviceId).JunosPortConfig(junosPortConfig).Execute()

updateSiteLocalSwitchPortConfig



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
	junosPortConfig := *openapiclient.NewJunosPortConfig("Usage_example") // JunosPortConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesDevicesWiredAPI.UpdateSiteLocalSwitchPortConfig(context.Background(), siteId, deviceId).JunosPortConfig(junosPortConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesWiredAPI.UpdateSiteLocalSwitchPortConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateSiteLocalSwitchPortConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **junosPortConfig** | [**JunosPortConfig**](JunosPortConfig.md) |  | 

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

