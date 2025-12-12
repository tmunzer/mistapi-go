
# Switch Bgp Config Neighbor

## Structure

`SwitchBgpConfigNeighbor`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NeighborAs` | [`models.SwitchBgpConfigNeighborNeighborAs`](../../doc/models/containers/switch-bgp-config-neighbor-neighbor-as.md) | Required | This is a container for any-of cases. |
| `HoldTime` | [`*models.SwitchBgpConfigHoldTime`](../../doc/models/containers/switch-bgp-config-hold-time.md) | Optional | Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535. |
| `ExportPolicy` | `*string` | Optional | Export policy must match one of the policy names defined in the `routing_policies` property. |
| `ImportPolicy` | `*string` | Optional | Import policy must match one of the policy names defined in the `routing_policies` property. |
| `MultihopTtl` | `*int` | Optional | **Constraints**: `>= 1`, `<= 255` |

## Example (as JSON)

```json
{
  "neighbor_as": 65000,
  "hold_time": 0,
  "export_policy": "export_policy0",
  "import_policy": "import_policy4",
  "multihop_ttl": 86
}
```

