
# Const Device Gateway

Gateway model capability definition returned by the constants API

## Structure

`ConstDeviceGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Defaults` | `map[string]string` | Optional | Object Key is the interface type name (e.g. "lan_ports", "wan_ports", ...) |
| `Description` | `*string` | Optional | Product description for the gateway model |
| `Experimental` | `*bool` | Optional | Whether this gateway model is marked experimental in the constants catalog<br><br>**Default**: `false` |
| `FansPluggable` | `*bool` | Optional | Whether the gateway model has field-replaceable fans<br><br>**Default**: `true` |
| `HaNode0Fpc` | `*int` | Optional | FPC number used for node0 in HA deployments |
| `HaNode1Fpc` | `*int` | Optional | FPC number used for node1 in HA deployments |
| `HasBgp` | `*bool` | Optional | Whether the gateway model supports BGP<br><br>**Default**: `false` |
| `HasFxp0` | `*bool` | Optional | Whether the gateway model includes an fxp0 management interface<br><br>**Default**: `true` |
| `HasHaControl` | `*bool` | Optional | Whether the gateway model has a dedicated HA control port<br><br>**Default**: `false` |
| `HasHaData` | `*bool` | Optional | Whether the gateway model has dedicated HA data ports<br><br>**Default**: `false` |
| `HasIrb` | `*bool` | Optional | Whether the gateway model supports IRB interfaces<br><br>**Default**: `false` |
| `HasPoeOut` | `*bool` | Optional | Whether the gateway model supports PoE output<br><br>**Default**: `true` |
| `HasSnapshot` | `*bool` | Optional | Whether the gateway model supports configuration snapshots<br><br>**Default**: `true` |
| `IrbDisabledByDefault` | `*bool` | Optional | Whether IRB interfaces are disabled by default on this gateway model<br><br>**Default**: `false` |
| `Model` | `*string` | Optional | Gateway model identifier for this capability definition |
| `NumberFans` | `*int` | Optional | Number of fans in the gateway model |
| `OcDevice` | `*bool` | Optional | Whether this gateway model is identified as an OpenConfig-managed device<br><br>**Default**: `false` |
| `Pic` | `map[string]string` | Optional | Object Key is the PIC number |
| `Ports` | [`*models.ConstDeviceGatewayPorts`](../../doc/models/const-device-gateway-ports.md) | Optional | Object Key is the interface name (e.g. "ge-0/0/1", ...) |
| `SubRequired` | `*string` | Optional | Subscription type required for this gateway model |
| `T128Device` | `*bool` | Optional | Whether this gateway model is a 128 Technology or SSR device<br><br>**Default**: `false` |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `gateway`<br><br>**Value**: `"gateway"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceGateway := models.ConstDeviceGateway{
        Defaults:             map[string]string{
            "key0": "defaults0",
            "key1": "defaults1",
            "key2": "defaults2",
        },
        Description:          models.ToPointer("description2"),
        Experimental:         models.ToPointer(false),
        FansPluggable:        models.ToPointer(true),
        HaNode0Fpc:           models.ToPointer(76),
        HasBgp:               models.ToPointer(false),
        HasFxp0:              models.ToPointer(true),
        HasHaControl:         models.ToPointer(false),
        HasHaData:            models.ToPointer(false),
        HasIrb:               models.ToPointer(false),
        HasPoeOut:            models.ToPointer(true),
        HasSnapshot:          models.ToPointer(true),
        IrbDisabledByDefault: models.ToPointer(false),
        OcDevice:             models.ToPointer(false),
        T128Device:           models.ToPointer(false),
        Type:                 "gateway",
    }

}
```

