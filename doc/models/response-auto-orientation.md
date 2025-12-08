
# Response Auto Orientation

## Structure

`ResponseAutoOrientation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Devices` | [`map[string]models.ResponseAutoOrientationDevice`](../../doc/models/response-auto-orientation-device.md) | Optional | Contains the validation status of each device. The Property Key is the device MAC Address. |
| `EstimatedRuntime` | `*int` | Optional | Estimated runtime for the process in seconds |
| `Reason` | `*string` | Optional | Provides the reason for the status. |
| `Started` | `*bool` | Optional | Indicates whether the auto orient process has started. |
| `Valid` | `*bool` | Optional | Indicates whether the auto orient request is valid. |
| `WifiInterrupting` | `*bool` | Optional | Indicates whether the auto orient process will interrupt WiFi traffic. |

## Example (as JSON)

```json
{
  "devices": {
    "key0": {
      "reason": "reason0",
      "valid": false
    },
    "key1": {
      "reason": "reason0",
      "valid": false
    },
    "key2": {
      "reason": "reason0",
      "valid": false
    }
  },
  "estimated_runtime": 142,
  "reason": "reason4",
  "started": false,
  "valid": false
}
```

