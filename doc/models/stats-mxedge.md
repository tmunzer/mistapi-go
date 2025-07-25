
# Stats Mxedge

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CpuStat` | [`*models.StatsMxedgeCpuStat`](../../doc/models/stats-mxedge-cpu-stat.md) | Optional | CPU/core stats list |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `FipsEnabled` | `*bool` | Optional | Indicate fips configuration on the device |
| `ForSite` | `*bool` | Optional | - |
| `Fwupdate` | [`*models.FwupdateStat`](../../doc/models/fwupdate-stat.md) | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `IdracVersion` | `*string` | Optional | IDRAC version of the mist edge device |
| `IpStat` | [`*models.StatsMxedgeIpStat`](../../doc/models/stats-mxedge-ip-stat.md) | Optional | IP stats |
| `LagStat` | [`map[string]models.StatsMxedgeLagStat`](../../doc/models/stats-mxedge-lag-stat.md) | Optional | Stat for LAG (Link Aggregation Group). Property key is the LAG name |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `*string` | Optional | - |
| `MemoryStat` | [`*models.StatsMxedgeMemoryStat`](../../doc/models/stats-mxedge-memory-stat.md) | Optional | Memory usage |
| `Model` | `*string` | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `MxagentRegistered` | `*bool` | Optional | - |
| `MxclusterId` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | The name of the tunnel |
| `NumTunnels` | `*int` | Optional | - |
| `OobIpConfig` | [`*models.MxedgeOobIpConfig`](../../doc/models/mxedge-oob-ip-config.md) | Optional | IPconfiguration of the Mist Edge out-of_band management interface |
| `OobIpStat` | [`*models.StatsMxedgeOobIpStat`](../../doc/models/stats-mxedge-oob-ip-stat.md) | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortStat` | [`map[string]models.StatsMxedgePortStat`](../../doc/models/stats-mxedge-port-stat.md) | Optional | - |
| `Serial` | `models.Optional[string]` | Optional | - |
| `ServiceStat` | [`map[string]models.StatsMxedgeServiceStat`](../../doc/models/stats-mxedge-service-stat.md) | Optional | Stat for each services |
| `Services` | `[]string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Status` | `*string` | Optional | - |
| `TuntermIpConfig` | [`*models.StatsMxedgeTuntermIpConfig`](../../doc/models/stats-mxedge-tunterm-ip-config.md) | Optional | - |
| `TuntermPortConfig` | [`*models.StatsMxedgeTuntermPortConfig`](../../doc/models/stats-mxedge-tunterm-port-config.md) | Optional | - |
| `TuntermRegistered` | `*bool` | Optional | - |
| `TuntermStat` | [`*models.StatsMxedgeTuntermStat`](../../doc/models/stats-mxedge-tunterm-stat.md) | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `VirtualizationType` | `*string` | Optional | Virtualization environment |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "for_site": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "idrac_version": "7.00.00.00",
  "lag_stat": {
    "lacp0": {
      "active_ports": [
        "port0",
        "port1"
      ]
    }
  },
  "last_seen": 1470417522,
  "mac": "020000a80cb4",
  "model": "ME-VM",
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
      "key0": null
    },
    "idle": 224,
    "interrupt": 80,
    "system": 80,
    "usage": 46,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "created_time": 71.86,
  "fips_enabled": false,
  "fwupdate": {
    "progress": 100,
    "status": "inprogress",
    "status_id": 70,
    "timestamp": 147.68,
    "will_retry": false,
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

