# MS Ps Tickets

```go
mSPsTickets := client.MSPsTickets()
```

## Class Name

`MSPsTickets`

## Methods

* [Count Msp Tickets](../../doc/controllers/ms-ps-tickets.md#count-msp-tickets)
* [List Msp Tickets](../../doc/controllers/ms-ps-tickets.md#list-msp-tickets)


# Count Msp Tickets

Count by Distinct Attributes of tickets

```go
CountMspTickets(
    ctx context.Context,
    mspId uuid.UUID,
    distinct *models.MspTicketsCountDistinctEnum,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.MspTicketsCountDistinctEnum`](../../doc/models/msp-tickets-count-distinct-enum.md) | Query, Optional | **Default**: `"status"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.MspTicketsCountDistinctEnum_STATUS

limit := 100

apiResponse, err := mSPsTickets.CountMspTickets(ctx, mspId, &distinct, &limit)
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


# List Msp Tickets

Get List of Tickets of a MSP

```go
ListMspTickets(
    ctx context.Context,
    mspId uuid.UUID,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[[]models.Ticket],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Ticket](../../doc/models/ticket.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

duration := "10m"

apiResponse, err := mSPsTickets.ListMspTickets(ctx, mspId, nil, nil, &duration)
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

