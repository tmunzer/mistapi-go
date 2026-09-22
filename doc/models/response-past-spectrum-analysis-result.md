
# Response Past Spectrum Analysis Result

Result of a past spectrum analysis

## Structure

`ResponsePastSpectrumAnalysisResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Radio band on which the spectrum analysis was run (for example, 24, 5, or 6) |
| `ChannelUsage` | [`[]models.ResponsePastSpectrumAnalysisChannelUsage`](../../doc/models/response-past-spectrum-analysis-channel-usage.md) | Optional | Per-channel utilization measurements captured during spectrum analysis |
| `DeviceId` | `*uuid.UUID` | Optional | AP device UUID used for a single-AP spectrum analysis |
| `DeviceIds` | `[]uuid.UUID` | Optional | AP device UUIDs used for a multi-AP spectrum analysis; maximum 5 devices<br><br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `5` |
| `FftSamples` | [`[]models.ResponsePastSpectrumAnalysisFftSample`](../../doc/models/response-past-spectrum-analysis-fft-sample.md) | Optional | List of FFT samples for the spectrum analysis |
| `Mac` | `*string` | Optional | AP MAC address for the access point that ran the spectrum analysis |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SpectrogramUrl` | `*string` | Optional | URL to the generated spectrogram image for the spectrum analysis |
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
        DeviceId:             models.ToPointer(uuid.MustParse("00001d14-0000-0000-0000-000000000000")),
        DeviceIds:            []uuid.UUID{
            uuid.MustParse("00000b29-0000-0000-0000-000000000000"),
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
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    }

}
```

