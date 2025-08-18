
# Response Past Spectrum Analysis Fft Sample

FFT sample data for a specific frequency

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePastSpectrumAnalysisFftSample`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Frequency` | `*float64` | Optional | Frequency in MHz |
| `Rssi` | `*float64` | Optional | RSSI in dBm |
| `Signal7` | `*float64` | Optional | RSSI in dBm |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "frequency": 2437.0,
  "rssi": -70.0,
  "signal7": -70.0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

