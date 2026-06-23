
# Response Past Spectrum Analysis Result

Result of a past spectrum analysis

## Structure

`ResponsePastSpectrumAnalysisResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Radio band on which the spectrum analysis was run (for example, 24, 5, or 6) |
| `ChannelUsage` | [`[]models.ResponsePastSpectrumAnalysisChannelUsage`](../../doc/models/response-past-spectrum-analysis-channel-usage.md) | Optional | Per-channel utilization measurements captured during spectrum analysis |
| `FftSamples` | [`[]models.ResponsePastSpectrumAnalysisFftSample`](../../doc/models/response-past-spectrum-analysis-fft-sample.md) | Optional | List of FFT samples for the spectrum analysis |
| `Mac` | `*string` | Optional | AP MAC address for the access point that ran the spectrum analysis |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Timestamp` | `*int` | Optional | Epoch timestamp, in seconds, when the spectrum analysis was run |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePastSpectrumAnalysisResult := models.ResponsePastSpectrumAnalysisResult{
        Band:                 models.ToPointer("band0"),
        ChannelUsage:         []models.ResponsePastSpectrumAnalysisChannelUsage{
            models.ResponsePastSpectrumAnalysisChannelUsage{
                Channel:              models.ToPointer(192),
                Noise:                models.ToPointer(float64(76.92)),
                NonWifi:              models.ToPointer(float64(164.5)),
                Wifi:                 models.ToPointer(float64(198.3)),
            },
            models.ResponsePastSpectrumAnalysisChannelUsage{
                Channel:              models.ToPointer(192),
                Noise:                models.ToPointer(float64(76.92)),
                NonWifi:              models.ToPointer(float64(164.5)),
                Wifi:                 models.ToPointer(float64(198.3)),
            },
        },
        FftSamples:           []models.ResponsePastSpectrumAnalysisFftSample{
            models.ResponsePastSpectrumAnalysisFftSample{
                Frequency:            models.ToPointer(float64(91.6)),
                Rssi:                 models.ToPointer(float64(42.86)),
                Signal7:              models.ToPointer(float64(18.34)),
            },
            models.ResponsePastSpectrumAnalysisFftSample{
                Frequency:            models.ToPointer(float64(91.6)),
                Rssi:                 models.ToPointer(float64(42.86)),
                Signal7:              models.ToPointer(float64(18.34)),
            },
        },
        Mac:                  models.ToPointer("mac2"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    }

}
```

