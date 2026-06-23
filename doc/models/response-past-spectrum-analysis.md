
# Response Past Spectrum Analysis

Paginated response for past site spectrum analysis records

## Structure

`ResponsePastSpectrumAnalysis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Epoch timestamp, in seconds, for the end of the spectrum analysis query window |
| `Limit` | `*int` | Optional | Maximum number of spectrum analysis records returned in this page |
| `Page` | `*int` | Optional | Current page number returned for the spectrum analysis results |
| `Results` | [`[]models.ResponsePastSpectrumAnalysisResult`](../../doc/models/response-past-spectrum-analysis-result.md) | Optional | Past spectrum analysis records returned for the site |
| `Start` | `*int` | Optional | Epoch timestamp, in seconds, for the start of the spectrum analysis query window |
| `Total` | `*int` | Optional | Number of spectrum analysis records available for the given time range |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePastSpectrumAnalysis := models.ResponsePastSpectrumAnalysis{
        End:                  models.ToPointer(92),
        Limit:                models.ToPointer(178),
        Page:                 models.ToPointer(192),
        Results:              []models.ResponsePastSpectrumAnalysisResult{
            models.ResponsePastSpectrumAnalysisResult{
                Band:                 models.ToPointer("band8"),
                ChannelUsage:         []models.ResponsePastSpectrumAnalysisChannelUsage{
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
                    models.ResponsePastSpectrumAnalysisFftSample{
                        Frequency:            models.ToPointer(float64(91.6)),
                        Rssi:                 models.ToPointer(float64(42.86)),
                        Signal7:              models.ToPointer(float64(18.34)),
                    },
                },
                Mac:                  models.ToPointer("mac0"),
                OrgId:                models.ToPointer(uuid.MustParse("00002492-0000-0000-0000-000000000000")),
            },
        },
        Start:                models.ToPointer(50),
    }

}
```

