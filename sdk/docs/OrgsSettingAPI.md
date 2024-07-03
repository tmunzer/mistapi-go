# \OrgsSettingAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgWirelessClientsBlocklist**](OrgsSettingAPI.md#CreateOrgWirelessClientsBlocklist) | **Post** /api/v1/orgs/{org_id}/setting/blacklist | createOrgWirelessClientsBlocklist
[**DeleteOrgWirelessClientsBlocklist**](OrgsSettingAPI.md#DeleteOrgWirelessClientsBlocklist) | **Delete** /api/v1/orgs/{org_id}/setting/blacklist | deleteOrgWirelessClientsBlocklist
[**GetOrgSettings**](OrgsSettingAPI.md#GetOrgSettings) | **Get** /api/v1/orgs/{org_id}/setting | getOrgSettings
[**SetOrgCustomBucket**](OrgsSettingAPI.md#SetOrgCustomBucket) | **Post** /api/v1/orgs/{org_id}/setting/pcap_bucket/setup | setOrgCustomBucket
[**UpdateOrgSettings**](OrgsSettingAPI.md#UpdateOrgSettings) | **Put** /api/v1/orgs/{org_id}/setting | updateOrgSettings
[**VerifyOrgCustomBucket**](OrgsSettingAPI.md#VerifyOrgCustomBucket) | **Post** /api/v1/orgs/{org_id}/setting/pcap_bucket/verify | verifyOrgCustomBucket



## CreateOrgWirelessClientsBlocklist

> MacAddresses CreateOrgWirelessClientsBlocklist(ctx, orgId).MacAddresses(macAddresses).Execute()

createOrgWirelessClientsBlocklist



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
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSettingAPI.CreateOrgWirelessClientsBlocklist(context.Background(), orgId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSettingAPI.CreateOrgWirelessClientsBlocklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgWirelessClientsBlocklist`: MacAddresses
	fmt.Fprintf(os.Stdout, "Response from `OrgsSettingAPI.CreateOrgWirelessClientsBlocklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWirelessClientsBlocklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

### Return type

[**MacAddresses**](MacAddresses.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgWirelessClientsBlocklist

> DeleteOrgWirelessClientsBlocklist(ctx, orgId).Execute()

deleteOrgWirelessClientsBlocklist



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSettingAPI.DeleteOrgWirelessClientsBlocklist(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSettingAPI.DeleteOrgWirelessClientsBlocklist``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrgWirelessClientsBlocklistRequest struct via the builder pattern


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


## GetOrgSettings

> OrgSetting GetOrgSettings(ctx, orgId).Execute()

getOrgSettings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSettingAPI.GetOrgSettings(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSettingAPI.GetOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSettings`: OrgSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgsSettingAPI.GetOrgSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetOrgCustomBucket

> ResponsePcapBucketConfig SetOrgCustomBucket(ctx, orgId).PcapBucket(pcapBucket).Execute()

setOrgCustomBucket



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
	pcapBucket := *openapiclient.NewPcapBucket("company-private-pcap") // PcapBucket | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSettingAPI.SetOrgCustomBucket(context.Background(), orgId).PcapBucket(pcapBucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSettingAPI.SetOrgCustomBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetOrgCustomBucket`: ResponsePcapBucketConfig
	fmt.Fprintf(os.Stdout, "Response from `OrgsSettingAPI.SetOrgCustomBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetOrgCustomBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pcapBucket** | [**PcapBucket**](PcapBucket.md) | Request Body | 

### Return type

[**ResponsePcapBucketConfig**](ResponsePcapBucketConfig.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSettings

> OrgSetting UpdateOrgSettings(ctx, orgId).OrgSetting(orgSetting).Execute()

updateOrgSettings



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
	orgSetting := *openapiclient.NewOrgSetting() // OrgSetting | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSettingAPI.UpdateOrgSettings(context.Background(), orgId).OrgSetting(orgSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSettingAPI.UpdateOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSettings`: OrgSetting
	fmt.Fprintf(os.Stdout, "Response from `OrgsSettingAPI.UpdateOrgSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orgSetting** | [**OrgSetting**](OrgSetting.md) | Request Body | 

### Return type

[**OrgSetting**](OrgSetting.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyOrgCustomBucket

> VerifyOrgCustomBucket(ctx, orgId).PcapBucketVerify(pcapBucketVerify).Execute()

verifyOrgCustomBucket



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
	pcapBucketVerify := *openapiclient.NewPcapBucketVerify("company-private-pcap", "eyJhbGciOiJIUzI1J9.eyJzdWIiOiIxMjM0joiMjgxOG5MDIyfQ.2rzcRvMA3Eg09NnjCAC-1EWMRtxAnFDM") // PcapBucketVerify | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSettingAPI.VerifyOrgCustomBucket(context.Background(), orgId).PcapBucketVerify(pcapBucketVerify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSettingAPI.VerifyOrgCustomBucket``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVerifyOrgCustomBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pcapBucketVerify** | [**PcapBucketVerify**](PcapBucketVerify.md) | Request Body | 

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

