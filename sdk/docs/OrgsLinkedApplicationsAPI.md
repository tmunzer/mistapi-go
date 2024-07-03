# \OrgsLinkedApplicationsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgOauthAppAccounts**](OrgsLinkedApplicationsAPI.md#AddOrgOauthAppAccounts) | **Post** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts | addOrgOauthAppAccounts
[**DeleteOrgOauthAppAuthorization**](OrgsLinkedApplicationsAPI.md#DeleteOrgOauthAppAuthorization) | **Delete** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts/{account_id} | deleteOrgOauthAppAuthorization
[**GetOrgOauthAppLinkedStatus**](OrgsLinkedApplicationsAPI.md#GetOrgOauthAppLinkedStatus) | **Get** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts | getOrgOauthAppAuthorizationUrl
[**LinkOrgToJuniperJuniperAccount**](OrgsLinkedApplicationsAPI.md#LinkOrgToJuniperJuniperAccount) | **Post** /api/v1/orgs/{org_id}/setting/juniper/link_accounts | linkOrgToJuniperJuniperAccount
[**UnlinkOrgFromJuniperCustomerId**](OrgsLinkedApplicationsAPI.md#UnlinkOrgFromJuniperCustomerId) | **Delete** /api/v1/orgs/{org_id}/setting/juniper/unlink_account | unlinkOrgFromJuniperCustomerId
[**UpdateOrgOauthAppAccounts**](OrgsLinkedApplicationsAPI.md#UpdateOrgOauthAppAccounts) | **Put** /api/v1/orgs/{org_id}/setting/{app_name}/link_accounts | updateOrgOauthAppAccounts



## AddOrgOauthAppAccounts

> AddOrgOauthAppAccounts(ctx, orgId, appName).AccountOauthAdd(accountOauthAdd).Execute()

addOrgOauthAppAccounts



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
	appName := "appName_example" // string | OAuth application name
	accountOauthAdd := openapiclient.account_oauth_add{AccountJamfConfig: openapiclient.NewAccountJamfConfig("ClientId_example", "ClientSecret_example", "junipertest.jamfcloud.com", "CompliantGroup1")} // AccountOauthAdd |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsLinkedApplicationsAPI.AddOrgOauthAppAccounts(context.Background(), orgId, appName).AccountOauthAdd(accountOauthAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLinkedApplicationsAPI.AddOrgOauthAppAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**appName** | **string** | OAuth application name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrgOauthAppAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountOauthAdd** | [**AccountOauthAdd**](AccountOauthAdd.md) |  | 

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


## DeleteOrgOauthAppAuthorization

> DeleteOrgOauthAppAuthorization(ctx, orgId, appName, accountId).Execute()

deleteOrgOauthAppAuthorization



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
	appName := "appName_example" // string | 
	accountId := "iojzXIJWEuiD73ZvydOfg" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsLinkedApplicationsAPI.DeleteOrgOauthAppAuthorization(context.Background(), orgId, appName, accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLinkedApplicationsAPI.DeleteOrgOauthAppAuthorization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**appName** | **string** |  | 
**accountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgOauthAppAuthorizationRequest struct via the builder pattern


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


## GetOrgOauthAppLinkedStatus

> ResponseOauthAppLink GetOrgOauthAppLinkedStatus(ctx, orgId, appName).Forward(forward).Execute()

getOrgOauthAppAuthorizationUrl



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
	appName := "appName_example" // string | OAuth application name
	forward := "forward_example" // string | Mist portal url to which backend needs to redirect after succesful OAuth authorization. **Required** to get the `authorization_url`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsLinkedApplicationsAPI.GetOrgOauthAppLinkedStatus(context.Background(), orgId, appName).Forward(forward).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLinkedApplicationsAPI.GetOrgOauthAppLinkedStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgOauthAppLinkedStatus`: ResponseOauthAppLink
	fmt.Fprintf(os.Stdout, "Response from `OrgsLinkedApplicationsAPI.GetOrgOauthAppLinkedStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**appName** | **string** | OAuth application name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgOauthAppLinkedStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forward** | **string** | Mist portal url to which backend needs to redirect after succesful OAuth authorization. **Required** to get the &#x60;authorization_url&#x60; | 

### Return type

[**ResponseOauthAppLink**](ResponseOauthAppLink.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkOrgToJuniperJuniperAccount

> AccountJuniperInfo LinkOrgToJuniperJuniperAccount(ctx, orgId).AccountJuniperConfig(accountJuniperConfig).Execute()

linkOrgToJuniperJuniperAccount



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
	accountJuniperConfig := *openapiclient.NewAccountJuniperConfig("password", "john@nmo.com") // AccountJuniperConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsLinkedApplicationsAPI.LinkOrgToJuniperJuniperAccount(context.Background(), orgId).AccountJuniperConfig(accountJuniperConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLinkedApplicationsAPI.LinkOrgToJuniperJuniperAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkOrgToJuniperJuniperAccount`: AccountJuniperInfo
	fmt.Fprintf(os.Stdout, "Response from `OrgsLinkedApplicationsAPI.LinkOrgToJuniperJuniperAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkOrgToJuniperJuniperAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountJuniperConfig** | [**AccountJuniperConfig**](AccountJuniperConfig.md) |  | 

### Return type

[**AccountJuniperInfo**](AccountJuniperInfo.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkOrgFromJuniperCustomerId

> UnlinkOrgFromJuniperCustomerId(ctx, orgId).AccountJuniperInfo(accountJuniperInfo).Execute()

unlinkOrgFromJuniperCustomerId



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
	accountJuniperInfo := *openapiclient.NewAccountJuniperInfo() // AccountJuniperInfo |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsLinkedApplicationsAPI.UnlinkOrgFromJuniperCustomerId(context.Background(), orgId).AccountJuniperInfo(accountJuniperInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLinkedApplicationsAPI.UnlinkOrgFromJuniperCustomerId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnlinkOrgFromJuniperCustomerIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountJuniperInfo** | [**AccountJuniperInfo**](AccountJuniperInfo.md) |  | 

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


## UpdateOrgOauthAppAccounts

> UpdateOrgOauthAppAccounts(ctx, orgId, appName).AccountOauthConfig(accountOauthConfig).Execute()

updateOrgOauthAppAccounts



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
	appName := "appName_example" // string | OAuth application name
	accountOauthConfig := *openapiclient.NewAccountOauthConfig("iojzXIJWEuiD73ZvydOfg") // AccountOauthConfig |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsLinkedApplicationsAPI.UpdateOrgOauthAppAccounts(context.Background(), orgId, appName).AccountOauthConfig(accountOauthConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsLinkedApplicationsAPI.UpdateOrgOauthAppAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**appName** | **string** | OAuth application name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgOauthAppAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountOauthConfig** | [**AccountOauthConfig**](AccountOauthConfig.md) |  | 

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

