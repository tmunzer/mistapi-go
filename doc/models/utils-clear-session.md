
# Utils Clear Session

To use five tuples to lookup the session to be cleared, all must be provided

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `ServiceName` | `*string` | Optional | Service name, only supported in SSR |
| `SessionIds` | `[]uuid.UUID` | Optional | Session identifiers to clear |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    utilsClearSession := models.UtilsClearSession{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        ServiceName:          models.ToPointer("internet-wan_and_lte"),
        SessionIds:           []uuid.UUID{
            uuid.MustParse("88776655-0123-4567-890a-112233445566"),
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

