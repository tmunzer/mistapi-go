
# Evpn Topology Response

EVPN topology metadata returned by EVPN topology APIs

## Structure

`EvpnTopologyResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `EvpnOptions` | [`*models.EvpnOptions`](../../doc/models/evpn-options.md) | Optional | EVPN topology generation options for campus fabric configuration |
| `ForSite` | `*bool` | Optional | Whether the EVPN topology is scoped to a site rather than the organization |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Display name for the EVPN topology |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Overwrite` | `*bool` | Optional | Whether generated EVPN configuration changes are applied automatically |
| `PodNames` | `map[string]string` | Optional | Property key is the pod number |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    evpnTopologyResponse := models.EvpnTopologyResponse{
        CreatedTime:          models.ToPointer(float64(54.02)),
        EvpnOptions:          models.ToPointer(models.EvpnOptions{
            AutoLoopbackSubnet:   models.ToPointer("auto_loopback_subnet4"),
            AutoLoopbackSubnet6:  models.ToPointer("auto_loopback_subnet60"),
            AutoRouterIdSubnet:   models.ToPointer("auto_router_id_subnet8"),
            AutoRouterIdSubnet6:  models.ToPointer("auto_router_id_subnet60"),
            CoreAsBorder:         models.ToPointer(false),
        }),
        ForSite:              models.ToPointer(false),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(24.94)),
        Name:                 models.ToPointer("CC"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

