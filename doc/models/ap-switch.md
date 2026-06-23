
# Ap Switch

Deprecated AP switch VLAN control settings for advanced per-port configuration

## Structure

`ApSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether deprecated AP switch VLAN control is enabled<br><br>**Default**: `false` |
| `Eth0` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | VLAN settings for one deprecated AP switch-config port |
| `Eth1` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | VLAN settings for one deprecated AP switch-config port |
| `Eth2` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | VLAN settings for one deprecated AP switch-config port |
| `Eth3` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | VLAN settings for one deprecated AP switch-config port |
| `Module` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | VLAN settings for one deprecated AP switch-config port |
| `Wds` | [`*models.ApSwitchSetting`](../../doc/models/ap-switch-setting.md) | Optional | VLAN settings for one deprecated AP switch-config port |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apSwitch := models.ApSwitch{
        Enabled:              models.ToPointer(false),
        Eth0:                 models.ToPointer(models.ApSwitchSetting{
            EnableVlan:           models.ToPointer(false),
            PortVlanId:           models.ToPointer(models.ApSwitchSettingPortVlanIdContainer.FromNumber(28)),
            VlanIds:              []int{
                114,
                115,
                116,
            },
        }),
        Eth1:                 models.ToPointer(models.ApSwitchSetting{
            EnableVlan:           models.ToPointer(false),
            PortVlanId:           models.ToPointer(models.ApSwitchSettingPortVlanIdContainer.FromNumber(60)),
            VlanIds:              []int{
                146,
                147,
                148,
            },
        }),
        Eth2:                 models.ToPointer(models.ApSwitchSetting{
            EnableVlan:           models.ToPointer(false),
            PortVlanId:           models.ToPointer(models.ApSwitchSettingPortVlanIdContainer.FromNumber(32)),
            VlanIds:              []int{
                118,
                119,
                120,
            },
        }),
        Eth3:                 models.ToPointer(models.ApSwitchSetting{
            EnableVlan:           models.ToPointer(false),
            PortVlanId:           models.ToPointer(models.ApSwitchSettingPortVlanIdContainer.FromNumber(150)),
            VlanIds:              []int{
                236,
                237,
            },
        }),
    }

}
```

