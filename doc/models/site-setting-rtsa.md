
# Site Setting Rtsa

managed mobility

## Structure

`SiteSettingRtsa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppWaking` | `*bool` | Optional | **Default**: `false` |
| `DisableDeadReckoning` | `*bool` | Optional | - |
| `DisablePressureSensor` | `*bool` | Optional | **Default**: `false` |
| `Enabled` | `*bool` | Optional | - |
| `TrackAsset` | `*bool` | Optional | asset tracking related<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "app_waking": false,
  "disable_pressure_sensor": false,
  "track_asset": false,
  "disable_dead_reckoning": false,
  "enabled": false
}
```

