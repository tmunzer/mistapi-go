# \SitesSLEsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSiteSleClassifierDetails**](SitesSLEsAPI.md#GetSiteSleClassifierDetails) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/classifier/{classifier}/summary | getSiteSleClassifierDetails
[**GetSiteSleHistogram**](SitesSLEsAPI.md#GetSiteSleHistogram) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/histogram | getSiteSleHistogram
[**GetSiteSleImpactSummary**](SitesSLEsAPI.md#GetSiteSleImpactSummary) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impact-summary | getSiteSleImpactSummary
[**GetSiteSleImpactedApplications**](SitesSLEsAPI.md#GetSiteSleImpactedApplications) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-applications | getSiteSleImpactedApplications
[**GetSiteSleImpactedAps**](SitesSLEsAPI.md#GetSiteSleImpactedAps) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-aps | getSiteSleImpactedAps
[**GetSiteSleImpactedChassis**](SitesSLEsAPI.md#GetSiteSleImpactedChassis) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-chassis | getSiteSleImpactedChassis
[**GetSiteSleImpactedGateways**](SitesSLEsAPI.md#GetSiteSleImpactedGateways) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-gateways | getSiteSleImpactedGateways
[**GetSiteSleImpactedInterfaces**](SitesSLEsAPI.md#GetSiteSleImpactedInterfaces) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-interfaces | getSiteSleImpactedInterfaces
[**GetSiteSleImpactedSwitches**](SitesSLEsAPI.md#GetSiteSleImpactedSwitches) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-switches | getSiteSleImpactedSwitches
[**GetSiteSleImpactedWiredClients**](SitesSLEsAPI.md#GetSiteSleImpactedWiredClients) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted-clients | getSiteSleImpactedWiredClients
[**GetSiteSleImpactedWirelessClients**](SitesSLEsAPI.md#GetSiteSleImpactedWirelessClients) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/impacted_users | getSiteSleImpactedWirelessClients
[**GetSiteSleMetricClassifiers**](SitesSLEsAPI.md#GetSiteSleMetricClassifiers) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/classifiers | getSiteSleMetricClassifiers
[**GetSiteSleSummary**](SitesSLEsAPI.md#GetSiteSleSummary) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/summary | getSiteSleSummary
[**GetSiteSleThreshold**](SitesSLEsAPI.md#GetSiteSleThreshold) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/threshold | getSiteSleThreshold
[**GetSiteSlesMetrics**](SitesSLEsAPI.md#GetSiteSlesMetrics) | **Get** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | getSiteSlesMetrics
[**ReplaceSiteSleThreshold**](SitesSLEsAPI.md#ReplaceSiteSleThreshold) | **Post** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/threshold | replaceSiteSleThreshold
[**UpdateSiteSleThreshold**](SitesSLEsAPI.md#UpdateSiteSleThreshold) | **Put** /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metric/{metric}/threshold | updateSiteSleThreshold



## GetSiteSleClassifierDetails

> SleClassifierSummary GetSiteSleClassifierDetails(ctx, siteId, scope, scopeId, metric, classifier).Start(start).End(end).Duration(duration).Execute()

getSiteSleClassifierDetails



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.sle_summary_scope("") // SleSummaryScope | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	classifier := "classifier_example" // string | 
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleClassifierDetails(context.Background(), siteId, scope, scopeId, metric, classifier).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleClassifierDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleClassifierDetails`: SleClassifierSummary
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleClassifierDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SleSummaryScope**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 
**classifier** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleClassifierDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**SleClassifierSummary**](SleClassifierSummary.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleHistogram

> SleHistogram GetSiteSleHistogram(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Execute()

getSiteSleHistogram



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_histogram_scope_parameters("") // SiteSleHistogramScopeParameters | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleHistogram(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleHistogram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleHistogram`: SleHistogram
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleHistogram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleHistogramScopeParameters**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**SleHistogram**](SleHistogram.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactSummary

> SleImpactSummary GetSiteSleImpactSummary(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Fields(fields).Classifier(classifier).Execute()

getSiteSleImpactSummary



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impact_summary_scope_parameters("") // SiteSleImpactSummaryScopeParameters | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	fields := openapiclient.site_sle_impact_summary_fields_parameter("") // SiteSleImpactSummaryFieldsParameter |  (optional)
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactSummary(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Fields(fields).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactSummary`: SleImpactSummary
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactSummaryScopeParameters**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **fields** | [**SiteSleImpactSummaryFieldsParameter**](SiteSleImpactSummaryFieldsParameter.md) |  | 
 **classifier** | **string** |  | 

### Return type

[**SleImpactSummary**](SleImpactSummary.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedApplications

> SleImpactedApplications GetSiteSleImpactedApplications(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedApplications



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_scope("") // SiteSleScope | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedApplications(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedApplications`: SleImpactedApplications
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleScope**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedApplications**](SleImpactedApplications.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedAps

> SleImpactedAps GetSiteSleImpactedAps(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedAps



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impacted_aps_scope_parameters("") // SiteSleImpactedApsScopeParameters | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedAps(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedAps`: SleImpactedAps
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedAps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactedApsScopeParameters**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedAps**](SleImpactedAps.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedChassis

> SleImpactedChassis GetSiteSleImpactedChassis(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedChassis



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impacted_chassis_scope_parameters("") // SiteSleImpactedChassisScopeParameters | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedChassis(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedChassis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedChassis`: SleImpactedChassis
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedChassis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactedChassisScopeParameters**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedChassisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedChassis**](SleImpactedChassis.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedGateways

> SleImpactedGateways GetSiteSleImpactedGateways(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedGateways



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impacted_gateways_scope_parameters("") // SiteSleImpactedGatewaysScopeParameters | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedGateways(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedGateways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedGateways`: SleImpactedGateways
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactedGatewaysScopeParameters**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedGateways**](SleImpactedGateways.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedInterfaces

> SleImpactedInterfaces GetSiteSleImpactedInterfaces(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedInterfaces



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impacted_interfaces_scope_parameters("") // SiteSleImpactedInterfacesScopeParameters | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedInterfaces(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedInterfaces`: SleImpactedInterfaces
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactedInterfacesScopeParameters**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedInterfaces**](SleImpactedInterfaces.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedSwitches

> SleImpactedSwitches GetSiteSleImpactedSwitches(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedSwitches



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impacted_switches_scope_parameters("") // SiteSleImpactedSwitchesScopeParameters | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedSwitches(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedSwitches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedSwitches`: SleImpactedSwitches
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedSwitches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactedSwitchesScopeParameters**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedSwitchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedSwitches**](SleImpactedSwitches.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedWiredClients

> SleImpactedClients GetSiteSleImpactedWiredClients(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedWiredClients



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impacted_clients_scope_parameters("") // SiteSleImpactedClientsScopeParameters | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedWiredClients(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedWiredClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedWiredClients`: SleImpactedClients
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedWiredClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactedClientsScopeParameters**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedWiredClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedClients**](SleImpactedClients.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleImpactedWirelessClients

> SleImpactedUsers GetSiteSleImpactedWirelessClients(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()

getSiteSleImpactedWirelessClients



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_impacted_users_scope_parameter("") // SiteSleImpactedUsersScopeParameter | 
	scopeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	classifier := "classifier_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleImpactedWirelessClients(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Classifier(classifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleImpactedWirelessClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleImpactedWirelessClients`: SleImpactedUsers
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleImpactedWirelessClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleImpactedUsersScopeParameter**](.md) |  | 
**scopeId** | **string** |  | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleImpactedWirelessClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **classifier** | **string** |  | 

### Return type

[**SleImpactedUsers**](SleImpactedUsers.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleMetricClassifiers

> []string GetSiteSleMetricClassifiers(ctx, siteId, scope, scopeId, metric).Execute()

getSiteSleMetricClassifiers



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_metric_classifiers_scope_parameters("") // SiteSleMetricClassifiersScopeParameters | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleMetricClassifiers(context.Background(), siteId, scope, scopeId, metric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleMetricClassifiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleMetricClassifiers`: []string
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleMetricClassifiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleMetricClassifiersScopeParameters**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleMetricClassifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

**[]string**

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleSummary

> SleSummary GetSiteSleSummary(ctx, siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Execute()

getSiteSleSummary



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_metric_summary_scope_parameters("") // SiteSleMetricSummaryScopeParameters | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleSummary(context.Background(), siteId, scope, scopeId, metric).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleSummary`: SleSummary
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleMetricSummaryScopeParameters**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**SleSummary**](SleSummary.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSleThreshold

> SleThreshold GetSiteSleThreshold(ctx, siteId, scope, scopeId, metric).Execute()

getSiteSleThreshold



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_threshold_scope_parameter("") // SiteSleThresholdScopeParameter | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSleThreshold(context.Background(), siteId, scope, scopeId, metric).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSleThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSleThreshold`: SleThreshold
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSleThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleThresholdScopeParameter**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSleThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**SleThreshold**](SleThreshold.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSiteSlesMetrics

> SiteSleMetrics GetSiteSlesMetrics(ctx, siteId, scope, scopeId).Execute()

getSiteSlesMetrics



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_metrics_scope_parameters("") // SiteSleMetricsScopeParameters | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.GetSiteSlesMetrics(context.Background(), siteId, scope, scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.GetSiteSlesMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteSlesMetrics`: SiteSleMetrics
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.GetSiteSlesMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleMetricsScopeParameters**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteSlesMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SiteSleMetrics**](SiteSleMetrics.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceSiteSleThreshold

> SleThreshold ReplaceSiteSleThreshold(ctx, siteId, scope, scopeId, metric).SleThreshold(sleThreshold).Execute()

replaceSiteSleThreshold



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_threshold_scope_parameter("") // SiteSleThresholdScopeParameter | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	sleThreshold := *openapiclient.NewSleThreshold() // SleThreshold |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.ReplaceSiteSleThreshold(context.Background(), siteId, scope, scopeId, metric).SleThreshold(sleThreshold).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.ReplaceSiteSleThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceSiteSleThreshold`: SleThreshold
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.ReplaceSiteSleThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleThresholdScopeParameter**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceSiteSleThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sleThreshold** | [**SleThreshold**](SleThreshold.md) |  | 

### Return type

[**SleThreshold**](SleThreshold.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSiteSleThreshold

> SleThreshold UpdateSiteSleThreshold(ctx, siteId, scope, scopeId, metric).SleThreshold(sleThreshold).Execute()

updateSiteSleThreshold



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
	siteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	scope := openapiclient.site_sle_threshold_scope_parameter("") // SiteSleThresholdScopeParameter | 
	scopeId := "scopeId_example" // string | * site_id if `scope`==`site` * device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway` * mac if `scope`==`client`
	metric := "metric_example" // string | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics
	sleThreshold := *openapiclient.NewSleThreshold() // SleThreshold |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesSLEsAPI.UpdateSiteSleThreshold(context.Background(), siteId, scope, scopeId, metric).SleThreshold(sleThreshold).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesSLEsAPI.UpdateSiteSleThreshold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteSleThreshold`: SleThreshold
	fmt.Fprintf(os.Stdout, "Response from `SitesSLEsAPI.UpdateSiteSleThreshold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**scope** | [**SiteSleThresholdScopeParameter**](.md) |  | 
**scopeId** | **string** | * site_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;site&#x60; * device_id if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;ap&#x60;, &#x60;scope&#x60;&#x3D;&#x3D;&#x60;switch&#x60; or &#x60;scope&#x60;&#x3D;&#x3D;&#x60;gateway&#x60; * mac if &#x60;scope&#x60;&#x3D;&#x3D;&#x60;client&#x60; | 
**metric** | **string** | values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteSleThresholdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sleThreshold** | [**SleThreshold**](SleThreshold.md) |  | 

### Return type

[**SleThreshold**](SleThreshold.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

