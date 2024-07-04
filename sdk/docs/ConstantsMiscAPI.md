# \ConstantsMiscAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGatewayDefaultConfig**](ConstantsMiscAPI.md#GetGatewayDefaultConfig) | **Get** /api/v1/const/default_gateway_config | getGatewayDefaultConfig
[**GetLicenseTypes**](ConstantsMiscAPI.md#GetLicenseTypes) | **Get** /api/v1/const/license_types | getLicenseTypes
[**ListApChannels**](ConstantsMiscAPI.md#ListApChannels) | **Get** /api/v1/const/ap_channels | listApChannels
[**ListApplications**](ConstantsMiscAPI.md#ListApplications) | **Get** /api/v1/const/applications | listApplications
[**ListCountryCodes**](ConstantsMiscAPI.md#ListCountryCodes) | **Get** /api/v1/const/countries | listCountryCodes
[**ListGatewayApplications**](ConstantsMiscAPI.md#ListGatewayApplications) | **Get** /api/v1/const/gateway_applications | listGatewayApplications
[**ListInsightMetrics**](ConstantsMiscAPI.md#ListInsightMetrics) | **Get** /api/v1/const/insight_metrics | listInsightMetrics
[**ListSiteLanguages**](ConstantsMiscAPI.md#ListSiteLanguages) | **Get** /api/v1/const/languages | listSiteLanguages
[**ListTrafficTypes**](ConstantsMiscAPI.md#ListTrafficTypes) | **Get** /api/v1/const/traffic_types | listTrafficTypes



## GetGatewayDefaultConfig

> map[string]interface{} GetGatewayDefaultConfig(ctx).Model(model).Ha(ha).Execute()

getGatewayDefaultConfig



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
	model := "model_example" // string | model the default gateway config is intended (as the default LAN/WAN port can differ)
	ha := "ha_example" // string | whether the config is intended for HA (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.GetGatewayDefaultConfig(context.Background()).Model(model).Ha(ha).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.GetGatewayDefaultConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGatewayDefaultConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.GetGatewayDefaultConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayDefaultConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | model the default gateway config is intended (as the default LAN/WAN port can differ) | 
 **ha** | **string** | whether the config is intended for HA | 

### Return type

**map[string]interface{}**

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseTypes

> []ConstLicenseType GetLicenseTypes(ctx).Execute()

getLicenseTypes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.GetLicenseTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.GetLicenseTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseTypes`: []ConstLicenseType
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.GetLicenseTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseTypesRequest struct via the builder pattern


### Return type

[**[]ConstLicenseType**](ConstLicenseType.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApChannels

> ConstApChannel ListApChannels(ctx).CountryCode(countryCode).Execute()

listApChannels



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
	countryCode := "US" // string | country code, in two-character (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.ListApChannels(context.Background()).CountryCode(countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.ListApChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApChannels`: ConstApChannel
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.ListApChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** | country code, in two-character | 

### Return type

[**ConstApChannel**](ConstApChannel.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> []ConstApplicationDefinition ListApplications(ctx).Execute()

listApplications



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.ListApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: []ConstApplicationDefinition
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.ListApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


### Return type

[**[]ConstApplicationDefinition**](ConstApplicationDefinition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCountryCodes

> []ConstCountry ListCountryCodes(ctx).Execute()

listCountryCodes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.ListCountryCodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.ListCountryCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCountryCodes`: []ConstCountry
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.ListCountryCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCountryCodesRequest struct via the builder pattern


### Return type

[**[]ConstCountry**](ConstCountry.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGatewayApplications

> []ConstGatewayApplicationsDefinition ListGatewayApplications(ctx).Execute()

listGatewayApplications



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.ListGatewayApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.ListGatewayApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGatewayApplications`: []ConstGatewayApplicationsDefinition
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.ListGatewayApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGatewayApplicationsRequest struct via the builder pattern


### Return type

[**[]ConstGatewayApplicationsDefinition**](ConstGatewayApplicationsDefinition.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightMetrics

> map[string]ConstInsightMetricsProperty ListInsightMetrics(ctx).Execute()

listInsightMetrics



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.ListInsightMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.ListInsightMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInsightMetrics`: map[string]ConstInsightMetricsProperty
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.ListInsightMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInsightMetricsRequest struct via the builder pattern


### Return type

[**map[string]ConstInsightMetricsProperty**](ConstInsightMetricsProperty.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSiteLanguages

> []ConstLanguage ListSiteLanguages(ctx).Execute()

listSiteLanguages



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.ListSiteLanguages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.ListSiteLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteLanguages`: []ConstLanguage
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.ListSiteLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteLanguagesRequest struct via the builder pattern


### Return type

[**[]ConstLanguage**](ConstLanguage.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrafficTypes

> []ConstTrafficType ListTrafficTypes(ctx).Execute()

listTrafficTypes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConstantsMiscAPI.ListTrafficTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConstantsMiscAPI.ListTrafficTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTrafficTypes`: []ConstTrafficType
	fmt.Fprintf(os.Stdout, "Response from `ConstantsMiscAPI.ListTrafficTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTrafficTypesRequest struct via the builder pattern


### Return type

[**[]ConstTrafficType**](ConstTrafficType.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

