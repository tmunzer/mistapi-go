# \SitesVPNsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSiteVpnsDerived**](SitesVPNsAPI.md#ListSiteVpnsDerived) | **Get** /api/v1/sites/{site_id}/vpns/derived | listSiteVpnsDerived



## ListSiteVpnsDerived

> []Vpn ListSiteVpnsDerived(ctx, siteId).Resolve(resolve).Execute()

listSiteVpnsDerived



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
	resolve := true // bool | whether resolve the site variables (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesVPNsAPI.ListSiteVpnsDerived(context.Background(), siteId).Resolve(resolve).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesVPNsAPI.ListSiteVpnsDerived``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteVpnsDerived`: []Vpn
	fmt.Fprintf(os.Stdout, "Response from `SitesVPNsAPI.ListSiteVpnsDerived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteVpnsDerivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolve** | **bool** | whether resolve the site variables | [default to false]

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

