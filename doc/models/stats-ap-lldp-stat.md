
# Stats Ap Lldp Stat

LLDP neighbor information and power negotiations. For backward compatibility, when multiple neighbors exist, only information from the first neighbor is displayed.

## Structure

`StatsApLldpStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChassisId` | `models.Optional[string]` | Optional, Read-only | LLDP neighbor chassis identifier |
| `LldpMedSupported` | `models.Optional[bool]` | Optional, Read-only | Whether it support LLDP-MED |
| `MgmtAddr` | `models.Optional[string]` | Optional, Read-only | Management IP address of the switch |
| `MgmtAddrs` | `[]string` | Optional | Management IP addresses advertised by an LLDP neighbor |
| `PortDesc` | `models.Optional[string]` | Optional, Read-only | Port description, e.g. “2/20”, “Port 2 on Switch0” |
| `PortId` | `models.Optional[string]` | Optional, Read-only | LLDP neighbor port identifier |
| `PowerAllocated` | `models.Optional[float64]` | Optional, Read-only | In mW, power allocated by PSE |
| `PowerAvail` | `*int` | Optional | In mW, total Power Avail at AP from pwr source |
| `PowerBudget` | `*int` | Optional | In mW, surplus if positive or deficit if negative |
| `PowerConstrained` | `*bool` | Optional | Whether power is insufficient |
| `PowerDraw` | `models.Optional[float64]` | Optional, Read-only | In mW, total power needed by PD |
| `PowerNeeded` | `*int` | Optional | In mW, total Power needed incl Peripherals |
| `PowerOpmode` | `*string` | Optional | Power operating mode negotiated through LLDP |
| `PowerRequestCount` | `models.Optional[int]` | Optional, Read-only | Number of negotiations, if it keeps increasing, we don’ t have a stable power |
| `PowerRequested` | `models.Optional[float64]` | Optional, Read-only | In mW, power requested by PD |
| `PowerSrc` | `*string` | Optional | Single power source (DC Input / PoE 802.3at / PoE 802.3af / PoE 802.3bt / MULTI-PD / LLDP / ? (unknown)). |
| `PowerSrcs` | `[]string` | Optional | Power sources reported through LLDP for an AP |
| `SystemDesc` | `models.Optional[string]` | Optional, Read-only | Description provided by switch |
| `SystemName` | `models.Optional[string]` | Optional, Read-only | Name of the switch |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsApLldpStat := models.StatsApLldpStat{
        ChassisId:            models.NewOptional(models.ToPointer("chassis_id8")),
        LldpMedSupported:     models.NewOptional(models.ToPointer(false)),
        MgmtAddr:             models.NewOptional(models.ToPointer("mgmt_addr6")),
        MgmtAddrs:            []string{
            "mgmt_addrs1",
        },
        PortDesc:             models.NewOptional(models.ToPointer("2/20")),
        PortId:               models.NewOptional(models.ToPointer("ge-0/0/4")),
        SystemDesc:           models.NewOptional(models.ToPointer("uniper Networks, Inc. ex4300-48t internet router, kernel JUNOS 20.4R3-S7.2, Build date: 2023-04-21 19:47:18 UTC Copyright (c) 1996-2023 Juniper Networks, Inc.")),
        SystemName:           models.NewOptional(models.ToPointer("Core-AE23")),
    }

}
```

