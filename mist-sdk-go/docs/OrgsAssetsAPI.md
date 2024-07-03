# \OrgsAssetsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgAssetsByDistanceField**](OrgsAssetsAPI.md#CountOrgAssetsByDistanceField) | **Get** /api/v1/orgs/{org_id}/stats/assets/count | countOrgAssetsByDistanceField
[**CreateOrgAsset**](OrgsAssetsAPI.md#CreateOrgAsset) | **Post** /api/v1/orgs/{org_id}/assets | createOrgAsset
[**DeleteOrgAsset**](OrgsAssetsAPI.md#DeleteOrgAsset) | **Delete** /api/v1/orgs/{org_id}/assets/{asset_id} | deleteOrgAsset
[**GetOrgAsset**](OrgsAssetsAPI.md#GetOrgAsset) | **Get** /api/v1/orgs/{org_id}/assets/{asset_id} | getOrgAsset
[**ImportOrgAssets**](OrgsAssetsAPI.md#ImportOrgAssets) | **Post** /api/v1/orgs/{org_id}/assets/import | importOrgAssets
[**ListOrgAssets**](OrgsAssetsAPI.md#ListOrgAssets) | **Get** /api/v1/orgs/{org_id}/assets | listOrgAssets
[**ListOrgAssetsStats**](OrgsAssetsAPI.md#ListOrgAssetsStats) | **Get** /api/v1/orgs/{org_id}/stats/assets | listOrgAssetsStats
[**SearchOrgAssets**](OrgsAssetsAPI.md#SearchOrgAssets) | **Get** /api/v1/orgs/{org_id}/stats/assets/search | searchOrgAssets
[**UpdateOrgAsset**](OrgsAssetsAPI.md#UpdateOrgAsset) | **Put** /api/v1/orgs/{org_id}/assets/{asset_id} | updateOrgAsset



## CountOrgAssetsByDistanceField

> RepsonseCount CountOrgAssetsByDistanceField(ctx, orgId).Distinct(distinct).Execute()

countOrgAssetsByDistanceField



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	distinct := openapiclient.org_asset_count_distinct("") // OrgAssetCountDistinct |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetsAPI.CountOrgAssetsByDistanceField(context.Background(), orgId).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.CountOrgAssetsByDistanceField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgAssetsByDistanceField`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetsAPI.CountOrgAssetsByDistanceField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgAssetsByDistanceFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgAssetCountDistinct**](OrgAssetCountDistinct.md) |  | 

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


## CreateOrgAsset

> Asset CreateOrgAsset(ctx, orgId).Asset(asset).Execute()

createOrgAsset



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	asset := *openapiclient.NewAsset("Mac_example", "Name_example") // Asset | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetsAPI.CreateOrgAsset(context.Background(), orgId).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.CreateOrgAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetsAPI.CreateOrgAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAssetRequest struct via the builder pattern


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


## DeleteOrgAsset

> DeleteOrgAsset(ctx, orgId, assetId).Execute()

deleteOrgAsset



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAssetsAPI.DeleteOrgAsset(context.Background(), orgId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.DeleteOrgAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgAssetRequest struct via the builder pattern


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


## GetOrgAsset

> Asset GetOrgAsset(ctx, orgId, assetId).Execute()

getOrgAsset



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetsAPI.GetOrgAsset(context.Background(), orgId, assetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.GetOrgAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetsAPI.GetOrgAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAssetRequest struct via the builder pattern


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


## ImportOrgAssets

> ImportOrgAssets(ctx, orgId).AssetImport(assetImport).Execute()

importOrgAssets



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetImport := []openapiclient.AssetImport{*openapiclient.NewAssetImport("Mac_example", "Name_example")} // []AssetImport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAssetsAPI.ImportOrgAssets(context.Background(), orgId).AssetImport(assetImport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.ImportOrgAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportOrgAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## ListOrgAssets

> []Asset ListOrgAssets(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgAssets



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetsAPI.ListOrgAssets(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.ListOrgAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAssets`: []Asset
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetsAPI.ListOrgAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAssetsRequest struct via the builder pattern


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


## ListOrgAssetsStats

> []StatsAsset ListOrgAssetsStats(ctx, orgId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listOrgAssetsStats



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetsAPI.ListOrgAssetsStats(context.Background(), orgId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.ListOrgAssetsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAssetsStats`: []StatsAsset
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetsAPI.ListOrgAssetsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAssetsStatsRequest struct via the builder pattern


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


## SearchOrgAssets

> ResponseStatsAssets SearchOrgAssets(ctx, orgId).SiteId(siteId).Mac(mac).DeviceName(deviceName).Name(name).MapId(mapId).IbeaconUuid(ibeaconUuid).IbeaconMajor(ibeaconMajor).IbeaconMinor(ibeaconMinor).EddystoneUidNamespace(eddystoneUidNamespace).EddystoneUidInstance(eddystoneUidInstance).EddystoneUrl(eddystoneUrl).ApMac(apMac).Beam(beam).Rssi(rssi).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgAssets



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	mac := "mac_example" // string |  (optional)
	deviceName := "deviceName_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	mapId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ibeaconUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	ibeaconMajor := "ibeaconMajor_example" // string |  (optional)
	ibeaconMinor := "ibeaconMinor_example" // string |  (optional)
	eddystoneUidNamespace := "eddystoneUidNamespace_example" // string |  (optional)
	eddystoneUidInstance := "eddystoneUidInstance_example" // string |  (optional)
	eddystoneUrl := "eddystoneUrl_example" // string |  (optional)
	apMac := "apMac_example" // string |  (optional)
	beam := int32(56) // int32 |  (optional)
	rssi := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetsAPI.SearchOrgAssets(context.Background(), orgId).SiteId(siteId).Mac(mac).DeviceName(deviceName).Name(name).MapId(mapId).IbeaconUuid(ibeaconUuid).IbeaconMajor(ibeaconMajor).IbeaconMinor(ibeaconMinor).EddystoneUidNamespace(eddystoneUidNamespace).EddystoneUidInstance(eddystoneUidInstance).EddystoneUrl(eddystoneUrl).ApMac(apMac).Beam(beam).Rssi(rssi).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.SearchOrgAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgAssets`: ResponseStatsAssets
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetsAPI.SearchOrgAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** |  | 
 **mac** | **string** |  | 
 **deviceName** | **string** |  | 
 **name** | **string** |  | 
 **mapId** | **string** |  | 
 **ibeaconUuid** | **string** |  | 
 **ibeaconMajor** | **string** |  | 
 **ibeaconMinor** | **string** |  | 
 **eddystoneUidNamespace** | **string** |  | 
 **eddystoneUidInstance** | **string** |  | 
 **eddystoneUrl** | **string** |  | 
 **apMac** | **string** |  | 
 **beam** | **int32** |  | 
 **rssi** | **int32** |  | 
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


## UpdateOrgAsset

> Asset UpdateOrgAsset(ctx, orgId, assetId).Asset(asset).Execute()

updateOrgAsset



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	assetId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	asset := *openapiclient.NewAsset("Mac_example", "Name_example") // Asset | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAssetsAPI.UpdateOrgAsset(context.Background(), orgId, assetId).Asset(asset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAssetsAPI.UpdateOrgAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgAsset`: Asset
	fmt.Fprintf(os.Stdout, "Response from `OrgsAssetsAPI.UpdateOrgAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgAssetRequest struct via the builder pattern


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

