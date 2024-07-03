# \ConstantsEventsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListClientEventsDefinitions**](ConstantsEventsAPI.md#ListClientEventsDefinitions) | **Get** /api/v1/const/client_events | listClientEventsDefinitions
[**ListDeviceEventsDefinitions**](ConstantsEventsAPI.md#ListDeviceEventsDefinitions) | **Get** /api/v1/const/device_events | listDeviceEventsDefinitions
[**ListMxEdgeEventsDefinitions**](ConstantsEventsAPI.md#ListMxEdgeEventsDefinitions) | **Get** /api/v1/const/mxedge_events | listMxEdgeEventsDefinitions
[**ListNacEventsDefinitions**](ConstantsEventsAPI.md#ListNacEventsDefinitions) | **Get** /api/v1/const/nac_events | listNacEventsDefinitions
[**ListOtherDeviceEventsDefinitions**](ConstantsEventsAPI.md#ListOtherDeviceEventsDefinitions) | **Get** /api/v1/const/otherdevice_events | listOtherDeviceEventsDefinitions
[**ListSystemEventsDefinitions**](ConstantsEventsAPI.md#ListSystemEventsDefinitions) | **Get** /api/v1/const/system_events | listSystemEventsDefinitions



## ListClientEventsDefinitions

> []ConstEvent ListClientEventsDefinitions(ctx).Execute()

listClientEventsDefinitions



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
	resp, r, err := apiClient.ConstantsEventsAPI.ListClientEventsDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsEventsAPI.ListClientEventsDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListClientEventsDefinitions`: []ConstEvent
	fmt.Fprintf(os.Stdout, "Response from `ConstantsEventsAPI.ListClientEventsDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListClientEventsDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstEvent**](ConstEvent.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceEventsDefinitions

> []ConstEvent ListDeviceEventsDefinitions(ctx).Execute()

listDeviceEventsDefinitions



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
	resp, r, err := apiClient.ConstantsEventsAPI.ListDeviceEventsDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsEventsAPI.ListDeviceEventsDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDeviceEventsDefinitions`: []ConstEvent
	fmt.Fprintf(os.Stdout, "Response from `ConstantsEventsAPI.ListDeviceEventsDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceEventsDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstEvent**](ConstEvent.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMxEdgeEventsDefinitions

> []ConstEvent ListMxEdgeEventsDefinitions(ctx).Execute()

listMxEdgeEventsDefinitions



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
	resp, r, err := apiClient.ConstantsEventsAPI.ListMxEdgeEventsDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsEventsAPI.ListMxEdgeEventsDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMxEdgeEventsDefinitions`: []ConstEvent
	fmt.Fprintf(os.Stdout, "Response from `ConstantsEventsAPI.ListMxEdgeEventsDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMxEdgeEventsDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstEvent**](ConstEvent.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNacEventsDefinitions

> []ConstNacEvent ListNacEventsDefinitions(ctx).Execute()

listNacEventsDefinitions



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
	resp, r, err := apiClient.ConstantsEventsAPI.ListNacEventsDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsEventsAPI.ListNacEventsDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNacEventsDefinitions`: []ConstNacEvent
	fmt.Fprintf(os.Stdout, "Response from `ConstantsEventsAPI.ListNacEventsDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNacEventsDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstNacEvent**](ConstNacEvent.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOtherDeviceEventsDefinitions

> []ConstEvent ListOtherDeviceEventsDefinitions(ctx).Execute()

listOtherDeviceEventsDefinitions



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
	resp, r, err := apiClient.ConstantsEventsAPI.ListOtherDeviceEventsDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsEventsAPI.ListOtherDeviceEventsDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOtherDeviceEventsDefinitions`: []ConstEvent
	fmt.Fprintf(os.Stdout, "Response from `ConstantsEventsAPI.ListOtherDeviceEventsDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOtherDeviceEventsDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstEvent**](ConstEvent.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemEventsDefinitions

> []ConstEvent ListSystemEventsDefinitions(ctx).Execute()

listSystemEventsDefinitions



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
	resp, r, err := apiClient.ConstantsEventsAPI.ListSystemEventsDefinitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsEventsAPI.ListSystemEventsDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSystemEventsDefinitions`: []ConstEvent
	fmt.Fprintf(os.Stdout, "Response from `ConstantsEventsAPI.ListSystemEventsDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemEventsDefinitionsRequest struct via the builder pattern


### Return type

[**[]ConstEvent**](ConstEvent.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

