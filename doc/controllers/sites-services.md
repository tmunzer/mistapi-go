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

Count by Distinct Attributes of Service Path Events

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
    mac *string,
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
| `distinct` | [`*models.SiteServiceEventsCountDistinctEnum`](../../doc/models/site-service-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `mac`, `model`, `policy`, `port_id`, `site_id`, `type`, `vpn_name`, `vpn_path`<br><br>**Default**: `"type"` |
| `mType` | `*string` | Query, Optional | Event type, e.g. GW_SERVICE_PATH_DOWN |
| `text` | `*string` | Query, Optional | Description of the event including the reason it is triggered |
| `vpnName` | `*string` | Query, Optional | Filter results by vpn name |
| `vpnPath` | `*string` | Query, Optional | Filter results by vpn path |
| `policy` | `*string` | Query, Optional | Service policy associated with that specific path |
| `portId` | `*string` | Query, Optional | Filter results by port identifier |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `version` | `*string` | Query, Optional | Filter results by software version |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
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

distinct := models.SiteServiceEventsCountDistinctEnum_ENUMTYPE

duration := "10m"

limit := 100

apiResponse, err := sitesServices.CountSiteServicePathEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# List Site Services Derived

Get the list of derived Services for a Site

```go
ListSiteServicesDerived(
    ctx context.Context,
    siteId uuid.UUID,
    resolve *bool) (
    models.ApiResponse[[]models.Service],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `resolve` | `*bool` | Query, Optional | Whether resolve the site variables<br><br>**Default**: `false` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.Service](../../doc/models/service.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resolve := false

apiResponse, err := sitesServices.ListSiteServicesDerived(ctx, siteId, &resolve)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


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
    mac *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseEventsPathSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Event type, e.g. GW_SERVICE_PATH_DOWN |
| `text` | `*string` | Query, Optional | Description of the event including the reason it is triggered |
| `peerPortId` | `*string` | Query, Optional | Port ID of the peer gateway |
| `peerMac` | `*string` | Query, Optional | MAC address of the peer gateway |
| `vpnName` | `*string` | Query, Optional | Filter results by vpn name |
| `vpnPath` | `*string` | Query, Optional | Filter results by vpn path |
| `policy` | `*string` | Query, Optional | Service policy associated with that specific path |
| `portId` | `*string` | Query, Optional | Filter results by port identifier |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `version` | `*string` | Query, Optional | Filter results by software version |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsPathSearch](../../doc/models/response-events-path-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesServices.SearchSiteServicePathEvents(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

