# \MSPsSSOAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspSso**](MSPsSSOAPI.md#CreateMspSso) | **Post** /api/v1/msps/{msp_id}/ssos | createMspSso
[**DeleteMspSso**](MSPsSSOAPI.md#DeleteMspSso) | **Delete** /api/v1/msps/{msp_id}/ssos/{sso_id} | deleteMspSso
[**DownloadMspSsoSamlMetadata**](MSPsSSOAPI.md#DownloadMspSsoSamlMetadata) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id}/metadata.xml | downloadMspSsoSamlMetadata
[**GetMspSso**](MSPsSSOAPI.md#GetMspSso) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id} | getMspSso
[**GetMspSsoSamlMetadata**](MSPsSSOAPI.md#GetMspSsoSamlMetadata) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id}/metadata | getMspSsoSamlMetadata
[**ListMspSsoLatestFailures**](MSPsSSOAPI.md#ListMspSsoLatestFailures) | **Get** /api/v1/msps/{msp_id}/ssos/{sso_id}/failures | listMspSsoLatestFailures
[**ListMspSsos**](MSPsSSOAPI.md#ListMspSsos) | **Get** /api/v1/msps/{msp_id}/ssos | listMspSsos
[**UpdateMspSso**](MSPsSSOAPI.md#UpdateMspSso) | **Put** /api/v1/msps/{msp_id}/ssos/{sso_id} | updateMspSso



## CreateMspSso

> Sso CreateMspSso(ctx, mspId).Sso(sso).Execute()

createMspSso



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sso := *openapiclient.NewSso("Name_example") // Sso | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSOAPI.CreateMspSso(context.Background(), mspId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.CreateMspSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMspSso`: Sso
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSOAPI.CreateMspSso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMspSsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sso** | [**Sso**](Sso.md) | Request Body | 

### Return type

[**Sso**](Sso.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMspSso

> DeleteMspSso(ctx, mspId, ssoId).Execute()

deleteMspSso



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsSSOAPI.DeleteMspSso(context.Background(), mspId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.DeleteMspSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMspSsoRequest struct via the builder pattern


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


## DownloadMspSsoSamlMetadata

> *os.File DownloadMspSsoSamlMetadata(ctx, mspId, ssoId).Execute()

downloadMspSsoSamlMetadata



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSOAPI.DownloadMspSsoSamlMetadata(context.Background(), mspId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.DownloadMspSsoSamlMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadMspSsoSamlMetadata`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSOAPI.DownloadMspSsoSamlMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMspSsoSamlMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspSso

> Sso GetMspSso(ctx, mspId, ssoId).Execute()

getMspSso



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSOAPI.GetMspSso(context.Background(), mspId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.GetMspSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspSso`: Sso
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSOAPI.GetMspSso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspSsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sso**](Sso.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspSsoSamlMetadata

> SsoSamlMetadata GetMspSsoSamlMetadata(ctx, mspId, ssoId).Execute()

getMspSsoSamlMetadata



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSOAPI.GetMspSsoSamlMetadata(context.Background(), mspId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.GetMspSsoSamlMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspSsoSamlMetadata`: SsoSamlMetadata
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSOAPI.GetMspSsoSamlMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspSsoSamlMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SsoSamlMetadata**](SsoSamlMetadata.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMspSsoLatestFailures

> ResponseSsoFailureSearch ListMspSsoLatestFailures(ctx, mspId, ssoId).Execute()

listMspSsoLatestFailures



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSOAPI.ListMspSsoLatestFailures(context.Background(), mspId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.ListMspSsoLatestFailures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspSsoLatestFailures`: ResponseSsoFailureSearch
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSOAPI.ListMspSsoLatestFailures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspSsoLatestFailuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseSsoFailureSearch**](ResponseSsoFailureSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMspSsos

> []Sso ListMspSsos(ctx, mspId).Execute()

listMspSsos



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSOAPI.ListMspSsos(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.ListMspSsos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspSsos`: []Sso
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSOAPI.ListMspSsos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspSsosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Sso**](Sso.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMspSso

> Sso UpdateMspSso(ctx, mspId, ssoId).Sso(sso).Execute()

updateMspSso



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sso := *openapiclient.NewSso("Name_example") // Sso | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSOAPI.UpdateMspSso(context.Background(), mspId, ssoId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSOAPI.UpdateMspSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMspSso`: Sso
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSOAPI.UpdateMspSso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMspSsoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sso** | [**Sso**](Sso.md) | Request Body | 

### Return type

[**Sso**](Sso.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

