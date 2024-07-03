# \OrgsClientsMarvisAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgMarvisClientInvites**](OrgsClientsMarvisAPI.md#CreateOrgMarvisClientInvites) | **Post** /api/v1/orgs/{org_id}/marvisinvites | createOrgMarvisClientInvites
[**DeleteOrgMarvisClientInvite**](OrgsClientsMarvisAPI.md#DeleteOrgMarvisClientInvite) | **Delete** /api/v1/orgs/{org_id}/marvisinvites/{marvisinvite_id} | deleteOrgMarvisClientInvite
[**GetOrgMarvisClientInvites**](OrgsClientsMarvisAPI.md#GetOrgMarvisClientInvites) | **Get** /api/v1/orgs/{org_id}/marvisinvites/{marvisinvite_id} | getOrgMarvisClientInvites
[**ListOrgMarvisClientInvites**](OrgsClientsMarvisAPI.md#ListOrgMarvisClientInvites) | **Get** /api/v1/orgs/{org_id}/marvisinvites | listOrgMarvisClientInvites
[**UpdateOrgMarvisClientInvite**](OrgsClientsMarvisAPI.md#UpdateOrgMarvisClientInvite) | **Put** /api/v1/orgs/{org_id}/marvisinvites/{marvisinvite_id} | updateOrgMarvisClientInvites



## CreateOrgMarvisClientInvites

> MarvisClient CreateOrgMarvisClientInvites(ctx, orgId).MarvisClient(marvisClient).Execute()

createOrgMarvisClientInvites



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
	marvisClient := *openapiclient.NewMarvisClient() // MarvisClient |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsMarvisAPI.CreateOrgMarvisClientInvites(context.Background(), orgId).MarvisClient(marvisClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsMarvisAPI.CreateOrgMarvisClientInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgMarvisClientInvites`: MarvisClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsMarvisAPI.CreateOrgMarvisClientInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgMarvisClientInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marvisClient** | [**MarvisClient**](MarvisClient.md) |  | 

### Return type

[**MarvisClient**](MarvisClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgMarvisClientInvite

> DeleteOrgMarvisClientInvite(ctx, orgId, marvisinviteId).Execute()

deleteOrgMarvisClientInvite



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
	marvisinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsClientsMarvisAPI.DeleteOrgMarvisClientInvite(context.Background(), orgId, marvisinviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsMarvisAPI.DeleteOrgMarvisClientInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**marvisinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMarvisClientInviteRequest struct via the builder pattern


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


## GetOrgMarvisClientInvites

> MarvisClient GetOrgMarvisClientInvites(ctx, orgId, marvisinviteId).Execute()

getOrgMarvisClientInvites



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
	marvisinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsMarvisAPI.GetOrgMarvisClientInvites(context.Background(), orgId, marvisinviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsMarvisAPI.GetOrgMarvisClientInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMarvisClientInvites`: MarvisClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsMarvisAPI.GetOrgMarvisClientInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**marvisinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMarvisClientInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MarvisClient**](MarvisClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMarvisClientInvites

> []MarvisClient ListOrgMarvisClientInvites(ctx, orgId).Execute()

listOrgMarvisClientInvites



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
	resp, r, err := apiClient.OrgsClientsMarvisAPI.ListOrgMarvisClientInvites(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsMarvisAPI.ListOrgMarvisClientInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgMarvisClientInvites`: []MarvisClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsMarvisAPI.ListOrgMarvisClientInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMarvisClientInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MarvisClient**](MarvisClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMarvisClientInvite

> MarvisClient UpdateOrgMarvisClientInvite(ctx, orgId, marvisinviteId).MarvisClient(marvisClient).Execute()

updateOrgMarvisClientInvites



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
	marvisinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	marvisClient := *openapiclient.NewMarvisClient() // MarvisClient |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsClientsMarvisAPI.UpdateOrgMarvisClientInvite(context.Background(), orgId, marvisinviteId).MarvisClient(marvisClient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsClientsMarvisAPI.UpdateOrgMarvisClientInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgMarvisClientInvite`: MarvisClient
	fmt.Fprintf(os.Stdout, "Response from `OrgsClientsMarvisAPI.UpdateOrgMarvisClientInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**marvisinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMarvisClientInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **marvisClient** | [**MarvisClient**](MarvisClient.md) |  | 

### Return type

[**MarvisClient**](MarvisClient.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

