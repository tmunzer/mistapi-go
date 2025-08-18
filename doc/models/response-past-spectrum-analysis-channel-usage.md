
# Response Past Spectrum Analysis Channel Usage

Channel usage data for a specific channel

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePastSpectrumAnalysisChannelUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `*int` | Optional | Channel number |
| `Noise` | `*float64` | Optional | Noise level in dBm |
| `NonWifi` | `*float64` | Optional | Percentage of channel usage by non-WiFi signals in the range [0, 1] |
| `Wifi` | `*float64` | Optional | Percentage of channel usage by WiFi in the range [0, 1] |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": 36,
  "noise": -90.0,
  "non_wifi": 0.87,
  "wifi": 0.13,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

