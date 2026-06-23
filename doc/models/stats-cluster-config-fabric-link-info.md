
# Stats Cluster Config Fabric Link Info

Fabric link status details for a gateway cluster

## Structure

`StatsClusterConfigFabricLinkInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DataPlaneNotifiedStatus` | `*string` | Optional | Fabric link data-plane notification status reported by the gateway |
| `Interface` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `InternalStatus` | `*string` | Optional | Gateway-reported internal status for the fabric link |
| `State` | `*string` | Optional | Administrative state of the fabric link |
| `Status` | `*string` | Optional | Operational status of the fabric link |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsClusterConfigFabricLinkInfo := models.StatsClusterConfigFabricLinkInfo{
        DataPlaneNotifiedStatus: models.ToPointer("DataPlaneNotifiedStatus8"),
        Interface:               []string{
            "Interface6",
            "Interface7",
            "Interface8",
        },
        InternalStatus:          models.ToPointer("InternalStatus8"),
        State:                   models.ToPointer("State0"),
        Status:                  models.ToPointer("Status6"),
    }

}
```

