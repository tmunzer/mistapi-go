
# Sw Routing Policy

Switch routing policy made of ordered match-action terms

## Structure

`SwRoutingPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Terms` | [`[]models.SwRoutingPolicyTerm`](../../doc/models/sw-routing-policy-term.md) | Optional | Ordered switch routing policy terms; at least one term is required<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    swRoutingPolicy := models.SwRoutingPolicy{
        Terms:                []models.SwRoutingPolicyTerm{
            models.SwRoutingPolicyTerm{
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
                Name:                 "name8",
            },
            models.SwRoutingPolicyTerm{
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
                Name:                 "name8",
            },
            models.SwRoutingPolicyTerm{
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
                Name:                 "name8",
            },
        },
    }

}
```

