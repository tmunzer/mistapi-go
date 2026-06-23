
# Sw Routing Policy Term

Switch routing policy term with match criteria and actions

## Structure

`SwRoutingPolicyTerm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Actions` | [`*models.SwRoutingPolicyTermAction`](../../doc/models/sw-routing-policy-term-action.md) | Optional | Actions applied to routes matched by a switch routing policy term |
| `Matching` | [`*models.SwRoutingPolicyTermMatching`](../../doc/models/sw-routing-policy-term-matching.md) | Optional | Route match criteria for a switch routing policy term; all specified criteria must match |
| `Name` | `string` | Required | Display name of the switch routing policy term |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    swRoutingPolicyTerm := models.SwRoutingPolicyTerm{
        Actions:              models.ToPointer(models.SwRoutingPolicyTermAction{
            Accept:               models.ToPointer(false),
            Community:            []string{
                "community4",
                "community5",
            },
            LocalPreference:      models.ToPointer(models.RoutingPolicyLocalPreferenceContainer.FromString("String5")),
            PrependAsPath:        []string{
                "prepend_as_path9",
                "prepend_as_path8",
                "prepend_as_path7",
            },
        }),
        Matching:             models.ToPointer(models.SwRoutingPolicyTermMatching{
            AsPath:               []models.BgpAs{
                models.BgpAsContainer.FromString("String3"),
            },
            Community:            []string{
                "community4",
            },
            Prefix:               []string{
                "prefix5",
                "prefix6",
                "prefix7",
            },
            Protocol:             []models.SwRoutingPolicyTermMatchingProtocolEnum{
                models.SwRoutingPolicyTermMatchingProtocolEnum_BGP,
                models.SwRoutingPolicyTermMatchingProtocolEnum_DIRECT,
            },
        }),
        Name:                 "name0",
    }

}
```

