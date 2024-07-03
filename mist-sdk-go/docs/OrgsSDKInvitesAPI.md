# \OrgsSDKInvitesAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSdkInvite**](OrgsSDKInvitesAPI.md#ActivateSdkInvite) | **Post** /api/v1/mobile/verify/{secret} | activateSdkInvite
[**CreateSdkInvite**](OrgsSDKInvitesAPI.md#CreateSdkInvite) | **Post** /api/v1/orgs/{org_id}/sdkinvites | createSdkInvite
[**GetSdkInvite**](OrgsSDKInvitesAPI.md#GetSdkInvite) | **Get** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id} | getSdkInvite
[**GetSdkInviteQrCode**](OrgsSDKInvitesAPI.md#GetSdkInviteQrCode) | **Get** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id}/qrcode | getSdkInviteQrCode
[**ListSdkInvites**](OrgsSDKInvitesAPI.md#ListSdkInvites) | **Get** /api/v1/orgs/{org_id}/sdkinvites | listSdkInvites
[**RevokeSdkInvite**](OrgsSDKInvitesAPI.md#RevokeSdkInvite) | **Delete** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id} | revokeSdkInvite
[**SendSdkInviteEmail**](OrgsSDKInvitesAPI.md#SendSdkInviteEmail) | **Post** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id}/email | sendSdkInviteEmail
[**SendSdkInviteSms**](OrgsSDKInvitesAPI.md#SendSdkInviteSms) | **Post** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id}/sms | sendSdkInviteSms
[**UpdateSdkInvite**](OrgsSDKInvitesAPI.md#UpdateSdkInvite) | **Put** /api/v1/orgs/{org_id}/sdkinvites/{sdkinvite_id} | updateSdkInvite



## ActivateSdkInvite

> ResponseMobileVerifySecret ActivateSdkInvite(ctx, secret).DeviceIdString(deviceIdString).Execute()

activateSdkInvite



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
	secret := "secret_example" // string | 
	deviceIdString := *openapiclient.NewDeviceIdString("DeviceId_example") // DeviceIdString |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKInvitesAPI.ActivateSdkInvite(context.Background(), secret).DeviceIdString(deviceIdString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.ActivateSdkInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateSdkInvite`: ResponseMobileVerifySecret
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKInvitesAPI.ActivateSdkInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secret** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSdkInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceIdString** | [**DeviceIdString**](DeviceIdString.md) |  | 

### Return type

[**ResponseMobileVerifySecret**](ResponseMobileVerifySecret.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSdkInvite

> Sdkinvite CreateSdkInvite(ctx, orgId).Sdkinvite(sdkinvite).Execute()

createSdkInvite



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
	sdkinvite := *openapiclient.NewSdkinvite("Name_example") // Sdkinvite | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKInvitesAPI.CreateSdkInvite(context.Background(), orgId).Sdkinvite(sdkinvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.CreateSdkInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSdkInvite`: Sdkinvite
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKInvitesAPI.CreateSdkInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSdkInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sdkinvite** | [**Sdkinvite**](Sdkinvite.md) | Request Body | 

### Return type

[**Sdkinvite**](Sdkinvite.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdkInvite

> Sdkinvite GetSdkInvite(ctx, orgId, sdkinviteId).Execute()

getSdkInvite



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
	sdkinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKInvitesAPI.GetSdkInvite(context.Background(), orgId, sdkinviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.GetSdkInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSdkInvite`: Sdkinvite
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKInvitesAPI.GetSdkInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdkinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdkInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Sdkinvite**](Sdkinvite.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSdkInviteQrCode

> *os.File GetSdkInviteQrCode(ctx, orgId, sdkinviteId).Execute()

getSdkInviteQrCode



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
	sdkinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKInvitesAPI.GetSdkInviteQrCode(context.Background(), orgId, sdkinviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.GetSdkInviteQrCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSdkInviteQrCode`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKInvitesAPI.GetSdkInviteQrCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdkinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSdkInviteQrCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSdkInvites

> []Sdkinvite ListSdkInvites(ctx, orgId).Execute()

listSdkInvites



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKInvitesAPI.ListSdkInvites(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.ListSdkInvites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSdkInvites`: []Sdkinvite
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKInvitesAPI.ListSdkInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSdkInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Sdkinvite**](Sdkinvite.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeSdkInvite

> RevokeSdkInvite(ctx, orgId, sdkinviteId).Execute()

revokeSdkInvite



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
	sdkinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSDKInvitesAPI.RevokeSdkInvite(context.Background(), orgId, sdkinviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.RevokeSdkInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdkinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeSdkInviteRequest struct via the builder pattern


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


## SendSdkInviteEmail

> SendSdkInviteEmail(ctx, orgId, sdkinviteId).EmailString(emailString).Execute()

sendSdkInviteEmail



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
	sdkinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	emailString := *openapiclient.NewEmailString("Email_example") // EmailString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSDKInvitesAPI.SendSdkInviteEmail(context.Background(), orgId, sdkinviteId).EmailString(emailString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.SendSdkInviteEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdkinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSdkInviteEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailString** | [**EmailString**](EmailString.md) | Request Body | 

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


## SendSdkInviteSms

> SendSdkInviteSms(ctx, orgId, sdkinviteId).SdkInviteSms(sdkInviteSms).Execute()

sendSdkInviteSms



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
	sdkinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sdkInviteSms := *openapiclient.NewSdkInviteSms("Number_example") // SdkInviteSms | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsSDKInvitesAPI.SendSdkInviteSms(context.Background(), orgId, sdkinviteId).SdkInviteSms(sdkInviteSms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.SendSdkInviteSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdkinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSdkInviteSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sdkInviteSms** | [**SdkInviteSms**](SdkInviteSms.md) | Request Body | 

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


## UpdateSdkInvite

> Sdkinvite UpdateSdkInvite(ctx, orgId, sdkinviteId).Sdkinvite(sdkinvite).Execute()

updateSdkInvite



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
	sdkinviteId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	sdkinvite := *openapiclient.NewSdkinvite("Name_example") // Sdkinvite | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsSDKInvitesAPI.UpdateSdkInvite(context.Background(), orgId, sdkinviteId).Sdkinvite(sdkinvite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsSDKInvitesAPI.UpdateSdkInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSdkInvite`: Sdkinvite
	fmt.Fprintf(os.Stdout, "Response from `OrgsSDKInvitesAPI.UpdateSdkInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**sdkinviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSdkInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sdkinvite** | [**Sdkinvite**](Sdkinvite.md) | Request Body | 

### Return type

[**Sdkinvite**](Sdkinvite.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

