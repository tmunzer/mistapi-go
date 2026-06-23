# Orgs RF Templates

```go
orgsRFTemplates := client.OrgsRFTemplates()
```

## Class Name

`OrgsRFTemplates`

## Methods

* [Create Org Rf Template](../../doc/controllers/orgs-rf-templates.md#create-org-rf-template)
* [Delete Org Rf Template](../../doc/controllers/orgs-rf-templates.md#delete-org-rf-template)
* [Get Org Rf Template](../../doc/controllers/orgs-rf-templates.md#get-org-rf-template)
* [List Org Rf Templates](../../doc/controllers/orgs-rf-templates.md#list-org-rf-templates)
* [Update Org Rf Template](../../doc/controllers/orgs-rf-templates.md#update-org-rf-template)


# Create Org Rf Template

Create an organization RF template with 2.4, 5, and 6 GHz radio
settings, country code, scanning behavior, antenna gain, and model-specific
overrides.

To assign a RF template to a site, use the [Update Site](../../doc/controllers/sites.md#update-site-info) endpoint and specify the RF template ID in the `rftemplate_id` field of the request body.

```go
CreateOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.RfTemplate) (
    models.ApiResponse[models.RfTemplate],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RfTemplate`](../../doc/models/rf-template.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.RfTemplate](../../doc/models/rf-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RfTemplate{
    Name:                 "name6",
}

apiResponse, err := orgsRFTemplates.CreateOrgRfTemplate(ctx, orgId, &body)
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
  "ant_gain_24": 0,
  "ant_gain_5": 0,
  "band_24": {
    "allow_rrm_disable": true,
    "ant_gain": 0,
    "bandwidth": 20,
    "channels": [
      1,
      6,
      11
    ],
    "disabled": false,
    "power_max": 11,
    "power_min": 3,
    "preamble": "short"
  },
  "band_24_usage": "auto",
  "band_5": {
    "allow_rrm_disable": false,
    "ant_gain": 0,
    "bandwidth": 80,
    "channels": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      149,
      153,
      157,
      161
    ],
    "disabled": false,
    "power_max": 16,
    "power_min": 9,
    "preamble": "short"
  },
  "country_code": "FR",
  "created_time": 1594743723,
  "id": "b3f20330-f76a-49f1-bc65-0d8727140b1d",
  "model_specific": {},
  "modified_time": 1613582192,
  "name": "Lab",
  "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4"
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


# Delete Org Rf Template

Delete an organization RF template by template ID so it can no longer be applied to sites.

```go
DeleteOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `rftemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rftemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsRFTemplates.DeleteOrgRfTemplate(ctx, orgId, rftemplateId)
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
    fmt.Println(resp.StatusCode)
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


# Get Org Rf Template

Retrieve details for a specific organization RF template, including radio settings, country code, scanning behavior, antenna gain, and model-specific overrides.

```go
GetOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID) (
    models.ApiResponse[models.RfTemplate],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `rftemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.RfTemplate](../../doc/models/rf-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rftemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsRFTemplates.GetOrgRfTemplate(ctx, orgId, rftemplateId)
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
  "ant_gain_24": 0,
  "ant_gain_5": 0,
  "band_24": {
    "allow_rrm_disable": true,
    "ant_gain": 0,
    "bandwidth": 20,
    "channels": [
      1,
      6,
      11
    ],
    "disabled": false,
    "power_max": 11,
    "power_min": 3,
    "preamble": "short"
  },
  "band_24_usage": "auto",
  "band_5": {
    "allow_rrm_disable": false,
    "ant_gain": 0,
    "bandwidth": 80,
    "channels": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      149,
      153,
      157,
      161
    ],
    "disabled": false,
    "power_max": 16,
    "power_min": 9,
    "preamble": "short"
  },
  "country_code": "FR",
  "created_time": 1594743723,
  "id": "b3f20330-f76a-49f1-bc65-0d8727140b1d",
  "model_specific": {},
  "modified_time": 1613582192,
  "name": "Lab",
  "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4"
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


# List Org Rf Templates

List organization RF templates used by RRM to apply radio settings across sites.

```go
ListOrgRfTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.RfTemplate],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.RfTemplate](../../doc/models/rf-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsRFTemplates.ListOrgRfTemplates(ctx, orgId, &limit, &page)
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
[
  {
    "ant_gain_24": 0,
    "ant_gain_5": 0,
    "band_24": {
      "allow_rrm_disable": true,
      "ant_gain": 0,
      "bandwidth": 20,
      "channels": [
        1,
        6,
        11
      ],
      "disabled": false,
      "power_max": 11,
      "power_min": 3,
      "preamble": "short"
    },
    "band_24_usage": "auto",
    "band_5": {
      "allow_rrm_disable": false,
      "ant_gain": 0,
      "bandwidth": 80,
      "channels": [
        36,
        40,
        44,
        48,
        52,
        56,
        60,
        64,
        100,
        104,
        149,
        153,
        157,
        161
      ],
      "disabled": false,
      "power_max": 16,
      "power_min": 9,
      "preamble": "short"
    },
    "country_code": "FR",
    "created_time": 1594743723,
    "id": "b3f20330-f76a-49f1-bc65-0d8727140b1d",
    "model_specific": {},
    "modified_time": 1613582192,
    "name": "Lab",
    "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Update Org Rf Template

Update an organization RF template, including radio settings, country code, scanning behavior, antenna gain, and model-specific overrides.

```go
UpdateOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID,
    body *models.RfTemplate) (
    models.ApiResponse[models.RfTemplate],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `rftemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RfTemplate`](../../doc/models/rf-template.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.RfTemplate](../../doc/models/rf-template.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rftemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RfTemplate{
    Name:                 "name6",
}

apiResponse, err := orgsRFTemplates.UpdateOrgRfTemplate(ctx, orgId, rftemplateId, &body)
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
  "ant_gain_24": 0,
  "ant_gain_5": 0,
  "band_24": {
    "allow_rrm_disable": true,
    "ant_gain": 0,
    "bandwidth": 20,
    "channels": [
      1,
      6,
      11
    ],
    "disabled": false,
    "power_max": 11,
    "power_min": 3,
    "preamble": "short"
  },
  "band_24_usage": "auto",
  "band_5": {
    "allow_rrm_disable": false,
    "ant_gain": 0,
    "bandwidth": 80,
    "channels": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      149,
      153,
      157,
      161
    ],
    "disabled": false,
    "power_max": 16,
    "power_min": 9,
    "preamble": "short"
  },
  "country_code": "FR",
  "created_time": 1594743723,
  "id": "b3f20330-f76a-49f1-bc65-0d8727140b1d",
  "model_specific": {},
  "modified_time": 1613582192,
  "name": "Lab",
  "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4"
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

