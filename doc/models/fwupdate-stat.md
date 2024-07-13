
# Fwupdate Stat

## Structure

`FwupdateStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Progress` | `models.Optional[int]` | Optional | - |
| `Status` | `models.Optional[string]` | Optional | - |
| `StatusId` | `models.Optional[int]` | Optional | - |
| `Timestamp` | `models.Optional[float64]` | Optional | - |
| `WillRetry` | `models.Optional[bool]` | Optional | - |

## Example (as JSON)

```json
{
  "progress": 10,
  "status": "inprogress",
  "status_id": 5,
  "timestamp": 1716480189.8164835,
  "will_retry": false
}
```

