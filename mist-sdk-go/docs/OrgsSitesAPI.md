# \OrgsSitesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgSites**](OrgsSitesAPI.md#CountOrgSites) | **Get** /api/v1/orgs/{org_id}/sites/count | countOrgSites
[**CreateOrgSite**](OrgsSitesAPI.md#CreateOrgSite) | **Post** /api/v1/orgs/{org_id}/sites | createOrgSite
[**ListOrgSites**](OrgsSitesAPI.md#ListOrgSites) | **Get** /api/v1/orgs/{org_id}/sites | listOrgSites
[**SearchOrgSites**](OrgsSitesAPI.md#SearchOrgSites) | **Get** /api/v1/orgs/{org_id}/sites/search | searchOrgSites



## CountOrgSites

> RepsonseCount CountOrgSites(ctx, orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgSites



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
	distinct := openapiclient.org_sites_count_distinct("") // OrgSitesCountDistinct |  (optional) (default to "id")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitesAPI.CountOrgSites(context.Background(), orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitesAPI.CountOrgSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgSites`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitesAPI.CountOrgSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgSitesCountDistinct**](OrgSitesCountDistinct.md) |  | [default to &quot;id&quot;]
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


## CreateOrgSite

> Site CreateOrgSite(ctx, orgId).Site(site).Execute()

createOrgSite



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
	site := *openapiclient.NewSite("Mist Office") // Site | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitesAPI.CreateOrgSite(context.Background(), orgId).Site(site).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitesAPI.CreateOrgSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSite`: Site
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitesAPI.CreateOrgSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **site** | [**Site**](Site.md) | Request Body | 

### Return type

[**Site**](Site.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSites

> []Site ListOrgSites(ctx, orgId).Limit(limit).Page(page).Execute()

listOrgSites



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
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitesAPI.ListOrgSites(context.Background(), orgId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitesAPI.ListOrgSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSites`: []Site
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitesAPI.ListOrgSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**[]Site**](Site.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgSites

> ResponseSiteSearch SearchOrgSites(ctx, orgId).AnalyticEnabled(analyticEnabled).AppWaking(appWaking).AssetEnabled(assetEnabled).AutoUpgradeEnabled(autoUpgradeEnabled).AutoUpgradeVersion(autoUpgradeVersion).CountryCode(countryCode).HoneypotEnabled(honeypotEnabled).Id(id).LocateUnconnected(locateUnconnected).MeshEnabled(meshEnabled).Name(name).RogueEnabled(rogueEnabled).RemoteSyslogEnabled(remoteSyslogEnabled).RtsaEnabled(rtsaEnabled).VnaEnabled(vnaEnabled).WifiEnabled(wifiEnabled).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgSites



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
	analyticEnabled := true // bool | if Advanced Analytic feature is enabled (optional)
	appWaking := true // bool | if App Waking feature is enabled (optional)
	assetEnabled := true // bool | if Asset Tracking is enabled (optional)
	autoUpgradeEnabled := true // bool | if Auto Upgrade feature is enabled (optional)
	autoUpgradeVersion := "autoUpgradeVersion_example" // string | if Auto Upgrade feature is enabled (optional)
	countryCode := "countryCode_example" // string | site country code (optional)
	honeypotEnabled := true // bool | if Honeypot detection is enabled (optional)
	id := "id_example" // string | site id (optional)
	locateUnconnected := true // bool | if unconnected client are located (optional)
	meshEnabled := true // bool | if Mesh feature is enabled (optional)
	name := "name_example" // string | site name (optional)
	rogueEnabled := true // bool | if Rogue detection is enabled (optional)
	remoteSyslogEnabled := true // bool | if Remote Syslog is enabled (optional)
	rtsaEnabled := true // bool | if managed mobility feature is enabled (optional)
	vnaEnabled := true // bool | if Virtual Network Assistant is enabled (optional)
	wifiEnabled := true // bool | if WIFI feature is enabled (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSitesAPI.SearchOrgSites(context.Background(), orgId).AnalyticEnabled(analyticEnabled).AppWaking(appWaking).AssetEnabled(assetEnabled).AutoUpgradeEnabled(autoUpgradeEnabled).AutoUpgradeVersion(autoUpgradeVersion).CountryCode(countryCode).HoneypotEnabled(honeypotEnabled).Id(id).LocateUnconnected(locateUnconnected).MeshEnabled(meshEnabled).Name(name).RogueEnabled(rogueEnabled).RemoteSyslogEnabled(remoteSyslogEnabled).RtsaEnabled(rtsaEnabled).VnaEnabled(vnaEnabled).WifiEnabled(wifiEnabled).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSitesAPI.SearchOrgSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgSites`: ResponseSiteSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsSitesAPI.SearchOrgSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **analyticEnabled** | **bool** | if Advanced Analytic feature is enabled | 
 **appWaking** | **bool** | if App Waking feature is enabled | 
 **assetEnabled** | **bool** | if Asset Tracking is enabled | 
 **autoUpgradeEnabled** | **bool** | if Auto Upgrade feature is enabled | 
 **autoUpgradeVersion** | **string** | if Auto Upgrade feature is enabled | 
 **countryCode** | **string** | site country code | 
 **honeypotEnabled** | **bool** | if Honeypot detection is enabled | 
 **id** | **string** | site id | 
 **locateUnconnected** | **bool** | if unconnected client are located | 
 **meshEnabled** | **bool** | if Mesh feature is enabled | 
 **name** | **string** | site name | 
 **rogueEnabled** | **bool** | if Rogue detection is enabled | 
 **remoteSyslogEnabled** | **bool** | if Remote Syslog is enabled | 
 **rtsaEnabled** | **bool** | if managed mobility feature is enabled | 
 **vnaEnabled** | **bool** | if Virtual Network Assistant is enabled | 
 **wifiEnabled** | **bool** | if WIFI feature is enabled | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseSiteSearch**](ResponseSiteSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

