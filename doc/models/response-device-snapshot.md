
# Response Device Snapshot

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceSnapshot`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StatusId` | `*string` | Optional | Internal status id |
| `Staus` | [`*models.ResponseDeviceSnapshotStatusEnum`](../../doc/models/response-device-snapshot-status-enum.md) | Optional | enum: `error`, `inprogress`, `scheduled`, `starting`, `success` |
| `Timestamp` | `*float64` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "status_id": "status_id6",
  "staus": "error",
  "timestamp": 197.34,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

