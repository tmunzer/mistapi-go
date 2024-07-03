# \SitesWebhooksAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountSiteWebhooksDeliveries**](SitesWebhooksAPI.md#CountSiteWebhooksDeliveries) | **Get** /api/v1/sites/{site_id}/webhooks/{webhook_id}/events/count | countSiteWebhooksDeliveries
[**CreateSiteWebhook**](SitesWebhooksAPI.md#CreateSiteWebhook) | **Post** /api/v1/sites/{site_id}/webhooks | createSiteWebhook
[**DeleteSiteWebhook**](SitesWebhooksAPI.md#DeleteSiteWebhook) | **Delete** /api/v1/sites/{site_id}/webhooks/{webhook_id} | deleteSiteWebhook
[**GetSiteWebhook**](SitesWebhooksAPI.md#GetSiteWebhook) | **Get** /api/v1/sites/{site_id}/webhooks/{webhook_id} | getSiteWebhook
[**ListSiteWebhooks**](SitesWebhooksAPI.md#ListSiteWebhooks) | **Get** /api/v1/sites/{site_id}/webhooks | listSiteWebhooks
[**PingSiteWebhook**](SitesWebhooksAPI.md#PingSiteWebhook) | **Post** /api/v1/sites/{site_id}/webhooks/{webhook_id}/ping | pingSiteWebhook
[**SearchSiteWebhooksDeliveries**](SitesWebhooksAPI.md#SearchSiteWebhooksDeliveries) | **Get** /api/v1/sites/{site_id}/webhooks/{webhook_id}/events/search | searchSiteWebhooksDeliveries
[**UpdateSiteWebhook**](SitesWebhooksAPI.md#UpdateSiteWebhook) | **Put** /api/v1/sites/{site_id}/webhooks/{webhook_id} | updateSiteWebhook



## CountSiteWebhooksDeliveries

> RepsonseCount CountSiteWebhooksDeliveries(ctx, siteId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Execute()

countSiteWebhooksDeliveries



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
	resp, r, err := apiClient.SitesWebhooksAPI.CountSiteWebhooksDeliveries(context.Background(), siteId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Distinct(distinct).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.CountSiteWebhooksDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountSiteWebhooksDeliveries`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `SitesWebhooksAPI.CountSiteWebhooksDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountSiteWebhooksDeliveriesRequest struct via the builder pattern


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


## CreateSiteWebhook

> Webhook CreateSiteWebhook(ctx, siteId).Webhook(webhook).Execute()

createSiteWebhook



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
	webhook := *openapiclient.NewWebhook() // Webhook | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWebhooksAPI.CreateSiteWebhook(context.Background(), siteId).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.CreateSiteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSiteWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `SitesWebhooksAPI.CreateSiteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSiteWebhookRequest struct via the builder pattern


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


## DeleteSiteWebhook

> DeleteSiteWebhook(ctx, siteId, webhookId).Execute()

deleteSiteWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesWebhooksAPI.DeleteSiteWebhook(context.Background(), siteId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.DeleteSiteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSiteWebhookRequest struct via the builder pattern


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


## GetSiteWebhook

> Webhook GetSiteWebhook(ctx, siteId, webhookId).Execute()

getSiteWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWebhooksAPI.GetSiteWebhook(context.Background(), siteId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.GetSiteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSiteWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `SitesWebhooksAPI.GetSiteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSiteWebhookRequest struct via the builder pattern


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


## ListSiteWebhooks

> []Webhook ListSiteWebhooks(ctx, siteId).Page(page).Limit(limit).Execute()

listSiteWebhooks



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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWebhooksAPI.ListSiteWebhooks(context.Background(), siteId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.ListSiteWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSiteWebhooks`: []Webhook
	fmt.Fprintf(os.Stdout, "Response from `SitesWebhooksAPI.ListSiteWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSiteWebhooksRequest struct via the builder pattern


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


## PingSiteWebhook

> PingSiteWebhook(ctx, siteId, webhookId).Execute()

pingSiteWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SitesWebhooksAPI.PingSiteWebhook(context.Background(), siteId, webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.PingSiteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingSiteWebhookRequest struct via the builder pattern


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


## SearchSiteWebhooksDeliveries

> SearchWebhookDelivery SearchSiteWebhooksDeliveries(ctx, siteId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Start(start).End(end).Duration(duration).Limit(limit).Execute()

searchSiteWebhooksDeliveries



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
	resp, r, err := apiClient.SitesWebhooksAPI.SearchSiteWebhooksDeliveries(context.Background(), siteId, webhookId).Error_(error_).StatusCode(statusCode).Status(status).Topic(topic).Start(start).End(end).Duration(duration).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.SearchSiteWebhooksDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSiteWebhooksDeliveries`: SearchWebhookDelivery
	fmt.Fprintf(os.Stdout, "Response from `SitesWebhooksAPI.SearchSiteWebhooksDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchSiteWebhooksDeliveriesRequest struct via the builder pattern


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


## UpdateSiteWebhook

> Webhook UpdateSiteWebhook(ctx, siteId, webhookId).Webhook(webhook).Execute()

updateSiteWebhook



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
	webhookId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	webhook := *openapiclient.NewWebhook() // Webhook | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesWebhooksAPI.UpdateSiteWebhook(context.Background(), siteId, webhookId).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesWebhooksAPI.UpdateSiteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSiteWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `SitesWebhooksAPI.UpdateSiteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**siteId** | **string** |  | 
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSiteWebhookRequest struct via the builder pattern


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

