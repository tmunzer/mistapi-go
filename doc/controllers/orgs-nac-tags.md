# Orgs NAC Tags

```go
orgsNACTags := client.OrgsNACTags()
```

## Class Name

`OrgsNACTags`

## Methods

* [Create Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#create-org-nac-tag)
* [Delete Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#delete-org-nac-tag)
* [Get Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#get-org-nac-tag)
* [List Org Nac Tags](../../doc/controllers/orgs-nac-tags.md#list-org-nac-tags)
* [Update Org Nac Tag](../../doc/controllers/orgs-nac-tags.md#update-org-nac-tag)


# Create Org Nac Tag

Create Org NAC Tag

```go
CreateOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.NacTag) (
    models.ApiResponse[models.NacTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NacTag`](../../doc/models/nac-tag.md) | Body, Optional | - |

## Response Type

[`models.NacTag`](../../doc/models/nac-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NacTag{
    AllowUsermacOverride: models.ToPointer(false),
    EgressVlanNames:      []string{
        "1vlan-30",
        "1vlan-20",
        "2-vlan10",
    },
    MatchAll:             models.ToPointer(false),
    Name:                 "name6",
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RadiusAttrs:          []string{
        "Idle-Timeout=600",
        "Termination-Action=RADIUS-Request",
    },
    RadiusVendorAttrs:    []string{
        "PaloAlto-Admin-Role=superuser",
        "PaloAlto-Panorama-Admin-Role=administrator",
    },
    SessionTimeout:       models.ToPointer(86000),
    Type:                 models.NacTagTypeEnum("radius_vendor_attrs"),
}

apiResponse, err := orgsNACTags.CreateOrgNacTag(ctx, orgId, &body)
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
  "match": "client_mac",
  "name": "cameras",
  "type": "match",
  "values": [
    "010203040506",
    "abcdef*"
  ]
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


# Delete Org Nac Tag

Delete Org NAC Tag

```go
DeleteOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nactagId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nactagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsNACTags.DeleteOrgNacTag(ctx, orgId, nactagId)
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


# Get Org Nac Tag

Get Org NAC Tag

```go
GetOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID) (
    models.ApiResponse[models.NacTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nactagId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.NacTag`](../../doc/models/nac-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nactagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsNACTags.GetOrgNacTag(ctx, orgId, nactagId)
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
  "match": "client_mac",
  "name": "cameras",
  "type": "match",
  "values": [
    "010203040506",
    "abcdef*"
  ]
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


# List Org Nac Tags

Get List of Org NAC Tags

```go
ListOrgNacTags(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    name *string,
    match *string,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.NacTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Type of NAC Tag |
| `name` | `*string` | Query, Optional | Name of NAC Tag |
| `match` | `*string` | Query, Optional | Type of NAC Tag |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.NacTag`](../../doc/models/nac-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







page := 1

limit := 100

apiResponse, err := orgsNACTags.ListOrgNacTags(ctx, orgId, nil, nil, nil, &page, &limit)
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


# Update Org Nac Tag

Update Org NAC Tag

```go
UpdateOrgNacTag(
    ctx context.Context,
    orgId uuid.UUID,
    nactagId uuid.UUID,
    body *models.NacTag) (
    models.ApiResponse[models.NacTag],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `nactagId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.NacTag`](../../doc/models/nac-tag.md) | Body, Optional | - |

## Response Type

[`models.NacTag`](../../doc/models/nac-tag.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

nactagId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.NacTag{
    AllowUsermacOverride: models.ToPointer(false),
    EgressVlanNames:      []string{
        "1vlan-30",
        "1vlan-20",
        "2-vlan10",
    },
    MatchAll:             models.ToPointer(false),
    Name:                 "name6",
    OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    RadiusAttrs:          []string{
        "Idle-Timeout=600",
        "Termination-Action=RADIUS-Request",
    },
    RadiusVendorAttrs:    []string{
        "PaloAlto-Admin-Role=superuser",
        "PaloAlto-Panorama-Admin-Role=administrator",
    },
    SessionTimeout:       models.ToPointer(86000),
    Type:                 models.NacTagTypeEnum("radius_vendor_attrs"),
}

apiResponse, err := orgsNACTags.UpdateOrgNacTag(ctx, orgId, nactagId, &body)
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
  "match": "client_mac",
  "name": "cameras",
  "type": "match",
  "values": [
    "010203040506",
    "abcdef*"
  ]
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
