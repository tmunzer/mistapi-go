# \OrgsMarvisAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TroubleshootOrg**](OrgsMarvisAPI.md#TroubleshootOrg) | **Get** /api/v1/orgs/{org_id}/troubleshoot | troubleshootOrg



## TroubleshootOrg

> ResponseTroubleshoot TroubleshootOrg(ctx, orgId).Mac(mac).SiteId(siteId).Start(start).End(end).Type_(type_).Execute()

troubleshootOrg



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
	mac := "mac_example" // string | **required** when troubleshooting device or a client (optional)
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | **required** when troubleshooting site (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	type_ := openapiclient.troubleshoot_type("") // TroubleshootType | when troubleshooting site, type of network to troubleshoot (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMarvisAPI.TroubleshootOrg(context.Background(), orgId).Mac(mac).SiteId(siteId).Start(start).End(end).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMarvisAPI.TroubleshootOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TroubleshootOrg`: ResponseTroubleshoot
	fmt.Fprintf(os.Stdout, "Response from `OrgsMarvisAPI.TroubleshootOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTroubleshootOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mac** | **string** | **required** when troubleshooting device or a client | 
 **siteId** | **string** | **required** when troubleshooting site | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **type_** | [**TroubleshootType**](TroubleshootType.md) | when troubleshooting site, type of network to troubleshoot | 

### Return type

[**ResponseTroubleshoot**](ResponseTroubleshoot.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

