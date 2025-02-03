
# Response Async License

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAsyncLicense`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Completed` | `[]string` | Optional | - |
| `Detail` | [`*models.ResponseAsyncLicenseDetail`](../../doc/models/response-async-license-detail.md) | Optional | detail claim status per device |
| `Failed` | `*int` | Optional | Current failed number of device |
| `Incompleted` | `[]string` | Optional | Current incompleted lists (macs) |
| `Processed` | `*int` | Optional | Current proceseed number of device |
| `ScheduledAt` | `*int` | Optional | epoch time of aysnc claim scheduled |
| `Status` | [`*models.ResponseAsyncLicenseStatusEnum`](../../doc/models/response-async-license-status-enum.md) | Optional | processing status of async. enum: `prepared`, `ongoing`, `done` |
| `Succeed` | `*int` | Optional | Current succeed number of device |
| `Timestamp` | `*int` | Optional | epoch time of last reporting time |
| `Total` | `*int` | Optional | total number of device included in claim |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "completed": [
    "completed8",
    "completed7",
    "completed6"
  ],
  "detail": {
    "mac": "mac0",
    "status": "status2",
    "timestamp": 22.64,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "failed": 170,
  "incompleted": [
    "incompleted8",
    "incompleted7",
    "incompleted6"
  ],
  "processed": 152,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

