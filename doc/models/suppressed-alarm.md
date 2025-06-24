
# Suppressed Alarm

*This model accepts additional fields of type interface{}.*

## Structure

`SuppressedAlarm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Applies` | [`*models.SuppressedAlarmApplies`](../../doc/models/suppressed-alarm-applies.md) | Optional | If `scope`==`site`. Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration` |
| `Duration` | `*float64` | Optional | Duration, in seconds. Maximum duration is 86400 * 180 (180 days). 0 is to un-suppress alarms<br><br>**Default**: `3600`<br><br>**Constraints**: `>= 0`, `<= 15552000` |
| `ScheduledTime` | `*int` | Optional | Epoch_time in seconds, Default as now, accepted time range is from now to now + 7 days |
| `Scope` | [`*models.SuppressedAlarmScopeEnum`](../../doc/models/suppressed-alarm-scope-enum.md) | Optional | level of scope. enum: `org`, `site`<br><br>**Default**: `"site"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "scheduled_time": 8,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

