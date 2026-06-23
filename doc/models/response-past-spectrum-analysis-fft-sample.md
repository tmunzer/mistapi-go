
# Response Past Spectrum Analysis Fft Sample

FFT sample data for a specific frequency

## Structure

`ResponsePastSpectrumAnalysisFftSample`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Frequency` | `*float64` | Optional | Sample frequency, in MHz, for this FFT point |
| `Rssi` | `*float64` | Optional | Received signal strength, in dBm, for this FFT sample |
| `Signal7` | `*float64` | Optional | Additional signal level, in dBm, reported for this FFT sample |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responsePastSpectrumAnalysisFftSample := models.ResponsePastSpectrumAnalysisFftSample{
        Frequency:            models.ToPointer(float64(2437)),
        Rssi:                 models.ToPointer(float64(-70)),
        Signal7:              models.ToPointer(float64(-70)),
    }

}
```

