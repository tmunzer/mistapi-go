# Sites Stats-Apps

```go
sitesStatsApps := client.SitesStatsApps()
```

## Class Name

`SitesStatsApps`


# Count Site Apps

Count by Distinct Attributes of Applications

```go
CountSiteApps(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteAppsCountDistinctEnum,
    deviceMac *string,
    app *string,
    wired *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteAppsCountDistinctEnum`](../../doc/models/site-apps-count-distinct-enum.md) | Query, Optional | Default for wireless devices is `ap`. Default for wired devices is `device_mac` |
| `deviceMac` | `*string` | Query, Optional | MAC of the device |
| `app` | `*string` | Query, Optional | Application name |
| `wired` | `*string` | Query, Optional | If a device is wired or wireless. Default is False. |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



deviceMac := "001122334455"

app := "Example App"

wired := "False"

limit := 100

apiResponse, err := sitesStatsApps.CountSiteApps(ctx, siteId, nil, &deviceMac, &app, &wired, &limit)
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

