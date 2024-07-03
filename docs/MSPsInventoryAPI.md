# \MSPsInventoryAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspInventoryByMac**](MSPsInventoryAPI.md#GetMspInventoryByMac) | **Get** /api/v1/msps/{msp_id}/inventory/{device_mac} | getMspInventoryByMac



## GetMspInventoryByMac

> ResponseMspInventoryDevice GetMspInventoryByMac(ctx, mspId, deviceMac).Execute()

getMspInventoryByMac



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsInventoryAPI.GetMspInventoryByMac(context.Background(), mspId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsInventoryAPI.GetMspInventoryByMac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspInventoryByMac`: ResponseMspInventoryDevice
	fmt.Fprintf(os.Stdout, "Response from `MSPsInventoryAPI.GetMspInventoryByMac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspInventoryByMacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseMspInventoryDevice**](ResponseMspInventoryDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

