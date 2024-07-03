# \MSPsAdminsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMspAdmin**](MSPsAdminsAPI.md#GetMspAdmin) | **Get** /api/v1/msps/{msp_id}/admins/{admin_id} | getMspAdmin
[**InviteMspAdmin**](MSPsAdminsAPI.md#InviteMspAdmin) | **Post** /api/v1/msps/{msp_id}/invites | inviteMspAdmin
[**ListMspAdmins**](MSPsAdminsAPI.md#ListMspAdmins) | **Get** /api/v1/msps/{msp_id}/admins | listMspAdmins
[**RevokeMspAdmin**](MSPsAdminsAPI.md#RevokeMspAdmin) | **Delete** /api/v1/msps/{msp_id}/admins/{admin_id} | revokeMspAdmin
[**UninviteMspAdmin**](MSPsAdminsAPI.md#UninviteMspAdmin) | **Delete** /api/v1/msps/{msp_id}/invites/{invite_id} | uninviteMspAdmin
[**UpdateMspAdmin**](MSPsAdminsAPI.md#UpdateMspAdmin) | **Put** /api/v1/msps/{msp_id}/admins/{admin_id} | updateMspAdmin
[**UpdateMspAdminInvite**](MSPsAdminsAPI.md#UpdateMspAdminInvite) | **Put** /api/v1/msps/{msp_id}/invites/{invite_id} | updateMspAdminInvite



## GetMspAdmin

> Admin GetMspAdmin(ctx, mspId, adminId).Execute()

getMspAdmin



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
	adminId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsAdminsAPI.GetMspAdmin(context.Background(), mspId, adminId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsAdminsAPI.GetMspAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMspAdmin`: Admin
	fmt.Fprintf(os.Stdout, "Response from `MSPsAdminsAPI.GetMspAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**adminId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMspAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Admin**](Admin.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteMspAdmin

> Admin InviteMspAdmin(ctx, mspId).Admin(admin).Execute()

inviteMspAdmin



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
	admin := *openapiclient.NewAdmin("jsnow@abc.com", "John", "Sno") // Admin | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsAdminsAPI.InviteMspAdmin(context.Background(), mspId).Admin(admin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsAdminsAPI.InviteMspAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InviteMspAdmin`: Admin
	fmt.Fprintf(os.Stdout, "Response from `MSPsAdminsAPI.InviteMspAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteMspAdminRequest struct via the builder pattern


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


## ListMspAdmins

> []Admin ListMspAdmins(ctx, mspId).Execute()

listMspAdmins



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
	resp, r, err := apiClient.MSPsAdminsAPI.ListMspAdmins(context.Background(), mspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsAdminsAPI.ListMspAdmins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMspAdmins`: []Admin
	fmt.Fprintf(os.Stdout, "Response from `MSPsAdminsAPI.ListMspAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMspAdminsRequest struct via the builder pattern


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


## RevokeMspAdmin

> RevokeMspAdmin(ctx, mspId, adminId).Execute()

revokeMspAdmin



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
	adminId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsAdminsAPI.RevokeMspAdmin(context.Background(), mspId, adminId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsAdminsAPI.RevokeMspAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**adminId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeMspAdminRequest struct via the builder pattern


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


## UninviteMspAdmin

> UninviteMspAdmin(ctx, mspId, inviteId).Execute()

uninviteMspAdmin



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
	inviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MSPsAdminsAPI.UninviteMspAdmin(context.Background(), mspId, inviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsAdminsAPI.UninviteMspAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**inviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninviteMspAdminRequest struct via the builder pattern


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


## UpdateMspAdmin

> Admin UpdateMspAdmin(ctx, mspId, adminId).Admin(admin).Execute()

updateMspAdmin



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
	adminId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	admin := *openapiclient.NewAdmin("jsnow@abc.com", "John", "Sno") // Admin | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsAdminsAPI.UpdateMspAdmin(context.Background(), mspId, adminId).Admin(admin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsAdminsAPI.UpdateMspAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMspAdmin`: Admin
	fmt.Fprintf(os.Stdout, "Response from `MSPsAdminsAPI.UpdateMspAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**adminId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMspAdminRequest struct via the builder pattern


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


## UpdateMspAdminInvite

> Admin UpdateMspAdminInvite(ctx, mspId, inviteId).Admin(admin).Execute()

updateMspAdminInvite



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
	inviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	admin := *openapiclient.NewAdmin("jsnow@abc.com", "John", "Sno") // Admin | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MSPsAdminsAPI.UpdateMspAdminInvite(context.Background(), mspId, inviteId).Admin(admin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MSPsAdminsAPI.UpdateMspAdminInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMspAdminInvite`: Admin
	fmt.Fprintf(os.Stdout, "Response from `MSPsAdminsAPI.UpdateMspAdminInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mspId** | **string** |  | 
**inviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMspAdminInviteRequest struct via the builder pattern


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

