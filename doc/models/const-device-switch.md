
# Const Device Switch

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Alias` | `*string` | Optional | - |
| `Defaults` | [`*models.ConstDeviceSwitchDefault`](../../doc/models/const-device-switch-default.md) | Optional | - |
| `Description` | `*string` | Optional | - |
| `Display` | `*string` | Optional | - |
| `EvolvedOs` | `*bool` | Optional | **Default**: `false` |
| `EvpnRiType` | `*string` | Optional | - |
| `Experimental` | `*bool` | Optional | **Default**: `false` |
| `FansPluggable` | `*bool` | Optional | **Default**: `false` |
| `HasBgp` | `*bool` | Optional | **Default**: `false` |
| `HasEts` | `*bool` | Optional | **Default**: `false` |
| `HasEvpn` | `*bool` | Optional | **Default**: `false` |
| `HasIrb` | `*bool` | Optional | **Default**: `false` |
| `HasPoeOut` | `*bool` | Optional | **Default**: `false` |
| `HasSnapshot` | `*bool` | Optional | **Default**: `true` |
| `HasVc` | `*bool` | Optional | **Default**: `true` |
| `Model` | `*string` | Optional | - |
| `Modular` | `*bool` | Optional | **Default**: `false` |
| `NoShapingRate` | `*bool` | Optional | **Default**: `false` |
| `NumberFans` | `*int` | Optional | - |
| `OcDevice` | `*bool` | Optional | **Default**: `false` |
| `OobInterface` | `*string` | Optional | - |
| `PacketActionDropOnly` | `*bool` | Optional | **Default**: `false` |
| `Pic` | `map[string]string` | Optional | Object Key is the PIC number |
| `SubRequired` | `*string` | Optional | - |
| `Type` | `string` | Required, Constant | Device Type. enum: `switch`<br>**Default**: `"switch"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "alias": "EX4100-48P-CHAS",
  "description": "Juniper EX4100 Series",
  "display": "EX4100-48P",
  "evolved_os": false,
  "evpn_ri_type": "mac-vrf",
  "experimental": false,
  "fans_pluggable": true,
  "has_bgp": true,
  "has_ets": false,
  "has_evpn": true,
  "has_irb": true,
  "has_poe_out": true,
  "has_snapshot": true,
  "has_vc": true,
  "model": "EX4100-48P",
  "modular": false,
  "no_shaping_rate": false,
  "number_fans": 2,
  "oc_device": true,
  "oob_interface": "re0:mgmt-0, re1:mgmt-0",
  "packet_action_drop_only": false,
  "pic": {
    "0": "ge*48",
    "1": "qsfp+*4",
    "2": "sfp+*4 (uplink)"
  },
  "type": "switch",
  "defaults": {
    "_ports": "_ports6",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

