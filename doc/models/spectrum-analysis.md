
# Spectrum Analysis

Request parameters for starting AP spectrum analysis at a site

*This model accepts additional fields of type interface{}.*

## Structure

`SpectrumAnalysis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | [`models.SpectrumAnalysisBandEnum`](../../doc/models/spectrum-analysis-band-enum.md) | Required | Band for spectrum analysis. enum: `24`, `5`, `6` |
| `Channels` | `[]int` | Optional | Optional list of channels to scan. If not specified, all supported channels will be scanned |
| `DeviceId` | `*uuid.UUID` | Optional | Device ID of the AP that is performing spectrum analysis |
| `Duration` | `*int` | Optional | Length of the spectrum analysis run, in seconds<br><br>**Default**: `300`<br><br>**Constraints**: `>= 60`, `<= 600` |
| `Format` | [`*models.SpectrumAnalysisFormatEnum`](../../doc/models/spectrum-analysis-format-enum.md) | Optional | Format of the spectrum analysis data. enum: `json`, `stream`<br><br>**Default**: `"json"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    spectrumAnalysis := models.SpectrumAnalysis{
        Band:                 models.SpectrumAnalysisBandEnum_ENUM6,
        Channels:             []int{
            36,
            40,
            44,
            48,
        },
        DeviceId:             models.ToPointer(uuid.MustParse("00000570-0000-0000-0000-000000000000")),
        Duration:             models.ToPointer(300),
        Format:               models.ToPointer(models.SpectrumAnalysisFormatEnum_JSON),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

