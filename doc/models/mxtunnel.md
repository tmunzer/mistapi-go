
# Mxtunnel

MxTunnel

## Structure

`Mxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AnchorMxtunnelIds` | `[]uuid.UUID` | Optional | list of anchor mxtunnels used for forming edge to edge tunnels |
| `AutoPreemption` | [`*models.AutoPreemption`](../../doc/models/auto-preemption.md) | Optional | schedule to preempt ap’s which are not connected to preferred peer |
| `CreatedTime` | `*float64` | Optional | - |
| `ForSite` | `*bool` | Optional | - |
| `HelloInterval` | `models.Optional[int]` | Optional | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by `hello_retries`.<br>**Default**: `60`<br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `models.Optional[int]` | Optional | **Default**: `7`<br>**Constraints**: `>= 2`, `<= 30` |
| `Id` | `*uuid.UUID` | Optional | - |
| `Ipsec` | [`*models.MxtunnelIpsec`](../../doc/models/mxtunnel-ipsec.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `Mtu` | `*int` | Optional | 0 to enable PMTU, 552-1500 to start PMTU with a lower MTU<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 1500` |
| `MxclusterIds` | `[]uuid.UUID` | Optional | list of mxclusters to deploy this tunnel to |
| `Name` | `models.Optional[string]` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Protocol` | [`*models.MxtunnelProtocolEnum`](../../doc/models/mxtunnel-protocol-enum.md) | Optional | **Default**: `"udp"` |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `VlanIds` | `[]int` | Optional | list of vlan_ids that will be used |

## Example (as JSON)

```json
{
  "hello_interval": 60,
  "hello_retries": 7,
  "mtu": 0,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "protocol": "udp",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "anchor_mxtunnel_ids": [
    "00000b0c-0000-0000-0000-000000000000"
  ],
  "auto_preemption": {
    "day_of_week": "tue",
    "enabled": false,
    "time_of_day": "time_of_day4"
  },
  "created_time": 245.94,
  "for_site": false
}
```

