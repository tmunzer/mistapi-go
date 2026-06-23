
# Gw Routing Policy

Gateway routing policy made of ordered match-action terms

## Structure

`GwRoutingPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Terms` | [`[]models.GwRoutingPolicyTerm`](../../doc/models/gw-routing-policy-term.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gwRoutingPolicy := models.GwRoutingPolicy{
        Terms:                []models.GwRoutingPolicyTerm{
            models.GwRoutingPolicyTerm{
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
            },
        },
    }

}
```

