# \OrgsAPITokensAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgApiToken**](OrgsAPITokensAPI.md#CreateOrgApiToken) | **Post** /api/v1/orgs/{org_id}/apitokens | createOrgApiToken
[**DeleteOrgApiToken**](OrgsAPITokensAPI.md#DeleteOrgApiToken) | **Delete** /api/v1/orgs/{org_id}/apitokens/{apitoken_id} | deleteOrgApiToken
[**GetOrgApiToken**](OrgsAPITokensAPI.md#GetOrgApiToken) | **Get** /api/v1/orgs/{org_id}/apitokens/{apitoken_id} | getOrgApiToken
[**ListOrgApiTokens**](OrgsAPITokensAPI.md#ListOrgApiTokens) | **Get** /api/v1/orgs/{org_id}/apitokens | listOrgApiTokens
[**UpdateOrgApiToken**](OrgsAPITokensAPI.md#UpdateOrgApiToken) | **Put** /api/v1/orgs/{org_id}/apitokens/{apitoken_id} | updateOrgApiToken



## CreateOrgApiToken

> OrgApitoken CreateOrgApiToken(ctx, orgId).OrgApitoken(orgApitoken).Execute()

createOrgApiToken



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
	orgApitoken := *openapiclient.NewOrgApitoken() // OrgApitoken |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPITokensAPI.CreateOrgApiToken(context.Background(), orgId).OrgApitoken(orgApitoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPITokensAPI.CreateOrgApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgApiToken`: OrgApitoken
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPITokensAPI.CreateOrgApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgApitoken** | [**OrgApitoken**](OrgApitoken.md) |  | 

### Return type

[**OrgApitoken**](OrgApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgApiToken

> DeleteOrgApiToken(ctx, orgId, apitokenId).Execute()

deleteOrgApiToken



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
	apitokenId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAPITokensAPI.DeleteOrgApiToken(context.Background(), orgId, apitokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPITokensAPI.DeleteOrgApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**apitokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgApiTokenRequest struct via the builder pattern


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


## GetOrgApiToken

> OrgApitoken GetOrgApiToken(ctx, orgId, apitokenId).Execute()

getOrgApiToken



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
	apitokenId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPITokensAPI.GetOrgApiToken(context.Background(), orgId, apitokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPITokensAPI.GetOrgApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgApiToken`: OrgApitoken
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPITokensAPI.GetOrgApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**apitokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgApitoken**](OrgApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgApiTokens

> []OrgApitoken ListOrgApiTokens(ctx, orgId).Execute()

listOrgApiTokens



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPITokensAPI.ListOrgApiTokens(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPITokensAPI.ListOrgApiTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgApiTokens`: []OrgApitoken
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPITokensAPI.ListOrgApiTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgApiTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrgApitoken**](OrgApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgApiToken

> OrgApitoken UpdateOrgApiToken(ctx, orgId, apitokenId).OrgApitoken(orgApitoken).Execute()

updateOrgApiToken



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
	apitokenId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	orgApitoken := *openapiclient.NewOrgApitoken() // OrgApitoken | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPITokensAPI.UpdateOrgApiToken(context.Background(), orgId, apitokenId).OrgApitoken(orgApitoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPITokensAPI.UpdateOrgApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgApiToken`: OrgApitoken
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPITokensAPI.UpdateOrgApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**apitokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orgApitoken** | [**OrgApitoken**](OrgApitoken.md) | Request Body | 

### Return type

[**OrgApitoken**](OrgApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

