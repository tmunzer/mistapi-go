# \OrgsTicketsAPI

All URIs are relative to *https://api.mist.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrgTicketComment**](OrgsTicketsAPI.md#AddOrgTicketComment) | **Post** /api/v1/orgs/{org_id}/tickets/{ticket_id}/comments | addOrgTicketComment
[**CountOrgTickets**](OrgsTicketsAPI.md#CountOrgTickets) | **Get** /api/v1/orgs/{org_id}/tickets/count | countOrgTickets
[**CreateOrgTicket**](OrgsTicketsAPI.md#CreateOrgTicket) | **Post** /api/v1/orgs/{org_id}/tickets | createOrgTicket
[**GetOrgTicket**](OrgsTicketsAPI.md#GetOrgTicket) | **Get** /api/v1/orgs/{org_id}/tickets/{ticket_id} | getOrgTicket
[**GetOrgTicketAttachment**](OrgsTicketsAPI.md#GetOrgTicketAttachment) | **Get** /api/v1/orgs/{org_id}/tickets/{ticket_id}/attachments/{attachment_id} | GetOrgTicketAttachment
[**ListOrgTickets**](OrgsTicketsAPI.md#ListOrgTickets) | **Get** /api/v1/orgs/{org_id}/tickets | listOrgTickets
[**UpdateOrgTicket**](OrgsTicketsAPI.md#UpdateOrgTicket) | **Put** /api/v1/orgs/{org_id}/tickets/{ticket_id} | updateOrgTicket
[**UploadrgTicketAttachment**](OrgsTicketsAPI.md#UploadrgTicketAttachment) | **Post** /api/v1/orgs/{org_id}/tickets/{ticket_id}/attachments | UploadrgTicketAttachment



## AddOrgTicketComment

> Ticket AddOrgTicketComment(ctx, orgId, ticketId).TicketComment(ticketComment).Execute()

addOrgTicketComment



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
	ticketId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ticketComment := *openapiclient.NewTicketComment("Author_example", "Comment_example", int32(123)) // TicketComment | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTicketsAPI.AddOrgTicketComment(context.Background(), orgId, ticketId).TicketComment(ticketComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.AddOrgTicketComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOrgTicketComment`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `OrgsTicketsAPI.AddOrgTicketComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOrgTicketCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ticketComment** | [**TicketComment**](TicketComment.md) | Request Body | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountOrgTickets

> RepsonseCount CountOrgTickets(ctx, orgId).Distinct(distinct).Execute()

countOrgTickets



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
	distinct := openapiclient.org_tickets_count_distinct("") // OrgTicketsCountDistinct |  (optional) (default to "status")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTicketsAPI.CountOrgTickets(context.Background(), orgId).Distinct(distinct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.CountOrgTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountOrgTickets`: RepsonseCount
	fmt.Fprintf(os.Stdout, "Response from `OrgsTicketsAPI.CountOrgTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountOrgTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **distinct** | [**OrgTicketsCountDistinct**](OrgTicketsCountDistinct.md) |  | [default to &quot;status&quot;]

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


## CreateOrgTicket

> Ticket CreateOrgTicket(ctx, orgId).Ticket(ticket).Execute()

createOrgTicket



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
	ticket := *openapiclient.NewTicket("Subject_example", "Type_example") // Ticket | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTicketsAPI.CreateOrgTicket(context.Background(), orgId).Ticket(ticket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.CreateOrgTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgTicket`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `OrgsTicketsAPI.CreateOrgTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ticket** | [**Ticket**](Ticket.md) | Request Body | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgTicket

> Ticket GetOrgTicket(ctx, orgId, ticketId).Start(start).End(end).Duration(duration).Execute()

getOrgTicket



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
	ticketId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTicketsAPI.GetOrgTicket(context.Background(), orgId, ticketId).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.GetOrgTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgTicket`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `OrgsTicketsAPI.GetOrgTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgTicketAttachment

> TicketAttachment GetOrgTicketAttachment(ctx, orgId, ticketId, attachmentId).Start(start).End(end).Duration(duration).Execute()

GetOrgTicketAttachment



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
	ticketId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	attachmentId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTicketsAPI.GetOrgTicketAttachment(context.Background(), orgId, ticketId, attachmentId).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.GetOrgTicketAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgTicketAttachment`: TicketAttachment
	fmt.Fprintf(os.Stdout, "Response from `OrgsTicketsAPI.GetOrgTicketAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ticketId** | **string** |  | 
**attachmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgTicketAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**TicketAttachment**](TicketAttachment.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgTickets

> []Ticket ListOrgTickets(ctx, orgId).Start(start).End(end).Duration(duration).Execute()

listOrgTickets



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
	start := int32(56) // int32 | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified (optional)
	end := int32(56) // int32 | end datetime, can be epoch or relative time like -1d, -2h; now if not specified (optional)
	duration := "10m" // string | duration like 7d, 2w (optional) (default to "1d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTicketsAPI.ListOrgTickets(context.Background(), orgId).Start(start).End(end).Duration(duration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.ListOrgTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgTickets`: []Ticket
	fmt.Fprintf(os.Stdout, "Response from `OrgsTicketsAPI.ListOrgTickets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified | 
 **end** | **int32** | end datetime, can be epoch or relative time like -1d, -2h; now if not specified | 
 **duration** | **string** | duration like 7d, 2w | [default to &quot;1d&quot;]

### Return type

[**[]Ticket**](Ticket.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgTicket

> Ticket UpdateOrgTicket(ctx, orgId, ticketId).Ticket(ticket).Execute()

updateOrgTicket



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
	ticketId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	ticket := *openapiclient.NewTicket("Subject_example", "Type_example") // Ticket | Request Body (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsTicketsAPI.UpdateOrgTicket(context.Background(), orgId, ticketId).Ticket(ticket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.UpdateOrgTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgTicket`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `OrgsTicketsAPI.UpdateOrgTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ticket** | [**Ticket**](Ticket.md) | Request Body | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadrgTicketAttachment

> UploadrgTicketAttachment(ctx, orgId, ticketId).File(file).Execute()

UploadrgTicketAttachment



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
	ticketId := "000000ab-00ab-00ab-00ab-0000000000ab" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | ekahau or ibwave file (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsTicketsAPI.UploadrgTicketAttachment(context.Background(), orgId, ticketId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsTicketsAPI.UploadrgTicketAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadrgTicketAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | ekahau or ibwave file | 

### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form_data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

