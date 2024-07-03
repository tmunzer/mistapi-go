# \OrgsDevicesZscalerAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgZscalerCredential**](OrgsDevicesZscalerAPI.md#DeleteOrgZscalerCredential) | **Delete** /api/v1/orgs/{org_id}/setting/zscaler/setup | deleteOrgZscalerCredential
[**GetOrgZscalerCredential**](OrgsDevicesZscalerAPI.md#GetOrgZscalerCredential) | **Get** /api/v1/orgs/{org_id}/setting/zscaler/setup | getOrgZscalerCredential
[**SetupOrgZscalerCredential**](OrgsDevicesZscalerAPI.md#SetupOrgZscalerCredential) | **Post** /api/v1/orgs/{org_id}/setting/zscaler/setup | setupOrgZscalerCredential



## DeleteOrgZscalerCredential

> DeleteOrgZscalerCredential(ctx, orgId).Execute()

deleteOrgZscalerCredential



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
	r, err := apiClient.OrgsDevicesZscalerAPI.DeleteOrgZscalerCredential(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesZscalerAPI.DeleteOrgZscalerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgZscalerCredentialRequest struct via the builder pattern


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


## GetOrgZscalerCredential

> AccountZscalerInfo GetOrgZscalerCredential(ctx, orgId).Execute()

getOrgZscalerCredential



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
	resp, r, err := apiClient.OrgsDevicesZscalerAPI.GetOrgZscalerCredential(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesZscalerAPI.GetOrgZscalerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgZscalerCredential`: AccountZscalerInfo
	fmt.Fprintf(os.Stdout, "Response from `OrgsDevicesZscalerAPI.GetOrgZscalerCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgZscalerCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountZscalerInfo**](AccountZscalerInfo.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetupOrgZscalerCredential

> SetupOrgZscalerCredential(ctx, orgId).AccountZscalerConfig(accountZscalerConfig).Execute()

setupOrgZscalerCredential



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
	accountZscalerConfig := *openapiclient.NewAccountZscalerConfig("zscalerbeta.net", "K35vrZcK3JvrZc", "password", "john@nmo.com") // AccountZscalerConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDevicesZscalerAPI.SetupOrgZscalerCredential(context.Background(), orgId).AccountZscalerConfig(accountZscalerConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDevicesZscalerAPI.SetupOrgZscalerCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetupOrgZscalerCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountZscalerConfig** | [**AccountZscalerConfig**](AccountZscalerConfig.md) |  | 

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

