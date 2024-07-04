# \OrgsMxEdgesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgMxEdgeImage**](OrgsMxEdgesAPI.md#AddOrgMxEdgeImage) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/image/{image_number} | addOrgMxEdgeImage
[**AssignOrgMxEdgeToSite**](OrgsMxEdgesAPI.md#AssignOrgMxEdgeToSite) | **Post** /api/v1/orgs/{org_id}/mxedges/assign | assignOrgMxEdgeToSite
[**BounceOrgMxEdgeDataPorts**](OrgsMxEdgesAPI.md#BounceOrgMxEdgeDataPorts) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/services/tunterm/bounce_port | bounceOrgMxEdgeDataPorts
[**ClaimOrgMxEdge**](OrgsMxEdgesAPI.md#ClaimOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges/claim | claimOrgMxEdge
[**ControlOrgMxEdgeServices**](OrgsMxEdgesAPI.md#ControlOrgMxEdgeServices) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/services/{name}/{action} | controlOrgMxEdgeServices
[**CountOrgMxEdges**](OrgsMxEdgesAPI.md#CountOrgMxEdges) | **Get** /api/v1/orgs/{org_id}/mxedges/count | countOrgMxEdges
[**CountOrgSiteMxEdgeEvents**](OrgsMxEdgesAPI.md#CountOrgSiteMxEdgeEvents) | **Get** /api/v1/orgs/{org_id}/mxedges/events/count | countOrgSiteMxEdgeEvents
[**CreateOrgMxEdge**](OrgsMxEdgesAPI.md#CreateOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges | createOrgMxEdge
[**DeleteOrgMxEdge**](OrgsMxEdgesAPI.md#DeleteOrgMxEdge) | **Delete** /api/v1/orgs/{org_id}/mxedges/{mxedge_id} | deleteOrgMxEdge
[**DeleteOrgMxEdgeImage**](OrgsMxEdgesAPI.md#DeleteOrgMxEdgeImage) | **Delete** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/image/{image_number} | deleteOrgMxEdgeImage
[**GetOrgMxEdge**](OrgsMxEdgesAPI.md#GetOrgMxEdge) | **Get** /api/v1/orgs/{org_id}/mxedges/{mxedge_id} | getOrgMxEdge
[**GetOrgMxEdgeStats**](OrgsMxEdgesAPI.md#GetOrgMxEdgeStats) | **Get** /api/v1/orgs/{org_id}/stats/mxedges/{mxedge_id} | getOrgMxEdgeStats
[**GetOrgMxEdgeUpgradeInfo**](OrgsMxEdgesAPI.md#GetOrgMxEdgeUpgradeInfo) | **Get** /api/v1/orgs/{org_id}/mxedges/version | getOrgMxEdgeUpgradeInfo
[**ListOrgMxEdges**](OrgsMxEdgesAPI.md#ListOrgMxEdges) | **Get** /api/v1/orgs/{org_id}/mxedges | listOrgMxEdges
[**ListOrgMxEdgesStats**](OrgsMxEdgesAPI.md#ListOrgMxEdgesStats) | **Get** /api/v1/orgs/{org_id}/stats/mxedges | listOrgMxEdgesStats
[**RestartOrgMxEdge**](OrgsMxEdgesAPI.md#RestartOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/restart | restartOrgMxEdge
[**SearchOrgMistEdgeEvents**](OrgsMxEdgesAPI.md#SearchOrgMistEdgeEvents) | **Get** /api/v1/orgs/{org_id}/mxedges/events/search | searchOrgMistEdgeEvents
[**SearchOrgMxEdges**](OrgsMxEdgesAPI.md#SearchOrgMxEdges) | **Get** /api/v1/orgs/{org_id}/mxedges/search | searchOrgMxEdges
[**UnassignOrgMxEdgeFromSite**](OrgsMxEdgesAPI.md#UnassignOrgMxEdgeFromSite) | **Post** /api/v1/orgs/{org_id}/mxedges/unassign | unassignOrgMxEdgeFromSite
[**UnregisterOrgMxEdge**](OrgsMxEdgesAPI.md#UnregisterOrgMxEdge) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/unregister | unregisterOrgMxEdge
[**UpdateOrgMxEdge**](OrgsMxEdgesAPI.md#UpdateOrgMxEdge) | **Put** /api/v1/orgs/{org_id}/mxedges/{mxedge_id} | updateOrgMxEdge
[**UploadOrgMxEdgeSupportFiles**](OrgsMxEdgesAPI.md#UploadOrgMxEdgeSupportFiles) | **Post** /api/v1/orgs/{org_id}/mxedges/{mxedge_id}/support | uploadOrgMxEdgeSupportFiles



## AddOrgMxEdgeImage

> AddOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber).ImageImport(imageImport).Execute()

addOrgMxEdgeImage



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	imageNumber := int32(56) // int32 | 
	imageImport := *openapiclient.NewImageImport("TODO") // ImageImport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.AddOrgMxEdgeImage(context.Background(), orgId, mxedgeId, imageNumber).ImageImport(imageImport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.AddOrgMxEdgeImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 
**imageNumber** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrgMxEdgeImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **imageImport** | [**ImageImport**](ImageImport.md) |  | 

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


## AssignOrgMxEdgeToSite

> ResponseAssignSuccess AssignOrgMxEdgeToSite(ctx, orgId).MxedgesAssign(mxedgesAssign).Execute()

assignOrgMxEdgeToSite



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
	mxedgesAssign := *openapiclient.NewMxedgesAssign([]string{"MxedgeIds_example"}, "43e9c864-a7e4-4310-8031-d9817d2c5a43") // MxedgesAssign | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.AssignOrgMxEdgeToSite(context.Background(), orgId).MxedgesAssign(mxedgesAssign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.AssignOrgMxEdgeToSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignOrgMxEdgeToSite`: ResponseAssignSuccess
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.AssignOrgMxEdgeToSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrgMxEdgeToSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgesAssign** | [**MxedgesAssign**](MxedgesAssign.md) | Request Body | 

### Return type

[**ResponseAssignSuccess**](ResponseAssignSuccess.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BounceOrgMxEdgeDataPorts

> BounceOrgMxEdgeDataPorts(ctx, orgId, mxedgeId).UtilsTuntermBouncePort(utilsTuntermBouncePort).Execute()

bounceOrgMxEdgeDataPorts



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	utilsTuntermBouncePort := *openapiclient.NewUtilsTuntermBouncePort([]string{"Ports_example"}) // UtilsTuntermBouncePort |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.BounceOrgMxEdgeDataPorts(context.Background(), orgId, mxedgeId).UtilsTuntermBouncePort(utilsTuntermBouncePort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.BounceOrgMxEdgeDataPorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBounceOrgMxEdgeDataPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsTuntermBouncePort** | [**UtilsTuntermBouncePort**](UtilsTuntermBouncePort.md) |  | 

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


## ClaimOrgMxEdge

> ResponseClaimMxEdge ClaimOrgMxEdge(ctx, orgId).CodeString(codeString).Execute()

claimOrgMxEdge



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
	codeString := *openapiclient.NewCodeString("Code_example") // CodeString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.ClaimOrgMxEdge(context.Background(), orgId).CodeString(codeString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.ClaimOrgMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimOrgMxEdge`: ResponseClaimMxEdge
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.ClaimOrgMxEdge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimOrgMxEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeString** | [**CodeString**](CodeString.md) | Request Body | 

### Return type

[**ResponseClaimMxEdge**](ResponseClaimMxEdge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ControlOrgMxEdgeServices

> ControlOrgMxEdgeServices(ctx, orgId, mxedgeId, name, action).Execute()

controlOrgMxEdgeServices



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	name := openapiclient.mxedge_service_name("") // MxedgeServiceName | 
	action := openapiclient.mxedge_service_action("") // MxedgeServiceAction | restart or start or stop

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.ControlOrgMxEdgeServices(context.Background(), orgId, mxedgeId, name, action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.ControlOrgMxEdgeServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 
**name** | [**MxedgeServiceName**](.md) |  | 
**action** | [**MxedgeServiceAction**](.md) | restart or start or stop | 

### Other Parameters

Other parameters are passed through a pointer to a apiControlOrgMxEdgeServicesRequest struct via the builder pattern


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


## CountOrgMxEdges

> RepsonseCount CountOrgMxEdges(ctx, orgId).Distinct(distinct).MxedgeId(mxedgeId).SiteId(siteId).MxclusterId(mxclusterId).Model(model).Distro(distro).TuntermVersion(tuntermVersion).Sort(sort).Stats(stats).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

countOrgMxEdges



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
	distinct := openapiclient.org_mxedge_count_distinct("") // OrgMxedgeCountDistinct |  (optional) (default to "model")
	mxedgeId := "mxedgeId_example" // string | mist edge id (optional)
	siteId := "siteId_example" // string | mist edge site id (optional)
	mxclusterId := "mxclusterId_example" // string | mist edge cluster id (optional)
	model := "model_example" // string | model name (optional)
	distro := "distro_example" // string | debian code name(buster, bullseye) (optional)
	tuntermVersion := "tuntermVersion_example" // string | tunterm version (optional)
	sort := "sort_example" // string | sort options, -prefix represents DESC order, default is -last_seen (optional)
	stats := true // bool | whether to return device stats, default is false (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.CountOrgMxEdges(context.Background(), orgId).Distinct(distinct).MxedgeId(mxedgeId).SiteId(siteId).MxclusterId(mxclusterId).Model(model).Distro(distro).TuntermVersion(tuntermVersion).Sort(sort).Stats(stats).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.CountOrgMxEdges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgMxEdges`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.CountOrgMxEdges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgMxEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgMxedgeCountDistinct**](OrgMxedgeCountDistinct.md) |  | [default to &quot;model&quot;]
 **mxedgeId** | **string** | mist edge id | 
 **siteId** | **string** | mist edge site id | 
 **mxclusterId** | **string** | mist edge cluster id | 
 **model** | **string** | model name | 
 **distro** | **string** | debian code name(buster, bullseye) | 
 **tuntermVersion** | **string** | tunterm version | 
 **sort** | **string** | sort options, -prefix represents DESC order, default is -last_seen | 
 **stats** | **bool** | whether to return device stats, default is false | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

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


## CountOrgSiteMxEdgeEvents

> RepsonseCount CountOrgSiteMxEdgeEvents(ctx, orgId).Distinct(distinct).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countOrgSiteMxEdgeEvents



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
	distinct := openapiclient.org_mxedge_events_count_distinct("") // OrgMxedgeEventsCountDistinct |  (optional) (default to "mxedge_id")
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
	resp, r, err := apiClient.OrgsMxEdgesAPI.CountOrgSiteMxEdgeEvents(context.Background(), orgId).Distinct(distinct).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.CountOrgSiteMxEdgeEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgSiteMxEdgeEvents`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.CountOrgSiteMxEdgeEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgSiteMxEdgeEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgMxedgeEventsCountDistinct**](OrgMxedgeEventsCountDistinct.md) |  | [default to &quot;mxedge_id&quot;]
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


## CreateOrgMxEdge

> Mxedge CreateOrgMxEdge(ctx, orgId).Mxedge(mxedge).Execute()

createOrgMxEdge



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
	mxedge := *openapiclient.NewMxedge("ME-100", "Guest") // Mxedge | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.CreateOrgMxEdge(context.Background(), orgId).Mxedge(mxedge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.CreateOrgMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgMxEdge`: Mxedge
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.CreateOrgMxEdge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgMxEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedge** | [**Mxedge**](Mxedge.md) | Request Body | 

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


## DeleteOrgMxEdge

> DeleteOrgMxEdge(ctx, orgId, mxedgeId).Execute()

deleteOrgMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.DeleteOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.DeleteOrgMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMxEdgeRequest struct via the builder pattern


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


## DeleteOrgMxEdgeImage

> DeleteOrgMxEdgeImage(ctx, orgId, mxedgeId, imageNumber).Execute()

deleteOrgMxEdgeImage



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	imageNumber := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.DeleteOrgMxEdgeImage(context.Background(), orgId, mxedgeId, imageNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.DeleteOrgMxEdgeImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 
**imageNumber** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMxEdgeImageRequest struct via the builder pattern


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


## GetOrgMxEdge

> Mxedge GetOrgMxEdge(ctx, orgId, mxedgeId).Execute()

getOrgMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.GetOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.GetOrgMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMxEdge`: Mxedge
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.GetOrgMxEdge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMxEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Mxedge**](Mxedge.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgMxEdgeStats

> MxedgeStats GetOrgMxEdgeStats(ctx, orgId, mxedgeId).Execute()

getOrgMxEdgeStats



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.GetOrgMxEdgeStats(context.Background(), orgId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.GetOrgMxEdgeStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMxEdgeStats`: MxedgeStats
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.GetOrgMxEdgeStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMxEdgeStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MxedgeStats**](MxedgeStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgMxEdgeUpgradeInfo

> []MxedgeUpgradeInfoItems GetOrgMxEdgeUpgradeInfo(ctx, orgId).Channel(channel).Execute()

getOrgMxEdgeUpgradeInfo



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
	channel := openapiclient.get_org_mxedge_upgrade_info_channel("") // GetOrgMxedgeUpgradeInfoChannel | upgrade channel to follow, stable (default) / beta / alpha (optional) (default to "stable")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.GetOrgMxEdgeUpgradeInfo(context.Background(), orgId).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.GetOrgMxEdgeUpgradeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMxEdgeUpgradeInfo`: []MxedgeUpgradeInfoItems
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.GetOrgMxEdgeUpgradeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMxEdgeUpgradeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | [**GetOrgMxedgeUpgradeInfoChannel**](GetOrgMxedgeUpgradeInfoChannel.md) | upgrade channel to follow, stable (default) / beta / alpha | [default to &quot;stable&quot;]

### Return type

[**[]MxedgeUpgradeInfoItems**](MxedgeUpgradeInfoItems.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMxEdges

> []Mxedge ListOrgMxEdges(ctx, orgId).ForSites(forSites).Limit(limit).Page(page).Execute()

listOrgMxEdges



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
	forSites := openapiclient.mxedge_for_site("") // MxedgeForSite | filter for site level mist edges (optional) (default to "any")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.ListOrgMxEdges(context.Background(), orgId).ForSites(forSites).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.ListOrgMxEdges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgMxEdges`: []Mxedge
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.ListOrgMxEdges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMxEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forSites** | [**MxedgeForSite**](MxedgeForSite.md) | filter for site level mist edges | [default to &quot;any&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

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


## ListOrgMxEdgesStats

> []MxedgeStats ListOrgMxEdgesStats(ctx, orgId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).ForSite(forSite).Execute()

listOrgMxEdgesStats



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
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	forSite := true // bool | filter for site level mist edges (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.ListOrgMxEdgesStats(context.Background(), orgId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).ForSite(forSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.ListOrgMxEdgesStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgMxEdgesStats`: []MxedgeStats
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.ListOrgMxEdgesStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMxEdgesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **forSite** | **bool** | filter for site level mist edges | [default to false]

### Return type

[**[]MxedgeStats**](MxedgeStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartOrgMxEdge

> RestartOrgMxEdge(ctx, orgId, mxedgeId).Execute()

restartOrgMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.RestartOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.RestartOrgMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartOrgMxEdgeRequest struct via the builder pattern


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


## SearchOrgMistEdgeEvents

> ResponseMxedgeEventsSearch SearchOrgMistEdgeEvents(ctx, orgId).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchOrgMistEdgeEvents



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
	resp, r, err := apiClient.OrgsMxEdgesAPI.SearchOrgMistEdgeEvents(context.Background(), orgId).MxedgeId(mxedgeId).MxclusterId(mxclusterId).Type_(type_).Service(service).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.SearchOrgMistEdgeEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgMistEdgeEvents`: ResponseMxedgeEventsSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.SearchOrgMistEdgeEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgMistEdgeEventsRequest struct via the builder pattern


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


## SearchOrgMxEdges

> ResponseMxedgeSearch SearchOrgMxEdges(ctx, orgId).MxedgeId(mxedgeId).SiteId(siteId).MxclusterId(mxclusterId).Model(model).Distro(distro).TuntermVersion(tuntermVersion).Sort(sort).Stats(stats).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

searchOrgMxEdges



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
	mxedgeId := "mxedgeId_example" // string | mist edge id (optional)
	siteId := "siteId_example" // string | mist edge site id (optional)
	mxclusterId := "mxclusterId_example" // string | mist edge cluster id (optional)
	model := "model_example" // string | model name (optional)
	distro := "distro_example" // string | debian code name(buster, bullseye) (optional)
	tuntermVersion := "tuntermVersion_example" // string | tunterm version (optional)
	sort := "sort_example" // string | sort options, -prefix represents DESC order, default is -last_seen (optional)
	stats := true // bool | whether to return device stats, default is false (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.SearchOrgMxEdges(context.Background(), orgId).MxedgeId(mxedgeId).SiteId(siteId).MxclusterId(mxclusterId).Model(model).Distro(distro).TuntermVersion(tuntermVersion).Sort(sort).Stats(stats).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.SearchOrgMxEdges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgMxEdges`: ResponseMxedgeSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.SearchOrgMxEdges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgMxEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgeId** | **string** | mist edge id | 
 **siteId** | **string** | mist edge site id | 
 **mxclusterId** | **string** | mist edge cluster id | 
 **model** | **string** | model name | 
 **distro** | **string** | debian code name(buster, bullseye) | 
 **tuntermVersion** | **string** | tunterm version | 
 **sort** | **string** | sort options, -prefix represents DESC order, default is -last_seen | 
 **stats** | **bool** | whether to return device stats, default is false | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**ResponseMxedgeSearch**](ResponseMxedgeSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignOrgMxEdgeFromSite

> ResponseAssignSuccess UnassignOrgMxEdgeFromSite(ctx, orgId).MxedgesUnassign(mxedgesUnassign).Execute()

unassignOrgMxEdgeFromSite



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
	mxedgesUnassign := *openapiclient.NewMxedgesUnassign([]string{"MxedgeIds_example"}) // MxedgesUnassign | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.UnassignOrgMxEdgeFromSite(context.Background(), orgId).MxedgesUnassign(mxedgesUnassign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.UnassignOrgMxEdgeFromSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnassignOrgMxEdgeFromSite`: ResponseAssignSuccess
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.UnassignOrgMxEdgeFromSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignOrgMxEdgeFromSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgesUnassign** | [**MxedgesUnassign**](MxedgesUnassign.md) | Request Body | 

### Return type

[**ResponseAssignSuccess**](ResponseAssignSuccess.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterOrgMxEdge

> UnregisterOrgMxEdge(ctx, orgId, mxedgeId).Execute()

unregisterOrgMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.UnregisterOrgMxEdge(context.Background(), orgId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.UnregisterOrgMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterOrgMxEdgeRequest struct via the builder pattern


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


## UpdateOrgMxEdge

> Mxedge UpdateOrgMxEdge(ctx, orgId, mxedgeId).Mxedge(mxedge).Execute()

updateOrgMxEdge



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mxedge := *openapiclient.NewMxedge("ME-100", "Guest") // Mxedge | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxEdgesAPI.UpdateOrgMxEdge(context.Background(), orgId, mxedgeId).Mxedge(mxedge).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.UpdateOrgMxEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgMxEdge`: Mxedge
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxEdgesAPI.UpdateOrgMxEdge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMxEdgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mxedge** | [**Mxedge**](Mxedge.md) | Request Body | 

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


## UploadOrgMxEdgeSupportFiles

> UploadOrgMxEdgeSupportFiles(ctx, orgId, mxedgeId).Execute()

uploadOrgMxEdgeSupportFiles



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
	mxedgeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxEdgesAPI.UploadOrgMxEdgeSupportFiles(context.Background(), orgId, mxedgeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxEdgesAPI.UploadOrgMxEdgeSupportFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxedgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadOrgMxEdgeSupportFilesRequest struct via the builder pattern


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

