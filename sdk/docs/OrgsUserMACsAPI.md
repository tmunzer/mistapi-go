# \OrgsUserMACsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgUserMacs**](OrgsUserMACsAPI.md#CreateOrgUserMacs) | **Post** /api/v1/orgs/{org_id}/usermacs | createOrgUserMacs
[**DeleteOrgUserMac**](OrgsUserMACsAPI.md#DeleteOrgUserMac) | **Delete** /api/v1/orgs/{org_id}/usermacs/search/{usermac_id} | deleteOrgUserMac
[**GetOrgUserMac**](OrgsUserMACsAPI.md#GetOrgUserMac) | **Get** /api/v1/orgs/{org_id}/usermacs/search/{usermac_id} | getOrgUserMac
[**ImportOrgUserMacs**](OrgsUserMACsAPI.md#ImportOrgUserMacs) | **Post** /api/v1/orgs/{org_id}/usermacs/import | importOrgUserMacs
[**ListOrgUserMacs**](OrgsUserMACsAPI.md#ListOrgUserMacs) | **Get** /api/v1/orgs/{org_id}/usermacs | listOrgUserMacs
[**SearchOrgUserMacs**](OrgsUserMACsAPI.md#SearchOrgUserMacs) | **Get** /api/v1/orgs/{org_id}/usermacs/search | searchOrgUserMacs



## CreateOrgUserMacs

> UserMacImport CreateOrgUserMacs(ctx, orgId).UserMac(userMac).Execute()

createOrgUserMacs



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
	userMac := *openapiclient.NewUserMac() // UserMac |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsUserMACsAPI.CreateOrgUserMacs(context.Background(), orgId).UserMac(userMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsUserMACsAPI.CreateOrgUserMacs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgUserMacs`: UserMacImport
	fmt.Fprintf(os.Stdout, "Response from `OrgsUserMACsAPI.CreateOrgUserMacs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgUserMacsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userMac** | [**UserMac**](UserMac.md) |  | 

### Return type

[**UserMacImport**](UserMacImport.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgUserMac

> DeleteOrgUserMac(ctx, orgId, usermacId).Execute()

deleteOrgUserMac



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
	usermacId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsUserMACsAPI.DeleteOrgUserMac(context.Background(), orgId, usermacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsUserMACsAPI.DeleteOrgUserMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**usermacId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgUserMacRequest struct via the builder pattern


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


## GetOrgUserMac

> UserMac GetOrgUserMac(ctx, orgId, usermacId).Execute()

getOrgUserMac



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
	usermacId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsUserMACsAPI.GetOrgUserMac(context.Background(), orgId, usermacId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsUserMACsAPI.GetOrgUserMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgUserMac`: UserMac
	fmt.Fprintf(os.Stdout, "Response from `OrgsUserMACsAPI.GetOrgUserMac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**usermacId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgUserMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserMac**](UserMac.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOrgUserMacs

> ImportOrgUserMacs(ctx, orgId).UserMac(userMac).Execute()

importOrgUserMacs



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
	userMac := []openapiclient.UserMac{*openapiclient.NewUserMac()} // []UserMac |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsUserMACsAPI.ImportOrgUserMacs(context.Background(), orgId).UserMac(userMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsUserMACsAPI.ImportOrgUserMacs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiImportOrgUserMacsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userMac** | [**[]UserMac**](UserMac.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgUserMacs

> []UserMac ListOrgUserMacs(ctx, orgId).Blacklisted(blacklisted).ForGuestWifi(forGuestWifi).CrossSite(crossSite).SiteId(siteId).Page(page).Limit(limit).Execute()

listOrgUserMacs



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
	blacklisted := true // bool |  (optional)
	forGuestWifi := true // bool |  (optional)
	crossSite := true // bool |  (optional)
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsUserMACsAPI.ListOrgUserMacs(context.Background(), orgId).Blacklisted(blacklisted).ForGuestWifi(forGuestWifi).CrossSite(crossSite).SiteId(siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsUserMACsAPI.ListOrgUserMacs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgUserMacs`: []UserMac
	fmt.Fprintf(os.Stdout, "Response from `OrgsUserMACsAPI.ListOrgUserMacs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgUserMacsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blacklisted** | **bool** |  | 
 **forGuestWifi** | **bool** |  | 
 **crossSite** | **bool** |  | 
 **siteId** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]UserMac**](UserMac.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgUserMacs

> UserMacsSearch SearchOrgUserMacs(ctx, orgId).Mac(mac).Labels(labels).Page(page).Limit(limit).Execute()

searchOrgUserMacs



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
	mac := "mac_example" // string | mac address (optional)
	labels := []string{"Inner_example"} // []string | optional, array of strings of labels (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsUserMACsAPI.SearchOrgUserMacs(context.Background(), orgId).Mac(mac).Labels(labels).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsUserMACsAPI.SearchOrgUserMacs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgUserMacs`: UserMacsSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsUserMACsAPI.SearchOrgUserMacs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgUserMacsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | mac address | 
 **labels** | **[]string** | optional, array of strings of labels | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**UserMacsSearch**](UserMacsSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

