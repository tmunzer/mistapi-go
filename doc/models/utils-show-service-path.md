
# Utils Show Service Path

Service path lookup request for SSR devices

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowServicePath`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `ServiceName` | `*string` | Optional | Exact service name for which to display the service path |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowServicePath := models.UtilsShowServicePath{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        ServiceName:          models.ToPointer("any"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

