# \ConstantsModelsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDeviceModels**](ConstantsModelsAPI.md#ListDeviceModels) | **Get** /api/v1/const/device_models | listDeviceModels
[**ListMxEdgeModels**](ConstantsModelsAPI.md#ListMxEdgeModels) | **Get** /api/v1/const/mxedge_models | listMxEdgeModels
[**ListSupportedOtherDeviceModels**](ConstantsModelsAPI.md#ListSupportedOtherDeviceModels) | **Get** /api/v1/const/otherdevice_models | listSupportedOtherDeviceModels



## ListDeviceModels

> []ConstDeviceModel ListDeviceModels(ctx).Execute()

listDeviceModels



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
	resp, r, err := apiClient.ConstantsModelsAPI.ListDeviceModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsModelsAPI.ListDeviceModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeviceModels`: []ConstDeviceModel
	fmt.Fprintf(os.Stdout, "Response from `ConstantsModelsAPI.ListDeviceModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceModelsRequest struct via the builder pattern


### Return type

[**[]ConstDeviceModel**](ConstDeviceModel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMxEdgeModels

> []ConstMxedgeModel ListMxEdgeModels(ctx).Execute()

listMxEdgeModels



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
	resp, r, err := apiClient.ConstantsModelsAPI.ListMxEdgeModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsModelsAPI.ListMxEdgeModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMxEdgeModels`: []ConstMxedgeModel
	fmt.Fprintf(os.Stdout, "Response from `ConstantsModelsAPI.ListMxEdgeModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMxEdgeModelsRequest struct via the builder pattern


### Return type

[**[]ConstMxedgeModel**](ConstMxedgeModel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportedOtherDeviceModels

> []ConstOtherDeviceModel ListSupportedOtherDeviceModels(ctx).Execute()

listSupportedOtherDeviceModels



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
	resp, r, err := apiClient.ConstantsModelsAPI.ListSupportedOtherDeviceModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsModelsAPI.ListSupportedOtherDeviceModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSupportedOtherDeviceModels`: []ConstOtherDeviceModel
	fmt.Fprintf(os.Stdout, "Response from `ConstantsModelsAPI.ListSupportedOtherDeviceModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSupportedOtherDeviceModelsRequest struct via the builder pattern


### Return type

[**[]ConstOtherDeviceModel**](ConstOtherDeviceModel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

