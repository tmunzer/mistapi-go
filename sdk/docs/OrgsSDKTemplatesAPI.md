# \OrgsSDKTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSdkTemplate**](OrgsSDKTemplatesAPI.md#CreateSdkTemplate) | **Post** /api/v1/orgs/{org_id}/sdktemplates | createSdkTemplate
[**DeleteSdkTemplate**](OrgsSDKTemplatesAPI.md#DeleteSdkTemplate) | **Delete** /api/v1/orgs/{org_id}/sdktemplates/{sdktemplate_id} | deleteSdkTemplate
[**GetSdkTemplate**](OrgsSDKTemplatesAPI.md#GetSdkTemplate) | **Get** /api/v1/orgs/{org_id}/sdktemplates/{sdktemplate_id} | getSdkTemplate
[**ListSdkTemplates**](OrgsSDKTemplatesAPI.md#ListSdkTemplates) | **Get** /api/v1/orgs/{org_id}/sdktemplates | listSdkTemplates
[**UpdateSdkTemplate**](OrgsSDKTemplatesAPI.md#UpdateSdkTemplate) | **Put** /api/v1/orgs/{org_id}/sdktemplates/{sdktemplate_id} | updateSdkTemplate



## CreateSdkTemplate

> Sdktemplate CreateSdkTemplate(ctx, orgId).Sdktemplate(sdktemplate).Execute()

createSdkTemplate



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
	sdktemplate := *openapiclient.NewSdktemplate("Name_example") // Sdktemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKTemplatesAPI.CreateSdkTemplate(context.Background(), orgId).Sdktemplate(sdktemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKTemplatesAPI.CreateSdkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSdkTemplate`: Sdktemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKTemplatesAPI.CreateSdkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdktemplate** | [**Sdktemplate**](Sdktemplate.md) | Request Body | 

### Return type

[**Sdktemplate**](Sdktemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSdkTemplate

> DeleteSdkTemplate(ctx, orgId, sdktemplateId).Execute()

deleteSdkTemplate



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
	sdktemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSDKTemplatesAPI.DeleteSdkTemplate(context.Background(), orgId, sdktemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKTemplatesAPI.DeleteSdkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdktemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSdkTemplateRequest struct via the builder pattern


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


## GetSdkTemplate

> Sdktemplate GetSdkTemplate(ctx, orgId, sdktemplateId).Execute()

getSdkTemplate



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
	sdktemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKTemplatesAPI.GetSdkTemplate(context.Background(), orgId, sdktemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKTemplatesAPI.GetSdkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSdkTemplate`: Sdktemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKTemplatesAPI.GetSdkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdktemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sdktemplate**](Sdktemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSdkTemplates

> []Sdktemplate ListSdkTemplates(ctx, orgId).Execute()

listSdkTemplates



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
	resp, r, err := apiClient.OrgsSDKTemplatesAPI.ListSdkTemplates(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKTemplatesAPI.ListSdkTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSdkTemplates`: []Sdktemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKTemplatesAPI.ListSdkTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSdkTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Sdktemplate**](Sdktemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSdkTemplate

> Sdktemplate UpdateSdkTemplate(ctx, orgId, sdktemplateId).Sdktemplate(sdktemplate).Execute()

updateSdkTemplate



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
	sdktemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sdktemplate := *openapiclient.NewSdktemplate("Name_example") // Sdktemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKTemplatesAPI.UpdateSdkTemplate(context.Background(), orgId, sdktemplateId).Sdktemplate(sdktemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKTemplatesAPI.UpdateSdkTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSdkTemplate`: Sdktemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKTemplatesAPI.UpdateSdkTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdktemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSdkTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sdktemplate** | [**Sdktemplate**](Sdktemplate.md) | Request Body | 

### Return type

[**Sdktemplate**](Sdktemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

