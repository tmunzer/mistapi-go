# \SitesMxEdgesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteMxEdgeEvents**](SitesMxEdgesAPI.md#CountSiteMxEdgeEvents) | **Get** /api/v1/sites/{site_id}/mxedges/events/count | countSiteMxEdgeEvents
[**CreateSiteMxEdge**](SitesMxEdgesAPI.md#CreateSiteMxEdge) | **Post** /api/v1/sites/{site_id}/mxedges | createSiteMxEdge
[**DeleteSiteMxEdge**](SitesMxEdgesAPI.md#DeleteSiteMxEdge) | **Delete** /api/v1/sites/{site_id}/mxedges/{mxedge_id} | deleteSiteMxEdge
[**GetSiteMxEdge**](SitesMxEdgesAPI.md#GetSiteMxEdge) | **Get** /api/v1/sites/{site_id}/mxedges/{mxedge_id} | getSiteMxEdge
[**ListSiteMxEdges**](SitesMxEdgesAPI.md#ListSiteMxEdges) | **Get** /api/v1/sites/{site_id}/mxedges | listSiteMxEdges
[**SearchSiteMistEdgeEvents**](SitesMxEdgesAPI.md#SearchSiteMistEdgeEvents) | **Get** /api/v1/sites/{site_id}/mxedges/events/search | searchSiteMistEdgeEvents
[**UpdateSiteMxEdge**](SitesMxEdgesAPI.md#UpdateSiteMxEdge) | **Put** /api/v1/sites/{site_id}/mxedges/{mxedge_id} | updateSiteMxEdge
[**UploadSiteMxEdgeSupportFiles**](SitesMxEdgesAPI.md#UploadSiteMxEdgeSupportFiles) | **Post** /api/v1/sites/{site_id}/mxedges/{mxedge_id}/support | uploadSiteMxEdgeSupportFiles



## CountSiteMxEdgeEvents

> RepsonseCount CountSiteMxEdgeEvents(ctx, siteId).Distinct(distinct).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countSiteMxEdgeEvents



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
	distinct := openapiclient.site_mxedge_events_count_distinct("") // SiteMxedgeEventsCountDistinct |  (optional) (default to "mxedge_id")
	mxedgeId := "mxedgeId_example" // string | mist edge id (optional)
	mxclusterId := "mxclusterId_example" // string | mist edge cluster id (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	service := "service_example" // string | service running on mist edge(mxagent, tunterm etc) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMxEdgesAPI.CountSiteMxEdgeEvents(context.Background(), siteId).Distinct(distinct).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.CountSiteMxEdgeEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteMxEdgeEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesMxEdgesAPI.CountSiteMxEdgeEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteMxEdgeEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteMxedgeEventsCountDistinct**](SiteMxedgeEventsCountDistinct.md) |  | [default to &quot;mxedge_id&quot;]
 **mxedgeId** | **string** | mist edge id | 
 **mxclusterId** | **string** | mist edge cluster id | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **service** | **string** | service running on mist edge(mxagent, tunterm etc) | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

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


## CreateSiteMxEdge

> Mxedge CreateSiteMxEdge(ctx, siteId).Mxedge(mxedge).Execute()

createSiteMxEdge



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
	mxedge := *openapiclient.NewMxedge("ME-100", "Guest") // Mxedge |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMxEdgesAPI.CreateSiteMxEdge(context.Background(), siteId).Mxedge(mxedge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.CreateSiteMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteMxEdge`: Mxedge
	fmt.Fprintf(os.Stdout, "Response from `SitesMxEdgesAPI.CreateSiteMxEdge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteMxEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedge** | [**Mxedge**](Mxedge.md) |  | 

### Return type

[**Mxedge**](Mxedge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteMxEdge

> DeleteSiteMxEdge(ctx, siteId, mxedgeId).Execute()

deleteSiteMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMxEdgesAPI.DeleteSiteMxEdge(context.Background(), siteId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.DeleteSiteMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteMxEdgeRequest struct via the builder pattern


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


## GetSiteMxEdge

> GetSiteMxEdge(ctx, siteId, mxedgeId).Execute()

getSiteMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMxEdgesAPI.GetSiteMxEdge(context.Background(), siteId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.GetSiteMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteMxEdgeRequest struct via the builder pattern


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


## ListSiteMxEdges

> []Mxedge ListSiteMxEdges(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteMxEdges



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
	resp, r, err := apiClient.SitesMxEdgesAPI.ListSiteMxEdges(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.ListSiteMxEdges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteMxEdges`: []Mxedge
	fmt.Fprintf(os.Stdout, "Response from `SitesMxEdgesAPI.ListSiteMxEdges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteMxEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Mxedge**](Mxedge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteMistEdgeEvents

> ResponseMxedgeEventsSearch SearchSiteMistEdgeEvents(ctx, siteId).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchSiteMistEdgeEvents



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
	mxedgeId := "mxedgeId_example" // string | mist edge id (optional)
	mxclusterId := "mxclusterId_example" // string | mist edge cluster id (optional)
	type_ := "type__example" // string | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) (optional)
	service := "service_example" // string | service running on mist edge(mxagent, tunterm etc) (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMxEdgesAPI.SearchSiteMistEdgeEvents(context.Background(), siteId).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.SearchSiteMistEdgeEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteMistEdgeEvents`: ResponseMxedgeEventsSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesMxEdgesAPI.SearchSiteMistEdgeEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteMistEdgeEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgeId** | **string** | mist edge id | 
 **mxclusterId** | **string** | mist edge cluster id | 
 **type_** | **string** | see [listDeviceEventsDefinitions]($e/Constants%20Events/listDeviceEventsDefinitions) | 
 **service** | **string** | service running on mist edge(mxagent, tunterm etc) | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**ResponseMxedgeEventsSearch**](ResponseMxedgeEventsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteMxEdge

> Mxedge UpdateSiteMxEdge(ctx, siteId, mxedgeId).Mxedge(mxedge).Execute()

updateSiteMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mxedge := *openapiclient.NewMxedge("ME-100", "Guest") // Mxedge |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesMxEdgesAPI.UpdateSiteMxEdge(context.Background(), siteId, mxedgeId).Mxedge(mxedge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.UpdateSiteMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteMxEdge`: Mxedge
	fmt.Fprintf(os.Stdout, "Response from `SitesMxEdgesAPI.UpdateSiteMxEdge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteMxEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mxedge** | [**Mxedge**](Mxedge.md) |  | 

### Return type

[**Mxedge**](Mxedge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSiteMxEdgeSupportFiles

> UploadSiteMxEdgeSupportFiles(ctx, siteId, mxedgeId).Execute()

uploadSiteMxEdgeSupportFiles



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesMxEdgesAPI.UploadSiteMxEdgeSupportFiles(context.Background(), siteId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesMxEdgesAPI.UploadSiteMxEdgeSupportFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSiteMxEdgeSupportFilesRequest struct via the builder pattern


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

