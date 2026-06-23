
# Response Org Suppress Alarm Item

Suppressed alarm entry for an organization or site alarm scope

## Structure

`ResponseOrgSuppressAlarmItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Number of seconds that alarms remain suppressed. Maximum duration is 86400 * 14 (14 days). Use 0 to remove suppression. |
| `ExpireTime` | `*int` | Optional | Epoch timestamp, in seconds, when the alarm suppression expires |
| `ScheduledTime` | `*int` | Optional | Epoch timestamp, in seconds, when the alarm suppression starts |
| `Scope` | [`*models.SuppressedAlarmScopeEnum`](../../doc/models/suppressed-alarm-scope-enum.md) | Optional | level of scope. enum: `org`, `site`<br><br>**Default**: `"site"` |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseOrgSuppressAlarmItem := models.ResponseOrgSuppressAlarmItem{
        Duration:             models.ToPointer(4),
        ExpireTime:           models.ToPointer(180),
        ScheduledTime:        models.ToPointer(110),
        Scope:                models.ToPointer(models.SuppressedAlarmScopeEnum_SITE),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

