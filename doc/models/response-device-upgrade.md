
# Response Device Upgrade

## Structure

`ResponseDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`models.UpgradeInfoStatusEnum`](../../doc/models/upgrade-info-status-enum.md) | Required | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `Timestamp` | `float64` | Required | Epoch (seconds) |

## Example (as JSON)

```json
{
  "status": "success",
  "timestamp": 104.0
}
```

