# Constants Definitions

```go
constantsDefinitions := client.ConstantsDefinitions()
```

## Class Name

`ConstantsDefinitions`

## Methods

* [List Ap Led Definition](../../doc/controllers/constants-definitions.md#list-ap-led-definition)
* [List App Category Definitions](../../doc/controllers/constants-definitions.md#list-app-category-definitions)
* [List App Sub Category Definitions](../../doc/controllers/constants-definitions.md#list-app-sub-category-definitions)
* [List Applications](../../doc/controllers/constants-definitions.md#list-applications)
* [List Country Codes](../../doc/controllers/constants-definitions.md#list-country-codes)
* [List Gateway Applications](../../doc/controllers/constants-definitions.md#list-gateway-applications)
* [List Traffic Types](../../doc/controllers/constants-definitions.md#list-traffic-types)


# List Ap Led Definition

Get List of AP LED definition

```go
ListApLedDefinition(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstApLed],
    error)
```

## Response Type

[`[]models.ConstApLed`](../../doc/models/const-ap-led.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListApLedDefinition(ctx)
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
    "code": "01",
    "description": "LED not working",
    "key": "LED_FAILURE",
    "name": "LED Failure"
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


# List App Category Definitions

Get List of definitions of all the supported Application Categories. The example field contains an example payload as you would recieve in the alarm webhook output.

```go
ListAppCategoryDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstAppCategoryDefinition],
    error)
```

## Response Type

[`[]models.ConstAppCategoryDefinition`](../../doc/models/const-app-category-definition.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListAppCategoryDefinitions(ctx)
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
    "display": "Images",
    "filters": {
      "srx": [
        "Enhanced_Images_Media",
        "Enhanced_Web_Images",
        "Enhanced_Image_Servers"
      ]
    },
    "key": "Images"
  },
  {
    "display": "Standard",
    "includes": [
      "Adult",
      "FileSharing",
      "Games",
      "Images",
      "Malware",
      "NewsAndReference",
      "Recreation",
      "Religion",
      "Security",
      "Sports",
      "Technology",
      "Violence"
    ],
    "key": "Standard"
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


# List App Sub Category Definitions

Get List of definitions of all the supported Application sub-categories. The example field contains an example payload as you would recieve in the alarm webhook output.

```go
ListAppSubCategoryDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstAppSubcategoryDefinition],
    error)
```

## Response Type

[`[]models.ConstAppSubcategoryDefinition`](../../doc/models/const-app-subcategory-definition.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListAppSubCategoryDefinitions(ctx)
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
    "display": "Office Documents",
    "key": "Office_Documents",
    "traffic_type": "data_interactive"
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


# List Applications

Get List of a list of applications that Juniper-Mist APs recognize

```go
ListApplications(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstApplicationDefinition],
    error)
```

## Response Type

[`[]models.ConstApplicationDefinition`](../../doc/models/const-application-definition.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListApplications(ctx)
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
    "app_id": true,
    "app_image_url": "",
    "app_probe": true,
    "category": "FileSharing",
    "group": "File Sharing",
    "key": "dropbox",
    "name": "Dropbox",
    "signature_based": true,
    "ssr_app_id": true
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


# List Country Codes

Get List of List of available Country Codes

```go
ListCountryCodes(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstCountry],
    error)
```

## Response Type

[`[]models.ConstCountry`](../../doc/models/const-country.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListCountryCodes(ctx)
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
    "alpha2": "FR",
    "certified": true,
    "name": "France",
    "numeric": 250
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


# List Gateway Applications

Get the full list of applications that we recognize

```go
ListGatewayApplications(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstGatewayApplicationsDefinition],
    error)
```

## Response Type

[`[]models.ConstGatewayApplicationsDefinition`](../../doc/models/const-gateway-applications-definition.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListGatewayApplications(ctx)
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
    "app_id": true,
    "key": "4shared",
    "name": "4shared",
    "ssr_app_id": true
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


# List Traffic Types

Get List of identified traffic

```go
ListTrafficTypes(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstTrafficType],
    error)
```

## Response Type

[`[]models.ConstTrafficType`](../../doc/models/const-traffic-type.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListTrafficTypes(ctx)
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
    "display": "VoIP Video",
    "dscp": 32,
    "failover_policy": "non_revertible",
    "max_jitter": 250,
    "max_latency": 1500,
    "max_loss": 35,
    "name": "voip_video",
    "traffic_class": "medium"
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

