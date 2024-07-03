# \OrgsAlarmTemplatesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgAlarmTemplate**](OrgsAlarmTemplatesAPI.md#CreateOrgAlarmTemplate) | **Post** /api/v1/orgs/{org_id}/alarmtemplates | createOrgAlarmTemplate
[**DeleteOrgAlarmTemplate**](OrgsAlarmTemplatesAPI.md#DeleteOrgAlarmTemplate) | **Delete** /api/v1/orgs/{org_id}/alarmtemplates/{alarmtemplate_id} | deleteOrgAlarmTemplate
[**GetOrgAlarmTemplate**](OrgsAlarmTemplatesAPI.md#GetOrgAlarmTemplate) | **Get** /api/v1/orgs/{org_id}/alarmtemplates/{alarmtemplate_id} | getOrgAlarmTemplate
[**ListOrgAlarmTemplates**](OrgsAlarmTemplatesAPI.md#ListOrgAlarmTemplates) | **Get** /api/v1/orgs/{org_id}/alarmtemplates | listOrgAlarmTemplates
[**ListOrgSuppressedAlarms**](OrgsAlarmTemplatesAPI.md#ListOrgSuppressedAlarms) | **Get** /api/v1/orgs/{org_id}/alarmtemplates/suppress | listOrgSuppressedAlarms
[**SuppressOrgAlarm**](OrgsAlarmTemplatesAPI.md#SuppressOrgAlarm) | **Post** /api/v1/orgs/{org_id}/alarmtemplates/suppress | suppressOrgAlarm
[**UnsuppressOrgSuppressedAlarms**](OrgsAlarmTemplatesAPI.md#UnsuppressOrgSuppressedAlarms) | **Delete** /api/v1/orgs/{org_id}/alarmtemplates/suppress | unsuppressOrgSuppressedAlarms
[**UpdateOrgAlarmTemplate**](OrgsAlarmTemplatesAPI.md#UpdateOrgAlarmTemplate) | **Put** /api/v1/orgs/{org_id}/alarmtemplates/{alarmtemplate_id} | updateOrgAlarmTemplate



## CreateOrgAlarmTemplate

> AlarmTemplate CreateOrgAlarmTemplate(ctx, orgId).AlarmTemplate(alarmTemplate).Execute()

createOrgAlarmTemplate



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
	alarmTemplate := *openapiclient.NewAlarmTemplate(*openapiclient.NewDelivery(true), map[string]AlarmTemplateRule{"key": *openapiclient.NewAlarmTemplateRule()}) // AlarmTemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAlarmTemplatesAPI.CreateOrgAlarmTemplate(context.Background(), orgId).AlarmTemplate(alarmTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.CreateOrgAlarmTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgAlarmTemplate`: AlarmTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAlarmTemplatesAPI.CreateOrgAlarmTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgAlarmTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alarmTemplate** | [**AlarmTemplate**](AlarmTemplate.md) | Request Body | 

### Return type

[**AlarmTemplate**](AlarmTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgAlarmTemplate

> DeleteOrgAlarmTemplate(ctx, orgId, alarmtemplateId).Execute()

deleteOrgAlarmTemplate



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
	alarmtemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmTemplatesAPI.DeleteOrgAlarmTemplate(context.Background(), orgId, alarmtemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.DeleteOrgAlarmTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**alarmtemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgAlarmTemplateRequest struct via the builder pattern


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


## GetOrgAlarmTemplate

> AlarmTemplate GetOrgAlarmTemplate(ctx, orgId, alarmtemplateId).Execute()

getOrgAlarmTemplate



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
	alarmtemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAlarmTemplatesAPI.GetOrgAlarmTemplate(context.Background(), orgId, alarmtemplateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.GetOrgAlarmTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgAlarmTemplate`: AlarmTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAlarmTemplatesAPI.GetOrgAlarmTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**alarmtemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAlarmTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlarmTemplate**](AlarmTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgAlarmTemplates

> []AlarmTemplate ListOrgAlarmTemplates(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgAlarmTemplates



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
	resp, r, err := apiClient.OrgsAlarmTemplatesAPI.ListOrgAlarmTemplates(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.ListOrgAlarmTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAlarmTemplates`: []AlarmTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAlarmTemplatesAPI.ListOrgAlarmTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAlarmTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]AlarmTemplate**](AlarmTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSuppressedAlarms

> ResponseOrgSuppressAlarm ListOrgSuppressedAlarms(ctx, orgId).Scope(scope).Execute()

listOrgSuppressedAlarms



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
	scope := openapiclient.suppressed_alarm_scope("") // SuppressedAlarmScope |  (optional) (default to "site")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAlarmTemplatesAPI.ListOrgSuppressedAlarms(context.Background(), orgId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.ListOrgSuppressedAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSuppressedAlarms`: ResponseOrgSuppressAlarm
	fmt.Fprintf(os.Stdout, "Response from `OrgsAlarmTemplatesAPI.ListOrgSuppressedAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSuppressedAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**SuppressedAlarmScope**](SuppressedAlarmScope.md) |  | [default to &quot;site&quot;]

### Return type

[**ResponseOrgSuppressAlarm**](ResponseOrgSuppressAlarm.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuppressOrgAlarm

> SuppressOrgAlarm(ctx, orgId).SuppressedAlarm(suppressedAlarm).Execute()

suppressOrgAlarm



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
	suppressedAlarm := *openapiclient.NewSuppressedAlarm() // SuppressedAlarm |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmTemplatesAPI.SuppressOrgAlarm(context.Background(), orgId).SuppressedAlarm(suppressedAlarm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.SuppressOrgAlarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuppressOrgAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressedAlarm** | [**SuppressedAlarm**](SuppressedAlarm.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsuppressOrgSuppressedAlarms

> UnsuppressOrgSuppressedAlarms(ctx, orgId).Execute()

unsuppressOrgSuppressedAlarms



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmTemplatesAPI.UnsuppressOrgSuppressedAlarms(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.UnsuppressOrgSuppressedAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsuppressOrgSuppressedAlarmsRequest struct via the builder pattern


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


## UpdateOrgAlarmTemplate

> AlarmTemplate UpdateOrgAlarmTemplate(ctx, orgId, alarmtemplateId).AlarmTemplate(alarmTemplate).Execute()

updateOrgAlarmTemplate



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
	alarmtemplateId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	alarmTemplate := *openapiclient.NewAlarmTemplate(*openapiclient.NewDelivery(true), map[string]AlarmTemplateRule{"key": *openapiclient.NewAlarmTemplateRule()}) // AlarmTemplate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAlarmTemplatesAPI.UpdateOrgAlarmTemplate(context.Background(), orgId, alarmtemplateId).AlarmTemplate(alarmTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmTemplatesAPI.UpdateOrgAlarmTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgAlarmTemplate`: AlarmTemplate
	fmt.Fprintf(os.Stdout, "Response from `OrgsAlarmTemplatesAPI.UpdateOrgAlarmTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**alarmtemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgAlarmTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alarmTemplate** | [**AlarmTemplate**](AlarmTemplate.md) | Request Body | 

### Return type

[**AlarmTemplate**](AlarmTemplate.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

