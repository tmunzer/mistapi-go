
# Ap Template Matching Rule

Model and name match criteria for an AP template port configuration

## Structure

`ApTemplateMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MatchModel` | `*string` | Optional | Access point model matched by this AP template rule<br><br>**Constraints**: *Minimum Length*: `1` |
| `Name` | `*string` | Optional | Display label for this AP template matching rule<br><br>**Constraints**: *Minimum Length*: `1` |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the interface(s) name (e.g. "eth1,eth2") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apTemplateMatchingRule := models.ApTemplateMatchingRule{
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
            "key2": models.ApPortConfig{
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
    }

}
```

