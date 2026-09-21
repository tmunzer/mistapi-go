
# Spectrum Analysis Response

Response returned after starting AP spectrum analysis

## Structure

`SpectrumAnalysisResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | [`*models.SpectrumAnalysisBandEnum`](../../doc/models/spectrum-analysis-band-enum.md) | Optional | Band for spectrum analysis. enum: `24`, `5`, `6` |
| `Channels` | `[]int` | Optional | - |
| `DeviceId` | `*uuid.UUID` | Optional | AP device UUID used for a single-AP spectrum analysis |
| `DeviceIds` | `[]uuid.UUID` | Optional | AP device UUIDs used for a multi-AP spectrum analysis; maximum 5 devices<br><br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `5` |
| `Duration` | `*int` | Optional | - |
| `Format` | [`*models.SpectrumAnalysisFormatEnum`](../../doc/models/spectrum-analysis-format-enum.md) | Optional | Format of the spectrum analysis data. enum: `json`, `stream`<br><br>**Default**: `"json"` |
| `InvalidDeviceIds` | `map[string][]uuid.UUID` | Optional | AP device IDs that failed validation, grouped by reason; available for multi-AP requests |
| `SessionId` | `uuid.UUID` | Required | Spectrum analysis session identifier used to correlate WebSocket output |
| `StartedTime` | `*int` | Optional | Epoch timestamp when spectrum analysis started |
| `Width` | `*int` | Optional | Channel width used during spectrum analysis, in MHz |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    spectrumAnalysisResponse := models.SpectrumAnalysisResponse{
        Band:                 models.ToPointer(models.SpectrumAnalysisBandEnum_ENUM5),
        Channels:             []int{
            95,
            96,
            97,
        },
        DeviceId:             models.ToPointer(uuid.MustParse("00000382-0000-0000-0000-000000000000")),
        DeviceIds:            []uuid.UUID{
            uuid.MustParse("000018a7-0000-0000-0000-000000000000"),
        },
        Duration:             models.ToPointer(92),
        Format:               models.ToPointer(models.SpectrumAnalysisFormatEnum_JSON),
        SessionId:            uuid.MustParse("000000ea-0000-0000-0000-000000000000"),
    }

}
```

