
# Ha Cluster Delete

Request body identifying the HA cluster node to remove

*This model accepts additional fields of type interface{}.*

## Structure

`HaClusterDelete`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Node0 MAC address identifying the HA cluster to delete |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    haClusterDelete := models.HaClusterDelete{
        Mac:                  models.ToPointer("aff827549235"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

