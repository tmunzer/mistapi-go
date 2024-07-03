# \OrgsIDPProfilesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgIdpProfile**](OrgsIDPProfilesAPI.md#CreateOrgIdpProfile) | **Post** /api/v1/orgs/{org_id}/idpprofiles | createOrgIdpProfile
[**DeleteOrgIdpProfile**](OrgsIDPProfilesAPI.md#DeleteOrgIdpProfile) | **Delete** /api/v1/orgs/{org_id}/idpprofiles/{idpprofile_id} | deleteOrgIdpProfile
[**GetOrgIdpProfile**](OrgsIDPProfilesAPI.md#GetOrgIdpProfile) | **Get** /api/v1/orgs/{org_id}/idpprofiles/{idpprofile_id} | getOrgIdpProfile
[**ListOrgIdpProfiles**](OrgsIDPProfilesAPI.md#ListOrgIdpProfiles) | **Get** /api/v1/orgs/{org_id}/idpprofiles | listOrgIdpProfiles
[**UpdateOrgIdpProfile**](OrgsIDPProfilesAPI.md#UpdateOrgIdpProfile) | **Put** /api/v1/orgs/{org_id}/idpprofiles/{idpprofile_id} | updateOrgIdpProfile



## CreateOrgIdpProfile

> IdpProfile CreateOrgIdpProfile(ctx, orgId).IdpProfile(idpProfile).Execute()

createOrgIdpProfile



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
	idpProfile := *openapiclient.NewIdpProfile() // IdpProfile |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsIDPProfilesAPI.CreateOrgIdpProfile(context.Background(), orgId).IdpProfile(idpProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsIDPProfilesAPI.CreateOrgIdpProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgIdpProfile`: IdpProfile
	fmt.Fprintf(os.Stdout, "Response from `OrgsIDPProfilesAPI.CreateOrgIdpProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgIdpProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idpProfile** | [**IdpProfile**](IdpProfile.md) |  | 

### Return type

[**IdpProfile**](IdpProfile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgIdpProfile

> DeleteOrgIdpProfile(ctx, orgId, idpprofileId).Execute()

deleteOrgIdpProfile



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
	idpprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsIDPProfilesAPI.DeleteOrgIdpProfile(context.Background(), orgId, idpprofileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsIDPProfilesAPI.DeleteOrgIdpProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**idpprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgIdpProfileRequest struct via the builder pattern


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


## GetOrgIdpProfile

> IdpProfile GetOrgIdpProfile(ctx, orgId, idpprofileId).Execute()

getOrgIdpProfile



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
	idpprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsIDPProfilesAPI.GetOrgIdpProfile(context.Background(), orgId, idpprofileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsIDPProfilesAPI.GetOrgIdpProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgIdpProfile`: IdpProfile
	fmt.Fprintf(os.Stdout, "Response from `OrgsIDPProfilesAPI.GetOrgIdpProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**idpprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgIdpProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdpProfile**](IdpProfile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgIdpProfiles

> []IdpProfile ListOrgIdpProfiles(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgIdpProfiles



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
	resp, r, err := apiClient.OrgsIDPProfilesAPI.ListOrgIdpProfiles(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsIDPProfilesAPI.ListOrgIdpProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgIdpProfiles`: []IdpProfile
	fmt.Fprintf(os.Stdout, "Response from `OrgsIDPProfilesAPI.ListOrgIdpProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgIdpProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]IdpProfile**](IdpProfile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgIdpProfile

> UpdateOrgIdpProfile(ctx, orgId, idpprofileId).IdpProfile(idpProfile).Execute()

updateOrgIdpProfile



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
	idpprofileId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	idpProfile := *openapiclient.NewIdpProfile() // IdpProfile |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsIDPProfilesAPI.UpdateOrgIdpProfile(context.Background(), orgId, idpprofileId).IdpProfile(idpProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsIDPProfilesAPI.UpdateOrgIdpProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**idpprofileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgIdpProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idpProfile** | [**IdpProfile**](IdpProfile.md) |  | 

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

