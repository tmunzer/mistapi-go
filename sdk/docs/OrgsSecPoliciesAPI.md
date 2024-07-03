# \OrgsSecPoliciesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSecPolicies**](OrgsSecPoliciesAPI.md#CreateOrgSecPolicies) | **Post** /api/v1/orgs/{org_id}/secpolicies | createOrgSecPolicies
[**DeleteOrgSecPolicy**](OrgsSecPoliciesAPI.md#DeleteOrgSecPolicy) | **Delete** /api/v1/orgs/{org_id}/secpolicies/{secpolicy_id} | deleteOrgSecPolicy
[**GetOrgSecPolicy**](OrgsSecPoliciesAPI.md#GetOrgSecPolicy) | **Get** /api/v1/orgs/{org_id}/secpolicies/{secpolicy_id} | getOrgSecPolicy
[**ListOrgSecPolicies**](OrgsSecPoliciesAPI.md#ListOrgSecPolicies) | **Get** /api/v1/orgs/{org_id}/secpolicies | listOrgSecPolicies
[**UpdateOrgSecPolicies**](OrgsSecPoliciesAPI.md#UpdateOrgSecPolicies) | **Put** /api/v1/orgs/{org_id}/secpolicies/{secpolicy_id} | updateOrgSecPolicies



## CreateOrgSecPolicies

> Secpolicy CreateOrgSecPolicies(ctx, orgId).Secpolicy(secpolicy).Execute()

createOrgSecPolicies



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
	secpolicy := *openapiclient.NewSecpolicy() // Secpolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSecPoliciesAPI.CreateOrgSecPolicies(context.Background(), orgId).Secpolicy(secpolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSecPoliciesAPI.CreateOrgSecPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSecPolicies`: Secpolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsSecPoliciesAPI.CreateOrgSecPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSecPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secpolicy** | [**Secpolicy**](Secpolicy.md) |  | 

### Return type

[**Secpolicy**](Secpolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgSecPolicy

> DeleteOrgSecPolicy(ctx, orgId, secpolicyId).Execute()

deleteOrgSecPolicy



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
	secpolicyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSecPoliciesAPI.DeleteOrgSecPolicy(context.Background(), orgId, secpolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSecPoliciesAPI.DeleteOrgSecPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**secpolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSecPolicyRequest struct via the builder pattern


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


## GetOrgSecPolicy

> Secpolicy GetOrgSecPolicy(ctx, orgId, secpolicyId).Execute()

getOrgSecPolicy



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
	secpolicyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSecPoliciesAPI.GetOrgSecPolicy(context.Background(), orgId, secpolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSecPoliciesAPI.GetOrgSecPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSecPolicy`: Secpolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsSecPoliciesAPI.GetOrgSecPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**secpolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSecPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Secpolicy**](Secpolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSecPolicies

> []Secpolicy ListOrgSecPolicies(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgSecPolicies



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSecPoliciesAPI.ListOrgSecPolicies(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSecPoliciesAPI.ListOrgSecPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSecPolicies`: []Secpolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsSecPoliciesAPI.ListOrgSecPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSecPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Secpolicy**](Secpolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSecPolicies

> Secpolicy UpdateOrgSecPolicies(ctx, orgId, secpolicyId).Secpolicy(secpolicy).Execute()

updateOrgSecPolicies



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
	secpolicyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	secpolicy := *openapiclient.NewSecpolicy() // Secpolicy | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSecPoliciesAPI.UpdateOrgSecPolicies(context.Background(), orgId, secpolicyId).Secpolicy(secpolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSecPoliciesAPI.UpdateOrgSecPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSecPolicies`: Secpolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsSecPoliciesAPI.UpdateOrgSecPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**secpolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSecPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secpolicy** | [**Secpolicy**](Secpolicy.md) | Request Body | 

### Return type

[**Secpolicy**](Secpolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

