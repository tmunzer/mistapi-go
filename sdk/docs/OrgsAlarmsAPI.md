# \OrgsAlarmsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckOrgAlarm**](OrgsAlarmsAPI.md#AckOrgAlarm) | **Post** /api/v1/orgs/{org_id}/alarms/{alarm_id}/ack | ackOrgAlarm
[**AckOrgAllAlarms**](OrgsAlarmsAPI.md#AckOrgAllAlarms) | **Post** /api/v1/orgs/{org_id}/alarms/ack_all | ackOrgAllAlarms
[**AckOrgMultipleAlarms**](OrgsAlarmsAPI.md#AckOrgMultipleAlarms) | **Post** /api/v1/orgs/{org_id}/alarms/ack | ackOrgMultipleAlarms
[**CountOrgAlarms**](OrgsAlarmsAPI.md#CountOrgAlarms) | **Get** /api/v1/orgs/{org_id}/alarms/count | countOrgAlarms
[**SearchOrgAlarms**](OrgsAlarmsAPI.md#SearchOrgAlarms) | **Get** /api/v1/orgs/{org_id}/alarms/search | searchOrgAlarms
[**UnackOrgAllArlarms**](OrgsAlarmsAPI.md#UnackOrgAllArlarms) | **Post** /api/v1/orgs/{org_id}/alarms/unack_all | unackOrgAllArlarms
[**UnackOrgMultipleAlarms**](OrgsAlarmsAPI.md#UnackOrgMultipleAlarms) | **Post** /api/v1/orgs/{org_id}/alarms/unack | unackOrgMultipleAlarms



## AckOrgAlarm

> AckOrgAlarm(ctx, orgId, alarmId).NoteString(noteString).Execute()

ackOrgAlarm



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
	alarmId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	noteString := *openapiclient.NewNoteString() // NoteString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmsAPI.AckOrgAlarm(context.Background(), orgId, alarmId).NoteString(noteString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmsAPI.AckOrgAlarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**alarmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAckOrgAlarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **noteString** | [**NoteString**](NoteString.md) | Request Body | 

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


## AckOrgAllAlarms

> AckOrgAllAlarms(ctx, orgId).NoteString(noteString).Execute()

ackOrgAllAlarms



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
	noteString := *openapiclient.NewNoteString() // NoteString |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmsAPI.AckOrgAllAlarms(context.Background(), orgId).NoteString(noteString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmsAPI.AckOrgAllAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAckOrgAllAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteString** | [**NoteString**](NoteString.md) |  | 

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


## AckOrgMultipleAlarms

> AckOrgMultipleAlarms(ctx, orgId).Alarms(alarms).Execute()

ackOrgMultipleAlarms



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
	alarms := *openapiclient.NewAlarms([]string{"AlarmIds_example"}) // Alarms | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmsAPI.AckOrgMultipleAlarms(context.Background(), orgId).Alarms(alarms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmsAPI.AckOrgMultipleAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAckOrgMultipleAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alarms** | [**Alarms**](Alarms.md) | Request Body | 

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


## CountOrgAlarms

> RepsonseCount CountOrgAlarms(ctx, orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countOrgAlarms



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
	distinct := "distinct_example" // string |  (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAlarmsAPI.CountOrgAlarms(context.Background(), orgId).Distinct(distinct).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmsAPI.CountOrgAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgAlarms`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsAlarmsAPI.CountOrgAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | **string** |  | 
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**RepsonseCount**](RepsonseCount.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOrgAlarms

> AlarmSearchResult SearchOrgAlarms(ctx, orgId).SiteId(siteId).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchOrgAlarms



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
	siteId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | site ID (optional)
	type_ := "type__example" // string | alarm type (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAlarmsAPI.SearchOrgAlarms(context.Background(), orgId).SiteId(siteId).Type_(type_).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmsAPI.SearchOrgAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgAlarms`: AlarmSearchResult
	fmt.Fprintf(os.Stdout, "Response from `OrgsAlarmsAPI.SearchOrgAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteId** | **string** | site ID | 
 **type_** | **string** | alarm type | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**AlarmSearchResult**](AlarmSearchResult.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnackOrgAllArlarms

> UnackOrgAllArlarms(ctx, orgId).NoteString(noteString).Execute()

unackOrgAllArlarms



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
	noteString := *openapiclient.NewNoteString() // NoteString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmsAPI.UnackOrgAllArlarms(context.Background(), orgId).NoteString(noteString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmsAPI.UnackOrgAllArlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnackOrgAllArlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteString** | [**NoteString**](NoteString.md) | Request Body | 

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


## UnackOrgMultipleAlarms

> UnackOrgMultipleAlarms(ctx, orgId).Alarms(alarms).Execute()

unackOrgMultipleAlarms



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
	alarms := *openapiclient.NewAlarms([]string{"AlarmIds_example"}) // Alarms | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAlarmsAPI.UnackOrgMultipleAlarms(context.Background(), orgId).Alarms(alarms).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAlarmsAPI.UnackOrgMultipleAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnackOrgMultipleAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alarms** | [**Alarms**](Alarms.md) | Request Body | 

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

