# \OrgsWxTagsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWxTag**](OrgsWxTagsAPI.md#CreateOrgWxTag) | **Post** /api/v1/orgs/{org_id}/wxtags | createOrgWxTag
[**DeleteOrgWxTag**](OrgsWxTagsAPI.md#DeleteOrgWxTag) | **Delete** /api/v1/orgs/{org_id}/wxtags/{wxtag_id} | deleteOrgWxTag
[**GetOrgApplicationList**](OrgsWxTagsAPI.md#GetOrgApplicationList) | **Get** /api/v1/orgs/{org_id}/wxtags/apps | getOrgApplicationList
[**GetOrgCurrentMatchingClientsOfAWxTag**](OrgsWxTagsAPI.md#GetOrgCurrentMatchingClientsOfAWxTag) | **Get** /api/v1/orgs/{org_id}/wxtags/{wxtag_id}/clients | getOrgCurrentMatchingClientsOfAWxTag
[**GetOrgWxTag**](OrgsWxTagsAPI.md#GetOrgWxTag) | **Get** /api/v1/orgs/{org_id}/wxtags/{wxtag_id} | getOrgWxTag
[**ListOrgWxTags**](OrgsWxTagsAPI.md#ListOrgWxTags) | **Get** /api/v1/orgs/{org_id}/wxtags | listOrgWxTags
[**UpdateOrgWxTag**](OrgsWxTagsAPI.md#UpdateOrgWxTag) | **Put** /api/v1/orgs/{org_id}/wxtags/{wxtag_id} | updateOrgWxTag



## CreateOrgWxTag

> WxlanTag CreateOrgWxTag(ctx, orgId).WxlanTag(wxlanTag).Execute()

createOrgWxTag



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
	wxlanTag := *openapiclient.NewWxlanTag("Name_example", openapiclient.wxlan_tag_type("")) // WxlanTag | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxTagsAPI.CreateOrgWxTag(context.Background(), orgId).WxlanTag(wxlanTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTagsAPI.CreateOrgWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgWxTag`: WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTagsAPI.CreateOrgWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWxTagRequest struct via the builder pattern


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


## DeleteOrgWxTag

> DeleteOrgWxTag(ctx, orgId, wxtagId).Execute()

deleteOrgWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsWxTagsAPI.DeleteOrgWxTag(context.Background(), orgId, wxtagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTagsAPI.DeleteOrgWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWxTagRequest struct via the builder pattern


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


## GetOrgApplicationList

> []SearchWxtagAppsItem GetOrgApplicationList(ctx, orgId).Execute()

getOrgApplicationList



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
	resp, r, err := apiClient.OrgsWxTagsAPI.GetOrgApplicationList(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTagsAPI.GetOrgApplicationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgApplicationList`: []SearchWxtagAppsItem
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTagsAPI.GetOrgApplicationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgApplicationListRequest struct via the builder pattern


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


## GetOrgCurrentMatchingClientsOfAWxTag

> []WxtagClient GetOrgCurrentMatchingClientsOfAWxTag(ctx, orgId, wxtagId).Execute()

getOrgCurrentMatchingClientsOfAWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxTagsAPI.GetOrgCurrentMatchingClientsOfAWxTag(context.Background(), orgId, wxtagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTagsAPI.GetOrgCurrentMatchingClientsOfAWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCurrentMatchingClientsOfAWxTag`: []WxtagClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTagsAPI.GetOrgCurrentMatchingClientsOfAWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCurrentMatchingClientsOfAWxTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]WxtagClient**](WxtagClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgWxTag

> WxlanTag GetOrgWxTag(ctx, orgId, wxtagId).Execute()

getOrgWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxTagsAPI.GetOrgWxTag(context.Background(), orgId, wxtagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTagsAPI.GetOrgWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgWxTag`: WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTagsAPI.GetOrgWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWxTagRequest struct via the builder pattern


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


## ListOrgWxTags

> []WxlanTag ListOrgWxTags(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgWxTags



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
	resp, r, err := apiClient.OrgsWxTagsAPI.ListOrgWxTags(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTagsAPI.ListOrgWxTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgWxTags`: []WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTagsAPI.ListOrgWxTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWxTagsRequest struct via the builder pattern


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


## UpdateOrgWxTag

> WxlanTag UpdateOrgWxTag(ctx, orgId, wxtagId).WxlanTag(wxlanTag).Execute()

updateOrgWxTag



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
	wxtagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	wxlanTag := *openapiclient.NewWxlanTag("Name_example", openapiclient.wxlan_tag_type("")) // WxlanTag | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWxTagsAPI.UpdateOrgWxTag(context.Background(), orgId, wxtagId).WxlanTag(wxlanTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWxTagsAPI.UpdateOrgWxTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgWxTag`: WxlanTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsWxTagsAPI.UpdateOrgWxTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**wxtagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWxTagRequest struct via the builder pattern


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

