
# Ssr Upgrade Multi

## Structure

`SsrUpgradeMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | [`*models.SsrUpgradeChannelEnum`](../../doc/models/ssr-upgrade-channel-enum.md) | Optional | upgrade channel to follow<br>**Default**: `"stable"` |
| `DeviceIds` | `[]uuid.UUID` | Required | list of 128T device IDs to upgrade |
| `RebootAt` | `*int` | Optional | reboot start time in epoch seconds, default is start_time, -1 disables reboot |
| `StartTime` | `*int` | Optional | 128T firmware download start time in epoch seconds, default is now, -1 disables download |
| `Strategy` | [`*models.SsrUpgradeStrategyEnum`](../../doc/models/ssr-upgrade-strategy-enum.md) | Optional | * `big_bang`: upgrade all at once<br>* `serial`: one at a time<br>**Default**: `"big_bang"` |
| `Version` | `*string` | Optional | 128T firmware version to upgrade (e.g. 5.3.0-93)<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "channel": "stable",
  "device_ids": [
    "000023c3-0000-0000-0000-000000000000"
  ],
  "strategy": "big_bang",
  "reboot_at": 122,
  "start_time": 192,
  "version": "version2"
}
```
