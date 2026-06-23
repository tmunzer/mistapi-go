
# Site Setting Ap Matching Rule

AP matching rule for selecting APs and applying port configuration

## Structure

`SiteSettingApMatchingRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MatchModel` | `*string` | Optional | AP model matched by this rule |
| `Name` | `*string` | Optional | Display name of the AP matching rule |
| `PortConfig` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the interface(s) (e.g. "eth1,eth2") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingApMatchingRule := models.SiteSettingApMatchingRule{
        MatchModel:           models.ToPointer("AP12"),
        Name:                 models.ToPointer("AP12"),
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

