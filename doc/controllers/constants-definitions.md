# Constants Definitions

```go
constantsDefinitions := client.ConstantsDefinitions()
```

## Class Name

`ConstantsDefinitions`

## Methods

* [List Ap Channels](../../doc/controllers/constants-definitions.md#list-ap-channels)
* [List Ap L Esl Versions](../../doc/controllers/constants-definitions.md#list-ap-l-esl-versions)
* [List Ap Led Definition](../../doc/controllers/constants-definitions.md#list-ap-led-definition)
* [List App Category Definitions](../../doc/controllers/constants-definitions.md#list-app-category-definitions)
* [List App Sub Category Definitions](../../doc/controllers/constants-definitions.md#list-app-sub-category-definitions)
* [List Applications](../../doc/controllers/constants-definitions.md#list-applications)
* [List Country Codes](../../doc/controllers/constants-definitions.md#list-country-codes)
* [List Fingerprint Types](../../doc/controllers/constants-definitions.md#list-fingerprint-types)
* [List Gateway Applications](../../doc/controllers/constants-definitions.md#list-gateway-applications)
* [List Insight Metrics](../../doc/controllers/constants-definitions.md#list-insight-metrics)
* [List License Types](../../doc/controllers/constants-definitions.md#list-license-types)
* [List Marvis Client Versions](../../doc/controllers/constants-definitions.md#list-marvis-client-versions)
* [List Site Languages](../../doc/controllers/constants-definitions.md#list-site-languages)
* [List States](../../doc/controllers/constants-definitions.md#list-states)
* [List Traffic Types](../../doc/controllers/constants-definitions.md#list-traffic-types)
* [List Webhook Topics](../../doc/controllers/constants-definitions.md#list-webhook-topics)


# List Ap Channels

Get List of List of Available channels per country code

```go
ListApChannels(
    ctx context.Context,
    countryCode *string) (
    models.ApiResponse[models.ConstApChannel],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `countryCode` | `*string` | Query, Optional | Country code, in two-character<br><br>**Constraints**: *Pattern*: `^[a-zA-Z]{2}$` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ConstApChannel](../../doc/models/const-ap-channel.md).

## Example Usage

```go
ctx := context.Background()

countryCode := "US"

apiResponse, err := constantsDefinitions.ListApChannels(ctx, &countryCode)
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
  "band24_40mhz_allowed": true,
  "band24_channels": {
    "20": [
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      10,
      11
    ],
    "40": [
      1,
      2,
      3,
      4,
      5,
      6,
      7,
      8,
      9,
      10,
      11
    ]
  },
  "band24_enabled": true,
  "band5_channels": {
    "20": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161,
      165
    ],
    "40": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161
    ],
    "80": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161
    ],
    "dfs": [
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144
    ],
    "outdoor": [
      36,
      40,
      44,
      48,
      52,
      56,
      60,
      64,
      100,
      104,
      108,
      112,
      116,
      120,
      124,
      128,
      132,
      136,
      140,
      144,
      149,
      153,
      157,
      161,
      165
    ]
  },
  "band5_enabled": true,
  "band6_channels": {
    "160": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221
    ],
    "20": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221,
      225,
      229,
      233
    ],
    "40": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221,
      225,
      229
    ],
    "80": [
      1,
      5,
      9,
      13,
      17,
      21,
      25,
      29,
      33,
      37,
      41,
      45,
      49,
      53,
      57,
      61,
      65,
      69,
      73,
      77,
      81,
      85,
      89,
      93,
      97,
      101,
      105,
      109,
      113,
      117,
      121,
      125,
      129,
      133,
      137,
      141,
      145,
      149,
      153,
      157,
      161,
      165,
      169,
      173,
      177,
      181,
      185,
      189,
      193,
      197,
      201,
      205,
      209,
      213,
      217,
      221
    ],
    "psc": [
      5,
      21,
      37,
      53,
      69,
      85,
      101,
      117,
      133,
      149,
      165,
      181,
      197,
      213,
      229
    ]
  },
  "band6_enabled": true,
  "certified": true,
  "code": 840,
  "dfs_ok": true,
  "key": "US",
  "name": "United States",
  "uses": "US_FCC"
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


# List Ap L Esl Versions

Get Available AP ESL Versions

```go
ListApLEslVersions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstApEslVersion],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstApEslVersion](../../doc/models/const-ap-esl-version.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListApLEslVersions(ctx)
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
    "esl_version": "2.5.1",
    "model": "AP34"
  },
  {
    "esl_version": "2.5.0",
    "model": "AP43"
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


# List Ap Led Definition

Get List of AP LED definition

```go
ListApLedDefinition(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstApLed],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstApLed](../../doc/models/const-ap-led.md).

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

Get List of definitions of all the supported Application Categories. The example field contains an example payload as you would receive in the alarm webhook output.

```go
ListAppCategoryDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstAppCategoryDefinition],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstAppCategoryDefinition](../../doc/models/const-app-category-definition.md).

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

Get List of definitions of all the supported Application sub-categories. The example field contains an example payload as you would receive in the alarm webhook output.

```go
ListAppSubCategoryDefinitions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstAppSubcategoryDefinition],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstAppSubcategoryDefinition](../../doc/models/const-app-subcategory-definition.md).

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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstApplicationDefinition](../../doc/models/const-application-definition.md).

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

Get List of available Country Codes

```go
ListCountryCodes(
    ctx context.Context,
    extend *bool) (
    models.ApiResponse[[]models.ConstCountry],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `extend` | `*bool` | Query, Optional | Will include more country codes if true<br><br>**Default**: `false` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstCountry](../../doc/models/const-country.md).

## Example Usage

```go
ctx := context.Background()

extend := false

apiResponse, err := constantsDefinitions.ListCountryCodes(ctx, &extend)
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


# List Fingerprint Types

Get List of supported fingerprint attribute values

* family
* model
* mfg
* os_type

This information can be used in the [Mist NAC Rules]($h/Orgs%20NAC%20Rules/_overview) `matching` attribute.

```go
ListFingerprintTypes(
    ctx context.Context) (
    models.ApiResponse[models.ConstFingerprintTypes],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ConstFingerprintTypes](../../doc/models/const-fingerprint-types.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListFingerprintTypes(ctx)
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
  "family": [
    "2Wire Router",
    "3Com Switches",
    "ACTi Corporation Network Camera",
    "APC Video Equipment",
    "APC-Schneider UPS",
    "Aastra VoIP",
    "Acer",
    "Actiontec Wireless Router",
    "Aerohive Access Point",
    "Alcatel",
    "Alcatel VoIP",
    "Amazon Echo"
  ],
  "mfg": [
    "100fio Networks Technology llc",
    "10NET COMMUNICATIONS/DCA",
    "11wave Technonlogy Co.,Ltd",
    "12Sided Technology, LLC",
    "1Net Corporation",
    "1Verge Internet Technology (Beijing) Co., Ltd."
  ],
  "model": [
    "10T Lite",
    "10T Pro",
    "10th Gen",
    "11 Lite",
    "11 Pro",
    "11 Pro Max"
  ],
  "os": [
    "Android",
    "Apple OS",
    "Asha Platform OS"
  ]
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


# List Gateway Applications

Get the full list of applications that we recognize

```go
ListGatewayApplications(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstGatewayApplicationsDefinition],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstGatewayApplicationsDefinition](../../doc/models/const-gateway-applications-definition.md).

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


# List Insight Metrics

List Insight Metrics

```go
ListInsightMetrics(
    ctx context.Context) (
    models.ApiResponse[map[string]models.ConstInsightMetricsProperty],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [map[string]models.ConstInsightMetricsProperty](../../doc/models/const-insight-metrics-property.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListInsightMetrics(ctx)
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
  "bytes": {
    "description": "Aggregated bytes over time",
    "example": [
      185,
      197,
      250
    ],
    "intervals": {
      "10m": {
        "interval": 600,
        "max_age": 86400
      },
      "1h": {
        "interval": 3600,
        "max_age": 1209600
      }
    },
    "report_durations": {
      "1d": {
        "duration": 86400,
        "interval": 3600
      },
      "1w": {
        "duration": 604800,
        "interval": 3600
      }
    },
    "report_scopes": [
      "site",
      "org"
    ],
    "scopes": [
      "site",
      "ap",
      "client"
    ],
    "type": "timeseries",
    "unit": "byte"
  },
  "num_clients": {
    "description": "Number of client over time",
    "example": [
      18,
      null,
      15
    ],
    "intervals": {
      "10m": {
        "interval": 600,
        "max_age": 86400
      },
      "1h": {
        "interval": 3600,
        "max_age": 1209600
      }
    },
    "report_durations": {
      "1d": {
        "duration": 86400,
        "interval": 3600
      },
      "1w": {
        "duration": 604800,
        "interval": 3600
      }
    },
    "report_scopes": [
      "site",
      "org"
    ],
    "scopes": [
      "site",
      "ap",
      "device"
    ],
    "type": "timeseries",
    "unit": ""
  }
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


# List License Types

Get License Types

```go
ListLicenseTypes(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstLicenseType],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstLicenseType](../../doc/models/const-license-type.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListLicenseTypes(ctx)
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
    "description": "Wired Assurance 12",
    "includes": [
      "sub_ex12a",
      "sub_ex12p"
    ],
    "key": "sub_ex12",
    "name": "SUB-EX12"
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


# List Marvis Client Versions

Get List of the available Marvis Client Versions.

```go
ListMarvisClientVersions(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstMarvisClientVersion],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstMarvisClientVersion](../../doc/models/const-marvis-client-version.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListMarvisClientVersions(ctx)
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
    "label": "default",
    "notes": "",
    "os": "android",
    "url": "https://mobile.mist.com/installers/marvisclient/android/1.1.9/marvisclient-installer.apk",
    "version": "1.1.9"
  },
  {
    "label": "default",
    "notes": "",
    "os": "macos",
    "url": "https://mobile.mist.com/installers/marvisclient/macos/0.100.29/marvisclient-installer.dmg",
    "version": "0.100.29"
  },
  {
    "label": "default",
    "notes": "",
    "os": "windows",
    "url": "https://mobile.mist.com/installers/marvisclient/windows/0.100.26/marvisclient-installer.zip",
    "version": "0.100.26"
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


# List Site Languages

Get List of Languages

```go
ListSiteLanguages(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstLanguage],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstLanguage](../../doc/models/const-language.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListSiteLanguages(ctx)
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
    "display": "English (US)",
    "display_native": "English (US)",
    "key": "en-US"
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


# List States

Get List of ISO States based on country code

```go
ListStates(
    ctx context.Context,
    countryCode string) (
    models.ApiResponse[[]models.ConstState],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `countryCode` | `string` | Query, Required | Country code, in [two-character](../../doc/controllers/constants-definitions.md#list-country-codes) |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstState](../../doc/models/const-state.md).

## Example Usage

```go
ctx := context.Background()

countryCode := "US"

apiResponse, err := constantsDefinitions.ListStates(ctx, countryCode)
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
    "iso_code": "AK",
    "name": "Alaska"
  },
  {
    "iso_code": "AL",
    "name": "Alabama"
  },
  {
    "iso_code": "AS",
    "name": "American Samoa"
  },
  {
    "iso_code": "AZ",
    "name": "Arizona"
  },
  {
    "iso_code": "CA",
    "name": "California"
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

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstTrafficType](../../doc/models/const-traffic-type.md).

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


# List Webhook Topics

Get List of the available Webhook Topics.

```go
ListWebhookTopics(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstWebhookTopic],
    error)
```

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [[]models.ConstWebhookTopic](../../doc/models/const-webhook-topic.md).

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsDefinitions.ListWebhookTopics(ctx)
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
    "for_org": true,
    "has_delivery_results": true,
    "key": "alarms"
  },
  {
    "key": "asset-raw"
  },
  {
    "key": "asset-raw-rssi"
  },
  {
    "for_org": true,
    "has_delivery_results": true,
    "key": "audits"
  },
  {
    "for_org": true,
    "key": "client-info"
  },
  {
    "for_org": true,
    "key": "client-join"
  },
  {
    "key": "client-latency"
  },
  {
    "for_org": true,
    "key": "client-sessions"
  },
  {
    "allows_single_event_per_message": true,
    "for_org": true,
    "key": "device-events"
  },
  {
    "for_org": true,
    "has_delivery_results": true,
    "key": "device-updowns"
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

