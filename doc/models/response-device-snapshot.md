
# Response Device Snapshot

## Structure

`ResponseDeviceSnapshot`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StatusId` | `*string` | Optional | the internal status id |
| `Staus` | [`*models.ResponseDeviceSnapshotStatusEnum`](../../doc/models/response-device-snapshot-status-enum.md) | Optional | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `Timestamp` | `*float64` | Optional | - |

## Example (as JSON)

```json
{
  "status_id": "status_id6",
  "staus": "error",
  "timestamp": 197.34
}
```

