
# Wlan Dynamic Vlan

For 802.1x

## Structure

`WlanDynamicVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultVlanId` | [`*models.WlanDynamicVlanDefaultVlanIdDeprecated`](../../doc/models/containers/wlan-dynamic-vlan-default-vlan-id-deprecated.md) | Optional | vlan_id to use when there’s no match from RADIUS |
| `DefaultVlanIds` | [`[]models.WlanDynamicVlanDefaultVlanId`](../../doc/models/containers/wlan-dynamic-vlan-default-vlan-id.md) | Optional | VLAN ID, VLAN range or variable to use when there’s no match from RADIUS |
| `Enabled` | `*bool` | Optional | Requires `vlan_enabled`==`true` to be set to `true`. Whether to enable dynamic vlan<br><br>**Default**: `false` |
| `LocalVlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | VLAN_ids to be locally bridged |
| `Type` | [`*models.WlanDynamicVlanTypeEnum`](../../doc/models/wlan-dynamic-vlan-type-enum.md) | Optional | standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco). enum: `airespace-interface-name`, `standard`<br><br>**Default**: `"standard"` |
| `Vlans` | `map[string]string` | Optional | Map between vlan_id (as string) to airespace interface names (comma-separated) or null for standard mapping<br><br>* if `dynamic_vlan.type`==`standard`, property key is the Vlan ID and property value is \"\"<br>* if `dynamic_vlan.type`==`airespace-interface-name`, property key is the Vlan ID and property value is the Airespace Interface Name |

## Example (as JSON)

```json
{
  "enabled": false,
  "type": "airespace-interface-name",
  "vlans": {
    "131": "default",
    "322": "fast,video"
  },
  "default_vlan_id": "String9",
  "default_vlan_ids": [
    "String5"
  ],
  "local_vlan_ids": [
    "String8",
    "String9"
  ]
}
```

