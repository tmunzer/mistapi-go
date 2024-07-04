# \OrgsAdminsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InviteOrgAdmin**](OrgsAdminsAPI.md#InviteOrgAdmin) | **Post** /api/v1/orgs/{org_id}/invites | inviteOrgAdmin
[**ListOrgAdmins**](OrgsAdminsAPI.md#ListOrgAdmins) | **Get** /api/v1/orgs/{org_id}/admins | listOrgAdmins
[**RevokeOrgAdmin**](OrgsAdminsAPI.md#RevokeOrgAdmin) | **Delete** /api/v1/orgs/{org_id}/admins/{admin_id} | revokeOrgAdmin
[**UninviteOrgAdmin**](OrgsAdminsAPI.md#UninviteOrgAdmin) | **Delete** /api/v1/orgs/{org_id}/invites/{invite_id} | uninviteOrgAdmin
[**UpdateOrgAdmin**](OrgsAdminsAPI.md#UpdateOrgAdmin) | **Put** /api/v1/orgs/{org_id}/admins/{admin_id} | updateOrgAdmin
[**UpdateOrgAdminInvite**](OrgsAdminsAPI.md#UpdateOrgAdminInvite) | **Put** /api/v1/orgs/{org_id}/invites/{invite_id} | updateOrgAdminInvite



## InviteOrgAdmin

> InviteOrgAdmin(ctx, orgId).Admin(admin).Execute()

inviteOrgAdmin



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
	admin := *openapiclient.NewAdmin("jsnow@abc.com", "John", "Sno") // Admin | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAdminsAPI.InviteOrgAdmin(context.Background(), orgId).Admin(admin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAdminsAPI.InviteOrgAdmin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiInviteOrgAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **admin** | [**Admin**](Admin.md) | Request Body | 

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


## ListOrgAdmins

> []Admin ListOrgAdmins(ctx, orgId).Execute()

listOrgAdmins



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
	resp, r, err := apiClient.OrgsAdminsAPI.ListOrgAdmins(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAdminsAPI.ListOrgAdmins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAdmins`: []Admin
	fmt.Fprintf(os.Stdout, "Response from `OrgsAdminsAPI.ListOrgAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAdminsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Admin**](Admin.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOrgAdmin

> RevokeOrgAdmin(ctx, orgId, adminId).Execute()

revokeOrgAdmin



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
	adminId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAdminsAPI.RevokeOrgAdmin(context.Background(), orgId, adminId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAdminsAPI.RevokeOrgAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**adminId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOrgAdminRequest struct via the builder pattern


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


## UninviteOrgAdmin

> UninviteOrgAdmin(ctx, orgId, inviteId).Execute()

uninviteOrgAdmin



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
	inviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAdminsAPI.UninviteOrgAdmin(context.Background(), orgId, inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAdminsAPI.UninviteOrgAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**inviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninviteOrgAdminRequest struct via the builder pattern


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


## UpdateOrgAdmin

> Admin UpdateOrgAdmin(ctx, orgId, adminId).Admin(admin).Execute()

updateOrgAdmin



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
	adminId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	admin := *openapiclient.NewAdmin("jsnow@abc.com", "John", "Sno") // Admin | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAdminsAPI.UpdateOrgAdmin(context.Background(), orgId, adminId).Admin(admin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAdminsAPI.UpdateOrgAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgAdmin`: Admin
	fmt.Fprintf(os.Stdout, "Response from `OrgsAdminsAPI.UpdateOrgAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**adminId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **admin** | [**Admin**](Admin.md) | Request Body | 

### Return type

[**Admin**](Admin.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgAdminInvite

> UpdateOrgAdminInvite(ctx, orgId, inviteId).Admin(admin).Execute()

updateOrgAdminInvite



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
	inviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	admin := *openapiclient.NewAdmin("jsnow@abc.com", "John", "Sno") // Admin | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAdminsAPI.UpdateOrgAdminInvite(context.Background(), orgId, inviteId).Admin(admin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAdminsAPI.UpdateOrgAdminInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**inviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgAdminInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **admin** | [**Admin**](Admin.md) | Request Body | 

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

