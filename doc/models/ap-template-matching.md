
# Ap Template Matching

Rules that select which AP template port configuration applies

## Structure

`ApTemplateMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether AP matching is enabled for this AP template |
| `Rules` | [`[]models.ApTemplateMatchingRule`](../../doc/models/ap-template-matching-rule.md) | Optional | Ordered AP template matching rules<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apTemplateMatching := models.ApTemplateMatching{
        Enabled:              models.ToPointer(false),
        Rules:                []models.ApTemplateMatchingRule{
            models.ApTemplateMatchingRule{
                MatchModel:           models.ToPointer("match_model0"),
                Name:                 models.ToPointer("name8"),
                PortConfig:           map[string]models.ApPortConfig{
                    "key0": models.ApPortConfig{
                        Disabled:             models.ToPointer(false),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(34),
                            Enabled:              models.ToPointer(false),
                            Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
                            Vlans:                map[string]string{
                                "key0": "vlans1",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_SITEMXEDGE),
                        MacAuthPreferred:     models.ToPointer(false),
                    },
                    "key1": models.ApPortConfig{
                        Disabled:             models.ToPointer(false),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(34),
                            Enabled:              models.ToPointer(false),
                            Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
                            Vlans:                map[string]string{
                                "key0": "vlans1",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_SITEMXEDGE),
                        MacAuthPreferred:     models.ToPointer(false),
                    },
                },
            },
            models.ApTemplateMatchingRule{
                MatchModel:           models.ToPointer("match_model0"),
                Name:                 models.ToPointer("name8"),
                PortConfig:           map[string]models.ApPortConfig{
                    "key0": models.ApPortConfig{
                        Disabled:             models.ToPointer(false),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(34),
                            Enabled:              models.ToPointer(false),
                            Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
                            Vlans:                map[string]string{
                                "key0": "vlans1",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_SITEMXEDGE),
                        MacAuthPreferred:     models.ToPointer(false),
                    },
                    "key1": models.ApPortConfig{
                        Disabled:             models.ToPointer(false),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(34),
                            Enabled:              models.ToPointer(false),
                            Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
                            Vlans:                map[string]string{
                                "key0": "vlans1",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_SITEMXEDGE),
                        MacAuthPreferred:     models.ToPointer(false),
                    },
                },
            },
            models.ApTemplateMatchingRule{
                MatchModel:           models.ToPointer("match_model0"),
                Name:                 models.ToPointer("name8"),
                PortConfig:           map[string]models.ApPortConfig{
                    "key0": models.ApPortConfig{
                        Disabled:             models.ToPointer(false),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(34),
                            Enabled:              models.ToPointer(false),
                            Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
                            Vlans:                map[string]string{
                                "key0": "vlans1",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_SITEMXEDGE),
                        MacAuthPreferred:     models.ToPointer(false),
                    },
                    "key1": models.ApPortConfig{
                        Disabled:             models.ToPointer(false),
                        DynamicVlan:          models.ToPointer(models.ApPortConfigDynamicVlan{
                            DefaultVlanId:        models.ToPointer(34),
                            Enabled:              models.ToPointer(false),
                            Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
                            Vlans:                map[string]string{
                                "key0": "vlans1",
                            },
                        }),
                        EnableMacAuth:        models.ToPointer(false),
                        Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_SITEMXEDGE),
                        MacAuthPreferred:     models.ToPointer(false),
                    },
                },
            },
        },
    }

}
```

