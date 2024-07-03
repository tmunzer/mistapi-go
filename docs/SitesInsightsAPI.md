# \SitesInsightsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteInsightMetrics**](SitesInsightsAPI.md#GetSiteInsightMetrics) | **Get** /api/v1/sites/{site_id}/insights/{metric} | getSiteInsightMetrics
[**GetSiteInsightMetricsForClient**](SitesInsightsAPI.md#GetSiteInsightMetricsForClient) | **Get** /api/v1/sites/{site_id}/insights/client/{client_mac}/{metric} | getSiteInsightMetricsForClient
[**GetSiteInsightMetricsForDevice**](SitesInsightsAPI.md#GetSiteInsightMetricsForDevice) | **Get** /api/v1/sites/{site_id}/insights/device/{device_mac}/{metric} | getSiteInsightMetricsForDevice



## GetSiteInsightMetrics

> InsightMetrics GetSiteInsightMetrics(ctx, siteId, metric).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()

getSiteInsightMetrics



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
	metric := "metric_example" // string | see /api/v1/const/insight_metrics for available metrics
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	interval := "10m" // string | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesInsightsAPI.GetSiteInsightMetrics(context.Background(), siteId, metric).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesInsightsAPI.GetSiteInsightMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteInsightMetrics`: InsightMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesInsightsAPI.GetSiteInsightMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**metric** | **string** | see /api/v1/const/insight_metrics for available metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteInsightMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **interval** | **string** | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 

### Return type

[**InsightMetrics**](InsightMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteInsightMetricsForClient

> InsightMetrics GetSiteInsightMetricsForClient(ctx, siteId, clientMac, metric).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()

getSiteInsightMetricsForClient



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
	clientMac := "0000000000ab" // string | 
	metric := "metric_example" // string | see /api/v1/const/insight_metrics for available metrics
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	interval := "10m" // string | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesInsightsAPI.GetSiteInsightMetricsForClient(context.Background(), siteId, clientMac, metric).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesInsightsAPI.GetSiteInsightMetricsForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteInsightMetricsForClient`: InsightMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesInsightsAPI.GetSiteInsightMetricsForClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 
**metric** | **string** | see /api/v1/const/insight_metrics for available metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteInsightMetricsForClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **interval** | **string** | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 

### Return type

[**InsightMetrics**](InsightMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteInsightMetricsForDevice

> ResponseDeviceMetrics GetSiteInsightMetricsForDevice(ctx, siteId, metric, deviceMac).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()

getSiteInsightMetricsForDevice



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
	metric := "metric_example" // string | see /api/v1/const/insight_metrics for available metrics
	deviceMac := "0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	interval := "10m" // string | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesInsightsAPI.GetSiteInsightMetricsForDevice(context.Background(), siteId, metric, deviceMac).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Interval(interval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesInsightsAPI.GetSiteInsightMetricsForDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteInsightMetricsForDevice`: ResponseDeviceMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesInsightsAPI.GetSiteInsightMetricsForDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**metric** | **string** | see /api/v1/const/insight_metrics for available metrics | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteInsightMetricsForDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **interval** | **string** | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 

### Return type

[**ResponseDeviceMetrics**](ResponseDeviceMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

