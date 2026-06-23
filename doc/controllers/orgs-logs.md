# Orgs Logs

```go
orgsLogs := client.OrgsLogs()
```

## Class Name

`OrgsLogs`

## Methods

* [Count Org Audit Logs](../../doc/controllers/orgs-logs.md#count-org-audit-logs)
* [List Org Audit Logs](../../doc/controllers/orgs-logs.md#list-org-audit-logs)
* [List Org Audit Logs Legacy](../../doc/controllers/orgs-logs.md#list-org-audit-logs-legacy)


# Count Org Audit Logs

Count organization audit log records, optionally grouped by `distinct` and filtered by administrator, site, message text, and time range.

```go
CountOrgAuditLogs(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgLogsCountDistinctEnum,
    adminId *uuid.UUID,
    adminName *string,
    siteId *uuid.UUID,
    message *string,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgLogsCountDistinctEnum`](../../doc/models/org-logs-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `admin_id`, `admin_name`, `message`, `site_id`<br><br>**Default**: `"admin_name"` |
| `adminId` | `*uuid.UUID` | Query, Optional | Filter audit log results by administrator identifier |
| `adminName` | `*string` | Query, Optional | Filter audit log results by one or more administrator names. Supports comma-separated values |
| `siteId` | `*uuid.UUID` | Query, Optional | Filter results by site identifier |
| `message` | `*string` | Query, Optional | Filter log results by message text |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

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
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Audit Logs

Get a list of change logs for the current Org

```go
ListOrgAuditLogs(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    adminName *string,
    message *string,
    sort *models.ListOrgLogsSortEnum,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseLogSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier. Accepts multiple comma-separated values. |
| `adminName` | `*string` | Query, Optional | Filter results by one or more administrator names or email addresses. Supports comma-separated values |
| `message` | `*string` | Query, Optional | Filter results by one or more message text values. Supports comma-separated values |
| `sort` | [`*models.ListOrgLogsSortEnum`](../../doc/models/list-org-logs-sort-enum.md) | Query, Optional | Field used to sort results; a leading `-` indicates descending order. enum: `-timestamp`, `admin_id`, `site_id`, `timestamp` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseLogSearch](../../doc/models/response-log-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"

adminName := "John Doe"

message := "Created a new site"

duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsLogs.ListOrgAuditLogs(ctx, orgId, &siteId, &adminName, &message, nil, nil, nil, &duration, &limit, &page)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Audit Logs Legacy

**This endpoint is deprecated.**

Get List of change logs for the current Org

```go
ListOrgAuditLogsLegacy(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    adminName *string,
    message *string,
    sort *models.ListOrgLogsSortEnum,
    start *string,
    end *string,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseLogSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier. Accepts multiple comma-separated values. |
| `adminName` | `*string` | Query, Optional | Filter results by one or more administrator names or email addresses. Supports comma-separated values |
| `message` | `*string` | Query, Optional | Filter results by one or more message text values. Supports comma-separated values |
| `sort` | [`*models.ListOrgLogsSortEnum`](../../doc/models/list-org-logs-sort-enum.md) | Query, Optional | Field used to sort results; a leading `-` indicates descending order. enum: `-timestamp`, `admin_id`, `site_id`, `timestamp` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseLogSearch](../../doc/models/response-log-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"

adminName := "John Doe"

message := "Created a new site"

duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsLogs.ListOrgAuditLogsLegacy(ctx, orgId, &siteId, &adminName, &message, nil, nil, nil, &duration, &limit, &page)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

