
# Switch Bgp Config Neighbor

Per-neighbor switch BGP session settings

## Structure

`SwitchBgpConfigNeighbor`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExportPolicy` | `*string` | Optional | Export policy must match one of the policy names defined in the `routing_policies` property. |
| `HoldTime` | `*int` | Optional | Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535. |
| `ImportPolicy` | `*string` | Optional | Import policy must match one of the policy names defined in the `routing_policies` property. |
| `MultihopTtl` | `*int` | Optional | Time-to-live value for multihop BGP sessions to this neighbor<br><br>**Constraints**: `>= 1`, `<= 255` |
| `NeighborAs` | [`models.BgpAs`](../../doc/models/containers/bgp-as.md) | Required | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchBgpConfigNeighbor := models.SwitchBgpConfigNeighbor{
        ExportPolicy:         models.ToPointer("export_policy4"),
        HoldTime:             models.ToPointer(218),
        ImportPolicy:         models.ToPointer("import_policy8"),
        MultihopTtl:          models.ToPointer(100),
        NeighborAs:           models.BgpAsContainer.FromNumber(65000),
    }

}
```

