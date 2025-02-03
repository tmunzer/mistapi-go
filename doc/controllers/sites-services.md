# Sites Services

```go
sitesServices := client.SitesServices()
```

## Class Name

`SitesServices`

## Methods

* [Count Site Service Path Events](../../doc/controllers/sites-services.md#count-site-service-path-events)
* [List Site Services Derived](../../doc/controllers/sites-services.md#list-site-services-derived)
* [Search Site Service Path Events](../../doc/controllers/sites-services.md#search-site-service-path-events)


# Count Site Service Path Events

Count Service Path Events

```go
CountSiteServicePathEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteServiceEventsCountDistinctEnum,
    mType *string,
    text *string,
    vpnName *string,
    vpnPath *string,
    policy *string,
    portId *string,
    model *string,
    version *string,
    timestamp *float64,
    mac *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteServiceEventsCountDistinctEnum`](../../doc/models/site-service-events-count-distinct-enum.md) | Query, Optional | **Default**: `"type"` |
| `mType` | `*string` | Query, Optional | Event type, e.g. GW_SERVICE_PATH_DOWN |
| `text` | `*string` | Query, Optional | Description of the event including the reason it is triggered |
| `vpnName` | `*string` | Query, Optional | Peer name |
| `vpnPath` | `*string` | Query, Optional | Peer path name |
| `policy` | `*string` | Query, Optional | Service policy associated with that specific path |
| `portId` | `*string` | Query, Optional | Network interface |
| `model` | `*string` | Query, Optional | Device model |
| `version` | `*string` | Query, Optional | Device firmware version |
| `timestamp` | `*float64` | Query, Optional | Start time, in epoch |
| `mac` | `*string` | Query, Optional | MAC address |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteServiceEventsCountDistinctEnum_ENUMTYPE

























duration := "10m"

limit := 100

apiResponse, err := sitesServices.CountSiteServicePathEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# List Site Services Derived

Retrieves the list of Services available for the Site

```go
ListSiteServicesDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.Service],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | Whether resolve the site variables<br>**Default**: `false` |

## Response Type

[`[]models.Service`](../../doc/models/service.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resolve := false

apiResponse, err := sitesServices.ListSiteServicesDerived(ctx, siteId, &resolve)
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
    "addresses": [
      "string"
    ],
    "apps": [
      "string"
    ],
    "dscp": 8,
    "hostnames": [
      "string"
    ],
    "max_jitter": 0,
    "max_latency": 0,
    "max_loss": 0,
    "name": "string",
    "specs": [
      {
        "port_range": "0",
        "protocol": "any"
      }
    ],
    "traffic_class": "best_effort",
    "traffic_type": "default",
    "type": "custom"
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


# Search Site Service Path Events

Search Service Path Events

```go
SearchSiteServicePathEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    text *string,
    peerPortId *string,
    peerMac *string,
    vpnName *string,
    vpnPath *string,
    policy *string,
    portId *string,
    model *string,
    version *string,
    timestamp *float64,
    mac *string,
    start *int,
    end *int,
    duration *string,
    limit *int) (
    models.ApiResponse[models.ResponseEventsPathSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Event type, e.g. GW_SERVICE_PATH_DOWN |
| `text` | `*string` | Query, Optional | Description of the event including the reason it is triggered |
| `peerPortId` | `*string` | Query, Optional | Port ID of the peer gateway |
| `peerMac` | `*string` | Query, Optional | MAC address of the peer gateway |
| `vpnName` | `*string` | Query, Optional | Peer name |
| `vpnPath` | `*string` | Query, Optional | Peer path name |
| `policy` | `*string` | Query, Optional | Service policy associated with that specific path |
| `portId` | `*string` | Query, Optional | Network interface |
| `model` | `*string` | Query, Optional | Device model |
| `version` | `*string` | Query, Optional | Device firmware version |
| `timestamp` | `*float64` | Query, Optional | Start time, in epoch |
| `mac` | `*string` | Query, Optional | MAC address |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |

## Response Type

[`models.ResponseEventsPathSearch`](../../doc/models/response-events-path-search.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





























duration := "10m"

limit := 100

apiResponse, err := sitesServices.SearchSiteServicePathEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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
  "end": 1697096379,
  "limit": 10,
  "results": [
    {
      "mac": "90ec7734b374",
      "model": "SSR120",
      "org_id": "a3c6718f-2823-4e48-bf5e-b841768a4c9b",
      "policy": "INTERNET",
      "port_id": "ge-1/0/6",
      "site_id": "4279edbd-1d24-41ea-9505-2eb26c8590fa",
      "text": "Peer Path Down",
      "timestamp": 1697037328.651775,
      "type": "GW_SERVICE_PATH_REMOVE",
      "version": "6.1.5-14.lts",
      "vpn_name": "Syracuse_HUB",
      "vpn_path": "Syracuse_HUB-Wan0"
    }
  ],
  "start": 1697009979,
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

