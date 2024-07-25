# Sites Stats-Wx Rules

```go
sitesStatsWxRules := client.SitesStatsWxRules()
```

## Class Name

`SitesStatsWxRules`


# Get Site Wx Rules Usage

Get Wxlan Rule usage

```go
GetSiteWxRulesUsage(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]models.WxruleStat],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.WxruleStat`](../../doc/models/wxrule-stat.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesStatsWxRules.GetSiteWxRulesUsage(ctx, siteId)
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
    "action": "allow",
    "client_mac": [
      "3bbbf819bb6f",
      "bd96cbc4910f"
    ],
    "dst_allow_wxtags": [
      "fff34466-eec0-3756-6765-381c728a6037",
      "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
    ],
    "dst_deny_wxtags": [
      "aaa34466-eec0-3756-6765-381c728a6037",
      "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
    ],
    "dst_wxtags": [
      "d4134466-eec0-3756-6765-381c728a6037",
      "1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3"
    ],
    "name": "Guest",
    "order": 1,
    "src_wxtags": [
      "8bfc2490-d726-3587-038d-cb2e71bd2330",
      "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
    ],
    "usage": {
      "1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3": {
        "num_flows": 60
      },
      "d4134466-eec0-3756-6765-381c728a6037": {
        "num_flows": 60
      }
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

