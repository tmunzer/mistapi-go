# \OrgsSSORolesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgSsoRole**](OrgsSSORolesAPI.md#CreateOrgSsoRole) | **Post** /api/v1/orgs/{org_id}/ssoroles | createOrgSsoRole
[**DeleteOrgSsoRole**](OrgsSSORolesAPI.md#DeleteOrgSsoRole) | **Delete** /api/v1/orgs/{org_id}/ssoroles/{ssorole_id} | deleteOrgSsoRole
[**GetOrgSsoRole**](OrgsSSORolesAPI.md#GetOrgSsoRole) | **Get** /api/v1/orgs/{org_id}/ssoroles/{ssorole_id} | getOrgSsoRole
[**ListOrgSsoRoles**](OrgsSSORolesAPI.md#ListOrgSsoRoles) | **Get** /api/v1/orgs/{org_id}/ssoroles | listOrgSsoRoles
[**UpdateOrgSsoRole**](OrgsSSORolesAPI.md#UpdateOrgSsoRole) | **Put** /api/v1/orgs/{org_id}/ssoroles/{ssorole_id} | updateOrgSsoRole



## CreateOrgSsoRole

> SsoRoleOrg CreateOrgSsoRole(ctx, orgId).SsoRoleOrg(ssoRoleOrg).Execute()

createOrgSsoRole



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoRoleOrg := *openapiclient.NewSsoRoleOrg("Name_example", []openapiclient.PrivilegeOrg{*openapiclient.NewPrivilegeOrg(openapiclient.privilege_org_role(""), openapiclient.privilege_org_scope(""))}) // SsoRoleOrg | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSORolesAPI.CreateOrgSsoRole(context.Background(), orgId).SsoRoleOrg(ssoRoleOrg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSORolesAPI.CreateOrgSsoRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgSsoRole`: SsoRoleOrg
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSORolesAPI.CreateOrgSsoRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSsoRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssoRoleOrg** | [**SsoRoleOrg**](SsoRoleOrg.md) | Request Body | 

### Return type

[**SsoRoleOrg**](SsoRoleOrg.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgSsoRole

> DeleteOrgSsoRole(ctx, orgId, ssoroleId).Execute()

deleteOrgSsoRole



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoroleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSSORolesAPI.DeleteOrgSsoRole(context.Background(), orgId, ssoroleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSORolesAPI.DeleteOrgSsoRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoroleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSsoRoleRequest struct via the builder pattern


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


## GetOrgSsoRole

> SsoRoleOrg GetOrgSsoRole(ctx, orgId, ssoroleId).Execute()

getOrgSsoRole



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoroleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSORolesAPI.GetOrgSsoRole(context.Background(), orgId, ssoroleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSORolesAPI.GetOrgSsoRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSsoRole`: SsoRoleOrg
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSORolesAPI.GetOrgSsoRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoroleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSsoRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SsoRoleOrg**](SsoRoleOrg.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSsoRoles

> []SsoRoleMsp ListOrgSsoRoles(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgSsoRoles



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSORolesAPI.ListOrgSsoRoles(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSORolesAPI.ListOrgSsoRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSsoRoles`: []SsoRoleMsp
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSORolesAPI.ListOrgSsoRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSsoRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]SsoRoleMsp**](SsoRoleMsp.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSsoRole

> SsoRoleOrg UpdateOrgSsoRole(ctx, orgId, ssoroleId).SsoRoleOrg(ssoRoleOrg).Execute()

updateOrgSsoRole



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
	orgId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoroleId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssoRoleOrg := *openapiclient.NewSsoRoleOrg("Name_example", []openapiclient.PrivilegeOrg{*openapiclient.NewPrivilegeOrg(openapiclient.privilege_org_role(""), openapiclient.privilege_org_scope(""))}) // SsoRoleOrg | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSSORolesAPI.UpdateOrgSsoRole(context.Background(), orgId, ssoroleId).SsoRoleOrg(ssoRoleOrg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSSORolesAPI.UpdateOrgSsoRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSsoRole`: SsoRoleOrg
	fmt.Fprintf(os.Stdout, "Response from `OrgsSSORolesAPI.UpdateOrgSsoRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ssoroleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSsoRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ssoRoleOrg** | [**SsoRoleOrg**](SsoRoleOrg.md) | Request Body | 

### Return type

[**SsoRoleOrg**](SsoRoleOrg.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

