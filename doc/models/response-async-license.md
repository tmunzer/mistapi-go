
# Response Async License

## Structure

`ResponseAsyncLicense`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Completed` | `[]string` | Optional | - |
| `Details` | [`[]models.ResponseAsyncLicenseDetail`](../../doc/models/response-async-license-detail.md) | Optional | - |
| `Failed` | `*int` | Optional | Current failed number of device |
| `Incompleted` | `[]string` | Optional | Current incompleted lists (macs) |
| `Processed` | `*int` | Optional | Current processed number of device |
| `ScheduledAt` | `*int` | Optional | epoch time of aysnc claim scheduled |
| `Status` | [`*models.ResponseAsyncLicenseStatusEnum`](../../doc/models/response-async-license-status-enum.md) | Optional | processing status of async. enum: `prepared`, `ongoing`, `done` |
| `Succeed` | `*int` | Optional | Current succeed number of device |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Total` | `*int` | Optional | total number of device included in claim |

## Example (as JSON)

```json
{
  "completed": [
    "completed8",
    "completed7",
    "completed6"
  ],
  "details": [
    {
      "mac": "mac4",
      "status": "status2",
      "timestamp": 235.48
    },
    {
      "mac": "mac4",
      "status": "status2",
      "timestamp": 235.48
    }
  ],
  "failed": 170,
  "incompleted": [
    "incompleted8",
    "incompleted7",
    "incompleted6"
  ],
  "processed": 152
}
```

