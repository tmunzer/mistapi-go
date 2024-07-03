# \MSPsSSORolesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMspSsoRole**](MSPsSSORolesAPI.md#CreateMspSsoRole) | **Post** /api/v1/msps/{msp_id}/ssoroles | createMspSsoRole
[**DeleteMspSsoRole**](MSPsSSORolesAPI.md#DeleteMspSsoRole) | **Delete** /api/v1/msps/{msp_id}/ssoroles/{ssorole_id} | deleteMspSsoRole
[**ListMspSsoRoles**](MSPsSSORolesAPI.md#ListMspSsoRoles) | **Get** /api/v1/msps/{msp_id}/ssoroles | listMspSsoRoles
[**UpdateMspSsoRole**](MSPsSSORolesAPI.md#UpdateMspSsoRole) | **Put** /api/v1/msps/{msp_id}/ssoroles/{ssorole_id} | updateMspSsoRole



## CreateMspSsoRole

> SsoRoleMsp CreateMspSsoRole(ctx, mspId).SsoRoleMsp(ssoRoleMsp).Execute()

createMspSsoRole



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
	ssoRoleMsp := *openapiclient.NewSsoRoleMsp("Name_example", []openapiclient.PrivilegeMsp{*openapiclient.NewPrivilegeMsp(openapiclient.privilege_msp_role(""), openapiclient.privilege_msp_scope(""))}) // SsoRoleMsp | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSORolesAPI.CreateMspSsoRole(context.Background(), mspId).SsoRoleMsp(ssoRoleMsp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSORolesAPI.CreateMspSsoRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMspSsoRole`: SsoRoleMsp
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSORolesAPI.CreateMspSsoRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMspSsoRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssoRoleMsp** | [**SsoRoleMsp**](SsoRoleMsp.md) | Request Body | 

### Return type

[**SsoRoleMsp**](SsoRoleMsp.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMspSsoRole

> DeleteMspSsoRole(ctx, mspId, ssoroleId).Execute()

deleteMspSsoRole



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
	ssoroleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsSSORolesAPI.DeleteMspSsoRole(context.Background(), mspId, ssoroleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSORolesAPI.DeleteMspSsoRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoroleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMspSsoRoleRequest struct via the builder pattern


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


## ListMspSsoRoles

> []SsoRoleOrg ListMspSsoRoles(ctx, mspId).Execute()

listMspSsoRoles



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSORolesAPI.ListMspSsoRoles(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSORolesAPI.ListMspSsoRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspSsoRoles`: []SsoRoleOrg
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSORolesAPI.ListMspSsoRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspSsoRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SsoRoleOrg**](SsoRoleOrg.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMspSsoRole

> SsoRoleMsp UpdateMspSsoRole(ctx, mspId, ssoroleId).SsoRoleMsp(ssoRoleMsp).Execute()

updateMspSsoRole



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
	ssoroleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoRoleMsp := *openapiclient.NewSsoRoleMsp("Name_example", []openapiclient.PrivilegeMsp{*openapiclient.NewPrivilegeMsp(openapiclient.privilege_msp_role(""), openapiclient.privilege_msp_scope(""))}) // SsoRoleMsp | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsSSORolesAPI.UpdateMspSsoRole(context.Background(), mspId, ssoroleId).SsoRoleMsp(ssoRoleMsp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsSSORolesAPI.UpdateMspSsoRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMspSsoRole`: SsoRoleMsp
	fmt.Fprintf(os.Stdout, "Response from `MSPsSSORolesAPI.UpdateMspSsoRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**ssoroleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMspSsoRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ssoRoleMsp** | [**SsoRoleMsp**](SsoRoleMsp.md) | Request Body | 

### Return type

[**SsoRoleMsp**](SsoRoleMsp.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

