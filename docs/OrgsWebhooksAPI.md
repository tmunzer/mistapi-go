# \OrgsWebhooksAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountOrgWebhooksDeliveries**](OrgsWebhooksAPI.md#CountOrgWebhooksDeliveries) | **Get** /api/v1/orgs/{org_id}/webhooks/{webhook_id}/events/count | countOrgWebhooksDeliveries
[**CreateOrgWebhook**](OrgsWebhooksAPI.md#CreateOrgWebhook) | **Post** /api/v1/orgs/{org_id}/webhooks | createOrgWebhook
[**DeleteOrgWebhook**](OrgsWebhooksAPI.md#DeleteOrgWebhook) | **Delete** /api/v1/orgs/{org_id}/webhooks/{webhook_id} | deleteOrgWebhook
[**GetOrgWebhook**](OrgsWebhooksAPI.md#GetOrgWebhook) | **Get** /api/v1/orgs/{org_id}/webhooks/{webhook_id} | getOrgWebhook
[**ListOrgWebhooks**](OrgsWebhooksAPI.md#ListOrgWebhooks) | **Get** /api/v1/orgs/{org_id}/webhooks | listOrgWebhooks
[**PingOrgWebhook**](OrgsWebhooksAPI.md#PingOrgWebhook) | **Post** /api/v1/orgs/{org_id}/webhooks/{webhook_id}/ping | pingOrgWebhook
[**SearchOrgWebhooksDeliveries**](OrgsWebhooksAPI.md#SearchOrgWebhooksDeliveries) | **Get** /api/v1/orgs/{org_id}/webhooks/{webhook_id}/events/search | searchOrgWebhooksDeliveries
[**UpdateOrgWebhook**](OrgsWebhooksAPI.md#UpdateOrgWebhook) | **Put** /api/v1/orgs/{org_id}/webhooks/{webhook_id} | updateOrgWebhook



## CountOrgWebhooksDeliveries

> RepsonseCount CountOrgWebhooksDeliveries(ctx, orgId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countOrgWebhooksDeliveries



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	error_ := "error__example" // string |  (optional)
	statusCode := int32(56) // int32 |  (optional)
	status := openapiclient.webhook_delivery_status("") // WebhookDeliveryStatus | webhook delivery status (optional)
	topic := openapiclient.webhook_delivery_topic("") // WebhookDeliveryTopic | webhook topic (optional)
	distinct := openapiclient.webhook_delivery_distinct("") // WebhookDeliveryDistinct |  (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWebhooksAPI.CountOrgWebhooksDeliveries(context.Background(), orgId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.CountOrgWebhooksDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgWebhooksDeliveries`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsWebhooksAPI.CountOrgWebhooksDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgWebhooksDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **error_** | **string** |  | 
 **statusCode** | **int32** |  | 
 **status** | [**WebhookDeliveryStatus**](WebhookDeliveryStatus.md) | webhook delivery status | 
 **topic** | [**WebhookDeliveryTopic**](WebhookDeliveryTopic.md) | webhook topic | 
 **distinct** | [**WebhookDeliveryDistinct**](WebhookDeliveryDistinct.md) |  | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

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


## CreateOrgWebhook

> Webhook CreateOrgWebhook(ctx, orgId).Webhook(webhook).Execute()

createOrgWebhook



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
	webhook := *openapiclient.NewWebhook() // Webhook | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWebhooksAPI.CreateOrgWebhook(context.Background(), orgId).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.CreateOrgWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `OrgsWebhooksAPI.CreateOrgWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhook** | [**Webhook**](Webhook.md) | Request Body | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgWebhook

> DeleteOrgWebhook(ctx, orgId, webhookId).Execute()

deleteOrgWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsWebhooksAPI.DeleteOrgWebhook(context.Background(), orgId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.DeleteOrgWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgWebhookRequest struct via the builder pattern


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


## GetOrgWebhook

> Webhook GetOrgWebhook(ctx, orgId, webhookId).Execute()

getOrgWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWebhooksAPI.GetOrgWebhook(context.Background(), orgId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.GetOrgWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `OrgsWebhooksAPI.GetOrgWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgWebhooks

> []Webhook ListOrgWebhooks(ctx, orgId).Page(page).Limit(limit).Execute()

listOrgWebhooks



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWebhooksAPI.ListOrgWebhooks(context.Background(), orgId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.ListOrgWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgWebhooks`: []Webhook
	fmt.Fprintf(os.Stdout, "Response from `OrgsWebhooksAPI.ListOrgWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 100]

### Return type

[**[]Webhook**](Webhook.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingOrgWebhook

> PingOrgWebhook(ctx, orgId, webhookId).Execute()

pingOrgWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsWebhooksAPI.PingOrgWebhook(context.Background(), orgId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.PingOrgWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingOrgWebhookRequest struct via the builder pattern


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


## SearchOrgWebhooksDeliveries

> SearchWebhookDelivery SearchOrgWebhooksDeliveries(ctx, orgId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchOrgWebhooksDeliveries



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	error_ := "error__example" // string |  (optional)
	statusCode := int32(56) // int32 |  (optional)
	status := openapiclient.webhook_delivery_status("") // WebhookDeliveryStatus | webhook delivery status (optional)
	topic := openapiclient.webhook_delivery_topic("") // WebhookDeliveryTopic | webhook topic (optional)
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWebhooksAPI.SearchOrgWebhooksDeliveries(context.Background(), orgId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.SearchOrgWebhooksDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOrgWebhooksDeliveries`: SearchWebhookDelivery
	fmt.Fprintf(os.Stdout, "Response from `OrgsWebhooksAPI.SearchOrgWebhooksDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchOrgWebhooksDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **error_** | **string** |  | 
 **statusCode** | **int32** |  | 
 **status** | [**WebhookDeliveryStatus**](WebhookDeliveryStatus.md) | webhook delivery status | 
 **topic** | [**WebhookDeliveryTopic**](WebhookDeliveryTopic.md) | webhook topic | 
 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]
 **limit** | **int32** |  | [default to 100]

### Return type

[**SearchWebhookDelivery**](SearchWebhookDelivery.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgWebhook

> Webhook UpdateOrgWebhook(ctx, orgId, webhookId).Webhook(webhook).Execute()

updateOrgWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	webhook := *openapiclient.NewWebhook() // Webhook | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsWebhooksAPI.UpdateOrgWebhook(context.Background(), orgId, webhookId).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsWebhooksAPI.UpdateOrgWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `OrgsWebhooksAPI.UpdateOrgWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webhook** | [**Webhook**](Webhook.md) | Request Body | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

