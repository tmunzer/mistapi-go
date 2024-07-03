# \OrgsVPNsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgPeerPathStats**](OrgsVPNsAPI.md#CountOrgPeerPathStats) | **Get** /api/v1/orgs/{org_id}/stats/vpn_peers/count | countOrgPeerPathStats
[**CreateOrgVpns**](OrgsVPNsAPI.md#CreateOrgVpns) | **Post** /api/v1/orgs/{org_id}/vpns | createOrgVpns
[**DeleteOrgVpn**](OrgsVPNsAPI.md#DeleteOrgVpn) | **Delete** /api/v1/orgs/{org_id}/vpns/{vpn_id} | deleteOrgVpn
[**GetOrgVpn**](OrgsVPNsAPI.md#GetOrgVpn) | **Get** /api/v1/orgs/{org_id}/vpns/{vpn_id} | getOrgVpn
[**ListOrgsVpns**](OrgsVPNsAPI.md#ListOrgsVpns) | **Get** /api/v1/orgs/{org_id}/vpns | listOrgsVpns
[**SearchOrgPeerPathStats**](OrgsVPNsAPI.md#SearchOrgPeerPathStats) | **Get** /api/v1/orgs/{org_id}/stats/vpn_peers/search | searchOrgPeerPathStats
[**UpdateOrgVpn**](OrgsVPNsAPI.md#UpdateOrgVpn) | **Put** /api/v1/orgs/{org_id}/vpns/{vpn_id} | updateOrgVpn



## CountOrgPeerPathStats

> RepsonseCount CountOrgPeerPathStats(ctx, orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgPeerPathStats



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
	distinct := "distinct_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsVPNsAPI.CountOrgPeerPathStats(context.Background(), orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVPNsAPI.CountOrgPeerPathStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgPeerPathStats`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsVPNsAPI.CountOrgPeerPathStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgPeerPathStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | **string** |  | 
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


## CreateOrgVpns

> Vpn CreateOrgVpns(ctx, orgId).Vpn(vpn).Execute()

createOrgVpns



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
	vpn := *openapiclient.NewVpn("Name_example", map[string]VpnPath{"key": *openapiclient.NewVpnPath()}) // Vpn |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsVPNsAPI.CreateOrgVpns(context.Background(), orgId).Vpn(vpn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVPNsAPI.CreateOrgVpns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgVpns`: Vpn
	fmt.Fprintf(os.Stdout, "Response from `OrgsVPNsAPI.CreateOrgVpns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vpn** | [**Vpn**](Vpn.md) |  | 

### Return type

[**Vpn**](Vpn.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgVpn

> DeleteOrgVpn(ctx, orgId, vpnId).Execute()

deleteOrgVpn



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
	vpnId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsVPNsAPI.DeleteOrgVpn(context.Background(), orgId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVPNsAPI.DeleteOrgVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**vpnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgVpnRequest struct via the builder pattern


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


## GetOrgVpn

> Vpn GetOrgVpn(ctx, orgId, vpnId).Execute()

getOrgVpn



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
	vpnId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsVPNsAPI.GetOrgVpn(context.Background(), orgId, vpnId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVPNsAPI.GetOrgVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgVpn`: Vpn
	fmt.Fprintf(os.Stdout, "Response from `OrgsVPNsAPI.GetOrgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**vpnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Vpn**](Vpn.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgsVpns

> []Vpn ListOrgsVpns(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgsVpns



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
	resp, r, err := apiClient.OrgsVPNsAPI.ListOrgsVpns(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVPNsAPI.ListOrgsVpns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgsVpns`: []Vpn
	fmt.Fprintf(os.Stdout, "Response from `OrgsVPNsAPI.ListOrgsVpns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgsVpnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Vpn**](Vpn.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgPeerPathStats

> VpnPeerStatSearch SearchOrgPeerPathStats(ctx, orgId).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchOrgPeerPathStats



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsVPNsAPI.SearchOrgPeerPathStats(context.Background(), orgId).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVPNsAPI.SearchOrgPeerPathStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgPeerPathStats`: VpnPeerStatSearch
	fmt.Fprintf(os.Stdout, "Response from `OrgsVPNsAPI.SearchOrgPeerPathStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgPeerPathStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**VpnPeerStatSearch**](VpnPeerStatSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgVpn

> Vpn UpdateOrgVpn(ctx, orgId, vpnId).Vpn(vpn).Execute()

updateOrgVpn



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
	vpnId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	vpn := *openapiclient.NewVpn("Name_example", map[string]VpnPath{"key": *openapiclient.NewVpnPath()}) // Vpn |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsVPNsAPI.UpdateOrgVpn(context.Background(), orgId, vpnId).Vpn(vpn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsVPNsAPI.UpdateOrgVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgVpn`: Vpn
	fmt.Fprintf(os.Stdout, "Response from `OrgsVPNsAPI.UpdateOrgVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**vpnId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vpn** | [**Vpn**](Vpn.md) |  | 

### Return type

[**Vpn**](Vpn.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

