# Constants Definitions

```go
constantsDefinitions := client.ConstantsDefinitions()
```

## Class Name

`ConstantsDefinitions`

## Methods

* [List Alarm Definitions](../../doc/controllers/constants-definitions.md#list-alarm-definitions)
* [List Ap Led Definition](../../doc/controllers/constants-definitions.md#list-ap-led-definition)
* [List App Category Definitions](../../doc/controllers/constants-definitions.md#list-app-category-definitions)
* [List App Sub Category Definitions](../../doc/controllers/constants-definitions.md#list-app-sub-category-definitions)


# List Alarm Definitions

Get List of brief definitions of all the supported alarm types.

The example field contains an example payload as you would recieve in the alarm webhook output.

HA cluster node names will be specified in the `node` field, if applicable.'

```go
ListAlarmDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstAlarmDefinition],
    error)
```

## Response Type

[`[]models.ConstAlarmDefinition`](../../doc/models/const-alarm-definition.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListAlarmDefinitions(ctx)
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
    "display": "Device offline",
    "example": {
      "aps": [
        "d420b02000fa"
      ],
      "count": 1,
      "group": "infrastructure",
      "hostnames": [
        "Vendor_AP2"
      ],
      "id": "e70c308f-7007-4866-9ecd-0d01842979ea",
      "last_seen": 1629753888,
      "org_id": "09dac91f-6e73-4100-89f7-698e0fafbb1b",
      "severity": "warn",
      "site_id": "dcfb31a1-d615-4361-8c95-b9dde05aa704",
      "timestamp": 1629753888,
      "type": "device_down"
    },
    "fields": [
      "aps",
      "hostnames"
    ],
    "group": "infrastructure",
    "key": "device_down",
    "marvis_suggestion_category": "string",
    "severity": "warn"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

