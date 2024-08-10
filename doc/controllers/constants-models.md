# Constants Models

```go
constantsModels := client.ConstantsModels()
```

## Class Name

`ConstantsModels`

## Methods

* [List Device Models](../../doc/controllers/constants-models.md#list-device-models)
* [List Mx Edge Models](../../doc/controllers/constants-models.md#list-mx-edge-models)
* [List Supported Other Device Models](../../doc/controllers/constants-models.md#list-supported-other-device-models)


# List Device Models

Get list of AP device models for the Mist Site

```go
ListDeviceModels(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstDeviceModel],
    error)
```

## Response Type

[`[]models.ConstDeviceModel`](../../doc/models/containers/const-device-model.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsModels.ListDeviceModels(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    responseBody := apiResponse.Data
    for _, item := range responseBody {
        if i, ok := item.AsConstDeviceAp(); ok {
            fmt.Println("Value narrowed down to models.ConstDeviceAp: ", *i)
        } else if i, ok := item.AsConstDeviceSwitch(); ok {
            fmt.Println("Value narrowed down to models.ConstDeviceSwitch: ", *i)
        } else if i, ok := item.AsConstDeviceGateway(); ok {
            fmt.Println("Value narrowed down to models.ConstDeviceGateway: ", *i)
        } else if i, ok := item.AsConstDeviceUnknown(); ok {
            fmt.Println("Value narrowed down to models.ConstDeviceUnknown: ", *i)
        }
    }

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


# List Mx Edge Models

Get List of available Mx Edge models

```go
ListMxEdgeModels(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstMxedgeModel],
    error)
```

## Response Type

[`[]models.ConstMxedgeModel`](../../doc/models/const-mxedge-model.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsModels.ListMxEdgeModels(ctx)
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
    "display": "X10",
    "model": "ME-X10",
    "ports": {
      "0": {
        "display": "xe0",
        "speed": 10000
      },
      "1": {
        "display": "xe1",
        "speed": 10000
      },
      "2": {
        "display": "xe2",
        "speed": 10000
      },
      "3": {
        "display": "xe3",
        "speed": 10000
      }
    }
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


# List Supported Other Device Models

Supported OtherDevice Models

```go
ListSupportedOtherDeviceModels(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstOtherDeviceModel],
    error)
```

## Response Type

[`[]models.ConstOtherDeviceModel`](../../doc/models/const-other-device-model.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsModels.ListSupportedOtherDeviceModels(ctx)
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
    "_vendor_model_id": "65",
    "display": "W1850",
    "model": "W1850",
    "type": "router",
    "vendor": "cradlepoint"
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

