# Orgs Logs

```go
orgsLogs := client.OrgsLogs()
```

## Class Name

`OrgsLogs`

## Methods

* [Count Org Audit Logs](../../doc/controllers/orgs-logs.md#count-org-audit-logs)
* [List Org Audit Logs](../../doc/controllers/orgs-logs.md#list-org-audit-logs)


# Count Org Audit Logs

Count by Distinct Attributes of Audit Logs

```go
CountOrgAuditLogs(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgLogsCountDistinctEnum,
    adminId *string,
    adminName *string,
    siteId *string,
    message *string,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgLogsCountDistinctEnum`](../../doc/models/org-logs-count-distinct-enum.md) | Query, Optional | - |
| `adminId` | `*string` | Query, Optional | - |
| `adminName` | `*string` | Query, Optional | - |
| `siteId` | `*string` | Query, Optional | - |
| `message` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgLogsCountDistinctEnum("admin_name")









page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsLogs.CountOrgAuditLogs(ctx, orgId, &distinct, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Org Audit Logs

Get List of change logs for the current Org

```go
ListOrgAuditLogs(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    adminName *string,
    message *string,
    sort *models.ListOrgLogsSortEnum,
    start *int,
    end *int,
    limit *int,
    page *int,
    duration *string) (
    models.ApiResponse[models.ResponseLogSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | site id |
| `adminName` | `*string` | Query, Optional | admin name or email |
| `message` | `*string` | Query, Optional | message |
| `sort` | [`*models.ListOrgLogsSortEnum`](../../doc/models/list-org-logs-sort-enum.md) | Query, Optional | sort order |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseLogSearch`](../../doc/models/response-log-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













limit := 100

page := 1

duration := "10m"

apiResponse, err := orgsLogs.ListOrgAuditLogs(ctx, orgId, nil, nil, nil, nil, nil, nil, &limit, &page, &duration)
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
      "admin_id": "72bfa2bd-e58a-4670-9d20-a1468f7a6f58",
      "admin_name": "test@mistsys.com",
      "id": "c6f9347b-b0a4-4a23-b927-fa9249f2ffb2",
      "message": "TEST AUDIT",
      "org_id": "469f6eca-6276-4993-bfeb-53cbbbba6f58",
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

