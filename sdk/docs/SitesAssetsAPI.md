# \SitesAssetsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteAssets**](SitesAssetsAPI.md#CountSiteAssets) | **Get** /api/v1/sites/{site_id}/stats/assets/count | countSiteAssets
[**CreateSiteAsset**](SitesAssetsAPI.md#CreateSiteAsset) | **Post** /api/v1/sites/{site_id}/assets | createSiteAsset
[**DeleteSiteAsset**](SitesAssetsAPI.md#DeleteSiteAsset) | **Delete** /api/v1/sites/{site_id}/assets/{asset_id} | deleteSiteAsset
[**GetSiteAsset**](SitesAssetsAPI.md#GetSiteAsset) | **Get** /api/v1/sites/{site_id}/assets/{asset_id} | getSiteAsset
[**GetSiteAssetStats**](SitesAssetsAPI.md#GetSiteAssetStats) | **Get** /api/v1/sites/{site_id}/stats/assets/asset_id | getSiteAssetStats
[**GetSiteAssetsOfInterest**](SitesAssetsAPI.md#GetSiteAssetsOfInterest) | **Get** /api/v1/sites/{site_id}/stats/filtered_assets | getSiteAssetsOfInterest
[**GetSiteDiscoveredAssetByMap**](SitesAssetsAPI.md#GetSiteDiscoveredAssetByMap) | **Get** /api/v1/sites/{site_id}/stats/maps/{map_id}/discovered_assets | getSiteDiscoveredAssetByMap
[**ImportSiteAssets**](SitesAssetsAPI.md#ImportSiteAssets) | **Post** /api/v1/sites/{site_id}/assets/import | importSiteAssets
[**ListSiteAssets**](SitesAssetsAPI.md#ListSiteAssets) | **Get** /api/v1/sites/{site_id}/assets | listSiteAssets
[**ListSiteAssetsStats**](SitesAssetsAPI.md#ListSiteAssetsStats) | **Get** /api/v1/sites/{site_id}/stats/assets | listSiteAssetsStats
[**ListSiteDiscoveredAssets**](SitesAssetsAPI.md#ListSiteDiscoveredAssets) | **Get** /api/v1/sites/{site_id}/stats/discovered_assets | listSiteDiscoveredAssets
[**SearchSiteAssets**](SitesAssetsAPI.md#SearchSiteAssets) | **Get** /api/v1/sites/{site_id}/stats/assets/search | searchSiteAssets
[**UpdateSiteAsset**](SitesAssetsAPI.md#UpdateSiteAsset) | **Put** /api/v1/sites/{site_id}/assets/{asset_id} | updateSiteAsset



## CountSiteAssets

> RepsonseCount CountSiteAssets(ctx, siteId).Distinct(distinct).Execute()

countSiteAssets



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
	distinct := openapiclient.site_assets_count_distinct("") // SiteAssetsCountDistinct |  (optional) (default to "map_id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.CountSiteAssets(context.Background(), siteId).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.CountSiteAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteAssets`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.CountSiteAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteAssetsCountDistinct**](SiteAssetsCountDistinct.md) |  | [default to &quot;map_id&quot;]

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


## CreateSiteAsset

> Asset CreateSiteAsset(ctx, siteId).Asset(asset).Execute()

createSiteAsset



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
	asset := *openapiclient.NewAsset("Mac_example", "Name_example") // Asset | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.CreateSiteAsset(context.Background(), siteId).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.CreateSiteAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.CreateSiteAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **asset** | [**Asset**](Asset.md) | Request Body | 

### Return type

[**Asset**](Asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteAsset

> DeleteSiteAsset(ctx, siteId, assetId).Execute()

deleteSiteAsset



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
	assetId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAssetsAPI.DeleteSiteAsset(context.Background(), siteId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.DeleteSiteAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteAssetRequest struct via the builder pattern


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


## GetSiteAsset

> Asset GetSiteAsset(ctx, siteId, assetId).Execute()

getSiteAsset



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
	assetId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.GetSiteAsset(context.Background(), siteId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.GetSiteAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.GetSiteAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Asset**](Asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAssetStats

> StatsAsset GetSiteAssetStats(ctx, siteId).Start(start).End(end).Duration(duration).Execute()

getSiteAssetStats



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
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.GetSiteAssetStats(context.Background(), siteId).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.GetSiteAssetStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAssetStats`: StatsAsset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.GetSiteAssetStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAssetStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**StatsAsset**](StatsAsset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteAssetsOfInterest

> []AssetOfInterest GetSiteAssetsOfInterest(ctx, siteId).Duration(duration).Start(start).End(end).Page(page).Limit(limit).Execute()

getSiteAssetsOfInterest



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
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.GetSiteAssetsOfInterest(context.Background(), siteId).Duration(duration).Start(start).End(end).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.GetSiteAssetsOfInterest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteAssetsOfInterest`: []AssetOfInterest
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.GetSiteAssetsOfInterest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteAssetsOfInterestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]AssetOfInterest**](AssetOfInterest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDiscoveredAssetByMap

> []StatsAsset GetSiteDiscoveredAssetByMap(ctx, siteId, mapId).Execute()

getSiteDiscoveredAssetByMap



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
	resp, r, err := apiClient.SitesAssetsAPI.GetSiteDiscoveredAssetByMap(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.GetSiteDiscoveredAssetByMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDiscoveredAssetByMap`: []StatsAsset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.GetSiteDiscoveredAssetByMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDiscoveredAssetByMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]StatsAsset**](StatsAsset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSiteAssets

> ImportSiteAssets(ctx, siteId).Upsert(upsert).AssetImport(assetImport).Execute()

importSiteAssets



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
	upsert := openapiclient.import_site_assets_upsert("") // ImportSiteAssetsUpsert | API will replace the assets with same mac if provided `upsert`==`True`, otherwise will report in errors in response. (optional) (default to "False")
	assetImport := []openapiclient.AssetImport{*openapiclient.NewAssetImport("Mac_example", "Name_example")} // []AssetImport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAssetsAPI.ImportSiteAssets(context.Background(), siteId).Upsert(upsert).AssetImport(assetImport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.ImportSiteAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSiteAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | [**ImportSiteAssetsUpsert**](ImportSiteAssetsUpsert.md) | API will replace the assets with same mac if provided &#x60;upsert&#x60;&#x3D;&#x3D;&#x60;True&#x60;, otherwise will report in errors in response. | [default to &quot;False&quot;]
 **assetImport** | [**[]AssetImport**](AssetImport.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteAssets

> []Asset ListSiteAssets(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteAssets



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.ListSiteAssets(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.ListSiteAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.ListSiteAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Asset**](Asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteAssetsStats

> []StatsAsset ListSiteAssetsStats(ctx, siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listSiteAssetsStats



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.ListSiteAssetsStats(context.Background(), siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.ListSiteAssetsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteAssetsStats`: []StatsAsset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.ListSiteAssetsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteAssetsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]StatsAsset**](StatsAsset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteDiscoveredAssets

> []Asset ListSiteDiscoveredAssets(ctx, siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listSiteDiscoveredAssets



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.ListSiteDiscoveredAssets(context.Background(), siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.ListSiteDiscoveredAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteDiscoveredAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.ListSiteDiscoveredAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteDiscoveredAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]Asset**](Asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteAssets

> ResponseStatsAssets SearchSiteAssets(ctx, siteId).Mac(mac).MapId(mapId).IbeaconUuid(ibeaconUuid).IbeaconMajor(ibeaconMajor).IbeaconMinor(ibeaconMinor).EddystoneUidNamespace(eddystoneUidNamespace).EddystoneUidInstance(eddystoneUidInstance).EddystoneUrl(eddystoneUrl).DeviceName(deviceName).By(by).Name(name).ApMac(apMac).Beam(beam).Rssi(rssi).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteAssets



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
	mac := "mac_example" // string |  (optional)
	mapId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ibeaconUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ibeaconMajor := int32(56) // int32 |  (optional)
	ibeaconMinor := int32(56) // int32 |  (optional)
	eddystoneUidNamespace := "eddystoneUidNamespace_example" // string |  (optional)
	eddystoneUidInstance := "eddystoneUidInstance_example" // string |  (optional)
	eddystoneUrl := "eddystoneUrl_example" // string |  (optional)
	deviceName := "deviceName_example" // string |  (optional)
	by := "by_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	apMac := "apMac_example" // string |  (optional)
	beam := "beam_example" // string |  (optional)
	rssi := "rssi_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.SearchSiteAssets(context.Background(), siteId).Mac(mac).MapId(mapId).IbeaconUuid(ibeaconUuid).IbeaconMajor(ibeaconMajor).IbeaconMinor(ibeaconMinor).EddystoneUidNamespace(eddystoneUidNamespace).EddystoneUidInstance(eddystoneUidInstance).EddystoneUrl(eddystoneUrl).DeviceName(deviceName).By(by).Name(name).ApMac(apMac).Beam(beam).Rssi(rssi).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.SearchSiteAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteAssets`: ResponseStatsAssets
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.SearchSiteAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** |  | 
 **mapId** | **string** |  | 
 **ibeaconUuid** | **string** |  | 
 **ibeaconMajor** | **int32** |  | 
 **ibeaconMinor** | **int32** |  | 
 **eddystoneUidNamespace** | **string** |  | 
 **eddystoneUidInstance** | **string** |  | 
 **eddystoneUrl** | **string** |  | 
 **deviceName** | **string** |  | 
 **by** | **string** |  | 
 **name** | **string** |  | 
 **apMac** | **string** |  | 
 **beam** | **string** |  | 
 **rssi** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseStatsAssets**](ResponseStatsAssets.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteAsset

> Asset UpdateSiteAsset(ctx, siteId, assetId).Asset(asset).Execute()

updateSiteAsset



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
	assetId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	asset := *openapiclient.NewAsset("Mac_example", "Name_example") // Asset | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAssetsAPI.UpdateSiteAsset(context.Background(), siteId, assetId).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAssetsAPI.UpdateSiteAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `SitesAssetsAPI.UpdateSiteAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asset** | [**Asset**](Asset.md) | Request Body | 

### Return type

[**Asset**](Asset.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

