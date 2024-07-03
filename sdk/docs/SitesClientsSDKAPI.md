# \SitesClientsSDKAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteSdkStats**](SitesClientsSDKAPI.md#GetSiteSdkStats) | **Get** /api/v1/sites/{site_id}/stats/sdkclients/{sdkclient_id} | getSiteSdkStats
[**GetSiteSdkStatsByMap**](SitesClientsSDKAPI.md#GetSiteSdkStatsByMap) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/sdkclients | getSiteSdkStatsByMap



## GetSiteSdkStats

> SdkstatsWirelessClient GetSiteSdkStats(ctx, siteId, sdkclientId).Execute()

getSiteSdkStats



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
	sdkclientId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsSDKAPI.GetSiteSdkStats(context.Background(), siteId, sdkclientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsSDKAPI.GetSiteSdkStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSdkStats`: SdkstatsWirelessClient
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsSDKAPI.GetSiteSdkStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**sdkclientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSdkStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SdkstatsWirelessClient**](SdkstatsWirelessClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSdkStatsByMap

> []SdkclientStat GetSiteSdkStatsByMap(ctx, siteId, mapId).Execute()

getSiteSdkStatsByMap



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesClientsSDKAPI.GetSiteSdkStatsByMap(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesClientsSDKAPI.GetSiteSdkStatsByMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSdkStatsByMap`: []SdkclientStat
	fmt.Fprintf(os.Stdout, "Response from `SitesClientsSDKAPI.GetSiteSdkStatsByMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSdkStatsByMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]SdkclientStat**](SdkclientStat.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

