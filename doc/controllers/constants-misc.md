# Constants Misc

```go
constantsMisc := client.ConstantsMisc()
```

## Class Name

`ConstantsMisc`

## Methods

* [Get Gateway Default Config](../../doc/controllers/constants-misc.md#get-gateway-default-config)
* [Get License Types](../../doc/controllers/constants-misc.md#get-license-types)
* [List Ap Channels](../../doc/controllers/constants-misc.md#list-ap-channels)
* [List Insight Metrics](../../doc/controllers/constants-misc.md#list-insight-metrics)
* [List Site Languages](../../doc/controllers/constants-misc.md#list-site-languages)


# Get Gateway Default Config

Generate Default Gateway Config

```go
GetGatewayDefaultConfig(
    ctx context.Context,
    model string,
    ha *string) (
    models.ApiResponse[interface{}],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `model` | `string` | Query, Required | model the default gateway config is intended (as the default LAN/WAN port can differ) |
| `ha` | `*string` | Query, Optional | whether the config is intended for HA |

## Response Type

`interface{}`

## Example Usage

```go
ctx := context.Background()

model := "model2"



apiResponse, err := constantsMisc.GetGatewayDefaultConfig(ctx, model, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response

```
{
  "dhcpd_config": {
    "lan": {
      "ip_end": "192.168.1.254",
      "ip_start": "192.168.1.2"
    }
  },
  "ip_configs": {
    "lan": {
      "ip": "192.168.1.1",
      "type": "static"
    }
  },
  "networks": {
    "lan": {
      "name": "lan",
      "subnet": "192.168.1.0/24",
      "vlan_id": 1
    }
  },
  "path_preferences": {
    "wan": {
      "paths": [
        {
          "name": "wan",
          "type": "wan"
        }
      ]
    }
  },
  "port_config": {
    "cl-1/0/0": {
      "ip_config": {
        "type": "dhcp"
      },
      "name": "lte",
      "usage": "wan",
      "wan_type": "lte"
    },
    "ge-0/0/0,ge-0/0/7": {
      "ip_config": {
        "type": "dhcp"
      },
      "name": "wan",
      "usage": "wan"
    },
    "ge-0/0/1-6": {
      "port_network": "lan",
      "usage": "lan"
    }
  },
  "service_policies": [
    {
      "action": "allow",
      "name": "Internet",
      "path_preference": "wan",
      "services": [
        "any"
      ],
      "tenants": [
        "lan"
      ]
    }
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


# Get License Types

Get License Types

```go
GetLicenseTypes(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstLicenseType],
    error)
```

## Response Type

[`[]models.ConstLicenseType`](../../doc/models/const-license-type.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsMisc.GetLicenseTypes(ctx)
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
| `countryCode` | `*string` | Query, Optional | country code, in two-character |

## Response Type

[`models.ConstApChannel`](../../doc/models/const-ap-channel.md)

## Example Usage

```go
ctx := context.Background()

countryCode := "US"

apiResponse, err := constantsMisc.ListApChannels(ctx, &countryCode)
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


# List Insight Metrics

List Insight Metrics

```go
ListInsightMetrics(
    ctx context.Context) (
    models.ApiResponse[map[string]models.ConstInsightMetricsProperty],
    error)
```

## Response Type

[`map[string]models.ConstInsightMetricsProperty`](../../doc/models/const-insight-metrics-property.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsMisc.ListInsightMetrics(ctx)
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
    "description": "aggregated bytes over time",
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
    "description": "number of client over time",
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


# List Site Languages

Get List of Languages

```go
ListSiteLanguages(
    ctx context.Context) (
    models.ApiResponse[[]models.ConstLanguage],
    error)
```

## Response Type

[`[]models.ConstLanguage`](../../doc/models/const-language.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := constantsMisc.ListSiteLanguages(ctx)
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

