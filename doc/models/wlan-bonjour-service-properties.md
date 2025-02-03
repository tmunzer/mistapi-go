
# Wlan Bonjour Service Properties

*This model accepts additional fields of type interface{}.*

## Structure

`WlanBonjourServiceProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableLocal` | `*bool` | Optional | Whether to prevent wireless clients to discover bonjour devices on the same WLAN<br>**Default**: `false` |
| `RadiusGroups` | `[]string` | Optional | Optional, if the service is further restricted for certain RADIUS groups |
| `Scope` | [`*models.WlanBonjourServicePropertiesScopeEnum`](../../doc/models/wlan-bonjour-service-properties-scope-enum.md) | Optional | how bonjour services should be discovered for the same WLAN. enum: `same_ap`, `same_map`, `same_site`<br>**Default**: `"same_site"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disable_local": false,
  "scope": "same_site",
  "radius_groups": [
    "radius_groups7"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

