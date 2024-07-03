# \SitesWANUsagesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWanUsage**](SitesWANUsagesAPI.md#CountSiteWanUsage) | **Get** /api/v1/sites/{site_id}/wan_usages/count | countSiteWanUsage
[**SearchSiteWanUsage**](SitesWANUsagesAPI.md#SearchSiteWanUsage) | **Get** /api/v1/sites/{site_id}/wan_usages/search | searchSiteWanUsage



## CountSiteWanUsage

> RepsonseCount CountSiteWanUsage(ctx, siteId).Mac(mac).PeerMac(peerMac).PortId(portId).PeerPortId(peerPortId).Policy(policy).Tenant(tenant).PathType(pathType).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteWanUsage



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
	mac := "mac_example" // string | MAC address (optional)
	peerMac := "peerMac_example" // string | Peer MAC address (optional)
	portId := "portId_example" // string | Port ID for the device (optional)
	peerPortId := "peerPortId_example" // string | Peer Port ID for the device (optional)
	policy := "policy_example" // string | policy for the wan path (optional)
	tenant := "tenant_example" // string | tenant network in which the packet is sent (optional)
	pathType := "pathType_example" // string | path_type of the port (optional)
	distinct := openapiclient.wan_usages_count_disctinct("") // WanUsagesCountDisctinct |  (optional) (default to "policy")
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWANUsagesAPI.CountSiteWanUsage(context.Background(), siteId).Mac(mac).PeerMac(peerMac).PortId(portId).PeerPortId(peerPortId).Policy(policy).Tenant(tenant).PathType(pathType).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWANUsagesAPI.CountSiteWanUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWanUsage`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesWANUsagesAPI.CountSiteWanUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWanUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | MAC address | 
 **peerMac** | **string** | Peer MAC address | 
 **portId** | **string** | Port ID for the device | 
 **peerPortId** | **string** | Peer Port ID for the device | 
 **policy** | **string** | policy for the wan path | 
 **tenant** | **string** | tenant network in which the packet is sent | 
 **pathType** | **string** | path_type of the port | 
 **distinct** | [**WanUsagesCountDisctinct**](WanUsagesCountDisctinct.md) |  | [default to &quot;policy&quot;]
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


## SearchSiteWanUsage

> SearchWanUsage SearchSiteWanUsage(ctx, siteId).Mac(mac).PeerMac(peerMac).PortId(portId).PeerPortId(peerPortId).Policy(policy).Tenant(tenant).PathType(pathType).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteWanUsage



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
	mac := "mac_example" // string | MAC address (optional)
	peerMac := "peerMac_example" // string | Peer MAC address (optional)
	portId := "portId_example" // string | Port ID for the device (optional)
	peerPortId := "peerPortId_example" // string | Peer Port ID for the device (optional)
	policy := "policy_example" // string | policy for the wan path (optional)
	tenant := "tenant_example" // string | tenant network in which the packet is sent (optional)
	pathType := "pathType_example" // string | path_type of the port (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWANUsagesAPI.SearchSiteWanUsage(context.Background(), siteId).Mac(mac).PeerMac(peerMac).PortId(portId).PeerPortId(peerPortId).Policy(policy).Tenant(tenant).PathType(pathType).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWANUsagesAPI.SearchSiteWanUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWanUsage`: SearchWanUsage
	fmt.Fprintf(os.Stdout, "Response from `SitesWANUsagesAPI.SearchSiteWanUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWanUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | MAC address | 
 **peerMac** | **string** | Peer MAC address | 
 **portId** | **string** | Port ID for the device | 
 **peerPortId** | **string** | Peer Port ID for the device | 
 **policy** | **string** | policy for the wan path | 
 **tenant** | **string** | tenant network in which the packet is sent | 
 **pathType** | **string** | path_type of the port | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**SearchWanUsage**](SearchWanUsage.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

