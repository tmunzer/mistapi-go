# Sites Clients-Wan

```go
sitesClientsWan := client.SitesClientsWan()
```

## Class Name

`SitesClientsWan`

## Methods

* [Count Site Wan Client Events](../../doc/controllers/sites-clients-wan.md#count-site-wan-client-events)
* [Count Site Wan Clients](../../doc/controllers/sites-clients-wan.md#count-site-wan-clients)
* [Search Site Wan Client Events](../../doc/controllers/sites-clients-wan.md#search-site-wan-client-events)
* [Search Site Wan Clients](../../doc/controllers/sites-clients-wan.md#search-site-wan-clients)


# Count Site Wan Client Events

Count by Distinct Attributes of Site WAN Client-Events

```go
CountSiteWanClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteWanClientEventsDistinctEnum,
    mType *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteWanClientEventsDistinctEnum`](../../doc/models/site-wan-client-events-distinct-enum.md) | Query, Optional | **Default**: `"type"` |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteWanClientEventsDistinctEnum_ENUMTYPE

duration := "10m"

limit := 100

apiResponse, err := sitesClientsWan.CountSiteWanClientEvents(ctx, siteId, &distinct, nil, nil, nil, &duration, &limit)
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


# Count Site Wan Clients

Count by Distinct Attributes of Site WAN Clients

```go
CountSiteWanClients(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteWanClientsCountDistinctEnum,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteWanClientsCountDistinctEnum`](../../doc/models/site-wan-clients-count-distinct-enum.md) | Query, Optional | **Default**: `"mac"` |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteWanClientsCountDistinctEnum_MAC

duration := "10m"

limit := 100

apiResponse, err := sitesClientsWan.CountSiteWanClients(ctx, siteId, &distinct, nil, nil, &duration, &limit)
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


# Search Site Wan Client Events

Search Site WAN Client Events

```go
SearchSiteWanClientEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    mac *string,
    hostname *string,
    ip *string,
    mfg *string,
    nacruleId *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.SearchEventsWanClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `mac` | `*string` | Query, Optional | Partial / full MAC address |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `ip` | `*string` | Query, Optional | Client IP |
| `mfg` | `*string` | Query, Optional | Manufacture |
| `nacruleId` | `*string` | Query, Optional | nacrule_id |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchEventsWanClient](../../doc/models/search-events-wan-client.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "0011223"

hostname := "test-hostname"

ip := "10.4.2.4"

mfg := "Juniper"

nacruleId := "00000000-0000-0000-0000-000000000000"

duration := "10m"

limit := 100

apiResponse, err := sitesClientsWan.SearchSiteWanClientEvents(ctx, siteId, nil, &mac, &hostname, &ip, &mfg, &nacruleId, nil, nil, &duration, &limit)
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
  "results": {
    "When": "2022-12-31 23:59:59.293000+00:00",
    "ev_type": "CLIENT_IP_ASSIGNED",
    "metadata": {},
    "org_id": "b0b9f142-aaba-11e6-aafc-0242ac110002",
    "random_mac": true,
    "site_id": "fc656275-b157-43fd-b922-5f4f341c19bf",
    "text": "DHCP Ack IP 192.168.88.216",
    "wcid": "62bbfb75-10d8-49d1-dec7-d2df91624287"
  },
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


# Search Site Wan Clients

Search Site WAN Clients

```go
SearchSiteWanClients(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    hostname *string,
    ip *string,
    mfg *string,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[models.SearchWanClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Partial / full MAC address |
| `hostname` | `*string` | Query, Optional | Partial / full hostname |
| `ip` | `*string` | Query, Optional | Client IP |
| `mfg` | `*string` | Query, Optional | Manufacture |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWanClient](../../doc/models/search-wan-client.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "001122334455"

hostname := "test-hostname"

ip := "10.2.52.4"

mfg := "Cisco"

duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesClientsWan.SearchSiteWanClients(ctx, siteId, &mac, &hostname, &ip, &mfg, nil, nil, &duration, &limit, &page)
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
  "results": [
    {
      "When": "2022-12-31T23:59:43.497+0000",
      "hostname": [
        "sonoszp"
      ],
      "ip": [
        "192.168.1.139"
      ],
      "last_hostname": "sonoszp",
      "last_ip": "192.168.1.139",
      "mfg": "Sonos",
      "org_id": "b4e16c72-d50e-4c03-a952-a3217e231e2c",
      "site_id": "f688779c-e335-4f88-8d7c-9c5e9964528b",
      "wcid": "8bbe7389-212b-c65d-2208-00fab2017936"
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

