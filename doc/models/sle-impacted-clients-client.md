
# Sle Impacted Clients Client

SLE impact row for a client

## Structure

`SleImpactedClientsClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Degraded` | `*int` | Optional | Portion of the SLE total that was degraded for this client |
| `Duration` | `*int` | Optional | Observation time represented by this client impact row |
| `Mac` | `*string` | Optional | Client MAC address for the impacted client |
| `Name` | `*string` | Optional | Display name for the client impact row |
| `Switches` | [`[]models.SleImpactedClientsClientSwitch`](../../doc/models/sle-impacted-clients-client-switch.md) | Optional | Switch associations for impacted clients |
| `Total` | `*int` | Optional | Overall SLE total measured for this client impact row |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedClientsClient := models.SleImpactedClientsClient{
        Degraded:             models.ToPointer(230),
        Duration:             models.ToPointer(80),
        Mac:                  models.ToPointer("mac8"),
        Name:                 models.ToPointer("name4"),
        Switches:             []models.SleImpactedClientsClientSwitch{
            models.SleImpactedClientsClientSwitch{
                Interfaces:           []string{
                    "interfaces9",
                },
                SwitchMac:            models.ToPointer("switch_mac6"),
                SwitchName:           models.ToPointer("switch_name0"),
            },
            models.SleImpactedClientsClientSwitch{
                Interfaces:           []string{
                    "interfaces9",
                },
                SwitchMac:            models.ToPointer("switch_mac6"),
                SwitchName:           models.ToPointer("switch_name0"),
            },
            models.SleImpactedClientsClientSwitch{
                Interfaces:           []string{
                    "interfaces9",
                },
                SwitchMac:            models.ToPointer("switch_mac6"),
                SwitchName:           models.ToPointer("switch_name0"),
            },
        },
    }

}
```

