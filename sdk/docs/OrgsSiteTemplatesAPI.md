# \OrgsSiteTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSiteTemplates**](OrgsSiteTemplatesAPI.md#CreateOrgSiteTemplates) | **Post** /api/v1/orgs/{org_id}/sitetemplates | createOrgSiteTemplates
[**DeleteOrgSiteTemplate**](OrgsSiteTemplatesAPI.md#DeleteOrgSiteTemplate) | **Delete** /api/v1/orgs/{org_id}/sitetemplates/{sitetemplate_id} | deleteOrgSiteTemplate
[**GetOrgSiteTemplate**](OrgsSiteTemplatesAPI.md#GetOrgSiteTemplate) | **Get** /api/v1/orgs/{org_id}/sitetemplates/{sitetemplate_id} | getOrgSiteTemplate
[**ListOrgSiteTemplates**](OrgsSiteTemplatesAPI.md#ListOrgSiteTemplates) | **Get** /api/v1/orgs/{org_id}/sitetemplates | listOrgSiteTemplates
[**UpdateOrgSiteTemplate**](OrgsSiteTemplatesAPI.md#UpdateOrgSiteTemplate) | **Put** /api/v1/orgs/{org_id}/sitetemplates/{sitetemplate_id} | updateOrgSiteTemplate



## CreateOrgSiteTemplates

> SiteTemplate CreateOrgSiteTemplates(ctx, orgId).SiteTemplate(siteTemplate).Execute()

createOrgSiteTemplates



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
	siteTemplate := *openapiclient.NewSiteTemplate() // SiteTemplate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSiteTemplatesAPI.CreateOrgSiteTemplates(context.Background(), orgId).SiteTemplate(siteTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSiteTemplatesAPI.CreateOrgSiteTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSiteTemplates`: SiteTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSiteTemplatesAPI.CreateOrgSiteTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSiteTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteTemplate** | [**SiteTemplate**](SiteTemplate.md) |  | 

### Return type

[**SiteTemplate**](SiteTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgSiteTemplate

> DeleteOrgSiteTemplate(ctx, orgId, sitetemplateId).Execute()

deleteOrgSiteTemplate



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
	sitetemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSiteTemplatesAPI.DeleteOrgSiteTemplate(context.Background(), orgId, sitetemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSiteTemplatesAPI.DeleteOrgSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sitetemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSiteTemplateRequest struct via the builder pattern


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


## GetOrgSiteTemplate

> SiteTemplate GetOrgSiteTemplate(ctx, orgId, sitetemplateId).Execute()

getOrgSiteTemplate



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
	sitetemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSiteTemplatesAPI.GetOrgSiteTemplate(context.Background(), orgId, sitetemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSiteTemplatesAPI.GetOrgSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSiteTemplate`: SiteTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSiteTemplatesAPI.GetOrgSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sitetemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SiteTemplate**](SiteTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSiteTemplates

> []SiteTemplate ListOrgSiteTemplates(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgSiteTemplates



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
	resp, r, err := apiClient.OrgsSiteTemplatesAPI.ListOrgSiteTemplates(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSiteTemplatesAPI.ListOrgSiteTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSiteTemplates`: []SiteTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSiteTemplatesAPI.ListOrgSiteTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSiteTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]SiteTemplate**](SiteTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSiteTemplate

> SiteTemplate UpdateOrgSiteTemplate(ctx, orgId, sitetemplateId).SiteTemplate(siteTemplate).Execute()

updateOrgSiteTemplate



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
	sitetemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	siteTemplate := *openapiclient.NewSiteTemplate() // SiteTemplate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSiteTemplatesAPI.UpdateOrgSiteTemplate(context.Background(), orgId, sitetemplateId).SiteTemplate(siteTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSiteTemplatesAPI.UpdateOrgSiteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSiteTemplate`: SiteTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSiteTemplatesAPI.UpdateOrgSiteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sitetemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSiteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteTemplate** | [**SiteTemplate**](SiteTemplate.md) |  | 

### Return type

[**SiteTemplate**](SiteTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

