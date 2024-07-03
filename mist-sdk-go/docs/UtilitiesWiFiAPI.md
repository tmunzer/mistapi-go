# \UtilitiesWiFiAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeauthSiteWirelessClientsConnectedToARogue**](UtilitiesWiFiAPI.md#DeauthSiteWirelessClientsConnectedToARogue) | **Post** /api/v1/sites/{site_id}/rogues/{rogue_bssid}/deauth_clients | deauthSiteWirelessClientsConnectedToARogue
[**DisconnectSiteMultipleClients**](UtilitiesWiFiAPI.md#DisconnectSiteMultipleClients) | **Post** /api/v1/sites/{site_id}/clients/disconnect | disconnectSiteMultipleClients
[**DisconnectSiteWirelessClient**](UtilitiesWiFiAPI.md#DisconnectSiteWirelessClient) | **Post** /api/v1/sites/{site_id}/clients/{client_mac}/disconnect | disconnectSiteWirelessClient
[**OptimizeSiteRrm**](UtilitiesWiFiAPI.md#OptimizeSiteRrm) | **Post** /api/v1/sites/{site_id}/rrm/optimize | optimizeSiteRrm
[**ReauthOrgDot1xWirelessClient**](UtilitiesWiFiAPI.md#ReauthOrgDot1xWirelessClient) | **Post** /api/v1/orgs/{org_id}/clients/{client_mac}/coa | reauthOrgDot1xWirelessClient
[**ReauthSiteDot1xWirelessClient**](UtilitiesWiFiAPI.md#ReauthSiteDot1xWirelessClient) | **Post** /api/v1/sites/{site_id}/clients/{client_mac}/coa | reauthSiteDot1xWirelessClient
[**ReprovisionSiteAllAps**](UtilitiesWiFiAPI.md#ReprovisionSiteAllAps) | **Post** /api/v1/sites/{site_id}/devices/reprovision | reprovisionSiteAllAps
[**ResetSiteAllApsToUseRrm**](UtilitiesWiFiAPI.md#ResetSiteAllApsToUseRrm) | **Post** /api/v1/sites/{site_id}/devices/reset_radio_config | resetSiteAllApsToUseRrm
[**TestSiteWlanTelstraSetup**](UtilitiesWiFiAPI.md#TestSiteWlanTelstraSetup) | **Post** /api/v1/utils/test_telstra | testSiteWlanTelstraSetup
[**TestSiteWlanTwilioSetup**](UtilitiesWiFiAPI.md#TestSiteWlanTwilioSetup) | **Post** /api/v1/utils/test_twilio | testSiteWlanTwilioSetup
[**UnauthorizeSiteMultipleClients**](UtilitiesWiFiAPI.md#UnauthorizeSiteMultipleClients) | **Post** /api/v1/sites/{site_id}/clients/unauthorize | unauthorizeSiteMultipleClients
[**UnauthorizeSiteWirelessClient**](UtilitiesWiFiAPI.md#UnauthorizeSiteWirelessClient) | **Post** /api/v1/sites/{site_id}/clients/{client_mac}/unauthorize | unauthorizeSiteWirelessClient
[**ZeroizeSiteFipsAllAps**](UtilitiesWiFiAPI.md#ZeroizeSiteFipsAllAps) | **Post** /api/v1/sites/{site_id}/devices/zeroize | zeroizeSiteFipsAllAps



## DeauthSiteWirelessClientsConnectedToARogue

> DeauthSiteWirelessClientsConnectedToARogue(ctx, siteId, rogueBssid).Execute()

deauthSiteWirelessClientsConnectedToARogue



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
	rogueBssid := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.DeauthSiteWirelessClientsConnectedToARogue(context.Background(), siteId, rogueBssid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.DeauthSiteWirelessClientsConnectedToARogue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**rogueBssid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeauthSiteWirelessClientsConnectedToARogueRequest struct via the builder pattern


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


## DisconnectSiteMultipleClients

> DisconnectSiteMultipleClients(ctx, siteId).MacAddresses(macAddresses).Execute()

disconnectSiteMultipleClients



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
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.DisconnectSiteMultipleClients(context.Background(), siteId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.DisconnectSiteMultipleClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectSiteMultipleClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

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


## DisconnectSiteWirelessClient

> DisconnectSiteWirelessClient(ctx, siteId, clientMac).Execute()

disconnectSiteWirelessClient



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
	clientMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.DisconnectSiteWirelessClient(context.Background(), siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.DisconnectSiteWirelessClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectSiteWirelessClientRequest struct via the builder pattern


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


## OptimizeSiteRrm

> OptimizeSiteRrm(ctx, siteId).UtilsRrmOptimize(utilsRrmOptimize).Execute()

optimizeSiteRrm



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
	utilsRrmOptimize := *openapiclient.NewUtilsRrmOptimize([]string{"Bands_example"}) // UtilsRrmOptimize | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.OptimizeSiteRrm(context.Background(), siteId).UtilsRrmOptimize(utilsRrmOptimize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.OptimizeSiteRrm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOptimizeSiteRrmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **utilsRrmOptimize** | [**UtilsRrmOptimize**](UtilsRrmOptimize.md) | Request Body | 

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


## ReauthOrgDot1xWirelessClient

> ReauthOrgDot1xWirelessClient(ctx, orgId, clientMac).Execute()

reauthOrgDot1xWirelessClient



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
	clientMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.ReauthOrgDot1xWirelessClient(context.Background(), orgId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.ReauthOrgDot1xWirelessClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReauthOrgDot1xWirelessClientRequest struct via the builder pattern


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


## ReauthSiteDot1xWirelessClient

> ReauthSiteDot1xWirelessClient(ctx, siteId, clientMac).Execute()

reauthSiteDot1xWirelessClient



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
	clientMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.ReauthSiteDot1xWirelessClient(context.Background(), siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.ReauthSiteDot1xWirelessClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReauthSiteDot1xWirelessClientRequest struct via the builder pattern


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


## ReprovisionSiteAllAps

> ReprovisionSiteAllAps(ctx, siteId).Execute()

reprovisionSiteAllAps



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.ReprovisionSiteAllAps(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.ReprovisionSiteAllAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReprovisionSiteAllApsRequest struct via the builder pattern


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


## ResetSiteAllApsToUseRrm

> ResetSiteAllApsToUseRrm(ctx, siteId).UtilsResetRadioConfig(utilsResetRadioConfig).Execute()

resetSiteAllApsToUseRrm



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
	utilsResetRadioConfig := *openapiclient.NewUtilsResetRadioConfig([]string{"Bands_example"}) // UtilsResetRadioConfig | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.ResetSiteAllApsToUseRrm(context.Background(), siteId).UtilsResetRadioConfig(utilsResetRadioConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.ResetSiteAllApsToUseRrm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetSiteAllApsToUseRrmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **utilsResetRadioConfig** | [**UtilsResetRadioConfig**](UtilsResetRadioConfig.md) | Request Body | 

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


## TestSiteWlanTelstraSetup

> TestSiteWlanTelstraSetup(ctx).TestTelstra(testTelstra).Execute()

testSiteWlanTelstraSetup



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
	testTelstra := *openapiclient.NewTestTelstra("123456", "abcdef", "+911122334455") // TestTelstra |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.TestSiteWlanTelstraSetup(context.Background()).TestTelstra(testTelstra).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.TestSiteWlanTelstraSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSiteWlanTelstraSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testTelstra** | [**TestTelstra**](TestTelstra.md) |  | 

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


## TestSiteWlanTwilioSetup

> TestSiteWlanTwilioSetup(ctx).TestTwilio(testTwilio).Execute()

testSiteWlanTwilioSetup



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
	testTwilio := *openapiclient.NewTestTwilio("+185051234567", "+19999999999", "2135be04736a1a0a314bce432d61721a", "AC5f4366878d193fb4865ab151739999eb") // TestTwilio | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.TestSiteWlanTwilioSetup(context.Background()).TestTwilio(testTwilio).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.TestSiteWlanTwilioSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestSiteWlanTwilioSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testTwilio** | [**TestTwilio**](TestTwilio.md) | Request Body | 

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


## UnauthorizeSiteMultipleClients

> UnauthorizeSiteMultipleClients(ctx, siteId).MacAddresses(macAddresses).Execute()

unauthorizeSiteMultipleClients



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
	macAddresses := *openapiclient.NewMacAddresses([]string{"Macs_example"}) // MacAddresses | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.UnauthorizeSiteMultipleClients(context.Background(), siteId).MacAddresses(macAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.UnauthorizeSiteMultipleClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnauthorizeSiteMultipleClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **macAddresses** | [**MacAddresses**](MacAddresses.md) | Request Body | 

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


## UnauthorizeSiteWirelessClient

> UnauthorizeSiteWirelessClient(ctx, siteId, clientMac).Execute()

unauthorizeSiteWirelessClient



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
	clientMac := "0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.UnauthorizeSiteWirelessClient(context.Background(), siteId, clientMac).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.UnauthorizeSiteWirelessClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**clientMac** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnauthorizeSiteWirelessClientRequest struct via the builder pattern


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


## ZeroizeSiteFipsAllAps

> ZeroizeSiteFipsAllAps(ctx, siteId).UtilsZeroiseFips(utilsZeroiseFips).Execute()

zeroizeSiteFipsAllAps



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
	utilsZeroiseFips := *openapiclient.NewUtilsZeroiseFips("Password_example") // UtilsZeroiseFips | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UtilitiesWiFiAPI.ZeroizeSiteFipsAllAps(context.Background(), siteId).UtilsZeroiseFips(utilsZeroiseFips).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UtilitiesWiFiAPI.ZeroizeSiteFipsAllAps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiZeroizeSiteFipsAllApsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **utilsZeroiseFips** | [**UtilsZeroiseFips**](UtilsZeroiseFips.md) | Request Body | 

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

