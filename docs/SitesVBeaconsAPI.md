# \SitesVBeaconsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteVBeacon**](SitesVBeaconsAPI.md#CreateSiteVBeacon) | **Post** /api/v1/sites/{site_id}/vbeacons | createSiteVBeacon
[**DeleteSiteVBeacon**](SitesVBeaconsAPI.md#DeleteSiteVBeacon) | **Delete** /api/v1/sites/{site_id}/vbeacons/{vbeacon_id} | deleteSiteVBeacon
[**GetSiteVBeacon**](SitesVBeaconsAPI.md#GetSiteVBeacon) | **Get** /api/v1/sites/{site_id}/vbeacons/{vbeacon_id} | getSiteVBeacon
[**ListSiteVBeacons**](SitesVBeaconsAPI.md#ListSiteVBeacons) | **Get** /api/v1/sites/{site_id}/vbeacons | listSiteVBeacons
[**UpdateSiteVBeacon**](SitesVBeaconsAPI.md#UpdateSiteVBeacon) | **Put** /api/v1/sites/{site_id}/vbeacons/{vbeacon_id} | updateSiteVBeacon



## CreateSiteVBeacon

> Vbeacon CreateSiteVBeacon(ctx, siteId).Vbeacon(vbeacon).Execute()

createSiteVBeacon



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
	vbeacon := *openapiclient.NewVbeacon() // Vbeacon | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesVBeaconsAPI.CreateSiteVBeacon(context.Background(), siteId).Vbeacon(vbeacon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesVBeaconsAPI.CreateSiteVBeacon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteVBeacon`: Vbeacon
	fmt.Fprintf(os.Stdout, "Response from `SitesVBeaconsAPI.CreateSiteVBeacon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteVBeaconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vbeacon** | [**Vbeacon**](Vbeacon.md) | Request Body | 

### Return type

[**Vbeacon**](Vbeacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteVBeacon

> DeleteSiteVBeacon(ctx, siteId, vbeaconId).Execute()

deleteSiteVBeacon



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
	vbeaconId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesVBeaconsAPI.DeleteSiteVBeacon(context.Background(), siteId, vbeaconId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesVBeaconsAPI.DeleteSiteVBeacon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**vbeaconId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteVBeaconRequest struct via the builder pattern


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


## GetSiteVBeacon

> Vbeacon GetSiteVBeacon(ctx, siteId, vbeaconId).Execute()

getSiteVBeacon



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
	vbeaconId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesVBeaconsAPI.GetSiteVBeacon(context.Background(), siteId, vbeaconId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesVBeaconsAPI.GetSiteVBeacon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteVBeacon`: Vbeacon
	fmt.Fprintf(os.Stdout, "Response from `SitesVBeaconsAPI.GetSiteVBeacon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**vbeaconId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteVBeaconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Vbeacon**](Vbeacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteVBeacons

> []Vbeacon ListSiteVBeacons(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteVBeacons



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
	resp, r, err := apiClient.SitesVBeaconsAPI.ListSiteVBeacons(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesVBeaconsAPI.ListSiteVBeacons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteVBeacons`: []Vbeacon
	fmt.Fprintf(os.Stdout, "Response from `SitesVBeaconsAPI.ListSiteVBeacons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteVBeaconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Vbeacon**](Vbeacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteVBeacon

> Vbeacon UpdateSiteVBeacon(ctx, siteId, vbeaconId).Vbeacon(vbeacon).Execute()

updateSiteVBeacon



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
	vbeaconId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	vbeacon := *openapiclient.NewVbeacon() // Vbeacon | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesVBeaconsAPI.UpdateSiteVBeacon(context.Background(), siteId, vbeaconId).Vbeacon(vbeacon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesVBeaconsAPI.UpdateSiteVBeacon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteVBeacon`: Vbeacon
	fmt.Fprintf(os.Stdout, "Response from `SitesVBeaconsAPI.UpdateSiteVBeacon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**vbeaconId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteVBeaconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vbeacon** | [**Vbeacon**](Vbeacon.md) | Request Body | 

### Return type

[**Vbeacon**](Vbeacon.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

