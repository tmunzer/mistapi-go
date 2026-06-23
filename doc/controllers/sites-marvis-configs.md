# Sites Marvis Configs

```go
sitesMarvisConfigs := client.SitesMarvisConfigs()
```

## Class Name

`SitesMarvisConfigs`

## Methods

* [Count Site Marvis Config Actions](../../doc/controllers/sites-marvis-configs.md#count-site-marvis-config-actions)
* [Delete Site Marvis Config Action](../../doc/controllers/sites-marvis-configs.md#delete-site-marvis-config-action)
* [Search Site Marvis Config Actions](../../doc/controllers/sites-marvis-configs.md#search-site-marvis-config-actions)
* [Submit Site Marvis Config Feedback](../../doc/controllers/sites-marvis-configs.md#submit-site-marvis-config-feedback)


# Count Site Marvis Config Actions

Count Marvis Config Actions for a site by a distinct field.

```go
CountSiteMarvisConfigActions(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *string,
    mac *string,
    mType *string,
    src *string,
    adminId *string,
    op *string,
    portId *string,
    vlanIds *int,
    reason *string,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.ResponseCount],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `distinct` | `*string` | Query, Optional | Field to count by. enum: `mac`, `type`, `src`, `admin_id`, `op`, `port_id`, `reason`, `vlan_ids`<br><br>**Default**: `"mac"` |
| `mac` | `*string` | Query, Optional | Filter by device MAC address |
| `mType` | `*string` | Query, Optional | Filter by config type (e.g. wired) |
| `src` | `*string` | Query, Optional | Filter by source of the config action (e.g. marvis) |
| `adminId` | `*string` | Query, Optional | Filter by admin ID |
| `op` | `*string` | Query, Optional | Filter by operation type (e.g. disable_port, enable_port, update_mtu, add_vlans_to_port) |
| `portId` | `*string` | Query, Optional | Filter by port identifier (e.g. ge-0/0/13) |
| `vlanIds` | `*int` | Query, Optional | Filter by VLAN ID |
| `reason` | `*string` | Query, Optional | Filter by reason for the config action (e.g. rogue_dhcp_server_detected) |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |

## Response Type

**200**: Count result

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseCount](../../doc/models/response-count.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := "mac"

limit := 100

duration := "10m"

apiResponse, err := sitesMarvisConfigs.CountSiteMarvisConfigActions(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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


# Delete Site Marvis Config Action

Delete a Marvis Config Action.

```go
DeleteSiteMarvisConfigAction(
    ctx context.Context,
    siteId uuid.UUID,
    id uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `id` | `uuid.UUID` | Template, Required | UUID of the Marvis Config Action |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

id := uuid.MustParse("00001770-0000-0000-0000-000000000000")

resp, err := sitesMarvisConfigs.DeleteSiteMarvisConfigAction(ctx, siteId, id)
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
    fmt.Println(resp.StatusCode)
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


# Search Site Marvis Config Actions

Search Marvis Config Actions for a site.

```go
SearchSiteMarvisConfigActions(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    mType *string,
    src *string,
    adminId *string,
    op *string,
    portId *string,
    vlanIds *int,
    reason *string,
    limit *int,
    start *string,
    end *string,
    duration *string) (
    models.ApiResponse[models.MarvisConfigActionsSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Filter by device MAC address |
| `mType` | `*string` | Query, Optional | Filter by config type (e.g. wired) |
| `src` | `*string` | Query, Optional | Filter by source of the config action (e.g. marvis) |
| `adminId` | `*string` | Query, Optional | Filter by admin ID |
| `op` | `*string` | Query, Optional | Filter by operation type (e.g. disable_port, enable_port, update_mtu, add_vlans_to_port) |
| `portId` | `*string` | Query, Optional | Filter by port identifier (e.g. ge-0/0/13) |
| `vlanIds` | `*int` | Query, Optional | Filter by VLAN ID |
| `reason` | `*string` | Query, Optional | Filter by reason for the config action (e.g. rogue_dhcp_server_detected) |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |

## Response Type

**200**: Paginated Marvis Config Actions search results

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.MarvisConfigActionsSearch](../../doc/models/marvis-config-actions-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

duration := "10m"

apiResponse, err := sitesMarvisConfigs.SearchSiteMarvisConfigActions(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, &limit, nil, nil, &duration)
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
  "end": 1775122221,
  "limit": 10,
  "results": [
    {
      "admin_id": "6d617276-0000-0000-3157-000000000000",
      "id": "05b46288-37d4-4860-9de4-1edc6e8d5363",
      "mac": "f8c1165aba00",
      "op": "disable_port",
      "org_id": "174260d5-cb22-4ea8-badb-c77a89acb0a9",
      "port_id": "ge-0/0/2",
      "reason": "rogue_dhcp_server_detected",
      "site_id": "437ac5f0-fc76-4a2b-87ab-d8d1e5c00405",
      "src": "marvis",
      "timestamp": 1775028130.405962,
      "type": "wired",
      "vlan_ids": []
    },
    {
      "admin_id": "6d617276-0000-0000-3157-000000000000",
      "id": "b1a81ed4-a7a2-4945-b01d-f54afe6d5cc4",
      "mac": "f8c1165aba00",
      "op": "add_vlans_to_port",
      "org_id": "174260d5-cb22-4ea8-badb-c77a89acb0a9",
      "port_id": "ge-0/0/12",
      "reason": "missing_vlans",
      "site_id": "437ac5f0-fc76-4a2b-87ab-d8d1e5c00405",
      "src": "marvis",
      "timestamp": 1774866716.15723,
      "type": "wired",
      "vlan_ids": [
        100,
        200
      ]
    }
  ],
  "start": 1775118621,
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


# Submit Site Marvis Config Feedback

Submit feedback on a Marvis-injected config action (e.g. mark as invalid).

```go
SubmitSiteMarvisConfigFeedback(
    ctx context.Context,
    siteId uuid.UUID,
    id uuid.UUID,
    body models.MarvisConfigFeedback) (
    models.ApiResponse[models.MarvisConfigFeedbackResponse],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `id` | `uuid.UUID` | Template, Required | UUID of the Marvis Config Action |
| `body` | [`models.MarvisConfigFeedback`](../../doc/models/marvis-config-feedback.md) | Body, Required | Request Body |

## Response Type

**200**: Marvis Config Feedback response

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.MarvisConfigFeedbackResponse](../../doc/models/marvis-config-feedback-response.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

id := uuid.MustParse("00001770-0000-0000-0000-000000000000")

body := models.MarvisConfigFeedback{
}

apiResponse, err := sitesMarvisConfigs.SubmitSiteMarvisConfigFeedback(ctx, siteId, id, body)
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
  "feedback_note": "this port config is intended, do not change anymore",
  "feedback_type": "invalid"
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

