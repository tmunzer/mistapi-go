# Sites Alarms

```go
sitesAlarms := client.SitesAlarms()
```

## Class Name

`SitesAlarms`

## Methods

* [Ack Site Alarm](../../doc/controllers/sites-alarms.md#ack-site-alarm)
* [Ack Site All Alarms](../../doc/controllers/sites-alarms.md#ack-site-all-alarms)
* [Ack Site Multiple Alarms](../../doc/controllers/sites-alarms.md#ack-site-multiple-alarms)
* [Count Site Alarms](../../doc/controllers/sites-alarms.md#count-site-alarms)
* [Search Site Alarms](../../doc/controllers/sites-alarms.md#search-site-alarms)
* [Subscribe Site Alarms](../../doc/controllers/sites-alarms.md#subscribe-site-alarms)
* [Unack Site Alarm](../../doc/controllers/sites-alarms.md#unack-site-alarm)
* [Unack Site All Arlarms](../../doc/controllers/sites-alarms.md#unack-site-all-arlarms)
* [Unack Site Multiple Alarms](../../doc/controllers/sites-alarms.md#unack-site-multiple-alarms)
* [Unsubscribe Site Alarms](../../doc/controllers/sites-alarms.md#unsubscribe-site-alarms)


# Ack Site Alarm

Ack Site Alarm

```go
AckSiteAlarm(
    ctx context.Context,
    siteId uuid.UUID,
    alarmId uuid.UUID,
    body *models.NoteString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `alarmId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NoteString`](../../doc/models/note-string.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

alarmId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NoteString{
    Note: models.ToPointer("maintenance window"),
}

resp, err := sitesAlarms.AckSiteAlarm(ctx, siteId, alarmId, &body)
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


# Ack Site All Alarms

Ack all Site Alarms

**N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.

```go
AckSiteAllAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.NoteString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NoteString`](../../doc/models/note-string.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NoteString{
    Note: models.ToPointer("maintenance window"),
}

resp, err := sitesAlarms.AckSiteAllAlarms(ctx, siteId, &body)
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


# Ack Site Multiple Alarms

Ack multiple Site Alarms

```go
AckSiteMultipleAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AlarmAck) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AlarmAck`](../../doc/models/alarm-ack.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AlarmAck{
    AlarmIds: []uuid.UUID{
        uuid.MustParse("ccb8c94d-ca56-4075-932f-1f2ab444ff2c"),
        uuid.MustParse("98ff4a3d-ec9b-4138-a42e-54fc3335179d"),
    },
    Note:     models.ToPointer("maintenance window"),
}

resp, err := sitesAlarms.AckSiteMultipleAlarms(ctx, siteId, &body)
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


# Count Site Alarms

Count Site Alarms

```go
CountSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.AlarmCountDisctinctEnum,
    ackAdminName *string,
    acked *bool,
    mType *string,
    severity *string,
    group *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.AlarmCountDisctinctEnum`](../../doc/models/alarm-count-disctinct-enum.md) | Query, Optional | Group by and count the alarms by some distinct field |
| `ackAdminName` | `*string` | Query, Optional | Name of the admins who have acked the alarms; accepts multiple values separated by comma |
| `acked` | `*bool` | Query, Optional | - |
| `mType` | `*string` | Query, Optional | Key-name of the alarms; accepts multiple values separated by comma |
| `severity` | `*string` | Query, Optional | Alarm severity; accepts multiple values separated by comma |
| `group` | `*string` | Query, Optional | Alarm group name; accepts multiple values separated by comma |
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

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.AlarmCountDisctinctEnum("type")











page := 1

limit := 100





duration := "10m"

apiResponse, err := sitesAlarms.CountSiteAlarms(ctx, siteId, &distinct, nil, nil, nil, nil, nil, &page, &limit, nil, nil, &duration)
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


# Search Site Alarms

Search Site Alarms

```go
SearchSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    ackAdminName *string,
    acked *bool,
    severity *string,
    group *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.AlarmSearchResult],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Key-name of the alarms; accepts multiple values separated by comma |
| `ackAdminName` | `*string` | Query, Optional | Name of the admins who have acked the alarms; accepts multiple values separated by comma |
| `acked` | `*bool` | Query, Optional | - |
| `severity` | `*string` | Query, Optional | Alarm severity; accepts multiple values separated by comma |
| `group` | `*string` | Query, Optional | Alarm group name; accepts multiple values separated by comma |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.AlarmSearchResult`](../../doc/models/alarm-search-result.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")











limit := 100





duration := "10m"

apiResponse, err := sitesAlarms.SearchSiteAlarms(ctx, siteId, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Subscribe Site Alarms

Subscribe to Site Alarms

```go
SubscribeSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesAlarms.SubscribeSiteAlarms(ctx, siteId)
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


# Unack Site Alarm

Unack Site Alarm

```go
UnackSiteAlarm(
    ctx context.Context,
    siteId uuid.UUID,
    alarmId uuid.UUID,
    body *models.NoteString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `alarmId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NoteString`](../../doc/models/note-string.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

alarmId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NoteString{
    Note: models.ToPointer("maintenance window"),
}

resp, err := sitesAlarms.UnackSiteAlarm(ctx, siteId, alarmId, &body)
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


# Unack Site All Arlarms

Unack all Site Alarms

**N.B.**: Batch size for multiple alarm ack and unack has to be less or or equal to 1000.

```go
UnackSiteAllArlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.NoteString) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NoteString`](../../doc/models/note-string.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NoteString{
    Note: models.ToPointer("maintenance window"),
}

resp, err := sitesAlarms.UnackSiteAllArlarms(ctx, siteId, &body)
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


# Unack Site Multiple Alarms

Unack multiple Site Alarms

```go
UnackSiteMultipleAlarms(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AlarmAck) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AlarmAck`](../../doc/models/alarm-ack.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AlarmAck{
    AlarmIds: []uuid.UUID{
        uuid.MustParse("ccb8c94d-ca56-4075-932f-1f2ab444ff2c"),
        uuid.MustParse("98ff4a3d-ec9b-4138-a42e-54fc3335179d"),
    },
    Note:     models.ToPointer("maintenance window"),
}

resp, err := sitesAlarms.UnackSiteMultipleAlarms(ctx, siteId, &body)
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


# Unsubscribe Site Alarms

Unsubscribe to Site Alarms

```go
UnsubscribeSiteAlarms(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesAlarms.UnsubscribeSiteAlarms(ctx, siteId)
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

