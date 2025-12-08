
# Fwupdate Stat

## Structure

`FwupdateStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Progress` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 100` |
| `Status` | [`models.Optional[models.FwupdateStatStatusEnum]`](../../doc/models/fwupdate-stat-status-enum.md) | Optional | enum: `inprogress`, `failed`, `upgraded`, `success`, `scheduled`, `error` |
| `StatusId` | `models.Optional[int]` | Optional | - |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `WillRetry` | `models.Optional[bool]` | Optional | - |

## Example (as JSON)

```json
{
  "progress": 10,
  "status_id": 5,
  "will_retry": false,
  "status": "scheduled",
  "timestamp": 231.24
}
```

