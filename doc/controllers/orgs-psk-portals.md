# Orgs Psk Portals

```go
orgsPskPortals := client.OrgsPskPortals()
```

## Class Name

`OrgsPskPortals`

## Methods

* [Count Org Psk Portal Logs](../../doc/controllers/orgs-psk-portals.md#count-org-psk-portal-logs)
* [Create Org Psk Portal](../../doc/controllers/orgs-psk-portals.md#create-org-psk-portal)
* [Delete Org Psk Portal](../../doc/controllers/orgs-psk-portals.md#delete-org-psk-portal)
* [Delete Org Psk Portal Image](../../doc/controllers/orgs-psk-portals.md#delete-org-psk-portal-image)
* [Get Org Psk Portal](../../doc/controllers/orgs-psk-portals.md#get-org-psk-portal)
* [List Org Psk Portal Logs](../../doc/controllers/orgs-psk-portals.md#list-org-psk-portal-logs)
* [List Org Psk Portals](../../doc/controllers/orgs-psk-portals.md#list-org-psk-portals)
* [Search Org Psk Portal Logs](../../doc/controllers/orgs-psk-portals.md#search-org-psk-portal-logs)
* [Update Org Psk Portal](../../doc/controllers/orgs-psk-portals.md#update-org-psk-portal)
* [Update Org Psk Portal Template](../../doc/controllers/orgs-psk-portals.md#update-org-psk-portal-template)
* [Upload Org Psk Portal Image](../../doc/controllers/orgs-psk-portals.md#upload-org-psk-portal-image)


# Count Org Psk Portal Logs

Count by Distinct Attributes of PskPortal Logs

```go
CountOrgPskPortalLogs(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgPskPortalLogsCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgPskPortalLogsCountDistinctEnum`](../../doc/models/org-psk-portal-logs-count-distinct-enum.md) | Query, Optional | **Default**: `"pskportal_id"` |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgPskPortalLogsCountDistinctEnum("pskportal_id")





duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsPskPortals.CountOrgPskPortalLogs(ctx, orgId, &distinct, nil, nil, &duration, &limit, &page)
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


# Create Org Psk Portal

Create Org Psk Portal

```go
CreateOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PskPortal) (
    models.ApiResponse[models.PskPortal],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PskPortal`](../../doc/models/psk-portal.md) | Body, Optional | - |

## Response Type

[`models.PskPortal`](../../doc/models/psk-portal.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.PskPortal{
    Auth:                         models.ToPointer(models.PskPortalAuthEnum("sso")),
    MaxUsage:                     models.ToPointer(0),
    Name:                         "string",
    RequiredFields:               []string{
        "string",
    },
    Role:                         models.ToPointer("string"),
    Ssid:                         "string",
    Sso:                          models.ToPointer(models.PskPortalSso{
        IdpCert:              models.ToPointer("string"),
        IdpSignAlgo:          models.ToPointer("string"),
        IdpSsoUrl:            models.ToPointer("string"),
        Issuer:               models.ToPointer("string"),
        NameidFormat:         models.ToPointer("string"),
    }),
    TemplateUrl:                  models.ToPointer("string"),
    Type:                         models.ToPointer(models.PskPortalTypeEnum("byod")),
    VlanId:                       models.ToPointer(models.VlanIdWithVariableContainer.FromNumber(10)),
}

apiResponse, err := orgsPskPortals.CreateOrgPskPortal(ctx, orgId, &body)
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


# Delete Org Psk Portal

Delete Org Psk Portal

```go
DeleteOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskportalId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskportalId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsPskPortals.DeleteOrgPskPortal(ctx, orgId, pskportalId)
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


# Delete Org Psk Portal Image

Delete background image for PskPortal

If image is not uploaded or is deleted, PskPortal will use default image.

```go
DeleteOrgPskPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskportalId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskportalId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsPskPortals.DeleteOrgPskPortalImage(ctx, orgId, pskportalId)
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


# Get Org Psk Portal

get Org Psk Portal Details

```go
GetOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID) (
    models.ApiResponse[models.PskPortal],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskportalId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.PskPortal`](../../doc/models/psk-portal.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskportalId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsPskPortals.GetOrgPskPortal(ctx, orgId, pskportalId)
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


# List Org Psk Portal Logs

Get the list of PSK Portals Logs

```go
ListOrgPskPortalLogs(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponsePskPortalLogsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponsePskPortalLogsSearch`](../../doc/models/response-psk-portal-logs-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsPskPortals.ListOrgPskPortalLogs(ctx, orgId, nil, nil, &duration, &limit, &page)
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
      "id": "8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09",
      "message": "Rotate PSK test@mist.com",
      "name_id": "test@mist.com",
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
      "psk_id": "608fe603-f9f0-4ce9-9473-04ef6c6ea749",
      "psk_name": "test@mist.com",
      "pskportal_id": "c1742c09-af35-4161-96ef-7dc65c6d5674",
      "timestamp": 1686346104.096
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


# List Org Psk Portals

Get List of Org Psk Portals

```go
ListOrgPskPortals(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.PskPortal],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.PskPortal`](../../doc/models/psk-portal.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsPskPortals.ListOrgPskPortals(ctx, orgId, &limit, &page)
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


# Search Org Psk Portal Logs

Search Org PSK Portal Logs

```go
SearchOrgPskPortalLogs(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int,
    pskName *string,
    pskId *string,
    pskportalId *string,
    id *uuid.UUID,
    adminName *string,
    adminId *string,
    nameId *uuid.UUID) (
    models.ApiResponse[models.ResponsePskPortalLogsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |
| `pskName` | `*string` | Query, Optional | - |
| `pskId` | `*string` | Query, Optional | - |
| `pskportalId` | `*string` | Query, Optional | - |
| `id` | `*uuid.UUID` | Query, Optional | audit_id |
| `adminName` | `*string` | Query, Optional | - |
| `adminId` | `*string` | Query, Optional | - |
| `nameId` | `*uuid.UUID` | Query, Optional | name_id used in SSO |

## Response Type

[`models.ResponsePskPortalLogsSearch`](../../doc/models/response-psk-portal-logs-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

page := 1















apiResponse, err := orgsPskPortals.SearchOrgPskPortalLogs(ctx, orgId, nil, nil, &duration, &limit, &page, nil, nil, nil, nil, nil, nil, nil)
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
      "id": "8a3dcaa7-80e3-4bb0-a75b-7bc6322cfd09",
      "message": "Rotate PSK test@mist.com",
      "name_id": "test@mist.com",
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
      "psk_id": "608fe603-f9f0-4ce9-9473-04ef6c6ea749",
      "psk_name": "test@mist.com",
      "pskportal_id": "c1742c09-af35-4161-96ef-7dc65c6d5674",
      "timestamp": 1686346104.096
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


# Update Org Psk Portal

update Org Psk Portal

```go
UpdateOrgPskPortal(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID,
    body *models.PskPortal) (
    models.ApiResponse[models.PskPortal],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskportalId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PskPortal`](../../doc/models/psk-portal.md) | Body, Optional | - |

## Response Type

[`models.PskPortal`](../../doc/models/psk-portal.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskportalId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.PskPortal{
    Auth:                         models.ToPointer(models.PskPortalAuthEnum("sso")),
    MaxUsage:                     models.ToPointer(0),
    Name:                         "string",
    RequiredFields:               []string{
        "string",
    },
    Role:                         models.ToPointer("string"),
    Ssid:                         "string",
    Sso:                          models.ToPointer(models.PskPortalSso{
        IdpCert:              models.ToPointer("string"),
        IdpSignAlgo:          models.ToPointer("string"),
        IdpSsoUrl:            models.ToPointer("string"),
        Issuer:               models.ToPointer("string"),
        NameidFormat:         models.ToPointer("string"),
    }),
    TemplateUrl:                  models.ToPointer("string"),
    Type:                         models.ToPointer(models.PskPortalTypeEnum("byod")),
    VlanId:                       models.ToPointer(models.VlanIdWithVariableContainer.FromNumber(10)),
}

apiResponse, err := orgsPskPortals.UpdateOrgPskPortal(ctx, orgId, pskportalId, &body)
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


# Update Org Psk Portal Template

Update Org Psk Portal Template

```go
UpdateOrgPskPortalTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID,
    body *models.PskPortalTemplate) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskportalId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PskPortalTemplate`](../../doc/models/psk-portal-template.md) | Body, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskportalId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.PskPortalTemplate{
    Alignment: models.ToPointer(models.PortalTemplateAlignmentEnum("center")),
    Color:     models.ToPointer("#1074bc"),
    PoweredBy: models.ToPointer(false),
}

resp, err := orgsPskPortals.UpdateOrgPskPortalTemplate(ctx, orgId, pskportalId, &body)
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


# Upload Org Psk Portal Image

Upload background image for PskPortal

```go
UploadOrgPskPortalImage(
    ctx context.Context,
    orgId uuid.UUID,
    pskportalId uuid.UUID,
    file *models.FileWrapper,
    json *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `pskportalId` | `uuid.UUID` | Template, Required | - |
| `file` | `*models.FileWrapper` | Form, Optional | Binary file |
| `json` | `*string` | Form, Optional | JSON string describing the upload |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

pskportalId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





resp, err := orgsPskPortals.UploadOrgPskPortalImage(ctx, orgId, pskportalId, nil, nil)
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

