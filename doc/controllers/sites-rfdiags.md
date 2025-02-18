# Sites Rfdiags

```go
sitesRfdiags := client.SitesRfdiags()
```

## Class Name

`SitesRfdiags`

## Methods

* [Delete Site Rfdiag Recording](../../doc/controllers/sites-rfdiags.md#delete-site-rfdiag-recording)
* [Download Site Rfdiag Recording](../../doc/controllers/sites-rfdiags.md#download-site-rfdiag-recording)
* [Get Site Rfdiag Recording](../../doc/controllers/sites-rfdiags.md#get-site-rfdiag-recording)
* [Get Site Site Rfdiag Recording](../../doc/controllers/sites-rfdiags.md#get-site-site-rfdiag-recording)
* [Start Site Recording](../../doc/controllers/sites-rfdiags.md#start-site-recording)
* [Stop Site Rfdiag Recording](../../doc/controllers/sites-rfdiags.md#stop-site-rfdiag-recording)
* [Update Site Rfdiag Recording](../../doc/controllers/sites-rfdiags.md#update-site-rfdiag-recording)


# Delete Site Rfdiag Recording

Delete Recording

```go
DeleteSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rfdiagId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rfdiagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesRfdiags.DeleteSiteRfdiagRecording(ctx, siteId, rfdiagId)
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


# Download Site Rfdiag Recording

Download Recording
Download raw_events blob

```go
DownloadSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    models.ApiResponse[[]byte],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rfdiagId` | `uuid.UUID` | Template, Required | - |

## Response Type

`[]byte`

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rfdiagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesRfdiags.DownloadSiteRfdiagRecording(ctx, siteId, rfdiagId)
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


# Get Site Rfdiag Recording

Get RF Diage Recording Details

```go
GetSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    models.ApiResponse[[]models.RfDiagInfoItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rfdiagId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.RfDiagInfoItem`](../../doc/models/rf-diag-info-item.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rfdiagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesRfdiags.GetSiteRfdiagRecording(ctx, siteId, rfdiagId)
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
    "asset_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "asset_name": "string",
    "client_name": "string",
    "duration": 0,
    "end_time": 0,
    "frame_count": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "name": "string",
    "next": "string",
    "raw_events": "string",
    "ready": true,
    "sdkclient_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "sdkclient_name": "string",
    "sdkclient_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "start_time": 0,
    "type": "sdkclient",
    "url": "string"
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


# Get Site Site Rfdiag Recording

List RF Glass Recording

```go
GetSiteSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[][]models.RfDiagInfoItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[][]models.RfDiagInfoItem`](../../doc/models/rf-diag-info-item.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesRfdiags.GetSiteSiteRfdiagRecording(ctx, siteId, nil, nil, &duration, &limit, &page)
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
  [
    {
      "asset_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "asset_name": "string",
      "client_name": "string",
      "duration": 0,
      "end_time": 0,
      "frame_count": 0,
      "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "mac": "string",
      "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "name": "string",
      "next": "string",
      "raw_events": "string",
      "ready": true,
      "sdkclient_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "sdkclient_name": "string",
      "sdkclient_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "start_time": 0,
      "type": "sdkclient",
      "url": "string"
    }
  ]
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


# Start Site Recording

Start RF Glass Recording

```go
StartSiteRecording(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.RfDiag) (
    models.ApiResponse[[]models.RfDiagInfoItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RfDiag`](../../doc/models/rf-diag.md) | Body, Optional | Request Body |

## Response Type

[`[]models.RfDiagInfoItem`](../../doc/models/rf-diag-info-item.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RfDiag{
    Duration:             models.ToPointer(180),
    Name:                 "name6",
    Type:                 models.RfClientTypeEnum_CLIENT,
}

apiResponse, err := sitesRfdiags.StartSiteRecording(ctx, siteId, &body)
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
    "asset_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "asset_name": "string",
    "client_name": "string",
    "duration": 0,
    "end_time": 0,
    "frame_count": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "name": "string",
    "next": "string",
    "raw_events": "string",
    "ready": true,
    "sdkclient_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "sdkclient_name": "string",
    "sdkclient_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "start_time": 0,
    "type": "sdkclient",
    "url": "string"
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


# Stop Site Rfdiag Recording

If the recording session is active for the given rfdiag_id, it will finish the recording. duration and end_time will be updated to reflect the correct values.

```go
StopSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rfdiagId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rfdiagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesRfdiags.StopSiteRfdiagRecording(ctx, siteId, rfdiagId)
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


# Update Site Rfdiag Recording

Update Recording

```go
UpdateSiteRfdiagRecording(
    ctx context.Context,
    siteId uuid.UUID,
    rfdiagId uuid.UUID,
    body *models.RfDiag) (
    models.ApiResponse[[]models.RfDiagInfoItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `rfdiagId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RfDiag`](../../doc/models/rf-diag.md) | Body, Optional | Request Body |

## Response Type

[`[]models.RfDiagInfoItem`](../../doc/models/rf-diag-info-item.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rfdiagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RfDiag{
    Duration:             models.ToPointer(180),
    Name:                 "name6",
    Type:                 models.RfClientTypeEnum_CLIENT,
}

apiResponse, err := sitesRfdiags.UpdateSiteRfdiagRecording(ctx, siteId, rfdiagId, &body)
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
    "asset_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "asset_name": "string",
    "client_name": "string",
    "duration": 0,
    "end_time": 0,
    "frame_count": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "map_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "name": "string",
    "next": "string",
    "raw_events": "string",
    "ready": true,
    "sdkclient_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "sdkclient_name": "string",
    "sdkclient_uuid": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "start_time": 0,
    "type": "sdkclient",
    "url": "string"
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

