
# Utils Show Session

Active session lookup request for device command output

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `ServiceName` | `*string` | Optional | The exact service name for which to display the active sessions |
| `SessionId` | `*string` | Optional | Identifier of the session to show in detail |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowSession := models.UtilsShowSession{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        ServiceName:          models.ToPointer("any"),
        SessionId:            models.ToPointer("session_id0"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

