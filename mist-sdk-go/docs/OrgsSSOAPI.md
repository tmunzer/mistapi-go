# \OrgsSSOAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSso**](OrgsSSOAPI.md#CreateOrgSso) | **Post** /api/v1/orgs/{org_id}/ssos | createOrgSso
[**DeleteOrgSso**](OrgsSSOAPI.md#DeleteOrgSso) | **Delete** /api/v1/orgs/{org_id}/ssos/{sso_id} | deleteOrgSso
[**DownloadOrgSsoSamlMetadata**](OrgsSSOAPI.md#DownloadOrgSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id}/metadata.xml | downloadOrgSsoSamlMetadata
[**GetOrgSso**](OrgsSSOAPI.md#GetOrgSso) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id} | getOrgSso
[**GetOrgSsoSamlMetadata**](OrgsSSOAPI.md#GetOrgSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id}/metadata | getOrgSsoSamlMetadata
[**ListOrgSsoLatestFailures**](OrgsSSOAPI.md#ListOrgSsoLatestFailures) | **Get** /api/v1/orgs/{org_id}/ssos/{sso_id}/failures | listOrgSsoLatestFailures
[**ListOrgSsos**](OrgsSSOAPI.md#ListOrgSsos) | **Get** /api/v1/orgs/{org_id}/ssos | listOrgSsos
[**UpdateOrgSso**](OrgsSSOAPI.md#UpdateOrgSso) | **Put** /api/v1/orgs/{org_id}/ssos/{sso_id} | updateOrgSso



## CreateOrgSso

> Sso CreateOrgSso(ctx, orgId).Sso(sso).Execute()

createOrgSso



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sso := *openapiclient.NewSso("Name_example") // Sso | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSOAPI.CreateOrgSso(context.Background(), orgId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.CreateOrgSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSso`: Sso
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSOAPI.CreateOrgSso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSsoRequest struct via the builder pattern


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


## DeleteOrgSso

> DeleteOrgSso(ctx, orgId, ssoId).Execute()

deleteOrgSso



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSSOAPI.DeleteOrgSso(context.Background(), orgId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.DeleteOrgSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSsoRequest struct via the builder pattern


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


## DownloadOrgSsoSamlMetadata

> *os.File DownloadOrgSsoSamlMetadata(ctx, orgId, ssoId).Execute()

downloadOrgSsoSamlMetadata



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSOAPI.DownloadOrgSsoSamlMetadata(context.Background(), orgId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.DownloadOrgSsoSamlMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadOrgSsoSamlMetadata`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSOAPI.DownloadOrgSsoSamlMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOrgSsoSamlMetadataRequest struct via the builder pattern


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


## GetOrgSso

> Sso GetOrgSso(ctx, orgId, ssoId).Execute()

getOrgSso



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSOAPI.GetOrgSso(context.Background(), orgId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.GetOrgSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSso`: Sso
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSOAPI.GetOrgSso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSsoRequest struct via the builder pattern


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


## GetOrgSsoSamlMetadata

> SsoSamlMetadata GetOrgSsoSamlMetadata(ctx, orgId, ssoId).Execute()

getOrgSsoSamlMetadata



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSOAPI.GetOrgSsoSamlMetadata(context.Background(), orgId, ssoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.GetOrgSsoSamlMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSsoSamlMetadata`: SsoSamlMetadata
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSOAPI.GetOrgSsoSamlMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSsoSamlMetadataRequest struct via the builder pattern


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


## ListOrgSsoLatestFailures

> ResponseSsoFailureSearch ListOrgSsoLatestFailures(ctx, orgId, ssoId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listOrgSsoLatestFailures



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSOAPI.ListOrgSsoLatestFailures(context.Background(), orgId, ssoId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.ListOrgSsoLatestFailures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSsoLatestFailures`: ResponseSsoFailureSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSOAPI.ListOrgSsoLatestFailures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSsoLatestFailuresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

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


## ListOrgSsos

> []Sso ListOrgSsos(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgSsos



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSOAPI.ListOrgSsos(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.ListOrgSsos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSsos`: []Sso
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSOAPI.ListOrgSsos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSsosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

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


## UpdateOrgSso

> Sso UpdateOrgSso(ctx, orgId, ssoId).Sso(sso).Execute()

updateOrgSso



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sso := *openapiclient.NewSso("Name_example") // Sso | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSOAPI.UpdateOrgSso(context.Background(), orgId, ssoId).Sso(sso).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSOAPI.UpdateOrgSso``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSso`: Sso
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSOAPI.UpdateOrgSso`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSsoRequest struct via the builder pattern


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

