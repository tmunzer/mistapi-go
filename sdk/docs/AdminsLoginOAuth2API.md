# \AdminsLoginOAuth2API

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauth2AuthorizationUrlForLogin**](AdminsLoginOAuth2API.md#GetOauth2AuthorizationUrlForLogin) | **Get** /api/v1/login/oauth/{provider} | getOauth2AuthorizationUrlForLogin
[**LoginOauth2**](AdminsLoginOAuth2API.md#LoginOauth2) | **Post** /api/v1/login/oauth/{provider} | loginOauth2
[**UnlinkOauth2Provider**](AdminsLoginOAuth2API.md#UnlinkOauth2Provider) | **Delete** /api/v1/login/oauth/{provider} | unlinkOauth2Provider



## GetOauth2AuthorizationUrlForLogin

> ResponseLoginOauthUrl GetOauth2AuthorizationUrlForLogin(ctx, provider).Forward(forward).Execute()

getOauth2AuthorizationUrlForLogin



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
	provider := "provider_example" // string | 
	forward := "http://manage.mist.com/oauth/callback.html" // string | callback URL (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminsLoginOAuth2API.GetOauth2AuthorizationUrlForLogin(context.Background(), provider).Forward(forward).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsLoginOAuth2API.GetOauth2AuthorizationUrlForLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauth2AuthorizationUrlForLogin`: ResponseLoginOauthUrl
	fmt.Fprintf(os.Stdout, "Response from `AdminsLoginOAuth2API.GetOauth2AuthorizationUrlForLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauth2AuthorizationUrlForLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forward** | **string** | callback URL | 

### Return type

[**ResponseLoginOauthUrl**](ResponseLoginOauthUrl.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginOauth2

> LoginOauth2(ctx, provider).CodeString(codeString).Execute()

loginOauth2



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
	provider := "provider_example" // string | 
	codeString := *openapiclient.NewCodeString("Code_example") // CodeString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminsLoginOAuth2API.LoginOauth2(context.Background(), provider).CodeString(codeString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsLoginOAuth2API.LoginOauth2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoginOauth2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeString** | [**CodeString**](CodeString.md) | Request Body | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkOauth2Provider

> UnlinkOauth2Provider(ctx, provider).Execute()

unlinkOauth2Provider



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminsLoginOAuth2API.UnlinkOauth2Provider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsLoginOAuth2API.UnlinkOauth2Provider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkOauth2ProviderRequest struct via the builder pattern


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

