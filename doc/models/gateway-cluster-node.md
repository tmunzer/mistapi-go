
# Gateway Cluster Node

Gateway cluster node identified by MAC address

## Structure

`GatewayClusterNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Gateway device MAC address. Format is `[0-9a-f]{12}` (e.g. "5684dae9ac8b") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayClusterNode := models.GatewayClusterNode{
        Mac:                  "mac6",
    }

}
```

