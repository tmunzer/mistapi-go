# \OrgsEVPNTopologiesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgEvpnTopology**](OrgsEVPNTopologiesAPI.md#CreateOrgEvpnTopology) | **Post** /api/v1/orgs/{org_id}/evpn_topologies | createOrgEvpnTopology
[**DeleteOrgEvpnTopology**](OrgsEVPNTopologiesAPI.md#DeleteOrgEvpnTopology) | **Delete** /api/v1/orgs/{org_id}/evpn_topologies/{evpn_topology_id} | deleteOrgEvpnTopology
[**GetOrgEvpnTolopogy**](OrgsEVPNTopologiesAPI.md#GetOrgEvpnTolopogy) | **Get** /api/v1/orgs/{org_id}/evpn_topologies/{evpn_topology_id} | getOrgEvpnTolopogy
[**ListOrgEvpnTopologies**](OrgsEVPNTopologiesAPI.md#ListOrgEvpnTopologies) | **Get** /api/v1/orgs/{org_id}/evpn_topologies | listOrgEvpnTopologies
[**UpdateOrgEvpnTopology**](OrgsEVPNTopologiesAPI.md#UpdateOrgEvpnTopology) | **Put** /api/v1/orgs/{org_id}/evpn_topologies/{evpn_topology_id} | updateOrgEvpnTopology



## CreateOrgEvpnTopology

> EvpnTopology CreateOrgEvpnTopology(ctx, orgId).EvpnTopology(evpnTopology).Execute()

createOrgEvpnTopology



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
	evpnTopology := *openapiclient.NewEvpnTopology([]openapiclient.EvpnTopologySwitch{*openapiclient.NewEvpnTopologySwitch()}) // EvpnTopology |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsEVPNTopologiesAPI.CreateOrgEvpnTopology(context.Background(), orgId).EvpnTopology(evpnTopology).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsEVPNTopologiesAPI.CreateOrgEvpnTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgEvpnTopology`: EvpnTopology
	fmt.Fprintf(os.Stdout, "Response from `OrgsEVPNTopologiesAPI.CreateOrgEvpnTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgEvpnTopologyRequest struct via the builder pattern


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


## DeleteOrgEvpnTopology

> DeleteOrgEvpnTopology(ctx, orgId, evpnTopologyId).Execute()

deleteOrgEvpnTopology



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
	evpnTopologyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsEVPNTopologiesAPI.DeleteOrgEvpnTopology(context.Background(), orgId, evpnTopologyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsEVPNTopologiesAPI.DeleteOrgEvpnTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**evpnTopologyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgEvpnTopologyRequest struct via the builder pattern


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


## GetOrgEvpnTolopogy

> EvpnTopology GetOrgEvpnTolopogy(ctx, orgId, evpnTopologyId).Execute()

getOrgEvpnTolopogy



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
	evpnTopologyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsEVPNTopologiesAPI.GetOrgEvpnTolopogy(context.Background(), orgId, evpnTopologyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsEVPNTopologiesAPI.GetOrgEvpnTolopogy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgEvpnTolopogy`: EvpnTopology
	fmt.Fprintf(os.Stdout, "Response from `OrgsEVPNTopologiesAPI.GetOrgEvpnTolopogy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**evpnTopologyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgEvpnTolopogyRequest struct via the builder pattern


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


## ListOrgEvpnTopologies

> []EvpnTopology ListOrgEvpnTopologies(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgEvpnTopologies



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsEVPNTopologiesAPI.ListOrgEvpnTopologies(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsEVPNTopologiesAPI.ListOrgEvpnTopologies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgEvpnTopologies`: []EvpnTopology
	fmt.Fprintf(os.Stdout, "Response from `OrgsEVPNTopologiesAPI.ListOrgEvpnTopologies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgEvpnTopologiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]EvpnTopology**](EvpnTopology.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgEvpnTopology

> EvpnTopology UpdateOrgEvpnTopology(ctx, orgId, evpnTopologyId).EvpnTopology(evpnTopology).Execute()

updateOrgEvpnTopology



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
	evpnTopologyId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	evpnTopology := *openapiclient.NewEvpnTopology([]openapiclient.EvpnTopologySwitch{*openapiclient.NewEvpnTopologySwitch()}) // EvpnTopology |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsEVPNTopologiesAPI.UpdateOrgEvpnTopology(context.Background(), orgId, evpnTopologyId).EvpnTopology(evpnTopology).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsEVPNTopologiesAPI.UpdateOrgEvpnTopology``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgEvpnTopology`: EvpnTopology
	fmt.Fprintf(os.Stdout, "Response from `OrgsEVPNTopologiesAPI.UpdateOrgEvpnTopology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**evpnTopologyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgEvpnTopologyRequest struct via the builder pattern


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

