# \SitesAnomalyAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteAnomalyEvents**](SitesAnomalyAPI.md#GetSiteAnomalyEvents) | **Get** /api/v1/sites/{site_id}/anomaly/{metric} | getSiteAnomalyEvents
[**GetSiteAnomalyEventsForClient**](SitesAnomalyAPI.md#GetSiteAnomalyEventsForClient) | **Get** /api/v1/sites/{site_id}/anomaly/client/{client_mac}/{metric} | getSiteAnomalyEventsForClient
[**GetSiteAnomalyEventsforDevice**](SitesAnomalyAPI.md#GetSiteAnomalyEventsforDevice) | **Get** /api/v1/sites/{site_id}/anomaly/device/{device_mac}/{metric} | getSiteAnomalyEventsforDevice



## GetSiteAnomalyEvents

> ResponseAnomalySearch GetSiteAnomalyEvents(ctx, siteId, metric).Execute()

getSiteAnomalyEvents



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAnomalyAPI.GetSiteAnomalyEvents(context.Background(), siteId, metric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAnomalyAPI.GetSiteAnomalyEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAnomalyEvents`: ResponseAnomalySearch
	fmt.Fprintf(os.Stdout, "Response from `SitesAnomalyAPI.GetSiteAnomalyEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**metric** | **string** | see /api/v1/const/insight_metrics for available metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAnomalyEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseAnomalySearch**](ResponseAnomalySearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAnomalyEventsForClient

> ResponseAnomalySearch GetSiteAnomalyEventsForClient(ctx, siteId, clientMac, metric).Execute()

getSiteAnomalyEventsForClient



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAnomalyAPI.GetSiteAnomalyEventsForClient(context.Background(), siteId, clientMac, metric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAnomalyAPI.GetSiteAnomalyEventsForClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAnomalyEventsForClient`: ResponseAnomalySearch
	fmt.Fprintf(os.Stdout, "Response from `SitesAnomalyAPI.GetSiteAnomalyEventsForClient`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSiteAnomalyEventsForClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResponseAnomalySearch**](ResponseAnomalySearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAnomalyEventsforDevice

> ResponseAnomalySearch GetSiteAnomalyEventsforDevice(ctx, siteId, metric, deviceMac).Execute()

getSiteAnomalyEventsforDevice



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAnomalyAPI.GetSiteAnomalyEventsforDevice(context.Background(), siteId, metric, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAnomalyAPI.GetSiteAnomalyEventsforDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAnomalyEventsforDevice`: ResponseAnomalySearch
	fmt.Fprintf(os.Stdout, "Response from `SitesAnomalyAPI.GetSiteAnomalyEventsforDevice`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSiteAnomalyEventsforDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResponseAnomalySearch**](ResponseAnomalySearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

