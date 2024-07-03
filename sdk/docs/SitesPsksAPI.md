# \SitesPsksAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSitePsk**](SitesPsksAPI.md#CreateSitePsk) | **Post** /api/v1/sites/{site_id}/psks | createSitePsk
[**DeleteSitePsk**](SitesPsksAPI.md#DeleteSitePsk) | **Delete** /api/v1/sites/{site_id}/psks/{psk_id} | deleteSitePsk
[**GetSitePsk**](SitesPsksAPI.md#GetSitePsk) | **Get** /api/v1/sites/{site_id}/psks/{psk_id} | getSitePsk
[**ImportSitePsks**](SitesPsksAPI.md#ImportSitePsks) | **Post** /api/v1/sites/{site_id}/psks/import | importSitePsks
[**ListSitePsks**](SitesPsksAPI.md#ListSitePsks) | **Get** /api/v1/sites/{site_id}/psks | listSitePsks
[**UpdateSiteMultiplePsks**](SitesPsksAPI.md#UpdateSiteMultiplePsks) | **Put** /api/v1/sites/{site_id}/psks | updateSiteMultiplePsks
[**UpdateSitePsk**](SitesPsksAPI.md#UpdateSitePsk) | **Put** /api/v1/sites/{site_id}/psks/{psk_id} | updateSitePsk



## CreateSitePsk

> Psk CreateSitePsk(ctx, siteId).Psk(psk).Execute()

createSitePsk



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
	psk := *openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example") // Psk | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesPsksAPI.CreateSitePsk(context.Background(), siteId).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesPsksAPI.CreateSitePsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSitePsk`: Psk
	fmt.Fprintf(os.Stdout, "Response from `SitesPsksAPI.CreateSitePsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSitePskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **psk** | [**Psk**](Psk.md) | Request Body | 

### Return type

[**Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSitePsk

> DeleteSitePsk(ctx, siteId, pskId).Execute()

deleteSitePsk



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
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesPsksAPI.DeleteSitePsk(context.Background(), siteId, pskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesPsksAPI.DeleteSitePsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**pskId** | **string** | PSK ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSitePskRequest struct via the builder pattern


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


## GetSitePsk

> Psk GetSitePsk(ctx, siteId, pskId).Execute()

getSitePsk



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
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesPsksAPI.GetSitePsk(context.Background(), siteId, pskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesPsksAPI.GetSitePsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSitePsk`: Psk
	fmt.Fprintf(os.Stdout, "Response from `SitesPsksAPI.GetSitePsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**pskId** | **string** | PSK ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitePskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSitePsks

> []Psk ImportSitePsks(ctx, siteId).Psk(psk).Execute()

importSitePsks



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
	psk := []openapiclient.Psk{*openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example")} // []Psk |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesPsksAPI.ImportSitePsks(context.Background(), siteId).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesPsksAPI.ImportSitePsks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportSitePsks`: []Psk
	fmt.Fprintf(os.Stdout, "Response from `SitesPsksAPI.ImportSitePsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSitePsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **psk** | [**[]Psk**](Psk.md) |  | 

### Return type

[**[]Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSitePsks

> []Psk ListSitePsks(ctx, siteId).Ssid(ssid).Role(role).Name(name).Page(page).Limit(limit).Execute()

listSitePsks



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
	ssid := "ssid_example" // string |  (optional)
	role := "role_example" // string |  (optional)
	name := "name_example" // string |  (optional) (default to "")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesPsksAPI.ListSitePsks(context.Background(), siteId).Ssid(ssid).Role(role).Name(name).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesPsksAPI.ListSitePsks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSitePsks`: []Psk
	fmt.Fprintf(os.Stdout, "Response from `SitesPsksAPI.ListSitePsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSitePsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssid** | **string** |  | 
 **role** | **string** |  | 
 **name** | **string** |  | [default to &quot;&quot;]
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteMultiplePsks

> []Psk UpdateSiteMultiplePsks(ctx, siteId).Psk(psk).Execute()

updateSiteMultiplePsks



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
	psk := []openapiclient.Psk{*openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example")} // []Psk |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesPsksAPI.UpdateSiteMultiplePsks(context.Background(), siteId).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesPsksAPI.UpdateSiteMultiplePsks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteMultiplePsks`: []Psk
	fmt.Fprintf(os.Stdout, "Response from `SitesPsksAPI.UpdateSiteMultiplePsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteMultiplePsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **psk** | [**[]Psk**](Psk.md) |  | 

### Return type

[**[]Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSitePsk

> Psk UpdateSitePsk(ctx, siteId, pskId).Psk(psk).Execute()

updateSitePsk



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
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID
	psk := *openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example") // Psk | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesPsksAPI.UpdateSitePsk(context.Background(), siteId, pskId).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesPsksAPI.UpdateSitePsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSitePsk`: Psk
	fmt.Fprintf(os.Stdout, "Response from `SitesPsksAPI.UpdateSitePsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**pskId** | **string** | PSK ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSitePskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **psk** | [**Psk**](Psk.md) | Request Body | 

### Return type

[**Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

