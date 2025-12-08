
# Mxedge Tunterm Multicast Config

## Structure

`MxedgeTuntermMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mdns` | [`*models.MxedgeTuntermMulticastMdns`](../../doc/models/mxedge-tunterm-multicast-mdns.md) | Optional | - |
| `Ssdp` | [`*models.MxedgeTuntermMulticastSsdp`](../../doc/models/mxedge-tunterm-multicast-ssdp.md) | Optional | - |

## Example (as JSON)

```json
{
  "mdns": {
    "enabled": false,
    "vlan_ids": [
      "vlan_ids4",
      "vlan_ids5"
    ]
  },
  "ssdp": {
    "enabled": false,
    "vlan_ids": [
      "vlan_ids2",
      "vlan_ids3",
      "vlan_ids4"
    ]
  }
}
```

