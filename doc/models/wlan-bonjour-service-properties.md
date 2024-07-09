
# Wlan Bonjour Service Properties

## Structure

`WlanBonjourServiceProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableLocal` | `*bool` | Optional | whether to prevent wireless clients to discover bonjour devices on the same WLAN<br>**Default**: `false` |
| `RadiusGroups` | `[]string` | Optional | optional, if the service is further restricted for certain RADIUS groups |
| `Scope` | [`*models.WlanBonjourServicePropertiesScopeEnum`](../../doc/models/wlan-bonjour-service-properties-scope-enum.md) | Optional | how bonjour services should be discovered for the same WLAN<br>**Default**: `"same_site"` |

## Example (as JSON)

```json
{
  "disable_local": false,
  "scope": "same_site",
  "radius_groups": [
    "radius_groups3",
    "radius_groups4",
    "radius_groups5"
  ]
}
```

