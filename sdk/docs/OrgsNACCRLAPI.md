# \OrgsNACCRLAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgNacCrl**](OrgsNACCRLAPI.md#DeleteOrgNacCrl) | **Delete** /api/v1/orgs/{org_id}/setting/mist_nac_crls/{naccrl_id} | deleteOrgNacCrl
[**GetOrgNacCrl**](OrgsNACCRLAPI.md#GetOrgNacCrl) | **Get** /api/v1/orgs/{org_id}/setting/mist_nac_crls | getOrgNacCrl
[**ImportOrgNacCrl**](OrgsNACCRLAPI.md#ImportOrgNacCrl) | **Post** /api/v1/orgs/{org_id}/setting/mist_nac_crls | importOrgNacCrl



## DeleteOrgNacCrl

> DeleteOrgNacCrl(ctx, orgId, naccrlId).Execute()

deleteOrgNacCrl



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
	naccrlId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsNACCRLAPI.DeleteOrgNacCrl(context.Background(), orgId, naccrlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACCRLAPI.DeleteOrgNacCrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**naccrlId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgNacCrlRequest struct via the builder pattern


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


## GetOrgNacCrl

> ResponseNacCrlFiles GetOrgNacCrl(ctx, orgId).Execute()

getOrgNacCrl



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
	resp, r, err := apiClient.OrgsNACCRLAPI.GetOrgNacCrl(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACCRLAPI.GetOrgNacCrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgNacCrl`: ResponseNacCrlFiles
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACCRLAPI.GetOrgNacCrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgNacCrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseNacCrlFiles**](ResponseNacCrlFiles.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOrgNacCrl

> NacCrlFile ImportOrgNacCrl(ctx, orgId).File(file).Json(json).Execute()

importOrgNacCrl



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
	file := os.NewFile(1234, "some_file") // *os.File | a binary .crl file (optional)
	json := "json_example" // string | json string with name for .crl file (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsNACCRLAPI.ImportOrgNacCrl(context.Background(), orgId).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsNACCRLAPI.ImportOrgNacCrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOrgNacCrl`: NacCrlFile
	fmt.Fprintf(os.Stdout, "Response from `OrgsNACCRLAPI.ImportOrgNacCrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportOrgNacCrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | a binary .crl file | 
 **json** | **string** | json string with name for .crl file (optional) | 

### Return type

[**NacCrlFile**](NacCrlFile.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

