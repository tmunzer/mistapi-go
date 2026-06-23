# Sites Devices

```go
sitesDevices := client.SitesDevices()
```

## Class Name

`SitesDevices`

## Methods

* [Add Site Device Image](../../doc/controllers/sites-devices.md#add-site-device-image)
* [Count Site Device Config History](../../doc/controllers/sites-devices.md#count-site-device-config-history)
* [Count Site Device Events](../../doc/controllers/sites-devices.md#count-site-device-events)
* [Count Site Device Last Config](../../doc/controllers/sites-devices.md#count-site-device-last-config)
* [Count Site Devices](../../doc/controllers/sites-devices.md#count-site-devices)
* [Delete Site Device Image](../../doc/controllers/sites-devices.md#delete-site-device-image)
* [Export Site Devices](../../doc/controllers/sites-devices.md#export-site-devices)
* [Get Site Device](../../doc/controllers/sites-devices.md#get-site-device)
* [Import Site Devices](../../doc/controllers/sites-devices.md#import-site-devices)
* [List Site Devices](../../doc/controllers/sites-devices.md#list-site-devices)
* [Search Site Device Config History](../../doc/controllers/sites-devices.md#search-site-device-config-history)
* [Search Site Device Events](../../doc/controllers/sites-devices.md#search-site-device-events)
* [Search Site Device Last Configs](../../doc/controllers/sites-devices.md#search-site-device-last-configs)
* [Search Site Devices](../../doc/controllers/sites-devices.md#search-site-devices)
* [Set Site Devices Gbp Tag](../../doc/controllers/sites-devices.md#set-site-devices-gbp-tag)
* [Update Site Device](../../doc/controllers/sites-devices.md#update-site-device)


# Add Site Device Image

Attach up to 3 images to a device

```go
AddSiteDeviceImage(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    imageNumber int,
    file models.FileWrapper,
    json *string) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | Image file content uploaded as multipart form data |
| `json` | `*string` | Form, Optional | Optional JSON metadata submitted with the image upload |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 110

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })

resp, err := sitesDevices.AddSiteDeviceImage(ctx, siteId, deviceId, imageNumber, file, nil)
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


# Count Site Device Config History

Counts the number of entries in device config history for distinct field with given filters

```go
CountSiteDeviceConfigHistory(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *string,
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
| `distinct` | `*string` | Query, Optional | Field used to group this count response |
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

duration := "10m"

limit := 100

apiResponse, err := sitesDevices.CountSiteDeviceConfigHistory(ctx, siteId, nil, nil, nil, nil, &duration, &limit)
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


# Count Site Device Events

Count device events for a site, optionally grouped by the `distinct` field and filtered by event attributes and time range. Use [Count Org Device Events](../../doc/controllers/orgs-devices.md#count-org-device-events) to count device events across the organization.

```go
CountSiteDeviceEvents(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDeviceEventsCountDistinctEnum,
    model *string,
    mType *string,
    typeCode *string,
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
| `distinct` | [`*models.SiteDeviceEventsCountDistinctEnum`](../../doc/models/site-device-events-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `mac`, `model`, `type`, `type_code`<br><br>**Default**: `"model"` |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `typeCode` | `*string` | Query, Optional | Filter results by event type code |
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

distinct := models.SiteDeviceEventsCountDistinctEnum_MODEL

duration := "10m"

limit := 100

apiResponse, err := sitesDevices.CountSiteDeviceEvents(ctx, siteId, &distinct, nil, nil, nil, nil, nil, &duration, &limit)
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


# Count Site Device Last Config

Counts the number of entries in device config history for distinct field with given filters

```go
CountSiteDeviceLastConfig(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDeviceLastConfigCountDistinctEnum,
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
| `distinct` | [`*models.SiteDeviceLastConfigCountDistinctEnum`](../../doc/models/site-device-last-config-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `mac`, `name`, `site_id`, `version`<br><br>**Default**: `"mac"` |
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

distinct := models.SiteDeviceLastConfigCountDistinctEnum_MAC

duration := "10m"

limit := 100

apiResponse, err := sitesDevices.CountSiteDeviceLastConfig(ctx, siteId, &distinct, nil, nil, &duration, &limit)
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


# Count Site Devices

Count devices for a site, optionally grouped by the `distinct` field and filtered by device inventory attributes and time range. Use [Count Org Devices](../../doc/controllers/orgs-devices.md#count-org-devices) to count devices across the organization.

```go
CountSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    distinct *models.SiteDevicesCountDistinctEnum,
    hostname *string,
    model *string,
    mac *string,
    version *string,
    mxtunnelStatus *string,
    mxedgeId *string,
    lldpSystemName *string,
    lldpSystemDesc *string,
    lldpPortId *string,
    lldpMgmtAddr *string,
    mapId *string,
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
| `distinct` | [`*models.SiteDevicesCountDistinctEnum`](../../doc/models/site-devices-count-distinct-enum.md) | Query, Optional | Field used to group this count response. enum: `hostname`, `lldp_mgmt_addr`, `lldp_port_id`, `lldp_system_desc`, `lldp_system_name`, `map_id`, `model`, `mxedge_id`, `mxtunnel_status`, `version`<br><br>**Default**: `"model"` |
| `hostname` | `*string` | Query, Optional | Filter results by hostname |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `version` | `*string` | Query, Optional | Filter results by software version |
| `mxtunnelStatus` | `*string` | Query, Optional | Filter AP results by Mist Tunnel status |
| `mxedgeId` | `*string` | Query, Optional | Filter results by Mist Edge identifier |
| `lldpSystemName` | `*string` | Query, Optional | Filter AP results by LLDP system name |
| `lldpSystemDesc` | `*string` | Query, Optional | Filter AP results by LLDP system description |
| `lldpPortId` | `*string` | Query, Optional | Filter AP results by LLDP port identifier |
| `lldpMgmtAddr` | `*string` | Query, Optional | Filter AP results by LLDP management IP address |
| `mapId` | `*string` | Query, Optional | Filter results by map identifier |
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

distinct := models.SiteDevicesCountDistinctEnum_MODEL

duration := "10m"

limit := 100

apiResponse, err := sitesDevices.CountSiteDevices(ctx, siteId, &distinct, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &duration, &limit)
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


# Delete Site Device Image

Delete image from a device

```go
DeleteSiteDeviceImage(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    imageNumber int) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `imageNumber` | `int` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

imageNumber := 110

resp, err := sitesDevices.DeleteSiteDeviceImage(ctx, siteId, deviceId, imageNumber)
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


# Export Site Devices

To download the exported device information

```go
ExportSiteDevices(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[[]byte],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []byte.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevices.ExportSiteDevices(ctx, siteId)
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

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Get Site Device

Get Device Configuration

```go
GetSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID) (
    models.ApiResponse[models.MistDevice],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type models.MistDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesDevices.GetSiteDevice(ctx, siteId, deviceId)
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
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceAp: ", *r)
    } else if r, ok := responseBody.AsDeviceSwitch(); ok {
        fmt.Println("Value narrowed down to models.DeviceSwitch: ", *r)
    } else if r, ok := responseBody.AsDeviceGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "aeroscout": {
    "enabled": false,
    "host": "aero.pvt.net",
    "locate_connected": true
  },
  "airista": {
    "enabled": false
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 3,
    "beacon_rate_mode": "custom",
    "beam_disabled": [
      1,
      3,
      6
    ],
    "custom_ble_packet_enabled": false,
    "custom_ble_packet_frame": "0x........",
    "custom_ble_packet_freq_msec": 300,
    "eddystone_uid_adv_power": -65,
    "eddystone_uid_beams": "2-4,7",
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 200,
    "eddystone_uid_instance": "5c5b35000001",
    "eddystone_uid_namespace": "2818e3868dec25629ede",
    "eddystone_url_adv_power": -65,
    "eddystone_url_beams": "2-4,7",
    "eddystone_url_enabled": true,
    "eddystone_url_freq_msec": 1000,
    "eddystone_url_url": "https://www.abc.com",
    "ibeacon_adv_power": -65,
    "ibeacon_beams": "2-4,7",
    "ibeacon_enabled": false,
    "ibeacon_freq_msec": 0,
    "ibeacon_major": 13,
    "ibeacon_minor": 138,
    "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
    "power": 6,
    "power_mode": "custom"
  },
  "centrak": {
    "enabled": false
  },
  "client_bridge": {
    "auth": {
      "psk": "foryoureyesonly",
      "type": "psk"
    },
    "enabled": false,
    "ssid": "Uplink-SSID"
  },
  "created_time": 0,
  "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "esl_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": false,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "for_site": true,
  "height": 2.75,
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6008",
  "image1_url": "string",
  "image2_url": "string",
  "image3_url": "string",
  "iot_config": {
    "A1": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A2": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A3": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A4": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "DI1": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DI2": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DO": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    }
  },
  "ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "dns_suffix": [
      ".mist.local",
      ".mist.com"
    ],
    "gateway": "10.2.1.254",
    "gateway6": "2607:f8b0:4005:808::1",
    "ip": "10.2.1.1",
    "ip6": "2607:f8b0:4005:808::2004",
    "mtu": 1500,
    "netmask": "255.255.255.0",
    "netmask6": "/32",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "led": {
    "brightness": 255,
    "enabled": true
  },
  "locked": true,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "mesh": {
    "enabled": false,
    "group": 1,
    "role": "base"
  },
  "modified_time": 0,
  "name": "conference room",
  "notes": "slightly off center",
  "ntp_servers": [
    "string"
  ],
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "orientation": 45,
  "poe_passthrough": false,
  "pwr_config": {
    "base": 2000,
    "prefer_usb_over_wifi": false
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "type": "ap",
  "uplink_port_config": {
    "dot1x": false,
    "keep_wlans_up_if_down": false
  },
  "usb_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": true,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "x": 53.5,
  "y": 173.1
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


# Import Site Devices

Import Information for Multiple Devices

CSV format:

```csv
mac,name,map_id,x,y,height,orientation,labels,band_24.power,band_24.bandwidth,band_24.channel,band_24.disabled,band_5.power,band_5.bandwidth,band_5.channel,band_5.disabled,band_6.power,band_6.bandwidth,band_6.channel,band_6.disabled
5c5b53010101,"AP 1",845a23bf-bed9-e43c-4c86-6fa474be7ae5,30,10,2.3,45,"guest, campus, vip",1,20,0,false,0,40,0,false,17,80,0,false
```

```go
ImportSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    file models.FileWrapper) (
    models.ApiResponse[[]models.ConfigDevice],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `file` | `models.FileWrapper` | Form, Required | Binary file payload to upload with this request |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []models.ConfigDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

file := getFile("dummy_file", func(err error) { log.Fatalln(err) })

apiResponse, err := sitesDevices.ImportSiteDevices(ctx, siteId, file)
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
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsDeviceAp(); ok {
            fmt.Println("Value narrowed down to models.DeviceAp: ", *i)
        } else if i, ok := item.AsDeviceSwitch(); ok {
            fmt.Println("Value narrowed down to models.DeviceSwitch: ", *i)
        } else if i, ok := item.AsDeviceGateway(); ok {
            fmt.Println("Value narrowed down to models.DeviceGateway: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
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


# List Site Devices

List devices in a site. Use [List Org Devices](../../doc/controllers/orgs-devices.md#list-org-devices) to retrieve devices across the organization, or [Search Org Devices](../../doc/controllers/orgs-devices.md#search-org-devices) when filters are needed.

```go
ListSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    mType *string,
    name *string,
    limit *int,
    page *int) (
    models.ApiResponse[[]models.ConfigDevice],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | `*string` | Query, Optional | Filter results by type<br><br>**Default**: `"ap"` |
| `name` | `*string` | Query, Optional | Filter results by name |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | Select the page number to return when using page-based pagination; starts at `1`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type []models.ConfigDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := "ap"

limit := 100

page := 1

apiResponse, err := sitesDevices.ListSiteDevices(ctx, siteId, &mType, nil, &limit, &page)
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
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsDeviceAp(); ok {
            fmt.Println("Value narrowed down to models.DeviceAp: ", *i)
        } else if i, ok := item.AsDeviceSwitch(); ok {
            fmt.Println("Value narrowed down to models.DeviceSwitch: ", *i)
        } else if i, ok := item.AsDeviceGateway(); ok {
            fmt.Println("Value narrowed down to models.DeviceGateway: ", *i)
        }
    }

    fmt.Println(apiResponse.Response.StatusCode)
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


# Search Site Device Config History

Search for entries in device config history

```go
SearchSiteDeviceConfigHistory(
    ctx context.Context,
    siteId uuid.UUID,
    mType *models.DeviceTypeDefaultApEnum,
    mac *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseConfigHistorySearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | Filter results by type. enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseConfigHistorySearch](../../doc/models/response-config-history-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

mType := models.DeviceTypeDefaultApEnum_AP

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesDevices.SearchSiteDeviceConfigHistory(ctx, siteId, &mType, nil, &limit, nil, nil, &duration, &sort, nil)
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
  "end": 1531862583,
  "limit": 10,
  "results": [
    {
      "channel_24": 11,
      "channel_5": 100,
      "radio_macs": [
        "5c5b352e000a",
        "5c5b352e000b",
        "5c5b352e000c"
      ],
      "radios": [
        {
          "band": "24",
          "channel": 11
        },
        {
          "band": "5",
          "channel": 100
        }
      ],
      "secpolicy_violated": false,
      "ssids": [
        "test24",
        "test5"
      ],
      "ssids_24": [
        "test24"
      ],
      "ssids_5": [
        "test5"
      ],
      "timestamp": 1531855856.643369,
      "version": "apfw-0.2.14754-cersei-75c8",
      "wlans": [
        {
          "auth": "psk",
          "bands": [
            "24"
          ],
          "id": "be22bba7-8e22-e1cf-5185-b880816fe2cf",
          "ssid": "test24",
          "vlan_ids": [
            "1"
          ]
        },
        {
          "auth": "psk",
          "bands": [
            "5"
          ],
          "id": "f8c18724-4118-3487-811a-f98964988604",
          "ssid": "test5",
          "vlan_ids": [
            "1"
          ]
        }
      ]
    }
  ],
  "start": 1531776183,
  "total": 1
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


# Search Site Device Events

Search device events for a site with filters for MAC address, model, event type, message text, and time range. Use [Search Org Device Events](../../doc/controllers/orgs-devices.md#search-org-device-events) to search device events across the organization.

```go
SearchSiteDeviceEvents(
    ctx context.Context,
    siteId uuid.UUID,
    mac *string,
    model *string,
    text *string,
    mType *string,
    lastBy *string,
    includes *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseEventsDevices],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `model` | `*string` | Query, Optional | Filter results by device model |
| `text` | `*string` | Query, Optional | Filter results by event message text |
| `mType` | `*string` | Query, Optional | See [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions) |
| `lastBy` | `*string` | Query, Optional | Return last/recent event for passed in field |
| `includes` | `*string` | Query, Optional | Keyword to include events from additional indices (e.g. ext_tunnel for prisma events) |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseEventsDevices](../../doc/models/response-events-devices.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

text := "Device connected"

lastBy := "port_id"

includes := "ext_tunnel"

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesDevices.SearchSiteDeviceEvents(ctx, siteId, nil, nil, &text, nil, &lastBy, &includes, &limit, nil, nil, &duration, &sort, nil)
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
  "end": 1531862583,
  "limit": 2,
  "next": "/api/v1/sites/8aaba0aa-09cc-44bd-9709-33b98040550c/devices/events/search?ap=5c5b350e0001&end=1531855849.000&limit=2&start=1531776183.0",
  "results": [
    {
      "chassis_mac": "60c78d939c0f",
      "count": 1,
      "device_type": "switch",
      "mac": "60c78d939c0f",
      "model": "EX4100-48MP",
      "org_id": "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f",
      "port_id": "ge-0/0/17",
      "site_id": "978c48e6-6ef6-11e6-8bbf-02e208b2d34f",
      "text": "ifIndex 533, ifAdminStatus up(1), ifOperStatus down(2), ifName ge-0/0/17",
      "timestamp": 1764236687.435,
      "type": "SW_PORT_DOWN",
      "version": "23.4R2-S4.11"
    },
    {
      "ap": "5c5b35d0077b",
      "device_type": "ap",
      "mac": "5c5b35d0077b",
      "model": "AP43",
      "org_id": "9777c1a0-6ef6-11e6-8bbf-02e208b2d34f",
      "site_id": "46fc665e-9706-4296-8fe2-78f42f2e67e4",
      "timestamp": 1764235684.467825,
      "type": "AP_CONFIGURED"
    }
  ],
  "start": 1531776183,
  "total": 14
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


# Search Site Device Last Configs

Search last known device configuration records for a site with filters for device type, MAC address, name, software version, certificate expiry, and time range. Use [Search Org Device Last Configs](../../doc/controllers/orgs-devices.md#search-org-device-last-configs) to search last known device configuration records across the organization.

```go
SearchSiteDeviceLastConfigs(
    ctx context.Context,
    siteId uuid.UUID,
    certExpiryDuration *string,
    deviceType *models.LastConfigDeviceTypeEnum,
    mac *string,
    version *string,
    name *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseConfigHistorySearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `certExpiryDuration` | `*string` | Query, Optional | Duration for expiring cert queries (format: 2d/3h/172800 seconds) |
| `deviceType` | [`*models.LastConfigDeviceTypeEnum`](../../doc/models/last-config-device-type-enum.md) | Query, Optional | Filter results by device type. enum: `ap`, `gateway`, `switch`, `mxedge`<br><br>**Default**: `"ap"` |
| `mac` | `*string` | Query, Optional | Filter results by MAC address |
| `version` | `*string` | Query, Optional | Filter results by software version |
| `name` | `*string` | Query, Optional | Filter results by name |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseConfigHistorySearch](../../doc/models/response-config-history-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

certExpiryDuration := "2d"

deviceType := models.LastConfigDeviceTypeEnum_AP

limit := 100

duration := "10m"

sort := "-site_id"

apiResponse, err := sitesDevices.SearchSiteDeviceLastConfigs(ctx, siteId, &certExpiryDuration, &deviceType, nil, nil, nil, &limit, nil, nil, &duration, &sort, nil)
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
  "end": 1531862583,
  "limit": 10,
  "results": [
    {
      "channel_24": 11,
      "channel_5": 100,
      "radio_macs": [
        "5c5b352e000a",
        "5c5b352e000b",
        "5c5b352e000c"
      ],
      "radios": [
        {
          "band": "24",
          "channel": 11
        },
        {
          "band": "5",
          "channel": 100
        }
      ],
      "secpolicy_violated": false,
      "ssids": [
        "test24",
        "test5"
      ],
      "ssids_24": [
        "test24"
      ],
      "ssids_5": [
        "test5"
      ],
      "timestamp": 1531855856.643369,
      "version": "apfw-0.2.14754-cersei-75c8",
      "wlans": [
        {
          "auth": "psk",
          "bands": [
            "24"
          ],
          "id": "be22bba7-8e22-e1cf-5185-b880816fe2cf",
          "ssid": "test24",
          "vlan_ids": [
            "1"
          ]
        },
        {
          "auth": "psk",
          "bands": [
            "5"
          ],
          "id": "f8c18724-4118-3487-811a-f98964988604",
          "ssid": "test5",
          "vlan_ids": [
            "1"
          ]
        }
      ]
    }
  ],
  "start": 1531776183,
  "total": 1
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


# Search Site Devices

Search devices in a site with filters for device type, identifiers, model, software version, radio settings, LLDP details, and other inventory attributes. Use [Search Org Devices](../../doc/controllers/orgs-devices.md#search-org-devices) to search devices across the organization.

```go
SearchSiteDevices(
    ctx context.Context,
    siteId uuid.UUID,
    band24Channel *int,
    band5Channel *int,
    band6Channel *int,
    band24Bandwidth *int,
    band5Bandwidth *int,
    band6Bandwidth *int,
    band24Power *int,
    band5Power *int,
    band6Power *int,
    clustered *bool,
    eth0PortSpeed *int,
    evpntopoId *uuid.UUID,
    extIp *string,
    hostname *string,
    ip *string,
    lastConfigStatus *string,
    lastHostname *string,
    lldpMgmtAddr *string,
    lldpPortId *string,
    lldpSystemDesc *string,
    lldpSystemName *string,
    mac *string,
    model *string,
    mxedgeId *uuid.UUID,
    mxedgeIds *string,
    mxtunnelStatus *models.SearchSiteDevicesMxtunnelStatusEnum,
    node *models.HaClusterNodeEnum,
    node0Mac *string,
    node1Mac *string,
    powerConstrained *bool,
    radiusStats *string,
    stats *bool,
    t128agentVersion *string,
    mType *models.DeviceTypeDefaultApEnum,
    version *string,
    limit *int,
    start *string,
    end *string,
    duration *string,
    sort *models.SearchSiteDevicesSortEnum,
    descSort *models.SearchSiteDevicesDescSortEnum,
    searchAfter *string) (
    models.ApiResponse[models.ResponseDeviceSearch],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `band24Channel` | `*int` | Query, Optional | When `type`==`ap`, Channel of band_24 |
| `band5Channel` | `*int` | Query, Optional | When `type`==`ap`, Channel of band_5 |
| `band6Channel` | `*int` | Query, Optional | When `type`==`ap`, Channel of band_6 |
| `band24Bandwidth` | `*int` | Query, Optional | When `type`==`ap`, Bandwidth of band_24 |
| `band5Bandwidth` | `*int` | Query, Optional | When `type`==`ap`, Bandwidth of band_5 |
| `band6Bandwidth` | `*int` | Query, Optional | When `type`==`ap`, Bandwidth of band_6 |
| `band24Power` | `*int` | Query, Optional | When `type`==`ap`, Power of band_24 |
| `band5Power` | `*int` | Query, Optional | When `type`==`ap`, Power of band_5 |
| `band6Power` | `*int` | Query, Optional | When `type`==`ap`, Power of band_6 |
| `clustered` | `*bool` | Query, Optional | When `type`==`gateway`, true / false |
| `eth0PortSpeed` | `*int` | Query, Optional | When `type`==`ap`, Port speed of eth0 |
| `evpntopoId` | `*uuid.UUID` | Query, Optional | When `type`==`switch`, EVPN topology id |
| `extIp` | `*string` | Query, Optional | Partial / full Device external ip. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `1.2.3.*` and `*.2.3.*` match `1.2.3.4`). Suffix-only wildcards (e.g. `*.2.3.4`) are not supported |
| `hostname` | `*string` | Query, Optional | Partial / full Device hostname. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `my-london*` and `*london*` match `my-london-1`). Suffix-only wildcards (e.g. `*london-1`) are not supported |
| `ip` | `*string` | Query, Optional | Partial / full Device IP address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `10.100.10.*` and `*100.10.*` match `10.100.10.54`). Suffix-only wildcards (e.g. `*.54`) are not supported |
| `lastConfigStatus` | `*string` | Query, Optional | When `type`==`switch` or `type`==`gateway`, last configuration status |
| `lastHostname` | `*string` | Query, Optional | Last hostname of the device. |
| `lldpMgmtAddr` | `*string` | Query, Optional | When `type`==`ap`, LLDP management IP address |
| `lldpPortId` | `*string` | Query, Optional | When `type`==`ap`, LLDP port id. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `ge-0/0/*` and `*-0/0/*` match `ge-0/0/30`). Suffix-only wildcards (e.g. `*switch-01`) are not supported |
| `lldpSystemDesc` | `*string` | Query, Optional | When `type`==`ap`, LLDP system description. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `Juniper Networks*` and `*Networks*` match `Juniper Networks, Inc.`). Suffix-only wildcards (e.g. `*switch-01`) are not supported |
| `lldpSystemName` | `*string` | Query, Optional | When `type`==`ap`, LLDP system name. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `my-switch*` and `*switch*` match `my-switch-01`). Suffix-only wildcards (e.g. `*switch-01`) are not supported |
| `mac` | `*string` | Query, Optional | Partial / full Device MAC address. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `001122*` and `*1122*` match `001122334455`). Suffix-only wildcards (e.g. `*4455`) are not supported |
| `model` | `*string` | Query, Optional | Partial / full Device model. Use `prefix*` for prefix search or `*substring*` for contains search (e.g. `AP4*` and `*P4*` match `AP43`). Suffix-only wildcards (e.g. `*43`) are not supported |
| `mxedgeId` | `*uuid.UUID` | Query, Optional | When `type`==`ap`, Mist Edge id, if AP is connecting to a Mist Edge |
| `mxedgeIds` | `*string` | Query, Optional | When `type`==`ap`, Comma separated list of Mist Edge id, if AP is connecting to a Mist Edge |
| `mxtunnelStatus` | [`*models.SearchSiteDevicesMxtunnelStatusEnum`](../../doc/models/search-site-devices-mxtunnel-status-enum.md) | Query, Optional | When `type`==`ap`, Mist Tunnel status used to filter results. enum: `down`, `up` |
| `node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Query, Optional | When `type`==`gateway`. enum: `node0`, `node1` |
| `node0Mac` | `*string` | Query, Optional | When `type`==`gateway`, node0 MAC address |
| `node1Mac` | `*string` | Query, Optional | When `type`==`gateway`, node1 MAC address |
| `powerConstrained` | `*bool` | Query, Optional | When `type`==`ap`, whether the AP is power constrained |
| `radiusStats` | `*string` | Query, Optional | When `type`==`switch` or `type`==`gateway`, Key-value pairs where the key<br>is the RADIUS server address and the value contains authentication statistics:<br><br>* <server_address> (string): IP address of the RADIUS server as the key<br>* `auth_accepts` (long): Number of accepted authentication requests<br>* `auth_rejects` (long): Number of rejected authentication requests<br>* `auth_timeouts` (long): Number of authentication timeouts<br>* `auth_server_status` (string): Status of the server. Possible values: `up`, `down`, `unreachable` |
| `stats` | `*bool` | Query, Optional | Whether to return device stats<br><br>**Default**: `false` |
| `t128agentVersion` | `*string` | Query, Optional | When `type`==`gateway` (SSR only), version of 128T agent |
| `mType` | [`*models.DeviceTypeDefaultApEnum`](../../doc/models/device-type-default-ap-enum.md) | Query, Optional | Device type used to filter results. enum: `ap`, `gateway`, `switch`<br><br>**Default**: `"ap"` |
| `version` | `*string` | Query, Optional | Filter results by software version |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `start` | `*string` | Query, Optional | Lower bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d` or `-1w` |
| `end` | `*string` | Query, Optional | Upper bound of the time range, as an epoch timestamp in seconds or a relative value such as `-1d`, `-2h`, or `now` |
| `duration` | `*string` | Query, Optional | Time range duration for the query, using relative units such as `10m`, `7d`, or `2w`<br><br>**Default**: `"1d"` |
| `sort` | [`*models.SearchSiteDevicesSortEnum`](../../doc/models/search-site-devices-sort-enum.md) | Query, Optional | Field used to sort results; a leading `-` indicates descending order. enum: `mac`, `model`, `sku`, `timestamp`<br><br>**Default**: `"timestamp"` |
| `descSort` | [`*models.SearchSiteDevicesDescSortEnum`](../../doc/models/search-site-devices-desc-sort-enum.md) | Query, Optional | Field used to sort results in descending order. enum: `mac`, `model`, `sku`, `timestamp` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseDeviceSearch](../../doc/models/response-device-search.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

eth0PortSpeed := 100

extIp := "1.2.3.4"

hostname := "my-london-1"

ip := "10.100.10.54"

lastConfigStatus := "success"

mac := "aabbccddeeff"

model := "AP43"

stats := false

mType := models.DeviceTypeDefaultApEnum_AP

limit := 100

duration := "10m"

sort := models.SearchSiteDevicesSortEnum_TIMESTAMP

apiResponse, err := sitesDevices.SearchSiteDevices(ctx, siteId, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, &eth0PortSpeed, nil, &extIp, &hostname, &ip, &lastConfigStatus, nil, nil, nil, nil, nil, &mac, &model, nil, nil, nil, nil, nil, nil, nil, nil, &stats, nil, &mType, nil, &limit, nil, nil, &duration, &sort, nil, nil)
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
  "next": "string",
  "results": [
    {
      "hostname": [
        "AP41-STB-3E5299-WH-2001",
        "AP41-STB-3E5299-WH-50",
        "AP41-STB-3E5299",
        "5c5b353e5299"
      ],
      "ip": "10.2.16.205",
      "lldp_mgmt_addr": "10.2.10.139",
      "lldp_port_desc": "GigabitEthernet1/0/1",
      "lldp_port_id": "Gi1/0/1",
      "lldp_system_desc": "Cisco IOS Software, C2960S Software (C2960S-UNIVERSALK9-M), Version 15.2(1)E1, RELEASE SOFTWARE (fc2)\nTechnical Support: https://www.cisco.com/techsupport\nCopyright (c) 1986-2013 by Cisco Systems, Inc.\nCompiled Fri 22-Nov-13 07:10 by prod_rel_team",
      "lldp_system_name": "ME-DC-1-ACC-SW",
      "mac": "5c5b353e5299",
      "model": "AP41",
      "mxedge_id": "00000000-0000-0000-1000-43a81f238391",
      "mxtunnel_status": "down",
      "org_id": "6748cfa6-4e12-11e6-9188-0242ac110007",
      "power_constrained": false,
      "power_opmode": "",
      "site_id": "a8178443-ecb5-461c-b854-f16627619ab3",
      "sku": "AP41-US",
      "timestamp": 1596588619.007,
      "uptime": 85280,
      "version": "0.7.20216",
      "wlans": [
        {
          "id": "28c36fc7-dc22-4960-9d81-34087511c2e5",
          "ssid": "Live-Demo-NAC"
        },
        {
          "id": "51b82e2b-f9e8-470b-a32a-cecde5501b0f",
          "ssid": "Live-Demo"
        }
      ]
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


# Set Site Devices Gbp Tag

Set GBP Tag for multiple devices

```go
SetSiteDevicesGbpTag(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.DevicesGbpTag) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.DevicesGbpTag`](../../doc/models/devices-gbp-tag.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.DevicesGbpTag{
    GbpTag:               220,
    Macs:                 []string{
        "683b679ac024",
    },
}

resp, err := sitesDevices.SetSiteDevicesGbpTag(ctx, siteId, &body)
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


# Update Site Device

Update Device Configuration

```go
UpdateSiteDevice(
    ctx context.Context,
    siteId uuid.UUID,
    deviceId uuid.UUID,
    body *models.MistDevice) (
    models.ApiResponse[models.MistDevice],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `deviceId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MistDevice`](../../doc/models/containers/mist-device.md) | Body, Optional | Mist-managed device object for an AP, switch, or gateway |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type models.MistDevice.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MistDeviceContainer.FromDeviceAp(models.DeviceAp{
    DeviceprofileId:      models.NewOptional(models.ToPointer(uuid.MustParse("6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9"))),
    DisableEth1:          models.ToPointer(false),
    DisableEth2:          models.ToPointer(false),
    DisableEth3:          models.ToPointer(false),
    DisableModule:        models.ToPointer(false),
    FlowControl:          models.ToPointer(false),
    Height:               models.ToPointer(float64(2.75)),
    MapId:                models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc")),
    Name:                 models.ToPointer("conference room"),
    Notes:                models.ToPointer("slightly off center"),
    Orientation:          models.ToPointer(45),
    PoePassthrough:       models.ToPointer(false),
    Type:                 "",
    Vars:                 map[string]string{
        "RADIUS_IP1": "172.31.2.5",
        "RADIUS_SECRET": "11s64632d",
    },
    X:                    models.ToPointer(float64(53.5)),
    Y:                    models.ToPointer(float64(173.1)),
})

apiResponse, err := sitesDevices.UpdateSiteDevice(ctx, siteId, deviceId, &body)
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
    responseBody := apiResponse.Data
    if r, ok := responseBody.AsDeviceAp(); ok {
        fmt.Println("Value narrowed down to models.DeviceAp: ", *r)
    } else if r, ok := responseBody.AsDeviceSwitch(); ok {
        fmt.Println("Value narrowed down to models.DeviceSwitch: ", *r)
    } else if r, ok := responseBody.AsDeviceGateway(); ok {
        fmt.Println("Value narrowed down to models.DeviceGateway: ", *r)
    }

    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "aeroscout": {
    "enabled": false,
    "host": "aero.pvt.net",
    "locate_connected": true
  },
  "airista": {
    "enabled": false
  },
  "ble_config": {
    "beacon_enabled": false,
    "beacon_rate": 3,
    "beacon_rate_mode": "custom",
    "beam_disabled": [
      1,
      3,
      6
    ],
    "custom_ble_packet_enabled": false,
    "custom_ble_packet_frame": "0x........",
    "custom_ble_packet_freq_msec": 300,
    "eddystone_uid_adv_power": -65,
    "eddystone_uid_beams": "2-4,7",
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 200,
    "eddystone_uid_instance": "5c5b35000001",
    "eddystone_uid_namespace": "2818e3868dec25629ede",
    "eddystone_url_adv_power": -65,
    "eddystone_url_beams": "2-4,7",
    "eddystone_url_enabled": true,
    "eddystone_url_freq_msec": 1000,
    "eddystone_url_url": "https://www.abc.com",
    "ibeacon_adv_power": -65,
    "ibeacon_beams": "2-4,7",
    "ibeacon_enabled": false,
    "ibeacon_freq_msec": 0,
    "ibeacon_major": 13,
    "ibeacon_minor": 138,
    "ibeacon_uuid": "f3f17139-704a-f03a-2786-0400279e37c3",
    "power": 6,
    "power_mode": "custom"
  },
  "centrak": {
    "enabled": false
  },
  "client_bridge": {
    "auth": {
      "psk": "foryoureyesonly",
      "type": "psk"
    },
    "enabled": false,
    "ssid": "Uplink-SSID"
  },
  "created_time": 0,
  "deviceprofile_id": "6f4bf402-45f9-2a56-6c8b-7f83d3bc98e9",
  "disable_eth1": false,
  "disable_eth2": false,
  "disable_eth3": false,
  "disable_module": false,
  "esl_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": false,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "for_site": true,
  "height": 2.75,
  "id": "497f6eca-6276-4993-bfeb-53cbbbba6008",
  "image1_url": "string",
  "image2_url": "string",
  "image3_url": "string",
  "iot_config": {
    "A1": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A2": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A3": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "A4": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    },
    "DI1": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DI2": {
      "enabled": false,
      "name": "string",
      "pullup": "internal"
    },
    "DO": {
      "enabled": false,
      "name": "motion",
      "output": true,
      "pullup": "internal",
      "value": 0
    }
  },
  "ip_config": {
    "dns": [
      "8.8.8.8",
      "4.4.4.4"
    ],
    "dns_suffix": [
      ".mist.local",
      ".mist.com"
    ],
    "gateway": "10.2.1.254",
    "gateway6": "2607:f8b0:4005:808::1",
    "ip": "10.2.1.1",
    "ip6": "2607:f8b0:4005:808::2004",
    "mtu": 1500,
    "netmask": "255.255.255.0",
    "netmask6": "/32",
    "type": "static",
    "type6": "static",
    "vlan_id": 1
  },
  "led": {
    "brightness": 255,
    "enabled": true
  },
  "locked": true,
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "mesh": {
    "enabled": false,
    "group": 1,
    "role": "base"
  },
  "modified_time": 0,
  "name": "conference room",
  "notes": "slightly off center",
  "ntp_servers": [
    "string"
  ],
  "org_id": "a40f5d1f-d889-42e9-94ea-b9b33585fc6b",
  "orientation": 45,
  "poe_passthrough": false,
  "pwr_config": {
    "base": 2000,
    "prefer_usb_over_wifi": false
  },
  "site_id": "72771e6a-6f5e-4de4-a5b9-1266c4197811",
  "type": "ap",
  "uplink_port_config": {
    "dot1x": false,
    "keep_wlans_up_if_down": false
  },
  "usb_config": {
    "cacert": "string",
    "channel": 3,
    "enabled": true,
    "host": "1.1.1.1",
    "port": 0,
    "type": "imagotag",
    "verify_cert": true,
    "vlan_id": 1
  },
  "vars": {
    "RADIUS_IP1": "172.31.2.5",
    "RADIUS_SECRET": "11s64632d"
  },
  "x": 53.5,
  "y": 173.1
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

