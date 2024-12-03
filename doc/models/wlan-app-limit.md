
# Wlan App Limit

bandwidth limiting for apps (applies to up/down)

*This model accepts additional fields of type interface{}.*

## Structure

`WlanAppLimit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apps` | `map[string]int` | Optional | Map from app key to bandwidth in kbps.<br>Property key is the app key, defined in Get Application List |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `WxtagIds` | `map[string]int` | Optional | Map from wxtag_id of Hostname Wxlan Tags to bandwidth in kbps<br>Property key is the wxtag id |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "apps": {
    "dropbox": 300,
    "netflix": 60
  },
  "enabled": false,
  "wxtag_ids": {
    "f99862d9-2726-931f-7559-3dfdf5d070d3": 30
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

