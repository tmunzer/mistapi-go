
# Const Device Gateway

## Structure

`ConstDeviceGateway`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Defaults` | `map[string]string` | Optional | Object Key is the interface type name (e.g. "lan_ports", "wan_ports", ...) |
| `Description` | `*string` | Optional | - |
| `Experimental` | `*bool` | Optional | **Default**: `false` |
| `FansPluggable` | `*bool` | Optional | **Default**: `true` |
| `HaNode0Fpc` | `*int` | Optional | - |
| `HaNode1Fpc` | `*int` | Optional | - |
| `HasBgp` | `*bool` | Optional | **Default**: `false` |
| `HasFxp0` | `*bool` | Optional | **Default**: `true` |
| `HasHaControl` | `*bool` | Optional | **Default**: `false` |
| `HasHaData` | `*bool` | Optional | **Default**: `false` |
| `HasIrb` | `*bool` | Optional | **Default**: `false` |
| `HasPoeOut` | `*bool` | Optional | **Default**: `true` |
| `HasSnapshot` | `*bool` | Optional | **Default**: `true` |
| `IrbDisabledByDefault` | `*bool` | Optional | **Default**: `false` |
| `Model` | `*string` | Optional | - |
| `NumberFans` | `*int` | Optional | - |
| `OcDevice` | `*bool` | Optional | **Default**: `false` |
| `Pic` | `map[string]string` | Optional | Object Key is the PIC number |
| `Ports` | [`*models.ConstDeviceGatewayPorts`](../../doc/models/const-device-gateway-ports.md) | Optional | Object Key is the interface name (e.g. "ge-0/0/1", ...) |
| `SubRequired` | `*string` | Optional | - |
| `T128Device` | `*bool` | Optional | **Default**: `false` |
| `Type` | `string` | Required, Constant | Device Type. enum: `gateway`<br>**Default**: `"gateway"` |

## Example (as JSON)

```json
{
  "experimental": false,
  "fans_pluggable": true,
  "has_bgp": false,
  "has_fxp0": true,
  "has_ha_control": false,
  "has_ha_data": false,
  "has_irb": false,
  "has_poe_out": true,
  "has_snapshot": true,
  "irb_disabled_by_default": false,
  "oc_device": false,
  "t128_device": false,
  "type": "gateway",
  "defaults": {
    "key0": "defaults6",
    "key1": "defaults7",
    "key2": "defaults8"
  },
  "description": "description8",
  "ha_node0_fpc": 116
}
```

