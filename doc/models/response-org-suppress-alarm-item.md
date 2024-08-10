
# Response Org Suppress Alarm Item

## Structure

`ResponseOrgSuppressAlarmItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | duration, in seconds. Maximum duration is 86400 * 14 (14 days). 0 is to un-suppress alarms. |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "duration": 158
}
```

