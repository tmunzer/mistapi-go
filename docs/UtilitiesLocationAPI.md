# \UtilitiesLocationAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSiteDevicesArbitratryBleBeacon**](UtilitiesLocationAPI.md#SendSiteDevicesArbitratryBleBeacon) | **Post** /api/v1/sites/{site_id}/devices/send_ble_beacon | sendSiteDevicesArbitratryBleBeacon



## SendSiteDevicesArbitratryBleBeacon

> SendSiteDevicesArbitratryBleBeacon(ctx, siteId).UtilsSendBleBeacon(utilsSendBleBeacon).Execute()

sendSiteDevicesArbitratryBleBeacon



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
	utilsSendBleBeacon := *openapiclient.NewUtilsSendBleBeacon() // UtilsSendBleBeacon |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesLocationAPI.SendSiteDevicesArbitratryBleBeacon(context.Background(), siteId).UtilsSendBleBeacon(utilsSendBleBeacon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesLocationAPI.SendSiteDevicesArbitratryBleBeacon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSiteDevicesArbitratryBleBeaconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **utilsSendBleBeacon** | [**UtilsSendBleBeacon**](UtilsSendBleBeacon.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

