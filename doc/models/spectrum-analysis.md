
# Spectrum Analysis

*This model accepts additional fields of type interface{}.*

## Structure

`SpectrumAnalysis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | [`models.SpectrumAnalysisBandEnum`](../../doc/models/spectrum-analysis-band-enum.md) | Required | Band for spectrum analysis. enum: `24`, `5`, `6` |
| `Channels` | `[]string` | Optional | Optional list of channels to scan. If not specified, all supported channels will be scanned |
| `DeviceId` | `*uuid.UUID` | Optional | Device ID of the AP that is performing spectrum analysis |
| `Duration` | `*int` | Optional | Duration of the spectrum analysis in seconds<br><br>**Default**: `300`<br><br>**Constraints**: `>= 60`, `<= 600` |
| `Format` | [`*models.SpectrumAnalysisFormatEnum`](../../doc/models/spectrum-analysis-format-enum.md) | Optional | Format of the spectrum analysis data. enum: `json`, `stream`<br><br>**Default**: `"json"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "6",
  "channels": [
    "36",
    "40",
    "44",
    "48"
  ],
  "duration": 300,
  "format": "json",
  "device_id": "00001a40-0000-0000-0000-000000000000",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

