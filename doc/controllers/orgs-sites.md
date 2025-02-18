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

Count Sites

```go
CountOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgSitesCountDistinctEnum,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgSitesCountDistinctEnum`](../../doc/models/org-sites-count-distinct-enum.md) | Query, Optional | **Default**: `"id"` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgSitesCountDistinctEnum_ID





duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsSites.CountOrgSites(ctx, orgId, &distinct, nil, nil, &duration, &limit, &page)
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


# Create Org Site

Create Org Site

```go
CreateOrgSite(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.Site) (
    models.ApiResponse[models.Site],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.Site`](../../doc/models/site.md) | Body, Optional | Request Body |

## Response Type

[`models.Site`](../../doc/models/site.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.Site{
    Address:              models.ToPointer("1601 S. Deanza Blvd., Cupertino, CA, 95014"),
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
    Notes:                models.ToPointer("string"),
    RftemplateId:         models.NewOptional(models.ToPointer(uuid.MustParse("bb8a9017-1e36-5d6c-6f2b-551abe8a76a2"))),
    SecpolicyId:          models.NewOptional(models.ToPointer(uuid.MustParse("3bcd0beb-5d0a-4cbd-92c1-14aea91e98ef"))),
    SitegroupIds:         []uuid.UUID{
        uuid.MustParse("497f6eca-6276-4997-bfeb-53cbbbba6f3b"),
    },
    Timezone:             models.ToPointer("America/Los_Angeles"),
    AdditionalProperties: map[string]interface{}{
        "apporttemplate_id": interface{}("string"),
    },
}

apiResponse, err := orgsSites.CreateOrgSite(ctx, orgId, &body)
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
  "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
  "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
  "apporttemplate_id": "string",
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# List Org Sites

Get List of Org Sites

```go
ListOrgSites(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.Site],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.Site`](../../doc/models/site.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

apiResponse, err := orgsSites.ListOrgSites(ctx, orgId, &limit, &page)
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
    "address": "1601 S. Deanza Blvd., Cupertino, CA, 95014",
    "alarmtemplate_id": "684dfc5c-fe77-2290-eb1d-ef3d677fe168",
    "apporttemplate_id": "string",
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Search Org Sites

Search Sites

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
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.ResponseSiteSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `analyticEnabled` | `*bool` | Query, Optional | If Advanced Analytic feature is enabled |
| `appWaking` | `*bool` | Query, Optional | If App Waking feature is enabled |
| `assetEnabled` | `*bool` | Query, Optional | If Asset Tracking is enabled |
| `autoUpgradeEnabled` | `*bool` | Query, Optional | If Auto Upgrade feature is enabled |
| `autoUpgradeVersion` | `*string` | Query, Optional | If Auto Upgrade feature is enabled |
| `countryCode` | `*string` | Query, Optional | Site country code |
| `honeypotEnabled` | `*bool` | Query, Optional | If Honeypot detection is enabled |
| `id` | `*string` | Query, Optional | Site id |
| `locateUnconnected` | `*bool` | Query, Optional | If unconnected client are located |
| `meshEnabled` | `*bool` | Query, Optional | If Mesh feature is enabled |
| `name` | `*string` | Query, Optional | Site name |
| `rogueEnabled` | `*bool` | Query, Optional | If Rogue detection is enabled |
| `remoteSyslogEnabled` | `*bool` | Query, Optional | If Remote Syslog is enabled |
| `rtsaEnabled` | `*bool` | Query, Optional | If managed mobility feature is enabled |
| `vnaEnabled` | `*bool` | Query, Optional | If Virtual Network Assistant is enabled |
| `wifiEnabled` | `*bool` | Query, Optional | If WIFI feature is enabled |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |

## Response Type

[`models.ResponseSiteSearch`](../../doc/models/response-site-search.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

































limit := 100





duration := "10m"

apiResponse, err := orgsSites.SearchOrgSites(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

