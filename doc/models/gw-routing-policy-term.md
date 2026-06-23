
# Gw Routing Policy Term

Gateway routing policy term with match criteria and actions

## Structure

`GwRoutingPolicyTerm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Actions` | [`*models.GwRoutingPolicyTermAction`](../../doc/models/gw-routing-policy-term-action.md) | Optional | Actions applied to routes matched by a gateway routing policy term |
| `Matching` | [`*models.GwRoutingPolicyTermMatching`](../../doc/models/gw-routing-policy-term-matching.md) | Optional | Route match criteria for a gateway routing policy term; all specified criteria must match |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gwRoutingPolicyTerm := models.GwRoutingPolicyTerm{
        Actions:              models.ToPointer(models.GwRoutingPolicyTermAction{
            Accept:               models.ToPointer(false),
            AddCommunity:         []string{
                "add_community5",
                "add_community6",
                "add_community7",
            },
            AddTargetVrfs:        []string{
                "add_target_vrfs1",
            },
            Community:            []string{
                "community4",
                "community5",
            },
            ExcludeAsPath:        []string{
                "exclude_as_path0",
                "exclude_as_path1",
                "exclude_as_path2",
            },
        }),
        Matching:             models.ToPointer(models.GwRoutingPolicyTermMatching{
            AsPath:               []models.BgpAs{
                models.BgpAsContainer.FromString("String3"),
            },
            Community:            []string{
                "community4",
            },
            Network:              []string{
                "network7",
                "network8",
                "network9",
            },
            Prefix:               []string{
                "prefix5",
                "prefix6",
                "prefix7",
            },
            Protocol:             []models.GwRoutingPolicyTermMatchingProtocolEnum{
                models.GwRoutingPolicyTermMatchingProtocolEnum_AGGREGATE,
                models.GwRoutingPolicyTermMatchingProtocolEnum_BGP,
            },
        }),
    }

}
```

