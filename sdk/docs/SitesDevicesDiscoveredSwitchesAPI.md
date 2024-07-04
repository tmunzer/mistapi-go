# \SitesDevicesDiscoveredSwitchesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteDiscoveredSwitches**](SitesDevicesDiscoveredSwitchesAPI.md#CountSiteDiscoveredSwitches) | **Get** /api/v1/sites/{site_id}/stats/discovered_switches/count | countSiteDiscoveredSwitches
[**GetSiteDiscoveredSwitchesMetrics**](SitesDevicesDiscoveredSwitchesAPI.md#GetSiteDiscoveredSwitchesMetrics) | **Get** /api/v1/sites/{site_id}/stats/discovered_switches/metrics | getSiteDiscoveredSwitchesMetrics
[**SearchSiteDiscoveredSwitches**](SitesDevicesDiscoveredSwitchesAPI.md#SearchSiteDiscoveredSwitches) | **Get** /api/v1/sites/{site_id}/stats/discovered_switches/search | searchSiteDiscoveredSwitches
[**SearchSiteDiscoveredSwitchesMetrics**](SitesDevicesDiscoveredSwitchesAPI.md#SearchSiteDiscoveredSwitchesMetrics) | **Get** /api/v1/sites/{site_id}/stats/discovered_switch_metrics/search | searchSiteDiscoveredSwitchesMetrics



## CountSiteDiscoveredSwitches

> RepsonseCount CountSiteDiscoveredSwitches(ctx, siteId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteDiscoveredSwitches



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
	distinct := openapiclient.site_discovered_switches_count_distinct("") // SiteDiscoveredSwitchesCountDistinct |  (optional) (default to "system_name")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.CountSiteDiscoveredSwitches(context.Background(), siteId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesDiscoveredSwitchesAPI.CountSiteDiscoveredSwitches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteDiscoveredSwitches`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesDiscoveredSwitchesAPI.CountSiteDiscoveredSwitches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteDiscoveredSwitchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteDiscoveredSwitchesCountDistinct**](SiteDiscoveredSwitchesCountDistinct.md) |  | [default to &quot;system_name&quot;]
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


## GetSiteDiscoveredSwitchesMetrics

> ResponseDswitchesMetrics GetSiteDiscoveredSwitchesMetrics(ctx, siteId).Threshold(threshold).SystemName(systemName).Execute()

getSiteDiscoveredSwitchesMetrics



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
	threshold := "threshold_example" // string | configurable # ap per switch threshold, default 12 (optional)
	systemName := "systemName_example" // string | system name for switch level metrics, optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.GetSiteDiscoveredSwitchesMetrics(context.Background(), siteId).Threshold(threshold).SystemName(systemName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesDiscoveredSwitchesAPI.GetSiteDiscoveredSwitchesMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDiscoveredSwitchesMetrics`: ResponseDswitchesMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesDiscoveredSwitchesAPI.GetSiteDiscoveredSwitchesMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDiscoveredSwitchesMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **threshold** | **string** | configurable # ap per switch threshold, default 12 | 
 **systemName** | **string** | system name for switch level metrics, optional | 

### Return type

[**ResponseDswitchesMetrics**](ResponseDswitchesMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteDiscoveredSwitches

> ResponseDiscoveredSwitches SearchSiteDiscoveredSwitches(ctx, siteId).Adopted(adopted).SystemName(systemName).Hostname(hostname).Vendor(vendor).Model(model).Version(version).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteDiscoveredSwitches



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
	adopted := true // bool |  (optional)
	systemName := "systemName_example" // string |  (optional)
	hostname := "hostname_example" // string |  (optional)
	vendor := "vendor_example" // string |  (optional)
	model := "model_example" // string |  (optional)
	version := "version_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitches(context.Background(), siteId).Adopted(adopted).SystemName(systemName).Hostname(hostname).Vendor(vendor).Model(model).Version(version).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteDiscoveredSwitches`: ResponseDiscoveredSwitches
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteDiscoveredSwitchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adopted** | **bool** |  | 
 **systemName** | **string** |  | 
 **hostname** | **string** |  | 
 **vendor** | **string** |  | 
 **model** | **string** |  | 
 **version** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseDiscoveredSwitches**](ResponseDiscoveredSwitches.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteDiscoveredSwitchesMetrics

> ResponseDiscoveredSwitchMetrics SearchSiteDiscoveredSwitchesMetrics(ctx, siteId).Scope(scope).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteDiscoveredSwitchesMetrics



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
	scope := openapiclient.discovered_switches_metric_scope("") // DiscoveredSwitchesMetricScope | metric scope (optional) (default to "site")
	type_ := openapiclient.discovered_switch_metric_type("") // DiscoveredSwitchMetricType | metric type (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitchesMetrics(context.Background(), siteId).Scope(scope).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitchesMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteDiscoveredSwitchesMetrics`: ResponseDiscoveredSwitchMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesDevicesDiscoveredSwitchesAPI.SearchSiteDiscoveredSwitchesMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteDiscoveredSwitchesMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**DiscoveredSwitchesMetricScope**](DiscoveredSwitchesMetricScope.md) | metric scope | [default to &quot;site&quot;]
 **type_** | [**DiscoveredSwitchMetricType**](DiscoveredSwitchMetricType.md) | metric type | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseDiscoveredSwitchMetrics**](ResponseDiscoveredSwitchMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

