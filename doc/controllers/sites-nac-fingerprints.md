# Sites NAC Fingerprints

```go
sitesNACFingerprints := client.SitesNACFingerprints()
```

## Class Name

`SitesNACFingerprints`

## Methods

* [Count Site Client Fingerprints](../../doc/controllers/sites-nac-fingerprints.md#count-site-client-fingerprints)
* [Search Site Client Fingerprints](../../doc/controllers/sites-nac-fingerprints.md#search-site-client-fingerprints)


# Count Site Client Fingerprints

Count Client Fingerprints

```go
CountSiteClientFingerprints(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.FingerprintsCountDistinctEnum,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.FingerprintsCountDistinctEnum`](../../doc/models/fingerprints-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `family`, `model`, `os`, `os_type`<br><br>**Default**: `"family"` |
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

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.FingerprintsCountDistinctEnum_FAMILY

duration := "10m"

limit := 100

apiResponse, err := sitesNACFingerprints.CountSiteClientFingerprints(ctx, siteId, &distinct, nil, nil, &duration, &limit)
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


# Search Site Client Fingerprints

Search Client Fingerprints

```go
SearchSiteClientFingerprints(
    ctx context.Context,
    siteId uuid.UUID,
    family *string,
    clientType *models.NacAccessTypeEnum,
    model *string,
    mfg *string,
    os *string,
    osType *string,
    mac *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    interval *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.FingerprintSearchResult],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `family` | `*string` | Query, Optional | Device Category of the client device |
| `clientType` | [`*models.NacAccessTypeEnum`](../../doc/models/nac-access-type-enum.md) | Query, Optional | Filter results by client type. enum: `wireless`, `wired`, `vty` |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `mfg` | `*string` | Query, Optional | Manufacturer name of the client device |
| `os` | `*string` | Query, Optional | Operating System name and version of the client device |
| `osType` | `*string` | Query, Optional | Operating system name of the client device |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order.<br><br>**Default**: `"wxid"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.FingerprintSearchResult](../../doc/models/fingerprint-search-result.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

family := "EX Series Switch"

clientType := models.NacAccessTypeEnum_WIRED

model := "ex4100-f-12p"

mfg := "Juniper Networks, Inc."

os := "JUNOS 22.3R1.12"

osType := "JUNOS"

mac := "d420b080516d"

limit := 100

duration := "10m"

interval := "10m"

sort := "-site_id"

apiResponse, err := sitesNACFingerprints.SearchSiteClientFingerprints(ctx, siteId, &family, &clientType, &model, &mfg, &os, &osType, &mac, &limit, nil, nil, &duration, &interval, &sort, nil)
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
  "end": 1735678700,
  "limit": 10,
  "results": [
    {
      "family": "Apple",
      "mac": "d420b080516e",
      "mfg": "Apple, Inc.",
      "model": "Unknown",
      "org_id": "bb2fb165-0931-49c7-a1b8-9b5814326b7d",
      "os": "iOS 18.1.1",
      "os_type": "iOS",
      "random_mac": true,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 1735678662.58
    },
    {
      "family": "EX Series Switch",
      "mac": "d420b080516d",
      "mfg": "Juniper Networks, Inc.",
      "model": "ex4100-f-12p",
      "org_id": "b6bc08f3-60a3-402b-8f0d-caf9132a1e9a",
      "os": "JUNOS 22.3R1.12",
      "os_type": "JUNOS",
      "random_mac": false,
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "timestamp": 1735669092.932
    }
  ],
  "start": 1735678650,
  "total": 2
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

