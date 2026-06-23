
# Ospf Area

Property key is the OSPF Area (Area should be a number (0-255) / IP address)

## Structure

`OspfArea`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IncludeLoopback` | `*bool` | Optional | Whether loopback interfaces are included in this OSPF area<br><br>**Default**: `false` |
| `Networks` | [`map[string]models.OspfAreasNetwork`](../../doc/models/ospf-areas-network.md) | Optional | OSPF network settings keyed by network name |
| `Type` | [`*models.OspfAreaTypeEnum`](../../doc/models/ospf-area-type-enum.md) | Optional | OSPF type. enum: `default`, `nssa`, `stub`<br><br>**Default**: `"default"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ospfArea := models.OspfArea{
        IncludeLoopback:      models.ToPointer(false),
        Networks:             map[string]models.OspfAreasNetwork{
            "corp": models.OspfAreasNetwork{
                AuthKeys:               map[string]string{
                    "1": "auth-key-1",
                },
                AuthPassword:           models.ToPointer("auth_password4"),
                AuthType:               models.ToPointer(models.OspfAreaNetworkAuthTypeEnum_MD5),
                BfdMinimumInterval:     models.ToPointer(500),
                DeadInterval:           models.ToPointer(40),
                HelloInterval:          models.ToPointer(10),
                InterfaceType:          models.ToPointer(models.OspfAreaNetworkInterfaceTypeEnum_NBMA),
                Metric:                 models.NewOptional(models.ToPointer(10000)),
            },
            "guest": models.OspfAreasNetwork{
                AuthKeys:               map[string]string{
                    "key0": "auth_keys4",
                },
                AuthPassword:           models.ToPointer("auth_password4"),
                AuthType:               models.ToPointer(models.OspfAreaNetworkAuthTypeEnum_PASSWORD),
                BfdMinimumInterval:     models.ToPointer(94),
                DeadInterval:           models.ToPointer(118),
                Passive:                models.ToPointer(true),
            },
        },
        Type:                 models.ToPointer(models.OspfAreaTypeEnum_ENUMDEFAULT),
    }

}
```

