
# Response Auto Map Assignment

## Structure

`ResponseAutoMapAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Devices` | [`map[string]models.ResponseAutoMapAssignmentDevice`](../../doc/models/response-auto-map-assignment-device.md) | Optional | Contains the validation status of each device. The property key is the device MAC address. |
| `EstimatedRuntime` | `*int` | Optional | Estimated runtime for the process in seconds |
| `Reason` | `*string` | Optional | Provides the reason for the status |
| `Started` | `*bool` | Optional | Indicates whether the auto map assignment process has started |
| `Valid` | `*bool` | Optional | Indicates whether the auto map assignment request is valid |

## Example (as JSON)

```json
{
  "devices": {
    "key0": {
      "reason": "reason0",
      "valid": false
    }
  },
  "estimated_runtime": 138,
  "reason": "reason2",
  "started": false,
  "valid": false
}
```

