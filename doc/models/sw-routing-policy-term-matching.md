
# Sw Routing Policy Term Matching

Route match criteria for a switch routing policy term; all specified criteria must match

## Structure

`SwRoutingPolicyTermMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsPath` | [`[]models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Community` | `[]string` | Optional | BGP community values used as routing-policy match criteria |
| `Prefix` | `[]string` | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `Protocol` | [`[]models.SwRoutingPolicyTermMatchingProtocolEnum`](../../doc/models/sw-routing-policy-term-matching-protocol-enum.md) | Optional | Routing protocols that can match a switch routing policy term |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    swRoutingPolicyTermMatching := models.SwRoutingPolicyTermMatching{
        AsPath:               []models.BgpAs{
            models.BgpAsContainer.FromString("String9"),
            models.BgpAsContainer.FromString("String8"),
            models.BgpAsContainer.FromString("String7"),
        },
        Community:            []string{
            "community8",
            "community9",
            "community0",
        },
        Prefix:               []string{
            "prefix9",
            "prefix0",
        },
        Protocol:             []models.SwRoutingPolicyTermMatchingProtocolEnum{
            models.SwRoutingPolicyTermMatchingProtocolEnum_STATIC,
        },
    }

}
```

