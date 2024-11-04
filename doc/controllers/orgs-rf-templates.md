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

Create Org RF Template

```go
CreateOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.RfTemplate) (
    models.ApiResponse[models.RfTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RfTemplate`](../../doc/models/rf-template.md) | Body, Optional | Request Body |

## Response Type

[`models.RfTemplate`](../../doc/models/rf-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RfTemplate{
    Name:            "name6",
}

apiResponse, err := orgsRFTemplates.CreateOrgRfTemplate(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Delete Org Rf Template

Delete Org RF Template

```go
DeleteOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `rftemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rftemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsRFTemplates.DeleteOrgRfTemplate(ctx, orgId, rftemplateId)
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


# Get Org Rf Template

Get Org RF Template Details

```go
GetOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID) (
    models.ApiResponse[models.RfTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `rftemplateId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.RfTemplate`](../../doc/models/rf-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rftemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsRFTemplates.GetOrgRfTemplate(ctx, orgId, rftemplateId)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Rf Templates

Get List of Org RF Template

```go
ListOrgRfTemplates(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.RfTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.RfTemplate`](../../doc/models/rf-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsRFTemplates.ListOrgRfTemplates(ctx, orgId, &limit, &page)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Update Org Rf Template

Update Org RF Template

```go
UpdateOrgRfTemplate(
    ctx context.Context,
    orgId uuid.UUID,
    rftemplateId uuid.UUID,
    body *models.RfTemplate) (
    models.ApiResponse[models.RfTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `rftemplateId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.RfTemplate`](../../doc/models/rf-template.md) | Body, Optional | Request Body |

## Response Type

[`models.RfTemplate`](../../doc/models/rf-template.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

rftemplateId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.RfTemplate{
    Name:            "name6",
}

apiResponse, err := orgsRFTemplates.UpdateOrgRfTemplate(ctx, orgId, rftemplateId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

