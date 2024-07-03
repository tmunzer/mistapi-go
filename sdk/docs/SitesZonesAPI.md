# \SitesZonesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteZoneSessions**](SitesZonesAPI.md#CountSiteZoneSessions) | **Get** /api/v1/sites/{site_id}/{zone_type}/count | countSiteZoneSessions
[**CreateSiteZone**](SitesZonesAPI.md#CreateSiteZone) | **Post** /api/v1/sites/{site_id}/zones | createSiteZone
[**DeleteSiteZone**](SitesZonesAPI.md#DeleteSiteZone) | **Delete** /api/v1/sites/{site_id}/zones/{zone_id} | deleteSiteZone
[**GetSiteZone**](SitesZonesAPI.md#GetSiteZone) | **Get** /api/v1/sites/{site_id}/zones/{zone_id} | getSiteZone
[**GetSiteZoneStats**](SitesZonesAPI.md#GetSiteZoneStats) | **Get** /api/v1/sites/{site_id}/stats/{zone_type}/{zone_id} | getSiteZoneStats
[**ListSiteZones**](SitesZonesAPI.md#ListSiteZones) | **Get** /api/v1/sites/{site_id}/zones | listSiteZones
[**ListSiteZonesStats**](SitesZonesAPI.md#ListSiteZonesStats) | **Get** /api/v1/sites/{site_id}/stats/zones | listSiteZonesStats
[**SearchSiteZoneSessions**](SitesZonesAPI.md#SearchSiteZoneSessions) | **Get** /api/v1/sites/{site_id}/{zone_type}/visits/search | searchSiteZoneSessions
[**UpdateSiteZone**](SitesZonesAPI.md#UpdateSiteZone) | **Put** /api/v1/sites/{site_id}/zones/{zone_id} | updateSiteZone



## CountSiteZoneSessions

> RepsonseCount CountSiteZoneSessions(ctx, siteId, zoneType).Distinct(distinct).UserType(userType).User(user).ScopeId(scopeId).Scope(scope).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteZoneSessions



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
	zoneType := "zoneType_example" // string | 
	distinct := openapiclient.site_zone_count_distinct("") // SiteZoneCountDistinct |  (optional) (default to "scope_id")
	userType := openapiclient.rf_client_type("") // RfClientType | user type (optional)
	user := "user_example" // string | client MAC / Asset MAC / SDK UUID (optional)
	scopeId := "scopeId_example" // string | if `scope`==`map`/`zone`/`rssizone`, the scope id (optional)
	scope := openapiclient.zone_scope("") // ZoneScope | scope (optional) (default to "site")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesZonesAPI.CountSiteZoneSessions(context.Background(), siteId, zoneType).Distinct(distinct).UserType(userType).User(user).ScopeId(scopeId).Scope(scope).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.CountSiteZoneSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteZoneSessions`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.CountSiteZoneSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**zoneType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteZoneSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **distinct** | [**SiteZoneCountDistinct**](SiteZoneCountDistinct.md) |  | [default to &quot;scope_id&quot;]
 **userType** | [**RfClientType**](RfClientType.md) | user type | 
 **user** | **string** | client MAC / Asset MAC / SDK UUID | 
 **scopeId** | **string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;map&#x60;/&#x60;zone&#x60;/&#x60;rssizone&#x60;, the scope id | 
 **scope** | [**ZoneScope**](ZoneScope.md) | scope | [default to &quot;site&quot;]
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

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


## CreateSiteZone

> Zone CreateSiteZone(ctx, siteId).Zone(zone).Execute()

createSiteZone



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
	zone := *openapiclient.NewZone() // Zone | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesZonesAPI.CreateSiteZone(context.Background(), siteId).Zone(zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.CreateSiteZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.CreateSiteZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | [**Zone**](Zone.md) | Request Body | 

### Return type

[**Zone**](Zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteZone

> DeleteSiteZone(ctx, siteId, zoneId).Execute()

deleteSiteZone



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
	zoneId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesZonesAPI.DeleteSiteZone(context.Background(), siteId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.DeleteSiteZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteZoneRequest struct via the builder pattern


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


## GetSiteZone

> Zone GetSiteZone(ctx, siteId, zoneId).Execute()

getSiteZone



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
	zoneId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesZonesAPI.GetSiteZone(context.Background(), siteId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.GetSiteZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.GetSiteZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Zone**](Zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteZoneStats

> ZoneStatsDetails GetSiteZoneStats(ctx, siteId, zoneType, zoneId).Execute()

getSiteZoneStats



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
	zoneType := "zoneType_example" // string | 
	zoneId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesZonesAPI.GetSiteZoneStats(context.Background(), siteId, zoneType, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.GetSiteZoneStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteZoneStats`: ZoneStatsDetails
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.GetSiteZoneStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**zoneType** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteZoneStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ZoneStatsDetails**](ZoneStatsDetails.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteZones

> []Zone ListSiteZones(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteZones



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
	resp, r, err := apiClient.SitesZonesAPI.ListSiteZones(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.ListSiteZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteZones`: []Zone
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.ListSiteZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Zone**](Zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteZonesStats

> []ZoneStats ListSiteZonesStats(ctx, siteId).MapId(mapId).Execute()

listSiteZonesStats



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
	mapId := "00000000-0000-0000-0000-000000000000" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesZonesAPI.ListSiteZonesStats(context.Background(), siteId).MapId(mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.ListSiteZonesStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteZonesStats`: []ZoneStats
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.ListSiteZonesStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteZonesStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mapId** | **string** |  | 

### Return type

[**[]ZoneStats**](ZoneStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteZoneSessions

> ResponseZoneSearch SearchSiteZoneSessions(ctx, siteId, zoneType).UserType(userType).User(user).ScopeId(scopeId).Scope(scope).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteZoneSessions



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
	zoneType := "zoneType_example" // string | 
	userType := openapiclient.rf_client_type("") // RfClientType | user type, client (default) / sdkclient / asset (optional)
	user := "user_example" // string | client MAC / Asset MAC / SDK UUID (optional)
	scopeId := "scopeId_example" // string | if `scope`==`map`/`zone`/`rssizone`, the scope id (optional)
	scope := openapiclient.visits_scope("") // VisitsScope | scope (optional) (default to "site")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesZonesAPI.SearchSiteZoneSessions(context.Background(), siteId, zoneType).UserType(userType).User(user).ScopeId(scopeId).Scope(scope).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.SearchSiteZoneSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteZoneSessions`: ResponseZoneSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.SearchSiteZoneSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**zoneType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteZoneSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userType** | [**RfClientType**](RfClientType.md) | user type, client (default) / sdkclient / asset | 
 **user** | **string** | client MAC / Asset MAC / SDK UUID | 
 **scopeId** | **string** | if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;map&#x60;/&#x60;zone&#x60;/&#x60;rssizone&#x60;, the scope id | 
 **scope** | [**VisitsScope**](VisitsScope.md) | scope | [default to &quot;site&quot;]
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseZoneSearch**](ResponseZoneSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteZone

> Zone UpdateSiteZone(ctx, siteId, zoneId).Zone(zone).Execute()

updateSiteZone



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
	zoneId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	zone := *openapiclient.NewZone() // Zone | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesZonesAPI.UpdateSiteZone(context.Background(), siteId, zoneId).Zone(zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesZonesAPI.UpdateSiteZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `SitesZonesAPI.UpdateSiteZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zone** | [**Zone**](Zone.md) | Request Body | 

### Return type

[**Zone**](Zone.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

