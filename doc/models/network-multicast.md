
# Network Multicast

Whether to enable multicast support (only PIM-sparse mode is supported)

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkMulticast`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableIgmp` | `*bool` | Optional | If the network will only be the source of the multicast traffic, IGMP can be disabled<br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Groups` | [`map[string]models.NetworkMulticastGroup`](../../doc/models/network-multicast-group.md) | Optional | Group address to RP (rendezvous point) mapping. Property Key is the CIDR (example "225.1.0.3/32") |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disable_igmp": false,
  "enabled": false,
  "groups": {
    "key0": {
      "rp_ip": "rp_ip4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

