
# Stats Device

## Structure

`StatsDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoPlacement` | [`*models.ApStatsAutoPlacement`](../../doc/models/ap-stats-auto-placement.md) | Optional | - |
| `AutoUpgradeStat` | [`*models.ApStatsAutoUpgrade`](../../doc/models/ap-stats-auto-upgrade.md) | Optional | - |
| `BleStat` | [`*models.ApStatsBle`](../../doc/models/ap-stats-ble.md) | Optional | - |
| `CertExpiry` | `models.Optional[int64]` | Optional | - |
| `ConfigReverted` | `models.Optional[bool]` | Optional | - |
| `CpuSystem` | `models.Optional[int64]` | Optional | - |
| `CpuUtil` | `models.Optional[int]` | Optional | - |
| `CreatedTime` | `models.Optional[int64]` | Optional | - |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `EnvStat` | [`*models.ApStatsEnvStat`](../../doc/models/ap-stats-env-stat.md) | Optional | device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage |
| `EslStat` | [`models.Optional[models.ApStatsEslStat]`](../../doc/models/ap-stats-esl-stat.md) | Optional | - |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional | - |
| `ExtIp` | `models.Optional[string]` | Optional | IP address |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `HwRev` | `models.Optional[string]` | Optional | device hardware revision number |
| `Id` | `*uuid.UUID` | Optional | serial |
| `InactiveWiredVlans` | `[]int` | Optional | - |
| `IotStat` | [`map[string]models.ApStatsIotStatAdditionalProperties`](../../doc/models/ap-stats-iot-stat-additional-properties.md) | Optional | - |
| `Ip` | `models.Optional[string]` | Optional | IP address |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | IP AP settings |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `L2tpStat` | [`map[string]models.ApStatsL2TpStat`](../../doc/models/ap-stats-l2-tp-stat.md) | Optional | l2tp tunnel status (key is the wxtunnel_id) |
| `LastSeen` | `models.Optional[float64]` | Optional | last seen timestamp |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | last trouble code of switch |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `LldpStat` | [`*models.ApStatsLldpStat`](../../doc/models/ap-stats-lldp-stat.md) | Optional | LLDP Stat (neighbor information, power negotiations) |
| `Locating` | `models.Optional[bool]` | Optional | - |
| `Locked` | `models.Optional[bool]` | Optional | whether this AP is considered locked (placement / orientation has been vetted) |
| `Mac` | `models.Optional[string]` | Optional | device mac |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | serial |
| `MemUsedKb` | `models.Optional[int64]` | Optional | - |
| `MeshDownlinks` | [`map[string]models.ApStatMeshDownlink`](../../doc/models/ap-stat-mesh-downlink.md) | Optional | Property key is the mesh downlink id (e.g `00000000-0000-0000-1000-5c5b35000010`) |
| `MeshUplink` | [`*models.ApStatMeshUplink`](../../doc/models/ap-stat-mesh-uplink.md) | Optional | - |
| `Model` | `models.Optional[string]` | Optional | device model |
| `ModifiedTime` | `models.Optional[int64]` | Optional | - |
| `Mount` | `models.Optional[string]` | Optional | - |
| `Name` | `models.Optional[string]` | Optional | device name if configured |
| `Notes` | `models.Optional[string]` | Optional | - |
| `NumClients` | `models.Optional[int]` | Optional | how many wireless clients are currently connected |
| `OrgId` | `models.Optional[uuid.UUID]` | Optional | serial |
| `PortStat` | [`models.Optional[map[string]models.ApStatsPortStat]`](../../doc/models/ap-stats-port-stat.md) | Optional | Property key is the port name (e.g. `eth0`) |
| `PowerBudget` | `models.Optional[int]` | Optional | in mW, surplus if positive or deficit if negative |
| `PowerConstrained` | `models.Optional[bool]` | Optional | whether insufficient power |
| `PowerOpmode` | `models.Optional[string]` | Optional | constrained mode |
| `PowerSrc` | `models.Optional[string]` | Optional | DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown) |
| `RadioConfig` | [`*models.ApStatsRadioConfig`](../../doc/models/ap-stats-radio-config.md) | Optional | - |
| `RadioStat` | [`*models.ApStatsRadioStat`](../../doc/models/ap-stats-radio-stat.md) | Optional | - |
| `RxBps` | `models.Optional[float64]` | Optional | - |
| `RxBytes` | `models.Optional[int64]` | Optional | - |
| `RxPkts` | `models.Optional[int]` | Optional | - |
| `Serial` | `models.Optional[string]` | Optional | serial |
| `SiteId` | `*uuid.UUID` | Optional | serial |
| `Status` | `models.Optional[string]` | Optional | - |
| `SwitchRedundancy` | [`*models.ApStatsSwitchRedundancy`](../../doc/models/ap-stats-switch-redundancy.md) | Optional | - |
| `TxBps` | `models.Optional[float64]` | Optional | - |
| `TxBytes` | `models.Optional[float64]` | Optional | - |
| `TxPkts` | `models.Optional[float64]` | Optional | - |
| `Type` | `models.Optional[string]` | Optional | device type, ap / ble /switch / gateway |
| `Uptime` | `models.Optional[float64]` | Optional | how long, in seconds, has the device been up (or rebooted) |
| `UsbStat` | [`*models.ApStatsUsbStat`](../../doc/models/ap-stats-usb-stat.md) | Optional | - |
| `Version` | `models.Optional[string]` | Optional | - |
| `X` | `models.Optional[float64]` | Optional | - |
| `Y` | `models.Optional[float64]` | Optional | - |
| `ApRedundancy` | [`*models.SwitchStatsApRedundancy1`](../../doc/models/switch-stats-ap-redundancy-1.md) | Optional | - |
| `ArpTableStats` | [`*models.ArpTableStats`](../../doc/models/arp-table-stats.md) | Optional | - |
| `Clients` | [`[]models.SwitchStatsClientItem`](../../doc/models/switch-stats-client-item.md) | Optional | - |
| `ClientsStats` | [`*models.SwitchStatsClientsStats`](../../doc/models/switch-stats-clients-stats.md) | Optional | - |
| `ConfigStatus` | `*string` | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `DhcpdStat` | [`*models.DhcpdStat`](../../doc/models/dhcpd-stat.md) | Optional | Property key is the network name |
| `FwVersionsOutofsync` | `*bool` | Optional | - |
| `HasPcap` | `*bool` | Optional | whether the switch supports packet capture |
| `Hostname` | `*string` | Optional | hostname reported by the device |
| `IfStat` | [`*models.IfStat`](../../doc/models/if-stat.md) | Optional | Property key is the interface name |
| `MacTableStats` | [`*models.MacTableStats`](../../doc/models/mac-table-stats.md) | Optional | - |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `ModuleStat` | [`[]models.ModuleStatItem`](../../doc/models/module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `RouteSummaryStats` | [`*models.RouteSummaryStats`](../../doc/models/route-summary-stats.md) | Optional | - |
| `ServiceStat` | [`*models.ServiceStat`](../../doc/models/service-stat.md) | Optional | - |
| `VcMac` | `models.Optional[string]` | Optional | - |
| `VcSetupInfo` | [`*models.SwitchStatsVcSetupInfo`](../../doc/models/switch-stats-vc-setup-info.md) | Optional | - |
| `ClusterConfig` | [`*models.ClusterConfigStats`](../../doc/models/cluster-config-stats.md) | Optional | - |
| `ClusterStat` | [`*models.GatewayStatsCluster`](../../doc/models/gateway-stats-cluster.md) | Optional | - |
| `ConductorName` | `*string` | Optional | - |
| `Cpu2Stat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `Dhcpd2Stat` | [`map[string]models.DhcpdStatLan`](../../doc/models/dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `If2Stat` | [`map[string]models.IfStatProperty`](../../doc/models/if-stat-property.md) | Optional | Property key is the interface name |
| `Ip2Stat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `IsHa` | `models.Optional[bool]` | Optional | - |
| `Memory2Stat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `Module2Stat` | [`[]models.ModuleStatItem`](../../doc/models/module-stat-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `NodeName` | `*string` | Optional | - |
| `RouterName` | `*string` | Optional | device name if configured |
| `Service2Stat` | [`map[string]models.ServiceStatProperty`](../../doc/models/service-stat-property.md) | Optional | - |
| `ServiceStatus` | [`*models.GatewayStatsServiceStatus`](../../doc/models/gateway-stats-service-status.md) | Optional | - |
| `Spu2Stat` | [`[]models.GatewayStatsSpuItem`](../../doc/models/gateway-stats-spu-item.md) | Optional | - |
| `SpuStat` | [`[]models.GatewayStatsSpuItem`](../../doc/models/gateway-stats-spu-item.md) | Optional | - |

## Example (as JSON)

```json
{
  "cert_expiry": 1534534392,
  "ext_ip": "73.92.124.103",
  "iot_stat": {
    "DI2": {
      "value": 0
    }
  },
  "ip": "10.2.9.159",
  "last_seen": 1470417522,
  "locating": false,
  "locked": true,
  "mac": "5c5b35000010",
  "map_id": "63eda950-c6da-11e4-a628-60f81dd250cc",
  "model": "AP200",
  "mount": "faceup",
  "name": "conference room",
  "power_budget": 1000,
  "power_constrained": false,
  "power_opmode": "[20] 6GHz(2x2) 5GHz(4x4) 2.4GHz(2x2).",
  "power_src": "PoE 802.3af",
  "rx_bps": 60003,
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "serial": "FXLH2015170017",
  "status": "connected",
  "tx_bps": 634301,
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "type": "ap",
  "uptime": 13500,
  "version": "0.14.12345",
  "x": 53.5,
  "y": 173.1,
  "has_pcap": false,
  "hostname": "sj-sw1",
  "node_name": "node0",
  "router_name": "sj1",
  "auto_placement": {
    "info": {
      "cluster_number": 112,
      "orientation_stats": 90,
      "probability_surface": {
        "radius": 74.96,
        "radius_m": 19.46,
        "x": 93.54
      }
    },
    "recommended_anchor": false,
    "status": "status4",
    "status_detail": "status_detail0",
    "use_auto_placement": false
  },
  "auto_upgrade_stat": {
    "lastcheck": 28
  },
  "ble_stat": {
    "beacon_enabled": false,
    "beacon_rate": 78,
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 132,
    "eddystone_uid_instance": "eddystone_uid_instance6"
  },
  "config_reverted": false
}
```

