
# Switch Stats

Switch statistics

## Structure

`SwitchStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApRedundancy` | [`*models.SwitchStatsApRedundancy`](../../doc/models/switch-stats-ap-redundancy.md) | Optional | - |
| `Clients` | [`[]models.SwitchStatsClient`](../../doc/models/switch-stats-client.md) | Optional | - |
| `CpuStat` | [`*models.CpuStat`](../../doc/models/cpu-stat.md) | Optional | - |
| `HasPcap` | `*bool` | Optional | whether the switch supports packet capture |
| `Hostname` | `*string` | Optional | hostname reported by the device |
| `IfStat` | [`map[string]models.SwitchStatsIfStat`](../../doc/models/switch-stats-if-stat.md) | Optional | Property key is the interface name |
| `Ip` | `*string` | Optional | - |
| `IpStat` | [`*models.IpStat`](../../doc/models/ip-stat.md) | Optional | - |
| `LastSeen` | `*int` | Optional | - |
| `LastTrouble` | [`*models.LastTrouble`](../../doc/models/last-trouble.md) | Optional | last trouble code of switch |
| `Mac` | [`*models.MemoryStat`](../../doc/models/memory-stat.md) | Optional | memory usage stat (for virtual chassis, memory usage of master RE) |
| `Model` | `*string` | Optional | - |
| `ModuleStat` | [`[]models.ModuleStat`](../../doc/models/module-stat.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Name` | `*string` | Optional | device name if configured |
| `NumClients` | [`*models.SwitchStatsNumClients`](../../doc/models/switch-stats-num-clients.md) | Optional | - |
| `Serial` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Uptime` | `*float64` | Optional | - |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "has_pcap": false,
  "hostname": "sj-sw1",
  "ip": "10.2.11.137",
  "last_seen": 1553203563,
  "model": "EX4600",
  "name": "sj-sw1",
  "serial": "TC3714190003",
  "status": "connected",
  "type": "switch",
  "uptime": 13501,
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
  "clients": [
    {
      "device_mac": "device_mac2",
      "hostname": "hostname6",
      "mac": "mac2",
      "port_id": "port_id8"
    },
    {
      "device_mac": "device_mac2",
      "hostname": "hostname6",
      "mac": "mac2",
      "port_id": "port_id8"
    },
    {
      "device_mac": "device_mac2",
      "hostname": "hostname6",
      "mac": "mac2",
      "port_id": "port_id8"
    }
  ],
  "cpu_stat": {
    "idle": 102.08,
    "interrupt": 215.84,
    "load_avg": [
      105.91
    ],
    "system": 13.6,
    "user": 204.52
  }
}
```

