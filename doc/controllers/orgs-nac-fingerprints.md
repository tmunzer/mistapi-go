# Orgs NAC Fingerprints

```go
orgsNACFingerprints := client.OrgsNACFingerprints()
```

## Class Name

`OrgsNACFingerprints`

## Methods

* [Count Org Client Fingerprints](../../doc/controllers/orgs-nac-fingerprints.md#count-org-client-fingerprints)
* [Search Org Client Fingerprints](../../doc/controllers/orgs-nac-fingerprints.md#search-org-client-fingerprints)


# Count Org Client Fingerprints

Count Client Fingerprints

```go
CountOrgClientFingerprints(
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

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.FingerprintsCountDistinctEnum`](../../doc/models/fingerprints-count-distinct-enum.md) | Query, Optional | **Default**: `"family"` |
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

distinct := models.FingerprintsCountDistinctEnum_FAMILY

duration := "10m"

limit := 100

apiResponse, err := orgsNACFingerprints.CountOrgClientFingerprints(ctx, siteId, &distinct, nil, nil, &duration, &limit)
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


# Search Org Client Fingerprints

Search Client Fingerprints

```go
SearchOrgClientFingerprints(
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
    sort *string) (
    models.ApiResponse[models.FingerprintSearchResult],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `family` | `*string` | Query, Optional | Device Category  of the client device |
| `clientType` | [`*models.NacAccessTypeEnum`](../../doc/models/nac-access-type-enum.md) | Query, Optional | Whether client is wired or wireless |
| `model` | `*string` | Query, Optional | Model name of the client device |
| `mfg` | `*string` | Query, Optional | Manufacturer name of the client device |
| `os` | `*string` | Query, Optional | Operating System name and version of the client device |
| `osType` | `*string` | Query, Optional | Operating system name of the client device |
| `mac` | `*string` | Query, Optional | MAC address of the client device |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `interval` | `*string` | Query, Optional | Aggregation works by giving a time range plus interval (e.g. 1d, 1h, 10m) where aggregation function would be applied to. |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order.<br><br>**Default**: `"wxid"` |

## Response Type

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

apiResponse, err := orgsNACFingerprints.SearchOrgClientFingerprints(ctx, siteId, &family, &clientType, &model, &mfg, &os, &osType, &mac, &limit, nil, nil, &duration, &interval, &sort)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

