# Orgs Wx Tags

```go
orgsWxTags := client.OrgsWxTags()
```

## Class Name

`OrgsWxTags`

## Methods

* [Create Org Wx Tag](../../doc/controllers/orgs-wx-tags.md#create-org-wx-tag)
* [Delete Org Wx Tag](../../doc/controllers/orgs-wx-tags.md#delete-org-wx-tag)
* [Get Org Application List](../../doc/controllers/orgs-wx-tags.md#get-org-application-list)
* [Get Org Current Matching Clients of a Wx Tag](../../doc/controllers/orgs-wx-tags.md#get-org-current-matching-clients-of-a-wx-tag)
* [Get Org Wx Tag](../../doc/controllers/orgs-wx-tags.md#get-org-wx-tag)
* [List Org Wx Tags](../../doc/controllers/orgs-wx-tags.md#list-org-wx-tags)
* [Update Org Wx Tag](../../doc/controllers/orgs-wx-tags.md#update-org-wx-tag)


# Create Org Wx Tag

Create WxLAN Tag

```go
CreateOrgWxTag(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.WxlanTag) (
    models.ApiResponse[models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTag`](../../doc/models/wxlan-tag.md) | Body, Optional | Request Body |

## Response Type

[`models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTag{
    Match:        models.ToPointer(models.WxlanTagMatchEnum("app")),
    Name:         "match app",
    Type:         models.WxlanTagTypeEnum("match"),
    Values:       []string{
        "gmail",
        "dropbox",
    },
}

apiResponse, err := orgsWxTags.CreateOrgWxTag(ctx, orgId, &body)
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


# Delete Org Wx Tag

Delete WxLAN Tag

```go
DeleteOrgWxTag(
    ctx context.Context,
    orgId uuid.UUID,
    wxtagId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsWxTags.DeleteOrgWxTag(ctx, orgId, wxtagId)
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


# Get Org Application List

Get Application List

```go
GetOrgApplicationList(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.SearchWxtagAppsItem],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.SearchWxtagAppsItem`](../../doc/models/search-wxtag-apps-item.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWxTags.GetOrgApplicationList(ctx, orgId)
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


# Get Org Current Matching Clients of a Wx Tag

Get Current Matching Clients of a WXLAN Tag

```go
GetOrgCurrentMatchingClientsOfAWxTag(
    ctx context.Context,
    orgId uuid.UUID,
    wxtagId uuid.UUID) (
    models.ApiResponse[[]models.WxtagClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.WxtagClient`](../../doc/models/wxtag-client.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWxTags.GetOrgCurrentMatchingClientsOfAWxTag(ctx, orgId, wxtagId)
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


# Get Org Wx Tag

Get WxLAN Tag Details

```go
GetOrgWxTag(
    ctx context.Context,
    orgId uuid.UUID,
    wxtagId uuid.UUID) (
    models.ApiResponse[models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsWxTags.GetOrgWxTag(ctx, orgId, wxtagId)
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


# List Org Wx Tags

Get List of Org WxLAN Tags

```go
ListOrgWxTags(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsWxTags.ListOrgWxTags(ctx, orgId, &limit, &page)
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


# Update Org Wx Tag

Update WxLAN Tag

```go
UpdateOrgWxTag(
    ctx context.Context,
    orgId uuid.UUID,
    wxtagId uuid.UUID,
    body *models.WxlanTag) (
    models.ApiResponse[models.WxlanTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `wxtagId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.WxlanTag`](../../doc/models/wxlan-tag.md) | Body, Optional | Request Body |

## Response Type

[`models.WxlanTag`](../../doc/models/wxlan-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

wxtagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.WxlanTag{
    Name:         "name6",
    Op:           models.ToPointer(models.WxlanTagOperationEnum("in")),
    OrgId:        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    SiteId:       models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    Type:         models.WxlanTagTypeEnum("subnet"),
    VlanId:       models.ToPointer(models.WxlanTagVlanIdContainer.FromNumber(1055)),
}

apiResponse, err := orgsWxTags.UpdateOrgWxTag(ctx, orgId, wxtagId, &body)
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

