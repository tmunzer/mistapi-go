# \UtilitiesPCAPsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgCapturingStatus**](UtilitiesPCAPsAPI.md#GetOrgCapturingStatus) | **Get** /api/v1/orgs/{org_id}/pcaps/capture | getOrgCapturingStatus
[**GetSiteCapturingStatus**](UtilitiesPCAPsAPI.md#GetSiteCapturingStatus) | **Get** /api/v1/sites/{site_id}/pcaps/capture | getSiteCapturingStatus
[**ListOrgPacketCaptures**](UtilitiesPCAPsAPI.md#ListOrgPacketCaptures) | **Get** /api/v1/orgs/{org_id}/pcaps | listOrgPacketCaptures
[**ListSitePacketCaptures**](UtilitiesPCAPsAPI.md#ListSitePacketCaptures) | **Get** /api/v1/sites/{site_id}/pcaps | listSitePacketCaptures
[**StartOrgPacketCapture**](UtilitiesPCAPsAPI.md#StartOrgPacketCapture) | **Post** /api/v1/orgs/{org_id}/pcaps/capture | startOrgPacketCapture
[**StartSitePacketCapture**](UtilitiesPCAPsAPI.md#StartSitePacketCapture) | **Post** /api/v1/sites/{site_id}/pcaps/capture | startSitePacketCapture
[**StopOrgPacketCapture**](UtilitiesPCAPsAPI.md#StopOrgPacketCapture) | **Delete** /api/v1/orgs/{org_id}/pcaps/capture | stopOrgPacketCapture
[**StopSitePacketCapture**](UtilitiesPCAPsAPI.md#StopSitePacketCapture) | **Delete** /api/v1/sites/{site_id}/pcaps/capture | stopSitePacketCapture
[**UpdateSitePacketCapture**](UtilitiesPCAPsAPI.md#UpdateSitePacketCapture) | **Put** /api/v1/sites/{site_id}/pcaps/{pcap_id} | updateSitePacketCapture



## GetOrgCapturingStatus

> ResponsePcapStatus GetOrgCapturingStatus(ctx, orgId).Execute()

getOrgCapturingStatus



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
	resp, r, err := apiClient.UtilitiesPCAPsAPI.GetOrgCapturingStatus(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.GetOrgCapturingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCapturingStatus`: ResponsePcapStatus
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesPCAPsAPI.GetOrgCapturingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCapturingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsePcapStatus**](ResponsePcapStatus.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCapturingStatus

> ResponsePcapStatus GetSiteCapturingStatus(ctx, siteId).Execute()

getSiteCapturingStatus



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
	resp, r, err := apiClient.UtilitiesPCAPsAPI.GetSiteCapturingStatus(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.GetSiteCapturingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteCapturingStatus`: ResponsePcapStatus
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesPCAPsAPI.GetSiteCapturingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteCapturingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponsePcapStatus**](ResponsePcapStatus.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgPacketCaptures

> ResponsePcapSearch ListOrgPacketCaptures(ctx, orgId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

listOrgPacketCaptures



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
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesPCAPsAPI.ListOrgPacketCaptures(context.Background(), orgId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.ListOrgPacketCaptures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgPacketCaptures`: ResponsePcapSearch
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesPCAPsAPI.ListOrgPacketCaptures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgPacketCapturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponsePcapSearch**](ResponsePcapSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSitePacketCaptures

> ResponsePcapSearch ListSitePacketCaptures(ctx, siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).ClientMac(clientMac).Execute()

listSitePacketCaptures



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	clientMac := "clientMac_example" // string | optional client mac filter (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesPCAPsAPI.ListSitePacketCaptures(context.Background(), siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).ClientMac(clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.ListSitePacketCaptures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSitePacketCaptures`: ResponsePcapSearch
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesPCAPsAPI.ListSitePacketCaptures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSitePacketCapturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **clientMac** | **string** | optional client mac filter | 

### Return type

[**ResponsePcapSearch**](ResponsePcapSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartOrgPacketCapture

> ResponsePcapStart StartOrgPacketCapture(ctx, orgId).CaptureOrg(captureOrg).Execute()

startOrgPacketCapture



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
	captureOrg := openapiclient.capture_org{CaptureMxedge: openapiclient.NewCaptureMxedge(openapiclient.capture_mxedge_type(""))} // CaptureOrg | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesPCAPsAPI.StartOrgPacketCapture(context.Background(), orgId).CaptureOrg(captureOrg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.StartOrgPacketCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartOrgPacketCapture`: ResponsePcapStart
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesPCAPsAPI.StartOrgPacketCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartOrgPacketCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captureOrg** | [**CaptureOrg**](CaptureOrg.md) | Request Body | 

### Return type

[**ResponsePcapStart**](ResponsePcapStart.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSitePacketCapture

> ResponsePcapStart StartSitePacketCapture(ctx, siteId).CaptureSite(captureSite).Execute()

startSitePacketCapture



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
	captureSite := openapiclient.capture_site{CaptureClient: openapiclient.NewCaptureClient(openapiclient.capture_client_type(""))} // CaptureSite | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesPCAPsAPI.StartSitePacketCapture(context.Background(), siteId).CaptureSite(captureSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.StartSitePacketCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartSitePacketCapture`: ResponsePcapStart
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesPCAPsAPI.StartSitePacketCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSitePacketCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captureSite** | [**CaptureSite**](CaptureSite.md) | Request Body | 

### Return type

[**ResponsePcapStart**](ResponsePcapStart.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopOrgPacketCapture

> StopOrgPacketCapture(ctx, orgId).Execute()

stopOrgPacketCapture



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
	r, err := apiClient.UtilitiesPCAPsAPI.StopOrgPacketCapture(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.StopOrgPacketCapture``: %v\n", err)
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

Other parameters are passed through a pointer to a apiStopOrgPacketCaptureRequest struct via the builder pattern


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


## StopSitePacketCapture

> StopSitePacketCapture(ctx, siteId).Execute()

stopSitePacketCapture



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
	r, err := apiClient.UtilitiesPCAPsAPI.StopSitePacketCapture(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.StopSitePacketCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSitePacketCaptureRequest struct via the builder pattern


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


## UpdateSitePacketCapture

> UpdateSitePacketCapture(ctx, siteId, pcapId).NotesString(notesString).Execute()

updateSitePacketCapture



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
	pcapId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	notesString := *openapiclient.NewNotesString() // NotesString |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesPCAPsAPI.UpdateSitePacketCapture(context.Background(), siteId, pcapId).NotesString(notesString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesPCAPsAPI.UpdateSitePacketCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**pcapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSitePacketCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **notesString** | [**NotesString**](NotesString.md) |  | 

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

