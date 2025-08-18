
# Response Running Spectrum Analysis

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseRunningSpectrumAnalysis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Band on which the spectrum analysis is running (e.g., 24, 5, 6) |
| `DeviceId` | `*uuid.UUID` | Optional | Device ID of the AP that is running spectrum analysis |
| `Duration` | `*int` | Optional | Duration of the spectrum analysis in seconds |
| `Format` | `*string` | Optional | Format of the spectrum analysis data (e.g., json, stream) |
| `StartedTime` | `*int` | Optional | Time when the spectrum analysis was started |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "band6",
  "device_id": "00001770-0000-0000-0000-000000000000",
  "duration": 26,
  "format": "format0",
  "started_time": 52,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

