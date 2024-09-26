# Constants Events

```go
constantsEvents := client.ConstantsEvents()
```

## Class Name

`ConstantsEvents`

## Methods

* [List Client Events Definitions](../../doc/controllers/constants-events.md#list-client-events-definitions)
* [List Device Events Definitions](../../doc/controllers/constants-events.md#list-device-events-definitions)
* [List Mx Edge Events Definitions](../../doc/controllers/constants-events.md#list-mx-edge-events-definitions)
* [List Nac Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions)
* [List Other Device Events Definitions](../../doc/controllers/constants-events.md#list-other-device-events-definitions)
* [List System Events Definitions](../../doc/controllers/constants-events.md#list-system-events-definitions)


# List Client Events Definitions

Get List of List of available Client Events

```go
ListClientEventsDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstEvent],
    error)
```

## Response Type

[`[]models.ConstEvent`](../../doc/models/const-event.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsEvents.ListClientEventsDefinitions(ctx)
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
    "display": "11r Association",
    "key": "CLIENT_AUTH_ASSOCIATION_11R"
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


# List Device Events Definitions

Get list of available Device Events

```go
ListDeviceEventsDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstEvent],
    error)
```

## Response Type

[`[]models.ConstEvent`](../../doc/models/const-event.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsEvents.ListDeviceEventsDefinitions(ctx)
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
    "description": "AP was assigned to a site",
    "display": "AP Assigned",
    "example": {
      "ap": "5c5b35000001",
      "audit_id": "e9a88814-fa81-5bdc-34b0-84e8735420e5",
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "timestamp": 1552408871,
      "type": "AP_ASSIGNED"
    },
    "key": "AP_ASSIGNED"
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


# List Mx Edge Events Definitions

Get List of available MX Edge Events

```go
ListMxEdgeEventsDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstEvent],
    error)
```

## Response Type

[`[]models.ConstEvent`](../../doc/models/const-event.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsEvents.ListMxEdgeEventsDefinitions(ctx)
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
    "description": "Config change on ME was triggered as a result of change made by user",
    "display": "ME Config changed by user",
    "example": {
      "audit_id": "e9a88814-fa81-5bdc-34b0-84e8735420e5",
      "mxcluster_id": "ed4665ed-c9ad-4835-8ca5-dda642765ad3",
      "mxedge_id": "387804a7-3474-85ce-15a2-f9a9684c9c9",
      "org_id": "2818e386-8dec-2562-9ede-5b8a0fbbdc71",
      "service": "mxagent",
      "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
      "timestamp": 1552408871,
      "type": "ME_CONFIG_CHANGED_BY_USER"
    },
    "key": "ME_CONFIG_CHANGED_BY_USER"
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


# List Nac Events Definitions

Get List of List of available NAC Client Events

```go
ListNacEventsDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstNacEvent],
    error)
```

## Response Type

[`[]models.ConstNacEvent`](../../doc/models/const-nac-event.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsEvents.ListNacEventsDefinitions(ctx)
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
    "ap": "5c5b355008c0",
    "bssid": "5c5b35548892",
    "cert_cn": "suriyas",
    "cert_expiry": 1711557441,
    "cert_issuer": "/DC=net/DC=jnpr/CN=Juniper Networks Issuing AWS1 CA",
    "cert_san_upn": [
      "suriyas@juniper.net"
    ],
    "cert_serial": "1300103d29e56ef083797bedc2000100103d29",
    "cert_subject": "/CN=suriyas/emailAddress=suriyas@juniper.net",
    "eap_type": "EAP-TLS",
    "nas_vendor": "Mist",
    "org_id": "94de66e8-556a-4d56-8780-a114620a5c42",
    "random_mac": true,
    "site_id": "b5a005ab-47d4-41f7-97bf-733f9cc252dd",
    "ssid": "Test_Suriya-SSID",
    "timestamp": 1685658478.438995,
    "type": "NAC_CLIENT_CERT_CHECK_SUCCESS",
    "username": "suriyas@juniper.net",
    "wcid": "b43637b0-f0d9-0a1d-1ec2-73c394a9f679"
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


# List Other Device Events Definitions

Supported Events Type

```go
ListOtherDeviceEventsDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstEvent],
    error)
```

## Response Type

[`[]models.ConstEvent`](../../doc/models/const-event.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsEvents.ListOtherDeviceEventsDefinitions(ctx)
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
    "display": "Connected to NCM",
    "example": {
      "device_mac": "5c5b351e13b5",
      "mac": "0030447771c0",
      "org_id": "c080ce4d-4e35-4373-bdc4-08df15d257f5",
      "site_id": "1df889ad-9111-4c0e-a00b-8a008b83eb68",
      "text": "Connected to NCM",
      "timestamp": 1675827825.765,
      "type": "CELLULAR_EDGE_CONNECTED_TO_NCM",
      "vendor": "cradlepoint"
    },
    "key": "CELLULAR_EDGE_CONNECTED_TO_NCM"
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


# List System Events Definitions

Get List of List of available System Events

```go
ListSystemEventsDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstEvent],
    error)
```

## Response Type

[`[]models.ConstEvent`](../../doc/models/const-event.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsEvents.ListSystemEventsDefinitions(ctx)
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
    "display": "AP Disconnect",
    "group": "ap_health",
    "key": "ap_disconnected"
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

