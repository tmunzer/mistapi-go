# \OrgsRFTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgRfTemplate**](OrgsRFTemplatesAPI.md#CreateOrgRfTemplate) | **Post** /api/v1/orgs/{org_id}/rftemplates | createOrgRfTemplate
[**DeleteOrgRfTemplate**](OrgsRFTemplatesAPI.md#DeleteOrgRfTemplate) | **Delete** /api/v1/orgs/{org_id}/rftemplates/{rftemplate_id} | deleteOrgRfTemplate
[**GetOrgRfTemplate**](OrgsRFTemplatesAPI.md#GetOrgRfTemplate) | **Get** /api/v1/orgs/{org_id}/rftemplates/{rftemplate_id} | getOrgRfTemplate
[**ListOrgRfTemplates**](OrgsRFTemplatesAPI.md#ListOrgRfTemplates) | **Get** /api/v1/orgs/{org_id}/rftemplates | listOrgRfTemplates
[**UpdateOrgRfTemplate**](OrgsRFTemplatesAPI.md#UpdateOrgRfTemplate) | **Put** /api/v1/orgs/{org_id}/rftemplates/{rftemplate_id} | updateOrgRfTemplate



## CreateOrgRfTemplate

> RfTemplate CreateOrgRfTemplate(ctx, orgId).RfTemplate(rfTemplate).Execute()

createOrgRfTemplate



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
	rfTemplate := *openapiclient.NewRfTemplate("Name_example") // RfTemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsRFTemplatesAPI.CreateOrgRfTemplate(context.Background(), orgId).RfTemplate(rfTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsRFTemplatesAPI.CreateOrgRfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgRfTemplate`: RfTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsRFTemplatesAPI.CreateOrgRfTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgRfTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rfTemplate** | [**RfTemplate**](RfTemplate.md) | Request Body | 

### Return type

[**RfTemplate**](RfTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgRfTemplate

> DeleteOrgRfTemplate(ctx, orgId, rftemplateId).Execute()

deleteOrgRfTemplate



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
	rftemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsRFTemplatesAPI.DeleteOrgRfTemplate(context.Background(), orgId, rftemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsRFTemplatesAPI.DeleteOrgRfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**rftemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgRfTemplateRequest struct via the builder pattern


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


## GetOrgRfTemplate

> RfTemplate GetOrgRfTemplate(ctx, orgId, rftemplateId).Execute()

getOrgRfTemplate



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
	rftemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsRFTemplatesAPI.GetOrgRfTemplate(context.Background(), orgId, rftemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsRFTemplatesAPI.GetOrgRfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgRfTemplate`: RfTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsRFTemplatesAPI.GetOrgRfTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**rftemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRfTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RfTemplate**](RfTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgRfTemplates

> []RfTemplate ListOrgRfTemplates(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgRfTemplates



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
	resp, r, err := apiClient.OrgsRFTemplatesAPI.ListOrgRfTemplates(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsRFTemplatesAPI.ListOrgRfTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgRfTemplates`: []RfTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsRFTemplatesAPI.ListOrgRfTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgRfTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]RfTemplate**](RfTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgRfTemplate

> RfTemplate UpdateOrgRfTemplate(ctx, orgId, rftemplateId).RfTemplate(rfTemplate).Execute()

updateOrgRfTemplate



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
	rftemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfTemplate := *openapiclient.NewRfTemplate("Name_example") // RfTemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsRFTemplatesAPI.UpdateOrgRfTemplate(context.Background(), orgId, rftemplateId).RfTemplate(rfTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsRFTemplatesAPI.UpdateOrgRfTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgRfTemplate`: RfTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsRFTemplatesAPI.UpdateOrgRfTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**rftemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRfTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rfTemplate** | [**RfTemplate**](RfTemplate.md) | Request Body | 

### Return type

[**RfTemplate**](RfTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

