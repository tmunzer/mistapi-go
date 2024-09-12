# Constants Models

```go
constantsModels := client.ConstantsModels()
```

## Class Name

`ConstantsModels`

## Methods

* [Get Gateway Default Config](../../doc/controllers/constants-models.md#get-gateway-default-config)
* [List Device Models](../../doc/controllers/constants-models.md#list-device-models)
* [List Mx Edge Models](../../doc/controllers/constants-models.md#list-mx-edge-models)
* [List Supported Other Device Models](../../doc/controllers/constants-models.md#list-supported-other-device-models)


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



apiResponse, err := constantsModels.GetGatewayDefaultConfig(ctx, model, nil)
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

## Example Response

```
[
  {
    "ap_type": "jewel",
    "band24": {
      "band5_channels_op": "low",
      "max_clients": 128,
      "max_power": 19,
      "min_power": 8
    },
    "band5": {
      "max_clients": 128,
      "max_power": 17,
      "min_power": 8
    },
    "band6": {
      "max_clients": 128,
      "max_power": 17,
      "min_power": 8
    },
    "band_24_usages": [
      "5"
    ],
    "ce_dfs_ok": true,
    "cisco_pace": true,
    "description": "AP-45",
    "disallowed_channels": {
      "property1": {
        "property1": [
          0
        ],
        "property2": [
          0
        ]
      },
      "property2": {
        "property1": [
          0
        ],
        "property2": [
          0
        ]
      }
    },
    "display": "AP45",
    "fcc_dfs_ok": true,
    "has_11ax": true,
    "has_compass": false,
    "has_ext_ant": true,
    "has_extio": false,
    "has_height": false,
    "has_module_port": true,
    "has_poe_out": true,
    "has_scanning_radio": true,
    "has_selectable_radio": true,
    "has_usb": true,
    "has_vble": true,
    "has_wifi_band24": true,
    "has_wifi_band5": true,
    "has_wifi_band6": true,
    "max_poe_out": 15400,
    "max_wlans": 0,
    "model": "AP45",
    "other_dfs_ok": true,
    "outdoor": false,
    "radios": {
      "r0": "6",
      "r1": "5",
      "r2": "24"
    },
    "shared_scanning_radio": true,
    "type": "ap",
    "unmanaged": true,
    "vble": {
      "beacon_rate": 4,
      "beams": 9,
      "power": 8
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

