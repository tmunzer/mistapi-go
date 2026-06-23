
# Wlan Dynamic Vlan

Dynamic VLAN assignment settings for 802.1X WLAN authentication

## Structure

`WlanDynamicVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultVlanId` | [`*models.WlanDynamicVlanDefaultVlanIdDeprecated`](../../doc/models/containers/wlan-dynamic-vlan-default-vlan-id-deprecated.md) | Optional | vlan_id to use when there’s no match from RADIUS |
| `DefaultVlanIds` | [`[]models.WlanDynamicVlanDefaultVlanId`](../../doc/models/containers/wlan-dynamic-vlan-default-vlan-id.md) | Optional | VLAN ID, VLAN range or variable to use when there’s no match from RADIUS |
| `Enabled` | `*bool` | Optional | Requires `vlan_enabled`==`true` to be set to `true`. Whether to enable dynamic vlan<br><br>**Default**: `false` |
| `LocalVlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | VLAN ID, either numeric or expressed as a template variable string |
| `Type` | [`*models.WlanDynamicVlanTypeEnum`](../../doc/models/wlan-dynamic-vlan-type-enum.md) | Optional | standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco). enum: `airespace-interface-name`, `standard`<br><br>**Default**: `"standard"` |
| `Vlans` | `map[string]string` | Optional | Map between vlan_id (as string) to airespace interface names (comma-separated) or null for standard mapping<br><br>* if `dynamic_vlan.type`==`standard`, property key is the VLAN ID and property value is \"\"<br>* if `dynamic_vlan.type`==`airespace-interface-name`, property key is the VLAN ID and property value is the Airespace Interface Name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanDynamicVlan := models.WlanDynamicVlan{
        DefaultVlanId:        models.ToPointer(models.WlanDynamicVlanDefaultVlanIdDeprecatedContainer.FromString("String9")),
        DefaultVlanIds:       []models.WlanDynamicVlanDefaultVlanId{
            models.WlanDynamicVlanDefaultVlanIdContainer.FromString("String5"),
            models.WlanDynamicVlanDefaultVlanIdContainer.FromString("String6"),
        },
        Enabled:              models.ToPointer(false),
        LocalVlanIds:         []models.VlanIdWithVariable{
            models.VlanIdWithVariableContainer.FromString("String8"),
            models.VlanIdWithVariableContainer.FromString("String9"),
            models.VlanIdWithVariableContainer.FromString("String0"),
        },
        Type:                 models.ToPointer(models.WlanDynamicVlanTypeEnum_AIRESPACEINTERFACENAME),
        Vlans:                map[string]string{
            "131": "default",
            "322": "fast,video",
        },
    }

}
```

