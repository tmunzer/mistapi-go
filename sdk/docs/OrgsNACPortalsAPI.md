# \OrgsNACPortalsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNacPortal**](OrgsNACPortalsAPI.md#CreateOrgNacPortal) | **Post** /api/v1/orgs/{org_id}/nacportals | createOrgNacPortal
[**DeleteOrgNacPortal**](OrgsNACPortalsAPI.md#DeleteOrgNacPortal) | **Delete** /api/v1/orgs/{org_id}/nacportals/{nacportal_id} | deleteOrgNacPortal
[**DownloadOrgNacPortalSsoSamlMetadata**](OrgsNACPortalsAPI.md#DownloadOrgNacPortalSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/saml_metadata.xml | downloadOrgNacPortalSsoSamlMetadata
[**GetOrgNacPortal**](OrgsNACPortalsAPI.md#GetOrgNacPortal) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id} | getOrgNacPortal
[**GetOrgNacPortalSsoSamlMetadata**](OrgsNACPortalsAPI.md#GetOrgNacPortalSsoSamlMetadata) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/saml_metadata | getOrgNacPortalSsoSamlMetadata
[**ListOrgNacPortalSsoLatestFailures**](OrgsNACPortalsAPI.md#ListOrgNacPortalSsoLatestFailures) | **Get** /api/v1/orgs/{org_id}/nacportals/{nacportal_id}/failures | listOrgNacPortalSsoLatestFailures
[**ListOrgNacPortals**](OrgsNACPortalsAPI.md#ListOrgNacPortals) | **Get** /api/v1/orgs/{org_id}/nacportals | listOrgNacPortals
[**UpdateOrgNacPortal**](OrgsNACPortalsAPI.md#UpdateOrgNacPortal) | **Put** /api/v1/orgs/{org_id}/nacportals/{nacportal_id} | updateOrgNacPortal



## CreateOrgNacPortal

> NacPortal CreateOrgNacPortal(ctx, orgId).NacPortal(nacPortal).Execute()

createOrgNacPortal



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
	nacPortal := *openapiclient.NewNacPortal() // NacPortal |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACPortalsAPI.CreateOrgNacPortal(context.Background(), orgId).NacPortal(nacPortal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.CreateOrgNacPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgNacPortal`: NacPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACPortalsAPI.CreateOrgNacPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgNacPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nacPortal** | [**NacPortal**](NacPortal.md) |  | 

### Return type

[**NacPortal**](NacPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgNacPortal

> DeleteOrgNacPortal(ctx, orgId, nacportalId).Execute()

deleteOrgNacPortal



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
	nacportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsNACPortalsAPI.DeleteOrgNacPortal(context.Background(), orgId, nacportalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.DeleteOrgNacPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgNacPortalRequest struct via the builder pattern


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


## DownloadOrgNacPortalSsoSamlMetadata

> *os.File DownloadOrgNacPortalSsoSamlMetadata(ctx, orgId, nacportalId).Execute()

downloadOrgNacPortalSsoSamlMetadata



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
	nacportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACPortalsAPI.DownloadOrgNacPortalSsoSamlMetadata(context.Background(), orgId, nacportalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.DownloadOrgNacPortalSsoSamlMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadOrgNacPortalSsoSamlMetadata`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACPortalsAPI.DownloadOrgNacPortalSsoSamlMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOrgNacPortalSsoSamlMetadataRequest struct via the builder pattern


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


## GetOrgNacPortal

> NacPortal GetOrgNacPortal(ctx, orgId, nacportalId).Execute()

getOrgNacPortal



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
	nacportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACPortalsAPI.GetOrgNacPortal(context.Background(), orgId, nacportalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.GetOrgNacPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgNacPortal`: NacPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACPortalsAPI.GetOrgNacPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNacPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NacPortal**](NacPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgNacPortalSsoSamlMetadata

> SsoSamlMetadata GetOrgNacPortalSsoSamlMetadata(ctx, orgId, nacportalId).Execute()

getOrgNacPortalSsoSamlMetadata



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
	nacportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACPortalsAPI.GetOrgNacPortalSsoSamlMetadata(context.Background(), orgId, nacportalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.GetOrgNacPortalSsoSamlMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgNacPortalSsoSamlMetadata`: SsoSamlMetadata
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACPortalsAPI.GetOrgNacPortalSsoSamlMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNacPortalSsoSamlMetadataRequest struct via the builder pattern


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


## ListOrgNacPortalSsoLatestFailures

> ResponseSsoFailureSearch ListOrgNacPortalSsoLatestFailures(ctx, orgId, nacportalId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listOrgNacPortalSsoLatestFailures



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
	nacportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACPortalsAPI.ListOrgNacPortalSsoLatestFailures(context.Background(), orgId, nacportalId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.ListOrgNacPortalSsoLatestFailures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgNacPortalSsoLatestFailures`: ResponseSsoFailureSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACPortalsAPI.ListOrgNacPortalSsoLatestFailures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNacPortalSsoLatestFailuresRequest struct via the builder pattern


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


## ListOrgNacPortals

> []NacPortal ListOrgNacPortals(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgNacPortals



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACPortalsAPI.ListOrgNacPortals(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.ListOrgNacPortals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgNacPortals`: []NacPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACPortalsAPI.ListOrgNacPortals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNacPortalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]NacPortal**](NacPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgNacPortal

> NacPortal UpdateOrgNacPortal(ctx, orgId, nacportalId).NacPortal(nacPortal).Execute()

updateOrgNacPortal



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
	nacportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nacPortal := *openapiclient.NewNacPortal() // NacPortal |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACPortalsAPI.UpdateOrgNacPortal(context.Background(), orgId, nacportalId).NacPortal(nacPortal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACPortalsAPI.UpdateOrgNacPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgNacPortal`: NacPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACPortalsAPI.UpdateOrgNacPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgNacPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nacPortal** | [**NacPortal**](NacPortal.md) |  | 

### Return type

[**NacPortal**](NacPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

