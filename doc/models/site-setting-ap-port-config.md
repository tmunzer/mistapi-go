
# Site Setting Ap Port Config

AP Ethernet port configuration overrides by model

## Structure

`SiteSettingApPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ModelSpecific` | [`map[string]models.ApPortConfig`](../../doc/models/ap-port-config.md) | Optional | Property key is the AP model (e.g. "AP32") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingApPortConfig := models.SiteSettingApPortConfig{
        ModelSpecific:        map[string]models.ApPortConfig{
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
                Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_WXTUNNEL),
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
                Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_WXTUNNEL),
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
                Forwarding:           models.ToPointer(models.ApPortConfigForwardingEnum_WXTUNNEL),
                MacAuthPreferred:     models.ToPointer(false),
            },
        },
    }

}
```

