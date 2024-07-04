# \MSPsSLEsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspSle**](MSPsSLEsAPI.md#GetMspSle) | **Get** /api/v1/msps/{msp_id}/insights/{metric} | getMspSle



## GetMspSle

> InsightMetrics GetMspSle(ctx, mspId, metric).Sle(sle).Duration(duration).Interval(interval).Start(start).End(end).Execute()

getMspSle



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	metric := "metric_example" // string | see /api/v1/const/insight_metrics for available metrics
	sle := "sle_example" // string | see /api/v1/const/insight_metrics for more details (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	interval := "10m" // string | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSLEsAPI.GetMspSle(context.Background(), mspId, metric).Sle(sle).Duration(duration).Interval(interval).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSLEsAPI.GetMspSle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspSle`: InsightMetrics
	fmt.Fprintf(os.Stdout, "Response from `MSPsSLEsAPI.GetMspSle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**metric** | **string** | see /api/v1/const/insight_metrics for available metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspSleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sle** | **string** | see /api/v1/const/insight_metrics for more details | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **interval** | **string** | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 

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

