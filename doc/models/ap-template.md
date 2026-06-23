
# Ap Template

Access point template configuration

*This model accepts additional fields of type interface{}.*

## Structure

`ApTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMatching` | [`models.ApTemplateMatching`](../../doc/models/ap-template-matching.md) | Required | Rules that select which AP template port configuration applies |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ForSite` | `*bool` | Optional, Read-only | Whether the AP template is scoped to a site rather than the organization |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Wifi` | [`*models.ApTemplateWifi`](../../doc/models/ap-template-wifi.md) | Optional | Wi-Fi behavior settings applied by an AP template |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    apTemplate := models.ApTemplate{
        ApMatching:           models.ApTemplateMatching{
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
        },
        CreatedTime:          models.ToPointer(float64(31.88)),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(47.08)),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

