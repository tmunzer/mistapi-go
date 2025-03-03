
# Wlan Bonjour

Bonjour gateway wlan settings

*This model accepts additional fields of type interface{}.*

## Structure

`WlanBonjour`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | `string` | Required | Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses |
| `Enabled` | `*bool` | Optional | Whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false<br>**Default**: `false` |
| `Services` | [`map[string]models.WlanBonjourServiceProperties`](../../doc/models/wlan-bonjour-service-properties.md) | Required | What services are allowed.<br>Property key is the service name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "additional_vlan_ids": "additional_vlan_ids0",
  "enabled": false,
  "services": {
    "airplay": {
      "radius_groups": [
        "teachers"
      ],
      "scope": "same_ap",
      "disable_local": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

