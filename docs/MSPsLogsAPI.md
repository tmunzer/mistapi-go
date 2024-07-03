# \MSPsLogsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountMspLogs**](MSPsLogsAPI.md#CountMspLogs) | **Get** /api/v1/msps/{msp_id}/logs/count | countMspLogs
[**ListMspLogs**](MSPsLogsAPI.md#ListMspLogs) | **Get** /api/v1/msps/{msp_id}/logs | listMspLogs



## CountMspLogs

> RepsonseCount CountMspLogs(ctx, mspId).Distinct(distinct).Execute()

countMspLogs



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	distinct := openapiclient.msp_logs_count_distinct("") // MspLogsCountDistinct |  (optional) (default to "admin_name")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsLogsAPI.CountMspLogs(context.Background(), mspId).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsLogsAPI.CountMspLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountMspLogs`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `MSPsLogsAPI.CountMspLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountMspLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**MspLogsCountDistinct**](MspLogsCountDistinct.md) |  | [default to &quot;admin_name&quot;]

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


## ListMspLogs

> ResponseLogSearch ListMspLogs(ctx, mspId).SiteId(siteId).AdminName(adminName).Message(message).Sort(sort).Start(start).End(end).Limit(limit).Page(page).Duration(duration).Execute()

listMspLogs



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	siteId := "siteId_example" // string | site id (optional)
	adminName := "adminName_example" // string | admin name or email (optional)
	message := "message_example" // string | message (optional)
	sort := openapiclient.list_msp_logs_sort("") // ListMspLogsSort | sort order (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsLogsAPI.ListMspLogs(context.Background(), mspId).SiteId(siteId).AdminName(adminName).Message(message).Sort(sort).Start(start).End(end).Limit(limit).Page(page).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsLogsAPI.ListMspLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspLogs`: ResponseLogSearch
	fmt.Fprintf(os.Stdout, "Response from `MSPsLogsAPI.ListMspLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | site id | 
 **adminName** | **string** | admin name or email | 
 **message** | **string** | message | 
 **sort** | [**ListMspLogsSort**](ListMspLogsSort.md) | sort order | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseLogSearch**](ResponseLogSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

