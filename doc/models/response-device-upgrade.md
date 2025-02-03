
# Response Device Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`models.UpgradeInfoStatusEnum`](../../doc/models/upgrade-info-status-enum.md) | Required | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `Timestamp` | `float64` | Required | Timestamp |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "status": "success",
  "timestamp": 104.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

