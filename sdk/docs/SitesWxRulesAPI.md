# \SitesWxRulesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWxRule**](SitesWxRulesAPI.md#CreateSiteWxRule) | **Post** /api/v1/sites/{site_id}/wxrules | createSiteWxRule
[**DeleteSiteWxRule**](SitesWxRulesAPI.md#DeleteSiteWxRule) | **Delete** /api/v1/sites/{site_id}/wxrules/{wxrule_id} | deleteSiteWxRule
[**GetSiteWxRule**](SitesWxRulesAPI.md#GetSiteWxRule) | **Get** /api/v1/sites/{site_id}/wxrules/{wxrule_id} | getSiteWxRule
[**GetSiteWxRulesDerived**](SitesWxRulesAPI.md#GetSiteWxRulesDerived) | **Get** /api/v1/sites/{site_id}/wxrules/derived | getSiteWxRulesDerived
[**GetSiteWxRulesUsage**](SitesWxRulesAPI.md#GetSiteWxRulesUsage) | **Get** /api/v1/sites/{site_id}/stats/wxrules | getSiteWxRulesUsage
[**ListSiteWxRules**](SitesWxRulesAPI.md#ListSiteWxRules) | **Get** /api/v1/sites/{site_id}/wxrules | listSiteWxRules
[**UpdateSiteWxRule**](SitesWxRulesAPI.md#UpdateSiteWxRule) | **Put** /api/v1/sites/{site_id}/wxrules/{wxrule_id} | updateSiteWxRule



## CreateSiteWxRule

> WxlanRule CreateSiteWxRule(ctx, siteId).WxlanRule(wxlanRule).Execute()

createSiteWxRule



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanRule := *openapiclient.NewWxlanRule(int32(1), []string{"SrcWxtags_example"}) // WxlanRule | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxRulesAPI.CreateSiteWxRule(context.Background(), siteId).WxlanRule(wxlanRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxRulesAPI.CreateSiteWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWxRule`: WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `SitesWxRulesAPI.CreateSiteWxRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWxRuleRequest struct via the builder pattern


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


## DeleteSiteWxRule

> DeleteSiteWxRule(ctx, siteId, wxruleId).Execute()

deleteSiteWxRule



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesWxRulesAPI.DeleteSiteWxRule(context.Background(), siteId, wxruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxRulesAPI.DeleteSiteWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteWxRuleRequest struct via the builder pattern


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


## GetSiteWxRule

> WxlanRule GetSiteWxRule(ctx, siteId, wxruleId).Execute()

getSiteWxRule



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxRulesAPI.GetSiteWxRule(context.Background(), siteId, wxruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxRulesAPI.GetSiteWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWxRule`: WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `SitesWxRulesAPI.GetSiteWxRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWxRuleRequest struct via the builder pattern


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


## GetSiteWxRulesDerived

> []WxlanRule GetSiteWxRulesDerived(ctx, siteId).Execute()

getSiteWxRulesDerived



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxRulesAPI.GetSiteWxRulesDerived(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxRulesAPI.GetSiteWxRulesDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWxRulesDerived`: []WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `SitesWxRulesAPI.GetSiteWxRulesDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWxRulesDerivedRequest struct via the builder pattern


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


## GetSiteWxRulesUsage

> []WxruleStat GetSiteWxRulesUsage(ctx, siteId).Execute()

getSiteWxRulesUsage



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxRulesAPI.GetSiteWxRulesUsage(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxRulesAPI.GetSiteWxRulesUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWxRulesUsage`: []WxruleStat
	fmt.Fprintf(os.Stdout, "Response from `SitesWxRulesAPI.GetSiteWxRulesUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWxRulesUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WxruleStat**](WxruleStat.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteWxRules

> []WxlanRule ListSiteWxRules(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteWxRules



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxRulesAPI.ListSiteWxRules(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxRulesAPI.ListSiteWxRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteWxRules`: []WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `SitesWxRulesAPI.ListSiteWxRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteWxRulesRequest struct via the builder pattern


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


## UpdateSiteWxRule

> WxlanRule UpdateSiteWxRule(ctx, siteId, wxruleId).WxlanRule(wxlanRule).Execute()

updateSiteWxRule



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxruleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanRule := *openapiclient.NewWxlanRule(int32(1), []string{"SrcWxtags_example"}) // WxlanRule | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxRulesAPI.UpdateSiteWxRule(context.Background(), siteId, wxruleId).WxlanRule(wxlanRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxRulesAPI.UpdateSiteWxRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteWxRule`: WxlanRule
	fmt.Fprintf(os.Stdout, "Response from `SitesWxRulesAPI.UpdateSiteWxRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteWxRuleRequest struct via the builder pattern


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

