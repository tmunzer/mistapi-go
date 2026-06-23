
# Ap Switch Setting

VLAN settings for one deprecated AP switch-config port

## Structure

`ApSwitchSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EnableVlan` | `*bool` | Optional | Whether VLAN tagging is enabled for this AP switch-config port |
| `PortVlanId` | [`*models.ApSwitchSettingPortVlanId`](../../doc/models/containers/ap-switch-setting-port-vlan-id.md) | Optional | Native VLAN ID, optional |
| `VlanIds` | `[]int` | Optional | List of VLAN IDs<br><br>**Constraints**: `>= 1`, `<= 4094` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apSwitchSetting := models.ApSwitchSetting{
        EnableVlan:           models.ToPointer(false),
        PortVlanId:           models.ToPointer(models.ApSwitchSettingPortVlanIdContainer.FromNumber(226)),
        VlanIds:              []int{
            56,
        },
    }

}
```

