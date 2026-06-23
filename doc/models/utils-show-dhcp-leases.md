
# Utils Show Dhcp Leases

DHCP leases command request

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowDhcpLeases`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Network` | `string` | Required | DHCP network for the leases, returns full table if not specified |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowDhcpLeases := models.UtilsShowDhcpLeases{
        Network:              "guest",
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

