
# Setting Ssr Auto Upgrade

auto_upgrade device first time it is onboarded

## Structure

`SettingSsrAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | [`*models.SsrUpgradeChannelEnum`](../../doc/models/ssr-upgrade-channel-enum.md) | Optional | upgrade channel to follow. enum: `alpha`, `beta`, `stable`<br><br>**Default**: `"stable"` |
| `CustomVersions` | `map[string]string` | Optional | Property key is the SSR model (e.g. "SSR130"). |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Version` | `*string` | Optional | Firmware version to deploy (e.g. 6.3.0-107.r1). Optional, used when custom_versions not specified |

## Example (as JSON)

```json
{
  "channel": "stable",
  "enabled": false,
  "version": "6.3.0-107.r1",
  "custom_versions": {
    "key0": "custom_versions3"
  }
}
```

