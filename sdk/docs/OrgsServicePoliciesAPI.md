# \OrgsServicePoliciesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgServicePolicy**](OrgsServicePoliciesAPI.md#CreateOrgServicePolicy) | **Post** /api/v1/orgs/{org_id}/servicepolicies | createOrgServicePolicy
[**DeleteOrgServicePolicy**](OrgsServicePoliciesAPI.md#DeleteOrgServicePolicy) | **Delete** /api/v1/orgs/{org_id}/servicepolicies/{servicepolicy_id} | deleteOrgServicePolicy
[**GetOrgServicePolicy**](OrgsServicePoliciesAPI.md#GetOrgServicePolicy) | **Get** /api/v1/orgs/{org_id}/servicepolicies/{servicepolicy_id} | getOrgServicePolicy
[**ListOrgServicePolicies**](OrgsServicePoliciesAPI.md#ListOrgServicePolicies) | **Get** /api/v1/orgs/{org_id}/servicepolicies | listOrgServicePolicies
[**UpdateOrgServicePolicy**](OrgsServicePoliciesAPI.md#UpdateOrgServicePolicy) | **Put** /api/v1/orgs/{org_id}/servicepolicies/{servicepolicy_id} | updateOrgServicePolicy



## CreateOrgServicePolicy

> ServicePolicy CreateOrgServicePolicy(ctx, orgId).ServicePolicy(servicePolicy).Execute()

createOrgServicePolicy



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
	servicePolicy := *openapiclient.NewServicePolicy() // ServicePolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsServicePoliciesAPI.CreateOrgServicePolicy(context.Background(), orgId).ServicePolicy(servicePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicePoliciesAPI.CreateOrgServicePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgServicePolicy`: ServicePolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicePoliciesAPI.CreateOrgServicePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgServicePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servicePolicy** | [**ServicePolicy**](ServicePolicy.md) |  | 

### Return type

[**ServicePolicy**](ServicePolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgServicePolicy

> DeleteOrgServicePolicy(ctx, orgId, servicepolicyId).Execute()

deleteOrgServicePolicy



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
	servicepolicyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsServicePoliciesAPI.DeleteOrgServicePolicy(context.Background(), orgId, servicepolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicePoliciesAPI.DeleteOrgServicePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**servicepolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgServicePolicyRequest struct via the builder pattern


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


## GetOrgServicePolicy

> ServicePolicy GetOrgServicePolicy(ctx, orgId, servicepolicyId).Execute()

getOrgServicePolicy



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
	servicepolicyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsServicePoliciesAPI.GetOrgServicePolicy(context.Background(), orgId, servicepolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicePoliciesAPI.GetOrgServicePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgServicePolicy`: ServicePolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicePoliciesAPI.GetOrgServicePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**servicepolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgServicePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServicePolicy**](ServicePolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgServicePolicies

> []ServicePolicy ListOrgServicePolicies(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgServicePolicies



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
	resp, r, err := apiClient.OrgsServicePoliciesAPI.ListOrgServicePolicies(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicePoliciesAPI.ListOrgServicePolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgServicePolicies`: []ServicePolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicePoliciesAPI.ListOrgServicePolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgServicePoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]ServicePolicy**](ServicePolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgServicePolicy

> ServicePolicy UpdateOrgServicePolicy(ctx, orgId, servicepolicyId).ServicePolicy(servicePolicy).Execute()

updateOrgServicePolicy



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
	servicepolicyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	servicePolicy := *openapiclient.NewServicePolicy() // ServicePolicy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsServicePoliciesAPI.UpdateOrgServicePolicy(context.Background(), orgId, servicepolicyId).ServicePolicy(servicePolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsServicePoliciesAPI.UpdateOrgServicePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgServicePolicy`: ServicePolicy
	fmt.Fprintf(os.Stdout, "Response from `OrgsServicePoliciesAPI.UpdateOrgServicePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**servicepolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgServicePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **servicePolicy** | [**ServicePolicy**](ServicePolicy.md) |  | 

### Return type

[**ServicePolicy**](ServicePolicy.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

