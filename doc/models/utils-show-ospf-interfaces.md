
# Utils Show Ospf Interfaces

OSPF interfaces command request for SSR and SRX devices

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowOspfInterfaces`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `PortId` | `*string` | Optional | Network interface filter for OSPF interface output |
| `Vrf` | `*string` | Optional | Routing instance or VRF filter for OSPF interface output |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowOspfInterfaces := models.UtilsShowOspfInterfaces{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        PortId:               models.ToPointer("ge-0/0/3"),
        Vrf:                  models.ToPointer("lan"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

