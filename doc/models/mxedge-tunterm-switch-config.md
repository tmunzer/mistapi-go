
# Mxedge Tunterm Switch Config

Switch VLAN settings for one tunnel termination port

## Structure

`MxedgeTuntermSwitchConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortVlanId` | `*int` | Optional | Untagged VLAN ID for this tunnel termination switch port |
| `VlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | VLAN ID, either numeric or expressed as a template variable string |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermSwitchConfig := models.MxedgeTuntermSwitchConfig{
        PortVlanId:           models.ToPointer(16),
        VlanIds:              []models.VlanIdWithVariable{
            models.VlanIdWithVariableContainer.FromString("String3"),
            models.VlanIdWithVariableContainer.FromString("String4"),
            models.VlanIdWithVariableContainer.FromString("String5"),
        },
    }

}
```

