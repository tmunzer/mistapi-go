# Sites Wx Tags

```go
sitesWxTags := client.SitesWxTags()
```

## Class Name

`SitesWxTags`

## Methods

* [Create Site Wx Tag](../../doc/controllers/sites-wx-tags.md#create-site-wx-tag)
* [Delete Site Wx Tag](../../doc/controllers/sites-wx-tags.md#delete-site-wx-tag)
* [Get Site Application List](../../doc/controllers/sites-wx-tags.md#get-site-application-list)
* [Get Site Current Matching Clients of a Wx Tag](../../doc/controllers/sites-wx-tags.md#get-site-current-matching-clients-of-a-wx-tag)
* [Get Site Wx Tag](../../doc/controllers/sites-wx-tags.md#get-site-wx-tag)
* [List Site Wx Tags](../../doc/controllers/sites-wx-tags.md#list-site-wx-tags)
* [Update Site Wx Tag](../../doc/controllers/sites-wx-tags.md#update-site-wx-tag)


# Create Site Wx Tag

Create Site WxTag

```go
CreateSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.WxlanTag) (
    models.ApiResponse[models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTag`](../../doc/models/wxlan-tag.md) | Body, Optional | Request Body |

## Response Type

[`models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTag{
    Match:        models.ToPointer(models.WxlanTagMatchEnum("app")),
    Name:         "match app",
    Type:         models.WxlanTagTypeEnum("match"),
    Values:       []string{
        "gmail",
        "dropbox",
    },
}

apiResponse, err := sitesWxTags.CreateSiteWxTag(ctx, siteId, &body)
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
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "last_ips": [
    "string"
  ],
  "mac": "string",
  "match": "wlan_id",
  "modified_time": 0,
  "name": "string",
  "op": "in",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "resource_mac": "string",
  "services": [
    "string"
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "specs": [
    {
      "port_range": "string",
      "protocol": "tcp",
      "subnet": [
        "string"
      ]
    }
  ],
  "subnet": "string",
  "type": "match",
  "values": [
    "string"
  ]
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


# Delete Site Wx Tag

Delete Site WxTag

```go
DeleteSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    wxtagId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesWxTags.DeleteSiteWxTag(ctx, siteId, wxtagId)
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


# Get Site Application List

Get Application List

```go
GetSiteApplicationList(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]models.SearchWxtagAppsItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.SearchWxtagAppsItem`](../../doc/models/search-wxtag-apps-item.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesWxTags.GetSiteApplicationList(ctx, siteId)
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
    "group": "Emails",
    "key": "gmail",
    "name": "Gmail - web/app"
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


# Get Site Current Matching Clients of a Wx Tag

Get Current Matching Clients of a WXLAN Tag

```go
GetSiteCurrentMatchingClientsOfAWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    wxtagId uuid.UUID) (
    models.ApiResponse[[]models.WxtagMatching],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.WxtagMatching`](../../doc/models/wxtag-matching.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesWxTags.GetSiteCurrentMatchingClientsOfAWxTag(ctx, siteId, wxtagId)
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
    "mac": "5684dae9ac8b",
    "since": 1428939600
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


# Get Site Wx Tag

Get Site WxTag Details

```go
GetSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    wxtagId uuid.UUID) (
    models.ApiResponse[models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesWxTags.GetSiteWxTag(ctx, siteId, wxtagId)
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
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "last_ips": [
    "string"
  ],
  "mac": "string",
  "match": "wlan_id",
  "modified_time": 0,
  "name": "string",
  "op": "in",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "resource_mac": "string",
  "services": [
    "string"
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "specs": [
    {
      "port_range": "string",
      "protocol": "tcp",
      "subnet": [
        "string"
      ]
    }
  ],
  "subnet": "string",
  "type": "match",
  "values": [
    "string"
  ]
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


# List Site Wx Tags

Get List of Site WxTags

```go
ListSiteWxTags(
    ctx context.Context,
    siteId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := sitesWxTags.ListSiteWxTags(ctx, siteId, &limit, &page)
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
    "last_ips": [
      "string"
    ],
    "mac": "string",
    "match": "wlan_id",
    "modified_time": 0,
    "name": "string",
    "op": "in",
    "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "resource_mac": "string",
    "services": [
      "string"
    ],
    "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
    "specs": [
      {
        "port_range": "string",
        "protocol": "tcp",
        "subnet": [
          "string"
        ]
      }
    ],
    "subnet": "string",
    "type": "match",
    "values": [
      "string"
    ]
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


# Update Site Wx Tag

Update Site WxTag

```go
UpdateSiteWxTag(
    ctx context.Context,
    siteId uuid.UUID,
    wxtagId uuid.UUID,
    body *models.WxlanTag) (
    models.ApiResponse[models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTag`](../../doc/models/wxlan-tag.md) | Body, Optional | Request Body |

## Response Type

[`models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTag{
    Match:        models.ToPointer(models.WxlanTagMatchEnum("app")),
    Name:         "match app",
    Type:         models.WxlanTagTypeEnum("match"),
    Values:       []string{
        "gmail",
        "dropbox",
    },
}

apiResponse, err := sitesWxTags.UpdateSiteWxTag(ctx, siteId, wxtagId, &body)
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
  "created_time": 0,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "last_ips": [
    "string"
  ],
  "mac": "string",
  "match": "wlan_id",
  "modified_time": 0,
  "name": "string",
  "op": "in",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "resource_mac": "string",
  "services": [
    "string"
  ],
  "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "specs": [
    {
      "port_range": "string",
      "protocol": "tcp",
      "subnet": [
        "string"
      ]
    }
  ],
  "subnet": "string",
  "type": "match",
  "values": [
    "string"
  ]
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

