
# Discovered Switch Metric

Time-series metric result for discovered switch health or compliance

## Structure

`DiscoveredSwitchMetric`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Adopted` | `*bool` | Optional | Whether the discovered switch in this metric has been adopted into Mist management |
| `Aps` | [`[]models.DiscoveredSwitchMetricAp`](../../doc/models/discovered-switch-metric-ap.md) | Optional | AP attachment details included in discovered switch metrics |
| `ChassisId` | `[]string` | Optional | LLDP chassis identifiers associated with a discovered switch metric |
| `Hostname` | `*string` | Optional | Switch hostname associated with this metric result |
| `MgmtAddr` | `*string` | Optional | Management IP address associated with this metric result |
| `Model` | `*string` | Optional | Switch model associated with this metric result |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Scope` | `*string` | Optional | Aggregation scope for the discovered switch metric result |
| `Score` | `*int` | Optional | Compliance or health score for the discovered switch metric |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `SystemDesc` | `*string` | Optional | LLDP system description associated with this metric result |
| `SystemName` | `*string` | Optional | LLDP system name associated with this metric result |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | Metric category represented by this discovered switch metric |
| `Vendor` | `*string` | Optional | Switch vendor associated with this metric result |
| `Version` | `*string` | Optional | Software version associated with this metric result |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    discoveredSwitchMetric := models.DiscoveredSwitchMetric{
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
            models.DiscoveredSwitchMetricAp{
                Hostname:             models.ToPointer("hostname0"),
                Mac:                  models.ToPointer("mac8"),
                PoeStatus:            models.ToPointer(false),
                Port:                 models.ToPointer("port4"),
                PortId:               models.ToPointer("port_id4"),
            },
        },
        ChassisId:            []string{
            "chassis_id4",
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
    }

}
```

