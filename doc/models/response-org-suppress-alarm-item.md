
# Response Org Suppress Alarm Item

## Structure

`ResponseOrgSuppressAlarmItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Duration, in seconds. Maximum duration is 86400 * 14 (14 days). 0 is to un-suppress alarms. |
| `ExpireTime` | `*int` | Optional | - |
| `ScheduledTime` | `*int` | Optional | - |
| `Scope` | [`*models.SuppressedAlarmScopeEnum`](../../doc/models/suppressed-alarm-scope-enum.md) | Optional | level of scope. enum: `org`, `site`<br><br>**Default**: `"site"` |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "scope": "site",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "duration": 158,
  "expire_time": 230,
  "scheduled_time": 8
}
```

