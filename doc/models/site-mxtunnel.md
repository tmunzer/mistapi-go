
# Site Mxtunnel

Site MxTunnel

*This model accepts additional fields of type interface{}.*

## Structure

`SiteMxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalMxtunnels` | [`map[string]models.SiteMxtunnelAdditionalMxtunnel`](../../doc/models/site-mxtunnel-additional-mxtunnel.md) | Optional | - |
| `ApSubnets` | `[]string` | Optional | list of subnets where we allow AP to establish Mist Tunnels from |
| `AutoPreemption` | [`*models.AutoPreemption`](../../doc/models/auto-preemption.md) | Optional | schedule to preempt ap’s which are not connected to preferred peer |
| `Clusters` | [`[]models.SiteMxtunnelCluster`](../../doc/models/site-mxtunnel-cluster.md) | Optional | for AP, how to connect to tunterm or radsecproxy |
| `CreatedTime` | `*float64` | Optional | when the object has been created, in epoch |
| `Enabled` | `*bool` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `HelloInterval` | `*int` | Optional | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries<br>**Default**: `60`<br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `*int` | Optional | **Default**: `7`<br>**Constraints**: `>= 2`, `<= 30` |
| `Hosts` | `[]string` | Optional | hostnames or IPs where a Mist Tunnel will use as the Peer (i.e. they are reachable from AP) |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | when the object has been modified for the last time, in epoch |
| `Mtu` | `*int` | Optional | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 1500` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Protocol` | [`*models.MxtunnelProtocolEnum`](../../doc/models/mxtunnel-protocol-enum.md) | Optional | enum: `ip`, `udp`<br>**Default**: `"udp"` |
| `Radsec` | [`*models.SiteMxtunnelRadsec`](../../doc/models/site-mxtunnel-radsec.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `VlanIds` | `[]int` | Optional | list of vlan_ids that will be used |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "hello_interval": 60,
  "hello_retries": 3,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "mtu": 1100,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "protocol": "udp",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "additional_mxtunnels": {
    "key0": {
      "clusters": [
        {
          "name": "name6",
          "tunterm_hosts": [
            "tunterm_hosts6"
          ],
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "name": "name6",
          "tunterm_hosts": [
            "tunterm_hosts6"
          ],
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "hello_interval": 38,
      "hello_retries": 186,
      "protocol": "ip",
      "vlan_ids": [
        142,
        143
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "clusters": [
        {
          "name": "name6",
          "tunterm_hosts": [
            "tunterm_hosts6"
          ],
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "name": "name6",
          "tunterm_hosts": [
            "tunterm_hosts6"
          ],
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        }
      ],
      "hello_interval": 38,
      "hello_retries": 186,
      "protocol": "ip",
      "vlan_ids": [
        142,
        143
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "ap_subnets": [
    "ap_subnets4",
    "ap_subnets5",
    "ap_subnets6"
  ],
  "auto_preemption": {
    "day_of_week": "mon",
    "enabled": false,
    "time_of_day": "time_of_day4",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "clusters": [
    {
      "name": "name6",
      "tunterm_hosts": [
        "tunterm_hosts6"
      ],
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "created_time": 248.86,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

