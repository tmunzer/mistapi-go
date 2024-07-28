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
    models.ApiResponse[[]models.ListDeviceModelsResponse],
    error)
```

## Response Type

[`[]models.ListDeviceModelsResponse`](../../doc/models/containers/list-device-models-response.md)

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

## Example Response *(as JSON)*

```json
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
    "extio": {
      "property1": {
        "default_dir": "IN",
        "input": true,
        "output": true
      },
      "property2": {
        "default_dir": "IN",
        "input": true,
        "output": true
      }
    },
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
    "outdoor": true,
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
  },
  {
    "alias": "EX4100-48P-CHAS",
    "defaults": {
      "_ports": "ge-0/0/0-47, et-0/1/0-3, xe-0/2/0-3, ge-0/2/0-3"
    },
    "description": "Juniper EX4100 Series",
    "display": "EX4100-48P",
    "evpn_ri_type": "mac-vrf",
    "fans_pluggable": true,
    "has_bgp": true,
    "has_evpn": true,
    "has_irb": true,
    "has_poe_out": true,
    "model": "EX4100-48P",
    "number_fans": 2,
    "oc_device": true,
    "oob_interface": "re0:mgmt-0, re1:mgmt-0",
    "pic": {
      "0": "ge*48",
      "1": "qsfp+*4",
      "2": "sfp+*4 (uplink)"
    },
    "sub_required": "string",
    "type": "switch"
  },
  {
    "defaults": {
      "ha_control_port": "ge-0/0/1",
      "ha_data_ports": "ge-0/0/2,ge-3/0/2",
      "ha_fxp0_port": "ge-0/0/0",
      "ha_lan_ports": "ge-0/0/4,ge-3/0/4",
      "ha_wan_ports": "ge-0/0/3,ge-3/0/3",
      "lan_ports": "ge-0/0/1-6",
      "lte_wan_ports": "cl-1/0/0",
      "wan_ports": "ge-0/0/0,ge-0/0/7"
    },
    "description": "Juniper SRX320 Series",
    "fans_pluggable": false,
    "ha_node1_fpc": 3,
    "has_bgp": true,
    "has_fxp0": false,
    "has_irb": true,
    "model": "SRX320",
    "number_fans": 1,
    "oc_device": true,
    "pic": {
      "0": "ge*6, sfp*2"
    },
    "sub_required": "wan1",
    "type": "gateway"
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

