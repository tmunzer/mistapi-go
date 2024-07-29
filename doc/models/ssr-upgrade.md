
# Ssr Upgrade

## Structure

`SsrUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | [`*models.SsrUpgradeChannelEnum`](../../doc/models/ssr-upgrade-channel-enum.md) | Optional | upgrade channel to follow. enum: `alpha`, `beta`, `stable`<br>**Default**: `"stable"` |
| `RebootAt` | `*int` | Optional | eboot start time in epoch seconds, default is start_time, -1 disables reboot |
| `StartTime` | `*int` | Optional | 128T firmware download start time in epoch seconds, default is now, -1 disables download |
| `Version` | `string` | Required | 128T firmware version to upgrade (e.g. 5.3.0-93)<br>**Default**: `"stable"`<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "channel": "stable",
  "version": "stable",
  "reboot_at": 50,
  "start_time": 236
}
```

