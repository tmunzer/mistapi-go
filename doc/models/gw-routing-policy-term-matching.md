
# Gw Routing Policy Term Matching

Route match criteria for a gateway routing policy term; all specified criteria must match

## Structure

`GwRoutingPolicyTermMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsPath` | [`[]models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Community` | `[]string` | Optional | BGP community values used as routing-policy match criteria |
| `Network` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Prefix` | `[]string` | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `Protocol` | [`[]models.GwRoutingPolicyTermMatchingProtocolEnum`](../../doc/models/gw-routing-policy-term-matching-protocol-enum.md) | Optional | Routing protocols that can match a gateway routing policy term |
| `RouteExists` | [`*models.GwRoutingPolicyTermMatchingRouteExists`](../../doc/models/gw-routing-policy-term-matching-route-exists.md) | Optional | Route-existence match condition for a gateway routing policy term |
| `VpnNeighborMac` | `[]string` | Optional | Overlay neighbor MAC addresses used for bgp_config where `via`==`vpn` |
| `VpnPath` | `[]string` | Optional | Overlay path names used for bgp_config where `via`==`vpn`; order is significant |
| `VpnPathSla` | [`*models.GwRoutingPolicyTermMatchingVpnPathSla`](../../doc/models/gw-routing-policy-term-matching-vpn-path-sla.md) | Optional | SLA thresholds used when matching VPN paths |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gwRoutingPolicyTermMatching := models.GwRoutingPolicyTermMatching{
        AsPath:               []models.BgpAs{
            models.BgpAsContainer.FromString("String3"),
            models.BgpAsContainer.FromString("String2"),
            models.BgpAsContainer.FromString("String1"),
        },
        Community:            []string{
            "community4",
            "community5",
            "community6",
        },
        Network:              []string{
            "network7",
            "network8",
        },
        Prefix:               []string{
            "prefix5",
            "prefix6",
        },
        Protocol:             []models.GwRoutingPolicyTermMatchingProtocolEnum{
            models.GwRoutingPolicyTermMatchingProtocolEnum_AGGREGATE,
        },
    }

}
```

