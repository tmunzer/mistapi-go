# \SitesBeaconsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteBeacon**](SitesBeaconsAPI.md#CreateSiteBeacon) | **Post** /api/v1/sites/{site_id}/beacons | createSiteBeacon
[**DeleteSiteBeacons**](SitesBeaconsAPI.md#DeleteSiteBeacons) | **Delete** /api/v1/sites/{site_id}/beacons/{beacon_id} | deleteSiteBeacons
[**GetSiteBeacon**](SitesBeaconsAPI.md#GetSiteBeacon) | **Get** /api/v1/sites/{site_id}/beacons/{beacon_id} | getSiteBeacon
[**ListSiteBeacons**](SitesBeaconsAPI.md#ListSiteBeacons) | **Get** /api/v1/sites/{site_id}/beacons | listSiteBeacons
[**ListSiteBeaconsStats**](SitesBeaconsAPI.md#ListSiteBeaconsStats) | **Get** /api/v1/sites/{site_id}/stats/beacons | listSiteBeaconsStats
[**UpdateSiteBeacons**](SitesBeaconsAPI.md#UpdateSiteBeacons) | **Put** /api/v1/sites/{site_id}/beacons/{beacon_id} | updateSiteBeacons



## CreateSiteBeacon

> Beacon CreateSiteBeacon(ctx, siteId).Beacon(beacon).Execute()

createSiteBeacon



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
	beacon := *openapiclient.NewBeacon() // Beacon | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesBeaconsAPI.CreateSiteBeacon(context.Background(), siteId).Beacon(beacon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesBeaconsAPI.CreateSiteBeacon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteBeacon`: Beacon
	fmt.Fprintf(os.Stdout, "Response from `SitesBeaconsAPI.CreateSiteBeacon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteBeaconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beacon** | [**Beacon**](Beacon.md) | Request Body | 

### Return type

[**Beacon**](Beacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteBeacons

> DeleteSiteBeacons(ctx, siteId, beaconId).Execute()

deleteSiteBeacons



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
	beaconId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesBeaconsAPI.DeleteSiteBeacons(context.Background(), siteId, beaconId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesBeaconsAPI.DeleteSiteBeacons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**beaconId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteBeaconsRequest struct via the builder pattern


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


## GetSiteBeacon

> Beacon GetSiteBeacon(ctx, siteId, beaconId).Execute()

getSiteBeacon



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
	beaconId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesBeaconsAPI.GetSiteBeacon(context.Background(), siteId, beaconId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesBeaconsAPI.GetSiteBeacon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteBeacon`: Beacon
	fmt.Fprintf(os.Stdout, "Response from `SitesBeaconsAPI.GetSiteBeacon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**beaconId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteBeaconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Beacon**](Beacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteBeacons

> []Beacon ListSiteBeacons(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteBeacons



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
	resp, r, err := apiClient.SitesBeaconsAPI.ListSiteBeacons(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesBeaconsAPI.ListSiteBeacons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteBeacons`: []Beacon
	fmt.Fprintf(os.Stdout, "Response from `SitesBeaconsAPI.ListSiteBeacons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteBeaconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Beacon**](Beacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteBeaconsStats

> []BeaconStatsItems ListSiteBeaconsStats(ctx, siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listSiteBeaconsStats



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
	resp, r, err := apiClient.SitesBeaconsAPI.ListSiteBeaconsStats(context.Background(), siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesBeaconsAPI.ListSiteBeaconsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteBeaconsStats`: []BeaconStatsItems
	fmt.Fprintf(os.Stdout, "Response from `SitesBeaconsAPI.ListSiteBeaconsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteBeaconsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]BeaconStatsItems**](BeaconStatsItems.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteBeacons

> Beacon UpdateSiteBeacons(ctx, siteId, beaconId).Beacon(beacon).Execute()

updateSiteBeacons



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
	beaconId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	beacon := *openapiclient.NewBeacon() // Beacon | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesBeaconsAPI.UpdateSiteBeacons(context.Background(), siteId, beaconId).Beacon(beacon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesBeaconsAPI.UpdateSiteBeacons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteBeacons`: Beacon
	fmt.Fprintf(os.Stdout, "Response from `SitesBeaconsAPI.UpdateSiteBeacons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**beaconId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteBeaconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **beacon** | [**Beacon**](Beacon.md) | Request Body | 

### Return type

[**Beacon**](Beacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

