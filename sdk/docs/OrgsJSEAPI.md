# \OrgsJSEAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgJsecCredential**](OrgsJSEAPI.md#DeleteOrgJsecCredential) | **Delete** /api/v1/orgs/{org_id}/setting/jse/setup | deleteOrgJsecCredential
[**GetOrgJseInfo**](OrgsJSEAPI.md#GetOrgJseInfo) | **Get** /api/v1/orgs/{org_id}/setting/jse/info | getOrgJseInfo
[**GetOrgJsecCredential**](OrgsJSEAPI.md#GetOrgJsecCredential) | **Get** /api/v1/orgs/{org_id}/setting/jse/setup | getOrgJsecCredential
[**SetupOrgJsecCredential**](OrgsJSEAPI.md#SetupOrgJsecCredential) | **Post** /api/v1/orgs/{org_id}/setting/jse/setup | setupOrgJsecCredential



## DeleteOrgJsecCredential

> DeleteOrgJsecCredential(ctx, orgId).Execute()

deleteOrgJsecCredential



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsJSEAPI.DeleteOrgJsecCredential(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSEAPI.DeleteOrgJsecCredential``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrgJsecCredentialRequest struct via the builder pattern


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


## GetOrgJseInfo

> AccountJseInfo GetOrgJseInfo(ctx, orgId).Execute()

getOrgJseInfo



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsJSEAPI.GetOrgJseInfo(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSEAPI.GetOrgJseInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgJseInfo`: AccountJseInfo
	fmt.Fprintf(os.Stdout, "Response from `OrgsJSEAPI.GetOrgJseInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgJseInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountJseInfo**](AccountJseInfo.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgJsecCredential

> AccountJseInfo GetOrgJsecCredential(ctx, orgId).Execute()

getOrgJsecCredential



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsJSEAPI.GetOrgJsecCredential(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSEAPI.GetOrgJsecCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgJsecCredential`: AccountJseInfo
	fmt.Fprintf(os.Stdout, "Response from `OrgsJSEAPI.GetOrgJsecCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgJsecCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountJseInfo**](AccountJseInfo.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetupOrgJsecCredential

> AccountJseInfo SetupOrgJsecCredential(ctx, orgId).AccountJseConfig(accountJseConfig).Execute()

setupOrgJsecCredential



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
	accountJseConfig := *openapiclient.NewAccountJseConfig("foryoureyesonly", "john@abc.com") // AccountJseConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsJSEAPI.SetupOrgJsecCredential(context.Background(), orgId).AccountJseConfig(accountJseConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsJSEAPI.SetupOrgJsecCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetupOrgJsecCredential`: AccountJseInfo
	fmt.Fprintf(os.Stdout, "Response from `OrgsJSEAPI.SetupOrgJsecCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetupOrgJsecCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountJseConfig** | [**AccountJseConfig**](AccountJseConfig.md) |  | 

### Return type

[**AccountJseInfo**](AccountJseInfo.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

