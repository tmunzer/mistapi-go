
# Stats Cluster Config Control Link Info

Control link status details for a gateway cluster

## Structure

`StatsClusterConfigControlLinkInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Control link interface name reported by the gateway cluster |
| `Status` | `*string` | Optional | Operational state reported for the control link |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsClusterConfigControlLinkInfo := models.StatsClusterConfigControlLinkInfo{
        Name:                 models.ToPointer("name8"),
        Status:               models.ToPointer("status0"),
    }

}
```

