# Orgs Stats-Sites

```go
orgsStatsSites := client.OrgsStatsSites()
```

## Class Name

`OrgsStatsSites`


# List Org Site Stats

Get List of Org Site Stats

```go
ListOrgSiteStats(
    ctx context.Context,
    orgId uuid.UUID,
    start *int,
    end *int,
    duration *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.StatsSite],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `start` | `*int` | Query, Optional | Start datetime, can be epoch or relative time like -1d, -1w; -1d if not specified |
| `end` | `*int` | Query, Optional | End datetime, can be epoch or relative time like -1d, -2h; now if not specified |
| `duration` | `*string` | Query, Optional | Duration like 7d, 2w<br>**Default**: `"1d"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`[]models.StatsSite`](../../doc/models/stats-site.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





duration := "10m"

limit := 100

page := 1

apiResponse, err := orgsStatsSites.ListOrgSiteStats(ctx, orgId, nil, nil, &duration, &limit, &page)
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
    "address": "1601 S De Anza Blvd, Cupertino, CA 95014, USA",
    "alarmtemplate_id": null,
    "analyticEnabled": true,
    "aptemplate_id": null,
    "country_code": "US",
    "created_time": 1472591606,
    "engagementEnabled": true,
    "gatewaytemplate_id": "e571f2a2-d748-4ad4-bd6c-895467957c21",
    "id": "83bc290a-b76d-47fa-a294-d34e47f30f7f",
    "lat": 37.295553,
    "latlng": {
      "lat": 37.295553,
      "lng": -122.033007
    },
    "lng": -122.033007,
    "modified_time": 1728057857,
    "msp_id": "a9af4951-a1de-4520-b398-c95a58947349",
    "name": "Live-Demo",
    "networktemplate_id": "964cb213-deb2-469d-8c1e-a5f8661c6886",
    "notes": "This site is used for demonstration purposes.",
    "num_ap": 17,
    "num_ap_connected": 14,
    "num_clients": 14,
    "num_devices": 26,
    "num_devices_connected": 22,
    "num_gateway": 1,
    "num_gateway_connected": 1,
    "num_switch": 8,
    "num_switch_connected": 7,
    "org_id": "b9814b40-ac4b-4424-86a8-b787eb68b86a",
    "rftemplate_id": "2c134c07-3c57-46b3-a53b-8aea92ed7234",
    "secpolicy_id": null,
    "sitegroup_ids": [
      "5644a432-eea9-4a2f-a30a-ddaf4dbc79cf",
      "5fc0f305-f626-49db-8869-10b87f201bba",
      "882796ef-190b-405e-98ef-cb487140cf64"
    ],
    "sitetemplate_id": null,
    "timezone": "America/Los_Angeles",
    "tzoffset": 960
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

