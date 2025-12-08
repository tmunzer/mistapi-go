
# Network Multicast

Whether to enable multicast support (only PIM-sparse mode is supported)

## Structure

`NetworkMulticast`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableIgmp` | `*bool` | Optional | If the network will only be the source of the multicast traffic, IGMP can be disabled<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Groups` | [`map[string]models.NetworkMulticastGroup`](../../doc/models/network-multicast-group.md) | Optional | Group address to RP (rendezvous point) mapping. Property Key is the CIDR (example "225.1.0.3/32") |

## Example (as JSON)

```json
{
  "disable_igmp": false,
  "enabled": false,
  "groups": {
    "key0": {
      "rp_ip": "rp_ip4"
    }
  }
}
```

