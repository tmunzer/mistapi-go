# \OrgsSitegroupsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSiteGroup**](OrgsSitegroupsAPI.md#CreateOrgSiteGroup) | **Post** /api/v1/orgs/{org_id}/sitegroups | createOrgSiteGroup
[**DeleteOrgSiteGroup**](OrgsSitegroupsAPI.md#DeleteOrgSiteGroup) | **Delete** /api/v1/orgs/{org_id}/sitegroups/{sitegroup_id} | deleteOrgSiteGroup
[**GetOrgSiteGroup**](OrgsSitegroupsAPI.md#GetOrgSiteGroup) | **Get** /api/v1/orgs/{org_id}/sitegroups/{sitegroup_id} | getOrgSiteGroup
[**ListOrgSiteGroups**](OrgsSitegroupsAPI.md#ListOrgSiteGroups) | **Get** /api/v1/orgs/{org_id}/sitegroups | listOrgSiteGroups
[**UpdateOrgSiteGroup**](OrgsSitegroupsAPI.md#UpdateOrgSiteGroup) | **Put** /api/v1/orgs/{org_id}/sitegroups/{sitegroup_id} | updateOrgSiteGroup



## CreateOrgSiteGroup

> Sitegroup CreateOrgSiteGroup(ctx, orgId).Sitegroup(sitegroup).Execute()

createOrgSiteGroup



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
	sitegroup := *openapiclient.NewSitegroup("Name_example") // Sitegroup | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitegroupsAPI.CreateOrgSiteGroup(context.Background(), orgId).Sitegroup(sitegroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitegroupsAPI.CreateOrgSiteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSiteGroup`: Sitegroup
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitegroupsAPI.CreateOrgSiteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSiteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sitegroup** | [**Sitegroup**](Sitegroup.md) | Request Body | 

### Return type

[**Sitegroup**](Sitegroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgSiteGroup

> DeleteOrgSiteGroup(ctx, orgId, sitegroupId).Execute()

deleteOrgSiteGroup



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
	sitegroupId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSitegroupsAPI.DeleteOrgSiteGroup(context.Background(), orgId, sitegroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitegroupsAPI.DeleteOrgSiteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sitegroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSiteGroupRequest struct via the builder pattern


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


## GetOrgSiteGroup

> Sitegroup GetOrgSiteGroup(ctx, orgId, sitegroupId).Execute()

getOrgSiteGroup



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
	sitegroupId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitegroupsAPI.GetOrgSiteGroup(context.Background(), orgId, sitegroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitegroupsAPI.GetOrgSiteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSiteGroup`: Sitegroup
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitegroupsAPI.GetOrgSiteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sitegroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSiteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sitegroup**](Sitegroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSiteGroups

> []Sitegroup ListOrgSiteGroups(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgSiteGroups



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
	resp, r, err := apiClient.OrgsSitegroupsAPI.ListOrgSiteGroups(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitegroupsAPI.ListOrgSiteGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSiteGroups`: []Sitegroup
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitegroupsAPI.ListOrgSiteGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSiteGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Sitegroup**](Sitegroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSiteGroup

> Sitegroup UpdateOrgSiteGroup(ctx, orgId, sitegroupId).NameString(nameString).Execute()

updateOrgSiteGroup



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
	sitegroupId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	nameString := *openapiclient.NewNameString() // NameString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitegroupsAPI.UpdateOrgSiteGroup(context.Background(), orgId, sitegroupId).NameString(nameString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitegroupsAPI.UpdateOrgSiteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSiteGroup`: Sitegroup
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitegroupsAPI.UpdateOrgSiteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sitegroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSiteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nameString** | [**NameString**](NameString.md) | Request Body | 

### Return type

[**Sitegroup**](Sitegroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

