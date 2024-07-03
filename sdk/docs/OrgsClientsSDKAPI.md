# \OrgsClientsSDKAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSdkClient**](OrgsClientsSDKAPI.md#UpdateSdkClient) | **Put** /api/v1/orgs/{org_id}/sdkclients/{sdkclient_id} | updateSdkClient



## UpdateSdkClient

> UpdateSdkClient(ctx, orgId, sdkclientId).NameString(nameString).Execute()

updateSdkClient



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
	sdkclientId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nameString := *openapiclient.NewNameString() // NameString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsClientsSDKAPI.UpdateSdkClient(context.Background(), orgId, sdkclientId).NameString(nameString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsSDKAPI.UpdateSdkClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdkclientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSdkClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nameString** | [**NameString**](NameString.md) | Request Body | 

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

