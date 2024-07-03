# \MSPsTicketsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountMspTickets**](MSPsTicketsAPI.md#CountMspTickets) | **Get** /api/v1/msps/{msp_id}/tickets/count | countMspTickets
[**ListMspTickets**](MSPsTicketsAPI.md#ListMspTickets) | **Get** /api/v1/msps/{msp_id}/tickets | listMspTickets



## CountMspTickets

> RepsonseCount CountMspTickets(ctx, mspId).Distinct(distinct).Execute()

countMspTickets



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
	distinct := openapiclient.msp_tickets_count_distinct("") // MspTicketsCountDistinct |  (optional) (default to "status")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsTicketsAPI.CountMspTickets(context.Background(), mspId).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsTicketsAPI.CountMspTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountMspTickets`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `MSPsTicketsAPI.CountMspTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountMspTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**MspTicketsCountDistinct**](MspTicketsCountDistinct.md) |  | [default to &quot;status&quot;]

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


## ListMspTickets

> []Ticket ListMspTickets(ctx, mspId).Start(start).End(end).Duration(duration).Execute()

listMspTickets



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
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsTicketsAPI.ListMspTickets(context.Background(), mspId).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsTicketsAPI.ListMspTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspTickets`: []Ticket
	fmt.Fprintf(os.Stdout, "Response from `MSPsTicketsAPI.ListMspTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]Ticket**](Ticket.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

