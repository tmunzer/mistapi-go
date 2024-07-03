# \UtilitiesWANAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearSiteSsrArpCache**](UtilitiesWANAPI.md#ClearSiteSsrArpCache) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_arp | clearSiteSsrArpCache
[**ClearSiteSsrBgpRoutes**](UtilitiesWANAPI.md#ClearSiteSsrBgpRoutes) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/clear_bgp | clearSiteBgpRoutes
[**GetSiteSsrAndSrxRoutes**](UtilitiesWANAPI.md#GetSiteSsrAndSrxRoutes) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_route | getSiteDeviceRoutes
[**GetSiteSsrAndSrxSessions**](UtilitiesWANAPI.md#GetSiteSsrAndSrxSessions) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_session | getSiteSsrAndSrxSessions
[**GetSiteSsrServicePath**](UtilitiesWANAPI.md#GetSiteSsrServicePath) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/show_service_path | getSiteSsrServicePath
[**ReleaseSiteSsrDhcpLease**](UtilitiesWANAPI.md#ReleaseSiteSsrDhcpLease) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/release_dhcp | releaseSiteSsrDhcpLease
[**ServicePingFromSsr**](UtilitiesWANAPI.md#ServicePingFromSsr) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/service_ping | servicePingFromSsr
[**TestSiteSsrDnsResolution**](UtilitiesWANAPI.md#TestSiteSsrDnsResolution) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/resolve_dns | testSiteSsrDnsResolution



## ClearSiteSsrArpCache

> WebsocketSession ClearSiteSsrArpCache(ctx, siteId, deviceId).UtilsClearArp(utilsClearArp).Execute()

clearSiteSsrArpCache



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
	utilsClearArp := *openapiclient.NewUtilsClearArp() // UtilsClearArp |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesWANAPI.ClearSiteSsrArpCache(context.Background(), siteId, deviceId).UtilsClearArp(utilsClearArp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.ClearSiteSsrArpCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearSiteSsrArpCache`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesWANAPI.ClearSiteSsrArpCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearSiteSsrArpCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsClearArp** | [**UtilsClearArp**](UtilsClearArp.md) |  | 

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


## ClearSiteSsrBgpRoutes

> WebsocketSession ClearSiteSsrBgpRoutes(ctx, siteId, deviceId).UtilsClearBgp(utilsClearBgp).Execute()

clearSiteBgpRoutes



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
	utilsClearBgp := *openapiclient.NewUtilsClearBgp("Neighbor_example", openapiclient.utils_clear_bgp_type("")) // UtilsClearBgp |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesWANAPI.ClearSiteSsrBgpRoutes(context.Background(), siteId, deviceId).UtilsClearBgp(utilsClearBgp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.ClearSiteSsrBgpRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearSiteSsrBgpRoutes`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesWANAPI.ClearSiteSsrBgpRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearSiteSsrBgpRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsClearBgp** | [**UtilsClearBgp**](UtilsClearBgp.md) |  | 

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


## GetSiteSsrAndSrxRoutes

> WebsocketSession GetSiteSsrAndSrxRoutes(ctx, siteId, deviceId).UtilsShowRoute(utilsShowRoute).Execute()

getSiteDeviceRoutes



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
	utilsShowRoute := *openapiclient.NewUtilsShowRoute() // UtilsShowRoute | all attributes are optional (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesWANAPI.GetSiteSsrAndSrxRoutes(context.Background(), siteId, deviceId).UtilsShowRoute(utilsShowRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.GetSiteSsrAndSrxRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSsrAndSrxRoutes`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesWANAPI.GetSiteSsrAndSrxRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSsrAndSrxRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsShowRoute** | [**UtilsShowRoute**](UtilsShowRoute.md) | all attributes are optional | 

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


## GetSiteSsrAndSrxSessions

> WebsocketSession GetSiteSsrAndSrxSessions(ctx, siteId, deviceId).UtilsShowSession(utilsShowSession).Execute()

getSiteSsrAndSrxSessions



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
	utilsShowSession := *openapiclient.NewUtilsShowSession() // UtilsShowSession |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesWANAPI.GetSiteSsrAndSrxSessions(context.Background(), siteId, deviceId).UtilsShowSession(utilsShowSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.GetSiteSsrAndSrxSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSsrAndSrxSessions`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesWANAPI.GetSiteSsrAndSrxSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSsrAndSrxSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsShowSession** | [**UtilsShowSession**](UtilsShowSession.md) |  | 

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


## GetSiteSsrServicePath

> WebsocketSession GetSiteSsrServicePath(ctx, siteId, deviceId).UtilsShowServicePath(utilsShowServicePath).Execute()

getSiteSsrServicePath



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
	utilsShowServicePath := *openapiclient.NewUtilsShowServicePath() // UtilsShowServicePath |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesWANAPI.GetSiteSsrServicePath(context.Background(), siteId, deviceId).UtilsShowServicePath(utilsShowServicePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.GetSiteSsrServicePath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSsrServicePath`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesWANAPI.GetSiteSsrServicePath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSsrServicePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsShowServicePath** | [**UtilsShowServicePath**](UtilsShowServicePath.md) |  | 

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


## ReleaseSiteSsrDhcpLease

> ReleaseSiteSsrDhcpLease(ctx, siteId, deviceId).UtilsReleaseDhcp(utilsReleaseDhcp).Execute()

releaseSiteSsrDhcpLease



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
	utilsReleaseDhcp := *openapiclient.NewUtilsReleaseDhcp("wan-interface") // UtilsReleaseDhcp |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWANAPI.ReleaseSiteSsrDhcpLease(context.Background(), siteId, deviceId).UtilsReleaseDhcp(utilsReleaseDhcp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.ReleaseSiteSsrDhcpLease``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReleaseSiteSsrDhcpLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsReleaseDhcp** | [**UtilsReleaseDhcp**](UtilsReleaseDhcp.md) |  | 

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


## ServicePingFromSsr

> WebsocketSession ServicePingFromSsr(ctx, siteId, deviceId).UtilsServicePing(utilsServicePing).Execute()

servicePingFromSsr



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
	utilsServicePing := *openapiclient.NewUtilsServicePing("Host_example", "Service_example") // UtilsServicePing | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesWANAPI.ServicePingFromSsr(context.Background(), siteId, deviceId).UtilsServicePing(utilsServicePing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.ServicePingFromSsr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServicePingFromSsr`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesWANAPI.ServicePingFromSsr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiServicePingFromSsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **utilsServicePing** | [**UtilsServicePing**](UtilsServicePing.md) | Request Body | 

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


## TestSiteSsrDnsResolution

> WebsocketSession TestSiteSsrDnsResolution(ctx, siteId, deviceId).Execute()

testSiteSsrDnsResolution



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
	resp, r, err := apiClient.UtilitiesWANAPI.TestSiteSsrDnsResolution(context.Background(), siteId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWANAPI.TestSiteSsrDnsResolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestSiteSsrDnsResolution`: WebsocketSession
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesWANAPI.TestSiteSsrDnsResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSiteSsrDnsResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WebsocketSession**](WebsocketSession.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

