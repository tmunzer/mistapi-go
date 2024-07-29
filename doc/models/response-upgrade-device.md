
# Response Upgrade Device

## Structure

`ResponseUpgradeDevice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`models.UpgradeInfoStatusEnum`](../../doc/models/upgrade-info-status-enum.md) | Required | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `Timestamp` | `float64` | Required | timestamp |

## Example (as JSON)

```json
{
  "status": "success",
  "timestamp": 190.9
}
```

