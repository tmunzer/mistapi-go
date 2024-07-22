# Sites Psks

```go
sitesPsks := client.SitesPsks()
```

## Class Name

`SitesPsks`

## Methods

* [Create Site Psk](../../doc/controllers/sites-psks.md#create-site-psk)
* [Delete Site Psk](../../doc/controllers/sites-psks.md#delete-site-psk)
* [Get Site Psk](../../doc/controllers/sites-psks.md#get-site-psk)
* [Import Site Psks](../../doc/controllers/sites-psks.md#import-site-psks)
* [List Site Psks](../../doc/controllers/sites-psks.md#list-site-psks)
* [Update Site Multiple Psks](../../doc/controllers/sites-psks.md#update-site-multiple-psks)
* [Update Site Psk](../../doc/controllers/sites-psks.md#update-site-psk)


# Create Site Psk

Create Site PSK

When `usage`==`macs`, corresponding "macs" field will hold a list consisting of client mac addresses (["xx:xx:xx:xx:xx",...]) or mac patterns(["xx:xx:*","xx*",...]) or both (["xx:xx:xx:xx:xx:xx", "xx:*", ...]). This list is capped at 5000

```go
CreateSitePsk(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.Psk) (
    models.ApiResponse[models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Psk`](../../doc/models/psk.md) | Body, Optional | Request Body |

## Response Type

[`models.Psk`](../../doc/models/psk.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Psk{
    Mac:                    models.ToPointer("string"),
    Name:                   "string",
    Passphrase:             "stringst",
    Ssid:                   "string",
    Usage:                  models.ToPointer(models.PskUsageEnum("multi")),
    VlanId:                 models.ToPointer(models.PskVlanIdContainer.FromNumber(1)),
}

apiResponse, err := sitesPsks.CreateSitePsk(ctx, siteId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Delete Site Psk

Delete Site PSK

```go
DeleteSitePsk(
    ctx context.Context,
    siteId uuid.UUID,
    pskId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesPsks.DeleteSitePsk(ctx, siteId, pskId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Site Psk

Get Site PSK Details

```go
GetSitePsk(
    ctx context.Context,
    siteId uuid.UUID,
    pskId uuid.UUID) (
    models.ApiResponse[models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |

## Response Type

[`models.Psk`](../../doc/models/psk.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesPsks.GetSitePsk(ctx, siteId, pskId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Import Site Psks

Import PSK from CSV file or JSON

## CSV File Format

```csv
PSK Import CSV File Format:
name,ssid,passphrase,usage,vlan_id,mac
Common,warehouse,foryoureyesonly,single,35,a31425f31278
Justin,reception,visible,multi,1002
```

```go
ImportSitePsks(
    ctx context.Context,
    siteId uuid.UUID,
    file *models.FileWrapper) (
    models.ApiResponse[[]models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `file` | `*models.FileWrapper` | Form, Optional | - |

## Response Type

[`[]models.Psk`](../../doc/models/psk.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := sitesPsks.ImportSitePsks(ctx, siteId, nil)
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
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "passphrase": "stringst",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ssid": "string",
    "usage": "multi",
    "vlan_id": 1
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# List Site Psks

Get List of Site PSKs

```go
ListSitePsks(
    ctx context.Context,
    siteId uuid.UUID,
    ssid *string,
    role *string,
    name *string,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `ssid` | `*string` | Query, Optional | - |
| `role` | `*string` | Query, Optional | - |
| `name` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.Psk`](../../doc/models/psk.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







page := 1

limit := 100

apiResponse, err := sitesPsks.ListSitePsks(ctx, siteId, nil, nil, nil, &page, &limit)
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
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "passphrase": "stringst",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ssid": "string",
    "usage": "multi",
    "vlan_id": 1
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Site Multiple Psks

Update multiple PSKs

```go
UpdateSiteMultiplePsks(
    ctx context.Context,
    siteId uuid.UUID,
    body []models.Psk) (
    models.ApiResponse[[]models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`[]models.Psk`](../../doc/models/psk.md) | Body, Optional | - |

## Response Type

[`[]models.Psk`](../../doc/models/psk.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := []models.Psk{
    models.Psk{
        ExpireTime:             models.NewOptional(models.ToPointer(1614990263)),
        Macs:                   []string{
            "11:22:33:44:55:66",
            "aa:bb:",
            "53",
        },
        MaxUsage:               models.ToPointer(0),
        Name:                   "name6",
        NotifyExpiry:           models.ToPointer(false),
        OrgId:                  models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Passphrase:             "passphrase6",
        SiteId:                 models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Ssid:                   "ssid6",
        Usage:                  models.ToPointer(models.PskUsageEnum("multi")),
    },
}

apiResponse, err := sitesPsks.UpdateSiteMultiplePsks(ctx, siteId, body)
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
    "created_time": 0,
    "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "mac": "string",
    "modified_time": 0,
    "name": "string",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "passphrase": "stringst",
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "ssid": "string",
    "usage": "multi",
    "vlan_id": 1
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Update Site Psk

Update Site PSK

```go
UpdateSitePsk(
    ctx context.Context,
    siteId uuid.UUID,
    pskId uuid.UUID,
    body *models.Psk) (
    models.ApiResponse[models.Psk],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `pskId` | `uuid.UUID` | Template, Required | PSK ID |
| `body` | [`*models.Psk`](../../doc/models/psk.md) | Body, Optional | Request Body |

## Response Type

[`models.Psk`](../../doc/models/psk.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Psk{
    Mac:                    models.ToPointer("string"),
    Name:                   "string",
    Passphrase:             "stringst",
    Ssid:                   "string",
    Usage:                  models.ToPointer(models.PskUsageEnum("multi")),
    VlanId:                 models.ToPointer(models.PskVlanIdContainer.FromNumber(1)),
}

apiResponse, err := sitesPsks.UpdateSitePsk(ctx, siteId, pskId, &body)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

