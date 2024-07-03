# \OrgsPsksAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgPsk**](OrgsPsksAPI.md#CreateOrgPsk) | **Post** /api/v1/orgs/{org_id}/psks | createOrgPsk
[**DeleteOrgPsk**](OrgsPsksAPI.md#DeleteOrgPsk) | **Delete** /api/v1/orgs/{org_id}/psks/{psk_id} | deleteOrgPsk
[**DeleteOrgPskList**](OrgsPsksAPI.md#DeleteOrgPskList) | **Post** /api/v1/orgs/{org_id}/psks/delete | deleteOrgPskList
[**DeleteOrgPskOldPassphrase**](OrgsPsksAPI.md#DeleteOrgPskOldPassphrase) | **Post** /api/v1/orgs/{org_id}/psks/{psk_id}/delete_old_passphrase | deleteOrgPskOldPassphrase
[**GetOrgPsk**](OrgsPsksAPI.md#GetOrgPsk) | **Get** /api/v1/orgs/{org_id}/psks/{psk_id} | getOrgPsk
[**ImportOrgPsks**](OrgsPsksAPI.md#ImportOrgPsks) | **Post** /api/v1/orgs/{org_id}/psks/import | importOrgPsks
[**ListOrgPsks**](OrgsPsksAPI.md#ListOrgPsks) | **Get** /api/v1/orgs/{org_id}/psks | listOrgPsks
[**UpdateOrgMultiplePsks**](OrgsPsksAPI.md#UpdateOrgMultiplePsks) | **Put** /api/v1/orgs/{org_id}/psks | updateOrgMultiplePsks
[**UpdateOrgPsk**](OrgsPsksAPI.md#UpdateOrgPsk) | **Put** /api/v1/orgs/{org_id}/psks/{psk_id} | updateOrgPsk



## CreateOrgPsk

> Psk CreateOrgPsk(ctx, orgId).Upsert(upsert).Psk(psk).Execute()

createOrgPsk



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
	upsert := true // bool | if a key exists with the same `name`, replace it with the new one (optional)
	psk := *openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example") // Psk | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPsksAPI.CreateOrgPsk(context.Background(), orgId).Upsert(upsert).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.CreateOrgPsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgPsk`: Psk
	fmt.Fprintf(os.Stdout, "Response from `OrgsPsksAPI.CreateOrgPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **bool** | if a key exists with the same &#x60;name&#x60;, replace it with the new one | 
 **psk** | [**Psk**](Psk.md) | Request Body | 

### Return type

[**Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgPsk

> DeleteOrgPsk(ctx, orgId, pskId).Execute()

deleteOrgPsk



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
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsPsksAPI.DeleteOrgPsk(context.Background(), orgId, pskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.DeleteOrgPsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskId** | **string** | PSK ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgPskRequest struct via the builder pattern


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


## DeleteOrgPskList

> DeleteOrgPskList(ctx, orgId).PskIdList(pskIdList).Execute()

deleteOrgPskList



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
	pskIdList := *openapiclient.NewPskIdList() // PskIdList |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsPsksAPI.DeleteOrgPskList(context.Background(), orgId).PskIdList(pskIdList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.DeleteOrgPskList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrgPskListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pskIdList** | [**PskIdList**](PskIdList.md) |  | 

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


## DeleteOrgPskOldPassphrase

> Psk DeleteOrgPskOldPassphrase(ctx, orgId, pskId).Execute()

deleteOrgPskOldPassphrase



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
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPsksAPI.DeleteOrgPskOldPassphrase(context.Background(), orgId, pskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.DeleteOrgPskOldPassphrase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrgPskOldPassphrase`: Psk
	fmt.Fprintf(os.Stdout, "Response from `OrgsPsksAPI.DeleteOrgPskOldPassphrase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskId** | **string** | PSK ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgPskOldPassphraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgPsk

> Psk GetOrgPsk(ctx, orgId, pskId).Execute()

getOrgPsk



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
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPsksAPI.GetOrgPsk(context.Background(), orgId, pskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.GetOrgPsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgPsk`: Psk
	fmt.Fprintf(os.Stdout, "Response from `OrgsPsksAPI.GetOrgPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskId** | **string** | PSK ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOrgPsks

> []Psk ImportOrgPsks(ctx, orgId).Psk(psk).Execute()

importOrgPsks



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
	psk := []openapiclient.Psk{*openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example")} // []Psk |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPsksAPI.ImportOrgPsks(context.Background(), orgId).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.ImportOrgPsks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOrgPsks`: []Psk
	fmt.Fprintf(os.Stdout, "Response from `OrgsPsksAPI.ImportOrgPsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportOrgPsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **psk** | [**[]Psk**](Psk.md) |  | 

### Return type

[**[]Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgPsks

> []Psk ListOrgPsks(ctx, orgId).Name(name).Ssid(ssid).Role(role).Page(page).Limit(limit).Execute()

listOrgPsks



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
	name := "psk_name" // string |  (optional)
	ssid := "ssid_example" // string |  (optional)
	role := "role_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPsksAPI.ListOrgPsks(context.Background(), orgId).Name(name).Ssid(ssid).Role(role).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.ListOrgPsks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgPsks`: []Psk
	fmt.Fprintf(os.Stdout, "Response from `OrgsPsksAPI.ListOrgPsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgPsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 
 **ssid** | **string** |  | 
 **role** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMultiplePsks

> []Psk UpdateOrgMultiplePsks(ctx, orgId).Psk(psk).Execute()

updateOrgMultiplePsks



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
	psk := []openapiclient.Psk{*openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example")} // []Psk |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPsksAPI.UpdateOrgMultiplePsks(context.Background(), orgId).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.UpdateOrgMultiplePsks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgMultiplePsks`: []Psk
	fmt.Fprintf(os.Stdout, "Response from `OrgsPsksAPI.UpdateOrgMultiplePsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMultiplePsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **psk** | [**[]Psk**](Psk.md) |  | 

### Return type

[**[]Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgPsk

> Psk UpdateOrgPsk(ctx, orgId, pskId).Psk(psk).Execute()

updateOrgPsk



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
	pskId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | PSK ID
	psk := *openapiclient.NewPsk("Name_example", "Passphrase_example", "Ssid_example") // Psk | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsPsksAPI.UpdateOrgPsk(context.Background(), orgId, pskId).Psk(psk).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsPsksAPI.UpdateOrgPsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgPsk`: Psk
	fmt.Fprintf(os.Stdout, "Response from `OrgsPsksAPI.UpdateOrgPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**pskId** | **string** | PSK ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **psk** | [**Psk**](Psk.md) | Request Body | 

### Return type

[**Psk**](Psk.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

