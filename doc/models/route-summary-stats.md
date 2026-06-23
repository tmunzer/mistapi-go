
# Route Summary Stats

Route table capacity and usage summary

## Structure

`RouteSummaryStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FibRoutes` | `*int` | Optional | Number of routes installed in the forwarding information base |
| `MaxUnicastRoutesSupported` | `*int` | Optional | Supported maximum number of unicast routes |
| `RibRoutes` | `*int` | Optional | Number of routes present in the routing information base |
| `TotalRoutes` | `*int` | Optional | Aggregate number of routes reported by the device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    routeSummaryStats := models.RouteSummaryStats{
        FibRoutes:                 models.ToPointer(224),
        MaxUnicastRoutesSupported: models.ToPointer(218),
        RibRoutes:                 models.ToPointer(86),
        TotalRoutes:               models.ToPointer(46),
    }

}
```

