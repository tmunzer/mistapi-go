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
    adminId *uuid.UUID,
    adminName *string,
    siteId *uuid.UUID,
    message *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgLogsCountDistinctEnum`](../../doc/models/org-logs-count-distinct-enum.md) | Query, Optional | **Default**: `"admin_name"` |
| `adminId` | `*uuid.UUID` | Query, Optional | - |
| `adminName` | `*string` | Query, Optional | - |
| `siteId` | `*uuid.UUID` | Query, Optional | - |
| `message` | `*string` | Query, Optional | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgLogsCountDistinctEnum_ADMINNAME

adminId := uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")

adminName := "John Doe"

siteId := uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")

message := "Created a new site"





duration := "10m"

limit := 100

apiResponse, err := orgsLogs.CountOrgAuditLogs(ctx, orgId, &distinct, &adminId, &adminName, &siteId, &message, nil, nil, &duration, &limit)
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


# List Org Audit Logs

Get List of change logs for the current Org

```go
ListOrgAuditLogs(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *uuid.UUID,
    adminName *string,
    message *string,
    sort *models.ListOrgLogsSortEnum,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*uuid.UUID` | Query, Optional | Site id |
| `adminName` | `*string` | Query, Optional | Admin name or email |
| `message` | `*string` | Query, Optional | Message |
| `sort` | [`*models.ListOrgLogsSortEnum`](../../doc/models/list-org-logs-sort-enum.md) | Query, Optional | Sort order |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseLogSearch](../../doc/models/response-log-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("4ac1dcf4-9d8b-7211-65c4-057819f0862b")

adminName := "John Doe"

message := "Created a new site"







duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsLogs.ListOrgAuditLogs(ctx, orgId, &siteId, &adminName, &message, nil, nil, nil, &duration, &limit, &page)
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
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
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

