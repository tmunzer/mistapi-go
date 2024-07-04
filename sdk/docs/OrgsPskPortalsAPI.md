# \OrgsPskPortalsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgPskPortalLogs**](OrgsPskPortalsAPI.md#CountOrgPskPortalLogs) | **Get** /api/v1/orgs/{org_id}/pskportals/logs/count | countOrgPskPortalLogs
[**CreateOrgPskPortal**](OrgsPskPortalsAPI.md#CreateOrgPskPortal) | **Post** /api/v1/orgs/{org_id}/pskportals | createOrgPskPortal
[**DeleteOrgPskPortal**](OrgsPskPortalsAPI.md#DeleteOrgPskPortal) | **Delete** /api/v1/orgs/{org_id}/pskportals/{pskportal_id} | deleteOrgPskPortal
[**DeleteOrgPskPortalImage**](OrgsPskPortalsAPI.md#DeleteOrgPskPortalImage) | **Delete** /api/v1/orgs/{org_id}/pskportals/{pskportal_id}/portal_image | deleteOrgPskPortalImage
[**GetOrgPskPortal**](OrgsPskPortalsAPI.md#GetOrgPskPortal) | **Get** /api/v1/orgs/{org_id}/pskportals/{pskportal_id} | getOrgPskPortal
[**ListOrgPskPortalLogs**](OrgsPskPortalsAPI.md#ListOrgPskPortalLogs) | **Get** /api/v1/orgs/{org_id}/pskportals/logs | listOrgPskPortalLogs
[**ListOrgPskPortals**](OrgsPskPortalsAPI.md#ListOrgPskPortals) | **Get** /api/v1/orgs/{org_id}/pskportals | listOrgPskPortals
[**SearchOrgPskPortalLogs**](OrgsPskPortalsAPI.md#SearchOrgPskPortalLogs) | **Get** /api/v1/orgs/{org_id}/pskportals/logs/search | searchOrgPskPortalLogs
[**UpdateOrgPskPortal**](OrgsPskPortalsAPI.md#UpdateOrgPskPortal) | **Put** /api/v1/orgs/{org_id}/pskportals/{pskportal_id} | updateOrgPskPortal
[**UpdateOrgPskPortalTemplate**](OrgsPskPortalsAPI.md#UpdateOrgPskPortalTemplate) | **Put** /api/v1/orgs/{org_id}/pskportals/{pskportal_id}/portal_template | updateOrgPskPortalTemplate
[**UploadOrgPskPortalImage**](OrgsPskPortalsAPI.md#UploadOrgPskPortalImage) | **Post** /api/v1/orgs/{org_id}/pskportals/{pskportal_id}/portal_image | uploadOrgPskPortalImage



## CountOrgPskPortalLogs

> RepsonseCount CountOrgPskPortalLogs(ctx, orgId).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

countOrgPskPortalLogs



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
	distinct := openapiclient.org_psk_portal_logs_count_distinct("") // OrgPskPortalLogsCountDistinct |  (optional) (default to "pskportal_id")
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPskPortalsAPI.CountOrgPskPortalLogs(context.Background(), orgId).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.CountOrgPskPortalLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgPskPortalLogs`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsPskPortalsAPI.CountOrgPskPortalLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgPskPortalLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgPskPortalLogsCountDistinct**](OrgPskPortalLogsCountDistinct.md) |  | [default to &quot;pskportal_id&quot;]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**RepsonseCount**](RepsonseCount.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgPskPortal

> PskPortal CreateOrgPskPortal(ctx, orgId).PskPortal(pskPortal).Execute()

createOrgPskPortal



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
	pskPortal := *openapiclient.NewPskPortal("Name_example", "Ssid_example") // PskPortal |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPskPortalsAPI.CreateOrgPskPortal(context.Background(), orgId).PskPortal(pskPortal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.CreateOrgPskPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgPskPortal`: PskPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsPskPortalsAPI.CreateOrgPskPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgPskPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pskPortal** | [**PskPortal**](PskPortal.md) |  | 

### Return type

[**PskPortal**](PskPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgPskPortal

> DeleteOrgPskPortal(ctx, orgId, pskportalId).Execute()

deleteOrgPskPortal



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
	pskportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsPskPortalsAPI.DeleteOrgPskPortal(context.Background(), orgId, pskportalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.DeleteOrgPskPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgPskPortalRequest struct via the builder pattern


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


## DeleteOrgPskPortalImage

> DeleteOrgPskPortalImage(ctx, orgId, pskportalId).Execute()

deleteOrgPskPortalImage



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
	pskportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsPskPortalsAPI.DeleteOrgPskPortalImage(context.Background(), orgId, pskportalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.DeleteOrgPskPortalImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgPskPortalImageRequest struct via the builder pattern


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


## GetOrgPskPortal

> PskPortal GetOrgPskPortal(ctx, orgId, pskportalId).Execute()

getOrgPskPortal



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
	pskportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPskPortalsAPI.GetOrgPskPortal(context.Background(), orgId, pskportalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.GetOrgPskPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgPskPortal`: PskPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsPskPortalsAPI.GetOrgPskPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgPskPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PskPortal**](PskPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgPskPortalLogs

> ResponsePskPortalLogsSearch ListOrgPskPortalLogs(ctx, orgId).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()

listOrgPskPortalLogs



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
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPskPortalsAPI.ListOrgPskPortalLogs(context.Background(), orgId).Start(start).End(end).Duration(duration).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.ListOrgPskPortalLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgPskPortalLogs`: ResponsePskPortalLogsSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsPskPortalsAPI.ListOrgPskPortalLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgPskPortalLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**ResponsePskPortalLogsSearch**](ResponsePskPortalLogsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgPskPortals

> []PskPortal ListOrgPskPortals(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgPskPortals



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPskPortalsAPI.ListOrgPskPortals(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.ListOrgPskPortals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgPskPortals`: []PskPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsPskPortalsAPI.ListOrgPskPortals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgPskPortalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]PskPortal**](PskPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgPskPortalLogs

> ResponsePskPortalLogsSearch SearchOrgPskPortalLogs(ctx, orgId).Start(start).End(end).Duration(duration).Limit(limit).Page(page).PskName(pskName).PskId(pskId).PskportalId(pskportalId).Id(id).AdminName(adminName).AdminId(adminId).NameId(nameId).Execute()

searchOrgPskPortalLogs



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
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)
	pskName := "pskName_example" // string |  (optional)
	pskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	pskportalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | audit_id (optional)
	adminName := "adminName_example" // string |  (optional)
	adminId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	nameId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | name_id used in SSO (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPskPortalsAPI.SearchOrgPskPortalLogs(context.Background(), orgId).Start(start).End(end).Duration(duration).Limit(limit).Page(page).PskName(pskName).PskId(pskId).PskportalId(pskportalId).Id(id).AdminName(adminName).AdminId(adminId).NameId(nameId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.SearchOrgPskPortalLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgPskPortalLogs`: ResponsePskPortalLogsSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsPskPortalsAPI.SearchOrgPskPortalLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgPskPortalLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]
 **pskName** | **string** |  | 
 **pskId** | **string** |  | 
 **pskportalId** | **string** |  | 
 **id** | **string** | audit_id | 
 **adminName** | **string** |  | 
 **adminId** | **string** |  | 
 **nameId** | **string** | name_id used in SSO | 

### Return type

[**ResponsePskPortalLogsSearch**](ResponsePskPortalLogsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgPskPortal

> PskPortal UpdateOrgPskPortal(ctx, orgId, pskportalId).PskPortal(pskPortal).Execute()

updateOrgPskPortal



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
	pskportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	pskPortal := *openapiclient.NewPskPortal("Name_example", "Ssid_example") // PskPortal |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPskPortalsAPI.UpdateOrgPskPortal(context.Background(), orgId, pskportalId).PskPortal(pskPortal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.UpdateOrgPskPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgPskPortal`: PskPortal
	fmt.Fprintf(os.Stdout, "Response from `OrgsPskPortalsAPI.UpdateOrgPskPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgPskPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pskPortal** | [**PskPortal**](PskPortal.md) |  | 

### Return type

[**PskPortal**](PskPortal.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgPskPortalTemplate

> UpdateOrgPskPortalTemplate(ctx, orgId, pskportalId).PskPortalTemplate(pskPortalTemplate).Execute()

updateOrgPskPortalTemplate



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
	pskportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	pskPortalTemplate := *openapiclient.NewPskPortalTemplate() // PskPortalTemplate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsPskPortalsAPI.UpdateOrgPskPortalTemplate(context.Background(), orgId, pskportalId).PskPortalTemplate(pskPortalTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.UpdateOrgPskPortalTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgPskPortalTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pskPortalTemplate** | [**PskPortalTemplate**](PskPortalTemplate.md) |  | 

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


## UploadOrgPskPortalImage

> UploadOrgPskPortalImage(ctx, orgId, pskportalId).File(file).Json(json).Execute()

uploadOrgPskPortalImage



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
	pskportalId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | Binary file (optional)
	json := "json_example" // string | JSON string describing the upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsPskPortalsAPI.UploadOrgPskPortalImage(context.Background(), orgId, pskportalId).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPskPortalsAPI.UploadOrgPskPortalImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskportalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadOrgPskPortalImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | Binary file | 
 **json** | **string** | JSON string describing the upload | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

