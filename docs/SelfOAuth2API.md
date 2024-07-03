# \SelfOAuth2API

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOauth2UrlForLinking**](SelfOAuth2API.md#GetOauth2UrlForLinking) | **Get** /api/v1/self/oauth/{provider} | getOauth2UrlForLinking
[**LinkOauth2MistAccount**](SelfOAuth2API.md#LinkOauth2MistAccount) | **Post** /api/v1/self/oauth/{provider} | linkOauth2MistAccount



## GetOauth2UrlForLinking

> ResponseSelfOauthUrl GetOauth2UrlForLinking(ctx, provider).Forward(forward).Execute()

getOauth2UrlForLinking



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
	provider := "provider_example" // string | 
	forward := "http://manage.mist.com/oauth/callback.html" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfOAuth2API.GetOauth2UrlForLinking(context.Background(), provider).Forward(forward).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfOAuth2API.GetOauth2UrlForLinking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOauth2UrlForLinking`: ResponseSelfOauthUrl
	fmt.Fprintf(os.Stdout, "Response from `SelfOAuth2API.GetOauth2UrlForLinking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauth2UrlForLinkingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forward** | **string** |  | 

### Return type

[**ResponseSelfOauthUrl**](ResponseSelfOauthUrl.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkOauth2MistAccount

> ResponseSelfOauthLinkSuccess LinkOauth2MistAccount(ctx, provider).CodeString(codeString).Execute()

linkOauth2MistAccount



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
	provider := "provider_example" // string | 
	codeString := *openapiclient.NewCodeString("Code_example") // CodeString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfOAuth2API.LinkOauth2MistAccount(context.Background(), provider).CodeString(codeString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfOAuth2API.LinkOauth2MistAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkOauth2MistAccount`: ResponseSelfOauthLinkSuccess
	fmt.Fprintf(os.Stdout, "Response from `SelfOAuth2API.LinkOauth2MistAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkOauth2MistAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeString** | [**CodeString**](CodeString.md) | Request Body | 

### Return type

[**ResponseSelfOauthLinkSuccess**](ResponseSelfOauthLinkSuccess.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

