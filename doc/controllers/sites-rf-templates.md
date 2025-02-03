# Sites RF Templates

```go
sitesRFTemplates := client.SitesRFTemplates()
```

## Class Name

`SitesRFTemplates`


# List Site Rf Template Derived

Get derived RF Templates for Site

```go
ListSiteRfTemplateDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.RfTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | Whether resolve the site variables |

## Response Type

[`[]models.RfTemplate`](../../doc/models/rf-template.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := sitesRFTemplates.ListSiteRfTemplateDerived(ctx, siteId, nil)
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

