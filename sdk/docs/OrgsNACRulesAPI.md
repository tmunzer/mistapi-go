# \OrgsNACRulesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNacRule**](OrgsNACRulesAPI.md#CreateOrgNacRule) | **Post** /api/v1/orgs/{org_id}/nacrules | createOrgNacRule
[**DeleteOrgNacRule**](OrgsNACRulesAPI.md#DeleteOrgNacRule) | **Delete** /api/v1/orgs/{org_id}/nacrules/{nacrule_id} | deleteOrgNacRule
[**GetOrgNacRule**](OrgsNACRulesAPI.md#GetOrgNacRule) | **Get** /api/v1/orgs/{org_id}/nacrules/{nacrule_id} | getOrgNacRule
[**ListOrgNacRules**](OrgsNACRulesAPI.md#ListOrgNacRules) | **Get** /api/v1/orgs/{org_id}/nacrules | listOrgNacRules
[**UpdateOrgNacRule**](OrgsNACRulesAPI.md#UpdateOrgNacRule) | **Put** /api/v1/orgs/{org_id}/nacrules/{nacrule_id} | updateOrgNacRule



## CreateOrgNacRule

> NacRule CreateOrgNacRule(ctx, orgId).NacRule(nacRule).Execute()

createOrgNacRule



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
	nacRule := *openapiclient.NewNacRule(openapiclient.nac_rule_action(""), "Name_example") // NacRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACRulesAPI.CreateOrgNacRule(context.Background(), orgId).NacRule(nacRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACRulesAPI.CreateOrgNacRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgNacRule`: NacRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACRulesAPI.CreateOrgNacRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgNacRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nacRule** | [**NacRule**](NacRule.md) |  | 

### Return type

[**NacRule**](NacRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgNacRule

> DeleteOrgNacRule(ctx, orgId, nacruleId).Execute()

deleteOrgNacRule



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
	nacruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsNACRulesAPI.DeleteOrgNacRule(context.Background(), orgId, nacruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACRulesAPI.DeleteOrgNacRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgNacRuleRequest struct via the builder pattern


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


## GetOrgNacRule

> NacRule GetOrgNacRule(ctx, orgId, nacruleId).Execute()

getOrgNacRule



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
	nacruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACRulesAPI.GetOrgNacRule(context.Background(), orgId, nacruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACRulesAPI.GetOrgNacRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgNacRule`: NacRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACRulesAPI.GetOrgNacRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNacRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NacRule**](NacRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgNacRules

> []NacRule ListOrgNacRules(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgNacRules



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
	resp, r, err := apiClient.OrgsNACRulesAPI.ListOrgNacRules(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACRulesAPI.ListOrgNacRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgNacRules`: []NacRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACRulesAPI.ListOrgNacRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNacRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]NacRule**](NacRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgNacRule

> NacRule UpdateOrgNacRule(ctx, orgId, nacruleId).NacRule(nacRule).Execute()

updateOrgNacRule



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
	nacruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nacRule := *openapiclient.NewNacRule(openapiclient.nac_rule_action(""), "Name_example") // NacRule |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACRulesAPI.UpdateOrgNacRule(context.Background(), orgId, nacruleId).NacRule(nacRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACRulesAPI.UpdateOrgNacRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgNacRule`: NacRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACRulesAPI.UpdateOrgNacRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nacruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgNacRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nacRule** | [**NacRule**](NacRule.md) |  | 

### Return type

[**NacRule**](NacRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

