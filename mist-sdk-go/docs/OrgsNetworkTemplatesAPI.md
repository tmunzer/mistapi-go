# \OrgsNetworkTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNetworkTemplate**](OrgsNetworkTemplatesAPI.md#CreateOrgNetworkTemplate) | **Post** /api/v1/orgs/{org_id}/networktemplates | createOrgNetworkTemplate
[**DeleteOrgNetworkTemplate**](OrgsNetworkTemplatesAPI.md#DeleteOrgNetworkTemplate) | **Delete** /api/v1/orgs/{org_id}/networktemplates/{networktemplate_id} | deleteOrgNetworkTemplate
[**GetOrgNetworkTemplate**](OrgsNetworkTemplatesAPI.md#GetOrgNetworkTemplate) | **Get** /api/v1/orgs/{org_id}/networktemplates/{networktemplate_id} | getOrgNetworkTemplate
[**ListOrgNetworkTemplates**](OrgsNetworkTemplatesAPI.md#ListOrgNetworkTemplates) | **Get** /api/v1/orgs/{org_id}/networktemplates | listOrgNetworkTemplates
[**UpdateOrgNetworkTemplates**](OrgsNetworkTemplatesAPI.md#UpdateOrgNetworkTemplates) | **Put** /api/v1/orgs/{org_id}/networktemplates/{networktemplate_id} | updateOrgNetworkTemplates



## CreateOrgNetworkTemplate

> NetworkTemplate CreateOrgNetworkTemplate(ctx, orgId).NetworkTemplate(networkTemplate).Execute()

createOrgNetworkTemplate



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
	networkTemplate := *openapiclient.NewNetworkTemplate() // NetworkTemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNetworkTemplatesAPI.CreateOrgNetworkTemplate(context.Background(), orgId).NetworkTemplate(networkTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworkTemplatesAPI.CreateOrgNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgNetworkTemplate`: NetworkTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworkTemplatesAPI.CreateOrgNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgNetworkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTemplate** | [**NetworkTemplate**](NetworkTemplate.md) | Request Body | 

### Return type

[**NetworkTemplate**](NetworkTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgNetworkTemplate

> DeleteOrgNetworkTemplate(ctx, orgId, networktemplateId).Execute()

deleteOrgNetworkTemplate



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
	networktemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsNetworkTemplatesAPI.DeleteOrgNetworkTemplate(context.Background(), orgId, networktemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworkTemplatesAPI.DeleteOrgNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**networktemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgNetworkTemplateRequest struct via the builder pattern


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


## GetOrgNetworkTemplate

> NetworkTemplate GetOrgNetworkTemplate(ctx, orgId, networktemplateId).Execute()

getOrgNetworkTemplate



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
	networktemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNetworkTemplatesAPI.GetOrgNetworkTemplate(context.Background(), orgId, networktemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworkTemplatesAPI.GetOrgNetworkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgNetworkTemplate`: NetworkTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworkTemplatesAPI.GetOrgNetworkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**networktemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNetworkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkTemplate**](NetworkTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgNetworkTemplates

> []NetworkTemplate ListOrgNetworkTemplates(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgNetworkTemplates



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
	resp, r, err := apiClient.OrgsNetworkTemplatesAPI.ListOrgNetworkTemplates(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworkTemplatesAPI.ListOrgNetworkTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgNetworkTemplates`: []NetworkTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworkTemplatesAPI.ListOrgNetworkTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNetworkTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]NetworkTemplate**](NetworkTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgNetworkTemplates

> NetworkTemplate UpdateOrgNetworkTemplates(ctx, orgId, networktemplateId).NetworkTemplate(networkTemplate).Execute()

updateOrgNetworkTemplates



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
	networktemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	networkTemplate := *openapiclient.NewNetworkTemplate() // NetworkTemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNetworkTemplatesAPI.UpdateOrgNetworkTemplates(context.Background(), orgId, networktemplateId).NetworkTemplate(networkTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNetworkTemplatesAPI.UpdateOrgNetworkTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgNetworkTemplates`: NetworkTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsNetworkTemplatesAPI.UpdateOrgNetworkTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**networktemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgNetworkTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkTemplate** | [**NetworkTemplate**](NetworkTemplate.md) | Request Body | 

### Return type

[**NetworkTemplate**](NetworkTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

