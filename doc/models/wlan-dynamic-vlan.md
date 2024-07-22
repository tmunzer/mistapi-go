
# Wlan Dynamic Vlan

for 802.1x

## Structure

`WlanDynamicVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultVlanId` | [`models.Optional[models.WlanDynamicVlanDefaultVlanId]`](../../doc/models/containers/wlan-dynamic-vlan-default-vlan-id.md) | Optional | This is a container for one-of cases. |
| `Enabled` | `*bool` | Optional | whether to enable dynamic vlan<br>**Default**: `false` |
| `LocalVlanIds` | [`[]models.WlanDynamicVlanLocalVlanIds`](../../doc/models/containers/wlan-dynamic-vlan-local-vlan-ids.md) | Optional | This is Array of a container for one-of cases. |
| `Type` | [`*models.WlanDynamicVlanTypeEnum`](../../doc/models/wlan-dynamic-vlan-type-enum.md) | Optional | standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco)<br>**Default**: `"standard"` |
| `Vlans` | `map[string]string` | Optional | map between vlan_id (as string) to airespace interface names (comma-separated) or null for stndard mapping<br><br>* if `dynamic_vlan.type`==`standard`, property key is the Vlan ID and property value is ""<br>* if `dynamic_vlan.type`==`airespace-interface-name`, property key is the Vlan ID and property value is the Airespace Interface Name |

## Example (as JSON)

```json
{
  "default_vlan_id": 999,
  "enabled": false,
  "type": "airespace-interface-name",
  "vlans": {
    "131": "default",
    "322": "fast,video"
  },
  "local_vlan_ids": [
    "String2"
  ]
}
```

