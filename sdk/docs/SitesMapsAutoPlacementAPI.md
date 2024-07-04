# \SitesMapsAutoPlacementAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSiteApAutoplacement**](SitesMapsAutoPlacementAPI.md#ClearSiteApAutoplacement) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/clear_autoplacement | clearSiteApAutoplacement
[**ConfirmSiteApLocalizationData**](SitesMapsAutoPlacementAPI.md#ConfirmSiteApLocalizationData) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/use_auto_ap_values | confirmSiteApLocalizationData
[**DeleteSiteApAutoplacement**](SitesMapsAutoPlacementAPI.md#DeleteSiteApAutoplacement) | **Delete** /api/v1/sites/{site_id}/maps/{map_id}/auto_placement | deleteSiteApAutoplacement
[**GetSiteApAutoPlacement**](SitesMapsAutoPlacementAPI.md#GetSiteApAutoPlacement) | **Get** /api/v1/sites/{site_id}/maps/{map_id}/auto_placement | getSiteApAutoplacement
[**RunSiteApAutoplacement**](SitesMapsAutoPlacementAPI.md#RunSiteApAutoplacement) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/auto_placement | runSiteApAutoplacement



## ClearSiteApAutoplacement

> ClearSiteApAutoplacement(ctx, siteId, mapId).MacAddresses(macAddresses).Execute()

clearSiteApAutoplacement



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
	r, err := apiClient.SitesMapsAutoPlacementAPI.ClearSiteApAutoplacement(context.Background(), siteId, mapId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoPlacementAPI.ClearSiteApAutoplacement``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClearSiteApAutoplacementRequest struct via the builder pattern


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


## ConfirmSiteApLocalizationData

> ConfirmSiteApLocalizationData(ctx, siteId, mapId).UseAutoApValues(useAutoApValues).Execute()

confirmSiteApLocalizationData



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
	useAutoApValues := *openapiclient.NewUseAutoApValues() // UseAutoApValues |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAutoPlacementAPI.ConfirmSiteApLocalizationData(context.Background(), siteId, mapId).UseAutoApValues(useAutoApValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoPlacementAPI.ConfirmSiteApLocalizationData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiConfirmSiteApLocalizationDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **useAutoApValues** | [**UseAutoApValues**](UseAutoApValues.md) |  | 

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


## DeleteSiteApAutoplacement

> DeleteSiteApAutoplacement(ctx, siteId, mapId).Execute()

deleteSiteApAutoplacement



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAutoPlacementAPI.DeleteSiteApAutoplacement(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoPlacementAPI.DeleteSiteApAutoplacement``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteApAutoplacementRequest struct via the builder pattern


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


## GetSiteApAutoPlacement

> ResponseAutoPlacementInfo GetSiteApAutoPlacement(ctx, siteId, mapId).Execute()

getSiteApAutoplacement



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMapsAutoPlacementAPI.GetSiteApAutoPlacement(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoPlacementAPI.GetSiteApAutoPlacement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteApAutoPlacement`: ResponseAutoPlacementInfo
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAutoPlacementAPI.GetSiteApAutoPlacement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteApAutoPlacementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseAutoPlacementInfo**](ResponseAutoPlacementInfo.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSiteApAutoplacement

> RunSiteApAutoplacement(ctx, siteId, mapId).AutoPlacement(autoPlacement).Execute()

runSiteApAutoplacement



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
	autoPlacement := *openapiclient.NewAutoPlacement() // AutoPlacement |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAutoPlacementAPI.RunSiteApAutoplacement(context.Background(), siteId, mapId).AutoPlacement(autoPlacement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAutoPlacementAPI.RunSiteApAutoplacement``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRunSiteApAutoplacementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoPlacement** | [**AutoPlacement**](AutoPlacement.md) |  | 

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

