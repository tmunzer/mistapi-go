
# Response Discovered Switch Metrics

Paginated response for discovered switch metric search results

## Structure

`ResponseDiscoveredSwitchMetrics`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | Epoch timestamp, in seconds, for the end of the metric search window |
| `Limit` | `int` | Required | Maximum number of discovered switch metric records returned in this page |
| `Next` | `*string` | Optional | Pagination cursor or URL for retrieving the next page of discovered switch metric records |
| `Results` | [`[]models.DiscoveredSwitchMetric`](../../doc/models/discovered-switch-metric.md) | Required | Discovered switch metric records returned by a search response |
| `Start` | `float64` | Required | Epoch timestamp, in seconds, for the start of the metric search window |
| `Total` | `int` | Required | Number of discovered switch metric records matching the search filters across all pages |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseDiscoveredSwitchMetrics := models.ResponseDiscoveredSwitchMetrics{
        End:                  float64(74.88),
        Limit:                150,
        Next:                 models.ToPointer("next8"),
        Results:              []models.DiscoveredSwitchMetric{
            models.DiscoveredSwitchMetric{
                Adopted:              models.ToPointer(false),
                Aps:                  []models.DiscoveredSwitchMetricAp{
                    models.DiscoveredSwitchMetricAp{
                        Hostname:             models.ToPointer("hostname0"),
                        Mac:                  models.ToPointer("mac8"),
                        PoeStatus:            models.ToPointer(false),
                        Port:                 models.ToPointer("port4"),
                        PortId:               models.ToPointer("port_id4"),
                    },
                    models.DiscoveredSwitchMetricAp{
                        Hostname:             models.ToPointer("hostname0"),
                        Mac:                  models.ToPointer("mac8"),
                        PoeStatus:            models.ToPointer(false),
                        Port:                 models.ToPointer("port4"),
                        PortId:               models.ToPointer("port_id4"),
                    },
                },
                ChassisId:            []string{
                    "chassis_id0",
                    "chassis_id1",
                    "chassis_id2",
                },
                Hostname:             models.ToPointer("SW-HLAB-ea2e00"),
                MgmtAddr:             models.ToPointer("10.10.20.42"),
                OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
                Scope:                models.ToPointer("site"),
                Score:                models.ToPointer(100),
                SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
                SystemDesc:           models.ToPointer("Juniper Networks, Inc. ex4100-f-12p Ethernet Switch, kernel JUNOS 22.4R3.25, Build date: 2024-02-10 00:49:09 UTC Copyright (c) 1996-2024 Juniper Networks, Inc."),
                SystemName:           models.ToPointer("SW-HLAB-ea2e00"),
                Type:                 models.ToPointer("inactive_wired_vlans"),
            },
        },
        Start:                float64(30.94),
        Total:                56,
    }

}
```

