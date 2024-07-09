# Orgs Site Templates

```go
orgsSiteTemplates := client.OrgsSiteTemplates()
```

## Class Name

`OrgsSiteTemplates`

## Methods

* [Create Org Site Templates](../../doc/controllers/orgs-site-templates.md#create-org-site-templates)
* [Delete Org Site Template](../../doc/controllers/orgs-site-templates.md#delete-org-site-template)
* [Get Org Site Template](../../doc/controllers/orgs-site-templates.md#get-org-site-template)
* [List Org Site Templates](../../doc/controllers/orgs-site-templates.md#list-org-site-templates)
* [Update Org Site Template](../../doc/controllers/orgs-site-templates.md#update-org-site-template)


# Create Org Site Templates

Create Org Site Template

```go
CreateOrgSiteTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.SiteTemplate) (
    models.ApiResponse[models.SiteTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SiteTemplate`](../../doc/models/site-template.md) | Body, Optional | - |

## Response Type

[`models.SiteTemplate`](../../doc/models/site-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SiteTemplate{
    Vars:        map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
}

apiResponse, err := orgsSiteTemplates.CreateOrgSiteTemplates(ctx, orgId, &body)
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
  "auto_upgrade": {
    "day_of_week": "mon",
    "enabled": true,
    "time_of_day": "string",
    "version": "string"
  },
  "name": "string",
  "vars": {
    "SSID_STR": "string",
    "VLAN_ID": "string"
  }
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


# Delete Org Site Template

Delete Org Site Template

```go
DeleteOrgSiteTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sitetemplateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sitetemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sitetemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSiteTemplates.DeleteOrgSiteTemplate(ctx, orgId, sitetemplateId)
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


# Get Org Site Template

Get Org Site Template

```go
GetOrgSiteTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sitetemplateId uuid.UUID) (
    models.ApiResponse[models.SiteTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sitetemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.SiteTemplate`](../../doc/models/site-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sitetemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSiteTemplates.GetOrgSiteTemplate(ctx, orgId, sitetemplateId)
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
  "auto_upgrade": {
    "day_of_week": "mon",
    "enabled": true,
    "time_of_day": "string",
    "version": "string"
  },
  "name": "string",
  "vars": {
    "SSID_STR": "string",
    "VLAN_ID": "string"
  }
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


# List Org Site Templates

Get List of Org Site Templates

```go
ListOrgSiteTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    page *int,
    limit *int) (
    models.ApiResponse[[]models.SiteTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |

## Response Type

[`[]models.SiteTemplate`](../../doc/models/site-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

page := 1

limit := 100

apiResponse, err := orgsSiteTemplates.ListOrgSiteTemplates(ctx, orgId, &page, &limit)
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
    "auto_upgrade": {
      "day_of_week": "mon",
      "enabled": true,
      "time_of_day": "string",
      "version": "string"
    },
    "name": "string",
    "vars": {
      "SSID_STR": "string",
      "VLAN_ID": "string"
    }
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


# Update Org Site Template

Update Org Site Template

```go
UpdateOrgSiteTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    sitetemplateId uuid.UUID,
    body *models.SiteTemplate) (
    models.ApiResponse[models.SiteTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sitetemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.SiteTemplate`](../../doc/models/site-template.md) | Body, Optional | - |

## Response Type

[`models.SiteTemplate`](../../doc/models/site-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sitetemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SiteTemplate{
    Vars:        map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
}

apiResponse, err := orgsSiteTemplates.UpdateOrgSiteTemplate(ctx, orgId, sitetemplateId, &body)
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
  "auto_upgrade": {
    "day_of_week": "mon",
    "enabled": true,
    "time_of_day": "string",
    "version": "string"
  },
  "name": "string",
  "vars": {
    "SSID_STR": "string",
    "VLAN_ID": "string"
  }
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
