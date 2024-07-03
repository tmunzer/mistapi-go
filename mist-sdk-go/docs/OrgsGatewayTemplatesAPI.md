# \OrgsGatewayTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgGatewayTemplate**](OrgsGatewayTemplatesAPI.md#CreateOrgGatewayTemplate) | **Post** /api/v1/orgs/{org_id}/gatewaytemplates | createOrgGatewayTemplate
[**DeleteOrgGatewayTemplate**](OrgsGatewayTemplatesAPI.md#DeleteOrgGatewayTemplate) | **Delete** /api/v1/orgs/{org_id}/gatewaytemplates/{gatewaytemplate_id} | deleteOrgGatewayTemplate
[**GetOrgGatewayTemplate**](OrgsGatewayTemplatesAPI.md#GetOrgGatewayTemplate) | **Get** /api/v1/orgs/{org_id}/gatewaytemplates/{gatewaytemplate_id} | getOrgGatewayTemplate
[**ListOrgGatewayTemplates**](OrgsGatewayTemplatesAPI.md#ListOrgGatewayTemplates) | **Get** /api/v1/orgs/{org_id}/gatewaytemplates | listOrgGatewayTemplates
[**UpdateOrgGatewayTemplate**](OrgsGatewayTemplatesAPI.md#UpdateOrgGatewayTemplate) | **Put** /api/v1/orgs/{org_id}/gatewaytemplates/{gatewaytemplate_id} | updateOrgGatewayTemplate



## CreateOrgGatewayTemplate

> GatewayTemplate CreateOrgGatewayTemplate(ctx, orgId).GatewayTemplate(gatewayTemplate).Execute()

createOrgGatewayTemplate



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
	gatewayTemplate := *openapiclient.NewGatewayTemplate("gw_template") // GatewayTemplate | Gateway Template (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGatewayTemplatesAPI.CreateOrgGatewayTemplate(context.Background(), orgId).GatewayTemplate(gatewayTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGatewayTemplatesAPI.CreateOrgGatewayTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgGatewayTemplate`: GatewayTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsGatewayTemplatesAPI.CreateOrgGatewayTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgGatewayTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayTemplate** | [**GatewayTemplate**](GatewayTemplate.md) | Gateway Template | 

### Return type

[**GatewayTemplate**](GatewayTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgGatewayTemplate

> DeleteOrgGatewayTemplate(ctx, orgId, gatewaytemplateId).Execute()

deleteOrgGatewayTemplate



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
	gatewaytemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsGatewayTemplatesAPI.DeleteOrgGatewayTemplate(context.Background(), orgId, gatewaytemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGatewayTemplatesAPI.DeleteOrgGatewayTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**gatewaytemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgGatewayTemplateRequest struct via the builder pattern


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


## GetOrgGatewayTemplate

> GatewayTemplate GetOrgGatewayTemplate(ctx, orgId, gatewaytemplateId).Execute()

getOrgGatewayTemplate



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
	gatewaytemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGatewayTemplatesAPI.GetOrgGatewayTemplate(context.Background(), orgId, gatewaytemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGatewayTemplatesAPI.GetOrgGatewayTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgGatewayTemplate`: GatewayTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsGatewayTemplatesAPI.GetOrgGatewayTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**gatewaytemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgGatewayTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayTemplate**](GatewayTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgGatewayTemplates

> []GatewayTemplate ListOrgGatewayTemplates(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgGatewayTemplates



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
	resp, r, err := apiClient.OrgsGatewayTemplatesAPI.ListOrgGatewayTemplates(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGatewayTemplatesAPI.ListOrgGatewayTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgGatewayTemplates`: []GatewayTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsGatewayTemplatesAPI.ListOrgGatewayTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgGatewayTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]GatewayTemplate**](GatewayTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgGatewayTemplate

> GatewayTemplate UpdateOrgGatewayTemplate(ctx, orgId, gatewaytemplateId).GatewayTemplate(gatewayTemplate).Execute()

updateOrgGatewayTemplate



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
	gatewaytemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	gatewayTemplate := *openapiclient.NewGatewayTemplate("gw_template") // GatewayTemplate | Gateway Template (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGatewayTemplatesAPI.UpdateOrgGatewayTemplate(context.Background(), orgId, gatewaytemplateId).GatewayTemplate(gatewayTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGatewayTemplatesAPI.UpdateOrgGatewayTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgGatewayTemplate`: GatewayTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsGatewayTemplatesAPI.UpdateOrgGatewayTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**gatewaytemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgGatewayTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gatewayTemplate** | [**GatewayTemplate**](GatewayTemplate.md) | Gateway Template | 

### Return type

[**GatewayTemplate**](GatewayTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

