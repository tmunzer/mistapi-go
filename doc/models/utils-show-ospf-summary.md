
# Utils Show Ospf Summary

OSPF summary command request for SSR and SRX devices

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowOspfSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `Vrf` | `*string` | Optional | Routing instance or VRF filter for OSPF summary output |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowOspfSummary := models.UtilsShowOspfSummary{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        Vrf:                  models.ToPointer("lan"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

