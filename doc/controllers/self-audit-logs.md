# Self Audit Logs

```go
selfAuditLogs := client.SelfAuditLogs()
```

## Class Name

`SelfAuditLogs`


# List Self Audit Logs

Get List of change logs across all Orgs for current admin
Audit logs records all administrative activities done by current admin across all orgs

```go
ListSelfAuditLogs(
    ctx context.Context,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseSelfAuditLogs],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSelfAuditLogs](../../doc/models/response-self-audit-logs.md).

## Example Usage

```go
ctx := context.Background()





duration := "10m"

limit := 100

page := 1

apiResponse, err := selfAuditLogs.ListSelfAuditLogs(ctx, nil, nil, &duration, &limit, &page)
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
  "page": 1,
  "results": [
    {
      "admin_id": "72bfa2bd-e58a-4670-9d20-a1468f7a6f58",
      "admin_name": "test@mistsys.com",
      "after": {
        "auth": {
          "type": "open"
        }
      },
      "before": {
        "auth": {
          "type": "psk"
        }
      },
      "id": "c6f9347b-b0a4-4a23-b927-fa9249f2ffb2",
      "message": "Update WLAN \"Corporate\"",
      "org_id": "423f6eca-6276-4994-bfeb-53cbbbba6f04",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "timestamp": 1431382121
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

