
# Utils Clear Bgp

Request to clear BGP sessions on a device

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearBgp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Neighbor` | `string` | Required | can be ip, ipv6, all<br><br>**Default**: `"all"` |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `Type` | [`models.UtilsClearBgpTypeEnum`](../../doc/models/utils-clear-bgp-type-enum.md) | Required | enum: `hard`, `in`, `out`, `soft`<br><br>**Default**: `"hard"` |
| `Vrf` | `*string` | Optional | Routing instance or VRF containing the BGP session |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsClearBgp := models.UtilsClearBgp{
        Neighbor:             "all",
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        Type:                 models.UtilsClearBgpTypeEnum_HARD,
        Vrf:                  models.ToPointer("vrf2"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

