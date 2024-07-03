# \OrgsWLANTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloneOrgTemplate**](OrgsWLANTemplatesAPI.md#CloneOrgTemplate) | **Post** /api/v1/orgs/{org_id}/templates/{template_id}/clone | cloneOrgTemplate
[**CreateOrgTemplate**](OrgsWLANTemplatesAPI.md#CreateOrgTemplate) | **Post** /api/v1/orgs/{org_id}/templates | createOrgTemplate
[**DeleteOrgTemplate**](OrgsWLANTemplatesAPI.md#DeleteOrgTemplate) | **Delete** /api/v1/orgs/{org_id}/templates/{template_id} | deleteOrgTemplate
[**GetOrgTemplate**](OrgsWLANTemplatesAPI.md#GetOrgTemplate) | **Get** /api/v1/orgs/{org_id}/templates/{template_id} | getOrgTemplate
[**ListOrgTemplates**](OrgsWLANTemplatesAPI.md#ListOrgTemplates) | **Get** /api/v1/orgs/{org_id}/templates | listOrgTemplates
[**UpdateOrgTemplate**](OrgsWLANTemplatesAPI.md#UpdateOrgTemplate) | **Put** /api/v1/orgs/{org_id}/templates/{template_id} | updateOrgTemplate



## CloneOrgTemplate

> Template CloneOrgTemplate(ctx, orgId, templateId).NameString(nameString).Execute()

cloneOrgTemplate



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
	templateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nameString := *openapiclient.NewNameString() // NameString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWLANTemplatesAPI.CloneOrgTemplate(context.Background(), orgId, templateId).NameString(nameString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWLANTemplatesAPI.CloneOrgTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneOrgTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `OrgsWLANTemplatesAPI.CloneOrgTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneOrgTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nameString** | [**NameString**](NameString.md) | Request Body | 

### Return type

[**Template**](Template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgTemplate

> Template CreateOrgTemplate(ctx, orgId).Template(template).Execute()

createOrgTemplate



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
	template := *openapiclient.NewTemplate("Name_example") // Template | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWLANTemplatesAPI.CreateOrgTemplate(context.Background(), orgId).Template(template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWLANTemplatesAPI.CreateOrgTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `OrgsWLANTemplatesAPI.CreateOrgTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **template** | [**Template**](Template.md) | Request Body | 

### Return type

[**Template**](Template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgTemplate

> DeleteOrgTemplate(ctx, orgId, templateId).Execute()

deleteOrgTemplate



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
	templateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsWLANTemplatesAPI.DeleteOrgTemplate(context.Background(), orgId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWLANTemplatesAPI.DeleteOrgTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgTemplateRequest struct via the builder pattern


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


## GetOrgTemplate

> Template GetOrgTemplate(ctx, orgId, templateId).Execute()

getOrgTemplate



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
	templateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWLANTemplatesAPI.GetOrgTemplate(context.Background(), orgId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWLANTemplatesAPI.GetOrgTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `OrgsWLANTemplatesAPI.GetOrgTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Template**](Template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgTemplates

> []Template ListOrgTemplates(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgTemplates



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
	resp, r, err := apiClient.OrgsWLANTemplatesAPI.ListOrgTemplates(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWLANTemplatesAPI.ListOrgTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgTemplates`: []Template
	fmt.Fprintf(os.Stdout, "Response from `OrgsWLANTemplatesAPI.ListOrgTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Template**](Template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgTemplate

> Template UpdateOrgTemplate(ctx, orgId, templateId).Template(template).Execute()

updateOrgTemplate



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
	templateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	template := *openapiclient.NewTemplate("Name_example") // Template | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWLANTemplatesAPI.UpdateOrgTemplate(context.Background(), orgId, templateId).Template(template).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWLANTemplatesAPI.UpdateOrgTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `OrgsWLANTemplatesAPI.UpdateOrgTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **template** | [**Template**](Template.md) | Request Body | 

### Return type

[**Template**](Template.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

