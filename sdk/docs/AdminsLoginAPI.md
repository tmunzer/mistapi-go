# \AdminsLoginAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](AdminsLoginAPI.md#Login) | **Post** /api/v1/login | login
[**TwoFactor**](AdminsLoginAPI.md#TwoFactor) | **Post** /api/v1/login/two_factor | twoFactor



## Login

> ResponseLoginSuccess Login(ctx).Login(login).Execute()

login



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
	login := *openapiclient.NewLogin("test@mistsys.com", "foryoureyesonly") // Login |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminsLoginAPI.Login(context.Background()).Login(login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsLoginAPI.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: ResponseLoginSuccess
	fmt.Fprintf(os.Stdout, "Response from `AdminsLoginAPI.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **login** | [**Login**](Login.md) |  | 

### Return type

[**ResponseLoginSuccess**](ResponseLoginSuccess.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TwoFactor

> TwoFactor(ctx).TwoFactorString(twoFactorString).Execute()

twoFactor



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
	twoFactorString := *openapiclient.NewTwoFactorString("123456") // TwoFactorString |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminsLoginAPI.TwoFactor(context.Background()).TwoFactorString(twoFactorString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsLoginAPI.TwoFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTwoFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **twoFactorString** | [**TwoFactorString**](TwoFactorString.md) |  | 

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

