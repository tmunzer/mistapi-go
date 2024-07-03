# \OrgsServicesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgService**](OrgsServicesAPI.md#CreateOrgService) | **Post** /api/v1/orgs/{org_id}/services | createOrgService
[**DeleteOrgService**](OrgsServicesAPI.md#DeleteOrgService) | **Delete** /api/v1/orgs/{org_id}/services/{service_id} | deleteOrgService
[**GetOrgService**](OrgsServicesAPI.md#GetOrgService) | **Get** /api/v1/orgs/{org_id}/services/{service_id} | getOrgService
[**ListOrgServices**](OrgsServicesAPI.md#ListOrgServices) | **Get** /api/v1/orgs/{org_id}/services | listOrgServices
[**UpdateOrgService**](OrgsServicesAPI.md#UpdateOrgService) | **Put** /api/v1/orgs/{org_id}/services/{service_id} | updateOrgService



## CreateOrgService

> Service CreateOrgService(ctx, orgId).Service(service).Execute()

createOrgService



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
	service := *openapiclient.NewService() // Service |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsServicesAPI.CreateOrgService(context.Background(), orgId).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicesAPI.CreateOrgService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgService`: Service
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicesAPI.CreateOrgService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **service** | [**Service**](Service.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgService

> DeleteOrgService(ctx, orgId, serviceId).Execute()

deleteOrgService



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
	serviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsServicesAPI.DeleteOrgService(context.Background(), orgId, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicesAPI.DeleteOrgService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgServiceRequest struct via the builder pattern


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


## GetOrgService

> Service GetOrgService(ctx, orgId, serviceId).Execute()

getOrgService



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
	serviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsServicesAPI.GetOrgService(context.Background(), orgId, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicesAPI.GetOrgService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgService`: Service
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicesAPI.GetOrgService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Service**](Service.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgServices

> []Service ListOrgServices(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgServices



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsServicesAPI.ListOrgServices(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicesAPI.ListOrgServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgServices`: []Service
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicesAPI.ListOrgServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Service**](Service.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgService

> Service UpdateOrgService(ctx, orgId, serviceId).Service(service).Execute()

updateOrgService



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
	serviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	service := *openapiclient.NewService() // Service |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsServicesAPI.UpdateOrgService(context.Background(), orgId, serviceId).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicesAPI.UpdateOrgService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgService`: Service
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicesAPI.UpdateOrgService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **service** | [**Service**](Service.md) |  | 

### Return type

[**Service**](Service.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

