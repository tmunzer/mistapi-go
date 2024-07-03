# \OrgsInventoryAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgInventory**](OrgsInventoryAPI.md#AddOrgInventory) | **Post** /api/v1/orgs/{org_id}/inventory | addOrgInventory
[**GetOrgInventory**](OrgsInventoryAPI.md#GetOrgInventory) | **Get** /api/v1/orgs/{org_id}/inventory | getOrgInventory
[**ReevaluateOrgAutoAssignment**](OrgsInventoryAPI.md#ReevaluateOrgAutoAssignment) | **Post** /api/v1/orgs/{org_id}/inventory/reevaluate_auto_assignment | reevaluateOrgAutoAssignment
[**ReplaceOrgDevices**](OrgsInventoryAPI.md#ReplaceOrgDevices) | **Post** /api/v1/orgs/{org_id}/inventory/replace | replaceOrgDevices
[**UpdateOrgInventoryAssignment**](OrgsInventoryAPI.md#UpdateOrgInventoryAssignment) | **Put** /api/v1/orgs/{org_id}/inventory | updateOrgInventoryAssignment



## AddOrgInventory

> ResponseInventory AddOrgInventory(ctx, orgId).RequestBody(requestBody).Execute()

addOrgInventory



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
	resp, r, err := apiClient.OrgsInventoryAPI.AddOrgInventory(context.Background(), orgId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsInventoryAPI.AddOrgInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOrgInventory`: ResponseInventory
	fmt.Fprintf(os.Stdout, "Response from `OrgsInventoryAPI.AddOrgInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrgInventoryRequest struct via the builder pattern


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


## GetOrgInventory

> []Inventory GetOrgInventory(ctx, orgId).Serial(serial).Model(model).Type_(type_).Mac(mac).SiteId(siteId).VcMac(vcMac).Vc(vc).Unassigned(unassigned).Limit(limit).Page(page).Execute()

getOrgInventory



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
	serial := "serial_example" // string | device serial (optional)
	model := "model_example" // string | device model (optional)
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	mac := "mac_example" // string | MAC address (optional)
	siteId := "siteId_example" // string | site id if assigned, null if not assigned (optional)
	vcMac := "vcMac_example" // string | Virtual Chassis MAC Address (optional)
	vc := true // bool | To display Virtual Chassis members (optional) (default to false)
	unassigned := true // bool | to display Unassigned devices (optional) (default to true)
	limit := int32(56) // int32 |  (optional) (default to 100)
	page := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsInventoryAPI.GetOrgInventory(context.Background(), orgId).Serial(serial).Model(model).Type_(type_).Mac(mac).SiteId(siteId).VcMac(vcMac).Vc(vc).Unassigned(unassigned).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsInventoryAPI.GetOrgInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgInventory`: []Inventory
	fmt.Fprintf(os.Stdout, "Response from `OrgsInventoryAPI.GetOrgInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serial** | **string** | device serial | 
 **model** | **string** | device model | 
 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **mac** | **string** | MAC address | 
 **siteId** | **string** | site id if assigned, null if not assigned | 
 **vcMac** | **string** | Virtual Chassis MAC Address | 
 **vc** | **bool** | To display Virtual Chassis members | [default to false]
 **unassigned** | **bool** | to display Unassigned devices | [default to true]
 **limit** | **int32** |  | [default to 100]
 **page** | **int32** |  | [default to 1]

### Return type

[**[]Inventory**](Inventory.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReevaluateOrgAutoAssignment

> ReevaluateOrgAutoAssignment(ctx, orgId).Execute()

reevaluateOrgAutoAssignment



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
	r, err := apiClient.OrgsInventoryAPI.ReevaluateOrgAutoAssignment(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsInventoryAPI.ReevaluateOrgAutoAssignment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReevaluateOrgAutoAssignmentRequest struct via the builder pattern


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


## ReplaceOrgDevices

> ResponseOrgInventoryChange ReplaceOrgDevices(ctx, orgId).ReplaceDevice(replaceDevice).Execute()

replaceOrgDevices



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
	replaceDevice := *openapiclient.NewReplaceDevice() // ReplaceDevice | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsInventoryAPI.ReplaceOrgDevices(context.Background(), orgId).ReplaceDevice(replaceDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsInventoryAPI.ReplaceOrgDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceOrgDevices`: ResponseOrgInventoryChange
	fmt.Fprintf(os.Stdout, "Response from `OrgsInventoryAPI.ReplaceOrgDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrgDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceDevice** | [**ReplaceDevice**](ReplaceDevice.md) | Request Body | 

### Return type

[**ResponseOrgInventoryChange**](ResponseOrgInventoryChange.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgInventoryAssignment

> ResponseOrgInventoryChange UpdateOrgInventoryAssignment(ctx, orgId).InventoryUpdate(inventoryUpdate).Execute()

updateOrgInventoryAssignment



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
	inventoryUpdate := *openapiclient.NewInventoryUpdate(openapiclient.inventory_update_operation("")) // InventoryUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsInventoryAPI.UpdateOrgInventoryAssignment(context.Background(), orgId).InventoryUpdate(inventoryUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsInventoryAPI.UpdateOrgInventoryAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgInventoryAssignment`: ResponseOrgInventoryChange
	fmt.Fprintf(os.Stdout, "Response from `OrgsInventoryAPI.UpdateOrgInventoryAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgInventoryAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inventoryUpdate** | [**InventoryUpdate**](InventoryUpdate.md) |  | 

### Return type

[**ResponseOrgInventoryChange**](ResponseOrgInventoryChange.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

