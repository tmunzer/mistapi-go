
# Stats Gateway Cluster

High-availability cluster state reported by a gateway

## Structure

`StatsGatewayCluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `State` | `models.Optional[string]` | Optional, Read-only | Current HA cluster state for the gateway |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsGatewayCluster := models.StatsGatewayCluster{
        State:                models.NewOptional(models.ToPointer("state0")),
    }

}
```

