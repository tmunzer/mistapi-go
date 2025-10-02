
# Setting Ssr Auto Upgrade

auto_upgrade device first time it is onboarded

*This model accepts additional fields of type interface{}.*

## Structure

`SettingSsrAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | [`*models.SsrUpgradeChannelEnum`](../../doc/models/ssr-upgrade-channel-enum.md) | Optional | upgrade channel to follow. enum: `alpha`, `beta`, `stable`<br><br>**Default**: `"stable"` |
| `CustomVersions` | `map[string]string` | Optional | Property key is the SSR model (e.g. "SSR130"). |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": "stable",
  "enabled": false,
  "custom_versions": {
    "key0": "custom_versions3"
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

