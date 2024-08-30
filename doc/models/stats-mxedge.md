
# Stats Mxedge

## Structure

`StatsMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CpuStat` | [`*models.StatsMxedgeCpuStat`](../../doc/models/stats-mxedge-cpu-stat.md) | Optional | CPU/core stats list |
| `CreatedTime` | `*float64` | Optional | - |
| `FipsEnabled` | `*bool` | Optional | alue indicating fips configuration on the device |
| `ForSite` | `*bool` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `IpStat` | [`*models.StatsMxedgeIpStat`](../../doc/models/stats-mxedge-ip-stat.md) | Optional | OOBM IP stats |
| `LagStat` | [`map[string]models.StatsMxedgeLagStat`](../../doc/models/stats-mxedge-lag-stat.md) | Optional | Stat for LAG (Link Aggregation Group). Property key is the LAG name |
| `LastSeen` | `*float64` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `MemoryStat` | [`*models.StatsMxedgeMemoryStat`](../../doc/models/stats-mxedge-memory-stat.md) | Optional | Memory usage |
| `Model` | `*string` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `MxagentRegistered` | `*bool` | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | The name of the tunnel |
| `NumTunnels` | `*int` | Optional | - |
| `OobIpConfig` | [`*models.MxedgeOobIpConfig`](../../doc/models/mxedge-oob-ip-config.md) | Optional | ip configuration of the Mist Edge out-of_band management interface |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortStat` | [`map[string]models.StatsMxedgePortStat`](../../doc/models/stats-mxedge-port-stat.md) | Optional | - |
| `SensorStat` | `*interface{}` | Optional | - |
| `Serial` | `models.Optional[string]` | Optional | - |
| `ServiceStat` | [`map[string]models.StatsMxedgeServiceStat`](../../doc/models/stats-mxedge-service-stat.md) | Optional | stat for each services |
| `Services` | `[]string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | `*string` | Optional | - |
| `TuntermIpConfig` | [`*models.StatsMxedgeTuntermIpConfig`](../../doc/models/stats-mxedge-tunterm-ip-config.md) | Optional | - |
| `TuntermPortConfig` | [`*models.StatsMxedgeTuntermPortConfig`](../../doc/models/stats-mxedge-tunterm-port-config.md) | Optional | - |
| `TuntermRegistered` | `*bool` | Optional | - |
| `TuntermStat` | [`*models.StatsMxedgeTuntermStat`](../../doc/models/stats-mxedge-tunterm-stat.md) | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `VirtualizationType` | `*string` | Optional | Virtualization environment |

## Example (as JSON)

```json
{
  "created_time": 1632684398.0,
  "for_site": false,
  "id": "00000000-0000-0000-1000-020000a80cb4",
  "lag_stat": {
    "lacp0": {
      "active_ports": [
        "port0",
        "port1"
      ]
    }
  },
  "last_seen": 1633721215,
  "mac": "020000a80cb4",
  "model": "ME-VM",
  "modified_time": 1633643629,
  "mxagent_registered": true,
  "mxcluster_id": "678bc339-7635-4556-bbc0-e77ad493ef8b",
  "name": "me-vm-1",
  "num_tunnels": 0,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "port_stat": {
    "port0": {
      "full_duplex": true,
      "mac": "9e294e49091d",
      "rx_bytes": 646898375700,
      "rx_errors": 0,
      "rx_pkts": 8784449574,
      "speed": 10000,
      "state": "forwarding",
      "tx_bytes": 647200748038,
      "tx_errors": 0,
      "tx_pkts": 8788647466,
      "up": true
    },
    "port1": {
      "full_duplex": true,
      "mac": "a270fe53437e",
      "rx_bytes": 647200437652,
      "rx_errors": 0,
      "rx_pkts": 8788644886,
      "speed": 10000,
      "state": "forwarding",
      "tx_bytes": 646898681650,
      "tx_errors": 0,
      "tx_pkts": 8784452092,
      "up": true
    }
  },
  "service_stat": {
    "mxagent": {
      "ext_ip": "99.0.86.164",
      "last_seen": 1633721215,
      "package_state": "Installed",
      "package_version": "3.1.1037-1",
      "running_state": "Running",
      "uptime": 21240
    },
    "tunterm": {
      "ext_ip": "99.0.86.164",
      "last_seen": 1633721203,
      "package_state": "Installed",
      "package_version": "0.1.2449+deb10",
      "running_state": "Running",
      "uptime": 76261
    }
  },
  "services": [
    "tunterm"
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "status": "connected",
  "tunterm_registered": true,
  "uptime": 76281,
  "virtualization_type": "KVM",
  "cpu_stat": {
    "cpus": {
      "key0": {
        "idle": 4.8,
        "interrupt": 118.56,
        "load_avg": [
          8.63
        ],
        "system": 110.88,
        "user": 107.24
      }
    },
    "idle": 224,
    "interrupt": 80,
    "system": 80,
    "usage": 46
  },
  "fips_enabled": false
}
```
