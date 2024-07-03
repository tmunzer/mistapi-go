# \OrgsWxRulesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWxRule**](OrgsWxRulesAPI.md#CreateOrgWxRule) | **Post** /api/v1/orgs/{org_id}/wxrules | createOrgWxRule
[**DeleteOrgWxRule**](OrgsWxRulesAPI.md#DeleteOrgWxRule) | **Delete** /api/v1/orgs/{org_id}/wxrules/{wxrule_id} | deleteOrgWxRule
[**GetOrgWxRule**](OrgsWxRulesAPI.md#GetOrgWxRule) | **Get** /api/v1/orgs/{org_id}/wxrules/{wxrule_id} | getOrgWxRule
[**GetOrgWxRulesDerived**](OrgsWxRulesAPI.md#GetOrgWxRulesDerived) | **Get** /api/v1/orgs/{org_id}/wxrules/derived | getOrgWxRulesDerived
[**ListOrgWxRules**](OrgsWxRulesAPI.md#ListOrgWxRules) | **Get** /api/v1/orgs/{org_id}/wxrules | listOrgWxRules
[**UpdateOrgWxRule**](OrgsWxRulesAPI.md#UpdateOrgWxRule) | **Put** /api/v1/orgs/{org_id}/wxrules/{wxrule_id} | updateOrgWxRule



## CreateOrgWxRule

> WxlanRule CreateOrgWxRule(ctx, orgId).WxlanRule(wxlanRule).Execute()

createOrgWxRule



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
	wxlanRule := *openapiclient.NewWxlanRule(int32(1), []string{"SrcWxtags_example"}) // WxlanRule | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxRulesAPI.CreateOrgWxRule(context.Background(), orgId).WxlanRule(wxlanRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxRulesAPI.CreateOrgWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgWxRule`: WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxRulesAPI.CreateOrgWxRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWxRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wxlanRule** | [**WxlanRule**](WxlanRule.md) | Request Body | 

### Return type

[**WxlanRule**](WxlanRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgWxRule

> DeleteOrgWxRule(ctx, orgId, wxruleId).Execute()

deleteOrgWxRule



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
	wxruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsWxRulesAPI.DeleteOrgWxRule(context.Background(), orgId, wxruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxRulesAPI.DeleteOrgWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWxRuleRequest struct via the builder pattern


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


## GetOrgWxRule

> WxlanRule GetOrgWxRule(ctx, orgId, wxruleId).Execute()

getOrgWxRule



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
	wxruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxRulesAPI.GetOrgWxRule(context.Background(), orgId, wxruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxRulesAPI.GetOrgWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgWxRule`: WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxRulesAPI.GetOrgWxRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWxRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WxlanRule**](WxlanRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgWxRulesDerived

> []WxlanRule GetOrgWxRulesDerived(ctx, orgId).Execute()

getOrgWxRulesDerived



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
	resp, r, err := apiClient.OrgsWxRulesAPI.GetOrgWxRulesDerived(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxRulesAPI.GetOrgWxRulesDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgWxRulesDerived`: []WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxRulesAPI.GetOrgWxRulesDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWxRulesDerivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WxlanRule**](WxlanRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgWxRules

> []WxlanRule ListOrgWxRules(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgWxRules



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
	resp, r, err := apiClient.OrgsWxRulesAPI.ListOrgWxRules(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxRulesAPI.ListOrgWxRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgWxRules`: []WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxRulesAPI.ListOrgWxRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWxRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]WxlanRule**](WxlanRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgWxRule

> WxlanRule UpdateOrgWxRule(ctx, orgId, wxruleId).WxlanRule(wxlanRule).Execute()

updateOrgWxRule



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
	wxruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanRule := *openapiclient.NewWxlanRule(int32(1), []string{"SrcWxtags_example"}) // WxlanRule | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxRulesAPI.UpdateOrgWxRule(context.Background(), orgId, wxruleId).WxlanRule(wxlanRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxRulesAPI.UpdateOrgWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgWxRule`: WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxRulesAPI.UpdateOrgWxRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWxRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wxlanRule** | [**WxlanRule**](WxlanRule.md) | Request Body | 

### Return type

[**WxlanRule**](WxlanRule.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

