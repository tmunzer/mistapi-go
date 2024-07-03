# \SitesUISettingsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteUiSettings**](SitesUISettingsAPI.md#CreateSiteUiSettings) | **Post** /api/v1/sites/{site_id}/uisettings | createSiteUiSettings
[**DeleteSiteUiSetting**](SitesUISettingsAPI.md#DeleteSiteUiSetting) | **Delete** /api/v1/sites/{site_id}/uisettings/{uisetting_id} | deleteSiteUiSetting
[**GetSiteUiSetting**](SitesUISettingsAPI.md#GetSiteUiSetting) | **Get** /api/v1/sites/{site_id}/uisettings/{uisetting_id} | getSiteUiSetting
[**GetSiteUiSettingDerived**](SitesUISettingsAPI.md#GetSiteUiSettingDerived) | **Get** /api/v1/sites/{site_id}/uisettings/derived | getSiteUiSettingDerived
[**ListSiteUiSettings**](SitesUISettingsAPI.md#ListSiteUiSettings) | **Get** /api/v1/sites/{site_id}/uisettings | listSiteUiSettings
[**UpdateSiteUiSetting**](SitesUISettingsAPI.md#UpdateSiteUiSetting) | **Post** /api/v1/sites/{site_id}/uisettings/{uisetting_id} | updateSiteUiSetting



## CreateSiteUiSettings

> UiSettings CreateSiteUiSettings(ctx, siteId).UiSettings(uiSettings).Execute()

createSiteUiSettings



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
	uiSettings := *openapiclient.NewUiSettings("Description of the databoard", "databoard") // UiSettings | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesUISettingsAPI.CreateSiteUiSettings(context.Background(), siteId).UiSettings(uiSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesUISettingsAPI.CreateSiteUiSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteUiSettings`: UiSettings
	fmt.Fprintf(os.Stdout, "Response from `SitesUISettingsAPI.CreateSiteUiSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteUiSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uiSettings** | [**UiSettings**](UiSettings.md) | Request Body | 

### Return type

[**UiSettings**](UiSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteUiSetting

> DeleteSiteUiSetting(ctx, siteId, uisettingId).Execute()

deleteSiteUiSetting



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
	uisettingId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesUISettingsAPI.DeleteSiteUiSetting(context.Background(), siteId, uisettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesUISettingsAPI.DeleteSiteUiSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**uisettingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteUiSettingRequest struct via the builder pattern


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


## GetSiteUiSetting

> []UiSettings GetSiteUiSetting(ctx, siteId, uisettingId).Execute()

getSiteUiSetting



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
	uisettingId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesUISettingsAPI.GetSiteUiSetting(context.Background(), siteId, uisettingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesUISettingsAPI.GetSiteUiSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteUiSetting`: []UiSettings
	fmt.Fprintf(os.Stdout, "Response from `SitesUISettingsAPI.GetSiteUiSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**uisettingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteUiSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]UiSettings**](UiSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteUiSettingDerived

> UiSettings GetSiteUiSettingDerived(ctx, siteId).Execute()

getSiteUiSettingDerived



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
	resp, r, err := apiClient.SitesUISettingsAPI.GetSiteUiSettingDerived(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesUISettingsAPI.GetSiteUiSettingDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteUiSettingDerived`: UiSettings
	fmt.Fprintf(os.Stdout, "Response from `SitesUISettingsAPI.GetSiteUiSettingDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteUiSettingDerivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UiSettings**](UiSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteUiSettings

> []UiSettings ListSiteUiSettings(ctx, siteId).Execute()

listSiteUiSettings



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
	resp, r, err := apiClient.SitesUISettingsAPI.ListSiteUiSettings(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesUISettingsAPI.ListSiteUiSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteUiSettings`: []UiSettings
	fmt.Fprintf(os.Stdout, "Response from `SitesUISettingsAPI.ListSiteUiSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteUiSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UiSettings**](UiSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteUiSetting

> UiSettings UpdateSiteUiSetting(ctx, siteId, uisettingId).UiSettings(uiSettings).Execute()

updateSiteUiSetting



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
	uisettingId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	uiSettings := *openapiclient.NewUiSettings("Description of the databoard", "databoard") // UiSettings | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesUISettingsAPI.UpdateSiteUiSetting(context.Background(), siteId, uisettingId).UiSettings(uiSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesUISettingsAPI.UpdateSiteUiSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteUiSetting`: UiSettings
	fmt.Fprintf(os.Stdout, "Response from `SitesUISettingsAPI.UpdateSiteUiSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**uisettingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteUiSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uiSettings** | [**UiSettings**](UiSettings.md) | Request Body | 

### Return type

[**UiSettings**](UiSettings.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

