# \MSPsOrgGroupsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspOrgGroup**](MSPsOrgGroupsAPI.md#CreateMspOrgGroup) | **Post** /api/v1/msps/{msp_id}/orggroups | createMspOrgGroup
[**DeleteMspOrgGroup**](MSPsOrgGroupsAPI.md#DeleteMspOrgGroup) | **Delete** /api/v1/msps/{msp_id}/orggroups/{orggroup_id} | deleteMspOrgGroup
[**GetMspOrgGroup**](MSPsOrgGroupsAPI.md#GetMspOrgGroup) | **Get** /api/v1/msps/{msp_id}/orggroups/{orggroup_id} | getMspOrgGroup
[**ListMspOrgGroups**](MSPsOrgGroupsAPI.md#ListMspOrgGroups) | **Get** /api/v1/msps/{msp_id}/orggroups | listMspOrgGroups
[**UpdateMspOrgGroup**](MSPsOrgGroupsAPI.md#UpdateMspOrgGroup) | **Put** /api/v1/msps/{msp_id}/orggroups/{orggroup_id} | updateMspOrgGroup



## CreateMspOrgGroup

> Orggroup CreateMspOrgGroup(ctx, mspId).Orggroup(orggroup).Execute()

createMspOrgGroup



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
	orggroup := *openapiclient.NewOrggroup("Name_example") // Orggroup | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgGroupsAPI.CreateMspOrgGroup(context.Background(), mspId).Orggroup(orggroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgGroupsAPI.CreateMspOrgGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMspOrgGroup`: Orggroup
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgGroupsAPI.CreateMspOrgGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMspOrgGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orggroup** | [**Orggroup**](Orggroup.md) | Request Body | 

### Return type

[**Orggroup**](Orggroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMspOrgGroup

> DeleteMspOrgGroup(ctx, mspId, orggroupId).Execute()

deleteMspOrgGroup



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
	orggroupId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsOrgGroupsAPI.DeleteMspOrgGroup(context.Background(), mspId, orggroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgGroupsAPI.DeleteMspOrgGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**orggroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMspOrgGroupRequest struct via the builder pattern


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


## GetMspOrgGroup

> Orggroup GetMspOrgGroup(ctx, mspId, orggroupId).Execute()

getMspOrgGroup



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
	orggroupId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgGroupsAPI.GetMspOrgGroup(context.Background(), mspId, orggroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgGroupsAPI.GetMspOrgGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspOrgGroup`: Orggroup
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgGroupsAPI.GetMspOrgGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**orggroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspOrgGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Orggroup**](Orggroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMspOrgGroups

> []Orggroup ListMspOrgGroups(ctx, mspId).Execute()

listMspOrgGroups



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
	resp, r, err := apiClient.MSPsOrgGroupsAPI.ListMspOrgGroups(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgGroupsAPI.ListMspOrgGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspOrgGroups`: []Orggroup
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgGroupsAPI.ListMspOrgGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspOrgGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Orggroup**](Orggroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMspOrgGroup

> Orggroup UpdateMspOrgGroup(ctx, mspId, orggroupId).Orggroup(orggroup).Execute()

updateMspOrgGroup



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
	orggroupId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	orggroup := *openapiclient.NewOrggroup("Name_example") // Orggroup | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsOrgGroupsAPI.UpdateMspOrgGroup(context.Background(), mspId, orggroupId).Orggroup(orggroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsOrgGroupsAPI.UpdateMspOrgGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMspOrgGroup`: Orggroup
	fmt.Fprintf(os.Stdout, "Response from `MSPsOrgGroupsAPI.UpdateMspOrgGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**orggroupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMspOrgGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orggroup** | [**Orggroup**](Orggroup.md) | Request Body | 

### Return type

[**Orggroup**](Orggroup.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

