# \OrgsLicensesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimOrgLicense**](OrgsLicensesAPI.md#ClaimOrgLicense) | **Post** /api/v1/orgs/{org_id}/claim | claimOrgLicense
[**GetOrgLicencesBySite**](OrgsLicensesAPI.md#GetOrgLicencesBySite) | **Get** /api/v1/orgs/{org_id}/licenses/usages | getOrgLicencesBySite
[**GetOrgLicencesSummary**](OrgsLicensesAPI.md#GetOrgLicencesSummary) | **Get** /api/v1/orgs/{org_id}/licenses | getOrgLicencesSummary
[**MoveOrDeleteOrgLicenseToAnotherOrg**](OrgsLicensesAPI.md#MoveOrDeleteOrgLicenseToAnotherOrg) | **Put** /api/v1/orgs/{org_id}/licenses | moveOrDeleteOrgLicenseToAnotherOrg



## ClaimOrgLicense

> ResponseClaimLicense ClaimOrgLicense(ctx, orgId).ClaimActivation(claimActivation).Execute()

claimOrgLicense



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
	claimActivation := *openapiclient.NewClaimActivation("Code_example", openapiclient.claim_type("")) // ClaimActivation | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsLicensesAPI.ClaimOrgLicense(context.Background(), orgId).ClaimActivation(claimActivation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLicensesAPI.ClaimOrgLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimOrgLicense`: ResponseClaimLicense
	fmt.Fprintf(os.Stdout, "Response from `OrgsLicensesAPI.ClaimOrgLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimOrgLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **claimActivation** | [**ClaimActivation**](ClaimActivation.md) | Request Body | 

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


## GetOrgLicencesBySite

> []LicenseUsageOrg GetOrgLicencesBySite(ctx, orgId).Execute()

getOrgLicencesBySite



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsLicensesAPI.GetOrgLicencesBySite(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLicensesAPI.GetOrgLicencesBySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgLicencesBySite`: []LicenseUsageOrg
	fmt.Fprintf(os.Stdout, "Response from `OrgsLicensesAPI.GetOrgLicencesBySite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgLicencesBySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]LicenseUsageOrg**](LicenseUsageOrg.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgLicencesSummary

> License GetOrgLicencesSummary(ctx, orgId).Execute()

getOrgLicencesSummary



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsLicensesAPI.GetOrgLicencesSummary(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLicensesAPI.GetOrgLicencesSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgLicencesSummary`: License
	fmt.Fprintf(os.Stdout, "Response from `OrgsLicensesAPI.GetOrgLicencesSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgLicencesSummaryRequest struct via the builder pattern


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


## MoveOrDeleteOrgLicenseToAnotherOrg

> MoveOrDeleteOrgLicenseToAnotherOrg(ctx, orgId).OrgLicenseAction(orgLicenseAction).Execute()

moveOrDeleteOrgLicenseToAnotherOrg



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
	orgLicenseAction := *openapiclient.NewOrgLicenseAction(openapiclient.org_license_action_operation("")) // OrgLicenseAction | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsLicensesAPI.MoveOrDeleteOrgLicenseToAnotherOrg(context.Background(), orgId).OrgLicenseAction(orgLicenseAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLicensesAPI.MoveOrDeleteOrgLicenseToAnotherOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrDeleteOrgLicenseToAnotherOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgLicenseAction** | [**OrgLicenseAction**](OrgLicenseAction.md) | Request Body | 

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

