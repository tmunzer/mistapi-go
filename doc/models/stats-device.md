
# Stats Device

## Structure

`StatsDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoPlacement` | [`*models.ApStatsAutoPlacement`](../../doc/models/ap-stats-auto-placement.md) | Optional | - |
| `BleConfig` | [`*models.ApStatsBleConfig`](../../doc/models/ap-stats-ble-config.md) | Optional | - |
| `BleStat` | [`*models.ApStatsBleStat`](../../doc/models/ap-stats-ble-stat.md) | Optional | - |
| `CertExpiry` | `*float64` | Optional | - |
| `EnvStat` | [`*models.ApStatsEnvStat`](../../doc/models/ap-stats-env-stat.md) | Optional | device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage |
| `EslStat` | [`*models.ApStatsEslStat`](../../doc/models/ap-stats-esl-stat.md) | Optional | - |
| `ExtIp` | `*string` | Optional | - |
| `Fwupdate` | [`*models.ApStatsFwupdate`](../../doc/models/ap-stats-fwupdate.md) | Optional | - |
| `IotStat` | [`map[string]models.ApStatsIotStatAdditionalProperties`](../../doc/models/ap-stats-iot-stat-additional-properties.md) | Optional | - |
| `Ip` | `*string` | Optional | IP address |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | IP AP settings |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `L2tpStat` | [`map[string]models.ApStatsL2TpStat`](../../doc/models/ap-stats-l2-tp-stat.md) | Optional | l2tp tunnel status (key is the wxtunnel_id) |
| `LastSeen` | `*float64` | Optional | last seen timestamp |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | last trouble code of switch |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `LldpStat` | [`*models.ApStatsLldpStat`](../../doc/models/ap-stats-lldp-stat.md) | Optional | LLDP Stat (neighbor information, power negotiations) |
| `Locating` | `*bool` | Optional | - |
| `Locked` | `*bool` | Optional | whether this AP is considered locked (placement / orientation has been vetted) |
| `Mac` | [`*models.MemoryStat1`](../../doc/models/memory-stat-1.md) | Optional | device mac |
| `MapId` | `*uuid.UUID` | Optional | - |
| `MeshDownlinks` | [`map[string]models.ApStatMeshDownlink`](../../doc/models/ap-stat-mesh-downlink.md) | Optional | Property key is the mesh downlink id (e.g `00000000-0000-0000-1000-5c5b35000010`) |
| `MeshUplink` | [`*models.ApStatMeshUplink`](../../doc/models/ap-stat-mesh-uplink.md) | Optional | - |
| `Model` | `*string` | Optional | device model |
| `Mount` | `*string` | Optional | - |
| `Name` | `*string` | Optional | device name if configured |
| `NumClients` | [`*models.SwitchStatsNumClients1`](../../doc/models/switch-stats-num-clients-1.md) | Optional | how many wireless clients are currently connected |
| `PortStat` | [`map[string]models.ApStatsPortStat`](../../doc/models/ap-stats-port-stat.md) | Optional | Property key is the port name (e.g. `eth0`) |
| `PowerBudget` | `*float64` | Optional | in mW, surplus if positive or deficit if negative |
| `PowerConstrained` | `*bool` | Optional | whether insufficient power |
| `PowerOpmode` | `*string` | Optional | constrained mode |
| `PowerSrc` | `*string` | Optional | DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown) |
| `RadioConfig` | [`*models.ApStatsRadioConfig`](../../doc/models/ap-stats-radio-config.md) | Optional | - |
| `RadioStat` | [`*models.ApStatsRadioStat`](../../doc/models/ap-stats-radio-stat.md) | Optional | - |
| `RxBps` | `*float64` | Optional | - |
| `RxBytes` | `*int64` | Optional | - |
| `RxPkts` | `*int` | Optional | - |
| `Serial` | `*string` | Optional | serial |
| `Status` | [`*models.ApStatsStatusEnum`](../../doc/models/ap-stats-status-enum.md) | Optional | - |
| `SwitchRedundancy` | [`*models.ApStatsSwitchRedundancy`](../../doc/models/ap-stats-switch-redundancy.md) | Optional | - |
| `TxBps` | `*float64` | Optional | - |
| `TxBytes` | `*float64` | Optional | - |
| `TxPkts` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | device type, ap / ble |
| `Uptime` | `*float64` | Optional | how long, in seconds, has the device been up (or rebooted) |
| `UsbStat` | [`*models.ApStatsUsbStat`](../../doc/models/ap-stats-usb-stat.md) | Optional | - |
| `Version` | `*string` | Optional | - |
| `X` | `*float64` | Optional | - |
| `Y` | `*float64` | Optional | - |
| `ApRedundancy` | [`*models.SwitchStatsApRedundancy1`](../../doc/models/switch-stats-ap-redundancy-1.md) | Optional | - |
| `Clients` | [`[]models.SwitchStatsClient`](../../doc/models/switch-stats-client.md) | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `HasPcap` | `*bool` | Optional | whether the switch supports packet capture |
| `Hostname` | `*string` | Optional | hostname reported by the device |
| `IfStat` | [`map[string]models.SwitchStatsIfStat`](../../doc/models/switch-stats-if-stat.md) | Optional | Property key is the interface name |
| `ModuleStat` | [`[]models.ModuleStat`](../../doc/models/module-stat.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `ClusterStat` | [`map[string]models.GatewayStatsClusterStat`](../../doc/models/gateway-stats-cluster-stat.md) | Optional | - |
| `Cpu2Stat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `Dhcpd2Stat` | [`map[string]models.GatewayStatsDhcpdStatLan`](../../doc/models/gateway-stats-dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `DhcpdStat` | [`map[string]models.GatewayStatsDhcpdStatLan`](../../doc/models/gateway-stats-dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `Ip2Stat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `Memory2Stat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `Module2Stat` | [`[]models.ModuleStat`](../../doc/models/module-stat.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Spu2Stat` | [`*models.GatewayStatsSpuStat`](../../doc/models/gateway-stats-spu-stat.md) | Optional | - |
| `SpuStat` | [`*models.GatewayStatsSpuStat`](../../doc/models/gateway-stats-spu-stat.md) | Optional | - |

## Example (as JSON)

```json
{
  "cert_expiry": 1534534392.0,
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
  "auto_placement": {
    "_id": "_id8",
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
    "status_detail": "status_detail0"
  },
  "ble_config": {
    "beacon_rate": 110,
    "beacon_rate_model": "beacon_rate_model2",
    "beam_disabled": [
      113,
      114,
      115
    ],
    "power": 212,
    "power_mode": "power_mode6"
  },
  "ble_stat": {
    "beacon_rate": 78,
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 132,
    "eddystone_uid_instance": "eddystone_uid_instance6",
    "eddystone_uid_namespace": "eddystone_uid_namespace8"
  },
  "env_stat": {
    "accel_x": 122.78,
    "accel_y": 222.08,
    "accel_z": 112.82,
    "ambient_temp": 66,
    "attitude": 104
  }
}
```

