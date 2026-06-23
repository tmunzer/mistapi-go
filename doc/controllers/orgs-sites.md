# Orgs Sites

```go
orgsSites := client.OrgsSites()
```

## Class Name

`OrgsSites`

## Methods

* [Count Org Sites](../../doc/controllers/orgs-sites.md#count-org-sites)
* [Create Org Site](../../doc/controllers/orgs-sites.md#create-org-site)
* [List Org Sites](../../doc/controllers/orgs-sites.md#list-org-sites)
* [Search Org Sites](../../doc/controllers/orgs-sites.md#search-org-sites)


# Count Org Sites

Count organization sites grouped by a distinct site attribute, such as country code, site name, feature flags, or enabled services.

```go
CountOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgSitesCountDistinctEnum,
    start *string,
    end *string,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgSitesCountDistinctEnum`](../../doc/models/org-sites-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `analytic_enabled`, `app_waking`, `asset_enabled`, `auto_upgrade_enabled`, `auto_upgrade_version`, `country_code`, `honeypot_enabled`, `id`, `locate_unconnected`, `mesh_enabled`, `name`, `remote_syslog_enabled`, `rogue_enabled`, `rtsa_enabled`, `vna_enabled`, `wifi_enabled`<br><br>**Default**: `"id"` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

**200**: Result of Count

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgSitesCountDistinctEnum_ID

duration := "10m"

limit := 100

apiResponse, err := orgsSites.CountOrgSites(ctx, orgId, &distinct, nil, nil, &duration, &limit)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Create Org Site

Create a site in the organization with location metadata, timezone, site group membership, and optional template or policy associations.

```go
CreateOrgSite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Site) (
    models.ApiResponse[models.Site],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Site`](../../doc/models/site.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.Site](../../doc/models/site.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Site{
    Address:              models.NewOptional(models.ToPointer("1601 S. Deanza Blvd., Cupertino, CA, 95014")),
    AlarmtemplateId:      models.NewOptional(models.ToPointer(uuid.MustParse("684dfc5c-fe77-2290-eb1d-ef3d677fe168"))),
    AptemplateId:         models.NewOptional(models.ToPointer(uuid.MustParse("16bdf952-ade2-4491-80b0-85ce506c760b"))),
    CountryCode:          models.ToPointer("US"),
    GatewaytemplateId:    models.NewOptional(models.ToPointer(uuid.MustParse("6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f"))),
    Latlng:               models.ToPointer(models.LatLng{
        Lat:                  float64(37.295833),
        Lng:                  float64(-122.032946),
    }),
    Name:                 "Mist Office",
    NetworktemplateId:    models.NewOptional(models.ToPointer(uuid.MustParse("12ae9bd2-e0ab-107b-72e8-a7a005565ec2"))),
    Notes:                models.NewOptional(models.ToPointer("string")),
    RftemplateId:         models.NewOptional(models.ToPointer(uuid.MustParse("bb8a9017-1e36-5d6c-6f2b-551abe8a76a2"))),
    SecpolicyId:          models.NewOptional(models.ToPointer(uuid.MustParse("3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef"))),
    SitegroupIds:         []uuid.UUID{
        uuid.MustParse("497f6eca-6276-4997-bfeb-53cbbbba6f3b"),
    },
    Timezone:             models.ToPointer("America/Los_Angeles"),
}

apiResponse, err := orgsSites.CreateOrgSite(ctx, orgId, &body)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
  "aptemplate_id": "16bdf952-ade2-4491-80b0-85ce506c760b",
  "country_code": "US",
  "created_time": 0,
  "gatewaytemplate_id": "6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f",
  "id": "497f6eca-6276-5005-bfeb-53cbbbba6f17",
  "latlng": {
    "lat": 37.295833,
    "lng": -122.032946
  },
  "modified_time": 0,
  "name": "Mist Office",
  "networktemplate_id": "12ae9bd2-e0ab-107b-72e8-a7a005565ec2",
  "notes": "string",
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "rftemplate_id": "bb8a9017-1e36-5d6c-6f2b-551abe8a76a2",
  "secpolicy_id": "3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef",
  "sitegroup_ids": [
    "497f6eca-6276-5006-bfeb-53cbbbba6f18"
  ],
  "timezone": "America/Los_Angeles"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# List Org Sites

List sites in the organization. Site records include location metadata, timezone, site group membership, and template or policy associations.

```go
ListOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Site],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Site](../../doc/models/site.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSites.ListOrgSites(ctx, orgId, &limit, &page)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
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
    "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
    "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
    "aptemplate_id": "16bdf952-ade2-4491-80b0-85ce506c760b",
    "country_code": "US",
    "created_time": 0,
    "gatewaytemplate_id": "6f9b2e75-9b2f-b5ae-81e3-e14c76f1a90f",
    "id": "497f6eca-6276-5007-bfeb-53cbbbba6f19",
    "latlng": {
      "lat": 37.295833,
      "lng": -122.032946
    },
    "modified_time": 0,
    "name": "Mist Office",
    "networktemplate_id": "12ae9bd2-e0ab-107b-72e8-a7a005565ec2",
    "notes": "string",
    "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
    "rftemplate_id": "bb8a9017-1e36-5d6c-6f2b-551abe8a76a2",
    "secpolicy_id": "3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef",
    "sitegroup_ids": [
      "497f6eca-6276-5008-bfeb-53cbbbba6f1a"
    ],
    "timezone": "America/Los_Angeles"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Sites

Search organization sites with filters for feature flags, country code, identifiers, names, upgrade settings, and other site attributes. Supports pagination and sorting.

```go
SearchOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    analyticEnabled *bool,
    appWaking *bool,
    assetEnabled *bool,
    autoUpgradeEnabled *bool,
    autoUpgradeVersion *string,
    countryCode *string,
    honeypotEnabled *bool,
    id *string,
    locateUnconnected *bool,
    meshEnabled *bool,
    name *string,
    rogueEnabled *bool,
    remoteSyslogEnabled *bool,
    rtsaEnabled *bool,
    vnaEnabled *bool,
    wifiEnabled *bool,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseSiteSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `analyticEnabled` | `*bool` | Query, Optional | If Advanced Analytic feature is enabled. Accepts multiple comma-separated boolean values. |
| `appWaking` | `*bool` | Query, Optional | If App Waking feature is enabled |
| `assetEnabled` | `*bool` | Query, Optional | If Asset Tracking is enabled. Accepts multiple comma-separated boolean values. |
| `autoUpgradeEnabled` | `*bool` | Query, Optional | Filter results by whether automatic upgrades are enabled |
| `autoUpgradeVersion` | `*string` | Query, Optional | Filter results by automatic upgrade version. Accepts multiple comma-separated values. |
| `countryCode` | `*string` | Query, Optional | Filter results by country code. Accepts multiple comma-separated values. |
| `honeypotEnabled` | `*bool` | Query, Optional | If Honeypot detection is enabled. Accepts multiple comma-separated boolean values. |
| `id` | `*string` | Query, Optional | Filter results by identifier. Accepts multiple comma-separated values. |
| `locateUnconnected` | `*bool` | Query, Optional | If unconnected client are located. Accepts multiple comma-separated boolean values. |
| `meshEnabled` | `*bool` | Query, Optional | If Mesh feature is enabled. Accepts multiple comma-separated boolean values. |
| `name` | `*string` | Query, Optional | Partial / full Site name. Case insensitive. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `my-site*` and `*site*` match `my-site-01`). Suffix-only wildcards (e.g. `*site-01`) are not supported. Accepts multiple comma-separated values. |
| `rogueEnabled` | `*bool` | Query, Optional | If Rogue detection is enabled. Accepts multiple comma-separated boolean values. |
| `remoteSyslogEnabled` | `*bool` | Query, Optional | If Remote Syslog is enabled |
| `rtsaEnabled` | `*bool` | Query, Optional | If managed mobility feature is enabled. Accepts multiple comma-separated boolean values. |
| `vnaEnabled` | `*bool` | Query, Optional | If Virtual Network Assistant is enabled. Accepts multiple comma-separated boolean values. |
| `wifiEnabled` | `*bool` | Query, Optional | If Wi-Fi feature is enabled |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSiteSearch](../../doc/models/response-site-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

autoUpgradeVersion := "custom,stable"

countryCode := "CA,FR"

id := "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"

name := "my-site-01,my-site*"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsSites.SearchOrgSites(ctx, orgId, nil, nil, nil, nil, &autoUpgradeVersion, &countryCode, nil, &id, nil, nil, &name, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "end": 0,
  "limit": 0,
  "next": "string",
  "results": [
    {
      "auto_upgrade_enabled": true,
      "auto_upgrade_version": "string",
      "country_code": "string",
      "honeypot_enabled": true,
      "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "name": "string",
      "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "site_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
      "timestamp": 0,
      "timezone": "string",
      "vna_enabled": true,
      "wifi_enabled": true
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

