
# Mxedge Tunterm Switch Configs

Custom VLAN settings for tunnel termination switch ports, if desired; property key is the port identifier

*This model accepts additional fields of type [models.MxedgeTuntermSwitchConfig](../../doc/models/mxedge-tunterm-switch-config.md).*

## Structure

`MxedgeTuntermSwitchConfigs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether custom tunnel termination switch VLAN settings are enabled |
| `AdditionalProperties` | [`map[string]models.MxedgeTuntermSwitchConfig`](../../doc/models/mxedge-tunterm-switch-config.md) | Optional | Switch VLAN settings for one tunnel termination port |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermSwitchConfigs := models.MxedgeTuntermSwitchConfigs{
        Enabled:              models.ToPointer(false),
        AdditionalProperties: map[string]models.MxedgeTuntermSwitchConfig{
            "exampleAdditionalProperty": models.MxedgeTuntermSwitchConfig{
                PortVlanId:           models.ToPointer(176),
                VlanIds:              []models.VlanIdWithVariable{
                    models.VlanIdWithVariableContainer.FromString("String7"),
                    models.VlanIdWithVariableContainer.FromString("String8"),
                },
            },
        },
    }

}
```

