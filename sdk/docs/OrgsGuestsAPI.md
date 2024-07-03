# \OrgsGuestsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgGuestAuthorizations**](OrgsGuestsAPI.md#CountOrgGuestAuthorizations) | **Get** /api/v1/orgs/{org_id}/guests/count | countOrgGuestAuthorizations
[**DeleteOrgGuestAuthorization**](OrgsGuestsAPI.md#DeleteOrgGuestAuthorization) | **Delete** /api/v1/orgs/{org_id}/guests/{guest_mac} | deleteOrgGuestAuthorization
[**GetOrgGuestAuthorization**](OrgsGuestsAPI.md#GetOrgGuestAuthorization) | **Get** /api/v1/orgs/{org_id}/guests/{guest_mac} | getOrgGuestAuthorization
[**ListOrgGuestAuthorizations**](OrgsGuestsAPI.md#ListOrgGuestAuthorizations) | **Get** /api/v1/orgs/{org_id}/guests | listOrgGuestAuthorizations
[**SearchOrgGuestAuthorization**](OrgsGuestsAPI.md#SearchOrgGuestAuthorization) | **Get** /api/v1/orgs/{org_id}/guests/search | searchOrgGuestAuthorization
[**UpdateOrgGuestAuthorization**](OrgsGuestsAPI.md#UpdateOrgGuestAuthorization) | **Put** /api/v1/orgs/{org_id}/guests/{guest_mac} | updateOrgGuestAuthorization



## CountOrgGuestAuthorizations

> RepsonseCount CountOrgGuestAuthorizations(ctx, orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgGuestAuthorizations



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
	distinct := openapiclient.org_guests_count_distinct("") // OrgGuestsCountDistinct |  (optional) (default to "auth_method")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGuestsAPI.CountOrgGuestAuthorizations(context.Background(), orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGuestsAPI.CountOrgGuestAuthorizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgGuestAuthorizations`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsGuestsAPI.CountOrgGuestAuthorizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgGuestAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgGuestsCountDistinct**](OrgGuestsCountDistinct.md) |  | [default to &quot;auth_method&quot;]
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

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


## DeleteOrgGuestAuthorization

> DeleteOrgGuestAuthorization(ctx, orgId, guestMac).Execute()

deleteOrgGuestAuthorization



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
	guestMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsGuestsAPI.DeleteOrgGuestAuthorization(context.Background(), orgId, guestMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGuestsAPI.DeleteOrgGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**guestMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgGuestAuthorizationRequest struct via the builder pattern


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


## GetOrgGuestAuthorization

> Guest GetOrgGuestAuthorization(ctx, orgId, guestMac).Execute()

getOrgGuestAuthorization



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
	guestMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGuestsAPI.GetOrgGuestAuthorization(context.Background(), orgId, guestMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGuestsAPI.GetOrgGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgGuestAuthorization`: Guest
	fmt.Fprintf(os.Stdout, "Response from `OrgsGuestsAPI.GetOrgGuestAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**guestMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgGuestAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Guest**](Guest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgGuestAuthorizations

> []Guest ListOrgGuestAuthorizations(ctx, orgId).Execute()

listOrgGuestAuthorizations



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGuestsAPI.ListOrgGuestAuthorizations(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGuestsAPI.ListOrgGuestAuthorizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgGuestAuthorizations`: []Guest
	fmt.Fprintf(os.Stdout, "Response from `OrgsGuestsAPI.ListOrgGuestAuthorizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgGuestAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Guest**](Guest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgGuestAuthorization

> ResponseGuestSearch SearchOrgGuestAuthorization(ctx, orgId).WlanId(wlanId).AuthMethod(authMethod).Ssid(ssid).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgGuestAuthorization



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
	wlanId := "00000000-0000-0000-0000-000000000000" // string | WLAN ID (optional)
	authMethod := "authMethod_example" // string | Authentication Methdo (optional)
	ssid := "ssid_example" // string | SSID (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGuestsAPI.SearchOrgGuestAuthorization(context.Background(), orgId).WlanId(wlanId).AuthMethod(authMethod).Ssid(ssid).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGuestsAPI.SearchOrgGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgGuestAuthorization`: ResponseGuestSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsGuestsAPI.SearchOrgGuestAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgGuestAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **string** | WLAN ID | 
 **authMethod** | **string** | Authentication Methdo | 
 **ssid** | **string** | SSID | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseGuestSearch**](ResponseGuestSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgGuestAuthorization

> Guest UpdateOrgGuestAuthorization(ctx, orgId, guestMac).Guest(guest).Execute()

updateOrgGuestAuthorization



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
	guestMac := "0000000000ab" // string | 
	guest := *openapiclient.NewGuest() // Guest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsGuestsAPI.UpdateOrgGuestAuthorization(context.Background(), orgId, guestMac).Guest(guest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsGuestsAPI.UpdateOrgGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgGuestAuthorization`: Guest
	fmt.Fprintf(os.Stdout, "Response from `OrgsGuestsAPI.UpdateOrgGuestAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**guestMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgGuestAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **guest** | [**Guest**](Guest.md) |  | 

### Return type

[**Guest**](Guest.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

