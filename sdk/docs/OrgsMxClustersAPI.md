# \OrgsMxClustersAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgMxEdgeCluster**](OrgsMxClustersAPI.md#CreateOrgMxEdgeCluster) | **Post** /api/v1/orgs/{org_id}/mxclusters | createOrgMxEdgeCluster
[**DeleteOrgMxEdgeCluster**](OrgsMxClustersAPI.md#DeleteOrgMxEdgeCluster) | **Delete** /api/v1/orgs/{org_id}/mxclusters/{mxcluster_id} | deleteOrgMxEdgeCluster
[**GetOrgMxEdgeCluster**](OrgsMxClustersAPI.md#GetOrgMxEdgeCluster) | **Get** /api/v1/orgs/{org_id}/mxclusters/{mxcluster_id} | getOrgMxEdgeCluster
[**ListOrgMxEdgeClusters**](OrgsMxClustersAPI.md#ListOrgMxEdgeClusters) | **Get** /api/v1/orgs/{org_id}/mxclusters | listOrgMxEdgeClusters
[**UpdateOrgMxEdgeCluster**](OrgsMxClustersAPI.md#UpdateOrgMxEdgeCluster) | **Put** /api/v1/orgs/{org_id}/mxclusters/{mxcluster_id} | updateOrgMxEdgeCluster



## CreateOrgMxEdgeCluster

> Mxcluster CreateOrgMxEdgeCluster(ctx, orgId).Mxcluster(mxcluster).Execute()

createOrgMxEdgeCluster



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
	mxcluster := *openapiclient.NewMxcluster() // Mxcluster | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxClustersAPI.CreateOrgMxEdgeCluster(context.Background(), orgId).Mxcluster(mxcluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxClustersAPI.CreateOrgMxEdgeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgMxEdgeCluster`: Mxcluster
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxClustersAPI.CreateOrgMxEdgeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgMxEdgeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mxcluster** | [**Mxcluster**](Mxcluster.md) | Request Body | 

### Return type

[**Mxcluster**](Mxcluster.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgMxEdgeCluster

> DeleteOrgMxEdgeCluster(ctx, orgId, mxclusterId).Execute()

deleteOrgMxEdgeCluster



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
	mxclusterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsMxClustersAPI.DeleteOrgMxEdgeCluster(context.Background(), orgId, mxclusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxClustersAPI.DeleteOrgMxEdgeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxclusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMxEdgeClusterRequest struct via the builder pattern


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


## GetOrgMxEdgeCluster

> Mxcluster GetOrgMxEdgeCluster(ctx, orgId, mxclusterId).Execute()

getOrgMxEdgeCluster



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
	mxclusterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxClustersAPI.GetOrgMxEdgeCluster(context.Background(), orgId, mxclusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxClustersAPI.GetOrgMxEdgeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMxEdgeCluster`: Mxcluster
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxClustersAPI.GetOrgMxEdgeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxclusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMxEdgeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Mxcluster**](Mxcluster.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMxEdgeClusters

> []Mxcluster ListOrgMxEdgeClusters(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgMxEdgeClusters



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
	resp, r, err := apiClient.OrgsMxClustersAPI.ListOrgMxEdgeClusters(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxClustersAPI.ListOrgMxEdgeClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgMxEdgeClusters`: []Mxcluster
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxClustersAPI.ListOrgMxEdgeClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMxEdgeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Mxcluster**](Mxcluster.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMxEdgeCluster

> Mxcluster UpdateOrgMxEdgeCluster(ctx, orgId, mxclusterId).Mxcluster(mxcluster).Execute()

updateOrgMxEdgeCluster



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
	mxclusterId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	mxcluster := *openapiclient.NewMxcluster() // Mxcluster | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsMxClustersAPI.UpdateOrgMxEdgeCluster(context.Background(), orgId, mxclusterId).Mxcluster(mxcluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsMxClustersAPI.UpdateOrgMxEdgeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgMxEdgeCluster`: Mxcluster
	fmt.Fprintf(os.Stdout, "Response from `OrgsMxClustersAPI.UpdateOrgMxEdgeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**mxclusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMxEdgeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mxcluster** | [**Mxcluster**](Mxcluster.md) | Request Body | 

### Return type

[**Mxcluster**](Mxcluster.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

