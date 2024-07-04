# \ConstantsDefinitionsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAlarmDefinitions**](ConstantsDefinitionsAPI.md#ListAlarmDefinitions) | **Get** /api/v1/const/alarm_defs | listAlarmDefinitions
[**ListApLedDefinition**](ConstantsDefinitionsAPI.md#ListApLedDefinition) | **Get** /api/v1/const/ap_led_status | listApLedDefinition
[**ListAppCategoryDefinitions**](ConstantsDefinitionsAPI.md#ListAppCategoryDefinitions) | **Get** /api/v1/const/app_categories | listAppCategoryDefinitions
[**ListAppSubCategoryDefinitions**](ConstantsDefinitionsAPI.md#ListAppSubCategoryDefinitions) | **Get** /api/v1/const/app_subcategories | listAppSubCategoryDefinitions



## ListAlarmDefinitions

> []ConstAlarmDefinition ListAlarmDefinitions(ctx).Execute()

listAlarmDefinitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsDefinitionsAPI.ListAlarmDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsDefinitionsAPI.ListAlarmDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlarmDefinitions`: []ConstAlarmDefinition
	fmt.Fprintf(os.Stdout, "Response from `ConstantsDefinitionsAPI.ListAlarmDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAlarmDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstAlarmDefinition**](ConstAlarmDefinition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApLedDefinition

> []ConstApLed ListApLedDefinition(ctx).Execute()

listApLedDefinition



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsDefinitionsAPI.ListApLedDefinition(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsDefinitionsAPI.ListApLedDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApLedDefinition`: []ConstApLed
	fmt.Fprintf(os.Stdout, "Response from `ConstantsDefinitionsAPI.ListApLedDefinition`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApLedDefinitionRequest struct via the builder pattern


### Return type

[**[]ConstApLed**](ConstApLed.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppCategoryDefinitions

> []ConstAppCategoryDefinition ListAppCategoryDefinitions(ctx).Execute()

listAppCategoryDefinitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsDefinitionsAPI.ListAppCategoryDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsDefinitionsAPI.ListAppCategoryDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppCategoryDefinitions`: []ConstAppCategoryDefinition
	fmt.Fprintf(os.Stdout, "Response from `ConstantsDefinitionsAPI.ListAppCategoryDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppCategoryDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstAppCategoryDefinition**](ConstAppCategoryDefinition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAppSubCategoryDefinitions

> []ConstAppSubcategoryDefinition ListAppSubCategoryDefinitions(ctx).Execute()

listAppSubCategoryDefinitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsDefinitionsAPI.ListAppSubCategoryDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsDefinitionsAPI.ListAppSubCategoryDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAppSubCategoryDefinitions`: []ConstAppSubcategoryDefinition
	fmt.Fprintf(os.Stdout, "Response from `ConstantsDefinitionsAPI.ListAppSubCategoryDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAppSubCategoryDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstAppSubcategoryDefinition**](ConstAppSubcategoryDefinition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

