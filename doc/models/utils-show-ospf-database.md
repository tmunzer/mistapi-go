
# Utils Show Ospf Database

OSPF database command request for SSR and SRX devices

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowOspfDatabase`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `SelfOriginate` | `*bool` | Optional | Whether to show only self-originated OSPF database entries<br><br>**Default**: `false` |
| `Vrf` | `*string` | Optional | Routing instance or VRF filter for OSPF database output |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowOspfDatabase := models.UtilsShowOspfDatabase{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        SelfOriginate:        models.ToPointer(false),
        Vrf:                  models.ToPointer("lan"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

