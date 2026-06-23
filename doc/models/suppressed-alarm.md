
# Suppressed Alarm

Request body for suppressing organization alarms

*This model accepts additional fields of type interface{}.*

## Structure

`SuppressedAlarm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Applies` | [`*models.SuppressedAlarmApplies`](../../doc/models/suppressed-alarm-applies.md) | Optional | If `scope`==`site`. Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration` |
| `Duration` | `*float64` | Optional | Suppression length, in seconds. Maximum duration is 86400 * 180 (180 days). 0 is to un-suppress alarms<br><br>**Default**: `3600`<br><br>**Constraints**: `>= 0`, `<= 15552000` |
| `ScheduledTime` | `*int` | Optional | Time when alarm suppression starts, in epoch seconds. Defaults to now; accepted range is from now to now + 7 days |
| `Scope` | [`*models.SuppressedAlarmScopeEnum`](../../doc/models/suppressed-alarm-scope-enum.md) | Optional | level of scope. enum: `org`, `site`<br><br>**Default**: `"site"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    suppressedAlarm := models.SuppressedAlarm{
        Applies:              models.ToPointer(models.SuppressedAlarmApplies{
            OrgId:                models.ToPointer(uuid.MustParse("00001bdc-0000-0000-0000-000000000000")),
            SiteIds:              []uuid.UUID{
                uuid.MustParse("00001706-0000-0000-0000-000000000000"),
                uuid.MustParse("00001707-0000-0000-0000-000000000000"),
                uuid.MustParse("00001708-0000-0000-0000-000000000000"),
            },
            SitegroupIds:         []uuid.UUID{
                uuid.MustParse("00000634-0000-0000-0000-000000000000"),
            },
        }),
        Duration:             models.ToPointer(float64(3600)),
        ScheduledTime:        models.ToPointer(234),
        Scope:                models.ToPointer(models.SuppressedAlarmScopeEnum_SITE),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

