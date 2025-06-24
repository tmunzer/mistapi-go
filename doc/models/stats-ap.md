
# Stats Ap

AP statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntMode` | [`*models.AntModeEnum`](../../doc/models/ant-mode-enum.md) | Optional | Antenna Mode for AP which supports selectable antennas. enum: `external`, `internal` |
| `AutoPlacement` | [`*models.StatsApAutoPlacement`](../../doc/models/stats-ap-auto-placement.md) | Optional | - |
| `AutoUpgradeStat` | [`*models.StatsApAutoUpgrade`](../../doc/models/stats-ap-auto-upgrade.md) | Optional | - |
| `BleStat` | [`*models.StatsApBle`](../../doc/models/stats-ap-ble.md) | Optional | - |
| `CertExpiry` | `models.Optional[float64]` | Optional | - |
| `ConfigReverted` | `models.Optional[bool]` | Optional | - |
| `CpuSystem` | `models.Optional[int64]` | Optional | - |
| `CpuUtil` | `models.Optional[int]` | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceprofileId` | `models.Optional[uuid.UUID]` | Optional | - |
| `EnvStat` | [`*models.StatsApEnvStat`](../../doc/models/stats-ap-env-stat.md) | Optional | Device environment, including CPU temperature, Ambient temperature, Humidity, Attitude, Pressure, Accelerometers, Magnetometers and vCore Voltage |
| `EslStat` | [`models.Optional[models.StatsApEslStat]`](../../doc/models/stats-ap-esl-stat.md) | Optional | - |
| `EvpntopoId` | `models.Optional[uuid.UUID]` | Optional | - |
| `ExpiringCerts` | `map[string]int` | Optional | Map of certificate serial numbers to their expiry timestamps (in epoch) for certificates expiring within 30 days. Property key is the certificate serial number |
| `ExtIp` | `models.Optional[string]` | Optional | - |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `Gps` | [`*models.StatsApGpsStat`](../../doc/models/stats-ap-gps-stat.md) | Optional | - |
| `HwRev` | `models.Optional[string]` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `InactiveWiredVlans` | `[]int` | Optional | - |
| `IotStat` | [`map[string]models.StatsApIotStatAdditionalProperties`](../../doc/models/stats-ap-iot-stat-additional-properties.md) | Optional | - |
| `Ip` | `models.Optional[string]` | Optional | - |
| `IpConfig` | [`*models.ApIpConfig`](../../doc/models/ap-ip-config.md) | Optional | IP AP settings |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `L2tpStat` | [`map[string]models.StatsApL2tpStat`](../../doc/models/stats-ap-l2-tp-stat.md) | Optional | L2TP tunnel status (key is the wxtunnel_id) |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | Last trouble code of switch |
| `Led` | [`*models.ApLed`](../../doc/models/ap-led.md) | Optional | LED AP settings |
| `LldpStat` | [`*models.StatsApLldpStat`](../../doc/models/stats-ap-lldp-stat.md) | Optional | LLDP neighbor information and power negotiations. For backward compatibility, when multiple neighbors exist, only information from the first neighbor is displayed. |
| `LldpStats` | [`map[string]models.StatsApLldpStat`](../../doc/models/stats-ap-lldp-stat.md) | Optional | Property key is the port name (e.g. "eth0", "eth1", ...). Map of ethernet ports to their respective LLDP neighbor information and power negotiations. Only present when multiple neighbors exist. |
| `Locating` | `models.Optional[bool]` | Optional | - |
| `Locked` | `models.Optional[bool]` | Optional | Whether this AP is considered locked (placement / orientation has been vetted) |
| `Mac` | `*string` | Required | Device mac |
| `MapId` | `models.Optional[uuid.UUID]` | Optional | - |
| `MemUsedKb` | `models.Optional[int64]` | Optional | - |
| `MeshDownlinks` | [`map[string]models.ApStatMeshDownlink`](../../doc/models/ap-stat-mesh-downlink.md) | Optional | Property key is the mesh downlink id (e.g. `00000000-0000-0000-1000-5c5b35000010`) |
| `MeshUplink` | [`*models.ApStatMeshUplink`](../../doc/models/ap-stat-mesh-uplink.md) | Optional | - |
| `Model` | `*string` | Required | Device model |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Mount` | `models.Optional[string]` | Optional | - |
| `Name` | `models.Optional[string]` | Optional | - |
| `Notes` | `models.Optional[string]` | Optional | - |
| `NumClients` | `models.Optional[int]` | Optional | How many wireless clients are currently connected |
| `NumWlans` | `*int` | Optional | How many WLANs are applied to the device |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortStat` | [`models.Optional[map[string]models.StatsApPortStat]`](../../doc/models/stats-ap-port-stat.md) | Optional | Property key is the port name (e.g. `eth0`) |
| `PowerBudget` | `models.Optional[int]` | Optional | In mW, surplus if positive or deficit if negative |
| `PowerConstrained` | `models.Optional[bool]` | Optional | Whether insufficient power |
| `PowerOpmode` | `models.Optional[string]` | Optional | Constrained mode |
| `PowerSrc` | `models.Optional[string]` | Optional | DC Input / PoE 802.3at / PoE 802.3af / LLDP / ? (unknown) |
| `RadioConfig` | [`*models.StatsApRadioConfig`](../../doc/models/stats-ap-radio-config.md) | Optional | - |
| `RadioStat` | [`*models.StatsApRadioStat`](../../doc/models/stats-ap-radio-stat.md) | Optional | - |
| `RxBps` | `models.Optional[int64]` | Optional | Rate of receiving traffic, bits/seconds, last known |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `Serial` | `models.Optional[string]` | Optional | Serial Number |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | `models.Optional[string]` | Optional | - |
| `SwitchRedundancy` | [`*models.StatsApSwitchRedundancy`](../../doc/models/stats-ap-switch-redundancy.md) | Optional | - |
| `TxBps` | `models.Optional[int64]` | Optional | Rate of transmitting traffic, bits/seconds, last known |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `Type` | `string` | Required, Constant | Device Type. enum: `ap`<br><br>**Value**: `"ap"` |
| `Uptime` | `models.Optional[float64]` | Optional | How long, in seconds, has the device been up (or rebooted) |
| `UsbStat` | [`*models.StatsApUsbStat`](../../doc/models/stats-ap-usb-stat.md) | Optional | - |
| `Version` | `models.Optional[string]` | Optional | - |
| `X` | `models.Optional[float64]` | Optional | - |
| `Y` | `models.Optional[float64]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ant_mode": "external",
  "cert_expiry": 1534534392.0,
  "ext_ip": "73.92.124.103",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
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
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "power_budget": 1000,
  "power_constrained": false,
  "power_opmode": "[20] 6GHz(2x2) 5GHz(4x4) 2.4GHz(2x2).",
  "power_src": "PoE 802.3af",
  "rx_bps": 60003,
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "serial": "FXLH2015170017",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "tx_bps": 634301,
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "type": "ap",
  "uptime": 13500,
  "version": "0.14.12345",
  "x": 53.5,
  "y": 173.1,
  "auto_placement": {
    "info": {
      "cluster_number": 112,
      "orientation_stats": 90,
      "probability_surface": {
        "radius": 74.96,
        "radius_m": 19.46,
        "x": 93.54,
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "recommended_anchor": false,
    "status": "status4",
    "status_detail": "status_detail0",
    "x": 15.0,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "auto_upgrade_stat": {
    "lastcheck": 28,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "ble_stat": {
    "beacon_enabled": false,
    "beacon_rate": 78,
    "eddystone_uid_enabled": false,
    "eddystone_uid_freq_msec": 132,
    "eddystone_uid_instance": "eddystone_uid_instance6",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

