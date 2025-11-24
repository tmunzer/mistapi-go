# Orgs Events

```go
orgsEvents := client.OrgsEvents()
```

## Class Name

`OrgsEvents`

## Methods

* [Count Org System Events](../../doc/controllers/orgs-events.md#count-org-system-events)
* [Search Org Events](../../doc/controllers/orgs-events.md#search-org-events)
* [Search Org System Events](../../doc/controllers/orgs-events.md#search-org-system-events)


# Count Org System Events

Count Org System Events

```go
CountOrgSystemEvents(
    ctx context.Context,
    orgId uuid.UUID,
    distinct *string,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `distinct` | `*string` | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := "site_id"

limit := 100

duration := "10m"

apiResponse, err := orgsEvents.CountOrgSystemEvents(ctx, orgId, &distinct, &limit, nil, nil, &duration)
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


# Search Org Events

Search Org events

Supported Event Types:

- CRADLEPOINT_SYNC_FAILED
- ORG_CA_CERT_ADDED
- ORG_CA_CERT_REGENERATED

```go
SearchOrgEvents(
    ctx context.Context,
    orgId uuid.UUID,
    mType *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseEventsOrgsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Event type |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsOrgsSearch](../../doc/models/response-events-orgs-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsEvents.SearchOrgEvents(ctx, orgId, nil, &limit, nil, nil, &duration, &sort, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
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


# Search Org System Events

Search Org System Events

```go
SearchOrgSystemEvents(
    ctx context.Context,
    orgId uuid.UUID,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseOrgSystemEventsSearch],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Start time (epoch timestamp in seconds, or relative string like "-1d", "-1w") |
| `end` | `*string` | Query, Optional | End time (epoch timestamp in seconds, or relative string like "-1d", "-2h", "now") |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseOrgSystemEventsSearch](../../doc/models/response-org-system-events-search.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := orgsEvents.SearchOrgSystemEvents(ctx, orgId, &limit, nil, nil, &duration, &sort, nil)
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
  "end": 1741312800,
  "limit": 1000,
  "results": [
    {
      "change_cat": "admin_action",
      "metadata": "{\"\\wlan_id\\\": \\\"None\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Delete WLAN \\\\\"Automation Test\\\\\"\\\"}",
      "org_id": "268e599f-5379-473f-b18b-4384e9b3f244",
      "scope": "org",
      "site_id": "dc47920f-52d5-499c-af72-dcd646764f84",
      "timestamp": 1741309621,
      "type": "delete-wlan"
    },
    {
      "change_cat": "admin_action",
      "metadata": "{\\\"wlan_id\\\": \\\"3855dc19-63fb-4490-b113-0176dc1cc8f4\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Add WLAN \\\\\"Automation Test\\\\\"\\\"}",
      "org_id": "268e599f-5379-473f-b18b-4384e9b3f244",
      "scope": "org",
      "site_id": "dc47920f-52d5-499c-af72-dcd646764f84",
      "timestamp": 1741309601,
      "type": "add-wlan"
    },
    {
      "change_cat": "admin_action",
      "metadata": "{\\\"template_id\\\": \\\"7e49acf4-6841-4e56-ad7d-68d0801cbba8\\\",\\\"admin_name\\\": \\\"API Automation api_automation@mistsys.com\\\",\\\"desc\\\": \\\"Add Template \\\\\"Automation template\\\\\"\\\"}",
      "org_id": "268e599f-5379-473f-b18b-4384e9b3f244",
      "scope": "org",
      "timestamp": 1741309280,
      "type": "add-template"
    }
  ],
  "start": 1741309200,
  "total": 3
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

