# \OrgsNACTagsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgNacTag**](OrgsNACTagsAPI.md#CreateOrgNacTag) | **Post** /api/v1/orgs/{org_id}/nactags | createOrgNacTag
[**DeleteOrgNacTag**](OrgsNACTagsAPI.md#DeleteOrgNacTag) | **Delete** /api/v1/orgs/{org_id}/nactags/{nactag_id} | deleteOrgNacTag
[**GetOrgNacTag**](OrgsNACTagsAPI.md#GetOrgNacTag) | **Get** /api/v1/orgs/{org_id}/nactags/{nactag_id} | getOrgNacTag
[**ListOrgNacTags**](OrgsNACTagsAPI.md#ListOrgNacTags) | **Get** /api/v1/orgs/{org_id}/nactags | listOrgNacTags
[**UpdateOrgNacTag**](OrgsNACTagsAPI.md#UpdateOrgNacTag) | **Put** /api/v1/orgs/{org_id}/nactags/{nactag_id} | updateOrgNacTag



## CreateOrgNacTag

> NacTag CreateOrgNacTag(ctx, orgId).NacTag(nacTag).Execute()

createOrgNacTag



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nacTag := *openapiclient.NewNacTag("Name_example", openapiclient.nac_tag_type("")) // NacTag |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACTagsAPI.CreateOrgNacTag(context.Background(), orgId).NacTag(nacTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACTagsAPI.CreateOrgNacTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgNacTag`: NacTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACTagsAPI.CreateOrgNacTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgNacTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nacTag** | [**NacTag**](NacTag.md) |  | 

### Return type

[**NacTag**](NacTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgNacTag

> DeleteOrgNacTag(ctx, orgId, nactagId).Execute()

deleteOrgNacTag



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nactagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsNACTagsAPI.DeleteOrgNacTag(context.Background(), orgId, nactagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACTagsAPI.DeleteOrgNacTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nactagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgNacTagRequest struct via the builder pattern


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


## GetOrgNacTag

> NacTag GetOrgNacTag(ctx, orgId, nactagId).Execute()

getOrgNacTag



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nactagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACTagsAPI.GetOrgNacTag(context.Background(), orgId, nactagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACTagsAPI.GetOrgNacTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgNacTag`: NacTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACTagsAPI.GetOrgNacTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nactagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNacTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NacTag**](NacTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgNacTags

> []NacTag ListOrgNacTags(ctx, orgId).Type_(type_).Name(name).Match(match).Page(page).Limit(limit).Execute()

listOrgNacTags



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	type_ := "type__example" // string | Type of NAC Tag (optional)
	name := "name_example" // string | Name of NAC Tag (optional)
	match := "match_example" // string | Type of NAC Tag (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACTagsAPI.ListOrgNacTags(context.Background(), orgId).Type_(type_).Name(name).Match(match).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACTagsAPI.ListOrgNacTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgNacTags`: []NacTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACTagsAPI.ListOrgNacTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgNacTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of NAC Tag | 
 **name** | **string** | Name of NAC Tag | 
 **match** | **string** | Type of NAC Tag | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]NacTag**](NacTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgNacTag

> NacTag UpdateOrgNacTag(ctx, orgId, nactagId).NacTag(nacTag).Execute()

updateOrgNacTag



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nactagId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nacTag := *openapiclient.NewNacTag("Name_example", openapiclient.nac_tag_type("")) // NacTag |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACTagsAPI.UpdateOrgNacTag(context.Background(), orgId, nactagId).NacTag(nacTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACTagsAPI.UpdateOrgNacTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgNacTag`: NacTag
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACTagsAPI.UpdateOrgNacTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**nactagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgNacTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nacTag** | [**NacTag**](NacTag.md) |  | 

### Return type

[**NacTag**](NacTag.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

