# Sites Licenses

```go
sitesLicenses := client.SitesLicenses()
```

## Class Name

`SitesLicenses`


# Get Site License Usage

This shows license usage (i.e. needed) based on the features enabled for the site.

```go
GetSiteLicenseUsage(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.LicenseUsageSite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.LicenseUsageSite`](../../doc/models/license-usage-site.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesLicenses.GetSiteLicenseUsage(ctx, siteId)
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
  "org_entitled": {
    "SUB-LOC": 30,
    "SUB-MAN": 60
  },
  "svna_enabled": true,
  "trial_enabled": true,
  "usages": {
    "SUB-LOC": 30,
    "SUB-MAN": 60
  },
  "vna_eligible": true,
  "vna_ui": true,
  "wvna_eligible": true
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

