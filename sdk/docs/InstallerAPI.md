# \InstallerAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstallerDeviceImage**](InstallerAPI.md#AddInstallerDeviceImage) | **Post** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/{image_name} | addInstallerDeviceImage
[**ClaimInstallerDevices**](InstallerAPI.md#ClaimInstallerDevices) | **Post** /api/v1/installer/orgs/{org_id}/devices | claimInstallerDevices
[**CreateInstallerMap**](InstallerAPI.md#CreateInstallerMap) | **Post** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/{map_id} | createInstallerMap
[**CreateInstallerVirtualChassis**](InstallerAPI.md#CreateInstallerVirtualChassis) | **Post** /api/v1/installer/orgs/{org_id}/devices/{fpc0_mac}/vc | createInstallerVirtualChassis
[**CreateOrUpdateInstallerSites**](InstallerAPI.md#CreateOrUpdateInstallerSites) | **Put** /api/v1/installer/orgs/{org_id}/sites/{site_name} | createOrUpdateInstallerSites
[**DeleteInstallerDeviceImage**](InstallerAPI.md#DeleteInstallerDeviceImage) | **Delete** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/{image_name} | deleteInstallerDeviceImage
[**DeleteInstallerMap**](InstallerAPI.md#DeleteInstallerMap) | **Delete** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/{map_id} | deleteInstallerMap
[**GetInstallerDeviceVirtualChassis**](InstallerAPI.md#GetInstallerDeviceVirtualChassis) | **Get** /api/v1/installer/orgs/{org_id}/devices/{fpc0_mac}/vc | getInstallerDeviceVirtualChassis
[**ImportInstallerMap**](InstallerAPI.md#ImportInstallerMap) | **Post** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/import | importInstallerMap
[**ListInstallerAlarmTemplates**](InstallerAPI.md#ListInstallerAlarmTemplates) | **Get** /api/v1/installer/orgs/{org_id}/alarmtemplates | listInstallerAlarmTemplates
[**ListInstallerDeviceProfiles**](InstallerAPI.md#ListInstallerDeviceProfiles) | **Get** /api/v1/installer/orgs/{org_id}/deviceprofiles | listInstallerDeviceProfiles
[**ListInstallerListOfRenctlyClaimedDevices**](InstallerAPI.md#ListInstallerListOfRenctlyClaimedDevices) | **Get** /api/v1/installer/orgs/{org_id}/devices | listInstallerListOfRenctlyClaimedDevices
[**ListInstallerMaps**](InstallerAPI.md#ListInstallerMaps) | **Get** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps | listInstallerMaps
[**ListInstallerRfTemplatesNames**](InstallerAPI.md#ListInstallerRfTemplatesNames) | **Get** /api/v1/installer/orgs/{org_id}/rftemplates | listInstallerRfTemplatesNames
[**ListInstallerSiteGroups**](InstallerAPI.md#ListInstallerSiteGroups) | **Get** /api/v1/installer/orgs/{org_id}/sitegroups | listInstallerSiteGroups
[**ListInstallerSites**](InstallerAPI.md#ListInstallerSites) | **Get** /api/v1/installer/orgs/{org_id}/sites | listInstallerSites
[**OptimizeInstallerRrm**](InstallerAPI.md#OptimizeInstallerRrm) | **Get** /api/v1/installer/sites/{site_name}/optimize | optimizeInstallerRrm
[**ProvisionInstallerDevices**](InstallerAPI.md#ProvisionInstallerDevices) | **Put** /api/v1/installer/orgs/{org_id}/devices/{device_mac} | provisionInstallerDevices
[**StartInstallerLocateDevice**](InstallerAPI.md#StartInstallerLocateDevice) | **Post** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/locate | startInstallerLocateDevice
[**StopInstallerLocateDevice**](InstallerAPI.md#StopInstallerLocateDevice) | **Post** /api/v1/installer/orgs/{org_id}/devices/{device_mac}/unlocate | stopInstallerLocateDevice
[**UnassignInstallerRecentlyClaimedDevice**](InstallerAPI.md#UnassignInstallerRecentlyClaimedDevice) | **Delete** /api/v1/installer/orgs/{org_id}/devices/{device_mac} | unassignInstallerRecentlyClaimedDevice
[**UpdateInstallerMap**](InstallerAPI.md#UpdateInstallerMap) | **Put** /api/v1/installer/orgs/{org_id}/sites/{site_name}/maps/{map_id} | updateInstallerMap
[**UpdateInstallerVirtualChassisMember**](InstallerAPI.md#UpdateInstallerVirtualChassisMember) | **Put** /api/v1/installer/orgs/{org_id}/devices/{fpc0_mac}/vc | updateInstallerVirtualChassisMember



## AddInstallerDeviceImage

> AddInstallerDeviceImage(ctx, orgId, imageName, deviceMac).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()

addInstallerDeviceImage



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
	imageName := "imageName_example" // string | 
	deviceMac := "0000000000ab" // string | 
	autoDeviceprofileAssignment := true // bool | whether to auto assign device to deviceprofile by name (optional)
	csv := os.NewFile(1234, "some_file") // *os.File | csv file for ap name mapping, optional (optional)
	file := os.NewFile(1234, "some_file") // *os.File | ekahau or ibwave file (optional)
	json := *openapiclient.NewMapImportJson(openapiclient.map_import_json_vendor_name("")) // MapImportJson |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.AddInstallerDeviceImage(context.Background(), orgId, imageName, deviceMac).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.AddInstallerDeviceImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**imageName** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInstallerDeviceImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **autoDeviceprofileAssignment** | **bool** | whether to auto assign device to deviceprofile by name | 
 **csv** | ***os.File** | csv file for ap name mapping, optional | 
 **file** | ***os.File** | ekahau or ibwave file | 
 **json** | [**MapImportJson**](MapImportJson.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimInstallerDevices

> ResponseInventory ClaimInstallerDevices(ctx, orgId).RequestBody(requestBody).Execute()

claimInstallerDevices



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
	requestBody := []string{"Property_example"} // []string | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.ClaimInstallerDevices(context.Background(), orgId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ClaimInstallerDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimInstallerDevices`: ResponseInventory
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ClaimInstallerDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimInstallerDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | Request Body | 

### Return type

[**ResponseInventory**](ResponseInventory.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstallerMap

> ModelMap CreateInstallerMap(ctx, orgId, siteName, mapId).ModelMap(modelMap).Execute()

createInstallerMap



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	modelMap := *openapiclient.NewMap() // ModelMap | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.CreateInstallerMap(context.Background(), orgId, siteName, mapId).ModelMap(modelMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.CreateInstallerMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstallerMap`: ModelMap
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.CreateInstallerMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**siteName** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstallerMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modelMap** | [**ModelMap**](ModelMap.md) | Request Body | 

### Return type

[**ModelMap**](ModelMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstallerVirtualChassis

> ResponseVirtualChassisConfig CreateInstallerVirtualChassis(ctx, orgId, fpc0Mac).VirtualChassisConfig(virtualChassisConfig).Execute()

createInstallerVirtualChassis



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
	fpc0Mac := "fpc0Mac_example" // string | FPC0 MAC Address
	virtualChassisConfig := *openapiclient.NewVirtualChassisConfig() // VirtualChassisConfig | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.CreateInstallerVirtualChassis(context.Background(), orgId, fpc0Mac).VirtualChassisConfig(virtualChassisConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.CreateInstallerVirtualChassis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstallerVirtualChassis`: ResponseVirtualChassisConfig
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.CreateInstallerVirtualChassis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**fpc0Mac** | **string** | FPC0 MAC Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstallerVirtualChassisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualChassisConfig** | [**VirtualChassisConfig**](VirtualChassisConfig.md) | Request Body | 

### Return type

[**ResponseVirtualChassisConfig**](ResponseVirtualChassisConfig.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateInstallerSites

> CreateOrUpdateInstallerSites(ctx, orgId, siteName).InstallerSite(installerSite).Execute()

createOrUpdateInstallerSites



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
	installerSite := *openapiclient.NewInstallerSite("1601 S. Deanza Blvd., Cupertino, CA, 95014", "US", *openapiclient.NewLatLng(float64(37.295833), float64(-122.032946)), "Mist Office") // InstallerSite | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.CreateOrUpdateInstallerSites(context.Background(), orgId, siteName).InstallerSite(installerSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.CreateOrUpdateInstallerSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**siteName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateInstallerSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installerSite** | [**InstallerSite**](InstallerSite.md) | Request Body | 

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


## DeleteInstallerDeviceImage

> DeleteInstallerDeviceImage(ctx, orgId, imageName, deviceMac).Execute()

deleteInstallerDeviceImage



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
	imageName := "imageName_example" // string | 
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.DeleteInstallerDeviceImage(context.Background(), orgId, imageName, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.DeleteInstallerDeviceImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**imageName** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstallerDeviceImageRequest struct via the builder pattern


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


## DeleteInstallerMap

> DeleteInstallerMap(ctx, orgId, siteName, mapId).Execute()

deleteInstallerMap



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.DeleteInstallerMap(context.Background(), orgId, siteName, mapId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.DeleteInstallerMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**siteName** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstallerMapRequest struct via the builder pattern


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


## GetInstallerDeviceVirtualChassis

> ResponseVirtualChassisConfig GetInstallerDeviceVirtualChassis(ctx, orgId, fpc0Mac).Execute()

getInstallerDeviceVirtualChassis



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
	fpc0Mac := "fpc0Mac_example" // string | FPC0 MAC Address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.GetInstallerDeviceVirtualChassis(context.Background(), orgId, fpc0Mac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.GetInstallerDeviceVirtualChassis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstallerDeviceVirtualChassis`: ResponseVirtualChassisConfig
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.GetInstallerDeviceVirtualChassis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**fpc0Mac** | **string** | FPC0 MAC Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstallerDeviceVirtualChassisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseVirtualChassisConfig**](ResponseVirtualChassisConfig.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportInstallerMap

> ResponseMapImport ImportInstallerMap(ctx, orgId, siteName).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()

importInstallerMap



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
	resp, r, err := apiClient.InstallerAPI.ImportInstallerMap(context.Background(), orgId, siteName).AutoDeviceprofileAssignment(autoDeviceprofileAssignment).Csv(csv).File(file).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ImportInstallerMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportInstallerMap`: ResponseMapImport
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ImportInstallerMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**siteName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportInstallerMapRequest struct via the builder pattern


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


## ListInstallerAlarmTemplates

> []InstallersItem ListInstallerAlarmTemplates(ctx, orgId).Execute()

listInstallerAlarmTemplates



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
	resp, r, err := apiClient.InstallerAPI.ListInstallerAlarmTemplates(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ListInstallerAlarmTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallerAlarmTemplates`: []InstallersItem
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ListInstallerAlarmTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallerAlarmTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstallersItem**](InstallersItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallerDeviceProfiles

> []InstallersItem ListInstallerDeviceProfiles(ctx, orgId).Type_(type_).Execute()

listInstallerDeviceProfiles



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
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.ListInstallerDeviceProfiles(context.Background(), orgId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ListInstallerDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallerDeviceProfiles`: []InstallersItem
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ListInstallerDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallerDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]

### Return type

[**[]InstallersItem**](InstallersItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallerListOfRenctlyClaimedDevices

> []InstallerDevice ListInstallerListOfRenctlyClaimedDevices(ctx, orgId).Model(model).SiteName(siteName).SiteId(siteId).Limit(limit).Page(page).Execute()

listInstallerListOfRenctlyClaimedDevices



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
	model := "model_example" // string | Device Model (optional)
	siteName := "siteName_example" // string | Site Name (optional)
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Site ID (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.ListInstallerListOfRenctlyClaimedDevices(context.Background(), orgId).Model(model).SiteName(siteName).SiteId(siteId).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ListInstallerListOfRenctlyClaimedDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallerListOfRenctlyClaimedDevices`: []InstallerDevice
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ListInstallerListOfRenctlyClaimedDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallerListOfRenctlyClaimedDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **string** | Device Model | 
 **siteName** | **string** | Site Name | 
 **siteId** | **string** | Site ID | 
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**[]InstallerDevice**](InstallerDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallerMaps

> []ModelMap ListInstallerMaps(ctx, orgId, siteName).Execute()

listInstallerMaps



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.ListInstallerMaps(context.Background(), orgId, siteName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ListInstallerMaps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallerMaps`: []ModelMap
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ListInstallerMaps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**siteName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallerMapsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ModelMap**](ModelMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallerRfTemplatesNames

> []InstallersItem ListInstallerRfTemplatesNames(ctx, orgId).Execute()

listInstallerRfTemplatesNames



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
	resp, r, err := apiClient.InstallerAPI.ListInstallerRfTemplatesNames(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ListInstallerRfTemplatesNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallerRfTemplatesNames`: []InstallersItem
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ListInstallerRfTemplatesNames`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallerRfTemplatesNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstallersItem**](InstallersItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallerSiteGroups

> []InstallersItem ListInstallerSiteGroups(ctx, orgId).Execute()

listInstallerSiteGroups



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
	resp, r, err := apiClient.InstallerAPI.ListInstallerSiteGroups(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ListInstallerSiteGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallerSiteGroups`: []InstallersItem
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ListInstallerSiteGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallerSiteGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstallersItem**](InstallersItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstallerSites

> []InstallerSite ListInstallerSites(ctx, orgId).Execute()

listInstallerSites



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
	resp, r, err := apiClient.InstallerAPI.ListInstallerSites(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ListInstallerSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstallerSites`: []InstallerSite
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.ListInstallerSites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstallerSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstallerSite**](InstallerSite.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OptimizeInstallerRrm

> OptimizeInstallerRrm(ctx, siteName).Execute()

optimizeInstallerRrm



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
	siteName := "siteName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.OptimizeInstallerRrm(context.Background(), siteName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.OptimizeInstallerRrm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOptimizeInstallerRrmRequest struct via the builder pattern


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


## ProvisionInstallerDevices

> ProvisionInstallerDevices(ctx, orgId, deviceMac).InstallerProvisionDevice(installerProvisionDevice).Execute()

provisionInstallerDevices



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
	deviceMac := "0000000000ab" // string | 
	installerProvisionDevice := *openapiclient.NewInstallerProvisionDevice("SJ1-AP1") // InstallerProvisionDevice | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.ProvisionInstallerDevices(context.Background(), orgId, deviceMac).InstallerProvisionDevice(installerProvisionDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.ProvisionInstallerDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisionInstallerDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **installerProvisionDevice** | [**InstallerProvisionDevice**](InstallerProvisionDevice.md) | Request Body | 

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


## StartInstallerLocateDevice

> StartInstallerLocateDevice(ctx, orgId, deviceMac).Execute()

startInstallerLocateDevice



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.StartInstallerLocateDevice(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.StartInstallerLocateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstallerLocateDeviceRequest struct via the builder pattern


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


## StopInstallerLocateDevice

> StopInstallerLocateDevice(ctx, orgId, deviceMac).Execute()

stopInstallerLocateDevice



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.StopInstallerLocateDevice(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.StopInstallerLocateDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstallerLocateDeviceRequest struct via the builder pattern


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


## UnassignInstallerRecentlyClaimedDevice

> UnassignInstallerRecentlyClaimedDevice(ctx, orgId, deviceMac).Execute()

unassignInstallerRecentlyClaimedDevice



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
	deviceMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstallerAPI.UnassignInstallerRecentlyClaimedDevice(context.Background(), orgId, deviceMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.UnassignInstallerRecentlyClaimedDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**deviceMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignInstallerRecentlyClaimedDeviceRequest struct via the builder pattern


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


## UpdateInstallerMap

> ModelMap UpdateInstallerMap(ctx, orgId, siteName, mapId).ModelMap(modelMap).Execute()

updateInstallerMap



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
	mapId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	modelMap := *openapiclient.NewMap() // ModelMap | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.UpdateInstallerMap(context.Background(), orgId, siteName, mapId).ModelMap(modelMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.UpdateInstallerMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstallerMap`: ModelMap
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.UpdateInstallerMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**siteName** | **string** |  | 
**mapId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstallerMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modelMap** | [**ModelMap**](ModelMap.md) | Request Body | 

### Return type

[**ModelMap**](ModelMap.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstallerVirtualChassisMember

> ResponseVirtualChassisConfig UpdateInstallerVirtualChassisMember(ctx, orgId, fpc0Mac).VirtualChassisUpdate(virtualChassisUpdate).Execute()

updateInstallerVirtualChassisMember



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
	fpc0Mac := "fpc0Mac_example" // string | FPC0 MAC Address
	virtualChassisUpdate := *openapiclient.NewVirtualChassisUpdate() // VirtualChassisUpdate | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallerAPI.UpdateInstallerVirtualChassisMember(context.Background(), orgId, fpc0Mac).VirtualChassisUpdate(virtualChassisUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallerAPI.UpdateInstallerVirtualChassisMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstallerVirtualChassisMember`: ResponseVirtualChassisConfig
	fmt.Fprintf(os.Stdout, "Response from `InstallerAPI.UpdateInstallerVirtualChassisMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**fpc0Mac** | **string** | FPC0 MAC Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstallerVirtualChassisMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **virtualChassisUpdate** | [**VirtualChassisUpdate**](VirtualChassisUpdate.md) | Request Body | 

### Return type

[**ResponseVirtualChassisConfig**](ResponseVirtualChassisConfig.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

