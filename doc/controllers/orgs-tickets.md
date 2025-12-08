# Orgs Tickets

```go
orgsTickets := client.OrgsTickets()
```

## Class Name

`OrgsTickets`

## Methods

* [Add Org Ticket Comment](../../doc/controllers/orgs-tickets.md#add-org-ticket-comment)
* [Count Org Tickets](../../doc/controllers/orgs-tickets.md#count-org-tickets)
* [Create Org Ticket](../../doc/controllers/orgs-tickets.md#create-org-ticket)
* [Get Org Ticket](../../doc/controllers/orgs-tickets.md#get-org-ticket)
* [Get Org Ticket Attachment](../../doc/controllers/orgs-tickets.md#get-org-ticket-attachment)
* [List Org Tickets](../../doc/controllers/orgs-tickets.md#list-org-tickets)
* [Update Org Ticket](../../doc/controllers/orgs-tickets.md#update-org-ticket)
* [Upload Org Ticket Attachment](../../doc/controllers/orgs-tickets.md#upload-org-ticket-attachment)


# Add Org Ticket Comment

Add Comment to support ticket

```go
AddOrgTicketComment(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    comment *string,
    file *string) (
    models.ApiResponse[models.Ticket],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ticketId` | `uuid.UUID` | Template, Required | - |
| `comment` | `*string` | Form, Optional | - |
| `file` | `*string` | Form, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Ticket](../../doc/models/ticket.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ticketId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

comment := "this is urgent"

apiResponse, err := orgsTickets.AddOrgTicketComment(ctx, orgId, ticketId, &comment, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "comments": [
    {
      "attachments": [
        {
          "content_type": "string",
          "content_url": "string"
        }
      ],
      "author": "string",
      "comment": "string",
      "created_at": 0
    }
  ],
  "created_at": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "requester": "string",
  "status": "open",
  "subject": "string",
  "type": "string",
  "updated_at": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Count Org Tickets

Count by Distinct Attributes of Org Tickets

```go
CountOrgTickets(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgTicketsCountDistinctEnum,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgTicketsCountDistinctEnum`](../../doc/models/org-tickets-count-distinct-enum.md) | Query, Optional | **Default**: `"status"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgTicketsCountDistinctEnum_STATUS

limit := 100

apiResponse, err := orgsTickets.CountOrgTickets(ctx, orgId, &distinct, &limit)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "distinct": "string",
  "end": 0,
  "limit": 0,
  "results": [
    {
      "count": 0,
      "property": "string"
    }
  ],
  "start": 0,
  "total": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Create Org Ticket

Create a support ticket

```go
CreateOrgTicket(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Ticket) (
    models.ApiResponse[models.Ticket],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Ticket`](../../doc/models/ticket.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Ticket](../../doc/models/ticket.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Ticket{
    Subject:              "string",
    Type:                 "question",
    AdditionalProperties: map[string]interface{}{
        "comment": interface{}("string"),
    },
}

apiResponse, err := orgsTickets.CreateOrgTicket(ctx, orgId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "comments": [
    {
      "attachments": [
        {
          "content_type": "string",
          "content_url": "string"
        }
      ],
      "author": "string",
      "comment": "string",
      "created_at": 0
    }
  ],
  "created_at": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "requester": "string",
  "status": "open",
  "subject": "string",
  "type": "string",
  "updated_at": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Ticket

Get support ticket details

```go
GetOrgTicket(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.Ticket],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ticketId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Ticket](../../doc/models/ticket.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ticketId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

apiResponse, err := orgsTickets.GetOrgTicket(ctx, orgId, ticketId, nil, nil, &duration)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "comments": [
    {
      "attachments": [
        {
          "content_type": "string",
          "content_url": "string"
        }
      ],
      "author": "string",
      "comment": "string",
      "created_at": 0
    }
  ],
  "created_at": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "requester": "string",
  "status": "open",
  "subject": "string",
  "type": "string",
  "updated_at": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Ticket Attachment

Get Org ticket Attachment

```go
GetOrgTicketAttachment(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    attachmentId uuid.UUID,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.TicketAttachment],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ticketId` | `uuid.UUID` | Template, Required | - |
| `attachmentId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.TicketAttachment](../../doc/models/ticket-attachment.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ticketId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

attachmentId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

apiResponse, err := orgsTickets.GetOrgTicketAttachment(ctx, orgId, ticketId, attachmentId, nil, nil, &duration)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "content_url": "https://api.mist.com/api/v1/forward/download?jwt=..."
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Tickets

Get List of Tickets of an Org

```go
ListOrgTickets(
    ctx context.Context,
    orgId uuid.UUID,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[[]models.Ticket],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Ticket](../../doc/models/ticket.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

apiResponse, err := orgsTickets.ListOrgTickets(ctx, orgId, nil, nil, &duration)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "comments": [
      {
        "attachments": [
          {
            "content_type": "string",
            "content_url": "string"
          }
        ],
        "author": "string",
        "comment": "string",
        "created_at": 0
      }
    ],
    "created_at": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "requester": "string",
    "status": "open",
    "subject": "string",
    "type": "string",
    "updated_at": 0
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Ticket

Update support ticket

```go
UpdateOrgTicket(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    body *models.Ticket) (
    models.ApiResponse[models.Ticket],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ticketId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Ticket`](../../doc/models/ticket.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Ticket](../../doc/models/ticket.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ticketId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Ticket{
    Subject:              "string",
    Type:                 "question",
    AdditionalProperties: map[string]interface{}{
        "comment": interface{}("string"),
    },
}

apiResponse, err := orgsTickets.UpdateOrgTicket(ctx, orgId, ticketId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "comments": [
    {
      "attachments": [
        {
          "content_type": "string",
          "content_url": "string"
        }
      ],
      "author": "string",
      "comment": "string",
      "created_at": 0
    }
  ],
  "created_at": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "requester": "string",
  "status": "open",
  "subject": "string",
  "type": "string",
  "updated_at": 0
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Upload Org Ticket Attachment

Get Org ticket Attachment

```go
UploadOrgTicketAttachment(
    ctx context.Context,
    orgId uuid.UUID,
    ticketId uuid.UUID,
    file *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `ticketId` | `uuid.UUID` | Template, Required | - |
| `file` | `*string` | Form, Optional | Ekahau or ibwave file |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

ticketId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsTickets.UploadOrgTicketAttachment(ctx, orgId, ticketId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

