# Orgs Clients-Wan

```go
orgsClientsWan := client.OrgsClientsWan()
```

## Class Name

`OrgsClientsWan`

## Methods

* [Count Org Wan Client Events](../../doc/controllers/orgs-clients-wan.md#count-org-wan-client-events)
* [Count Org Wan Clients](../../doc/controllers/orgs-clients-wan.md#count-org-wan-clients)
* [Search Org Wan Client Events](../../doc/controllers/orgs-clients-wan.md#search-org-wan-client-events)
* [Search Org Wan Clients](../../doc/controllers/orgs-clients-wan.md#search-org-wan-clients)


# Count Org Wan Client Events

Count by Distinct Attributes of Org WAN Client-Events

```go
CountOrgWanClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgWanClientsEventsCountDistinctEnum,
    mType *string,
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
| `distinct` | [`*models.OrgWanClientsEventsCountDistinctEnum`](../../doc/models/org-wan-clients-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `hostname`, `ip`, `mac`, `mfg`, `type`<br><br>**Default**: `"type"` |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
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

distinct := models.OrgWanClientsEventsCountDistinctEnum_ENUMTYPE

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWan.CountOrgWanClientEvents(ctx, orgId, &distinct, nil, nil, nil, &duration, &limit)
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


# Count Org Wan Clients

Count by Distinct Attributes of Org WAN Clients

```go
CountOrgWanClients(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgWanClientsCountDistinctEnum,
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
| `distinct` | [`*models.OrgWanClientsCountDistinctEnum`](../../doc/models/org-wan-clients-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `hostname`, `ip`, `mac`, `mfg`, `network`<br><br>**Default**: `"mac"` |
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

distinct := models.OrgWanClientsCountDistinctEnum_MAC

duration := "10m"

limit := 100

apiResponse, err := orgsClientsWan.CountOrgWanClients(ctx, orgId, &distinct, nil, nil, &duration, &limit)
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


# Search Org Wan Client Events

Search Org WAN Client Events

```go
SearchOrgWanClientEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    mac *string,
    hostname *string,
    ip *string,
    mfg *string,
    nacruleId *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.SearchEventsWanClient],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `mac` | `*string` | Query, Optional | Partial / full Client MAC address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `aabbcc*` and `*bbcc*` match `aabbccddeeff`). Suffix-only wildcards (e.g. `*bccddeeff`) are not supported |
| `hostname` | `*string` | Query, Optional | Partial / full Client hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `everest*` and `*rest*` match `my-everest-client`). Suffix-only wildcards (e.g. `*everest`) are not supported |
| `ip` | `*string` | Query, Optional | Partial / full Client IP address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `10.100.10.*` and `*100.10.*` match `10.100.10.54`). Suffix-only wildcards (e.g. `*.54`) are not supported |
| `mfg` | `*string` | Query, Optional | Partial / full Client manufacturer (e.g. "apple", "cisco", "juniper"). Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `Raspberry Pi*` and `*Pi*` match `Raspberry Pi Trading Ltd`). Suffix-only wildcards (e.g. `*Ltd`) are not supported |
| `nacruleId` | `*string` | Query, Optional | Filter results by NAC rule identifier |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchEventsWanClient](../../doc/models/search-events-wan-client.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mac := "aabbccddeeff"

hostname := "my-everest-client"

ip := "10.100.10.54"

mfg := "Raspberry Pi Trading Ltd"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWan.SearchOrgWanClientEvents(ctx, orgId, nil, &mac, &hostname, &ip, &mfg, nil, &limit, nil, nil, &duration, &sort, nil)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Search Org Wan Clients

Search Org WAN Clients

```go
SearchOrgWanClients(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *uuid.UUID,
    hostname *string,
    ip *string,
    ipSrc *string,
    mac *string,
    mfg *string,
    network *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.SearchWanClient],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*uuid.UUID` | Query, Optional | Filter results by site identifier |
| `hostname` | `*string` | Query, Optional | Partial / full Client hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `everest*` and `*rest*` match `my-everest-client`). Suffix-only wildcards (e.g. `*everest`) are not supported. Accepts multiple comma-separated values. |
| `ip` | `*string` | Query, Optional | Partial / full Client IP address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `10.100.10.*` and `*100.10.*` match `10.100.10.54`). Suffix-only wildcards (e.g. `*.54`) are not supported. Accepts multiple comma-separated values. |
| `ipSrc` | `*string` | Query, Optional | Filter results by source IP address |
| `mac` | `*string` | Query, Optional | Filter results by MAC address. Accepts multiple comma-separated values. |
| `mfg` | `*string` | Query, Optional | Filter results by manufacturer. Accepts multiple comma-separated values. |
| `network` | `*string` | Query, Optional | Partial / full Name of the network the client is/was connected to. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `my-corp*` and `*corp*` match `my-corp-network`). Suffix-only wildcards (e.g. `*corp`) are not supported. Accepts multiple comma-separated values. |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: Example response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.SearchWanClient](../../doc/models/search-wan-client.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := uuid.MustParse("72771e6a-6f5e-4de4-a5b9-1266c4197811")

hostname := "my-everest-client,my-everest*"

ip := "10.100.10.54,10.100.10.*"

ipSrc := "dhcp"

mac := "5c5b53010101,5c5b53020202"

mfg := "Unknown,Juniper Networks"

network := "my-corp-network,my-corp*"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsClientsWan.SearchOrgWanClients(ctx, orgId, &siteId, &hostname, &ip, &ipSrc, &mac, &mfg, &network, &limit, nil, nil, &duration, &sort, nil)
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
  "results": [
    {
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

