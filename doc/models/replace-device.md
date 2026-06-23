
# Replace Device

Request payload for replacing an inventory device with a claimed, unassigned device

*This model accepts additional fields of type interface{}.*

## Structure

`ReplaceDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Discard` | `[]string` | Optional | Attributes that should not be copied to the replacement device |
| `InventoryMac` | `*string` | Optional | MAC address of the claimed, unassigned inventory device that will replace the old device |
| `Mac` | `*string` | Optional | Device MAC address being replaced |
| `SiteId` | `*string` | Optional | Site containing the device being replaced |
| `TuntermPortConfig` | [`*models.TuntermPortConfig`](../../doc/models/tunterm-port-config.md) | Optional | Ethernet port configuration for tunnel termination interfaces |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    replaceDevice := models.ReplaceDevice{
        Discard:              []string{
            "discard6",
            "discard7",
            "discard8",
        },
        InventoryMac:         models.ToPointer("5c5b35000301"),
        Mac:                  models.ToPointer("5c5b35000101"),
        SiteId:               models.ToPointer("4ac1dcf4-9d8b-7211-65c4-057819f0862b"),
        TuntermPortConfig:    models.ToPointer(models.TuntermPortConfig{
            DownstreamPorts:            []string{
                "downstream_ports5",
            },
            SeparateUpstreamDownstream: models.ToPointer(false),
            UpstreamPortVlanId:         models.ToPointer(models.TuntermPortConfigUpstreamPortVlanIdContainer.FromString("String1")),
            UpstreamPorts:              []string{
                "upstream_ports0",
            },
        }),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

