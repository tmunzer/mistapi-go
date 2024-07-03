# \OrgsNACIDPAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateOrgIdpCredential**](OrgsNACIDPAPI.md#ValidateOrgIdpCredential) | **Post** /api/v1/orgs/{org_id}/mist_nac/test_idp | validateOrgIdpCredential



## ValidateOrgIdpCredential

> WebsocketSession ValidateOrgIdpCredential(ctx, orgId).UsernamePassword(usernamePassword).Execute()

validateOrgIdpCredential



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
	usernamePassword := *openapiclient.NewUsernamePassword() // UsernamePassword |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACIDPAPI.ValidateOrgIdpCredential(context.Background(), orgId).UsernamePassword(usernamePassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACIDPAPI.ValidateOrgIdpCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateOrgIdpCredential`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACIDPAPI.ValidateOrgIdpCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateOrgIdpCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usernamePassword** | [**UsernamePassword**](UsernamePassword.md) |  | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

