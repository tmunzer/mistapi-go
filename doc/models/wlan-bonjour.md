
# Wlan Bonjour

bonjour gateway wlan settings

## Structure

`WlanBonjour`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | `[]int` | Required | additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses |
| `Enabled` | `*bool` | Optional | whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false<br>**Default**: `false` |
| `Services` | [`map[string]models.WlanBonjourServiceProperties`](../../doc/models/wlan-bonjour-service-properties.md) | Required | what services are allowed.<br>Property key is the service name |

## Example (as JSON)

```json
{
  "additional_vlan_ids": [
    35,
    36,
    37
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

