
# Mxtunnel

MxTunnel

*This model accepts additional fields of type interface{}.*

## Structure

`Mxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AnchorMxtunnelIds` | `[]uuid.UUID` | Optional | List of anchor mxtunnels used for forming edge to edge tunnels |
| `AutoPreemption` | [`*models.AutoPreemption`](../../doc/models/auto-preemption.md) | Optional | Schedule to preempt ap’s which are not connected to preferred peer |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional | - |
| `HelloInterval` | `models.Optional[int]` | Optional | In seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by `hello_retries`.<br>**Default**: `60`<br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `models.Optional[int]` | Optional | **Default**: `7`<br>**Constraints**: `>= 2`, `<= 30` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Ipsec` | [`*models.MxtunnelIpsec`](../../doc/models/mxtunnel-ipsec.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Mtu` | `*int` | Optional | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 1500` |
| `MxclusterIds` | `[]uuid.UUID` | Optional | List of mxclusters to deploy this tunnel to |
| `Name` | `models.Optional[string]` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Protocol` | [`*models.MxtunnelProtocolEnum`](../../doc/models/mxtunnel-protocol-enum.md) | Optional | enum: `ip`, `udp`<br>**Default**: `"udp"` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `VlanIds` | `[]int` | Optional | List of vlan_ids that will be used |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "hello_interval": 60,
  "hello_retries": 7,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "mtu": 0,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "protocol": "udp",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "anchor_mxtunnel_ids": [
    "00000338-0000-0000-0000-000000000000",
    "00000339-0000-0000-0000-000000000000"
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
  "created_time": 69.9,
  "for_site": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

