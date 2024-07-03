# \AdminsLookupAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Lookup**](AdminsLookupAPI.md#Lookup) | **Post** /api/v1/login/lookup | lookup



## Lookup

> ResponseLoginLookup Lookup(ctx).EmailString(emailString).Execute()

lookup



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
	emailString := *openapiclient.NewEmailString("Email_example") // EmailString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminsLookupAPI.Lookup(context.Background()).EmailString(emailString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminsLookupAPI.Lookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Lookup`: ResponseLoginLookup
	fmt.Fprintf(os.Stdout, "Response from `AdminsLookupAPI.Lookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailString** | [**EmailString**](EmailString.md) | Request Body | 

### Return type

[**ResponseLoginLookup**](ResponseLoginLookup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

