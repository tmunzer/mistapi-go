
# Shell Node

Request body for creating a shell session on a specific HA node

*This model accepts additional fields of type interface{}.*

## Structure

`ShellNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    shellNode := models.ShellNode{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

