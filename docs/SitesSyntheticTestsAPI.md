# \SitesSyntheticTestsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteDeviceSyntheticTest**](SitesSyntheticTestsAPI.md#GetSiteDeviceSyntheticTest) | **Get** /api/v1/sites/{site_id}/devices/{device_id}/synthetic_test | getSiteDeviceSyntheticTest
[**GetSiteSyntheticTestStatus**](SitesSyntheticTestsAPI.md#GetSiteSyntheticTestStatus) | **Get** /api/v1/sites/{site_id}/synthetic_test | getSiteSyntheticTestStatus
[**SearchSiteSyntheticTest**](SitesSyntheticTestsAPI.md#SearchSiteSyntheticTest) | **Get** /api/v1/sites/{site_id}/synthetic_test/search | searchSiteSyntheticTest
[**StartSiteSwitchRadiusSyntheticTest**](SitesSyntheticTestsAPI.md#StartSiteSwitchRadiusSyntheticTest) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/check_radius_server | triggerSiteSwitchRadiusSyntheticTest
[**TriggerSiteDeviceSyntheticTest**](SitesSyntheticTestsAPI.md#TriggerSiteDeviceSyntheticTest) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/synthetic_test | triggerSiteDeviceSyntheticTest
[**TriggerSiteSyntheticTest**](SitesSyntheticTestsAPI.md#TriggerSiteSyntheticTest) | **Post** /api/v1/sites/{site_id}/synthetic_test | triggerSiteSyntheticTest



## GetSiteDeviceSyntheticTest

> GetSiteDeviceSyntheticTest(ctx, siteId, deviceId).Execute()

getSiteDeviceSyntheticTest



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesSyntheticTestsAPI.GetSiteDeviceSyntheticTest(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSyntheticTestsAPI.GetSiteDeviceSyntheticTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceSyntheticTestRequest struct via the builder pattern


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


## GetSiteSyntheticTestStatus

> SyntheticTestInfo GetSiteSyntheticTestStatus(ctx, siteId).Execute()

getSiteSyntheticTestStatus



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSyntheticTestsAPI.GetSiteSyntheticTestStatus(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSyntheticTestsAPI.GetSiteSyntheticTestStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSyntheticTestStatus`: SyntheticTestInfo
	fmt.Fprintf(os.Stdout, "Response from `SitesSyntheticTestsAPI.GetSiteSyntheticTestStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSyntheticTestStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticTestInfo**](SyntheticTestInfo.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSiteSyntheticTest

> ReponseSyntheticTestSearch SearchSiteSyntheticTest(ctx, siteId).Mac(mac).PortId(portId).VlanId(vlanId).By(by).Reason(reason).Type_(type_).Execute()

searchSiteSyntheticTest



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mac := "mac_example" // string | Device MAC Address (optional)
	portId := "portId_example" // string | port_id used to run the test (for SSR only) (optional)
	vlanId := "vlanId_example" // string | VLAN ID (optional)
	by := "by_example" // string | entity who triggers the test (optional)
	reason := "reason_example" // string | test failure reason (optional)
	type_ := openapiclient.synthetic_test_type("") // SyntheticTestType | synthetic test type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSyntheticTestsAPI.SearchSiteSyntheticTest(context.Background(), siteId).Mac(mac).PortId(portId).VlanId(vlanId).By(by).Reason(reason).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSyntheticTestsAPI.SearchSiteSyntheticTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteSyntheticTest`: ReponseSyntheticTestSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesSyntheticTestsAPI.SearchSiteSyntheticTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteSyntheticTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | Device MAC Address | 
 **portId** | **string** | port_id used to run the test (for SSR only) | 
 **vlanId** | **string** | VLAN ID | 
 **by** | **string** | entity who triggers the test | 
 **reason** | **string** | test failure reason | 
 **type_** | [**SyntheticTestType**](SyntheticTestType.md) | synthetic test type | 

### Return type

[**ReponseSyntheticTestSearch**](ReponseSyntheticTestSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSiteSwitchRadiusSyntheticTest

> WebsocketSession StartSiteSwitchRadiusSyntheticTest(ctx, siteId, deviceId).SyntheticTestRadiusServer(syntheticTestRadiusServer).Execute()

triggerSiteSwitchRadiusSyntheticTest



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	syntheticTestRadiusServer := *openapiclient.NewSyntheticTestRadiusServer("Password_example", "User_example") // SyntheticTestRadiusServer |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSyntheticTestsAPI.StartSiteSwitchRadiusSyntheticTest(context.Background(), siteId, deviceId).SyntheticTestRadiusServer(syntheticTestRadiusServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSyntheticTestsAPI.StartSiteSwitchRadiusSyntheticTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartSiteSwitchRadiusSyntheticTest`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `SitesSyntheticTestsAPI.StartSiteSwitchRadiusSyntheticTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSiteSwitchRadiusSyntheticTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syntheticTestRadiusServer** | [**SyntheticTestRadiusServer**](SyntheticTestRadiusServer.md) |  | 

### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerSiteDeviceSyntheticTest

> TriggerSiteDeviceSyntheticTest(ctx, siteId, deviceId).SyntheticTestDevice(syntheticTestDevice).Execute()

triggerSiteDeviceSyntheticTest



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	syntheticTestDevice := *openapiclient.NewSyntheticTestDevice(openapiclient.synthetic_test_type("")) // SyntheticTestDevice |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesSyntheticTestsAPI.TriggerSiteDeviceSyntheticTest(context.Background(), siteId, deviceId).SyntheticTestDevice(syntheticTestDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSyntheticTestsAPI.TriggerSiteDeviceSyntheticTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSiteDeviceSyntheticTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **syntheticTestDevice** | [**SyntheticTestDevice**](SyntheticTestDevice.md) |  | 

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


## TriggerSiteSyntheticTest

> ReponseSyntheticTest TriggerSiteSyntheticTest(ctx, siteId).SyntheticTest(syntheticTest).Execute()

triggerSiteSyntheticTest



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	syntheticTest := *openapiclient.NewSyntheticTest() // SyntheticTest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSyntheticTestsAPI.TriggerSiteSyntheticTest(context.Background(), siteId).SyntheticTest(syntheticTest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSyntheticTestsAPI.TriggerSiteSyntheticTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TriggerSiteSyntheticTest`: ReponseSyntheticTest
	fmt.Fprintf(os.Stdout, "Response from `SitesSyntheticTestsAPI.TriggerSiteSyntheticTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerSiteSyntheticTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticTest** | [**SyntheticTest**](SyntheticTest.md) |  | 

### Return type

[**ReponseSyntheticTest**](ReponseSyntheticTest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

