
# Stats Cluster Config Ethernet Connection Item

Ethernet interface connection state for a gateway cluster member

## Structure

`StatsClusterConfigEthernetConnectionItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Ethernet interface name reported for cluster connectivity |
| `Status` | `*string` | Optional | Operational state reported for the Ethernet interface |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsClusterConfigEthernetConnectionItem := models.StatsClusterConfigEthernetConnectionItem{
        Name:                 models.ToPointer("name6"),
        Status:               models.ToPointer("status8"),
    }

}
```

