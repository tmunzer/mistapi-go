# \SitesApplicationsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteApps**](SitesApplicationsAPI.md#CountSiteApps) | **Get** /api/v1/sites/{site_id}/stats/apps/count | countSiteApps
[**ListSiteApps**](SitesApplicationsAPI.md#ListSiteApps) | **Get** /api/v1/sites/{site_id}/apps | listSiteApps



## CountSiteApps

> RepsonseCount CountSiteApps(ctx, siteId).Distinct(distinct).DeviceMac(deviceMac).App(app).Wired(wired).Execute()

countSiteApps



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
	distinct := openapiclient.site_apps_count_distinct("") // SiteAppsCountDistinct | Default for wireless devices is `ap`. Default for wired devices is `device_mac` (optional)
	deviceMac := "deviceMac_example" // string | MAC of the device (optional)
	app := "app_example" // string | Application name (optional)
	wired := "wired_example" // string | If a device is wired or wireless. Default is False. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesApplicationsAPI.CountSiteApps(context.Background(), siteId).Distinct(distinct).DeviceMac(deviceMac).App(app).Wired(wired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesApplicationsAPI.CountSiteApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteApps`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesApplicationsAPI.CountSiteApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteAppsCountDistinct**](SiteAppsCountDistinct.md) | Default for wireless devices is &#x60;ap&#x60;. Default for wired devices is &#x60;device_mac&#x60; | 
 **deviceMac** | **string** | MAC of the device | 
 **app** | **string** | Application name | 
 **wired** | **string** | If a device is wired or wireless. Default is False. | 

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


## ListSiteApps

> []SiteApp ListSiteApps(ctx, siteId).Execute()

listSiteApps



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesApplicationsAPI.ListSiteApps(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesApplicationsAPI.ListSiteApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteApps`: []SiteApp
	fmt.Fprintf(os.Stdout, "Response from `SitesApplicationsAPI.ListSiteApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SiteApp**](SiteApp.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

