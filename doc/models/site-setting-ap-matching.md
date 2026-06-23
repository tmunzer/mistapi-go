
# Site Setting Ap Matching

Rules for applying AP port configuration by AP model or name

## Structure

`SiteSettingApMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether AP matching rules are enabled |
| `Rules` | [`[]models.SiteSettingApMatchingRule`](../../doc/models/site-setting-ap-matching-rule.md) | Optional | Ordered AP matching rules for applying port configuration |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingApMatching := models.SiteSettingApMatching{
        Enabled:              models.ToPointer(false),
        Rules:                []models.SiteSettingApMatchingRule{
            models.SiteSettingApMatchingRule{
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

