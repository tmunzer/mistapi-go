
# Sle Impacted Clients

Paginated list of clients impacted by an SLE metric

## Structure

`SleImpactedClients`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Classifier` | `*string` | Optional | Requested SLE classifier filter applied to the query |
| `Clients` | [`[]models.SleImpactedClientsClient`](../../doc/models/sle-impacted-clients-client.md) | Optional | Impacted client rows returned for an SLE query |
| `End` | `*int` | Optional | Last timestamp in the impacted clients window |
| `Failure` | `*string` | Optional | Requested SLE failure filter applied to the query |
| `Limit` | `*int` | Optional | Maximum number of impacted client rows returned per page |
| `Metric` | `*string` | Optional | SLE metric name used for the impacted clients query |
| `Page` | `*int` | Optional | Current page number for impacted client results |
| `Start` | `*int` | Optional | First timestamp in the impacted clients window |
| `TotalCount` | `*int` | Optional | Number of impacted client rows matching the query |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedClients := models.SleImpactedClients{
        Classifier:           models.ToPointer("classifier0"),
        Clients:              []models.SleImpactedClientsClient{
            models.SleImpactedClientsClient{
                Degraded:             models.ToPointer(18),
                Duration:             models.ToPointer(124),
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
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
            },
            models.SleImpactedClientsClient{
                Degraded:             models.ToPointer(18),
                Duration:             models.ToPointer(124),
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
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
            },
            models.SleImpactedClientsClient{
                Degraded:             models.ToPointer(18),
                Duration:             models.ToPointer(124),
                Mac:                  models.ToPointer("mac2"),
                Name:                 models.ToPointer("name8"),
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
            },
        },
        End:                  models.ToPointer(236),
        Failure:              models.ToPointer("failure8"),
        Limit:                models.ToPointer(190),
    }

}
```

