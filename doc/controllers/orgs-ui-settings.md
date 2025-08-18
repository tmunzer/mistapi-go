# Orgs UI Settings

```go
orgsUISettings := client.OrgsUISettings()
```

## Class Name

`OrgsUISettings`

## Methods

* [Create Org Ui Settings](../../doc/controllers/orgs-ui-settings.md#create-org-ui-settings)
* [Delete Org Ui Setting](../../doc/controllers/orgs-ui-settings.md#delete-org-ui-setting)
* [Get Org Ui Setting](../../doc/controllers/orgs-ui-settings.md#get-org-ui-setting)
* [List Org Ui Settings](../../doc/controllers/orgs-ui-settings.md#list-org-ui-settings)
* [Update Org Ui Setting](../../doc/controllers/orgs-ui-settings.md#update-org-ui-setting)


# Create Org Ui Settings

Create an Org UI settings/databoard

```go
CreateOrgUiSettings(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgUiSettings) (
    models.ApiResponse[models.OrgUiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgUiSettings`](../../doc/models/org-ui-settings.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgUiSettings](../../doc/models/org-ui-settings.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgUiSettings{
    Description:          models.ToPointer("AP related stats"),
    IsCustomDataboard:    models.ToPointer(true),
    Name:                 models.ToPointer("AP Stats"),
    Purpose:              models.ToPointer(models.OrgUiSettingsPurposeEnum_MARVISDASHBOARD),
    Tiles:                []models.OrgUiSettingsTile{
        models.OrgUiSettingsTile{
            Description:          models.ToPointer("User typed tile descr"),
            IsAutoTitle:          models.ToPointer(true),
            Name:                 models.ToPointer("List top 10 APs by bandwidth"),
            NlQuery:              models.ToPointer("List top 10 APs by bandwidth"),
            Position:             models.ToPointer(models.OrgUiSettingsTilePosition{
                Col:                  models.ToPointer(1),
                ColSpan:              models.ToPointer(2),
                Row:                  models.ToPointer(1),
                RowSpan:              models.ToPointer(1),
            }),
        },
    },
}

apiResponse, err := orgsUISettings.CreateOrgUiSettings(ctx, orgId, &body)
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
  "created_time": 1749083436,
  "description": "AP related stats",
  "for_site": false,
  "id": "9a702097-0dd3-48af-909b-2be4ff94d139",
  "isCustomDataboard": true,
  "modified_time": 1749083436,
  "name": "AP Stats",
  "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
  "purpose": "marvisdashboard",
  "site_id": "00000000-0000-0000-0000-000000000000",
  "tiles": [
    {
      "description": "User typed tile descr",
      "id": "3eef7c83-3d33-417a-a729-4772d4a1013a",
      "isAutoTitle": true,
      "name": "List top 10 APs by bandwidth",
      "nl_query": "List top 10 APs by bandwidth",
      "position": {
        "col": 1,
        "colSpan": 2,
        "row": 1,
        "rowSpan": 1
      }
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


# Delete Org Ui Setting

Delete an Org UI settings

```go
DeleteOrgUiSetting(
    ctx context.Context,
    orgId uuid.UUID,
    uisettingId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `uisettingId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

uisettingId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsUISettings.DeleteOrgUiSetting(ctx, orgId, uisettingId)
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


# Get Org Ui Setting

Get an Org UI settings/databoard

```go
GetOrgUiSetting(
    ctx context.Context,
    orgId uuid.UUID,
    uisettingId uuid.UUID) (
    models.ApiResponse[models.OrgUiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `uisettingId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgUiSettings](../../doc/models/org-ui-settings.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

uisettingId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsUISettings.GetOrgUiSetting(ctx, orgId, uisettingId)
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
  "created_time": 1749083436,
  "description": "AP related stats",
  "for_site": false,
  "id": "9a702097-0dd3-48af-909b-2be4ff94d139",
  "isCustomDataboard": true,
  "modified_time": 1749083436,
  "name": "AP Stats",
  "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
  "purpose": "marvisdashboard",
  "site_id": "00000000-0000-0000-0000-000000000000",
  "tiles": [
    {
      "description": "User typed tile descr",
      "id": "3eef7c83-3d33-417a-a729-4772d4a1013a",
      "isAutoTitle": true,
      "name": "List top 10 APs by bandwidth",
      "nl_query": "List top 10 APs by bandwidth",
      "position": {
        "col": 1,
        "colSpan": 2,
        "row": 1,
        "rowSpan": 1
      }
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


# List Org Ui Settings

List the Orgs UI settings/databoard

```go
ListOrgUiSettings(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.OrgUiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.OrgUiSettings](../../doc/models/org-ui-settings.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsUISettings.ListOrgUiSettings(ctx, orgId)
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
    "created_time": 1749083436,
    "description": "AP related stats",
    "for_site": false,
    "id": "9a702097-0dd3-48af-909b-2be4ff94d139",
    "isCustomDataboard": true,
    "modified_time": 1749083436,
    "name": "AP Stats",
    "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
    "purpose": "marvisdashboard",
    "site_id": "00000000-0000-0000-0000-000000000000",
    "tiles": [
      {
        "description": "User typed tile descr",
        "id": "3eef7c83-3d33-417a-a729-4772d4a1013a",
        "isAutoTitle": true,
        "name": "List top 10 APs by bandwidth",
        "nl_query": "List top 10 APs by bandwidth",
        "position": {
          "col": 1,
          "colSpan": 2,
          "row": 1,
          "rowSpan": 1
        }
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


# Update Org Ui Setting

Org UI settings/databoard

```go
UpdateOrgUiSetting(
    ctx context.Context,
    orgId uuid.UUID,
    uisettingId uuid.UUID,
    body *models.OrgUiSettings) (
    models.ApiResponse[models.OrgUiSettings],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `uisettingId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgUiSettings`](../../doc/models/org-ui-settings.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgUiSettings](../../doc/models/org-ui-settings.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

uisettingId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgUiSettings{
    Description:          models.ToPointer("AP related stats"),
    IsCustomDataboard:    models.ToPointer(true),
    Name:                 models.ToPointer("AP Stats"),
    Purpose:              models.ToPointer(models.OrgUiSettingsPurposeEnum_MARVISDASHBOARD),
    Tiles:                []models.OrgUiSettingsTile{
        models.OrgUiSettingsTile{
            Description:          models.ToPointer("User typed tile descr"),
            IsAutoTitle:          models.ToPointer(true),
            Name:                 models.ToPointer("List top 10 APs by bandwidth"),
            NlQuery:              models.ToPointer("List top 10 APs by bandwidth"),
            Position:             models.ToPointer(models.OrgUiSettingsTilePosition{
                Col:                  models.ToPointer(1),
                ColSpan:              models.ToPointer(2),
                Row:                  models.ToPointer(1),
                RowSpan:              models.ToPointer(1),
            }),
        },
    },
}

apiResponse, err := orgsUISettings.UpdateOrgUiSetting(ctx, orgId, uisettingId, &body)
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
  "created_time": 1749083436,
  "description": "AP related stats",
  "for_site": false,
  "id": "9a702097-0dd3-48af-909b-2be4ff94d139",
  "isCustomDataboard": true,
  "modified_time": 1749083436,
  "name": "AP Stats",
  "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
  "purpose": "marvisdashboard",
  "site_id": "00000000-0000-0000-0000-000000000000",
  "tiles": [
    {
      "description": "User typed tile descr",
      "id": "3eef7c83-3d33-417a-a729-4772d4a1013a",
      "isAutoTitle": true,
      "name": "List top 10 APs by bandwidth",
      "nl_query": "List top 10 APs by bandwidth",
      "position": {
        "col": 1,
        "colSpan": 2,
        "row": 1,
        "rowSpan": 1
      }
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

