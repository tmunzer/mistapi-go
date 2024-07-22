
# Wlan Bonjour

bonjour gateway wlan settings

## Structure

`WlanBonjour`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | [`[]models.WlanBonjourAdditionalVlanIds`](../../doc/models/containers/wlan-bonjour-additional-vlan-ids.md) | Required | This is Array of a container for one-of cases. |
| `Enabled` | `*bool` | Optional | whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false<br>**Default**: `false` |
| `Services` | [`map[string]models.WlanBonjourServiceProperties`](../../doc/models/wlan-bonjour-service-properties.md) | Required | what services are allowed.<br>Property key is the service name |

## Example (as JSON)

```json
{
  "additional_vlan_ids": [
    "String2",
    "String3",
    "String4"
  ],
  "enabled": false,
  "services": {
    "airplay": {
      "radius_groups": [
        "teachers"
      ],
      "scope": "same_ap",
      "disable_local": false
    }
  }
}
```

