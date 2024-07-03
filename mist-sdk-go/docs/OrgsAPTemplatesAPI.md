# \OrgsAPTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAptemplate**](OrgsAPTemplatesAPI.md#CreateOrgAptemplate) | **Post** /api/v1/orgs/{org_id}/aptemplates | createOrgAptemplate
[**DeleteOrgAptemplate**](OrgsAPTemplatesAPI.md#DeleteOrgAptemplate) | **Delete** /api/v1/orgs/{org_id}/aptemplates/{aptemplate_id} | deleteOrgAptemplate
[**GetOrgAptemplate**](OrgsAPTemplatesAPI.md#GetOrgAptemplate) | **Get** /api/v1/orgs/{org_id}/aptemplates/{aptemplate_id} | getOrgAptemplate
[**ListOrgAptemplates**](OrgsAPTemplatesAPI.md#ListOrgAptemplates) | **Get** /api/v1/orgs/{org_id}/aptemplates | listOrgAptemplates
[**UpdateOrgAptemplate**](OrgsAPTemplatesAPI.md#UpdateOrgAptemplate) | **Put** /api/v1/orgs/{org_id}/aptemplates/{aptemplate_id} | updateOrgAptemplate



## CreateOrgAptemplate

> ApTemplate CreateOrgAptemplate(ctx, orgId).ApTemplate(apTemplate).Execute()

createOrgAptemplate



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
	apTemplate := *openapiclient.NewApTemplate(*openapiclient.NewApTemplateMatching()) // ApTemplate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPTemplatesAPI.CreateOrgAptemplate(context.Background(), orgId).ApTemplate(apTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPTemplatesAPI.CreateOrgAptemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgAptemplate`: ApTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPTemplatesAPI.CreateOrgAptemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAptemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apTemplate** | [**ApTemplate**](ApTemplate.md) |  | 

### Return type

[**ApTemplate**](ApTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgAptemplate

> DeleteOrgAptemplate(ctx, orgId, aptemplateId).Execute()

deleteOrgAptemplate



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
	aptemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAPTemplatesAPI.DeleteOrgAptemplate(context.Background(), orgId, aptemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPTemplatesAPI.DeleteOrgAptemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**aptemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgAptemplateRequest struct via the builder pattern


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


## GetOrgAptemplate

> ApTemplate GetOrgAptemplate(ctx, orgId, aptemplateId).Execute()

getOrgAptemplate



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
	aptemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPTemplatesAPI.GetOrgAptemplate(context.Background(), orgId, aptemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPTemplatesAPI.GetOrgAptemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgAptemplate`: ApTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPTemplatesAPI.GetOrgAptemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**aptemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAptemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApTemplate**](ApTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgAptemplates

> []ApTemplate ListOrgAptemplates(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgAptemplates



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
	resp, r, err := apiClient.OrgsAPTemplatesAPI.ListOrgAptemplates(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPTemplatesAPI.ListOrgAptemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAptemplates`: []ApTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPTemplatesAPI.ListOrgAptemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAptemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]ApTemplate**](ApTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgAptemplate

> ApTemplate UpdateOrgAptemplate(ctx, orgId, aptemplateId).ApTemplate(apTemplate).Execute()

updateOrgAptemplate



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
	aptemplateId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apTemplate := *openapiclient.NewApTemplate(*openapiclient.NewApTemplateMatching()) // ApTemplate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPTemplatesAPI.UpdateOrgAptemplate(context.Background(), orgId, aptemplateId).ApTemplate(apTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPTemplatesAPI.UpdateOrgAptemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgAptemplate`: ApTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPTemplatesAPI.UpdateOrgAptemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**aptemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgAptemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apTemplate** | [**ApTemplate**](ApTemplate.md) |  | 

### Return type

[**ApTemplate**](ApTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

