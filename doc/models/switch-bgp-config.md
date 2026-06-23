
# Switch Bgp Config

Switch BGP configuration for a routing instance

## Structure

`SwitchBgpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | Authentication key used for BGP neighbor sessions, when configured |
| `BfdMinimumInterval` | `*int` | Optional | Minimum interval in milliseconds for BFD hello packets. A neighbor is considered failed when the device stops receiving replies after the specified interval. Value must be between 1 and 255000.<br><br>**Constraints**: `>= 1`, `<= 255000` |
| `ExportPolicy` | `*string` | Optional | Export policy must match one of the policy names defined in the `routing_policies` property. |
| `HoldTime` | `*int` | Optional | Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535. |
| `ImportPolicy` | `*string` | Optional | Import policy must match one of the policy names defined in the `routing_policies` property. |
| `LocalAs` | [`models.BgpAs`](../../doc/models/containers/bgp-as.md) | Required | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Neighbors` | [`map[string]models.SwitchBgpConfigNeighbor`](../../doc/models/switch-bgp-config-neighbor.md) | Optional | Switch BGP neighbor settings keyed by neighbor IP address |
| `Networks` | `[]string` | Optional | List of network names for BGP configuration. When a network is specified, a BGP group will be added to the VRF that network is part of. |
| `Type` | [`models.SwitchBgpConfigTypeEnum`](../../doc/models/switch-bgp-config-type-enum.md) | Required | BGP session type for this switch BGP configuration. enum: `external`, `internal` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchBgpConfig := models.SwitchBgpConfig{
        AuthKey:              models.ToPointer("auth_key8"),
        BfdMinimumInterval:   models.ToPointer(182),
        ExportPolicy:         models.ToPointer("export_policy0"),
        HoldTime:             models.ToPointer(46),
        ImportPolicy:         models.ToPointer("import_policy4"),
        LocalAs:              models.BgpAsContainer.FromNumber(65000),
        Type:                 models.SwitchBgpConfigTypeEnum_EXTERNAL,
    }

}
```

