
# Rrm Neighbors Neighbor

Neighbor AP observed by RRM

## Structure

`RrmNeighborsNeighbor`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | Neighbor AP MAC address observed by RRM |
| `Rssi` | `*int` | Optional | Observed RSSI for the neighbor AP, in dBm |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmNeighborsNeighbor := models.RrmNeighborsNeighbor{
        Mac:                  models.ToPointer("5c5b35000311"),
        Rssi:                 models.ToPointer(-66),
    }

}
```

