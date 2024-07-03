# \UtilitiesMxEdgeAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreemptSitesMxTunnel**](UtilitiesMxEdgeAPI.md#PreemptSitesMxTunnel) | **Post** /api/v1/sites/{site_id}/mxtunnels/{mxtunnel_id}/preempt_aps | preemptSitesMxTunnel



## PreemptSitesMxTunnel

> ResponseMxtunnelsPreemptAps PreemptSitesMxTunnel(ctx, siteId, mxtunnelId).Execute()

preemptSitesMxTunnel



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
	mxtunnelId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesMxEdgeAPI.PreemptSitesMxTunnel(context.Background(), siteId, mxtunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesMxEdgeAPI.PreemptSitesMxTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreemptSitesMxTunnel`: ResponseMxtunnelsPreemptAps
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesMxEdgeAPI.PreemptSitesMxTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**mxtunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreemptSitesMxTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseMxtunnelsPreemptAps**](ResponseMxtunnelsPreemptAps.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

