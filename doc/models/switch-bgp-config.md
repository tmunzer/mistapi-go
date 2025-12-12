
# Switch Bgp Config

## Structure

`SwitchBgpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`models.SwitchBgpConfigTypeEnum`](../../doc/models/switch-bgp-config-type-enum.md) | Required | enum: `external`, `internal` |
| `Networks` | `[]string` | Optional | List of network names for BGP configuration. When a network is specified, a BGP group will be added to the VRF that network is part of. |
| `BfdMinimumInterval` | `*int` | Optional | Minimum interval in milliseconds for BFD hello packets. A neighbor is considered failed when the device stops receiving replies after the specified interval. Value must be between 1 and 255000.<br><br>**Constraints**: `>= 1`, `<= 255000` |
| `LocalAs` | [`models.BgpAs`](../../doc/models/containers/bgp-as.md) | Required | BGP AS, value in range 1-4294967294 |
| `HoldTime` | [`*models.SwitchBgpConfigHoldTime`](../../doc/models/containers/switch-bgp-config-hold-time.md) | Optional | Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535. |
| `AuthKey` | `*string` | Optional | - |
| `Neighbors` | [`map[string]models.SwitchBgpConfigNeighbor`](../../doc/models/switch-bgp-config-neighbor.md) | Optional | Property key is the BGP Neighbor IP Address. |
| `ExportPolicy` | `*string` | Optional | Export policy must match one of the policy names defined in the `routing_policies` property. |
| `ImportPolicy` | `*string` | Optional | Import policy must match one of the policy names defined in the `routing_policies` property. |

## Example (as JSON)

```json
{
  "type": "external",
  "local_as": 65000,
  "networks": [
    "networks6",
    "networks7",
    "networks8"
  ],
  "bfd_minimum_interval": 174,
  "hold_time": 0,
  "auth_key": "auth_key2",
  "neighbors": {
    "key0": {
      "neighbor_as": "String3",
      "hold_time": 0,
      "export_policy": "export_policy0",
      "import_policy": "import_policy4",
      "multihop_ttl": 118
    }
  }
}
```

