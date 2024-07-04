# \UtilitiesUpgradeAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrgSsrUpgrade**](UtilitiesUpgradeAPI.md#CancelOrgSsrUpgrade) | **Post** /api/v1/orgs/{org_id}/ssr/upgrade/{upgrade_id}/cancel | cancelOrgSsrUpgrade
[**CancelSiteDeviceUpgrade**](UtilitiesUpgradeAPI.md#CancelSiteDeviceUpgrade) | **Post** /api/v1/sites/{site_id}/devices/upgrade/{upgrade_id}/cancel | cancelSiteDeviceUpgrade
[**GetOrgDeviceUpgrade**](UtilitiesUpgradeAPI.md#GetOrgDeviceUpgrade) | **Get** /api/v1/orgs/{org_id}/devices/upgrade/{upgrade_id} | getOrgDeviceUpgrade
[**GetOrgMxEdgeUpgrade**](UtilitiesUpgradeAPI.md#GetOrgMxEdgeUpgrade) | **Get** /api/v1/orgs/{org_id}/mxedges/upgrade/{upgrade_id} | getOrgMxEdgeUpgrade
[**GetSiteDeviceUpgrade**](UtilitiesUpgradeAPI.md#GetSiteDeviceUpgrade) | **Get** /api/v1/sites/{site_id}/devices/upgrade/{upgrade_id} | getSiteDeviceUpgrade
[**GetSiteSsrUpgrade**](UtilitiesUpgradeAPI.md#GetSiteSsrUpgrade) | **Get** /api/v1/sites/{site_id}/ssr/upgrade/{upgrade_id} | getSiteSsrUpgrade
[**ListOrgAvailableSsrVersions**](UtilitiesUpgradeAPI.md#ListOrgAvailableSsrVersions) | **Get** /api/v1/orgs/{org_id}/ssr/versions | listOrgAvailableSsrVersions
[**ListOrgDeviceUpgrades**](UtilitiesUpgradeAPI.md#ListOrgDeviceUpgrades) | **Get** /api/v1/orgs/{org_id}/devices/upgrade | listOrgDeviceUpgrades
[**ListOrgMxEdgeUpgrades**](UtilitiesUpgradeAPI.md#ListOrgMxEdgeUpgrades) | **Get** /api/v1/orgs/{org_id}/mxedges/upgrade | listOrgMxEdgeUpgrades
[**ListOrgSsrUpgrades**](UtilitiesUpgradeAPI.md#ListOrgSsrUpgrades) | **Get** /api/v1/orgs/{org_id}/ssr/upgrade | listOrgSsrUpgrades
[**ListSiteAvailableDeviceVersions**](UtilitiesUpgradeAPI.md#ListSiteAvailableDeviceVersions) | **Get** /api/v1/sites/{site_id}/devices/versions | listSiteAvailableDeviceVersions
[**ListSiteDeviceUpgrades**](UtilitiesUpgradeAPI.md#ListSiteDeviceUpgrades) | **Get** /api/v1/sites/{site_id}/devices/upgrade | listSiteDeviceUpgrades
[**UpgradeDevice**](UtilitiesUpgradeAPI.md#UpgradeDevice) | **Post** /api/v1/sites/{site_id}/devices/{device_id}/upgrade | upgradeDevice
[**UpgradeOrgDevices**](UtilitiesUpgradeAPI.md#UpgradeOrgDevices) | **Post** /api/v1/orgs/{org_id}/devices/upgrade | upgradeOrgDevices
[**UpgradeOrgJsiDevice**](UtilitiesUpgradeAPI.md#UpgradeOrgJsiDevice) | **Post** /api/v1/orgs/{org_id}/jsi/devices/{device_mac}/upgrade | upgradeOrgJsiDevice
[**UpgradeOrgMxEdges**](UtilitiesUpgradeAPI.md#UpgradeOrgMxEdges) | **Post** /api/v1/orgs/{org_id}/mxedges/upgrade | upgradeOrgMxEdges
[**UpgradeOrgSsrs**](UtilitiesUpgradeAPI.md#UpgradeOrgSsrs) | **Post** /api/v1/orgs/{org_id}/ssr/upgrade | upgradeOrgSsrs
[**UpgradeSiteDevices**](UtilitiesUpgradeAPI.md#UpgradeSiteDevices) | **Post** /api/v1/sites/{site_id}/devices/upgrade | upgradeSiteDevices
[**UpgradeSsr**](UtilitiesUpgradeAPI.md#UpgradeSsr) | **Post** /api/v1/sites/{site_id}/ssr/{device_id}/upgrade | upgradeSsr



## CancelOrgSsrUpgrade

> CancelOrgSsrUpgrade(ctx, orgId, upgradeId).Execute()

cancelOrgSsrUpgrade



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
	upgradeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesUpgradeAPI.CancelOrgSsrUpgrade(context.Background(), orgId, upgradeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.CancelOrgSsrUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**upgradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrgSsrUpgradeRequest struct via the builder pattern


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


## CancelSiteDeviceUpgrade

> CancelSiteDeviceUpgrade(ctx, siteId, upgradeId).Execute()

cancelSiteDeviceUpgrade



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	upgradeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesUpgradeAPI.CancelSiteDeviceUpgrade(context.Background(), siteId, upgradeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.CancelSiteDeviceUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**upgradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSiteDeviceUpgradeRequest struct via the builder pattern


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


## GetOrgDeviceUpgrade

> ResponseUpgradeOrgDevices GetOrgDeviceUpgrade(ctx, orgId, upgradeId).Execute()

getOrgDeviceUpgrade



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
	upgradeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.GetOrgDeviceUpgrade(context.Background(), orgId, upgradeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.GetOrgDeviceUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgDeviceUpgrade`: ResponseUpgradeOrgDevices
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.GetOrgDeviceUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**upgradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgDeviceUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseUpgradeOrgDevices**](ResponseUpgradeOrgDevices.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgMxEdgeUpgrade

> ResponseMxedgeUpgrade GetOrgMxEdgeUpgrade(ctx, orgId, upgradeId).Execute()

getOrgMxEdgeUpgrade



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
	upgradeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.GetOrgMxEdgeUpgrade(context.Background(), orgId, upgradeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.GetOrgMxEdgeUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMxEdgeUpgrade`: ResponseMxedgeUpgrade
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.GetOrgMxEdgeUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**upgradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMxEdgeUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseMxedgeUpgrade**](ResponseMxedgeUpgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteDeviceUpgrade

> ResponseDeviceUpgrade GetSiteDeviceUpgrade(ctx, siteId, upgradeId).Execute()

getSiteDeviceUpgrade



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	upgradeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.GetSiteDeviceUpgrade(context.Background(), siteId, upgradeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.GetSiteDeviceUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteDeviceUpgrade`: ResponseDeviceUpgrade
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.GetSiteDeviceUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**upgradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteDeviceUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseDeviceUpgrade**](ResponseDeviceUpgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSsrUpgrade

> ResponseSsrUpgradeStatus GetSiteSsrUpgrade(ctx, siteId, upgradeId).Execute()

getSiteSsrUpgrade



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	upgradeId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.GetSiteSsrUpgrade(context.Background(), siteId, upgradeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.GetSiteSsrUpgrade``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSsrUpgrade`: ResponseSsrUpgradeStatus
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.GetSiteSsrUpgrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**upgradeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSsrUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ResponseSsrUpgradeStatus**](ResponseSsrUpgradeStatus.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgAvailableSsrVersions

> []SsrVersion ListOrgAvailableSsrVersions(ctx, orgId).Channel(channel).Execute()

listOrgAvailableSsrVersions



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
	channel := "channel_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.ListOrgAvailableSsrVersions(context.Background(), orgId).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.ListOrgAvailableSsrVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgAvailableSsrVersions`: []SsrVersion
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.ListOrgAvailableSsrVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgAvailableSsrVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channel** | **string** |  | 

### Return type

[**[]SsrVersion**](SsrVersion.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgDeviceUpgrades

> []OrgDeviceUpgrade ListOrgDeviceUpgrades(ctx, orgId).Execute()

listOrgDeviceUpgrades



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
	resp, r, err := apiClient.UtilitiesUpgradeAPI.ListOrgDeviceUpgrades(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.ListOrgDeviceUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgDeviceUpgrades`: []OrgDeviceUpgrade
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.ListOrgDeviceUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgDeviceUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]OrgDeviceUpgrade**](OrgDeviceUpgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMxEdgeUpgrades

> []ResponseMxedgeUpgrade ListOrgMxEdgeUpgrades(ctx, orgId).Execute()

listOrgMxEdgeUpgrades



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
	resp, r, err := apiClient.UtilitiesUpgradeAPI.ListOrgMxEdgeUpgrades(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.ListOrgMxEdgeUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgMxEdgeUpgrades`: []ResponseMxedgeUpgrade
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.ListOrgMxEdgeUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMxEdgeUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ResponseMxedgeUpgrade**](ResponseMxedgeUpgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgSsrUpgrades

> []SsrUpgradeResponse ListOrgSsrUpgrades(ctx, orgId).Execute()

listOrgSsrUpgrades



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
	resp, r, err := apiClient.UtilitiesUpgradeAPI.ListOrgSsrUpgrades(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.ListOrgSsrUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgSsrUpgrades`: []SsrUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.ListOrgSsrUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSsrUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SsrUpgradeResponse**](SsrUpgradeResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteAvailableDeviceVersions

> []DeviceVersionItem ListSiteAvailableDeviceVersions(ctx, siteId).Type_(type_).Model(model).Execute()

listSiteAvailableDeviceVersions



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	type_ := openapiclient.device_type("") // DeviceType |  (optional) (default to "ap")
	model := "model_example" // string | fetch version for device model, use/combine with `type` as needed (for switch and gateway devices) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.ListSiteAvailableDeviceVersions(context.Background(), siteId).Type_(type_).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.ListSiteAvailableDeviceVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteAvailableDeviceVersions`: []DeviceVersionItem
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.ListSiteAvailableDeviceVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteAvailableDeviceVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**DeviceType**](DeviceType.md) |  | [default to &quot;ap&quot;]
 **model** | **string** | fetch version for device model, use/combine with &#x60;type&#x60; as needed (for switch and gateway devices) | 

### Return type

[**[]DeviceVersionItem**](DeviceVersionItem.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteDeviceUpgrades

> []ResponseSiteDeviceUpgrade ListSiteDeviceUpgrades(ctx, siteId).Status(status).Execute()

listSiteDeviceUpgrades



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	status := openapiclient.device_upgrade_status("") // DeviceUpgradeStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.ListSiteDeviceUpgrades(context.Background(), siteId).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.ListSiteDeviceUpgrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteDeviceUpgrades`: []ResponseSiteDeviceUpgrade
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.ListSiteDeviceUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteDeviceUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**DeviceUpgradeStatus**](DeviceUpgradeStatus.md) |  | 

### Return type

[**[]ResponseSiteDeviceUpgrade**](ResponseSiteDeviceUpgrade.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeDevice

> ResponseUpgradeDevice UpgradeDevice(ctx, siteId, deviceId).DeviceUpgrade(deviceUpgrade).Execute()

upgradeDevice



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceUpgrade := *openapiclient.NewDeviceUpgrade("Version_example") // DeviceUpgrade |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.UpgradeDevice(context.Background(), siteId, deviceId).DeviceUpgrade(deviceUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.UpgradeDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeDevice`: ResponseUpgradeDevice
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.UpgradeDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceUpgrade** | [**DeviceUpgrade**](DeviceUpgrade.md) |  | 

### Return type

[**ResponseUpgradeDevice**](ResponseUpgradeDevice.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeOrgDevices

> ResponseUpgradeOrgDevices UpgradeOrgDevices(ctx, orgId).UpgradeOrgDevices(upgradeOrgDevices).Execute()

upgradeOrgDevices



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
	upgradeOrgDevices := *openapiclient.NewUpgradeOrgDevices() // UpgradeOrgDevices |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.UpgradeOrgDevices(context.Background(), orgId).UpgradeOrgDevices(upgradeOrgDevices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.UpgradeOrgDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeOrgDevices`: ResponseUpgradeOrgDevices
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.UpgradeOrgDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeOrgDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeOrgDevices** | [**UpgradeOrgDevices**](UpgradeOrgDevices.md) |  | 

### Return type

[**ResponseUpgradeOrgDevices**](ResponseUpgradeOrgDevices.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeOrgJsiDevice

> UpgradeOrgJsiDevice(ctx, orgId, deviceMac).VersionString(versionString).Execute()

upgradeOrgJsiDevice



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
	versionString := *openapiclient.NewVersionString() // VersionString |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesUpgradeAPI.UpgradeOrgJsiDevice(context.Background(), orgId, deviceMac).VersionString(versionString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.UpgradeOrgJsiDevice``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpgradeOrgJsiDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **versionString** | [**VersionString**](VersionString.md) |  | 

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


## UpgradeOrgMxEdges

> UpgradeOrgMxEdges(ctx, orgId).MxedgeUpgradeMulti(mxedgeUpgradeMulti).Execute()

upgradeOrgMxEdges



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
	mxedgeUpgradeMulti := *openapiclient.NewMxedgeUpgradeMulti([]string{"MxedgeIds_example"}) // MxedgeUpgradeMulti | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesUpgradeAPI.UpgradeOrgMxEdges(context.Background(), orgId).MxedgeUpgradeMulti(mxedgeUpgradeMulti).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.UpgradeOrgMxEdges``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpgradeOrgMxEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxedgeUpgradeMulti** | [**MxedgeUpgradeMulti**](MxedgeUpgradeMulti.md) | Request Body | 

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


## UpgradeOrgSsrs

> SsrUpgradeResponse UpgradeOrgSsrs(ctx, orgId).SsrUpgradeMulti(ssrUpgradeMulti).Execute()

upgradeOrgSsrs



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
	ssrUpgradeMulti := *openapiclient.NewSsrUpgradeMulti([]string{"DeviceIds_example"}) // SsrUpgradeMulti |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.UpgradeOrgSsrs(context.Background(), orgId).SsrUpgradeMulti(ssrUpgradeMulti).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.UpgradeOrgSsrs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeOrgSsrs`: SsrUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.UpgradeOrgSsrs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeOrgSsrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ssrUpgradeMulti** | [**SsrUpgradeMulti**](SsrUpgradeMulti.md) |  | 

### Return type

[**SsrUpgradeResponse**](SsrUpgradeResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSiteDevices

> ResponseUpgradeSiteDevices UpgradeSiteDevices(ctx, siteId).UpgradeSiteDevices(upgradeSiteDevices).Execute()

upgradeSiteDevices



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	upgradeSiteDevices := *openapiclient.NewUpgradeSiteDevices() // UpgradeSiteDevices | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.UpgradeSiteDevices(context.Background(), siteId).UpgradeSiteDevices(upgradeSiteDevices).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.UpgradeSiteDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeSiteDevices`: ResponseUpgradeSiteDevices
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.UpgradeSiteDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSiteDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upgradeSiteDevices** | [**UpgradeSiteDevices**](UpgradeSiteDevices.md) | Request Body | 

### Return type

[**ResponseUpgradeSiteDevices**](ResponseUpgradeSiteDevices.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSsr

> SsrUpgradeResponse UpgradeSsr(ctx, siteId, deviceId).SsrUpgrade(ssrUpgrade).Execute()

upgradeSsr



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	deviceId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ssrUpgrade := *openapiclient.NewSsrUpgrade("Version_example") // SsrUpgrade |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UtilitiesUpgradeAPI.UpgradeSsr(context.Background(), siteId, deviceId).SsrUpgrade(ssrUpgrade).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesUpgradeAPI.UpgradeSsr``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpgradeSsr`: SsrUpgradeResponse
	fmt.Fprintf(os.Stdout, "Response from `UtilitiesUpgradeAPI.UpgradeSsr`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSsrRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ssrUpgrade** | [**SsrUpgrade**](SsrUpgrade.md) |  | 

### Return type

[**SsrUpgradeResponse**](SsrUpgradeResponse.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

