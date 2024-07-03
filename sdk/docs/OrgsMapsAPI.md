# \OrgsMapsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportOrgMapToSite**](OrgsMapsAPI.md#ImportOrgMapToSite) | **Post** /api/v1/orgs/{org_id}/sites/{site_name}/maps/import | importOrgMapToSite
[**ImportOrgMaps**](OrgsMapsAPI.md#ImportOrgMaps) | **Post** /api/v1/orgs/{org_id}/maps/import | importOrgMaps



## ImportOrgMapToSite

> ResponseMapImport ImportOrgMapToSite(ctx, orgId, siteName).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()

importOrgMapToSite



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
	siteName := "siteName_example" // string | 
	autoDeviceprofileAssignment := true // bool | whether to auto assign device to deviceprofile by name (optional)
	csv := os.NewFile(1234, "some_file") // *os.File | csv file for ap name mapping, optional (optional)
	file := os.NewFile(1234, "some_file") // *os.File | ekahau or ibwave file (optional)
	json := *openapiclient.NewMapImportJson(openapiclient.map_import_json_vendor_name("")) // MapImportJson |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMapsAPI.ImportOrgMapToSite(context.Background(), orgId, siteName).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMapsAPI.ImportOrgMapToSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOrgMapToSite`: ResponseMapImport
	fmt.Fprintf(os.Stdout, "Response from `OrgsMapsAPI.ImportOrgMapToSite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**siteName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportOrgMapToSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoDeviceprofileAssignment** | **bool** | whether to auto assign device to deviceprofile by name | 
 **csv** | ***os.File** | csv file for ap name mapping, optional | 
 **file** | ***os.File** | ekahau or ibwave file | 
 **json** | [**MapImportJson**](MapImportJson.md) |  | 

### Return type

[**ResponseMapImport**](ResponseMapImport.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportOrgMaps

> ResponseMapImport ImportOrgMaps(ctx, orgId).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()

importOrgMaps



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
	autoDeviceprofileAssignment := true // bool | whether to auto assign device to deviceprofile by name (optional)
	csv := os.NewFile(1234, "some_file") // *os.File | csv file for ap name mapping, optional (optional)
	file := os.NewFile(1234, "some_file") // *os.File | ekahau or ibwave file (optional)
	json := *openapiclient.NewMapOrgImportFileJson(openapiclient.map_org_import_file_json_vendor_name("")) // MapOrgImportFileJson |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMapsAPI.ImportOrgMaps(context.Background(), orgId).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMapsAPI.ImportOrgMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportOrgMaps`: ResponseMapImport
	fmt.Fprintf(os.Stdout, "Response from `OrgsMapsAPI.ImportOrgMaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportOrgMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoDeviceprofileAssignment** | **bool** | whether to auto assign device to deviceprofile by name | 
 **csv** | ***os.File** | csv file for ap name mapping, optional | 
 **file** | ***os.File** | ekahau or ibwave file | 
 **json** | [**MapOrgImportFileJson**](MapOrgImportFileJson.md) |  | 

### Return type

[**ResponseMapImport**](ResponseMapImport.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

