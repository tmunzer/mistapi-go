# \SelfAPITokenAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiToken**](SelfAPITokenAPI.md#CreateApiToken) | **Post** /api/v1/self/apitokens | createApiToken
[**DeleteApiToken**](SelfAPITokenAPI.md#DeleteApiToken) | **Delete** /api/v1/self/apitokens/{apitoken_id} | deleteApiToken
[**GetApiToken**](SelfAPITokenAPI.md#GetApiToken) | **Get** /api/v1/self/apitokens/{apitoken_id} | getApiToken
[**ListApiTokens**](SelfAPITokenAPI.md#ListApiTokens) | **Get** /api/v1/self/apitokens | listApiTokens
[**UpdateApiToken**](SelfAPITokenAPI.md#UpdateApiToken) | **Put** /api/v1/self/apitokens/{apitoken_id} | updateApiToken



## CreateApiToken

> []UserApitoken CreateApiToken(ctx).UserApitoken(userApitoken).Execute()

createApiToken



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
	userApitoken := *openapiclient.NewUserApitoken() // UserApitoken |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAPITokenAPI.CreateApiToken(context.Background()).UserApitoken(userApitoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPITokenAPI.CreateApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApiToken`: []UserApitoken
	fmt.Fprintf(os.Stdout, "Response from `SelfAPITokenAPI.CreateApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userApitoken** | [**UserApitoken**](UserApitoken.md) |  | 

### Return type

[**[]UserApitoken**](UserApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiToken

> DeleteApiToken(ctx, apitokenId).Execute()

deleteApiToken



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
	apitokenId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SelfAPITokenAPI.DeleteApiToken(context.Background(), apitokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPITokenAPI.DeleteApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apitokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiTokenRequest struct via the builder pattern


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


## GetApiToken

> UserApitoken GetApiToken(ctx, apitokenId).Execute()

getApiToken



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
	apitokenId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAPITokenAPI.GetApiToken(context.Background(), apitokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPITokenAPI.GetApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiToken`: UserApitoken
	fmt.Fprintf(os.Stdout, "Response from `SelfAPITokenAPI.GetApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apitokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserApitoken**](UserApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiTokens

> []UserApitoken ListApiTokens(ctx).Execute()

listApiTokens



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAPITokenAPI.ListApiTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPITokenAPI.ListApiTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiTokens`: []UserApitoken
	fmt.Fprintf(os.Stdout, "Response from `SelfAPITokenAPI.ListApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApiTokensRequest struct via the builder pattern


### Return type

[**[]UserApitoken**](UserApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiToken

> UserApitoken UpdateApiToken(ctx, apitokenId).UserApitoken(userApitoken).Execute()

updateApiToken



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
	apitokenId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	userApitoken := *openapiclient.NewUserApitoken() // UserApitoken |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAPITokenAPI.UpdateApiToken(context.Background(), apitokenId).UserApitoken(userApitoken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAPITokenAPI.UpdateApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApiToken`: UserApitoken
	fmt.Fprintf(os.Stdout, "Response from `SelfAPITokenAPI.UpdateApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apitokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userApitoken** | [**UserApitoken**](UserApitoken.md) |  | 

### Return type

[**UserApitoken**](UserApitoken.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

