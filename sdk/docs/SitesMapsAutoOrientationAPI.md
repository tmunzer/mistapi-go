# \SitesMapsAutoOrientationAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSiteApAutoOrient**](SitesMapsAutoOrientationAPI.md#ClearSiteApAutoOrient) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/clear_auto_orient | clearSiteApAutoOrient
[**DeleteSiteApAutoOrientation**](SitesMapsAutoOrientationAPI.md#DeleteSiteApAutoOrientation) | **Delete** /api/v1/sites/{site_id}/maps/{map_id}/auto_orient | deleteSiteApAutoOrientation
[**StartSiteApAutoOrientation**](SitesMapsAutoOrientationAPI.md#StartSiteApAutoOrientation) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/auto_orient | startSiteApAutoOrientation



## ClearSiteApAutoOrient

> ClearSiteApAutoOrient(ctx, siteId, mapId).MacAddresses(macAddresses).Execute()

clearSiteApAutoOrient



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAutoOrientationAPI.ClearSiteApAutoOrient(context.Background(), siteId, mapId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoOrientationAPI.ClearSiteApAutoOrient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearSiteApAutoOrientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macAddresses** | [**MacAddresses**](MacAddresses.md) |  | 

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


## DeleteSiteApAutoOrientation

> DeleteSiteApAutoOrientation(ctx, mapId, siteId).Execute()

deleteSiteApAutoOrientation



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAutoOrientationAPI.DeleteSiteApAutoOrientation(context.Background(), mapId, siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoOrientationAPI.DeleteSiteApAutoOrientation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **string** |  | 
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteApAutoOrientationRequest struct via the builder pattern


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


## StartSiteApAutoOrientation

> ResponseAutoOrientation StartSiteApAutoOrientation(ctx, mapId, siteId).AutoOrient(autoOrient).Execute()

startSiteApAutoOrientation



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	autoOrient := *openapiclient.NewAutoOrient() // AutoOrient |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMapsAutoOrientationAPI.StartSiteApAutoOrientation(context.Background(), mapId, siteId).AutoOrient(autoOrient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoOrientationAPI.StartSiteApAutoOrientation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartSiteApAutoOrientation`: ResponseAutoOrientation
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAutoOrientationAPI.StartSiteApAutoOrientation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mapId** | **string** |  | 
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSiteApAutoOrientationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoOrient** | [**AutoOrient**](AutoOrient.md) |  | 

### Return type

[**ResponseAutoOrientation**](ResponseAutoOrientation.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

