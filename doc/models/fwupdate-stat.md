
# Fwupdate Stat

*This model accepts additional fields of type interface{}.*

## Structure

`FwupdateStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Progress` | `models.Optional[int]` | Optional | **Constraints**: `>= 0`, `<= 100` |
| `Status` | [`models.Optional[models.FwupdateStatStatusEnum]`](../../doc/models/fwupdate-stat-status-enum.md) | Optional | enum: `inprogress`, `failed`, `upgraded` |
| `StatusId` | `models.Optional[int]` | Optional | - |
| `Timestamp` | `models.Optional[float64]` | Optional | - |
| `WillRetry` | `models.Optional[bool]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "progress": 10,
  "status_id": 5,
  "timestamp": 1716480189.8164835,
  "will_retry": false,
  "status": "failed",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

