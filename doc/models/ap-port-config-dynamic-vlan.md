
# Ap Port Config Dynamic Vlan

Dynamic VLAN assignment settings for AP port authentication

## Structure

`ApPortConfigDynamicVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultVlanId` | `*int` | Optional | Fallback VLAN ID used when RADIUS does not return a dynamic VLAN match<br><br>**Constraints**: `>= 1`, `<= 4094` |
| `Enabled` | `*bool` | Optional | Whether dynamic VLAN assignment is enabled for this AP port |
| `Type` | [`*models.ApPortConfigDynamicVlanTypeEnum`](../../doc/models/ap-port-config-dynamic-vlan-type-enum.md) | Optional | Mapping mode for interpreting dynamic VLAN attributes returned by RADIUS.\ \ enum: `airespace-interface-name`, where the VLAN is determined by parsing\ \ the RADIUS attribute as an Airespace interface name (e.g. "guest"\ \ would map to VLAN 100), or `standard`, where the VLAN is determined by parsing\ \ the RADIUS attribute as a standard VLAN ID number |
| `Vlans` | `map[string]string` | Optional | Mapping entries for RADIUS-assigned VLAN values on this AP port. For `type`==`airespace-interface-name`, the property key is the Airespace interface name returned by RADIUS (e.g. "guest"), and the value is the corresponding VLAN ID (e.g. 100). For `type`==`standard`, the property key is the VLAN ID number returned by RADIUS, and the value is ignored. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apPortConfigDynamicVlan := models.ApPortConfigDynamicVlan{
        DefaultVlanId:        models.ToPointer(999),
        Enabled:              models.ToPointer(false),
        Type:                 models.ToPointer(models.ApPortConfigDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
        Vlans:                map[string]string{
            "1-10": "",
            "user": "vlans8",
        },
    }

}
```

