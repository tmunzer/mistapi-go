
# Ap Switch

For people who want to fully control the vlans (advanced)

*This model accepts additional fields of type interface{}.*

## Structure

`ApSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Eth0` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | - |
| `Eth1` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | - |
| `Eth2` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | - |
| `Eth3` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | - |
| `Module` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | - |
| `Wds` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "eth0": {
    "enable_vlan": false,
    "port_vlan_id": 106,
    "vlan_ids": [
      114,
      115,
      116
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "eth1": {
    "enable_vlan": false,
    "port_vlan_id": 138,
    "vlan_ids": [
      146,
      147,
      148
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "eth2": {
    "enable_vlan": false,
    "port_vlan_id": 110,
    "vlan_ids": [
      118,
      119,
      120
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "eth3": {
    "enable_vlan": false,
    "port_vlan_id": 228,
    "vlan_ids": [
      236,
      237
    ],
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

