# Sites Auto Map Assignment

```go
sitesAutoMapAssignment := client.SitesAutoMapAssignment()
```

## Class Name

`SitesAutoMapAssignment`

## Methods

* [Apply Site Auto Map Assignment](../../doc/controllers/sites-auto-map-assignment.md#apply-site-auto-map-assignment)
* [Cancel Site Auto Map Assignment](../../doc/controllers/sites-auto-map-assignment.md#cancel-site-auto-map-assignment)
* [Clear Site Auto Map Assignment](../../doc/controllers/sites-auto-map-assignment.md#clear-site-auto-map-assignment)
* [Get Site Auto Map Assignment Status](../../doc/controllers/sites-auto-map-assignment.md#get-site-auto-map-assignment-status)
* [Start Site Auto Map Assignment](../../doc/controllers/sites-auto-map-assignment.md#start-site-auto-map-assignment)


# Apply Site Auto Map Assignment

Apply (accept) auto map assignment results for a site. Devices are associated with their assigned maps. Omit `map_ids` or provide an empty list to accept all pending assignments; provide specific `map_ids` for a partial accept.

```go
ApplySiteAutoMapAssignment(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AutoMapAssignmentRequest) (
    models.ApiResponse[models.ResponseAutoMapAssignmentApply],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AutoMapAssignmentRequest`](../../doc/models/auto-map-assignment-request.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseAutoMapAssignmentApply](../../doc/models/response-auto-map-assignment-apply.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AutoMapAssignmentRequest{
}

apiResponse, err := sitesAutoMapAssignment.ApplySiteAutoMapAssignment(ctx, siteId, &body)
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
  "accepted_maps": [
    "d3c42998-9012-4859-9743-6b9bee475309",
    "f7a21456-7891-4abc-def0-123456789abc"
  ],
  "message": "Accepted map assignments for map_ids: ['d3c42998-9012-4859-9743-6b9bee475309', 'f7a21456-7891-4abc-def0-123456789abc']"
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


# Cancel Site Auto Map Assignment

Cancel an in-progress auto map assignment operation for the site. Validates that auto map assignment is currently running, notifies all APs to fetch new configuration, and sends a cancel command to the orchestration service.

```go
CancelSiteAutoMapAssignment(
    ctx context.Context,
    siteId uuid.UUID) (
    http.Response,
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := sitesAutoMapAssignment.CancelSiteAutoMapAssignment(ctx, siteId)
if err != nil {
    switch typedErr := err.(type) {
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
| 400 | Auto map assignment not in progress | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Clear Site Auto Map Assignment

Clear (reject) auto map assignment results for a site without applying them. The cached assignment results are cleared. Omit `map_ids` or provide an empty list to reject all pending assignments; provide specific `map_ids` for a partial reject.

```go
ClearSiteAutoMapAssignment(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AutoMapAssignmentRequest) (
    models.ApiResponse[models.ResponseAutoMapAssignmentClear],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AutoMapAssignmentRequest`](../../doc/models/auto-map-assignment-request.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseAutoMapAssignmentClear](../../doc/models/response-auto-map-assignment-clear.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AutoMapAssignmentRequest{
}

apiResponse, err := sitesAutoMapAssignment.ClearSiteAutoMapAssignment(ctx, siteId, &body)
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
  "message": "Rejected map assignments for map_ids: ['d3c42998-9012-4859-9743-6b9bee475309', 'f7a21456-7891-4abc-def0-123456789abc']",
  "rejected_maps": [
    "d3c42998-9012-4859-9743-6b9bee475309",
    "f7a21456-7891-4abc-def0-123456789abc"
  ]
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


# Get Site Auto Map Assignment Status

Get the current status of auto map assignment for the site.

```go
GetSiteAutoMapAssignmentStatus(
    ctx context.Context,
    siteId uuid.UUID) (
    models.ApiResponse[models.ResponseAutoMapAssignmentInfo],
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseAutoMapAssignmentInfo](../../doc/models/response-auto-map-assignment-info.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := sitesAutoMapAssignment.GetSiteAutoMapAssignmentStatus(ctx, siteId)
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
  "est_time_left": 280.5,
  "start_time": 1678900062,
  "status": "in_progress",
  "time_updated": 1678900100
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


# Start Site Auto Map Assignment

Start the auto map assignment process for a site. The service automatically assigns APs to maps based on BLE ranging data and requires at least 3 APs with compatible firmware and model support for BLE.

Repeated POST requests while a site assignment is still running will be rejected.

```go
StartSiteAutoMapAssignment(
    ctx context.Context,
    siteId uuid.UUID,
    body *models.AutoMapAssignment) (
    models.ApiResponse[models.ResponseAutoMapAssignment],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `siteId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AutoMapAssignment`](../../doc/models/auto-map-assignment.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseAutoMapAssignment](../../doc/models/response-auto-map-assignment.md).

## Example Usage

```go
ctx := context.Background()

siteId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AutoMapAssignment{
    Dryrun:               models.ToPointer(false),
    ForceCollection:      models.ToPointer(false),
}

apiResponse, err := sitesAutoMapAssignment.StartSiteAutoMapAssignment(ctx, siteId, &body)
if err != nil {
    switch typedErr := err.(type) {
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
  "devices": {
    "5c5b35000001": {
      "reason": "Device meets the minimum requirements for auto map assignment",
      "valid": true
    },
    "5c5b35000002": {
      "reason": "Device meets the minimum requirements for auto map assignment",
      "valid": true
    },
    "5c5b35000003": {
      "reason": "Device meets the minimum requirements for auto map assignment",
      "valid": true
    }
  },
  "estimated_runtime": 300,
  "reason": "Started auto map assignment",
  "started": true,
  "valid": true
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Auto map assignment already in progress | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

