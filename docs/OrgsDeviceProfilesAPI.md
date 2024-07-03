# \OrgsDeviceProfilesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignOrgDeviceProfile**](OrgsDeviceProfilesAPI.md#AssignOrgDeviceProfile) | **Post** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id}/assign | assignOrgDeviceProfile
[**CreateOrgDeviceProfiles**](OrgsDeviceProfilesAPI.md#CreateOrgDeviceProfiles) | **Post** /api/v1/orgs/{org_id}/deviceprofiles | createOrgDeviceProfiles
[**DeleteOrgDeviceProfile**](OrgsDeviceProfilesAPI.md#DeleteOrgDeviceProfile) | **Delete** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id} | deleteOrgDeviceProfile
[**GetOrgDeviceProfile**](OrgsDeviceProfilesAPI.md#GetOrgDeviceProfile) | **Get** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id} | getOrgDeviceProfile
[**ListOrgDeviceProfiles**](OrgsDeviceProfilesAPI.md#ListOrgDeviceProfiles) | **Get** /api/v1/orgs/{org_id}/deviceprofiles | listOrgDeviceProfiles
[**UnassignOrgDeviceProfiles**](OrgsDeviceProfilesAPI.md#UnassignOrgDeviceProfiles) | **Post** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id}/unassign | unassignOrgDeviceProfiles
[**UpdateOrgDeviceProfile**](OrgsDeviceProfilesAPI.md#UpdateOrgDeviceProfile) | **Put** /api/v1/orgs/{org_id}/deviceprofiles/{deviceprofile_id} | updateOrgDeviceProfiles



## AssignOrgDeviceProfile

> ResponseAssignSuccess AssignOrgDeviceProfile(ctx, orgId, deviceprofileId).MacAddresses(macAddresses).Execute()

assignOrgDeviceProfile



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
	deviceprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDeviceProfilesAPI.AssignOrgDeviceProfile(context.Background(), orgId, deviceprofileId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDeviceProfilesAPI.AssignOrgDeviceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignOrgDeviceProfile`: ResponseAssignSuccess
	fmt.Fprintf(os.Stdout, "Response from `OrgsDeviceProfilesAPI.AssignOrgDeviceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignOrgDeviceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

### Return type

[**ResponseAssignSuccess**](ResponseAssignSuccess.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgDeviceProfiles

> Deviceprofile CreateOrgDeviceProfiles(ctx, orgId).Deviceprofile(deviceprofile).Execute()

createOrgDeviceProfiles



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
	deviceprofile := *openapiclient.NewDeviceprofile("gw_template") // Deviceprofile | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDeviceProfilesAPI.CreateOrgDeviceProfiles(context.Background(), orgId).Deviceprofile(deviceprofile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDeviceProfilesAPI.CreateOrgDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgDeviceProfiles`: Deviceprofile
	fmt.Fprintf(os.Stdout, "Response from `OrgsDeviceProfilesAPI.CreateOrgDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceprofile** | [**Deviceprofile**](Deviceprofile.md) | Request Body | 

### Return type

[**Deviceprofile**](Deviceprofile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgDeviceProfile

> DeleteOrgDeviceProfile(ctx, orgId, deviceprofileId).Execute()

deleteOrgDeviceProfile



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
	deviceprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsDeviceProfilesAPI.DeleteOrgDeviceProfile(context.Background(), orgId, deviceprofileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDeviceProfilesAPI.DeleteOrgDeviceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgDeviceProfileRequest struct via the builder pattern


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


## GetOrgDeviceProfile

> Deviceprofile GetOrgDeviceProfile(ctx, orgId, deviceprofileId).Execute()

getOrgDeviceProfile



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
	deviceprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDeviceProfilesAPI.GetOrgDeviceProfile(context.Background(), orgId, deviceprofileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDeviceProfilesAPI.GetOrgDeviceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgDeviceProfile`: Deviceprofile
	fmt.Fprintf(os.Stdout, "Response from `OrgsDeviceProfilesAPI.GetOrgDeviceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgDeviceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Deviceprofile**](Deviceprofile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgDeviceProfiles

> []Deviceprofile ListOrgDeviceProfiles(ctx, orgId).Type_(type_).Limit(limit).Page(page).Execute()

listOrgDeviceProfiles



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
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDeviceProfilesAPI.ListOrgDeviceProfiles(context.Background(), orgId).Type_(type_).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDeviceProfilesAPI.ListOrgDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgDeviceProfiles`: []Deviceprofile
	fmt.Fprintf(os.Stdout, "Response from `OrgsDeviceProfilesAPI.ListOrgDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**[]Deviceprofile**](Deviceprofile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignOrgDeviceProfiles

> ResponseAssignSuccess UnassignOrgDeviceProfiles(ctx, orgId, deviceprofileId).MacAddresses(macAddresses).Execute()

unassignOrgDeviceProfiles



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
	deviceprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDeviceProfilesAPI.UnassignOrgDeviceProfiles(context.Background(), orgId, deviceprofileId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDeviceProfilesAPI.UnassignOrgDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnassignOrgDeviceProfiles`: ResponseAssignSuccess
	fmt.Fprintf(os.Stdout, "Response from `OrgsDeviceProfilesAPI.UnassignOrgDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignOrgDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

### Return type

[**ResponseAssignSuccess**](ResponseAssignSuccess.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgDeviceProfile

> Deviceprofile UpdateOrgDeviceProfile(ctx, orgId, deviceprofileId).Deviceprofile(deviceprofile).Execute()

updateOrgDeviceProfiles



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
	deviceprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceprofile := *openapiclient.NewDeviceprofile("gw_template") // Deviceprofile | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsDeviceProfilesAPI.UpdateOrgDeviceProfile(context.Background(), orgId, deviceprofileId).Deviceprofile(deviceprofile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsDeviceProfilesAPI.UpdateOrgDeviceProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgDeviceProfile`: Deviceprofile
	fmt.Fprintf(os.Stdout, "Response from `OrgsDeviceProfilesAPI.UpdateOrgDeviceProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgDeviceProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceprofile** | [**Deviceprofile**](Deviceprofile.md) | Request Body | 

### Return type

[**Deviceprofile**](Deviceprofile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

