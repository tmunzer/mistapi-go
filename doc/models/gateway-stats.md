
# Gateway Stats

Gateway statistics

## Structure

`GatewayStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.ApRedundancy`](../../doc/models/ap-redundancy.md) | Optional | - |
| `ClusterStat` | [`map[string]models.GatewayStatsClusterStat`](../../doc/models/gateway-stats-cluster-stat.md) | Optional | - |
| `Cpu2Stat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `Dhcpd2Stat` | [`map[string]models.GatewayStatsDhcpdStatLan`](../../doc/models/gateway-stats-dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `DhcpdStat` | [`map[string]models.GatewayStatsDhcpdStatLan`](../../doc/models/gateway-stats-dhcpd-stat-lan.md) | Optional | Property key is the network name |
| `Hostname` | `*string` | Optional | hostname reported by the device |
| `Ip` | `*string` | Optional | IP address |
| `Ip2Stat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `LastSeen` | `*float64` | Optional | last seen timestamp |
| `Mac` | `string` | Required | device mac |
| `Memory2Stat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `MemoryStat` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `Model` | `*string` | Optional | device model |
| `Module2Stat` | [`[]models.ModuleStat`](../../doc/models/module-stat.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `ModuleStat` | [`[]models.ModuleStat`](../../doc/models/module-stat.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Name` | `*string` | Optional | device name if configured |
| `Serial` | `*string` | Optional | serial |
| `Spu2Stat` | [`*models.GatewayStatsSpuStat`](../../doc/models/gateway-stats-spu-stat.md) | Optional | - |
| `SpuStat` | [`*models.GatewayStatsSpuStat`](../../doc/models/gateway-stats-spu-stat.md) | Optional | - |
| `Status` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Uptime` | `*float64` | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "hostname": "sj1",
  "ip": "10.2.11.137",
  "last_seen": 1553203563,
  "mac": "dc38e1dbf3cd",
  "model": "SRX320",
  "name": "sj1",
  "serial": "TC3714190003",
  "status": "connected",
  "type": "gateway",
  "uptime": 3671219,
  "version": "18.4R1.8",
  "ap_redundancy": {
    "modules": {
      "key0": {
        "num_aps": 2,
        "num_aps_with_switch_redundancy": 254
      }
    },
    "num_aps": 246,
    "num_aps_with_switch_redundancy": 10
  },
  "cluster_stat": {
    "key0": {
      "status": "status6"
    },
    "key1": {
      "status": "status6"
    },
    "key2": {
      "status": "status6"
    }
  },
  "cpu2_stat": {
    "idle": 10.78,
    "interrupt": 124.54,
    "load_avg": [
      14.61
    ],
    "system": 151.1,
    "user": 113.22
  },
  "cpu_stat": {
    "idle": 102.08,
    "interrupt": 215.84,
    "load_avg": [
      105.91
    ],
    "system": 13.6,
    "user": 204.52
  },
  "dhcpd2_stat": {
    "key0": {
      "num_ips": 180,
      "num_leased": 236
    }
  }
}
```

