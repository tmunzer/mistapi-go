# Orgs Clients-Wired

```go
orgsClientsWired := client.OrgsClientsWired()
```

## Class Name

`OrgsClientsWired`

## Methods

* [Count Org Wired Clients](../../doc/controllers/orgs-clients-wired.md#count-org-wired-clients)
* [Search Org Wired Clients](../../doc/controllers/orgs-clients-wired.md#search-org-wired-clients)


# Count Org Wired Clients

Count by Distinct Attributes of Clients

Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

```go
CountOrgWiredClients(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *models.OrgWiredClientsCountDistinctEnum,
    page *int,
    limit *int,
    start *int,
    end *int,
    duration *string) (
    models.ApiResponse[models.RepsonseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.OrgWiredClientsCountDistinctEnum`](../../doc/models/org-wired-clients-count-distinct-enum.md) | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | - |
| `start` | `*int` | Query, Optional | start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | end datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | duration like 7d, 2w |

## Response Type

[`models.RepsonseCount`](../../doc/models/repsonse-count.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.OrgWiredClientsCountDistinctEnum("mac")

page := 1

limit := 100





duration := "10m"

apiResponse, err := orgsClientsWired.CountOrgWiredClients(ctx, orgId, &distinct, &page, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Search Org Wired Clients

Search for Wired Clients in org

Note: For list of avaialable `type` values, please refer to [listClientEventsDefinitions]($e/Constants%20Events/listClientEventsDefinitions)

```go
SearchOrgWiredClients(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    deviceMac *string,
    mac *string,
    portId *string,
    vlan *int,
    ip *string,
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
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Site ID |
| `deviceMac` | `*string` | Query, Optional | device mac |
| `mac` | `*string` | Query, Optional | client mac |
| `portId` | `*string` | Query, Optional | port id |
| `vlan` | `*int` | Query, Optional | vlan |
| `ip` | `*string` | Query, Optional | ip |
| `manufacture` | `*string` | Query, Optional | client manufacturer |
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

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



















limit := 100





duration := "10m"

apiResponse, err := orgsClientsWired.SearchOrgWiredClients(ctx, orgId, nil, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

