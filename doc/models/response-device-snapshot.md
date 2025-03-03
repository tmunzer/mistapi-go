
# Response Device Snapshot

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceSnapshot`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Status` | [`*models.ResponseDeviceSnapshotStatusEnum`](../../doc/models/response-device-snapshot-status-enum.md) | Optional | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `StatusId` | `*string` | Optional | Internal status id |
| `Timestamp` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "status": "starting",
  "status_id": "status_id6",
  "timestamp": 197.34,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

