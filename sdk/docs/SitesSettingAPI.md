# \SitesSettingAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWatchedStations**](SitesSettingAPI.md#CreateSiteWatchedStations) | **Post** /api/v1/sites/{site_id}/setting/watched_station | createSiteWatchedStations
[**CreateSiteWirelessClientsAllowlist**](SitesSettingAPI.md#CreateSiteWirelessClientsAllowlist) | **Post** /api/v1/sites/{site_id}/setting/whitelist | createSiteWirelessClientsAllowlist
[**CreateSiteWirelessClientsBlocklist**](SitesSettingAPI.md#CreateSiteWirelessClientsBlocklist) | **Post** /api/v1/sites/{site_id}/setting/blacklist | createSiteWirelessClientsBlocklist
[**DeleteSiteWatchedStations**](SitesSettingAPI.md#DeleteSiteWatchedStations) | **Delete** /api/v1/sites/{site_id}/setting/watched_station | deleteSiteWatchedStations
[**DeleteSiteWirelessClientsAllowlist**](SitesSettingAPI.md#DeleteSiteWirelessClientsAllowlist) | **Delete** /api/v1/sites/{site_id}/setting/whitelist | deleteSiteWirelessClientsAllowlist
[**DeleteSiteWirelessClientsBlocklist**](SitesSettingAPI.md#DeleteSiteWirelessClientsBlocklist) | **Delete** /api/v1/sites/{site_id}/setting/blacklist | deleteSiteWirelessClientsBlocklist
[**GetSiteSetting**](SitesSettingAPI.md#GetSiteSetting) | **Get** /api/v1/sites/{site_id}/setting | getSiteSetting
[**UpdateSiteSettings**](SitesSettingAPI.md#UpdateSiteSettings) | **Put** /api/v1/sites/{site_id}/setting | updateSiteSettings



## CreateSiteWatchedStations

> MacAddresses CreateSiteWatchedStations(ctx, siteId).MacAddresses(macAddresses).Execute()

createSiteWatchedStations



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
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSettingAPI.CreateSiteWatchedStations(context.Background(), siteId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.CreateSiteWatchedStations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWatchedStations`: MacAddresses
	fmt.Fprintf(os.Stdout, "Response from `SitesSettingAPI.CreateSiteWatchedStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWatchedStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

### Return type

[**MacAddresses**](MacAddresses.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteWirelessClientsAllowlist

> MacAddresses CreateSiteWirelessClientsAllowlist(ctx, siteId).MacAddresses(macAddresses).Execute()

createSiteWirelessClientsAllowlist



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
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSettingAPI.CreateSiteWirelessClientsAllowlist(context.Background(), siteId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.CreateSiteWirelessClientsAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWirelessClientsAllowlist`: MacAddresses
	fmt.Fprintf(os.Stdout, "Response from `SitesSettingAPI.CreateSiteWirelessClientsAllowlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWirelessClientsAllowlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

### Return type

[**MacAddresses**](MacAddresses.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSiteWirelessClientsBlocklist

> MacAddresses CreateSiteWirelessClientsBlocklist(ctx, siteId).MacAddresses(macAddresses).Execute()

createSiteWirelessClientsBlocklist



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
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSettingAPI.CreateSiteWirelessClientsBlocklist(context.Background(), siteId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.CreateSiteWirelessClientsBlocklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWirelessClientsBlocklist`: MacAddresses
	fmt.Fprintf(os.Stdout, "Response from `SitesSettingAPI.CreateSiteWirelessClientsBlocklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWirelessClientsBlocklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

### Return type

[**MacAddresses**](MacAddresses.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteWatchedStations

> DeleteSiteWatchedStations(ctx, siteId).Execute()

deleteSiteWatchedStations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesSettingAPI.DeleteSiteWatchedStations(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.DeleteSiteWatchedStations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteWatchedStationsRequest struct via the builder pattern


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


## DeleteSiteWirelessClientsAllowlist

> DeleteSiteWirelessClientsAllowlist(ctx, siteId).Execute()

deleteSiteWirelessClientsAllowlist



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesSettingAPI.DeleteSiteWirelessClientsAllowlist(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.DeleteSiteWirelessClientsAllowlist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteWirelessClientsAllowlistRequest struct via the builder pattern


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


## DeleteSiteWirelessClientsBlocklist

> DeleteSiteWirelessClientsBlocklist(ctx, siteId).Execute()

deleteSiteWirelessClientsBlocklist



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesSettingAPI.DeleteSiteWirelessClientsBlocklist(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.DeleteSiteWirelessClientsBlocklist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSiteWirelessClientsBlocklistRequest struct via the builder pattern


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


## GetSiteSetting

> SiteSetting GetSiteSetting(ctx, siteId).Execute()

getSiteSetting



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSettingAPI.GetSiteSetting(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.GetSiteSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSetting`: SiteSetting
	fmt.Fprintf(os.Stdout, "Response from `SitesSettingAPI.GetSiteSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SiteSetting**](SiteSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteSettings

> SiteSetting UpdateSiteSettings(ctx, siteId).SiteSetting(siteSetting).Execute()

updateSiteSettings



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
	siteSetting := *openapiclient.NewSiteSetting() // SiteSetting | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSettingAPI.UpdateSiteSettings(context.Background(), siteId).SiteSetting(siteSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSettingAPI.UpdateSiteSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteSettings`: SiteSetting
	fmt.Fprintf(os.Stdout, "Response from `SitesSettingAPI.UpdateSiteSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteSetting** | [**SiteSetting**](SiteSetting.md) | Request Body | 

### Return type

[**SiteSetting**](SiteSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

