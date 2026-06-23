
# Discovered Switch

Switch discovered from AP uplink LLDP data and site switch discovery

## Structure

`DiscoveredSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Adopted` | `*bool` | Optional | Whether the discovered switch has been adopted into Mist management |
| `ApRedundancy` | [`*models.ApRedundancy`](../../doc/models/ap-redundancy.md) | Optional | AP switch redundancy coverage summary |
| `Aps` | [`[]models.DiscoveredSwitchAp`](../../doc/models/discovered-switch-ap.md) | Optional | AP attachments observed on a discovered switch<br><br>**Constraints**: *Unique Items Required* |
| `ChassisId` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `ForSite` | `*bool` | Optional, Read-only | Whether the discovered switch is associated with the requested site |
| `MgmtAddr` | `*string` | Optional | Management IP address advertised by the discovered switch |
| `Model` | `*string` | Optional | Switch model reported for the discovered switch |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SystemDesc` | `*string` | Optional | LLDP system description advertised by the discovered switch |
| `SystemName` | `*string` | Optional | LLDP system name advertised by the discovered switch |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Vendor` | `*string` | Optional | Switch vendor reported for the discovered switch |
| `Version` | `*string` | Optional | Software version reported for the discovered switch |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    discoveredSwitch := models.DiscoveredSwitch{
        Adopted:              models.ToPointer(false),
        ApRedundancy:         models.ToPointer(models.ApRedundancy{
            Modules:                    map[string]models.ApRedundancyModule{
                "key0": models.ApRedundancyModule{
                    NumAps:                     models.ToPointer(2),
                    NumApsWithSwitchRedundancy: models.ToPointer(254),
                },
            },
            NumAps:                     models.ToPointer(246),
            NumApsWithSwitchRedundancy: models.ToPointer(10),
        }),
        Aps:                  []models.DiscoveredSwitchAp{
            models.DiscoveredSwitchAp{
                Hostname:             models.ToPointer("hostname0"),
                InactiveWiredVlans:   []int{
                    168,
                    169,
                    170,
                },
                Mac:                  models.ToPointer("mac8"),
                PoeStatus:            models.ToPointer(false),
                Port:                 models.ToPointer("port4"),
            },
        },
        ChassisId:            []string{
            "chassis_id2",
            "chassis_id3",
        },
        ForSite:              models.ToPointer(false),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

