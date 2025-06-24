# Orgs Site Templates

```go
orgsSiteTemplates := client.OrgsSiteTemplates()
```

## Class Name

`OrgsSiteTemplates`

## Methods

* [Create Org Site Template](../../doc/controllers/orgs-site-templates.md#create-org-site-template)
* [Delete Org Site Template](../../doc/controllers/orgs-site-templates.md#delete-org-site-template)
* [Get Org Site Template](../../doc/controllers/orgs-site-templates.md#get-org-site-template)
* [List Org Site Templates](../../doc/controllers/orgs-site-templates.md#list-org-site-templates)
* [Update Org Site Template](../../doc/controllers/orgs-site-templates.md#update-org-site-template)


# Create Org Site Template

Create Org Site Template

```go
CreateOrgSiteTemplate(
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SiteTemplate](../../doc/models/site-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SiteTemplate{
    Vars:                 map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
}

apiResponse, err := orgsSiteTemplates.CreateOrgSiteTemplate(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SiteTemplate](../../doc/models/site-template.md).

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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Site Templates

Get List of Org Site Templates

```go
ListOrgSiteTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.SiteTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.SiteTemplate](../../doc/models/site-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSiteTemplates.ListOrgSiteTemplates(ctx, orgId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SiteTemplate](../../doc/models/site-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

sitetemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.SiteTemplate{
    Vars:                 map[string]string{
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

