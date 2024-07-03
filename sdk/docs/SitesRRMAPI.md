# \SitesRRMAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteCurrentChannelPlanning**](SitesRRMAPI.md#GetSiteCurrentChannelPlanning) | **Get** /api/v1/sites/{site_id}/rrm/current | getSiteCurrentChannelPlanning
[**GetSiteCurrentRrmConsiderations**](SitesRRMAPI.md#GetSiteCurrentRrmConsiderations) | **Get** /api/v1/sites/{site_id}/rrm/current/devices/{device_id}/band/{band} | getSiteCurrentRrmConsiderations
[**GetSiteCurrentRrmNeighbors**](SitesRRMAPI.md#GetSiteCurrentRrmNeighbors) | **Get** /api/v1/sites/{site_id}/rrm/neighbors/band/{band} | getSiteCurrentRrmNeighbors
[**GetSiteRrmEvents**](SitesRRMAPI.md#GetSiteRrmEvents) | **Get** /api/v1/sites/{site_id}/rrm/events | getSiteRrmEvents



## GetSiteCurrentChannelPlanning

> Rrm GetSiteCurrentChannelPlanning(ctx, siteId).Execute()

getSiteCurrentChannelPlanning



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRRMAPI.GetSiteCurrentChannelPlanning(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRRMAPI.GetSiteCurrentChannelPlanning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteCurrentChannelPlanning`: Rrm
	fmt.Fprintf(os.Stdout, "Response from `SitesRRMAPI.GetSiteCurrentChannelPlanning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteCurrentChannelPlanningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Rrm**](Rrm.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCurrentRrmConsiderations

> ResponseRrmConsideration GetSiteCurrentRrmConsiderations(ctx, siteId, deviceId, band).Execute()

getSiteCurrentRrmConsiderations



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	band := "band_example" // string | 802.11 Band

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRRMAPI.GetSiteCurrentRrmConsiderations(context.Background(), siteId, deviceId, band).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRRMAPI.GetSiteCurrentRrmConsiderations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteCurrentRrmConsiderations`: ResponseRrmConsideration
	fmt.Fprintf(os.Stdout, "Response from `SitesRRMAPI.GetSiteCurrentRrmConsiderations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 
**band** | **string** | 802.11 Band | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteCurrentRrmConsiderationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResponseRrmConsideration**](ResponseRrmConsideration.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCurrentRrmNeighbors

> ResponseRrmNeighbors GetSiteCurrentRrmNeighbors(ctx, siteId, band).Page(page).Limit(limit).Execute()

getSiteCurrentRrmNeighbors



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
	band := "band_example" // string | 802.11 Band
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRRMAPI.GetSiteCurrentRrmNeighbors(context.Background(), siteId, band).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRRMAPI.GetSiteCurrentRrmNeighbors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteCurrentRrmNeighbors`: ResponseRrmNeighbors
	fmt.Fprintf(os.Stdout, "Response from `SitesRRMAPI.GetSiteCurrentRrmNeighbors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**band** | **string** | 802.11 Band | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteCurrentRrmNeighborsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**ResponseRrmNeighbors**](ResponseRrmNeighbors.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteRrmEvents

> ResponseEventsRrm GetSiteRrmEvents(ctx, siteId).Band(band).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

getSiteRrmEvents



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
	band := openapiclient.dot11_band("") // Dot11Band | 802.11 Band (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRRMAPI.GetSiteRrmEvents(context.Background(), siteId).Band(band).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRRMAPI.GetSiteRrmEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteRrmEvents`: ResponseEventsRrm
	fmt.Fprintf(os.Stdout, "Response from `SitesRRMAPI.GetSiteRrmEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRrmEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **band** | [**Dot11Band**](Dot11Band.md) | 802.11 Band | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseEventsRrm**](ResponseEventsRrm.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

