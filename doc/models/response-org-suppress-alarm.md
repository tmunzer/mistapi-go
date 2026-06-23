
# Response Org Suppress Alarm

Response containing currently suppressed organization alarm entries

## Structure

`ResponseOrgSuppressAlarm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Results` | [`[]models.ResponseOrgSuppressAlarmItem`](../../doc/models/response-org-suppress-alarm-item.md) | Optional | Suppressed alarm entries returned for the organization |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseOrgSuppressAlarm := models.ResponseOrgSuppressAlarm{
        Results:              []models.ResponseOrgSuppressAlarmItem{
            models.ResponseOrgSuppressAlarmItem{
                Duration:             models.ToPointer(202),
                ExpireTime:           models.ToPointer(238),
                ScheduledTime:        models.ToPointer(52),
                Scope:                models.ToPointer(models.SuppressedAlarmScopeEnum_ORG),
                SiteId:               models.ToPointer(uuid.MustParse("00001420-0000-0000-0000-000000000000")),
            },
            models.ResponseOrgSuppressAlarmItem{
                Duration:             models.ToPointer(202),
                ExpireTime:           models.ToPointer(238),
                ScheduledTime:        models.ToPointer(52),
                Scope:                models.ToPointer(models.SuppressedAlarmScopeEnum_ORG),
                SiteId:               models.ToPointer(uuid.MustParse("00001420-0000-0000-0000-000000000000")),
            },
            models.ResponseOrgSuppressAlarmItem{
                Duration:             models.ToPointer(202),
                ExpireTime:           models.ToPointer(238),
                ScheduledTime:        models.ToPointer(52),
                Scope:                models.ToPointer(models.SuppressedAlarmScopeEnum_ORG),
                SiteId:               models.ToPointer(uuid.MustParse("00001420-0000-0000-0000-000000000000")),
            },
        },
    }

}
```

