# \SitesEVPNTopologiesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSiteEvpnTopology**](SitesEVPNTopologiesAPI.md#CreateSiteEvpnTopology) | **Post** /api/v1/sites/{site_id}/evpn_topologies | createSiteEvpnTopology
[**DeleteSiteEvpnTopology**](SitesEVPNTopologiesAPI.md#DeleteSiteEvpnTopology) | **Delete** /api/v1/sites/{site_id}/evpn_topologies/{evpn_topology_id} | deleteSiteEvpnTopology
[**GetSiteEvpnTopology**](SitesEVPNTopologiesAPI.md#GetSiteEvpnTopology) | **Get** /api/v1/sites/{site_id}/evpn_topologies/{evpn_topology_id} | getSiteEvpnTopology
[**ListSiteEvpnTopologies**](SitesEVPNTopologiesAPI.md#ListSiteEvpnTopologies) | **Get** /api/v1/sites/{site_id}/evpn_topologies | listSiteEvpnTopologies
[**UpdateSiteEvpnTopology**](SitesEVPNTopologiesAPI.md#UpdateSiteEvpnTopology) | **Put** /api/v1/sites/{site_id}/evpn_topologies/{evpn_topology_id} | updateSiteEvpnTopology



## CreateSiteEvpnTopology

> EvpnTopology CreateSiteEvpnTopology(ctx, siteId).EvpnTopology(evpnTopology).Execute()

createSiteEvpnTopology



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
	evpnTopology := *openapiclient.NewEvpnTopology([]openapiclient.EvpnTopologySwitch{*openapiclient.NewEvpnTopologySwitch()}) // EvpnTopology |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesEVPNTopologiesAPI.CreateSiteEvpnTopology(context.Background(), siteId).EvpnTopology(evpnTopology).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEVPNTopologiesAPI.CreateSiteEvpnTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteEvpnTopology`: EvpnTopology
	fmt.Fprintf(os.Stdout, "Response from `SitesEVPNTopologiesAPI.CreateSiteEvpnTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteEvpnTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **evpnTopology** | [**EvpnTopology**](EvpnTopology.md) |  | 

### Return type

[**EvpnTopology**](EvpnTopology.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSiteEvpnTopology

> DeleteSiteEvpnTopology(ctx, siteId, evpnTopologyId).Execute()

deleteSiteEvpnTopology



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
	evpnTopologyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesEVPNTopologiesAPI.DeleteSiteEvpnTopology(context.Background(), siteId, evpnTopologyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEVPNTopologiesAPI.DeleteSiteEvpnTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**evpnTopologyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteEvpnTopologyRequest struct via the builder pattern


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


## GetSiteEvpnTopology

> GetSiteEvpnTopology(ctx, siteId, evpnTopologyId).Execute()

getSiteEvpnTopology



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
	evpnTopologyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesEVPNTopologiesAPI.GetSiteEvpnTopology(context.Background(), siteId, evpnTopologyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEVPNTopologiesAPI.GetSiteEvpnTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**evpnTopologyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteEvpnTopologyRequest struct via the builder pattern


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


## ListSiteEvpnTopologies

> EvpnTopology ListSiteEvpnTopologies(ctx, siteId).Execute()

listSiteEvpnTopologies



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesEVPNTopologiesAPI.ListSiteEvpnTopologies(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEVPNTopologiesAPI.ListSiteEvpnTopologies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteEvpnTopologies`: EvpnTopology
	fmt.Fprintf(os.Stdout, "Response from `SitesEVPNTopologiesAPI.ListSiteEvpnTopologies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteEvpnTopologiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EvpnTopology**](EvpnTopology.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteEvpnTopology

> EvpnTopology UpdateSiteEvpnTopology(ctx, siteId, evpnTopologyId).EvpnTopology(evpnTopology).Execute()

updateSiteEvpnTopology



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
	evpnTopologyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	evpnTopology := *openapiclient.NewEvpnTopology([]openapiclient.EvpnTopologySwitch{*openapiclient.NewEvpnTopologySwitch()}) // EvpnTopology |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesEVPNTopologiesAPI.UpdateSiteEvpnTopology(context.Background(), siteId, evpnTopologyId).EvpnTopology(evpnTopology).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesEVPNTopologiesAPI.UpdateSiteEvpnTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteEvpnTopology`: EvpnTopology
	fmt.Fprintf(os.Stdout, "Response from `SitesEVPNTopologiesAPI.UpdateSiteEvpnTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**evpnTopologyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteEvpnTopologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **evpnTopology** | [**EvpnTopology**](EvpnTopology.md) |  | 

### Return type

[**EvpnTopology**](EvpnTopology.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

