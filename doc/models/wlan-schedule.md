
# Wlan Schedule

WLAN operating schedule, default is disabled

## Structure

`WlanSchedule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | hours of operation filter, the available days (mon, tue, wed, thu, fri, sat, sun).<br><br>**Note**: If the dow is not defined then it’s treated as 00:00-23:59. |

## Example (as JSON)

```json
{
  "enabled": false,
  "hours": {
    "fri": "fri2",
    "mon": "mon8",
    "sat": "sat0",
    "sun": "sun6",
    "thu": "thu6"
  }
}
```

