
# Switch Bgp Config

## Structure

`SwitchBgpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | - |
| `BfdMinimumInterval` | `*int` | Optional | Minimum interval in milliseconds for BFD hello packets. A neighbor is considered failed when the device stops receiving replies after the specified interval. Value must be between 1 and 255000.<br><br>**Constraints**: `>= 1`, `<= 255000` |
| `ExportPolicy` | `*string` | Optional | Export policy must match one of the policy names defined in the `routing_policies` property. |
| `HoldTime` | `*int` | Optional | Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535. |
| `ImportPolicy` | `*string` | Optional | Import policy must match one of the policy names defined in the `routing_policies` property. |
| `LocalAs` | [`models.BgpAs`](../../doc/models/containers/bgp-as.md) | Required | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Neighbors` | [`map[string]models.SwitchBgpConfigNeighbor`](../../doc/models/switch-bgp-config-neighbor.md) | Optional | Property key is the BGP Neighbor IP Address. |
| `Networks` | `[]string` | Optional | List of network names for BGP configuration. When a network is specified, a BGP group will be added to the VRF that network is part of. |
| `Type` | [`models.SwitchBgpConfigTypeEnum`](../../doc/models/switch-bgp-config-type-enum.md) | Required | enum: `external`, `internal` |

## Example (as JSON)

```json
{
  "local_as": 65000,
  "type": "external",
  "auth_key": "auth_key2",
  "bfd_minimum_interval": 174,
  "export_policy": "export_policy4",
  "hold_time": 38,
  "import_policy": "import_policy8"
}
```

