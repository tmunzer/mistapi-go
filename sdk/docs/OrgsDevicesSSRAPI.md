# \OrgsDevicesSSRAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrg128TRegistrationCommands**](OrgsDevicesSSRAPI.md#GetOrg128TRegistrationCommands) | **Get** /api/v1/orgs/{org_id}/128routers/register_cmd | getOrg128TRegistrationCommands



## GetOrg128TRegistrationCommands

> ReponseRouter128tRegisterCmd GetOrg128TRegistrationCommands(ctx, orgId).Execute()

getOrg128TRegistrationCommands



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
	resp, r, err := apiClient.OrgsDevicesSSRAPI.GetOrg128TRegistrationCommands(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesSSRAPI.GetOrg128TRegistrationCommands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrg128TRegistrationCommands`: ReponseRouter128tRegisterCmd
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesSSRAPI.GetOrg128TRegistrationCommands`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrg128TRegistrationCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReponseRouter128tRegisterCmd**](ReponseRouter128tRegisterCmd.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

