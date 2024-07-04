# \SelfMFAAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateSecretFor2faVerification**](SelfMFAAPI.md#GenerateSecretFor2faVerification) | **Get** /api/v1/self/two_factor/token | generateSecretFor2faVerification
[**VerifyTwoFactor**](SelfMFAAPI.md#VerifyTwoFactor) | **Post** /api/v1/self/two_factor/verify | verifyTwoFactor



## GenerateSecretFor2faVerification

> ResponseTwoFactorJson GenerateSecretFor2faVerification(ctx).By(by).Execute()

generateSecretFor2faVerification



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
	by := openapiclient.mfa_secret_type("") // MfaSecretType | if `by`==`qrcode`, returns the secret as a qrcode image (optional) (default to "qrcode")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfMFAAPI.GenerateSecretFor2faVerification(context.Background()).By(by).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfMFAAPI.GenerateSecretFor2faVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSecretFor2faVerification`: ResponseTwoFactorJson
	fmt.Fprintf(os.Stdout, "Response from `SelfMFAAPI.GenerateSecretFor2faVerification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSecretFor2faVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **by** | [**MfaSecretType**](MfaSecretType.md) | if &#x60;by&#x60;&#x3D;&#x3D;&#x60;qrcode&#x60;, returns the secret as a qrcode image | [default to &quot;qrcode&quot;]

### Return type

[**ResponseTwoFactorJson**](ResponseTwoFactorJson.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyTwoFactor

> VerifyTwoFactor(ctx).TwoFactorCode(twoFactorCode).Execute()

verifyTwoFactor



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
	twoFactorCode := *openapiclient.NewTwoFactorCode("123456") // TwoFactorCode | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SelfMFAAPI.VerifyTwoFactor(context.Background()).TwoFactorCode(twoFactorCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfMFAAPI.VerifyTwoFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyTwoFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorCode** | [**TwoFactorCode**](TwoFactorCode.md) | Request Body | 

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

