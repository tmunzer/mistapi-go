# \OrgsTunnelsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgTunnelsStats**](OrgsTunnelsAPI.md#CountOrgTunnelsStats) | **Get** /api/v1/orgs/{org_id}/stats/tunnels/count | countOrgTunnelsStats
[**SearchOrgTunnelsStats**](OrgsTunnelsAPI.md#SearchOrgTunnelsStats) | **Get** /api/v1/orgs/{org_id}/stats/tunnels/search | searchOrgTunnelsStats



## CountOrgTunnelsStats

> RepsonseCount CountOrgTunnelsStats(ctx, orgId).Distinct(distinct).Type_(type_).Execute()

countOrgTunnelsStats



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
	distinct := openapiclient.org_tunnel_count_distinct("") // OrgTunnelCountDistinct | - If `type`==`wxtunnel`: wxtunnel_id / ap / remote_ip / remote_port / state / mxedge_id / mxcluster_id / site_id / peer_mxedge_id; default is wxtunnel_id  - If `type`==`wan`: mac / site_id / node / peer_ip / peer_host/ ip / tunnel_name / protocol / auth_algo / encrypt_algo / ike_version / last_event / up (optional) (default to "wxtunnel_id")
	type_ := openapiclient.org_tunnel_type_count("") // OrgTunnelTypeCount |  (optional) (default to "wxtunnel")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTunnelsAPI.CountOrgTunnelsStats(context.Background(), orgId).Distinct(distinct).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTunnelsAPI.CountOrgTunnelsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgTunnelsStats`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsTunnelsAPI.CountOrgTunnelsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgTunnelsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgTunnelCountDistinct**](OrgTunnelCountDistinct.md) | - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60;: wxtunnel_id / ap / remote_ip / remote_port / state / mxedge_id / mxcluster_id / site_id / peer_mxedge_id; default is wxtunnel_id  - If &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60;: mac / site_id / node / peer_ip / peer_host/ ip / tunnel_name / protocol / auth_algo / encrypt_algo / ike_version / last_event / up | [default to &quot;wxtunnel_id&quot;]
 **type_** | [**OrgTunnelTypeCount**](OrgTunnelTypeCount.md) |  | [default to &quot;wxtunnel&quot;]

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


## SearchOrgTunnelsStats

> ResponseTunnelSearch SearchOrgTunnelsStats(ctx, orgId).MxclusterId(mxclusterId).SiteId(siteId).WxtunnelId(wxtunnelId).Ap(ap).Mac(mac).Node(node).PeerIp(peerIp).PeerHost(peerHost).Ip(ip).TunnelName(tunnelName).Protocol(protocol).AuthAlgo(authAlgo).EncryptAlgo(encryptAlgo).IkeVersion(ikeVersion).Up(up).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchOrgTunnelsStats



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
	mxclusterId := "mxclusterId_example" // string | if `type`==`wxtunnel` (optional)
	siteId := "siteId_example" // string |  (optional)
	wxtunnelId := "wxtunnelId_example" // string | if `type`==`wxtunnel` (optional)
	ap := "ap_example" // string | if `type`==`wxtunnel` (optional)
	mac := "mac_example" // string | if `type`==`wan` (optional)
	node := "node_example" // string | if `type`==`wan` (optional)
	peerIp := "peerIp_example" // string | if `type`==`wan` (optional)
	peerHost := "peerHost_example" // string | if `type`==`wan` (optional)
	ip := "ip_example" // string | if `type`==`wan` (optional)
	tunnelName := "tunnelName_example" // string | if `type`==`wan` (optional)
	protocol := "protocol_example" // string | if `type`==`wan` (optional)
	authAlgo := "authAlgo_example" // string | if `type`==`wan` (optional)
	encryptAlgo := "encryptAlgo_example" // string | if `type`==`wan` (optional)
	ikeVersion := "ikeVersion_example" // string | if `type`==`wan` (optional)
	up := "up_example" // string | if `type`==`wan` (optional)
	type_ := openapiclient.tunnel_type("") // TunnelType |  (optional) (default to "wxtunnel")
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTunnelsAPI.SearchOrgTunnelsStats(context.Background(), orgId).MxclusterId(mxclusterId).SiteId(siteId).WxtunnelId(wxtunnelId).Ap(ap).Mac(mac).Node(node).PeerIp(peerIp).PeerHost(peerHost).Ip(ip).TunnelName(tunnelName).Protocol(protocol).AuthAlgo(authAlgo).EncryptAlgo(encryptAlgo).IkeVersion(ikeVersion).Up(up).Type_(type_).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTunnelsAPI.SearchOrgTunnelsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgTunnelsStats`: ResponseTunnelSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsTunnelsAPI.SearchOrgTunnelsStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgTunnelsStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxclusterId** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60; | 
 **siteId** | **string** |  | 
 **wxtunnelId** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60; | 
 **ap** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wxtunnel&#x60; | 
 **mac** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **node** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **peerIp** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **peerHost** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **ip** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **tunnelName** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **protocol** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **authAlgo** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **encryptAlgo** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **ikeVersion** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **up** | **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;wan&#x60; | 
 **type_** | [**TunnelType**](TunnelType.md) |  | [default to &quot;wxtunnel&quot;]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**ResponseTunnelSearch**](ResponseTunnelSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

