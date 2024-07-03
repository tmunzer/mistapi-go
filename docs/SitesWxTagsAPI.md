# \SitesWxTagsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteWxTag**](SitesWxTagsAPI.md#CreateSiteWxTag) | **Post** /api/v1/sites/{site_id}/wxtags | createSiteWxTag
[**DeleteSiteWxTag**](SitesWxTagsAPI.md#DeleteSiteWxTag) | **Delete** /api/v1/sites/{site_id}/wxtags/{wxtag_id} | deleteSiteWxTag
[**GetSiteApplicationList**](SitesWxTagsAPI.md#GetSiteApplicationList) | **Get** /api/v1/sites/{site_id}/wxtags/apps | getSiteApplicationList
[**GetSiteCurrentMatchingClientsOfAWxTag**](SitesWxTagsAPI.md#GetSiteCurrentMatchingClientsOfAWxTag) | **Get** /api/v1/sites/{site_id}/wxtags/{wxtag_id}/clients | getSiteCurrentMatchingClientsOfAWxTag
[**GetSiteWxTag**](SitesWxTagsAPI.md#GetSiteWxTag) | **Get** /api/v1/sites/{site_id}/wxtags/{wxtag_id} | getSiteWxTag
[**ListSiteWxTags**](SitesWxTagsAPI.md#ListSiteWxTags) | **Get** /api/v1/sites/{site_id}/wxtags | listSiteWxTags
[**UpdateSiteWxTag**](SitesWxTagsAPI.md#UpdateSiteWxTag) | **Put** /api/v1/sites/{site_id}/wxtags/{wxtag_id} | updateSiteWxTag



## CreateSiteWxTag

> WxlanTag CreateSiteWxTag(ctx, siteId).WxlanTag(wxlanTag).Execute()

createSiteWxTag



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
	wxlanTag := *openapiclient.NewWxlanTag("Name_example", openapiclient.wxlan_tag_type("")) // WxlanTag | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTagsAPI.CreateSiteWxTag(context.Background(), siteId).WxlanTag(wxlanTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTagsAPI.CreateSiteWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWxTag`: WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTagsAPI.CreateSiteWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWxTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wxlanTag** | [**WxlanTag**](WxlanTag.md) | Request Body | 

### Return type

[**WxlanTag**](WxlanTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteWxTag

> DeleteSiteWxTag(ctx, siteId, wxtagId).Execute()

deleteSiteWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesWxTagsAPI.DeleteSiteWxTag(context.Background(), siteId, wxtagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTagsAPI.DeleteSiteWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteWxTagRequest struct via the builder pattern


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


## GetSiteApplicationList

> []SearchWxtagAppsItem GetSiteApplicationList(ctx, siteId).Execute()

getSiteApplicationList



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
	resp, r, err := apiClient.SitesWxTagsAPI.GetSiteApplicationList(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTagsAPI.GetSiteApplicationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteApplicationList`: []SearchWxtagAppsItem
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTagsAPI.GetSiteApplicationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SearchWxtagAppsItem**](SearchWxtagAppsItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteCurrentMatchingClientsOfAWxTag

> []WxtagMatching GetSiteCurrentMatchingClientsOfAWxTag(ctx, siteId, wxtagId).Execute()

getSiteCurrentMatchingClientsOfAWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTagsAPI.GetSiteCurrentMatchingClientsOfAWxTag(context.Background(), siteId, wxtagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTagsAPI.GetSiteCurrentMatchingClientsOfAWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteCurrentMatchingClientsOfAWxTag`: []WxtagMatching
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTagsAPI.GetSiteCurrentMatchingClientsOfAWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteCurrentMatchingClientsOfAWxTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WxtagMatching**](WxtagMatching.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteWxTag

> WxlanTag GetSiteWxTag(ctx, siteId, wxtagId).Execute()

getSiteWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTagsAPI.GetSiteWxTag(context.Background(), siteId, wxtagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTagsAPI.GetSiteWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWxTag`: WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTagsAPI.GetSiteWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWxTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WxlanTag**](WxlanTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteWxTags

> []WxlanTag ListSiteWxTags(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteWxTags



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTagsAPI.ListSiteWxTags(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTagsAPI.ListSiteWxTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteWxTags`: []WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTagsAPI.ListSiteWxTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteWxTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]WxlanTag**](WxlanTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteWxTag

> WxlanTag UpdateSiteWxTag(ctx, siteId, wxtagId).WxlanTag(wxlanTag).Execute()

updateSiteWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanTag := *openapiclient.NewWxlanTag("Name_example", openapiclient.wxlan_tag_type("")) // WxlanTag | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWxTagsAPI.UpdateSiteWxTag(context.Background(), siteId, wxtagId).WxlanTag(wxlanTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWxTagsAPI.UpdateSiteWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteWxTag`: WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `SitesWxTagsAPI.UpdateSiteWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteWxTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wxlanTag** | [**WxlanTag**](WxlanTag.md) | Request Body | 

### Return type

[**WxlanTag**](WxlanTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

