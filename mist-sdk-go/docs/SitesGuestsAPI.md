# \SitesGuestsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteGuestAuthorizations**](SitesGuestsAPI.md#CountSiteGuestAuthorizations) | **Get** /api/v1/sites/{site_id}/guests/count | countSiteGuestAuthorizations
[**DeleteSiteGuestAuthorization**](SitesGuestsAPI.md#DeleteSiteGuestAuthorization) | **Delete** /api/v1/sites/{site_id}/guests/{guest_mac} | deleteSiteGuestAuthorization
[**GetSiteGuestAuthorization**](SitesGuestsAPI.md#GetSiteGuestAuthorization) | **Get** /api/v1/sites/{site_id}/guests/{guest_mac} | getSiteGuestAuthorization
[**ListSiteAllGuestAuthorizations**](SitesGuestsAPI.md#ListSiteAllGuestAuthorizations) | **Get** /api/v1/sites/{site_id}/guests | listSiteAllGuestAuthorizations
[**ListSiteAllGuestAuthorizationsDerived**](SitesGuestsAPI.md#ListSiteAllGuestAuthorizationsDerived) | **Get** /api/v1/sites/{site_id}/guests/derived | listSiteAllGuestAuthorizationsDerived
[**SearchSiteGuestAuthorization**](SitesGuestsAPI.md#SearchSiteGuestAuthorization) | **Get** /api/v1/sites/{site_id}/guests/search | searchSiteGuestAuthorization
[**UpdateSiteGuestAuthorization**](SitesGuestsAPI.md#UpdateSiteGuestAuthorization) | **Put** /api/v1/sites/{site_id}/guests/{guest_mac} | updateSiteGuestAuthorization



## CountSiteGuestAuthorizations

> RepsonseCount CountSiteGuestAuthorizations(ctx, siteId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteGuestAuthorizations



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
	distinct := openapiclient.site_guests_count_distinct("") // SiteGuestsCountDistinct |  (optional) (default to "auth_method")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesGuestsAPI.CountSiteGuestAuthorizations(context.Background(), siteId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesGuestsAPI.CountSiteGuestAuthorizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteGuestAuthorizations`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesGuestsAPI.CountSiteGuestAuthorizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteGuestAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**SiteGuestsCountDistinct**](SiteGuestsCountDistinct.md) |  | [default to &quot;auth_method&quot;]
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


## DeleteSiteGuestAuthorization

> DeleteSiteGuestAuthorization(ctx, siteId, guestMac).Execute()

deleteSiteGuestAuthorization



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
	guestMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesGuestsAPI.DeleteSiteGuestAuthorization(context.Background(), siteId, guestMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesGuestsAPI.DeleteSiteGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**guestMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteGuestAuthorizationRequest struct via the builder pattern


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


## GetSiteGuestAuthorization

> Guest GetSiteGuestAuthorization(ctx, siteId, guestMac).Execute()

getSiteGuestAuthorization



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
	guestMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesGuestsAPI.GetSiteGuestAuthorization(context.Background(), siteId, guestMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesGuestsAPI.GetSiteGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteGuestAuthorization`: Guest
	fmt.Fprintf(os.Stdout, "Response from `SitesGuestsAPI.GetSiteGuestAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**guestMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteGuestAuthorizationRequest struct via the builder pattern


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


## ListSiteAllGuestAuthorizations

> []Guest ListSiteAllGuestAuthorizations(ctx, siteId).WlanId(wlanId).Execute()

listSiteAllGuestAuthorizations



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
	wlanId := "wlanId_example" // string | UUID of single or multiple (Comma separated) WLAN under Site `site_id` (to filter by WLAN) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesGuestsAPI.ListSiteAllGuestAuthorizations(context.Background(), siteId).WlanId(wlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesGuestsAPI.ListSiteAllGuestAuthorizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteAllGuestAuthorizations`: []Guest
	fmt.Fprintf(os.Stdout, "Response from `SitesGuestsAPI.ListSiteAllGuestAuthorizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteAllGuestAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **string** | UUID of single or multiple (Comma separated) WLAN under Site &#x60;site_id&#x60; (to filter by WLAN) | 

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


## ListSiteAllGuestAuthorizationsDerived

> []Guest ListSiteAllGuestAuthorizationsDerived(ctx, siteId).WlanId(wlanId).CrossSite(crossSite).Execute()

listSiteAllGuestAuthorizationsDerived



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
	wlanId := "wlanId_example" // string | UUID of single or multiple (Comma separated) WLAN under Site `site_id` (to filter by WLAN) (optional)
	crossSite := true // bool | whether to get org level guests, default is false i.e get site level guests (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesGuestsAPI.ListSiteAllGuestAuthorizationsDerived(context.Background(), siteId).WlanId(wlanId).CrossSite(crossSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesGuestsAPI.ListSiteAllGuestAuthorizationsDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteAllGuestAuthorizationsDerived`: []Guest
	fmt.Fprintf(os.Stdout, "Response from `SitesGuestsAPI.ListSiteAllGuestAuthorizationsDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteAllGuestAuthorizationsDerivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **string** | UUID of single or multiple (Comma separated) WLAN under Site &#x60;site_id&#x60; (to filter by WLAN) | 
 **crossSite** | **bool** | whether to get org level guests, default is false i.e get site level guests | [default to false]

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


## SearchSiteGuestAuthorization

> ResponseGuestSearch SearchSiteGuestAuthorization(ctx, siteId).WlanId(wlanId).AuthMethod(authMethod).Ssid(ssid).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteGuestAuthorization



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
	wlanId := "00000000-0000-0000-0000-000000000000" // string |  (optional)
	authMethod := "authMethod_example" // string |  (optional)
	ssid := "ssid_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesGuestsAPI.SearchSiteGuestAuthorization(context.Background(), siteId).WlanId(wlanId).AuthMethod(authMethod).Ssid(ssid).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesGuestsAPI.SearchSiteGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteGuestAuthorization`: ResponseGuestSearch
	fmt.Fprintf(os.Stdout, "Response from `SitesGuestsAPI.SearchSiteGuestAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteGuestAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wlanId** | **string** |  | 
 **authMethod** | **string** |  | 
 **ssid** | **string** |  | 
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


## UpdateSiteGuestAuthorization

> Guest UpdateSiteGuestAuthorization(ctx, siteId, guestMac).Guest(guest).Execute()

updateSiteGuestAuthorization



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
	guestMac := "0000000000ab" // string | 
	guest := *openapiclient.NewGuest() // Guest | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesGuestsAPI.UpdateSiteGuestAuthorization(context.Background(), siteId, guestMac).Guest(guest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesGuestsAPI.UpdateSiteGuestAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteGuestAuthorization`: Guest
	fmt.Fprintf(os.Stdout, "Response from `SitesGuestsAPI.UpdateSiteGuestAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**guestMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteGuestAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **guest** | [**Guest**](Guest.md) | Request Body | 

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

