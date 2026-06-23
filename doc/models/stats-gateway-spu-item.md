
# Stats Gateway Spu Item

Services Processing Unit resource and session counters

## Structure

`StatsGatewaySpuItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SpuCpu` | `*int` | Optional | CPU utilization reported for the Services Processing Unit |
| `SpuCurrentSession` | `*int` | Optional | Current session count handled by the Services Processing Unit |
| `SpuMaxSession` | `*int` | Optional | Maximum sessions supported by the Services Processing Unit |
| `SpuMemory` | `*int` | Optional | Memory utilization reported for the Services Processing Unit |
| `SpuPendingSession` | `*int` | Optional | Pending session count on the Services Processing Unit |
| `SpuUptime` | `*int` | Optional | Elapsed time since the Services Processing Unit started, in seconds |
| `SpuValidSession` | `*int` | Optional | Valid session count currently tracked by the Services Processing Unit |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsGatewaySpuItem := models.StatsGatewaySpuItem{
        SpuCpu:               models.ToPointer(3670632),
        SpuCurrentSession:    models.ToPointer(215),
        SpuMaxSession:        models.ToPointer(131072),
        SpuMemory:            models.ToPointer(46),
        SpuPendingSession:    models.ToPointer(0),
        SpuUptime:            models.ToPointer(0),
    }

}
```

