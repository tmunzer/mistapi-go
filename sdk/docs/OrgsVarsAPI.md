# \OrgsVarsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchOrgVars**](OrgsVarsAPI.md#SearchOrgVars) | **Get** /api/v1/orgs/{org_id}/vars/search | searchOrgVars



## SearchOrgVars

> ResponseSearchVar SearchOrgVars(ctx, orgId).SiteId(siteId).Vars(vars).Src(src).Limit(limit).Page(page).Execute()

searchOrgVars



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
	siteId := "siteId_example" // string |  (optional)
	vars := "vars_example" // string |  (optional)
	src := openapiclient.var_source("") // VarSource |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsVarsAPI.SearchOrgVars(context.Background(), orgId).SiteId(siteId).Vars(vars).Src(src).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVarsAPI.SearchOrgVars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgVars`: ResponseSearchVar
	fmt.Fprintf(os.Stdout, "Response from `OrgsVarsAPI.SearchOrgVars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgVarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** |  | 
 **vars** | **string** |  | 
 **src** | [**VarSource**](VarSource.md) |  | 
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**ResponseSearchVar**](ResponseSearchVar.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

