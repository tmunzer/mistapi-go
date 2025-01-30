# MS Ps Logs

```go
mSPsLogs := client.MSPsLogs()
```

## Class Name

`MSPsLogs`

## Methods

* [Count Msp Audit Logs](../../doc/controllers/ms-ps-logs.md#count-msp-audit-logs)
* [List Msp Audit Logs](../../doc/controllers/ms-ps-logs.md#list-msp-audit-logs)


# Count Msp Audit Logs

Count by Distinct Attributes of Audit Logs

```go
CountMspAuditLogs(
    ctx context.Context,
    mspId uuid.UUID,
    distinct *models.MspLogsCountDistinctEnum) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.MspLogsCountDistinctEnum`](../../doc/models/msp-logs-count-distinct-enum.md) | Query, Optional | **Default**: `"admin_name"` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.MspLogsCountDistinctEnum_ADMINNAME

apiResponse, err := mSPsLogs.CountMspAuditLogs(ctx, mspId, &distinct)
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


# List Msp Audit Logs

Get list of change logs for the current MSP

```go
ListMspAuditLogs(
    ctx context.Context,
    mspId uuid.UUID,
    siteId *string,
    adminName *string,
    message *string,
    sort *models.ListMspLogsSortEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseLogSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | site id |
| `adminName` | `*string` | Query, Optional | admin name or email |
| `message` | `*string` | Query, Optional | message |
| `sort` | [`*models.ListMspLogsSortEnum`](../../doc/models/list-msp-logs-sort-enum.md) | Query, Optional | sort order |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponseLogSearch`](../../doc/models/response-log-search.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













duration := "10m"

limit := 100

page := 1

apiResponse, err := mSPsLogs.ListMspAuditLogs(ctx, mspId, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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
  "end": 1428954000,
  "limit": 100,
  "results": [
    {
      "admin_id": "48f4d7aa-97a0-43e1-81f7-74dbda4a9dae",
      "admin_name": "Chia-Wei Tang tangc@juniper.net",
      "id": "ac4415e5-7aef-4c79-9d17-7a7edd268e16",
      "message": "Accessed Org \"DELETE_ME\"",
      "org_id": "b0d0c697-c4c8-459a-bf61-bfe820aead98",
      "site_id": null,
      "src_ip": "165.225.242.194",
      "timestamp": 1729278563,
      "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:130.0) Gecko/20100101 Firefox/130.0"
    }
  ],
  "start": 1428939600,
  "total": 135
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

