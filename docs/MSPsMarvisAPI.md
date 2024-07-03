# \MSPsMarvisAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountMspsMarvisActions**](MSPsMarvisAPI.md#CountMspsMarvisActions) | **Get** /api/v1/msps/{msp_id}/suggestion/count | countMspsMarvisActions



## CountMspsMarvisActions

> ResponseCountMarvisActions CountMspsMarvisActions(ctx, mspId).Distinct(distinct).Limit(limit).Page(page).Execute()

countMspsMarvisActions



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
	distinct := openapiclient.msp_marvis_suggestions_count_distinct("") // MspMarvisSuggestionsCountDistinct |  (optional) (default to "org_id")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsMarvisAPI.CountMspsMarvisActions(context.Background(), mspId).Distinct(distinct).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsMarvisAPI.CountMspsMarvisActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountMspsMarvisActions`: ResponseCountMarvisActions
	fmt.Fprintf(os.Stdout, "Response from `MSPsMarvisAPI.CountMspsMarvisActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountMspsMarvisActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**MspMarvisSuggestionsCountDistinct**](MspMarvisSuggestionsCountDistinct.md) |  | [default to &quot;org_id&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**ResponseCountMarvisActions**](ResponseCountMarvisActions.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

