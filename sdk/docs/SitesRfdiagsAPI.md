# \SitesRfdiagsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSiteRfdiagRecording**](SitesRfdiagsAPI.md#DeleteSiteRfdiagRecording) | **Delete** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id} | deleteSiteRfdiagRecording
[**DownloadSiteRfdiagRecording**](SitesRfdiagsAPI.md#DownloadSiteRfdiagRecording) | **Get** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id}/download | downloadSiteRfdiagRecording
[**GetSiteRfdiagRecording**](SitesRfdiagsAPI.md#GetSiteRfdiagRecording) | **Get** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id} | getSiteRfdiagRecording
[**GetSiteSiteRfdiagRecording**](SitesRfdiagsAPI.md#GetSiteSiteRfdiagRecording) | **Get** /api/v1/sites/{site_id}/rfdiags | getSiteSiteRfdiagRecording
[**StartSiteRecording**](SitesRfdiagsAPI.md#StartSiteRecording) | **Post** /api/v1/sites/{site_id}/rfdiags | startSiteRecording
[**StopSiteRfdiagRecording**](SitesRfdiagsAPI.md#StopSiteRfdiagRecording) | **Post** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id}/stop | stopSiteRfdiagRecording
[**UpdateSiteRfdiagRecording**](SitesRfdiagsAPI.md#UpdateSiteRfdiagRecording) | **Put** /api/v1/sites/{site_id}/rfdiags/{rfdiag_id} | updateSiteRfdiagRecording



## DeleteSiteRfdiagRecording

> DeleteSiteRfdiagRecording(ctx, siteId, rfdiagId).Execute()

deleteSiteRfdiagRecording



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfdiagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesRfdiagsAPI.DeleteSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRfdiagsAPI.DeleteSiteRfdiagRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rfdiagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteRfdiagRecordingRequest struct via the builder pattern


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


## DownloadSiteRfdiagRecording

> *os.File DownloadSiteRfdiagRecording(ctx, siteId, rfdiagId).Execute()

downloadSiteRfdiagRecording



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfdiagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRfdiagsAPI.DownloadSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRfdiagsAPI.DownloadSiteRfdiagRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSiteRfdiagRecording`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SitesRfdiagsAPI.DownloadSiteRfdiagRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rfdiagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSiteRfdiagRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteRfdiagRecording

> []RfDiagInfoItem GetSiteRfdiagRecording(ctx, siteId, rfdiagId).Execute()

getSiteRfdiagRecording



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfdiagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRfdiagsAPI.GetSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRfdiagsAPI.GetSiteRfdiagRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteRfdiagRecording`: []RfDiagInfoItem
	fmt.Fprintf(os.Stdout, "Response from `SitesRfdiagsAPI.GetSiteRfdiagRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rfdiagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteRfdiagRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]RfDiagInfoItem**](RfDiagInfoItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSiteRfdiagRecording

> [][]RfDiagInfoItem GetSiteSiteRfdiagRecording(ctx, siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

getSiteSiteRfdiagRecording



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRfdiagsAPI.GetSiteSiteRfdiagRecording(context.Background(), siteId).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRfdiagsAPI.GetSiteSiteRfdiagRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSiteRfdiagRecording`: [][]RfDiagInfoItem
	fmt.Fprintf(os.Stdout, "Response from `SitesRfdiagsAPI.GetSiteSiteRfdiagRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSiteRfdiagRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[][]RfDiagInfoItem**](set.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSiteRecording

> []RfDiagInfoItem StartSiteRecording(ctx, siteId).RfDiag(rfDiag).Execute()

startSiteRecording



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfDiag := *openapiclient.NewRfDiag("Name_example", openapiclient.rf_client_type("")) // RfDiag | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRfdiagsAPI.StartSiteRecording(context.Background(), siteId).RfDiag(rfDiag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRfdiagsAPI.StartSiteRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartSiteRecording`: []RfDiagInfoItem
	fmt.Fprintf(os.Stdout, "Response from `SitesRfdiagsAPI.StartSiteRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSiteRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rfDiag** | [**RfDiag**](RfDiag.md) | Request Body | 

### Return type

[**[]RfDiagInfoItem**](RfDiagInfoItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopSiteRfdiagRecording

> StopSiteRfdiagRecording(ctx, siteId, rfdiagId).Execute()

stopSiteRfdiagRecording



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfdiagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesRfdiagsAPI.StopSiteRfdiagRecording(context.Background(), siteId, rfdiagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRfdiagsAPI.StopSiteRfdiagRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rfdiagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSiteRfdiagRecordingRequest struct via the builder pattern


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


## UpdateSiteRfdiagRecording

> []RfDiagInfoItem UpdateSiteRfdiagRecording(ctx, siteId, rfdiagId).RfDiag(rfDiag).Execute()

updateSiteRfdiagRecording



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfdiagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	rfDiag := *openapiclient.NewRfDiag("Name_example", openapiclient.rf_client_type("")) // RfDiag | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesRfdiagsAPI.UpdateSiteRfdiagRecording(context.Background(), siteId, rfdiagId).RfDiag(rfDiag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesRfdiagsAPI.UpdateSiteRfdiagRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteRfdiagRecording`: []RfDiagInfoItem
	fmt.Fprintf(os.Stdout, "Response from `SitesRfdiagsAPI.UpdateSiteRfdiagRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rfdiagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteRfdiagRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rfDiag** | [**RfDiag**](RfDiag.md) | Request Body | 

### Return type

[**[]RfDiagInfoItem**](RfDiagInfoItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

