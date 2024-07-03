# \SitesAlarmsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AckSiteAlarm**](SitesAlarmsAPI.md#AckSiteAlarm) | **Post** /api/v1/sites/{site_id}/alarms/{alarm_id}/ack | ackSiteAlarm
[**AckSiteAllAlarms**](SitesAlarmsAPI.md#AckSiteAllAlarms) | **Post** /api/v1/sites/{site_id}/alarms/ack_all | ackSiteAllAlarms
[**AckSiteMultipleAlarms**](SitesAlarmsAPI.md#AckSiteMultipleAlarms) | **Post** /api/v1/sites/{site_id}/alarms/ack | AckSiteMultipleAlarms
[**CountSiteAlarms**](SitesAlarmsAPI.md#CountSiteAlarms) | **Get** /api/v1/sites/{site_id}/alarms/count | countSiteAlarms
[**SearchSiteAlarms**](SitesAlarmsAPI.md#SearchSiteAlarms) | **Get** /api/v1/sites/{site_id}/alarms/search | searchSiteAlarms
[**SubscribeSiteAlarms**](SitesAlarmsAPI.md#SubscribeSiteAlarms) | **Post** /api/v1/sites/{site_id}/subscriptions | SubscribeSiteAlarms
[**UnackSiteAlarm**](SitesAlarmsAPI.md#UnackSiteAlarm) | **Post** /api/v1/sites/{site_id}/alarms/{alarm_id}/unack | unackSiteAlarm
[**UnackSiteAllArlarms**](SitesAlarmsAPI.md#UnackSiteAllArlarms) | **Post** /api/v1/sites/{site_id}/alarms/unack_all | unackSiteAllArlarms
[**UnackSiteMultipleAlarms**](SitesAlarmsAPI.md#UnackSiteMultipleAlarms) | **Post** /api/v1/sites/{site_id}/alarms/unack | unackSiteMultipleAlarms
[**UnsubscribeSiteAlarms**](SitesAlarmsAPI.md#UnsubscribeSiteAlarms) | **Delete** /api/v1/sites/{site_id}/subscriptions | UnsubscribeSiteAlarms



## AckSiteAlarm

> AckSiteAlarm(ctx, siteId, alarmId).NoteString(noteString).Execute()

ackSiteAlarm



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
	alarmId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	noteString := *openapiclient.NewNoteString() // NoteString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAlarmsAPI.AckSiteAlarm(context.Background(), siteId, alarmId).NoteString(noteString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.AckSiteAlarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**alarmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAckSiteAlarmRequest struct via the builder pattern


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


## AckSiteAllAlarms

> AckSiteAllAlarms(ctx, siteId).NoteString(noteString).Execute()

ackSiteAllAlarms



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
	noteString := *openapiclient.NewNoteString() // NoteString |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAlarmsAPI.AckSiteAllAlarms(context.Background(), siteId).NoteString(noteString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.AckSiteAllAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAckSiteAllAlarmsRequest struct via the builder pattern


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


## AckSiteMultipleAlarms

> AckSiteMultipleAlarms(ctx, siteId).AlarmAck(alarmAck).Execute()

AckSiteMultipleAlarms



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
	alarmAck := *openapiclient.NewAlarmAck([]string{"AlarmIds_example"}) // AlarmAck | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAlarmsAPI.AckSiteMultipleAlarms(context.Background(), siteId).AlarmAck(alarmAck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.AckSiteMultipleAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAckSiteMultipleAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alarmAck** | [**AlarmAck**](AlarmAck.md) | Request Body | 

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


## CountSiteAlarms

> RepsonseCount CountSiteAlarms(ctx, siteId).Distinct(distinct).AckAdminName(ackAdminName).Acked(acked).Type_(type_).Severity(severity).Group(group).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()

countSiteAlarms



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
	distinct := openapiclient.alarm_count_disctinct("") // AlarmCountDisctinct | Group by and count the alarms by some distinct field (optional) (default to "type")
	ackAdminName := "ackAdminName_example" // string | Name of the admins who have acked the alarms; accepts multiple values separated by comma (optional)
	acked := true // bool |  (optional)
	type_ := "type__example" // string | Key-name of the alarms; accepts multiple values separated by comma (optional)
	severity := "severity_example" // string | Alarm severity; accepts multiple values separated by comma (optional)
	group := "group_example" // string | Alarm group name; accepts multiple values separated by comma (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAlarmsAPI.CountSiteAlarms(context.Background(), siteId).Distinct(distinct).AckAdminName(ackAdminName).Acked(acked).Type_(type_).Severity(severity).Group(group).Page(page).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.CountSiteAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteAlarms`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesAlarmsAPI.CountSiteAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**AlarmCountDisctinct**](AlarmCountDisctinct.md) | Group by and count the alarms by some distinct field | [default to &quot;type&quot;]
 **ackAdminName** | **string** | Name of the admins who have acked the alarms; accepts multiple values separated by comma | 
 **acked** | **bool** |  | 
 **type_** | **string** | Key-name of the alarms; accepts multiple values separated by comma | 
 **severity** | **string** | Alarm severity; accepts multiple values separated by comma | 
 **group** | **string** | Alarm group name; accepts multiple values separated by comma | 
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


## SearchSiteAlarms

> AlarmSearchResult SearchSiteAlarms(ctx, siteId).Type_(type_).AckAdminName(ackAdminName).Acked(acked).Severity(severity).Group(group).Limit(limit).Start(start).End(end).Duration(duration).Execute()

searchSiteAlarms



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
	type_ := "type__example" // string | Key-name of the alarms; accepts multiple values separated by comma (optional)
	ackAdminName := "ackAdminName_example" // string | Name of the admins who have acked the alarms; accepts multiple values separated by comma (optional)
	acked := true // bool |  (optional)
	severity := "severity_example" // string | Alarm severity; accepts multiple values separated by comma (optional)
	group := "group_example" // string | Alarm group name; accepts multiple values separated by comma (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAlarmsAPI.SearchSiteAlarms(context.Background(), siteId).Type_(type_).AckAdminName(ackAdminName).Acked(acked).Severity(severity).Group(group).Limit(limit).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.SearchSiteAlarms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteAlarms`: AlarmSearchResult
	fmt.Fprintf(os.Stdout, "Response from `SitesAlarmsAPI.SearchSiteAlarms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Key-name of the alarms; accepts multiple values separated by comma | 
 **ackAdminName** | **string** | Name of the admins who have acked the alarms; accepts multiple values separated by comma | 
 **acked** | **bool** |  | 
 **severity** | **string** | Alarm severity; accepts multiple values separated by comma | 
 **group** | **string** | Alarm group name; accepts multiple values separated by comma | 
 **limit** | **int32** |  | [default to 100]
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

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


## SubscribeSiteAlarms

> SubscribeSiteAlarms(ctx, siteId).Execute()

SubscribeSiteAlarms



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
	r, err := apiClient.SitesAlarmsAPI.SubscribeSiteAlarms(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.SubscribeSiteAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscribeSiteAlarmsRequest struct via the builder pattern


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


## UnackSiteAlarm

> UnackSiteAlarm(ctx, siteId, alarmId).NoteString(noteString).Execute()

unackSiteAlarm



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
	alarmId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	noteString := *openapiclient.NewNoteString() // NoteString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAlarmsAPI.UnackSiteAlarm(context.Background(), siteId, alarmId).NoteString(noteString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.UnackSiteAlarm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**alarmId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnackSiteAlarmRequest struct via the builder pattern


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


## UnackSiteAllArlarms

> UnackSiteAllArlarms(ctx, siteId).NoteString(noteString).Execute()

unackSiteAllArlarms



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
	noteString := *openapiclient.NewNoteString() // NoteString | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAlarmsAPI.UnackSiteAllArlarms(context.Background(), siteId).NoteString(noteString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.UnackSiteAllArlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnackSiteAllArlarmsRequest struct via the builder pattern


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


## UnackSiteMultipleAlarms

> UnackSiteMultipleAlarms(ctx, siteId).AlarmAck(alarmAck).Execute()

unackSiteMultipleAlarms



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
	alarmAck := *openapiclient.NewAlarmAck([]string{"AlarmIds_example"}) // AlarmAck | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesAlarmsAPI.UnackSiteMultipleAlarms(context.Background(), siteId).AlarmAck(alarmAck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.UnackSiteMultipleAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnackSiteMultipleAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alarmAck** | [**AlarmAck**](AlarmAck.md) | Request Body | 

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


## UnsubscribeSiteAlarms

> UnsubscribeSiteAlarms(ctx, siteId).Execute()

UnsubscribeSiteAlarms



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
	r, err := apiClient.SitesAlarmsAPI.UnsubscribeSiteAlarms(context.Background(), siteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAlarmsAPI.UnsubscribeSiteAlarms``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUnsubscribeSiteAlarmsRequest struct via the builder pattern


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

