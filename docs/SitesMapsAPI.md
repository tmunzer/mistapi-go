# \SitesMapsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSiteMapImage**](SitesMapsAPI.md#AddSiteMapImage) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/image | addSiteMapImage
[**BulkAssignSiteApsToMap**](SitesMapsAPI.md#BulkAssignSiteApsToMap) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/set_map | bulkAssignSiteApsToMap
[**CreateSiteMap**](SitesMapsAPI.md#CreateSiteMap) | **Post** /api/v1/sites/{site_id}/maps | createSiteMap
[**DeleteSiteMap**](SitesMapsAPI.md#DeleteSiteMap) | **Delete** /api/v1/sites/{site_id}/maps/{map_id} | deleteSiteMap
[**DeleteSiteMapImage**](SitesMapsAPI.md#DeleteSiteMapImage) | **Delete** /api/v1/sites/{site_id}/maps/{map_id}/image | deleteSiteMapImage
[**GetSiteMap**](SitesMapsAPI.md#GetSiteMap) | **Get** /api/v1/sites/{site_id}/maps/{map_id} | getSiteMap
[**ImportSiteMaps**](SitesMapsAPI.md#ImportSiteMaps) | **Post** /api/v1/sites/{site_id}/maps/import | importSiteMaps
[**ImportSiteWayfindings**](SitesMapsAPI.md#ImportSiteWayfindings) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/wayfinding/import | importSiteWayfindings
[**ListSiteMaps**](SitesMapsAPI.md#ListSiteMaps) | **Get** /api/v1/sites/{site_id}/maps | listSiteMaps
[**ReplaceSiteMapImage**](SitesMapsAPI.md#ReplaceSiteMapImage) | **Post** /api/v1/sites/{site_id}/maps/{map_id}/replace | replaceSiteMapImage
[**UpdateSiteMap**](SitesMapsAPI.md#UpdateSiteMap) | **Put** /api/v1/sites/{site_id}/maps/{map_id} | updateSiteMap



## AddSiteMapImage

> AddSiteMapImage(ctx, siteId, mapId).File(file).Json(json).Execute()

addSiteMapImage



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
	file := os.NewFile(1234, "some_file") // *os.File | binary file
	json := "json_example" // string | JSON string describing your upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAPI.AddSiteMapImage(context.Background(), siteId, mapId).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.AddSiteMapImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddSiteMapImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | binary file | 
 **json** | **string** | JSON string describing your upload | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkAssignSiteApsToMap

> ResponseSetDevicesMap BulkAssignSiteApsToMap(ctx, siteId, mapId).MacAddresses(macAddresses).Execute()

bulkAssignSiteApsToMap



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
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMapsAPI.BulkAssignSiteApsToMap(context.Background(), siteId, mapId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.BulkAssignSiteApsToMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkAssignSiteApsToMap`: ResponseSetDevicesMap
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAPI.BulkAssignSiteApsToMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkAssignSiteApsToMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macAddresses** | [**MacAddresses**](MacAddresses.md) |  | 

### Return type

[**ResponseSetDevicesMap**](ResponseSetDevicesMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteMap

> ModelMap CreateSiteMap(ctx, siteId).ModelMap(modelMap).Execute()

createSiteMap



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
	modelMap := *openapiclient.NewMap() // ModelMap | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMapsAPI.CreateSiteMap(context.Background(), siteId).ModelMap(modelMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.CreateSiteMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteMap`: ModelMap
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAPI.CreateSiteMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelMap** | [**ModelMap**](ModelMap.md) | Request Body | 

### Return type

[**ModelMap**](ModelMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteMap

> DeleteSiteMap(ctx, siteId, mapId).Execute()

deleteSiteMap



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
	r, err := apiClient.SitesMapsAPI.DeleteSiteMap(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.DeleteSiteMap``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteMapRequest struct via the builder pattern


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


## DeleteSiteMapImage

> DeleteSiteMapImage(ctx, siteId, mapId).Execute()

deleteSiteMapImage



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
	r, err := apiClient.SitesMapsAPI.DeleteSiteMapImage(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.DeleteSiteMapImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteMapImageRequest struct via the builder pattern


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


## GetSiteMap

> ModelMap GetSiteMap(ctx, siteId, mapId).Execute()

getSiteMap



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
	resp, r, err := apiClient.SitesMapsAPI.GetSiteMap(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.GetSiteMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteMap`: ModelMap
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAPI.GetSiteMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ModelMap**](ModelMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSiteMaps

> ResponseMapImport ImportSiteMaps(ctx, siteId).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()

importSiteMaps



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
	autoDeviceprofileAssignment := true // bool | whether to auto assign device to deviceprofile by name (optional)
	csv := os.NewFile(1234, "some_file") // *os.File | csv file for ap name mapping, optional (optional)
	file := os.NewFile(1234, "some_file") // *os.File | ekahau or ibwave file (optional)
	json := *openapiclient.NewMapImportJson(openapiclient.map_import_json_vendor_name("")) // MapImportJson |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMapsAPI.ImportSiteMaps(context.Background(), siteId).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.ImportSiteMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportSiteMaps`: ResponseMapImport
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAPI.ImportSiteMaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSiteMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoDeviceprofileAssignment** | **bool** | whether to auto assign device to deviceprofile by name | 
 **csv** | ***os.File** | csv file for ap name mapping, optional | 
 **file** | ***os.File** | ekahau or ibwave file | 
 **json** | [**MapImportJson**](MapImportJson.md) |  | 

### Return type

[**ResponseMapImport**](ResponseMapImport.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSiteWayfindings

> ImportSiteWayfindings(ctx, siteId, mapId).WayfindingImportJson(wayfindingImportJson).Execute()

importSiteWayfindings



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
	wayfindingImportJson := openapiclient.wayfinding_import_json{MapJibestream: openapiclient.NewMapJibestream("199d6770-0f6f-407a-9bd5-fc33c7840194", "/9Nog3yDzcYj0bY91XJZQLCt+m9DXaIVhx+Ghk3ddd", int32(123), "https://api.jibestream.com", "b069b358-4c97-5319-1f8c-7c5ca64d6ab1", int32(223), float32(4), openapiclient.map_jibestream_vendor_name(""), int32(123))} // WayfindingImportJson | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAPI.ImportSiteWayfindings(context.Background(), siteId, mapId).WayfindingImportJson(wayfindingImportJson).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.ImportSiteWayfindings``: %v\n", err)
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

Other parameters are passed through a pointer to a apiImportSiteWayfindingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wayfindingImportJson** | [**WayfindingImportJson**](WayfindingImportJson.md) | Request Body | 

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


## ListSiteMaps

> []ModelMap ListSiteMaps(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteMaps



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
	resp, r, err := apiClient.SitesMapsAPI.ListSiteMaps(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.ListSiteMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteMaps`: []ModelMap
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAPI.ListSiteMaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]ModelMap**](ModelMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSiteMapImage

> ReplaceSiteMapImage(ctx, siteId, mapId).File(file).Json(json).Execute()

replaceSiteMapImage



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
	file := os.NewFile(1234, "some_file") // *os.File | 
	json := *openapiclient.NewMapSiteReplaceFileJson() // MapSiteReplaceFileJson |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMapsAPI.ReplaceSiteMapImage(context.Background(), siteId, mapId).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.ReplaceSiteMapImage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplaceSiteMapImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 
 **json** | [**MapSiteReplaceFileJson**](MapSiteReplaceFileJson.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteMap

> ModelMap UpdateSiteMap(ctx, siteId, mapId).ModelMap(modelMap).Execute()

updateSiteMap



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
	modelMap := *openapiclient.NewMap() // ModelMap | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMapsAPI.UpdateSiteMap(context.Background(), siteId, mapId).ModelMap(modelMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMapsAPI.UpdateSiteMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteMap`: ModelMap
	fmt.Fprintf(os.Stdout, "Response from `SitesMapsAPI.UpdateSiteMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **modelMap** | [**ModelMap**](ModelMap.md) | Request Body | 

### Return type

[**ModelMap**](ModelMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

