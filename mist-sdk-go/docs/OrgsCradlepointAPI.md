# \OrgsCradlepointAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgCradlepointConnection**](OrgsCradlepointAPI.md#DeleteOrgCradlepointConnection) | **Delete** /api/v1/orgs/{org_id}/setting/cradlepoint/setup | deleteOrgCradlepointConnection
[**SetupOrgCradlepointConnectionToMist**](OrgsCradlepointAPI.md#SetupOrgCradlepointConnectionToMist) | **Post** /api/v1/orgs/{org_id}/setting/cradlepoint/setup | setupOrgCradlepointConnectionToMist
[**SyncOrgCradlepointRouters**](OrgsCradlepointAPI.md#SyncOrgCradlepointRouters) | **Post** /api/v1/orgs/{org_id}/setting/cradlepoint/sync | syncOrgCradlepointRouters
[**UpdateOrgCradlepointConnectionToMist**](OrgsCradlepointAPI.md#UpdateOrgCradlepointConnectionToMist) | **Put** /api/v1/orgs/{org_id}/setting/cradlepoint/setup | updateOrgCradlepointConnectionToMist



## DeleteOrgCradlepointConnection

> DeleteOrgCradlepointConnection(ctx, orgId).Execute()

deleteOrgCradlepointConnection



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
	r, err := apiClient.OrgsCradlepointAPI.DeleteOrgCradlepointConnection(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsCradlepointAPI.DeleteOrgCradlepointConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrgCradlepointConnectionRequest struct via the builder pattern


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


## SetupOrgCradlepointConnectionToMist

> SetupOrgCradlepointConnectionToMist(ctx, orgId).AccountCradlepointConfig(accountCradlepointConfig).Execute()

setupOrgCradlepointConnectionToMist



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
	accountCradlepointConfig := *openapiclient.NewAccountCradlepointConfig() // AccountCradlepointConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsCradlepointAPI.SetupOrgCradlepointConnectionToMist(context.Background(), orgId).AccountCradlepointConfig(accountCradlepointConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsCradlepointAPI.SetupOrgCradlepointConnectionToMist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetupOrgCradlepointConnectionToMistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountCradlepointConfig** | [**AccountCradlepointConfig**](AccountCradlepointConfig.md) |  | 

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


## SyncOrgCradlepointRouters

> SyncOrgCradlepointRouters(ctx, orgId).Execute()

syncOrgCradlepointRouters



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
	r, err := apiClient.OrgsCradlepointAPI.SyncOrgCradlepointRouters(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsCradlepointAPI.SyncOrgCradlepointRouters``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSyncOrgCradlepointRoutersRequest struct via the builder pattern


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


## UpdateOrgCradlepointConnectionToMist

> UpdateOrgCradlepointConnectionToMist(ctx, orgId).AccountCradlepointConfig(accountCradlepointConfig).Execute()

updateOrgCradlepointConnectionToMist



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
	accountCradlepointConfig := *openapiclient.NewAccountCradlepointConfig() // AccountCradlepointConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsCradlepointAPI.UpdateOrgCradlepointConnectionToMist(context.Background(), orgId).AccountCradlepointConfig(accountCradlepointConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsCradlepointAPI.UpdateOrgCradlepointConnectionToMist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateOrgCradlepointConnectionToMistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountCradlepointConfig** | [**AccountCradlepointConfig**](AccountCradlepointConfig.md) |  | 

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

