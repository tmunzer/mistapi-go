
# Response Discovered Switches

Paginated response for discovered switch search results

## Structure

`ResponseDiscoveredSwitches`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Epoch timestamp, in seconds, for the end of the discovered switch search window |
| `Limit` | `int` | Required | Maximum number of discovered switch records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of discovered switch records |
| `Results` | [`[]models.DiscoveredSwitch`](../../doc/models/discovered-switch.md) | Required | Discovered switch records returned by a search response<br><br>**Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | Epoch timestamp, in seconds, for the start of the discovered switch search window |
| `Total` | `int` | Required | Number of discovered switch records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseDiscoveredSwitches := models.ResponseDiscoveredSwitches{
        End:                  float64(164.74),
        Limit:                80,
        Next:                 models.ToPointer("next4"),
        Results:              []models.DiscoveredSwitch{
            models.DiscoveredSwitch{
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
                    "chassis_id0",
                    "chassis_id1",
                    "chassis_id2",
                },
                ForSite:              models.ToPointer(false),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
            },
        },
        Start:                float64(120.8),
        Total:                82,
    }

}
```

