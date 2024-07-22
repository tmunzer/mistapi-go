
# Ap Switch

for people who want to fully control the vlans (advanced)

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

## Example (as JSON)

```json
{
  "enabled": false,
  "eth0": {
    "additional_vlan_ids": [
      143
    ],
    "enable_vlan": false,
    "port_vlan_id": 106,
    "vlan_ids": [
      114,
      115,
      116
    ]
  },
  "eth1": {
    "additional_vlan_ids": [
      175
    ],
    "enable_vlan": false,
    "port_vlan_id": 138,
    "vlan_ids": [
      146,
      147,
      148
    ]
  },
  "eth2": {
    "additional_vlan_ids": [
      147
    ],
    "enable_vlan": false,
    "port_vlan_id": 110,
    "vlan_ids": [
      118,
      119,
      120
    ]
  },
  "eth3": {
    "additional_vlan_ids": [
      9,
      10,
      11
    ],
    "enable_vlan": false,
    "port_vlan_id": 228,
    "vlan_ids": [
      236,
      237
    ]
  }
}
```

