# Sites Clients-Wired

```go
sitesClientsWired := client.SitesClientsWired()
```

## Class Name

`SitesClientsWired`

## Methods

* [Count Site Wired Clients](../../doc/controllers/sites-clients-wired.md#count-site-wired-clients)
* [Search Site Wired Clients](../../doc/controllers/sites-clients-wired.md#search-site-wired-clients)


# Count Site Wired Clients

Count by Distinct Attributes of Clients

```go
CountSiteWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteWiredClientsCountDistinctEnum,
    mac *string,
    deviceMac *string,
    portId *string,
    vlan *string,
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
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.SiteWiredClientsCountDistinctEnum`](../../doc/models/site-wired-clients-count-distinct-enum.md) | Query, Optional | - |
| `mac` | `*string` | Query, Optional | client mac |
| `deviceMac` | `*string` | Query, Optional | device mac |
| `portId` | `*string` | Query, Optional | port id |
| `vlan` | `*string` | Query, Optional | vlan |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |
| `limit` | `*int` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.SiteWiredClientsCountDistinctEnum("mac")













duration := "10m"

limit := 100

page := 1

apiResponse, err := sitesClientsWired.CountSiteWiredClients(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, &duration, &limit, &page)
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


# Search Site Wired Clients

Search Wired Clients

```go
SearchSiteWiredClients(
    ctx context.Context,
    siteId uuid.UUID,
    deviceMac *string,
    mac *string,
    ip *string,
    portId *string,
    vlan *string,
    manufacture *string,
    text *string,
    nacruleId *string,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.SearchWiredClient],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `*string` | Query, Optional | device mac |
| `mac` | `*string` | Query, Optional | client mac |
| `ip` | `*string` | Query, Optional | client ip |
| `portId` | `*string` | Query, Optional | port id |
| `vlan` | `*string` | Query, Optional | vlan |
| `manufacture` | `*string` | Query, Optional | manufacture |
| `text` | `*string` | Query, Optional | single entry of hostname/mac |
| `nacruleId` | `*string` | Query, Optional | nacrule_id |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.SearchWiredClient`](../../doc/models/search-wired-client.md)

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

















limit := 100





duration := "10m"

apiResponse, err := sitesClientsWired.SearchSiteWiredClients(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1648529800.8221116,
  "limit": 1000,
  "results": [
    {
      "device_mac": [
        "001122334455"
      ],
      "device_mac_port": [
        {
          "device_mac": "001122334455",
          "ip": "",
          "port_id": "et-0/0/1",
          "port_parent": "",
          "start": "2020-12-10T00:07:36.262+0000",
          "vlan": 1,
          "when": "2022-03-29T04:56:05.172+0000"
        }
      ],
      "ip": [
        "11.216.202.61"
      ],
      "mac": "112233445566",
      "org_id": "c168ddee-c14c-11e5-8e81-1258369c38a9",
      "port_id": [
        "et-0/0/1"
      ],
      "site_id": "c168ddee-c14c-11e5-8e81-1258369c38a9",
      "timestamp": 1571174567.807,
      "vlan": [
        0,
        1001
      ]
    }
  ],
  "start": 1648443400.8221116,
  "total": 1
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

