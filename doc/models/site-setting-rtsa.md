
# Site Setting Rtsa

Managed mobility

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingRtsa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppWaking` | `*bool` | Optional | **Default**: `false` |
| `DisableDeadReckoning` | `*bool` | Optional | - |
| `DisablePressureSensor` | `*bool` | Optional | **Default**: `false` |
| `Enabled` | `*bool` | Optional | - |
| `TrackAsset` | `*bool` | Optional | Asset tracking related<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "app_waking": false,
  "disable_pressure_sensor": false,
  "track_asset": false,
  "disable_dead_reckoning": false,
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

