# \SitesLocationAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSiteMlOverwriteForDevice**](SitesLocationAPI.md#ClearSiteMlOverwriteForDevice) | **Delete** /api/v1/sites/{site_id}/location/ml/device/{device_id} | clearSiteMlOverwriteForDevice
[**ClearSiteMlOverwriteForMap**](SitesLocationAPI.md#ClearSiteMlOverwriteForMap) | **Delete** /api/v1/sites/{site_id}/location/ml/map/{map_id} | clearSiteMlOverwriteForMap
[**GetSiteBeamCoverageOverview**](SitesLocationAPI.md#GetSiteBeamCoverageOverview) | **Get** /api/v1/sites/{site_id}/location/coverage | getSiteBeamCoverageOverview
[**GetSiteDefaultPlfForModels**](SitesLocationAPI.md#GetSiteDefaultPlfForModels) | **Get** /api/v1/sites/{site_id}/location/ml/defaults | getSiteDefaultPlfForModels
[**GetSiteMachineLearningCurrentStat**](SitesLocationAPI.md#GetSiteMachineLearningCurrentStat) | **Get** /api/v1/sites/{site_id}/location/ml/current | getSiteMachineLearningCurrentStat
[**OverwriteSiteMlForDevice**](SitesLocationAPI.md#OverwriteSiteMlForDevice) | **Put** /api/v1/sites/{site_id}/location/ml/device/{device_id} | overwriteSiteMlForDevice
[**OverwriteSiteMlForMap**](SitesLocationAPI.md#OverwriteSiteMlForMap) | **Put** /api/v1/sites/{site_id}/location/ml/map/{map_id} | overwriteSiteMlForMap
[**ResetSiteMlStatsByMap**](SitesLocationAPI.md#ResetSiteMlStatsByMap) | **Post** /api/v1/sites/{site_id}/location/ml/reset/map/{map_id} | resetSiteMlStatsByMap



## ClearSiteMlOverwriteForDevice

> ClearSiteMlOverwriteForDevice(ctx, siteId, deviceId).Execute()

clearSiteMlOverwriteForDevice



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesLocationAPI.ClearSiteMlOverwriteForDevice(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.ClearSiteMlOverwriteForDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearSiteMlOverwriteForDeviceRequest struct via the builder pattern


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


## ClearSiteMlOverwriteForMap

> ClearSiteMlOverwriteForMap(ctx, siteId, mapId).Execute()

clearSiteMlOverwriteForMap



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
	r, err := apiClient.SitesLocationAPI.ClearSiteMlOverwriteForMap(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.ClearSiteMlOverwriteForMap``: %v\n", err)
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

Other parameters are passed through a pointer to a apiClearSiteMlOverwriteForMapRequest struct via the builder pattern


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


## GetSiteBeamCoverageOverview

> ResponseLocationCoverage GetSiteBeamCoverageOverview(ctx, siteId).MapId(mapId).Type_(type_).ClientType(clientType).Duration(duration).Resolution(resolution).Start(start).End(end).Execute()

getSiteBeamCoverageOverview



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
	mapId := "00000000-0000-0000-0000-000000000000" // string | map_id (filter by map_id) (optional)
	type_ := openapiclient.rf_client_type("") // RfClientType |  (optional)
	clientType := "clientType_example" // string | client_type (as filter. optional) (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	resolution := openapiclient.resolution("") // Resolution |  (optional) (default to "default")
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesLocationAPI.GetSiteBeamCoverageOverview(context.Background(), siteId).MapId(mapId).Type_(type_).ClientType(clientType).Duration(duration).Resolution(resolution).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.GetSiteBeamCoverageOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteBeamCoverageOverview`: ResponseLocationCoverage
	fmt.Fprintf(os.Stdout, "Response from `SitesLocationAPI.GetSiteBeamCoverageOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteBeamCoverageOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapId** | **string** | map_id (filter by map_id) | 
 **type_** | [**RfClientType**](RfClientType.md) |  | 
 **clientType** | **string** | client_type (as filter. optional) | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **resolution** | [**Resolution**](Resolution.md) |  | [default to &quot;default&quot;]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 

### Return type

[**ResponseLocationCoverage**](ResponseLocationCoverage.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDefaultPlfForModels

> []map[string]interface{} GetSiteDefaultPlfForModels(ctx, siteId).Execute()

getSiteDefaultPlfForModels



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesLocationAPI.GetSiteDefaultPlfForModels(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.GetSiteDefaultPlfForModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDefaultPlfForModels`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SitesLocationAPI.GetSiteDefaultPlfForModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDefaultPlfForModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteMachineLearningCurrentStat

> []map[string]interface{} GetSiteMachineLearningCurrentStat(ctx, siteId).MapId(mapId).Execute()

getSiteMachineLearningCurrentStat



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
	mapId := "00000000-0000-0000-0000-000000000000" // string | map_id (as filter, optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesLocationAPI.GetSiteMachineLearningCurrentStat(context.Background(), siteId).MapId(mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.GetSiteMachineLearningCurrentStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteMachineLearningCurrentStat`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SitesLocationAPI.GetSiteMachineLearningCurrentStat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMachineLearningCurrentStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapId** | **string** | map_id (as filter, optional) | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteSiteMlForDevice

> []map[string]interface{} OverwriteSiteMlForDevice(ctx, siteId, deviceId).RequestBody(requestBody).Execute()

overwriteSiteMlForDevice



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
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	requestBody := map[string]MlOverwriteAdditionalProperties{"key": *openapiclient.NewMlOverwriteAdditionalProperties()} // map[string]MlOverwriteAdditionalProperties | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesLocationAPI.OverwriteSiteMlForDevice(context.Background(), siteId, deviceId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.OverwriteSiteMlForDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverwriteSiteMlForDevice`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SitesLocationAPI.OverwriteSiteMlForDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteSiteMlForDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | [**map[string]MlOverwriteAdditionalProperties**](MlOverwriteAdditionalProperties.md) | Request Body | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverwriteSiteMlForMap

> []map[string]interface{} OverwriteSiteMlForMap(ctx, siteId, mapId).RequestBody(requestBody).Execute()

overwriteSiteMlForMap



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
	requestBody := map[string]MlOverwriteAdditionalProperties{"key": *openapiclient.NewMlOverwriteAdditionalProperties()} // map[string]MlOverwriteAdditionalProperties | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesLocationAPI.OverwriteSiteMlForMap(context.Background(), siteId, mapId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.OverwriteSiteMlForMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OverwriteSiteMlForMap`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SitesLocationAPI.OverwriteSiteMlForMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOverwriteSiteMlForMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | [**map[string]MlOverwriteAdditionalProperties**](MlOverwriteAdditionalProperties.md) | Request Body | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSiteMlStatsByMap

> ResetSiteMlStatsByMap(ctx, siteId, mapId).Execute()

resetSiteMlStatsByMap



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
	r, err := apiClient.SitesLocationAPI.ResetSiteMlStatsByMap(context.Background(), siteId, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesLocationAPI.ResetSiteMlStatsByMap``: %v\n", err)
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

Other parameters are passed through a pointer to a apiResetSiteMlStatsByMapRequest struct via the builder pattern


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

