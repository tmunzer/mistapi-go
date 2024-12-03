
# Response Upgrade Device

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseUpgradeDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`models.UpgradeInfoStatusEnum`](../../doc/models/upgrade-info-status-enum.md) | Required | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `Timestamp` | `float64` | Required | timestamp |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "status": "success",
  "timestamp": 248.5,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

