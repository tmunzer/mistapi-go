# Sites Stats-Discovered Switches

```go
sitesStatsDiscoveredSwitches := client.SitesStatsDiscoveredSwitches()
```

## Class Name

`SitesStatsDiscoveredSwitches`

## Methods

* [Count Site Discovered Switches](../../doc/controllers/sites-stats-discovered-switches.md#count-site-discovered-switches)
* [List Site Discovered Switches Metrics](../../doc/controllers/sites-stats-discovered-switches.md#list-site-discovered-switches-metrics)
* [Search Site Discovered Switches](../../doc/controllers/sites-stats-discovered-switches.md#search-site-discovered-switches)
* [Search Site Discovered Switches Metrics](../../doc/controllers/sites-stats-discovered-switches.md#search-site-discovered-switches-metrics)


# Count Site Discovered Switches

Count Discovered Switches

```go
CountSiteDiscoveredSwitches(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDiscoveredSwitchesCountDistinctEnum,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteDiscoveredSwitchesCountDistinctEnum`](../../doc/models/site-discovered-switches-count-distinct-enum.md) | Query, Optional | **Default**: `"system_name"` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteDiscoveredSwitchesCountDistinctEnum_SYSTEMNAME

duration := "10m"

limit := 100

apiResponse, err := sitesStatsDiscoveredSwitches.CountSiteDiscoveredSwitches(ctx, siteId, &distinct, nil, nil, &duration, &limit)
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


# List Site Discovered Switches Metrics

Discovered switches related metrics, lists related switch system names & details if not compliant

```go
ListSiteDiscoveredSwitchesMetrics(
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
| `threshold` | `*string` | Query, Optional | Configurable # ap per switch threshold, default 12 |
| `systemName` | `*string` | Query, Optional | System name for switch level metrics, optional |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDswitchesMetrics](../../doc/models/response-dswitches-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

threshold := "12"

systemName := "switch1.example.com"

apiResponse, err := sitesStatsDiscoveredSwitches.ListSiteDiscoveredSwitchesMetrics(ctx, siteId, &threshold, &systemName)
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
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
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
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDiscoveredSwitches](../../doc/models/response-discovered-switches.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

adopted := true

systemName := "switch1.example.com"

hostname := "switch1"

vendor := "Cisco"

model := "WS-C3850-24P"

version := "1.0.0"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesStatsDiscoveredSwitches.SearchSiteDiscoveredSwitches(ctx, siteId, &adopted, &systemName, &hostname, &vendor, &model, &version, &limit, nil, nil, &duration, &sort, nil)
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
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseDiscoveredSwitchMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`*models.DiscoveredSwitchesMetricScopeEnum`](../../doc/models/discovered-switches-metric-scope-enum.md) | Query, Optional | Metric scope<br><br>**Default**: `"site"` |
| `mType` | [`*models.DiscoveredSwitchMetricTypeEnum`](../../doc/models/discovered-switch-metric-type-enum.md) | Query, Optional | Metric type |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDiscoveredSwitchMetrics](../../doc/models/response-discovered-switch-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.DiscoveredSwitchesMetricScopeEnum_SITE

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesStatsDiscoveredSwitches.SearchSiteDiscoveredSwitchesMetrics(ctx, siteId, &scope, nil, &limit, nil, nil, &duration, &sort, nil)
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

