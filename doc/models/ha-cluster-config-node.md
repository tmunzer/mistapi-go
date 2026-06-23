
# Ha Cluster Config Node

HA cluster inventory node identified by MAC address

## Structure

`HaClusterConfigNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Gateway node MAC address for an inventory node that is currently unassigned |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    haClusterConfigNode := models.HaClusterConfigNode{
        Mac:                  models.ToPointer("aff827549235"),
    }

}
```

