
# Wlan Bonjour

Bonjour gateway wlan settings

## Structure

`WlanBonjour`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | [`*models.AdditionalVlanIds`](../../doc/models/containers/additional-vlan-ids.md) | Optional | List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses |
| `Enabled` | `*bool` | Optional | Whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false<br><br>**Default**: `false` |
| `Services` | [`map[string]models.WlanBonjourServiceProperties`](../../doc/models/wlan-bonjour-service-properties.md) | Optional | What services are allowed.<br>Property key is the service name |

## Example (as JSON)

```json
{
  "enabled": false,
  "services": {
    "airplay": {
      "radius_groups": [
        "teachers"
      ],
      "scope": "same_ap",
      "disable_local": false
    }
  },
  "additional_vlan_ids": "String1"
}
```

