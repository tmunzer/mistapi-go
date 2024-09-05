# Sites Stats-Discovered Switches

```go
sitesStatsDiscoveredSwitches := client.SitesStatsDiscoveredSwitches()
```

## Class Name

`SitesStatsDiscoveredSwitches`

## Methods

* [Count Site Discovered Switches](../../doc/controllers/sites-stats-discovered-switches.md#count-site-discovered-switches)
* [Get Site Discovered Switches Metrics](../../doc/controllers/sites-stats-discovered-switches.md#get-site-discovered-switches-metrics)
* [Search Site Discovered Switches](../../doc/controllers/sites-stats-discovered-switches.md#search-site-discovered-switches)
* [Search Site Discovered Switches Metrics](../../doc/controllers/sites-stats-discovered-switches.md#search-site-discovered-switches-metrics)


# Count Site Discovered Switches

Count Discovered Switches

```go
CountSiteDiscoveredSwitches(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDiscoveredSwitchesCountDistinctEnum,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDiscoveredSwitchesCountDistinctEnum`](../../doc/models/site-discovered-switches-count-distinct-enum.md) | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDiscoveredSwitchesCountDistinctEnum("system_name")





duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesStatsDiscoveredSwitches.CountSiteDiscoveredSwitches(ctx, siteId, &distinct, nil, nil, &duration, &limit, &page)
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


# Get Site Discovered Switches Metrics

Discovered switches related metrics, lists related switch system names & details if not compliant

```go
GetSiteDiscoveredSwitchesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    threshold *string,
    systemName *string) (
    models.ApiResponse[models.ResponseDswitchesMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `threshold` | `*string` | Query, Optional | configurable # ap per switch threshold, default 12 |
| `systemName` | `*string` | Query, Optional | system name for switch level metrics, optional |

## Response Type

[`models.ResponseDswitchesMetrics`](../../doc/models/response-dswitches-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





apiResponse, err := sitesStatsDiscoveredSwitches.GetSiteDiscoveredSwitchesMetrics(ctx, siteId, nil, nil)
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
  "inactive_wired_vlans": {
    "details": {},
    "score": 100
  },
  "poe_compliance": {
    "details": {
      "total_aps": 63,
      "total_power": 981500
    },
    "score": 100
  },
  "switch_ap_affinity": {
    "details": {
      "system_name": [
        "mist-lab-ex2300c",
        "switch1"
      ],
      "threshold": 12
    },
    "score": 33.3333
  },
  "version_compliance": {
    "details": {
      "major_versions": [
        {
          "major_count": 2,
          "model": "EX2300-C-12P",
          "system_names": [
            "switch1",
            "mist-lab-ex2300c"
          ]
        },
        {
          "major_count": 1,
          "model": "EX4300-48P",
          "system_names": []
        }
      ],
      "total_switch_count": 5
    },
    "score": 75
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


# Search Site Discovered Switches

Search Discovered Switches

```go
SearchSiteDiscoveredSwitches(
    ctx context.Context,
    siteId uuid.UUID,
    adopted *bool,
    systemName *string,
    hostname *string,
    vendor *string,
    model *string,
    version *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDiscoveredSwitches],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `adopted` | `*bool` | Query, Optional | - |
| `systemName` | `*string` | Query, Optional | - |
| `hostname` | `*string` | Query, Optional | - |
| `vendor` | `*string` | Query, Optional | - |
| `model` | `*string` | Query, Optional | - |
| `version` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseDiscoveredSwitches`](../../doc/models/response-discovered-switches.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")













limit := 100





duration := "10m"

apiResponse, err := sitesStatsDiscoveredSwitches.SearchSiteDiscoveredSwitches(ctx, siteId, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1604496474.8978634,
  "limit": 1000,
  "results": [
    {
      "aps": [
        {
          "hostname": "ap41nearlab",
          "inactive_wired_vlans": [],
          "mac": "5c5b352e2001",
          "poe_status": true,
          "when": "2019-06-13T19:53:16.870+0000"
        }
      ],
      "mgmt_addr": "10.1.1.1",
      "model": "EX2300-C-12P",
      "org_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
      "site_id": "67970e46-4e12-11e6-9188-0242ac110007",
      "system_desc": "Juniper Networks, Inc. ex2300-c-12p Ethernet Switch, kernel JUNOS 18.2R2.6, Build date: 2018-12-07 13:19:04 UTC Copyright (c) 1996-2018 Juniper Networks, Inc.",
      "system_name": "mist-lab-ex2300c",
      "timestamp": 1560457177.037,
      "vendor": "Juniper Networks",
      "version": "18.2R2.6"
    }
  ],
  "start": 1604410074.8978484,
  "total": 1
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


# Search Site Discovered Switches Metrics

Search Discovered Switch Metrics

```go
SearchSiteDiscoveredSwitchesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    scope *models.DiscoveredSwitchesMetricScopeEnum,
    mType *models.DiscoveredSwitchMetricTypeEnum,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseDiscoveredSwitchMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`*models.DiscoveredSwitchesMetricScopeEnum`](../../doc/models/discovered-switches-metric-scope-enum.md) | Query, Optional | metric scope |
| `mType` | [`*models.DiscoveredSwitchMetricTypeEnum`](../../doc/models/discovered-switch-metric-type-enum.md) | Query, Optional | metric type |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.ResponseDiscoveredSwitchMetrics`](../../doc/models/response-discovered-switch-metrics.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.DiscoveredSwitchesMetricScopeEnum("site")



limit := 100





duration := "10m"

apiResponse, err := sitesStatsDiscoveredSwitches.SearchSiteDiscoveredSwitchesMetrics(ctx, siteId, &scope, nil, &limit, nil, nil, &duration)
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
  "end": 1675193686.0191767,
  "limit": 1,
  "next": "/api/v1/sites/f5fcbee5-fbca-45b3-8bf1-1619ede87879/stats/discovered_switch_metrics/search?end=1675193686.0191767&limit=1&search_after=%5B1675193400000%5D&start=1675107286.0191767",
  "results": [
    {
      "details": {},
      "org_id": "203d3d02-dbc0-4c1b-9f41-76896a3330f4",
      "scope": "site",
      "score": 100,
      "site_id": "f5fcbee5-fbca-45b3-8bf1-1619ede87879",
      "timestamp": 1675193400,
      "type": "inactive_wired_vlans"
    }
  ],
  "start": 1675107286.0191767,
  "total": 3
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

