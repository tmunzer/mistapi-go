
# Stats Mxedge

Statistics for a Mist Edge appliance returned by stats endpoints

## Structure

`StatsMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CpuStat` | [`*models.StatsMxedgeCpuStat`](../../doc/models/stats-mxedge-cpu-stat.md) | Optional | Aggregate and per-core CPU utilization statistics for a Mist Edge |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `FipsEnabled` | `*bool` | Optional | Whether FIPS mode is enabled on the Mist Edge |
| `ForSite` | `*bool` | Optional | Whether the Mist Edge is assigned directly to a site |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | Firmware update status for a device |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `IdracVersion` | `*string` | Optional | IDRAC version of the mist edge device |
| `InactiveVlanStrs` | [`*models.StatsMxedgeInactiveVlanStrs`](../../doc/models/stats-mxedge-inactive-vlan-strs.md) | Optional | Inactive wired/L2TP VLANs. Entries can be individual VLANs or ranges. |
| `IpStat` | [`*models.StatsMxedgeIpStat`](../../doc/models/stats-mxedge-ip-stat.md) | Optional | IP address statistics reported by a Mist Edge |
| `LagStat` | [`map[string]models.StatsMxedgeLagStat`](../../doc/models/stats-mxedge-lag-stat.md) | Optional | Link aggregation group statistics keyed by LAG name |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `Mac` | `*string` | Optional | Mist Edge MAC address reported by Mist |
| `Magic` | `*string` | Optional | Claim magic token associated with the Mist Edge |
| `MemoryStat` | [`*models.StatsMxedgeMemoryStat`](../../doc/models/stats-mxedge-memory-stat.md) | Optional | Memory usage counters reported by a Mist Edge |
| `Model` | `*string` | Optional | Mist Edge hardware or VM model |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MxagentRegistered` | `*bool` | Optional | Whether the Mist Edge management agent is registered with Mist |
| `MxclusterId` | `*uuid.UUID` | Optional | Mist Edge cluster identifier associated with this Mist Edge |
| `Name` | `*string` | Optional | Display name of the Mist Edge |
| `NumTunnels` | `*int` | Optional | Number of tunnels currently terminated by the Mist Edge |
| `OobIpConfig` | [`*models.MxedgeOobIpConfig`](../../doc/models/mxedge-oob-ip-config.md) | Optional | IP configuration for the Mist Edge out-of-band management interface |
| `OobIpStat` | [`*models.StatsMxedgeOobIpStat`](../../doc/models/stats-mxedge-oob-ip-stat.md) | Optional | Observed out-of-band management IP settings for a Mist Edge |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PortStat` | [`map[string]models.StatsMxedgePortStat`](../../doc/models/stats-mxedge-port-stat.md) | Optional | Mist Edge port statistics keyed by port name |
| `Serial` | `models.Optional[string]` | Optional | Device serial number reported for the Mist Edge, when available |
| `ServiceStat` | [`map[string]models.StatsMxedgeServiceStat`](../../doc/models/stats-mxedge-service-stat.md) | Optional | Stat for each services |
| `Services` | `[]string` | Optional | Mist Edge service names enabled or running on the appliance |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Status` | `*string` | Optional | Connection status reported for the Mist Edge, such as connected |
| `TuntermIpConfig` | [`*models.StatsMxedgeTuntermIpConfig`](../../doc/models/stats-mxedge-tunterm-ip-config.md) | Optional | Tunnel termination IP configuration reported by a Mist Edge |
| `TuntermPortConfig` | [`*models.StatsMxedgeTuntermPortConfig`](../../doc/models/stats-mxedge-tunterm-port-config.md) | Optional | Tunnel termination port role configuration reported by a Mist Edge |
| `TuntermRegistered` | `*bool` | Optional | Whether the tunnel termination service is registered with Mist |
| `TuntermStat` | [`*models.StatsMxedgeTuntermStat`](../../doc/models/stats-mxedge-tunterm-stat.md) | Optional | Tunnel termination monitoring status reported by a Mist Edge |
| `Uptime` | `*int` | Optional | Number of seconds the Mist Edge has been running since last boot |
| `VirtualizationType` | `*string` | Optional | Virtualization platform or environment running the Mist Edge VM |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsMxedge := models.StatsMxedge{
        CpuStat:              models.ToPointer(models.StatsMxedgeCpuStat{
            Cpus:                 map[string]models.CpuStat{
                "key0": nil,
            },
            Idle:                 models.ToPointer(224),
            Interrupt:            models.ToPointer(80),
            System:               models.ToPointer(80),
            Usage:                models.ToPointer(46),
        }),
        CreatedTime:          models.ToPointer(float64(79.58)),
        FipsEnabled:          models.ToPointer(false),
        ForSite:              models.ToPointer(false),
        Fwupdate:             models.ToPointer(models.FwupdateStat{
            Progress:             models.NewOptional(models.ToPointer(100)),
            Status:               models.NewOptional(models.ToPointer(models.FwupdateStatStatusEnum_INPROGRESS)),
            StatusId:             models.NewOptional(models.ToPointer(70)),
            Timestamp:            models.ToPointer(float64(147.68)),
            WillRetry:            models.NewOptional(models.ToPointer(false)),
        }),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        IdracVersion:         models.ToPointer("7.00.00.00"),
        LagStat:              map[string]models.StatsMxedgeLagStat{
            "lacp0": models.StatsMxedgeLagStat{
                ActivePorts:          []string{
                    "port0",
                    "port1",
                },
            },
        },
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Mac:                  models.ToPointer("020000a80cb4"),
        Model:                models.ToPointer("ME-VM"),
        MxagentRegistered:    models.ToPointer(true),
        MxclusterId:          models.ToPointer(uuid.MustParse("678bc339-7635-4556-bbc0-e77ad493ef8b")),
        Name:                 models.ToPointer("me-vm-1"),
        NumTunnels:           models.ToPointer(0),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PortStat:             map[string]models.StatsMxedgePortStat{
            "port0": models.StatsMxedgePortStat{
                FullDuplex:           models.ToPointer(true),
                Mac:                  models.ToPointer("9e294e49091d"),
                RxBytes:              models.NewOptional(models.ToPointer(int64(646898375700))),
                RxErrors:             models.ToPointer(0),
                RxPkts:               models.NewOptional(models.ToPointer(int64(8784449574))),
                Speed:                models.ToPointer(10000),
                State:                models.ToPointer("forwarding"),
                TxBytes:              models.NewOptional(models.ToPointer(int64(647200748038))),
                TxErrors:             models.ToPointer(0),
                TxPkts:               models.NewOptional(models.ToPointer(int64(8788647466))),
                Up:                   models.ToPointer(true),
            },
            "port1": models.StatsMxedgePortStat{
                FullDuplex:           models.ToPointer(true),
                Mac:                  models.ToPointer("a270fe53437e"),
                RxBytes:              models.NewOptional(models.ToPointer(int64(647200437652))),
                RxErrors:             models.ToPointer(0),
                RxPkts:               models.NewOptional(models.ToPointer(int64(8788644886))),
                Speed:                models.ToPointer(10000),
                State:                models.ToPointer("forwarding"),
                TxBytes:              models.NewOptional(models.ToPointer(int64(646898681650))),
                TxErrors:             models.ToPointer(0),
                TxPkts:               models.NewOptional(models.ToPointer(int64(8784452092))),
                Up:                   models.ToPointer(true),
            },
        },
        ServiceStat:          map[string]models.StatsMxedgeServiceStat{
            "mxagent": models.StatsMxedgeServiceStat{
                ExtIp:                models.ToPointer("99.0.86.164"),
                LastSeen:             models.ToPointer(float64(1633721215)),
                PackageState:         models.ToPointer("Installed"),
                PackageVersion:       models.ToPointer("3.1.1037-1"),
                RunningState:         models.ToPointer("Running"),
                Uptime:               models.ToPointer(21240),
            },
            "tunterm": models.StatsMxedgeServiceStat{
                ExtIp:                models.ToPointer("99.0.86.164"),
                LastSeen:             models.ToPointer(float64(1633721203)),
                PackageState:         models.ToPointer("Installed"),
                PackageVersion:       models.ToPointer("0.1.2449+deb10"),
                RunningState:         models.ToPointer("Running"),
                Uptime:               models.ToPointer(76261),
            },
        },
        Services:             []string{
            "tunterm",
        },
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Status:               models.ToPointer("connected"),
        TuntermRegistered:    models.ToPointer(true),
        Uptime:               models.ToPointer(76281),
        VirtualizationType:   models.ToPointer("KVM"),
    }

}
```

