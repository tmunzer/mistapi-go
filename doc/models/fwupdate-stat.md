
# Fwupdate Stat

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "progress": 10,
  "status": "inprogress",
  "status_id": 5,
  "timestamp": 1716480189.81648,
  "will_retry": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

