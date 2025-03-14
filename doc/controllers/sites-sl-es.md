# Sites SL Es

```go
sitesSLEs := client.SitesSLEs()
```

## Class Name

`SitesSLEs`

## Methods

* [Get Site Sle Classifier Details](../../doc/controllers/sites-sl-es.md#get-site-sle-classifier-details)
* [Get Site Sle Histogram](../../doc/controllers/sites-sl-es.md#get-site-sle-histogram)
* [Get Site Sle Impact Summary](../../doc/controllers/sites-sl-es.md#get-site-sle-impact-summary)
* [Get Site Sle Summary](../../doc/controllers/sites-sl-es.md#get-site-sle-summary)
* [Get Site Sle Threshold](../../doc/controllers/sites-sl-es.md#get-site-sle-threshold)
* [List Site Sle Impacted Applications](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-applications)
* [List Site Sle Impacted Aps](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-aps)
* [List Site Sle Impacted Chassis](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-chassis)
* [List Site Sle Impacted Gateways](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-gateways)
* [List Site Sle Impacted Interfaces](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-interfaces)
* [List Site Sle Impacted Switches](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-switches)
* [List Site Sle Impacted Wired Clients](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-wired-clients)
* [List Site Sle Impacted Wireless Clients](../../doc/controllers/sites-sl-es.md#list-site-sle-impacted-wireless-clients)
* [List Site Sle Metric Classifiers](../../doc/controllers/sites-sl-es.md#list-site-sle-metric-classifiers)
* [List Site Sles Metrics](../../doc/controllers/sites-sl-es.md#list-site-sles-metrics)
* [Replace Site Sle Threshold](../../doc/controllers/sites-sl-es.md#replace-site-sle-threshold)
* [Update Site Sle Threshold](../../doc/controllers/sites-sl-es.md#update-site-sle-threshold)


# Get Site Sle Classifier Details

Get SLE classifier details

```go
GetSiteSleClassifierDetails(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SleSummaryScopeEnum,
    scopeId string,
    metric string,
    classifier string,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SleClassifierSummary],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SleSummaryScopeEnum`](../../doc/models/sle-summary-scope-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `classifier` | `string` | Template, Required | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleClassifierSummary](../../doc/models/sle-classifier-summary.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SleSummaryScopeEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"

classifier := "classifier4"





duration := "10m"

apiResponse, err := sitesSLEs.GetSiteSleClassifierDetails(ctx, siteId, scope, scopeId, metric, classifier, nil, nil, &duration)
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
  "classifier": {
    "impact": {
      "num_aps": 2,
      "num_users": 17,
      "total_aps": 3,
      "total_users": 20
    },
    "interval": 3600,
    "name": "wifi-interference",
    "samples": {
      "degraded": [
        0,
        0,
        210.03334,
        3.1333334,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        102.5,
        108.03333,
        0,
        0,
        201.9,
        566.48334,
        135.63333,
        0
      ],
      "duration": [
        0,
        0,
        210.03334,
        3.1333334,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        102.5,
        108.03333,
        0,
        0,
        201.9,
        566.48334,
        135.63333,
        0
      ],
      "total": [
        1302.3,
        1289.0167,
        1396.3167,
        1423.6666,
        1439.2167,
        1414.7,
        1361.0834,
        1371.5834,
        1372.0667,
        1339.1,
        1374.3667,
        1369.9,
        1352.4833,
        1382.8,
        1426.7167,
        1425.6333,
        1403.9333,
        1420.75,
        1416.8334,
        1437.3334,
        1425.1,
        1485.3667,
        1426.4333,
        444.13333
      ]
    },
    "x_label": "seconds",
    "y_label": "user-minutes"
  },
  "end": 1627312871,
  "failures": [],
  "impact": {
    "num_aps": 2,
    "num_users": 21,
    "total_aps": 3,
    "total_users": 26
  },
  "metric": "capacity",
  "start": 1627226471
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


# Get Site Sle Histogram

Get the histogram for the SLE metric

```go
GetSiteSleHistogram(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleHistogramScopeParametersEnum,
    scopeId string,
    metric string,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SleHistogram],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleHistogramScopeParametersEnum`](../../doc/models/site-sle-histogram-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleHistogram](../../doc/models/sle-histogram.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleHistogramScopeParametersEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"





duration := "10m"

apiResponse, err := sitesSLEs.GetSiteSleHistogram(ctx, siteId, scope, scopeId, metric, nil, nil, &duration)
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
  "data": [
    {
      "range": [
        null,
        0
      ],
      "value": 0
    },
    {
      "range": [
        0,
        10
      ],
      "value": 0
    },
    {
      "range": [
        10,
        20
      ],
      "value": 5105
    },
    {
      "range": [
        20,
        30
      ],
      "value": 10616
    },
    {
      "range": [
        30,
        40
      ],
      "value": 40051
    },
    {
      "range": [
        40,
        50
      ],
      "value": 141201
    },
    {
      "range": [
        50,
        60
      ],
      "value": 949823
    },
    {
      "range": [
        60,
        70
      ],
      "value": 686308
    },
    {
      "range": [
        70,
        80
      ],
      "value": 177670
    },
    {
      "range": [
        80,
        90
      ],
      "value": 689
    },
    {
      "range": [
        90,
        100
      ],
      "value": 0
    },
    {
      "range": [
        100,
        null
      ],
      "value": 0
    }
  ],
  "end": 1627055181,
  "metric": "capacity",
  "start": 1626968781,
  "x_label": "available_bandwidth(%)",
  "y_label": "seconds"
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


# Get Site Sle Impact Summary

Get impact summary counts optionally filtered by classifier and failure type

* Wireless SLE Fields: `wlan`, `device_type`, `device_os` ,`band`, `ap`, `server`, `mxedge`
* Wired SLE Fields: `switch`, `client`, `vlan`, `interface`, `chassis`
* WAN SLE Fields: `gateway`, `client`, `interface`, `chassis`, `peer_path`, `gateway_zones`

```go
GetSiteSleImpactSummary(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactSummaryScopeParametersEnum,
    scopeId string,
    metric string,
    start *int,
    end *int,
    duration *string,
    fields *models.SiteSleImpactSummaryFieldsParameterEnum,
    classifier *string) (
    models.ApiResponse[models.SleImpactSummary],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactSummaryScopeParametersEnum`](../../doc/models/site-sle-impact-summary-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `fields` | [`*models.SiteSleImpactSummaryFieldsParameterEnum`](../../doc/models/site-sle-impact-summary-fields-parameter-enum.md) | Query, Optional | - |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactSummary](../../doc/models/sle-impact-summary.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactSummaryScopeParametersEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"





duration := "10m"





apiResponse, err := sitesSLEs.GetSiteSleImpactSummary(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil, nil)
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
  "ap": [
    {
      "ap_mac": "5c5b3550bd2e",
      "degraded": 1486,
      "duration": 0,
      "name": "ap43-off.lab",
      "total": 27406
    },
    {
      "ap_mac": "d420b083e17a",
      "degraded": 3,
      "duration": 0,
      "name": "ap33-ent.lab",
      "total": 1193
    }
  ],
  "band": [
    {
      "band": "24",
      "degraded": 1410,
      "duration": 0,
      "name": "24",
      "total": 28536
    },
    {
      "band": "5",
      "degraded": 78,
      "duration": 0,
      "name": "5",
      "total": 4679
    }
  ],
  "classifier": "",
  "device_os": [
    {
      "degraded": 1329,
      "device_os": "",
      "duration": 0,
      "name": "unknown",
      "total": 27165
    },
    {
      "degraded": 81,
      "device_os": "Linux",
      "duration": 0,
      "name": "Linux",
      "total": 1437
    },
    {
      "degraded": 36,
      "device_os": "Android 11",
      "duration": 0,
      "name": "Android 11",
      "total": 761
    },
    {
      "degraded": 39,
      "device_os": "14.6",
      "duration": 0,
      "name": "14.6",
      "total": 2413
    },
    {
      "degraded": 2,
      "device_os": "Catalina",
      "duration": 0,
      "name": "Catalina",
      "total": 1438
    }
  ],
  "device_type": [
    {
      "degraded": 1410,
      "device_type": "",
      "duration": 0,
      "name": "unknown",
      "total": 28603
    },
    {
      "degraded": 2,
      "device_type": "iPhone",
      "duration": 0,
      "name": "iPhone",
      "total": 1263
    },
    {
      "degraded": 36,
      "device_type": "OnePlus",
      "duration": 0,
      "name": "OnePlus",
      "total": 761
    },
    {
      "degraded": 37,
      "device_type": "iPad",
      "duration": 0,
      "name": "iPad",
      "total": 1150
    },
    {
      "degraded": 2,
      "device_type": "Mac",
      "duration": 0,
      "name": "Mac",
      "total": 1438
    }
  ],
  "end": 1627312734,
  "failure": "",
  "metric": "capacity",
  "start": 1627226334,
  "wlan": [
    {
      "degraded": 37,
      "duration": 0,
      "name": "MlN.ADM",
      "total": 1150,
      "wlan_id": "ba3f85fc-ba48-4d8f-ad89-152e5c42db18"
    },
    {
      "degraded": 1410,
      "duration": 0,
      "name": "MlN",
      "total": 28603,
      "wlan_id": "649a2336-b1e0-47bd-961c-f637dbe50e7b"
    },
    {
      "degraded": 41,
      "duration": 0,
      "name": "MlN.1X",
      "total": 3462,
      "wlan_id": "a937da77-fe3c-4784-86c4-f2134d7b1483"
    }
  ]
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


# Get Site Sle Summary

Get the summary for the SLE metric

```go
GetSiteSleSummary(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleMetricSummaryScopeParametersEnum,
    scopeId string,
    metric string,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SleSummary],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleMetricSummaryScopeParametersEnum`](../../doc/models/site-sle-metric-summary-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleSummary](../../doc/models/sle-summary.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleMetricSummaryScopeParametersEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"





duration := "10m"

apiResponse, err := sitesSLEs.GetSiteSleSummary(ctx, siteId, scope, scopeId, metric, nil, nil, &duration)
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
  "classifiers": [
    {
      "impact": {
        "num_aps": 1,
        "num_users": 4,
        "total_aps": 3,
        "total_users": 26
      },
      "interval": 3600,
      "name": "client-count",
      "samples": {
        "degraded": [
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          5.8,
          0,
          0,
          0,
          4.65,
          0,
          7.55,
          47.55,
          13.266666
        ],
        "duration": [
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          5.8,
          0,
          0,
          0,
          4.65,
          0,
          7.55,
          47.55,
          13.266666
        ],
        "total": [
          1302.3,
          1289.0167,
          1396.3167,
          1423.6666,
          1439.2167,
          1414.7,
          1361.0834,
          1371.5834,
          1372.0667,
          1339.1,
          1374.3667,
          1369.9,
          1352.4833,
          1382.8,
          1426.7167,
          1425.6333,
          1403.9333,
          1420.75,
          1416.8334,
          1437.3334,
          1425.1,
          1485.3667,
          1426.4333,
          289.83334
        ]
      },
      "x_label": "seconds",
      "y_label": "user-minutes"
    },
    {
      "impact": {
        "num_aps": 2,
        "num_users": 17,
        "total_aps": 3,
        "total_users": 26
      },
      "interval": 3600,
      "name": "wifi-interference",
      "samples": {
        "degraded": [
          0,
          0,
          210.03334,
          3.1333334,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          102.5,
          108.03333,
          0,
          0,
          201.9,
          566.48334,
          135.63333,
          0
        ],
        "duration": [
          0,
          0,
          210.03334,
          3.1333334,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          102.5,
          108.03333,
          0,
          0,
          201.9,
          566.48334,
          135.63333,
          0
        ],
        "total": [
          1302.3,
          1289.0167,
          1396.3167,
          1423.6666,
          1439.2167,
          1414.7,
          1361.0834,
          1371.5834,
          1372.0667,
          1339.1,
          1374.3667,
          1369.9,
          1352.4833,
          1382.8,
          1426.7167,
          1425.6333,
          1403.9333,
          1420.75,
          1416.8334,
          1437.3334,
          1425.1,
          1485.3667,
          1426.4333,
          289.83334
        ]
      },
      "x_label": "seconds",
      "y_label": "user-minutes"
    },
    {
      "impact": {
        "num_aps": 0,
        "num_users": 0,
        "total_aps": 3,
        "total_users": 26
      },
      "interval": 3600,
      "name": "client_usage",
      "samples": {
        "degraded": [
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0
        ],
        "duration": [
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0
        ],
        "total": [
          1302.3,
          1289.0167,
          1396.3167,
          1423.6666,
          1439.2167,
          1414.7,
          1361.0834,
          1371.5834,
          1372.0667,
          1339.1,
          1374.3667,
          1369.9,
          1352.4833,
          1382.8,
          1426.7167,
          1425.6333,
          1403.9333,
          1420.75,
          1416.8334,
          1437.3334,
          1425.1,
          1485.3667,
          1426.4333,
          289.83334
        ]
      },
      "x_label": "seconds",
      "y_label": "user-minutes"
    },
    {
      "impact": {
        "num_aps": 1,
        "num_users": 17,
        "total_aps": 3,
        "total_users": 26
      },
      "interval": 3600,
      "name": "non-wifi-interference",
      "samples": {
        "degraded": [
          0,
          0,
          0,
          0,
          16.65,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          31.15,
          17.616667,
          17.85,
          0,
          0,
          0,
          0
        ],
        "duration": [
          0,
          0,
          0,
          0,
          16.65,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          0,
          31.15,
          17.616667,
          17.85,
          0,
          0,
          0,
          0
        ],
        "total": [
          1302.3,
          1289.0167,
          1396.3167,
          1423.6666,
          1439.2167,
          1414.7,
          1361.0834,
          1371.5834,
          1372.0667,
          1339.1,
          1374.3667,
          1369.9,
          1352.4833,
          1382.8,
          1426.7167,
          1425.6333,
          1403.9333,
          1420.75,
          1416.8334,
          1437.3334,
          1425.1,
          1485.3667,
          1426.4333,
          289.83334
        ]
      },
      "x_label": "seconds",
      "y_label": "user-minutes"
    }
  ],
  "end": 1627312606,
  "events": [],
  "impact": {
    "num_aps": 2,
    "num_users": 21,
    "total_aps": 3,
    "total_users": 26
  },
  "sle": {
    "interval": 3600,
    "name": "capacity",
    "samples": {
      "degraded": [
        0,
        0,
        210.03334,
        3.1333334,
        16.65,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        0,
        5.8,
        102.5,
        139.18333,
        17.616667,
        22.5,
        201.9,
        574.0333,
        183.18333,
        13.266666
      ],
      "total": [
        1302.3,
        1289.0167,
        1396.3167,
        1423.6666,
        1439.2167,
        1414.7,
        1361.0834,
        1371.5834,
        1372.0667,
        1339.1,
        1374.3667,
        1369.9,
        1352.4833,
        1382.8,
        1426.7167,
        1425.6333,
        1403.9333,
        1420.75,
        1416.8334,
        1437.3334,
        1425.1,
        1485.3667,
        1426.4333,
        289.83334
      ],
      "value": [
        0.6764934,
        0.6783766,
        0.641645,
        0.6934629,
        0.68676674,
        0.6834809,
        0.6961604,
        0.6979584,
        0.7033722,
        0.70410794,
        0.7025278,
        0.70305353,
        0.70292175,
        0.7009334,
        0.69344264,
        0.68596864,
        0.5952168,
        0.62183666,
        0.68161446,
        0.65352744,
        0.6183489,
        0.54178274,
        0.6044712,
        0.66845906
      ]
    },
    "x_label": "seconds",
    "y_label": "%"
  },
  "start": 1627226206
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


# Get Site Sle Threshold

Get the SLE threshold

```go
GetSiteSleThreshold(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleThresholdScopeParameterEnum,
    scopeId string,
    metric string) (
    models.ApiResponse[models.SleThreshold],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleThresholdScopeParameterEnum`](../../doc/models/site-sle-threshold-scope-parameter-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleThreshold](../../doc/models/sle-threshold.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleThresholdScopeParameterEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"

apiResponse, err := sitesSLEs.GetSiteSleThreshold(ctx, siteId, scope, scopeId, metric)
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
  "default": -72,
  "direction": "left",
  "maximum": -60,
  "metric": "coverage",
  "minimum": -90,
  "threshold": "-66",
  "units": "dBm"
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


# List Site Sle Impacted Applications

For WAN SLEs. List the impacted interfaces optionally filtered by classifier and failure type

```go
ListSiteSleImpactedApplications(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleScopeEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedApplications],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleScopeEnum`](../../doc/models/site-sle-scope-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedApplications](../../doc/models/sle-impacted-applications.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleScopeEnum_ENUMSWITCH

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedApplications(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "apps": [
    {
      "app": "ZOOM",
      "degraded": 371103,
      "duration": 0,
      "name": "ZOOM",
      "threshold": 173,
      "total": 1771274
    }
  ],
  "classifier": "",
  "end": 1668760746,
  "failure": "",
  "limit": "1000",
  "metric": "application_health",
  "page": 1,
  "start": 1668121200,
  "total_count": 1
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


# List Site Sle Impacted Aps

For Wireless SLEs. Listimpacted APs optionally filtered by classifier and failure type

```go
ListSiteSleImpactedAps(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedApsScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedAps],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactedApsScopeParametersEnum`](../../doc/models/site-sle-impacted-aps-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedAps](../../doc/models/sle-impacted-aps.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactedApsScopeParametersEnum_SITE

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedAps(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "aps": [
    {
      "ap_mac": "5c5b35500000",
      "degraded": 1486,
      "duration": 0,
      "name": "ap43.lab",
      "total": 27377
    },
    {
      "ap_mac": "d420b0830000",
      "degraded": 3,
      "duration": 0,
      "name": "ap33.lab",
      "total": 1189
    }
  ],
  "classifier": "",
  "end": 1627313016,
  "failure": "",
  "limit": 1000,
  "metric": "capacity",
  "page": 1,
  "start": 1627226616,
  "total_count": 2
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


# List Site Sle Impacted Chassis

For Wired and WAN SLEs. List the impacted interfaces optionally filtered by classifier and failure type

```go
ListSiteSleImpactedChassis(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedChassisScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedChassis],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactedChassisScopeParametersEnum`](../../doc/models/site-sle-impacted-chassis-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedChassis](../../doc/models/sle-impacted-chassis.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactedChassisScopeParametersEnum_ENUMSWITCH

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedChassis(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "chassis": [
    {
      "chassis": "1",
      "degraded": 12.283334,
      "duration": 13655.167,
      "role": "master",
      "switch_mac": "d0dd49012345",
      "switch_name": "test-chassis",
      "total": 13655.167
    }
  ],
  "classifier": "",
  "end": 1668760643,
  "failure": "",
  "limit": 1000,
  "metric": "switch_health",
  "page": 1,
  "start": 1668121200,
  "total_count": 1
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


# List Site Sle Impacted Gateways

For WAN SLEs. List the impacted interfaces optionally filtered by classifier and failure type

```go
ListSiteSleImpactedGateways(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedGatewaysScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedGateways],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactedGatewaysScopeParametersEnum`](../../doc/models/site-sle-impacted-gateways-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedGateways](../../doc/models/sle-impacted-gateways.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactedGatewaysScopeParametersEnum_SITE

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedGateways(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "classifier": "",
  "end": 1668760746,
  "failure": "",
  "gateways": [
    {
      "degraded": 758573.1,
      "duration": 2770997,
      "gateway_mac": "fc3342001122",
      "gateway_model": "SRX320",
      "gateway_version": "20.4R1.12",
      "name": "test-SRX",
      "total": 2770997
    }
  ],
  "limit": 1000,
  "metric": "application_health",
  "page": 1,
  "start": 1668121200,
  "total_count": 1
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


# List Site Sle Impacted Interfaces

For Wired and WAN SLEs. List the impacted interfaces optionally filtered by classifier and failure type

```go
ListSiteSleImpactedInterfaces(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedInterfacesScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedInterfaces],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactedInterfacesScopeParametersEnum`](../../doc/models/site-sle-impacted-interfaces-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedInterfaces](../../doc/models/sle-impacted-interfaces.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactedInterfacesScopeParametersEnum_ENUMSWITCH

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedInterfaces(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "classifier": "",
  "end": 1668760198,
  "failure": "",
  "interfaces": [
    {
      "degraded": 11.583333,
      "duration": 765.4667,
      "interface_name": "ge-0/0/10",
      "switch_mac": "2c2131001122",
      "switch_name": "test-ex",
      "total": 765.4667
    },
    {
      "degraded": 191.08333,
      "duration": 13775.35,
      "interface_name": "xe-0/1/0",
      "switch_mac": "2c2131001122",
      "switch_name": "test-ex",
      "total": 13775.35
    }
  ],
  "limit": 1000,
  "metric": "switch_throughput",
  "page": 1,
  "start": 1668726000,
  "total_count": 5
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


# List Site Sle Impacted Switches

For Wired SLEs. List the impacted switches optionally filtered by classifier and failure type

```go
ListSiteSleImpactedSwitches(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedSwitchesScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedSwitches],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactedSwitchesScopeParametersEnum`](../../doc/models/site-sle-impacted-switches-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedSwitches](../../doc/models/sle-impacted-switches.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactedSwitchesScopeParametersEnum_SITE

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedSwitches(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "classifier": "",
  "end": 1668760198,
  "failure": "",
  "limit": 1000,
  "metric": "switch_throughput",
  "page": 1,
  "start": 1668726000,
  "switches": [
    {
      "degraded": 109.88333,
      "duration": 5753.75,
      "interface": [
        "ge-0/0/11",
        "xe-0/1/0"
      ],
      "name": "test-ex",
      "switch_mac": "2c2131001122",
      "switch_model": "EX2300-C-12P",
      "switch_version": "20.4R3-S3.4",
      "total": 5753.75
    }
  ],
  "total_count": 1
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


# List Site Sle Impacted Wired Clients

For Wired SLEs. List the impacted interfaces optionally filtered by classifier and failure type

```go
ListSiteSleImpactedWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedClientsScopeParametersEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedClients],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactedClientsScopeParametersEnum`](../../doc/models/site-sle-impacted-clients-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedClients](../../doc/models/sle-impacted-clients.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactedClientsScopeParametersEnum_ENUMSWITCH

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedWiredClients(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "classifier": "",
  "clients": [
    {
      "degraded": 40,
      "duration": 11014,
      "mac": "001122334455",
      "name": "test-device",
      "switches": [
        {
          "interfaces": [
            "ge-0/0/6"
          ],
          "switch_mac": "2c2131001122",
          "switch_name": "test-ex"
        }
      ],
      "total": 11014
    }
  ],
  "end": 1668760198,
  "failure": "",
  "limit": 1000,
  "metric": "switch_throughput",
  "page": 1,
  "start": 1668726000,
  "total_count": 1
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


# List Site Sle Impacted Wireless Clients

For Wireless SLEs. List the impacted wireless users optionally filtered by classifier and failure type

```go
ListSiteSleImpactedWirelessClients(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleImpactedUsersScopeParameterEnum,
    scopeId uuid.UUID,
    metric string,
    start *int,
    end *int,
    duration *string,
    classifier *string) (
    models.ApiResponse[models.SleImpactedUsers],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleImpactedUsersScopeParameterEnum`](../../doc/models/site-sle-impacted-users-scope-parameter-enum.md) | Template, Required | - |
| `scopeId` | `uuid.UUID` | Template, Required | - |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `classifier` | `*string` | Query, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleImpactedUsers](../../doc/models/sle-impacted-users.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleImpactedUsersScopeParameterEnum_AP

scopeId := uuid.MustParse("000008e8-0000-0000-0000-000000000000")

metric := "metric8"





duration := "10m"



apiResponse, err := sitesSLEs.ListSiteSleImpactedWirelessClients(ctx, siteId, scope, scopeId, metric, nil, nil, &duration, nil)
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
  "classifier": "",
  "end": 1627313103,
  "failure": "",
  "limit": 1000,
  "metric": "capacity",
  "page": 1,
  "start": 1627226703,
  "total_count": 21,
  "users": [
    {
      "ap_mac": "d420b0830000",
      "ap_name": "ap33.lab",
      "degraded": 2,
      "device_os": "14.6",
      "device_type": "iPhone",
      "duration": 1270,
      "mac": "dc080f360000",
      "name": "aPhone-20973",
      "ssid": "lab.1X",
      "total": 1270,
      "wlan_id": "a937da77-0000-0000-0000-f2134d7b1483"
    },
    {
      "ap_mac": "5c5b35500000",
      "ap_name": "ap43.lab",
      "degraded": 36,
      "device_os": "Android 11",
      "device_type": "OnePlus",
      "duration": 767,
      "mac": "4c4feedc0000",
      "name": "OnePlus-8",
      "ssid": "lab.1X",
      "total": 767,
      "wlan_id": "a937da77-0000-0000-0000-f2134d7b1483"
    },
    {
      "ap_mac": "5c5b35500000",
      "ap_name": "ap43.lab",
      "degraded": 2,
      "device_os": "Catalina",
      "device_type": "Mac",
      "duration": 1405,
      "mac": "a483e7390000",
      "name": "tmunzer-mbp",
      "ssid": "lab.1X",
      "total": 1405,
      "wlan_id": "a937da77-0000-0000-0000-f2134d7b1483"
    },
    {
      "ap_mac": "5c5b35500000",
      "ap_name": "ap43.lab",
      "degraded": 81,
      "device_os": "Linux",
      "device_type": "unknown",
      "duration": 1403,
      "mac": "5caafd0d0000",
      "name": "SonosZP",
      "ssid": "lab",
      "total": 1403,
      "wlan_id": "649a2336-0000-0000-0000-f637dbe50e7b"
    }
  ]
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


# List Site Sle Metric Classifiers

List classifiers for a specific metric

```go
ListSiteSleMetricClassifiers(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleMetricClassifiersScopeParametersEnum,
    scopeId string,
    metric string) (
    models.ApiResponse[[]string],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleMetricClassifiersScopeParametersEnum`](../../doc/models/site-sle-metric-classifiers-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []string.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleMetricClassifiersScopeParametersEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"

apiResponse, err := sitesSLEs.ListSiteSleMetricClassifiers(ctx, siteId, scope, scopeId, metric)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
[
  "asymmetry-uplink",
  "weak-signal",
  "asymmetry-downlink"
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


# List Site Sles Metrics

List the metrics for the given scope

```go
ListSiteSlesMetrics(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleMetricsScopeParametersEnum,
    scopeId string) (
    models.ApiResponse[models.SiteSleMetrics],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleMetricsScopeParametersEnum`](../../doc/models/site-sle-metrics-scope-parameters-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SiteSleMetrics](../../doc/models/site-sle-metrics.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleMetricsScopeParametersEnum_GATEWAY

scopeId := "scope_id0"

apiResponse, err := sitesSLEs.ListSiteSlesMetrics(ctx, siteId, scope, scopeId)
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
  "enabled": [
    "coverage",
    "capacity",
    "time-to-connect",
    "failed-to-connect",
    "roaming",
    "roaming-v2",
    "throughput",
    "switch_health",
    "switch_throughput",
    "switch_stc",
    "gateway-health",
    "application_health",
    "wan-link-health",
    "ap-availability"
  ],
  "supported": [
    "coverage",
    "capacity",
    "time-to-connect",
    "failed-to-connect",
    "roaming",
    "roaming-v2",
    "location_jitter",
    "location_latency",
    "throughput",
    "location_dropped-requests",
    "switch_health",
    "switch_throughput",
    "switch_stc",
    "gateway-health",
    "application_health",
    "wan-link-health",
    "ap-availability",
    "location_sdk-connect-time",
    "location_ble-hung"
  ]
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


# Replace Site Sle Threshold

Replace the SLE threshold

```go
ReplaceSiteSleThreshold(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleThresholdScopeParameterEnum,
    scopeId string,
    metric string,
    body *models.SleThreshold) (
    models.ApiResponse[models.SleThreshold],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleThresholdScopeParameterEnum`](../../doc/models/site-sle-threshold-scope-parameter-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `body` | [`*models.SleThreshold`](../../doc/models/sle-threshold.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleThreshold](../../doc/models/sle-threshold.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleThresholdScopeParameterEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"

body := models.SleThreshold{
    Maximum:              models.ToPointer(float64(-60)),
    Minimum:              models.ToPointer(float64(-90)),
}

apiResponse, err := sitesSLEs.ReplaceSiteSleThreshold(ctx, siteId, scope, scopeId, metric, &body)
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
  "default": -72,
  "direction": "left",
  "maximum": -60,
  "metric": "coverage",
  "minimum": -90,
  "threshold": "-66",
  "units": "dBm"
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


# Update Site Sle Threshold

Update the SLE threshold

```go
UpdateSiteSleThreshold(
    ctx context.Context,
    siteId uuid.UUID,
    scope models.SiteSleThresholdScopeParameterEnum,
    scopeId string,
    metric string,
    body *models.SleThreshold) (
    models.ApiResponse[models.SleThreshold],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `scope` | [`models.SiteSleThresholdScopeParameterEnum`](../../doc/models/site-sle-threshold-scope-parameter-enum.md) | Template, Required | - |
| `scopeId` | `string` | Template, Required | * site_id if `scope`==`site`<br>* device_id if `scope`==`ap`, `scope`==`switch` or `scope`==`gateway`<br>* mac if `scope`==`client` |
| `metric` | `string` | Template, Required | Values from /api/v1/sites/{site_id}/sle/{scope}/{scope_id}/metrics |
| `body` | [`*models.SleThreshold`](../../doc/models/sle-threshold.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SleThreshold](../../doc/models/sle-threshold.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

scope := models.SiteSleThresholdScopeParameterEnum_GATEWAY

scopeId := "scope_id0"

metric := "metric8"

body := models.SleThreshold{
    Maximum:              models.ToPointer(float64(-60)),
    Minimum:              models.ToPointer(float64(-90)),
}

apiResponse, err := sitesSLEs.UpdateSiteSleThreshold(ctx, siteId, scope, scopeId, metric, &body)
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
  "default": -72,
  "direction": "left",
  "maximum": -60,
  "metric": "coverage",
  "minimum": -90,
  "threshold": "-66",
  "units": "dBm"
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

