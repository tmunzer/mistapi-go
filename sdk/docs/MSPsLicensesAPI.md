# \MSPsLicensesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimMspLicence**](MSPsLicensesAPI.md#ClaimMspLicence) | **Post** /api/v1/msps/{msp_id}/claim | claimMspLicence
[**ListMspLicenses**](MSPsLicensesAPI.md#ListMspLicenses) | **Get** /api/v1/msps/{msp_id}/licenses | listMspLicenses
[**ListMspOrgLicenses**](MSPsLicensesAPI.md#ListMspOrgLicenses) | **Get** /api/v1/msps/{msp_id}/stats/licenses | listMspOrgLicenses
[**MoveOrDeleteMspLicenseToAnotherOrg**](MSPsLicensesAPI.md#MoveOrDeleteMspLicenseToAnotherOrg) | **Put** /api/v1/msps/{msp_id}/licenses | moveOrDeleteMspLicenseToAnotherOrg



## ClaimMspLicence

> ResponseClaimLicense ClaimMspLicence(ctx, mspId).CodeString(codeString).Execute()

claimMspLicence



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	codeString := *openapiclient.NewCodeString("Code_example") // CodeString |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsLicensesAPI.ClaimMspLicence(context.Background(), mspId).CodeString(codeString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsLicensesAPI.ClaimMspLicence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimMspLicence`: ResponseClaimLicense
	fmt.Fprintf(os.Stdout, "Response from `MSPsLicensesAPI.ClaimMspLicence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimMspLicenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeString** | [**CodeString**](CodeString.md) |  | 

### Return type

[**ResponseClaimLicense**](ResponseClaimLicense.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMspLicenses

> License ListMspLicenses(ctx, mspId).Execute()

listMspLicenses



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsLicensesAPI.ListMspLicenses(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsLicensesAPI.ListMspLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspLicenses`: License
	fmt.Fprintf(os.Stdout, "Response from `MSPsLicensesAPI.ListMspLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**License**](License.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMspOrgLicenses

> License ListMspOrgLicenses(ctx, mspId).Execute()

listMspOrgLicenses



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsLicensesAPI.ListMspOrgLicenses(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsLicensesAPI.ListMspOrgLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspOrgLicenses`: License
	fmt.Fprintf(os.Stdout, "Response from `MSPsLicensesAPI.ListMspOrgLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspOrgLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**License**](License.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrDeleteMspLicenseToAnotherOrg

> MoveOrDeleteMspLicenseToAnotherOrg(ctx, mspId).MspLicenseAction(mspLicenseAction).Execute()

moveOrDeleteMspLicenseToAnotherOrg



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
	mspId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mspLicenseAction := *openapiclient.NewMspLicenseAction(openapiclient.msp_license_action_operation("")) // MspLicenseAction |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsLicensesAPI.MoveOrDeleteMspLicenseToAnotherOrg(context.Background(), mspId).MspLicenseAction(mspLicenseAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsLicensesAPI.MoveOrDeleteMspLicenseToAnotherOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrDeleteMspLicenseToAnotherOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mspLicenseAction** | [**MspLicenseAction**](MspLicenseAction.md) |  | 

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

