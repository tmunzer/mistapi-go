
# Suppressed Alarm

## Structure

`SuppressedAlarm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Applies` | [`*models.SuppressedAlarmApplies`](../../doc/models/suppressed-alarm-applies.md) | Optional | if `scope`==`site`<br>Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration` |
| `Duration` | `*float64` | Optional | duration, in seconds. Maximum duration is 86400 * 180 (180 days). 0 is to un-suppress alarms<br>**Default**: `3600`<br>**Constraints**: `>= 0`, `<= 15552000` |
| `ScheduledTime` | `*int` | Optional | poch_time in seconds, Default as now, accepted time range is from now to now + 7 days |
| `Scope` | [`*models.SuppressedAlarmScopeEnum`](../../doc/models/suppressed-alarm-scope-enum.md) | Optional | level of scope<br>**Default**: `"site"` |

## Example (as JSON)

```json
{
  "duration": 3600.0,
  "scope": "site",
  "applies": {
    "org_id": "00001bdc-0000-0000-0000-000000000000",
    "site_ids": [
      "00001706-0000-0000-0000-000000000000",
      "00001707-0000-0000-0000-000000000000",
      "00001708-0000-0000-0000-000000000000"
    ],
    "sitegroup_ids": [
      "00000634-0000-0000-0000-000000000000"
    ]
  },
  "scheduled_time": 32
}
```
