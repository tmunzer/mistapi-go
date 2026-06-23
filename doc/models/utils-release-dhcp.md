
# Utils Release Dhcp

Request to release the DHCP lease on a device interface

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsReleaseDhcp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `PortId` | `string` | Required | Network interface whose current DHCP lease should be released<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsReleaseDhcp := models.UtilsReleaseDhcp{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        PortId:               "ge-0/0/1.10",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

