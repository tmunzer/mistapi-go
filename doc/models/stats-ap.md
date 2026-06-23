
# Stats Ap

AP runtime statistics, placement data, and inventory metadata

## Structure

`StatsAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntennaSelect` | [`*models.AntennaSelectEnum`](../../doc/models/antenna-select-enum.md) | Optional | Antenna Mode for AP which supports selectable antennas. enum: `""` (default), `external`, `internal` |
| `AutoPlacement` | [`*models.StatsApAutoPlacement`](../../doc/models/stats-ap-auto-placement.md) | Optional | Auto placement result and status for an AP |
| `AutoUpgradeStat` | [`*models.StatsApAutoUpgrade`](../../doc/models/stats-ap-auto-upgrade.md) | Optional | Auto-upgrade status for an AP |
| `BleStat` | [`*models.StatsApBle`](../../doc/models/stats-ap-ble.md) | Optional | BLE beacon and traffic statistics reported by an AP |
| `CertExpiry` | `models.Optional[float64]` | Optional, Read-only | Certificate expiry timestamp for the AP, in epoch seconds |
| `ConfigReverted` | `models.Optional[bool]` | Optional, Read-only | Whether the AP configuration was reverted by the device |
| `CpuSystem` | `models.Optional[int64]` | Optional, Read-only | CPU system utilization reported by the AP |
| `CpuUser` | `models.Optional[int]` | Optional, Read-only | CPU user utilization reported by the AP |
| `CpuUtil` | `models.Optional[int]` | Optional, Read-only | Total CPU utilization reported by the AP |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Device profile identifier applied to the AP, when present |
| `EnvStat` | [`*models.StatsApEnvStat`](../../doc/models/stats-ap-env-stat.md) | Optional | Device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage |
| `EslStat` | [`models.Optional[models.StatsApEslStat]`](../../doc/models/stats-ap-esl-stat.md) | Optional, Read-only | Electronic shelf label dongle status reported by an AP |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional, Read-only | EVPN topology identifier associated with the AP, when present |
| `ExpiringCerts` | `map[string]int` | Optional | Map of certificate serial numbers to their expiry timestamps (in epoch) for certificates expiring within 30 days. Property key is the certificate serial number |
| `ExtIp` | `models.Optional[string]` | Optional, Read-only | Public IP address observed for the AP |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | Firmware update status for a device |
| `GpsStat` | [`*models.StatsApGpsStat`](../../doc/models/stats-ap-gps-stat.md) | Optional | GPS location data reported by an AP |
| `HwRev` | `models.Optional[string]` | Optional, Read-only | Hardware revision reported by the AP |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `InactiveWiredVlans` | `[]int` | Optional | List of integer values |
| `IotStat` | [`map[string]models.StatsApIotStatAdditionalProperties`](../../doc/models/stats-ap-iot-stat-additional-properties.md) | Optional | IoT input statistics reported by an AP |
| `Ip` | `models.Optional[string]` | Optional, Read-only | Management IP address currently reported by the AP |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | Management IP addressing settings for an access point |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | Read-only IP addressing status reported by a device interface |
| `L2tpStat` | [`map[string]models.StatsApL2tpStat`](../../doc/models/stats-ap-l2-tp-stat.md) | Optional | L2TP tunnel status (key is the wxtunnel_id) |
| `LastSeen` | `models.Optional[float64]` | Optional, Read-only | Timestamp indicating when the entity was last seen |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | Last trouble indicator reported by a switch |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | Indicator light settings for an access point |
| `LldpStat` | [`*models.StatsApLldpStat`](../../doc/models/stats-ap-lldp-stat.md) | Optional | LLDP neighbor information and power negotiations. For backward compatibility, when multiple neighbors exist, only information from the first neighbor is displayed. |
| `LldpStats` | [`map[string]models.StatsApLldpStat`](../../doc/models/stats-ap-lldp-stat.md) | Optional | Property key is the port name (e.g. "eth0", "eth1", ...). Map of ethernet ports to their respective LLDP neighbor information and power negotiations. Only present when multiple neighbors exist. |
| `Locating` | `models.Optional[bool]` | Optional, Read-only | Whether AP locating mode is currently active |
| `Locked` | `models.Optional[bool]` | Optional, Read-only | Whether this AP is considered locked (placement / orientation has been vetted) |
| `Mac` | `models.Optional[string]` | Optional, Read-only | AP MAC address reported by Mist |
| `MapId` | `models.Optional[uuid.UUID]` | Optional, Read-only | Map identifier where the AP is placed, when available |
| `MemTotalKb` | `models.Optional[int64]` | Optional, Read-only | Total memory available on the AP, in kilobytes |
| `MemUsedKb` | `models.Optional[int64]` | Optional, Read-only | Memory currently used on the AP, in kilobytes |
| `MeshDownlinks` | [`map[string]models.ApStatMeshDownlink`](../../doc/models/ap-stat-mesh-downlink.md) | Optional | Property key is the mesh downlink id (e.g. `00000000-0000-0000-1000-5c5b35000010`) |
| `MeshUplink` | [`*models.ApStatMeshUplink`](../../doc/models/ap-stat-mesh-uplink.md) | Optional | Runtime statistics for this AP's mesh uplink |
| `Model` | `models.Optional[string]` | Optional, Read-only | AP model identifier reported by Mist |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Mount` | `models.Optional[string]` | Optional, Read-only | Orientation configured for the AP mount |
| `Name` | `models.Optional[string]` | Optional, Read-only | Display name of the AP |
| `Notes` | `models.Optional[string]` | Optional, Read-only | Free-form notes configured on the AP record |
| `NumClients` | `models.Optional[int]` | Optional, Read-only | How many wireless clients are currently connected |
| `NumWlans` | `*int` | Optional | How many WLANs are applied to the device |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PortStat` | [`models.Optional[map[string]models.StatsApPortStat]`](../../doc/models/stats-ap-port-stat.md) | Optional, Read-only | Property key is the port name (e.g. `eth0`) |
| `PowerBudget` | `models.Optional[int]` | Optional, Read-only | In mW, surplus if positive or deficit if negative |
| `PowerConstrained` | `models.Optional[bool]` | Optional, Read-only | Whether the AP is operating with insufficient power |
| `PowerOpmode` | `models.Optional[string]` | Optional, Read-only | Operating mode used while AP power is constrained |
| `PowerSrc` | `models.Optional[string]` | Optional, Read-only | DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown) |
| `RadioConfig` | [`*models.StatsApRadioConfig`](../../doc/models/stats-ap-radio-config.md) | Optional | Radio configuration reported by an AP |
| `RadioStat` | [`*models.StatsApRadioStat`](../../doc/models/stats-ap-radio-stat.md) | Optional | Per-band radio statistics reported by an AP |
| `RxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of receiving traffic, bits/seconds, last known |
| `RxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets received since connection |
| `Serial` | `models.Optional[string]` | Optional, Read-only | Device serial number for the AP |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Status` | `models.Optional[string]` | Optional, Read-only | Connection status reported for the AP |
| `SwitchRedundancy` | [`*models.StatsApSwitchRedundancy`](../../doc/models/stats-ap-switch-redundancy.md) | Optional | Switch redundancy status reported by an AP |
| `TxBps` | `models.Optional[int64]` | Optional, Read-only | Rate of transmitting traffic, bits/seconds, last known |
| `TxBytes` | `models.Optional[int64]` | Optional, Read-only | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional, Read-only | Amount of packets sent since connection |
| `Type` | `string` | Required, Constant, Read-only | Device Type. enum: `ap`<br><br>**Value**: `"ap"` |
| `Uptime` | `models.Optional[float64]` | Optional, Read-only | How long, in seconds, has the device been up (or rebooted) |
| `UsbStat` | [`*models.StatsApUsbStat`](../../doc/models/stats-ap-usb-stat.md) | Optional | USB peripheral status reported by an AP |
| `Version` | `models.Optional[string]` | Optional, Read-only | Firmware version running on the AP |
| `X` | `models.Optional[float64]` | Optional, Read-only | Map X coordinate of the AP placement, in pixels |
| `Y` | `models.Optional[float64]` | Optional, Read-only | Map Y coordinate of the AP placement, in pixels |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    statsAp := models.StatsAp{
        AntennaSelect:        models.ToPointer(models.AntennaSelectEnum_EXTERNAL),
        AutoPlacement:        models.ToPointer(models.StatsApAutoPlacement{
            Info:                 models.ToPointer(models.StatsApAutoPlacementInfo{
                ClusterNumber:        models.ToPointer(112),
                OrientationStats:     models.ToPointer(90),
                ProbabilitySurface:   models.ToPointer(models.StatsApAutoPlacementInfoProbabilitySurface{
                    Radius:               models.ToPointer(float64(74.96)),
                    RadiusM:              models.ToPointer(float64(19.46)),
                    X:                    models.ToPointer(float64(93.54)),
                }),
            }),
            RecommendedAnchor:    models.ToPointer(false),
            Status:               models.ToPointer("status4"),
            StatusDetail:         models.ToPointer("status_detail0"),
            X:                    models.ToPointer(float64(15)),
        }),
        AutoUpgradeStat:      models.ToPointer(models.StatsApAutoUpgrade{
            Lastcheck:            models.NewOptional(models.ToPointer(int64(28))),
        }),
        BleStat:              models.ToPointer(models.StatsApBle{
            BeaconEnabled:         models.NewOptional(models.ToPointer(false)),
            BeaconRate:            models.NewOptional(models.ToPointer(78)),
            EddystoneUidEnabled:   models.NewOptional(models.ToPointer(false)),
            EddystoneUidFreqMsec:  models.NewOptional(models.ToPointer(132)),
            EddystoneUidInstance:  models.NewOptional(models.ToPointer("eddystone_uid_instance6")),
        }),
        CertExpiry:           models.NewOptional(models.ToPointer(float64(1534534392))),
        ExtIp:                models.NewOptional(models.ToPointer("73.92.124.103")),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        IotStat:              map[string]models.StatsApIotStatAdditionalProperties{
            "DI2": models.StatsApIotStatAdditionalProperties{
                Value:                models.NewOptional(models.ToPointer(0)),
            },
        },
        Ip:                   models.NewOptional(models.ToPointer("10.2.9.159")),
        LastSeen:             models.NewOptional(models.ToPointer(float64(1470417522))),
        Locating:             models.NewOptional(models.ToPointer(false)),
        Locked:               models.NewOptional(models.ToPointer(true)),
        Mac:                  models.NewOptional(models.ToPointer("5c5b35000010")),
        MapId:                models.NewOptional(models.ToPointer(uuid.MustParse("63eda950-c6da-11e4-a628-60f81dd250cc"))),
        Model:                models.NewOptional(models.ToPointer("AP200")),
        Mount:                models.NewOptional(models.ToPointer("faceup")),
        Name:                 models.NewOptional(models.ToPointer("conference room")),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PowerBudget:          models.NewOptional(models.ToPointer(1000)),
        PowerConstrained:     models.NewOptional(models.ToPointer(false)),
        PowerOpmode:          models.NewOptional(models.ToPointer("[20] 6GHz(2x2) 5GHz(4x4) 2.4GHz(2x2).")),
        PowerSrc:             models.NewOptional(models.ToPointer("PoE 802.3af")),
        RxBps:                models.NewOptional(models.ToPointer(int64(60003))),
        RxBytes:              models.NewOptional(models.ToPointer(int64(8515104416))),
        RxPkts:               models.NewOptional(models.ToPointer(int64(57770567))),
        Serial:               models.NewOptional(models.ToPointer("FXLH2015170017")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        TxBps:                models.NewOptional(models.ToPointer(int64(634301))),
        TxBytes:              models.NewOptional(models.ToPointer(int64(211217389682))),
        TxPkts:               models.NewOptional(models.ToPointer(int64(812204062))),
        Type:                 "ap",
        Uptime:               models.NewOptional(models.ToPointer(float64(13500))),
        Version:              models.NewOptional(models.ToPointer("0.14.12345")),
        X:                    models.NewOptional(models.ToPointer(float64(53.5))),
        Y:                    models.NewOptional(models.ToPointer(float64(173.1))),
    }

}
```

