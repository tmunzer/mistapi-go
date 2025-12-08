
# Response Past Spectrum Analysis Channel Usage

Channel usage data for a specific channel

## Structure

`ResponsePastSpectrumAnalysisChannelUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `*int` | Optional | Channel number |
| `Noise` | `*float64` | Optional | Noise level in dBm |
| `NonWifi` | `*float64` | Optional | Percentage of channel usage by non-WiFi signals in the range [0, 1] |
| `Wifi` | `*float64` | Optional | Percentage of channel usage by WiFi in the range [0, 1] |

## Example (as JSON)

```json
{
  "channel": 36,
  "noise": -90.0,
  "non_wifi": 0.87,
  "wifi": 0.13
}
```

