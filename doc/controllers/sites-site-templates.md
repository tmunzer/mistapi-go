# Sites Site Templates

```go
sitesSiteTemplates := client.SitesSiteTemplates()
```

## Class Name

`SitesSiteTemplates`


# Get Site Site Template Derived

Get derived Site Templates for Site

```go
GetSiteSiteTemplateDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[models.SiteTemplate],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | whether resolve the site variables |

## Response Type

[`models.SiteTemplate`](../../doc/models/site-template.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



apiResponse, err := sitesSiteTemplates.GetSiteSiteTemplateDerived(ctx, siteId, nil)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

