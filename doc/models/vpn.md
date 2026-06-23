
# Vpn

Organization VPN overlay configuration

*This model accepts additional fields of type interface{}.*

## Structure

`Vpn`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the VPN configuration<br><br>**Constraints**: *Minimum Length*: `1` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PathSelection` | [`*models.VpnPathSelection`](../../doc/models/vpn-path-selection.md) | Optional | Only if `type`==`hub_spoke`; path selection behavior for VPN paths |
| `Paths` | [`map[string]models.VpnPath`](../../doc/models/vpn-path.md) | Required | For `type`==`hub_spoke`, Property key is the VPN name. For `type`==`mesh`, Property key is the Interface name |
| `Type` | [`*models.VpnModeEnum`](../../doc/models/vpn-mode-enum.md) | Optional | VPN topology mode for this configuration. enum: `hub_spoke`, `mesh`<br><br>**Default**: `"hub_spoke"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    vpn := models.Vpn{
        CreatedTime:          models.ToPointer(float64(229)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(105.96)),
        Name:                 "name0",
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PathSelection:        models.ToPointer(models.VpnPathSelection{
            Strategy:             models.ToPointer(models.VpnPathSelectionStrategyEnum_SIMPLE),
        }),
        Paths:                map[string]models.VpnPath{
            "key0": models.VpnPath{
                BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_BROADBAND),
                BfdUseTunnelMode:     models.ToPointer(false),
                Ip:                   models.ToPointer("ip8"),
                PeerPaths:            map[string]models.VpnPathPeerPathsPeer{
                    "key0": models.VpnPathPeerPathsPeer{
                        Preference:           models.ToPointer(144),
                    },
                },
                Pod:                  models.ToPointer(128),
            },
            "key1": models.VpnPath{
                BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_BROADBAND),
                BfdUseTunnelMode:     models.ToPointer(false),
                Ip:                   models.ToPointer("ip8"),
                PeerPaths:            map[string]models.VpnPathPeerPathsPeer{
                    "key0": models.VpnPathPeerPathsPeer{
                        Preference:           models.ToPointer(144),
                    },
                },
                Pod:                  models.ToPointer(128),
            },
            "key2": models.VpnPath{
                BfdProfile:           models.ToPointer(models.VpnPathBfdProfileEnum_BROADBAND),
                BfdUseTunnelMode:     models.ToPointer(false),
                Ip:                   models.ToPointer("ip8"),
                PeerPaths:            map[string]models.VpnPathPeerPathsPeer{
                    "key0": models.VpnPathPeerPathsPeer{
                        Preference:           models.ToPointer(144),
                    },
                },
                Pod:                  models.ToPointer(128),
            },
        },
        Type:                 models.ToPointer(models.VpnModeEnum_HUBSPOKE),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

