
# Utils Show Route

Route table lookup request for device command output

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Refresh duration in seconds; set only when `interval` is nonzero<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | Refresh interval in seconds for repeated command output<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `Neighbor` | `*string` | Optional | BGP neighbor IP address filter for received or advertised routes |
| `Node` | [`*models.HaClusterNode`](../../doc/models/ha-cluster-node.md) | Optional | HA cluster node selector for device operations |
| `Prefix` | `*string` | Optional | IPv4 or IPv6 prefix filter for route entries |
| `Protocol` | [`*models.UtilsShowRouteProtocolEnum`](../../doc/models/utils-show-route-protocol-enum.md) | Optional | enum: `any`, `bgp`, `direct`, `evpn`, `ospf`, `static`<br><br>**Default**: `"bgp"` |
| `Route` | `*string` | Optional | if neighbor is specified, received / advertised; if not specified, both will be shown<br><br>* for SSR, show bgp neighbors 10.250.18.202 received-routes/advertised-routes<br>* for SRX and Switches, show route receive-protocol/advertise-protocol bgp 192.168.255.12 |
| `Vrf` | `*string` | Optional | Routing instance or VRF filter for route entries |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowRoute := models.UtilsShowRoute{
        Duration:             models.ToPointer(0),
        Interval:             models.ToPointer(0),
        Neighbor:             models.ToPointer("192.168.4.1"),
        Node:                 models.ToPointer(models.HaClusterNode{
            Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
            AdditionalProperties: map[string]interface{}{
                "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
            },
        }),
        Prefix:               models.ToPointer("192.168.0.5/30"),
        Protocol:             models.ToPointer(models.UtilsShowRouteProtocolEnum_BGP),
        Route:                models.ToPointer("advertised"),
        Vrf:                  models.ToPointer("default"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

