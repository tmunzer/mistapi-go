
# Response Past Spectrum Analysis Channel Usage

Channel usage data for a specific channel

## Structure

`ResponsePastSpectrumAnalysisChannelUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `*int` | Optional | Radio channel measured by this channel usage sample |
| `Noise` | `*float64` | Optional | Measured noise floor, in dBm, for the channel |
| `NonWifi` | `*float64` | Optional | Percentage of channel usage by non-WiFi signals in the range [0, 1] |
| `Wifi` | `*float64` | Optional | Percentage of channel usage by WiFi in the range [0, 1] |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responsePastSpectrumAnalysisChannelUsage := models.ResponsePastSpectrumAnalysisChannelUsage{
        Channel:              models.ToPointer(36),
        Noise:                models.ToPointer(float64(-90)),
        NonWifi:              models.ToPointer(float64(0.87)),
        Wifi:                 models.ToPointer(float64(0.13)),
    }

}
```

