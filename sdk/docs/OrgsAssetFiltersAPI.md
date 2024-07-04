# \OrgsAssetFiltersAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAssetFilters**](OrgsAssetFiltersAPI.md#CreateOrgAssetFilters) | **Post** /api/v1/orgs/{org_id}/assetfilters | createOrgAssetFilters
[**DeleteOrgAssetFilter**](OrgsAssetFiltersAPI.md#DeleteOrgAssetFilter) | **Delete** /api/v1/orgs/{org_id}/assetfilters/{assetfilter_id} | deleteOrgAssetFilter
[**GetOrgAssetFilter**](OrgsAssetFiltersAPI.md#GetOrgAssetFilter) | **Get** /api/v1/orgs/{org_id}/assetfilters/{assetfilter_id} | getOrgAssetFilter
[**ListOrgAssetFilters**](OrgsAssetFiltersAPI.md#ListOrgAssetFilters) | **Get** /api/v1/orgs/{org_id}/assetfilters | listOrgAssetFilters
[**UpdateOrgAssetFilters**](OrgsAssetFiltersAPI.md#UpdateOrgAssetFilters) | **Put** /api/v1/orgs/{org_id}/assetfilters/{assetfilter_id} | updateOrgAssetFilters



## CreateOrgAssetFilters

> AssetFilter CreateOrgAssetFilters(ctx, orgId).AssetFilter(assetFilter).Execute()

createOrgAssetFilters



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetFilter := *openapiclient.NewAssetFilter("Visitor Tags") // AssetFilter |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetFiltersAPI.CreateOrgAssetFilters(context.Background(), orgId).AssetFilter(assetFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetFiltersAPI.CreateOrgAssetFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgAssetFilters`: AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetFiltersAPI.CreateOrgAssetFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAssetFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assetFilter** | [**AssetFilter**](AssetFilter.md) |  | 

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


## DeleteOrgAssetFilter

> DeleteOrgAssetFilter(ctx, orgId, assetfilterId).Execute()

deleteOrgAssetFilter



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetfilterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAssetFiltersAPI.DeleteOrgAssetFilter(context.Background(), orgId, assetfilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetFiltersAPI.DeleteOrgAssetFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**assetfilterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgAssetFilterRequest struct via the builder pattern


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


## GetOrgAssetFilter

> AssetFilter GetOrgAssetFilter(ctx, orgId, assetfilterId).Execute()

getOrgAssetFilter



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetfilterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetFiltersAPI.GetOrgAssetFilter(context.Background(), orgId, assetfilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetFiltersAPI.GetOrgAssetFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgAssetFilter`: AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetFiltersAPI.GetOrgAssetFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**assetfilterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAssetFilterRequest struct via the builder pattern


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


## ListOrgAssetFilters

> []AssetFilter ListOrgAssetFilters(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgAssetFilters



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetFiltersAPI.ListOrgAssetFilters(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetFiltersAPI.ListOrgAssetFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAssetFilters`: []AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetFiltersAPI.ListOrgAssetFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAssetFiltersRequest struct via the builder pattern


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


## UpdateOrgAssetFilters

> AssetFilter UpdateOrgAssetFilters(ctx, orgId, assetfilterId).AssetFilter(assetFilter).Execute()

updateOrgAssetFilters



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetfilterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetFilter := *openapiclient.NewAssetFilter("Visitor Tags") // AssetFilter | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetFiltersAPI.UpdateOrgAssetFilters(context.Background(), orgId, assetfilterId).AssetFilter(assetFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetFiltersAPI.UpdateOrgAssetFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgAssetFilters`: AssetFilter
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetFiltersAPI.UpdateOrgAssetFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**assetfilterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgAssetFiltersRequest struct via the builder pattern


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

