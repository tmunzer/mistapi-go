# \SitesAssetFiltersAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteAssetFilters**](SitesAssetFiltersAPI.md#CreateSiteAssetFilters) | **Post** /api/v1/sites/{site_id}/assetfilters | createSiteAssetFilters
[**DeleteSiteAssetFilter**](SitesAssetFiltersAPI.md#DeleteSiteAssetFilter) | **Delete** /api/v1/sites/{site_id}/assetfilters/{assetfilter_id} | deleteSiteAssetFilter
[**GetSiteAssetFilter**](SitesAssetFiltersAPI.md#GetSiteAssetFilter) | **Get** /api/v1/sites/{site_id}/assetfilters/{assetfilter_id} | getSiteAssetFilter
[**ListSiteAssetFilters**](SitesAssetFiltersAPI.md#ListSiteAssetFilters) | **Get** /api/v1/sites/{site_id}/assetfilters | listSiteAssetFilters
[**UpdateSiteAssetFilter**](SitesAssetFiltersAPI.md#UpdateSiteAssetFilter) | **Put** /api/v1/sites/{site_id}/assetfilters/{assetfilter_id} | updateSiteAssetFilter



## CreateSiteAssetFilters

> AssetFilter CreateSiteAssetFilters(ctx, siteId).AssetFilter(assetFilter).Execute()

createSiteAssetFilters



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
	assetFilter := *openapiclient.NewAssetFilter("Visitor Tags") // AssetFilter | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetFiltersAPI.CreateSiteAssetFilters(context.Background(), siteId).AssetFilter(assetFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetFiltersAPI.CreateSiteAssetFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteAssetFilters`: AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetFiltersAPI.CreateSiteAssetFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteAssetFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetFilter** | [**AssetFilter**](AssetFilter.md) | Request Body | 

### Return type

[**AssetFilter**](AssetFilter.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteAssetFilter

> DeleteSiteAssetFilter(ctx, siteId, assetfilterId).Execute()

deleteSiteAssetFilter



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
	assetfilterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAssetFiltersAPI.DeleteSiteAssetFilter(context.Background(), siteId, assetfilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetFiltersAPI.DeleteSiteAssetFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**assetfilterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteAssetFilterRequest struct via the builder pattern


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


## GetSiteAssetFilter

> AssetFilter GetSiteAssetFilter(ctx, siteId, assetfilterId).Execute()

getSiteAssetFilter



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
	assetfilterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetFiltersAPI.GetSiteAssetFilter(context.Background(), siteId, assetfilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetFiltersAPI.GetSiteAssetFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAssetFilter`: AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetFiltersAPI.GetSiteAssetFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**assetfilterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAssetFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AssetFilter**](AssetFilter.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteAssetFilters

> []AssetFilter ListSiteAssetFilters(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteAssetFilters



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetFiltersAPI.ListSiteAssetFilters(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetFiltersAPI.ListSiteAssetFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteAssetFilters`: []AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetFiltersAPI.ListSiteAssetFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteAssetFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]AssetFilter**](AssetFilter.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteAssetFilter

> AssetFilter UpdateSiteAssetFilter(ctx, siteId, assetfilterId).AssetFilter(assetFilter).Execute()

updateSiteAssetFilter



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
	assetfilterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetFilter := *openapiclient.NewAssetFilter("Visitor Tags") // AssetFilter | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetFiltersAPI.UpdateSiteAssetFilter(context.Background(), siteId, assetfilterId).AssetFilter(assetFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetFiltersAPI.UpdateSiteAssetFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteAssetFilter`: AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetFiltersAPI.UpdateSiteAssetFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**assetfilterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteAssetFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assetFilter** | [**AssetFilter**](AssetFilter.md) | Request Body | 

### Return type

[**AssetFilter**](AssetFilter.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

