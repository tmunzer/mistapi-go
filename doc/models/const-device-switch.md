
# Const Device Switch

Switch model capability definition returned by the constants API

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alias` | `*string` | Optional | Alternate model identifier for chassis-based switch definitions |
| `Defaults` | [`*models.ConstDeviceSwitchDefault`](../../doc/models/const-device-switch-default.md) | Optional | Default switch port ranges for a model |
| `Description` | `*string` | Optional | Product description for the switch model |
| `Display` | `*string` | Optional | User-facing model name shown for the switch |
| `EvolvedOs` | `*bool` | Optional | Whether the switch model runs Junos Evolved<br><br>**Default**: `false` |
| `EvpnRiType` | `*string` | Optional | EVPN routing-instance type used by this switch model |
| `Experimental` | `*bool` | Optional | Whether this switch model is marked experimental in the constants catalog<br><br>**Default**: `false` |
| `FansPluggable` | `*bool` | Optional | Whether the switch model has field-replaceable fans<br><br>**Default**: `false` |
| `HasBgp` | `*bool` | Optional | Whether the switch model supports BGP<br><br>**Default**: `false` |
| `HasEts` | `*bool` | Optional | Whether the switch model supports Enhanced Transmission Selection (ETS)<br><br>**Default**: `false` |
| `HasEvpn` | `*bool` | Optional | Whether the switch model supports EVPN<br><br>**Default**: `false` |
| `HasIrb` | `*bool` | Optional | Whether the switch model supports IRB interfaces<br><br>**Default**: `false` |
| `HasPoeOut` | `*bool` | Optional | Whether the switch model supports PoE output<br><br>**Default**: `false` |
| `HasSnapshot` | `*bool` | Optional | Whether the switch model supports configuration snapshots<br><br>**Default**: `true` |
| `HasVc` | `*bool` | Optional | Whether the switch model supports Virtual Chassis<br><br>**Default**: `true` |
| `Model` | `*string` | Optional | Switch model identifier for this capability definition |
| `Modular` | `*bool` | Optional | Whether the switch model has modular hardware<br><br>**Default**: `false` |
| `NoShapingRate` | `*bool` | Optional | Whether the switch model omits shaping-rate support<br><br>**Default**: `false` |
| `NumberFans` | `*int` | Optional | Number of fans in the switch model |
| `OcDevice` | `*bool` | Optional | Whether this switch model is identified as an OpenConfig-managed device<br><br>**Default**: `false` |
| `OobInterface` | `*string` | Optional | Out-of-band management interface names for this switch model |
| `PacketActionDropOnly` | `*bool` | Optional | Whether packet action support is limited to drop-only behavior<br><br>**Default**: `false` |
| `Pic` | `map[string]string` | Optional | Object Key is the PIC number |
| `SubRequired` | `*string` | Optional | Subscription type required for this switch model |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `switch`<br><br>**Value**: `"switch"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceSwitch := models.ConstDeviceSwitch{
        Alias:                models.ToPointer("EX4100-48P-CHAS"),
        Defaults:             models.ToPointer(models.ConstDeviceSwitchDefault{
            Ports:                models.ToPointer("_ports6"),
        }),
        Description:          models.ToPointer("Juniper EX4100 Series"),
        Display:              models.ToPointer("EX4100-48P"),
        EvolvedOs:            models.ToPointer(false),
        EvpnRiType:           models.ToPointer("mac-vrf"),
        Experimental:         models.ToPointer(false),
        FansPluggable:        models.ToPointer(true),
        HasBgp:               models.ToPointer(true),
        HasEts:               models.ToPointer(false),
        HasEvpn:              models.ToPointer(true),
        HasIrb:               models.ToPointer(true),
        HasPoeOut:            models.ToPointer(true),
        HasSnapshot:          models.ToPointer(true),
        HasVc:                models.ToPointer(true),
        Model:                models.ToPointer("EX4100-48P"),
        Modular:              models.ToPointer(false),
        NoShapingRate:        models.ToPointer(false),
        NumberFans:           models.ToPointer(2),
        OcDevice:             models.ToPointer(true),
        OobInterface:         models.ToPointer("re0:mgmt-0, re1:mgmt-0"),
        PacketActionDropOnly: models.ToPointer(false),
        Pic:                  map[string]string{
            "0": "ge*48",
            "1": "qsfp+*4",
            "2": "sfp+*4 (uplink)",
        },
        Type:                 "switch",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

