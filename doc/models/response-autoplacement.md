
# Response Autoplacement

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseAutoplacement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Devices` | [`map[string]models.ResponseAutoplacementDevice`](../../doc/models/response-autoplacement-device.md) | Optional | Property key is the AP MAC Address. Contains the validation status of each device. |
| `EstimatedRuntime` | `*int` | Optional | Estimated runtime for the process in seconds. |
| `Reason` | `*string` | Optional | Provides the reason for the status. |
| `Started` | `*bool` | Optional | Indicates whether the autoplacement process has started. |
| `Valid` | `*bool` | Optional | Indicates whether the autoplacement request is valid. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "devices": {
    "key0": {
      "reason": "reason0",
      "valid": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "estimated_runtime": 188,
  "reason": "reason0",
  "started": false,
  "valid": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

