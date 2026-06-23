
# Ha Cluster Config

Request to create an HA cluster from unassigned gateway inventory nodes

*This model accepts additional fields of type interface{}.*

## Structure

`HaClusterConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableAutoConfig` | `*bool` | Optional | This disables the default behavior of a cloud-ready switch/gateway being managed/configured by Mist. Setting this to `true` means you want to disable the default behavior and do not want the device to be Mist-managed. |
| `Managed` | `*bool` | Optional | An adopted switch/gateway will not be managed/configured by Mist by default. Setting this parameter to `true` enables the adopted switch/gateway to be managed/configured by Mist. |
| `MistConfigured` | `*bool` | Optional | whether the device can be configured by Mist or not. This deprecates `managed` (for adopted device) and `disable_auto_config` for claimed device) |
| `Nodes` | [`[]models.HaClusterConfigNode`](../../doc/models/ha-cluster-config-node.md) | Optional | Gateway inventory nodes used to create an HA cluster |
| `SiteId` | `*uuid.UUID` | Optional | Site where the HA cluster should be created |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    haClusterConfig := models.HaClusterConfig{
        DisableAutoConfig:    models.ToPointer(false),
        Managed:              models.ToPointer(false),
        MistConfigured:       models.ToPointer(false),
        Nodes:                []models.HaClusterConfigNode{
            models.HaClusterConfigNode{
                Mac:                  models.ToPointer("mac0"),
            },
        },
        SiteId:               models.ToPointer(uuid.MustParse("43e9c864-a7e4-4310-8031-d9817d2c5a43")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

