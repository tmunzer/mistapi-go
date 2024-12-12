# Orgs SL Es

```go
orgsSLEs := client.OrgsSLEs()
```

## Class Name

`OrgsSLEs`

## Methods

* [Get Org Sites Sle](../../doc/controllers/orgs-sl-es.md#get-org-sites-sle)
* [Get Org Sle](../../doc/controllers/orgs-sl-es.md#get-org-sle)


# Get Org Sites Sle

Get Org Sites SLE

```go
GetOrgSitesSle(
    ctx context.Context,
    orgId uuid.UUID,
    sle *models.OrgSiteSleTypeEnum,
    start *int,
    end *int,
    duration *string,
    interval *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseOrgSiteSle],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `sle` | [`*models.OrgSiteSleTypeEnum`](../../doc/models/org-site-sle-type-enum.md) | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponseOrgSiteSle`](../../doc/models/containers/response-org-site-sle.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")







duration := "10m"

interval := "10m"

limit := 100

page := 1

apiResponse, err := orgsSLEs.GetOrgSitesSle(ctx, orgId, nil, nil, nil, &duration, &interval, &limit, &page)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsOrgSiteSleWifi(); ok {
        fmt.Println("Value narrowed down to models.OrgSiteSleWifi: ", *r)
    } else if r, ok := responseBody.AsOrgSiteWiredWifi(); ok {
        fmt.Println("Value narrowed down to models.OrgSiteWiredWifi: ", *r)
    } else if r, ok := responseBody.AsOrgSiteWanWifi(); ok {
        fmt.Println("Value narrowed down to models.OrgSiteWanWifi: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "end": 1651323600,
  "interval": 3600,
  "limit": 1,
  "page": 2,
  "results": [
    {
      "application_health": 0.8250000047942866,
      "gateway-health": 1,
      "num_clients": 65,
      "num_gateways": 1,
      "site_id": "f5fcbee5-1234-5678-9101-1619ede87879",
      "wan-link-health": 0.9988471089272484
    }
  ],
  "start": 1651269600,
  "total": 4
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


# Get Org Sle

Get Org SLEs (all/worst sites, Mx Edges, ...)

```go
GetOrgSle(
    ctx context.Context,
    orgId uuid.UUID,
    metric string,
    sle *string,
    duration *string,
    interval *string,
    start *int,
    end *int) (
    models.ApiResponse[models.InsightMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | see /api/v1/const/insight_metrics for available metrics |
| `sle` | `*string` | Query, Optional | see [listInsightMetrics](../../doc/controllers/constants-definitions.md#list-insight-metrics) for more details |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w<br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |

## Response Type

[`models.InsightMetrics`](../../doc/models/insight-metrics.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

metric := "metric8"



duration := "10m"

interval := "10m"





apiResponse, err := orgsSLEs.GetOrgSle(ctx, orgId, metric, nil, &duration, &interval, nil, nil)
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
  "end": 1643706000,
  "interval": 3600,
  "limit": 100,
  "results": [
    {
      "ap-availability": 0.9919400860511628,
      "ap-health": 0.967607512909879,
      "capacity": 0.7484652273070254,
      "coverage": 0.91217567374857,
      "num_aps": 13,
      "num_clients": 12,
      "roaming": 0.991735537682683,
      "roaming-exp": 0.991735537682683,
      "site_id": "978c48e6-6ef6-11e6-8bbf-02e208b2d34f",
      "successful-connect": 0.46052632135780236,
      "throughput": 0.6775702123846302,
      "time-to-connect": 0.9349112447196916
    },
    {
      "ap-availability": 0.9990384613092129,
      "ap-health": 0.48201754375507955,
      "capacity": 0.9702673450306101,
      "coverage": 0.8335392334930375,
      "num_aps": 1,
      "num_clients": 6,
      "roaming": 1,
      "roaming-exp": 1,
      "site_id": "49ff76e0-a283-4e7d-b38d-041f1e9aff3c",
      "successful-connect": 1,
      "throughput": 0,
      "time-to-connect": 1
    },
    {
      "ap-availability": 1,
      "ap-health": 0.982456140612301,
      "capacity": 1,
      "coverage": 0.9276041182442488,
      "num_aps": 2,
      "num_clients": 3,
      "roaming": 1,
      "roaming-exp": 1,
      "site_id": "9b9b48f1-15a4-459e-86cc-9cbec9005983",
      "successful-connect": 1,
      "throughput": 1,
      "time-to-connect": 0.8125
    },
    {
      "ap-availability": 0.9981132070973234,
      "ap-health": 0.9991228068084047,
      "capacity": 1,
      "coverage": 1,
      "num_aps": 1,
      "num_clients": 0,
      "roaming": 1,
      "roaming-exp": 1,
      "site_id": "eb0e1671-7a6b-472b-94c3-c187dafe5274",
      "successful-connect": 1,
      "throughput": 0,
      "time-to-connect": 0.5
    }
  ],
  "start": 1643670000
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

