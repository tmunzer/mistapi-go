# \MSPsOrgsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspOrg**](MSPsOrgsAPI.md#CreateMspOrg) | **Post** /api/v1/msps/{msp_id}/orgs | createMspOrg
[**DeleteMspOrg**](MSPsOrgsAPI.md#DeleteMspOrg) | **Delete** /api/v1/msps/{msp_id}/orgs/{org_id} | deleteMspOrg
[**GetMspOrg**](MSPsOrgsAPI.md#GetMspOrg) | **Get** /api/v1/msps/{msp_id}/orgs/{org_id} | getMspOrg
[**ListMspOrgStats**](MSPsOrgsAPI.md#ListMspOrgStats) | **Get** /api/v1/msps/{msp_id}/stats/orgs | listMspOrgStats
[**ListMspOrgs**](MSPsOrgsAPI.md#ListMspOrgs) | **Get** /api/v1/msps/{msp_id}/orgs | listMspOrgs
[**ManageMspOrgs**](MSPsOrgsAPI.md#ManageMspOrgs) | **Put** /api/v1/msps/{msp_id}/orgs | manageMspOrgs
[**SearchMspOrgs**](MSPsOrgsAPI.md#SearchMspOrgs) | **Get** /api/v1/msps/{msp_id}/orgs/search | searchMspOrgs
[**UpdateMspOrg**](MSPsOrgsAPI.md#UpdateMspOrg) | **Put** /api/v1/msps/{msp_id}/orgs/{org_id} | updateMspOrg



## CreateMspOrg

> Org CreateMspOrg(ctx, mspId).Org(org).Execute()

createMspOrg



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
	org := *openapiclient.NewOrg("Org") // Org | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgsAPI.CreateMspOrg(context.Background(), mspId).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.CreateMspOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMspOrg`: Org
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgsAPI.CreateMspOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMspOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **org** | [**Org**](Org.md) | Request Body | 

### Return type

[**Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMspOrg

> DeleteMspOrg(ctx, mspId, orgId).Execute()

deleteMspOrg



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsOrgsAPI.DeleteMspOrg(context.Background(), mspId, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.DeleteMspOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMspOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMspOrg

> Org GetMspOrg(ctx, mspId, orgId).Execute()

getMspOrg



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgsAPI.GetMspOrg(context.Background(), mspId, orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.GetMspOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOrg`: Org
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgsAPI.GetMspOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMspOrgStats

> []OrgStats ListMspOrgStats(ctx, mspId).Page(page).Limit(limit).Execute()

listMspOrgStats



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgsAPI.ListMspOrgStats(context.Background(), mspId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.ListMspOrgStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspOrgStats`: []OrgStats
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgsAPI.ListMspOrgStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspOrgStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]OrgStats**](OrgStats.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMspOrgs

> []Org ListMspOrgs(ctx, mspId).Execute()

listMspOrgs



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
	resp, r, err := apiClient.MSPsOrgsAPI.ListMspOrgs(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.ListMspOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspOrgs`: []Org
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgsAPI.ListMspOrgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ManageMspOrgs

> ManageMspOrgs(ctx, mspId).MspOrgChange(mspOrgChange).Execute()

manageMspOrgs



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
	mspOrgChange := *openapiclient.NewMspOrgChange(openapiclient.msp_org_change_operation(""), []string{"OrgIds_example"}) // MspOrgChange | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsOrgsAPI.ManageMspOrgs(context.Background(), mspId).MspOrgChange(mspOrgChange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.ManageMspOrgs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiManageMspOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mspOrgChange** | [**MspOrgChange**](MspOrgChange.md) | Request Body | 

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


## SearchMspOrgs

> ResponseOrgSearch SearchMspOrgs(ctx, mspId).Name(name).OrgId(orgId).SubInsufficient(subInsufficient).TrialEnabled(trialEnabled).UsageTypes(usageTypes).Limit(limit).Execute()

searchMspOrgs



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
	name := "name_example" // string |  (optional) (default to "")
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | org id (optional)
	subInsufficient := true // bool | if this org has sufficient subscription (optional)
	trialEnabled := true // bool | if this org is under trial period (optional)
	usageTypes := []string{"Inner_example"} // []string | a list of types that enabled by usage (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgsAPI.SearchMspOrgs(context.Background(), mspId).Name(name).OrgId(orgId).SubInsufficient(subInsufficient).TrialEnabled(trialEnabled).UsageTypes(usageTypes).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.SearchMspOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMspOrgs`: ResponseOrgSearch
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgsAPI.SearchMspOrgs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchMspOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | [default to &quot;&quot;]
 **orgId** | **string** | org id | 
 **subInsufficient** | **bool** | if this org has sufficient subscription | 
 **trialEnabled** | **bool** | if this org is under trial period | 
 **usageTypes** | **[]string** | a list of types that enabled by usage | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**ResponseOrgSearch**](ResponseOrgSearch.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMspOrg

> Org UpdateMspOrg(ctx, mspId, orgId).Org(org).Execute()

updateMspOrg



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	org := *openapiclient.NewOrg("Org") // Org |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgsAPI.UpdateMspOrg(context.Background(), mspId, orgId).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgsAPI.UpdateMspOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMspOrg`: Org
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgsAPI.UpdateMspOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMspOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **org** | [**Org**](Org.md) |  | 

### Return type

[**Org**](Org.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

