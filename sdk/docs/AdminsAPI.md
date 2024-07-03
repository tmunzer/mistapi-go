# \AdminsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdminRegistrationInfo**](AdminsAPI.md#GetAdminRegistrationInfo) | **Get** /api/v1/register/recaptcha | getAdminRegistrationInfo
[**RegisterNewAdmin**](AdminsAPI.md#RegisterNewAdmin) | **Post** /api/v1/register | registerNewAdmin
[**VerifyAdminInvite**](AdminsAPI.md#VerifyAdminInvite) | **Post** /api/v1/invite/verify/{token} | verifyAdminInvite
[**VerifyRegistration**](AdminsAPI.md#VerifyRegistration) | **Post** /api/v1/register/verify/{token} | verifyRegistration



## GetAdminRegistrationInfo

> Recaptcha GetAdminRegistrationInfo(ctx).RecaptchaFlavor(recaptchaFlavor).Execute()

getAdminRegistrationInfo



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
	recaptchaFlavor := openapiclient.recaptcha_flavor("") // RecaptchaFlavor |  (optional) (default to "google")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminsAPI.GetAdminRegistrationInfo(context.Background()).RecaptchaFlavor(recaptchaFlavor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsAPI.GetAdminRegistrationInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdminRegistrationInfo`: Recaptcha
	fmt.Fprintf(os.Stdout, "Response from `AdminsAPI.GetAdminRegistrationInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRegistrationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recaptchaFlavor** | [**RecaptchaFlavor**](RecaptchaFlavor.md) |  | [default to &quot;google&quot;]

### Return type

[**Recaptcha**](Recaptcha.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNewAdmin

> RegisterNewAdmin(ctx).AdminInvite(adminInvite).Execute()

registerNewAdmin



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
	adminInvite := *openapiclient.NewAdminInvite("test@mistsys.com", "John", "Smith", "Smith LLC", "foryoureyesonly", "Recaptcha_example") // AdminInvite | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminsAPI.RegisterNewAdmin(context.Background()).AdminInvite(adminInvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsAPI.RegisterNewAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNewAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminInvite** | [**AdminInvite**](AdminInvite.md) | Request Body | 

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


## VerifyAdminInvite

> VerifyAdminInvite(ctx, token).Execute()

verifyAdminInvite



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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminsAPI.VerifyAdminInvite(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsAPI.VerifyAdminInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAdminInviteRequest struct via the builder pattern


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


## VerifyRegistration

> ResponseVerifyTokenSuccess VerifyRegistration(ctx, token).Execute()

verifyRegistration



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
	token := "token_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminsAPI.VerifyRegistration(context.Background(), token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsAPI.VerifyRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyRegistration`: ResponseVerifyTokenSuccess
	fmt.Fprintf(os.Stdout, "Response from `AdminsAPI.VerifyRegistration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseVerifyTokenSuccess**](ResponseVerifyTokenSuccess.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

