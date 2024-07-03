# \SelfAlarmsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAlarmSubscriptions**](SelfAlarmsAPI.md#ListAlarmSubscriptions) | **Get** /api/v1/self/subscriptions | listAlarmSubscriptions



## ListAlarmSubscriptions

> []ResponseSelfSubscription ListAlarmSubscriptions(ctx).Execute()

listAlarmSubscriptions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelfAlarmsAPI.ListAlarmSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelfAlarmsAPI.ListAlarmSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAlarmSubscriptions`: []ResponseSelfSubscription
	fmt.Fprintf(os.Stdout, "Response from `SelfAlarmsAPI.ListAlarmSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAlarmSubscriptionsRequest struct via the builder pattern


### Return type

[**[]ResponseSelfSubscription**](ResponseSelfSubscription.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

