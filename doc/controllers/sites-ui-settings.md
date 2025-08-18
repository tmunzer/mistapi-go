# Sites UI Settings

```go
sitesUISettings := client.SitesUISettings()
```

## Class Name

`SitesUISettings`

## Methods

* [Create Site Ui Settings](../../doc/controllers/sites-ui-settings.md#create-site-ui-settings)
* [Delete Site Ui Setting](../../doc/controllers/sites-ui-settings.md#delete-site-ui-setting)
* [Get Site Ui Setting](../../doc/controllers/sites-ui-settings.md#get-site-ui-setting)
* [List Site Ui Setting Derived](../../doc/controllers/sites-ui-settings.md#list-site-ui-setting-derived)
* [List Site Ui Settings](../../doc/controllers/sites-ui-settings.md#list-site-ui-settings)
* [Update Site Ui Setting](../../doc/controllers/sites-ui-settings.md#update-site-ui-setting)


# Create Site Ui Settings

Create a Site UI settings/databoard

```go
CreateSiteUiSettings(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.UiSettings) (
    models.ApiResponse[models.UiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UiSettings`](../../doc/models/ui-settings.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UiSettings](../../doc/models/ui-settings.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UiSettings{
    DefaultScopeId:       models.ToPointer("67970e46-4e12-11e6-9188-0242ad112847"),
    DefaultScopeType:     models.ToPointer("site"),
    DefaultTimeRange:     models.ToPointer(models.UiSettingsDefaultTimeRange{
        End:                  models.ToPointer(1508828400),
        EndDate:              models.ToPointer("10/23/2017"),
        Interval:             models.ToPointer("1d"),
        Name:                 models.ToPointer("This Week"),
        ShortName:            models.ToPointer("thisWeek"),
        Start:                models.ToPointer(1508655600),
        UsePreset:            models.ToPointer(true),
    }),
    Description:          "Description of the databoard",
    IsCustomDataboard:    models.ToPointer(true),
    IsScopeLinked:        models.ToPointer(true),
    IsTimeRangeLinked:    models.ToPointer(true),
    Name:                 models.ToPointer("New Databoard"),
    Purpose:              "databoard",
    Tiles:                []models.UiSettingsTile{
        models.UiSettingsTile{
            ChartBand:            models.ToPointer("2.4 ghz"),
            ChartColor:           models.ToPointer("#00B4AD"),
            ChartDirection:       models.ToPointer("tx + rx"),
            ChartRankBy:          models.ToPointer("string"),
            ChartType:            models.ToPointer("timeSeries"),
            Colspan:              models.ToPointer(5),
            Column:               models.ToPointer(1),
            HideEmptyRows:        models.ToPointer(true),
            Metric:               models.ToPointer(models.UiSettingsTileMetric{
                ApiName:              models.ToPointer("client_dhcp_latency"),
            }),
            Name:                 models.ToPointer("New Analysis"),
            Row:                  models.ToPointer(1),
            Rowspan:              models.ToPointer(2),
            ScopeId:              models.ToPointer("e0c767834b4c"),
            ScopeType:            models.ToPointer("client"),
            TimeRange:            models.ToPointer(models.UiSettingsTileTimeRange{
                End:                  models.ToPointer(float64(1508823743)),
                EndDate:              models.ToPointer("10/23/2017"),
                Interval:             models.ToPointer("1d"),
                Name:                 models.ToPointer("Past 7 Days"),
                ShortName:            models.ToPointer("7d"),
                Start:                models.ToPointer(1508223600),
                UsePreset:            models.ToPointer(true),
            }),
            TrendType:            models.ToPointer("line"),
            VizType:              models.ToPointer("averageTimeSeriesChart"),
        },
    },
}

apiResponse, err := sitesUISettings.CreateSiteUiSettings(ctx, siteId, &body)
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
  "created_time": 1508823803,
  "defaultScopeId": "67970e46-4e12-11e6-9188-0242ad112847",
  "defaultScopeType": "site",
  "defaultTimeRange": {
    "end": 1508828400,
    "endDate": "10/23/2017",
    "interval": "1d",
    "name": "This Week",
    "shortName": "thisWeek",
    "start": 1508655600,
    "usePreset": true
  },
  "description": "Description of the databoard",
  "for_site": true,
  "id": "3bdcc7e8-c04d-4512-b4fc-093da9057eb0",
  "isCustomDataboard": true,
  "isScopeLinked": true,
  "isTimeRangeLinked": true,
  "modified_time": 0,
  "name": "New Databoard",
  "org_id": "cc079380-5029-4d4a-9125-858de85731ff",
  "purpose": "databoard",
  "site_id": "67970e46-4e12-11e6-9188-0242ad112847",
  "tiles": [
    {
      "chartBand": "2.4 ghz",
      "chartColor": "#00B4AD",
      "chartDirection": "tx + rx",
      "chartRankBy": "string",
      "chartType": "timeSeries",
      "colspan": 5,
      "column": 1,
      "hideEmptyRows": true,
      "id": "7a9ab38c-cfc3-483d-b51a-0aec571fadc0",
      "metric": {
        "apiName": "client_dhcp_latency"
      },
      "name": "New Analysis",
      "row": 1,
      "rowspan": 2,
      "scopeId": "e0c767834b4c",
      "scopeType": "client",
      "sortedColumns": null,
      "timeRange": {
        "end": 1508823743,
        "endDate": "10/23/2017",
        "interval": "1d",
        "name": "Past 7 Days",
        "shortName": "7d",
        "start": 1508223600,
        "usePreset": true
      },
      "trendType": "line",
      "vizType": "averageTimeSeriesChart"
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


# Delete Site Ui Setting

Site UI settings

```go
DeleteSiteUiSetting(
    ctx context.Context,
    siteId uuid.UUID,
    uisettingId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `uisettingId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

uisettingId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesUISettings.DeleteSiteUiSetting(ctx, siteId, uisettingId)
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


# Get Site Ui Setting

Site UI settings

```go
GetSiteUiSetting(
    ctx context.Context,
    siteId uuid.UUID,
    uisettingId uuid.UUID) (
    models.ApiResponse[models.UiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `uisettingId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UiSettings](../../doc/models/ui-settings.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

uisettingId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesUISettings.GetSiteUiSetting(ctx, siteId, uisettingId)
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
  "created_time": 1508823803,
  "defaultScopeId": "67970e46-4e12-11e6-9188-0242ad112847",
  "defaultScopeType": "site",
  "defaultTimeRange": {
    "end": 1508828400,
    "endDate": "10/23/2017",
    "interval": "1d",
    "name": "This Week",
    "shortName": "thisWeek",
    "start": 1508655600,
    "usePreset": true
  },
  "description": "Description of the databoard",
  "for_site": true,
  "id": "3bdcc7e8-c04d-4512-b4fc-093da9057eb0",
  "isCustomDataboard": true,
  "isScopeLinked": true,
  "isTimeRangeLinked": true,
  "modified_time": 0,
  "name": "New Databoard",
  "org_id": "cc079380-5029-4d4a-9125-858de85731ff",
  "purpose": "databoard",
  "site_id": "67970e46-4e12-11e6-9188-0242ad112847",
  "tiles": [
    {
      "chartBand": "2.4 ghz",
      "chartColor": "#00B4AD",
      "chartDirection": "tx + rx",
      "chartRankBy": "string",
      "chartType": "timeSeries",
      "colspan": 5,
      "column": 1,
      "hideEmptyRows": true,
      "id": "7a9ab38c-cfc3-483d-b51a-0aec571fadc0",
      "metric": {
        "apiName": "client_dhcp_latency"
      },
      "name": "New Analysis",
      "row": 1,
      "rowspan": 2,
      "scopeId": "e0c767834b4c",
      "scopeType": "client",
      "sortedColumns": null,
      "timeRange": {
        "end": 1508823743,
        "endDate": "10/23/2017",
        "interval": "1d",
        "name": "Past 7 Days",
        "shortName": "7d",
        "start": 1508223600,
        "usePreset": true
      },
      "trendType": "line",
      "vizType": "averageTimeSeriesChart"
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


# List Site Ui Setting Derived

Get both site UI settings(for_site=true) and org UI settings (for_site=false)

```go
ListSiteUiSettingDerived(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.UiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UiSettings](../../doc/models/ui-settings.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesUISettings.ListSiteUiSettingDerived(ctx, siteId)
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
  "created_time": 1508823803,
  "defaultScopeId": "67970e46-4e12-11e6-9188-0242ad112847",
  "defaultScopeType": "site",
  "defaultTimeRange": {
    "end": 1508828400,
    "endDate": "10/23/2017",
    "interval": "1d",
    "name": "This Week",
    "shortName": "thisWeek",
    "start": 1508655600,
    "usePreset": true
  },
  "description": "Description of the databoard",
  "for_site": true,
  "id": "3bdcc7e8-c04d-4512-b4fc-093da9057eb0",
  "isCustomDataboard": true,
  "isScopeLinked": true,
  "isTimeRangeLinked": true,
  "modified_time": 0,
  "name": "New Databoard",
  "org_id": "cc079380-5029-4d4a-9125-858de85731ff",
  "purpose": "databoard",
  "site_id": "67970e46-4e12-11e6-9188-0242ad112847",
  "tiles": [
    {
      "chartBand": "2.4 ghz",
      "chartColor": "#00B4AD",
      "chartDirection": "tx + rx",
      "chartRankBy": "string",
      "chartType": "timeSeries",
      "colspan": 5,
      "column": 1,
      "hideEmptyRows": true,
      "id": "7a9ab38c-cfc3-483d-b51a-0aec571fadc0",
      "metric": {
        "apiName": "client_dhcp_latency"
      },
      "name": "New Analysis",
      "row": 1,
      "rowspan": 2,
      "scopeId": "e0c767834b4c",
      "scopeType": "client",
      "sortedColumns": null,
      "timeRange": {
        "end": 1508823743,
        "endDate": "10/23/2017",
        "interval": "1d",
        "name": "Past 7 Days",
        "shortName": "7d",
        "start": 1508223600,
        "usePreset": true
      },
      "trendType": "line",
      "vizType": "averageTimeSeriesChart"
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


# List Site Ui Settings

List the Site UI settings/databoard

```go
ListSiteUiSettings(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]models.UiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.UiSettings](../../doc/models/ui-settings.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesUISettings.ListSiteUiSettings(ctx, siteId)
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
    "created_time": 1508823803,
    "defaultScopeId": "67970e46-4e12-11e6-9188-0242ad112847",
    "defaultScopeType": "site",
    "defaultTimeRange": {
      "end": 1508828400,
      "endDate": "10/23/2017",
      "interval": "1d",
      "name": "This Week",
      "shortName": "thisWeek",
      "start": 1508655600,
      "usePreset": true
    },
    "description": "Description of the databoard",
    "for_site": true,
    "id": "3bdcc7e8-c04d-4512-b4fc-093da9057eb0",
    "isCustomDataboard": true,
    "isScopeLinked": true,
    "isTimeRangeLinked": true,
    "modified_time": 0,
    "name": "New Databoard",
    "org_id": "cc079380-5029-4d4a-9125-858de85731ff",
    "purpose": "databoard",
    "site_id": "67970e46-4e12-11e6-9188-0242ad112847",
    "tiles": [
      {
        "chartBand": "2.4 ghz",
        "chartColor": "#00B4AD",
        "chartDirection": "tx + rx",
        "chartRankBy": "string",
        "chartType": "timeSeries",
        "colspan": 5,
        "column": 1,
        "hideEmptyRows": true,
        "id": "7a9ab38c-cfc3-483d-b51a-0aec571fadc0",
        "metric": {
          "apiName": "client_dhcp_latency"
        },
        "name": "New Analysis",
        "row": 1,
        "rowspan": 2,
        "scopeId": "e0c767834b4c",
        "scopeType": "client",
        "sortedColumns": null,
        "timeRange": {
          "end": 1508823743,
          "endDate": "10/23/2017",
          "interval": "1d",
          "name": "Past 7 Days",
          "shortName": "7d",
          "start": 1508223600,
          "usePreset": true
        },
        "trendType": "line",
        "vizType": "averageTimeSeriesChart"
      }
    ]
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


# Update Site Ui Setting

Site UI settings

```go
UpdateSiteUiSetting(
    ctx context.Context,
    siteId uuid.UUID,
    uisettingId uuid.UUID,
    body *models.UiSettings) (
    models.ApiResponse[models.UiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `uisettingId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UiSettings`](../../doc/models/ui-settings.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.UiSettings](../../doc/models/ui-settings.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

uisettingId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UiSettings{
    DefaultScopeId:       models.ToPointer("67970e46-4e12-11e6-9188-0242ad112847"),
    DefaultScopeType:     models.ToPointer("site"),
    DefaultTimeRange:     models.ToPointer(models.UiSettingsDefaultTimeRange{
        End:                  models.ToPointer(1508828400),
        EndDate:              models.ToPointer("10/23/2017"),
        Interval:             models.ToPointer("1d"),
        Name:                 models.ToPointer("This Week"),
        ShortName:            models.ToPointer("thisWeek"),
        Start:                models.ToPointer(1508655600),
        UsePreset:            models.ToPointer(true),
    }),
    Description:          "Description of the databoard",
    IsCustomDataboard:    models.ToPointer(true),
    IsScopeLinked:        models.ToPointer(true),
    IsTimeRangeLinked:    models.ToPointer(true),
    Name:                 models.ToPointer("New Databoard"),
    Purpose:              "databoard",
    Tiles:                []models.UiSettingsTile{
        models.UiSettingsTile{
            ChartBand:            models.ToPointer("2.4 ghz"),
            ChartColor:           models.ToPointer("#00B4AD"),
            ChartDirection:       models.ToPointer("tx + rx"),
            ChartRankBy:          models.ToPointer("string"),
            ChartType:            models.ToPointer("timeSeries"),
            Colspan:              models.ToPointer(5),
            Column:               models.ToPointer(1),
            HideEmptyRows:        models.ToPointer(true),
            Metric:               models.ToPointer(models.UiSettingsTileMetric{
                ApiName:              models.ToPointer("client_dhcp_latency"),
            }),
            Name:                 models.ToPointer("New Analysis"),
            Row:                  models.ToPointer(1),
            Rowspan:              models.ToPointer(2),
            ScopeId:              models.ToPointer("e0c767834b4c"),
            ScopeType:            models.ToPointer("client"),
            TimeRange:            models.ToPointer(models.UiSettingsTileTimeRange{
                End:                  models.ToPointer(float64(1508823743)),
                EndDate:              models.ToPointer("10/23/2017"),
                Interval:             models.ToPointer("1d"),
                Name:                 models.ToPointer("Past 7 Days"),
                ShortName:            models.ToPointer("7d"),
                Start:                models.ToPointer(1508223600),
                UsePreset:            models.ToPointer(true),
            }),
            TrendType:            models.ToPointer("line"),
            VizType:              models.ToPointer("averageTimeSeriesChart"),
        },
    },
}

apiResponse, err := sitesUISettings.UpdateSiteUiSetting(ctx, siteId, uisettingId, &body)
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
  "created_time": 1508823803,
  "defaultScopeId": "67970e46-4e12-11e6-9188-0242ad112847",
  "defaultScopeType": "site",
  "defaultTimeRange": {
    "end": 1508828400,
    "endDate": "10/23/2017",
    "interval": "1d",
    "name": "This Week",
    "shortName": "thisWeek",
    "start": 1508655600,
    "usePreset": true
  },
  "description": "Description of the databoard",
  "for_site": true,
  "id": "3bdcc7e8-c04d-4512-b4fc-093da9057eb0",
  "isCustomDataboard": true,
  "isScopeLinked": true,
  "isTimeRangeLinked": true,
  "modified_time": 0,
  "name": "New Databoard",
  "org_id": "cc079380-5029-4d4a-9125-858de85731ff",
  "purpose": "databoard",
  "site_id": "67970e46-4e12-11e6-9188-0242ad112847",
  "tiles": [
    {
      "chartBand": "2.4 ghz",
      "chartColor": "#00B4AD",
      "chartDirection": "tx + rx",
      "chartRankBy": "string",
      "chartType": "timeSeries",
      "colspan": 5,
      "column": 1,
      "hideEmptyRows": true,
      "id": "7a9ab38c-cfc3-483d-b51a-0aec571fadc0",
      "metric": {
        "apiName": "client_dhcp_latency"
      },
      "name": "New Analysis",
      "row": 1,
      "rowspan": 2,
      "scopeId": "e0c767834b4c",
      "scopeType": "client",
      "sortedColumns": null,
      "timeRange": {
        "end": 1508823743,
        "endDate": "10/23/2017",
        "interval": "1d",
        "name": "Past 7 Days",
        "shortName": "7d",
        "start": 1508223600,
        "usePreset": true
      },
      "trendType": "line",
      "vizType": "averageTimeSeriesChart"
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

