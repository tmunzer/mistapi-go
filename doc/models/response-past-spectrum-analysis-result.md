
# Response Past Spectrum Analysis Result

Result of a past spectrum analysis

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePastSpectrumAnalysisResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Band on which the spectrum analysis was run (e.g., 24, 5, 6) |
| `ChannelUsage` | [`[]models.ResponsePastSpectrumAnalysisChannelUsage`](../../doc/models/response-past-spectrum-analysis-channel-usage.md) | Optional | - |
| `FftSamples` | [`[]models.ResponsePastSpectrumAnalysisFftSample`](../../doc/models/response-past-spectrum-analysis-fft-sample.md) | Optional | List of FFT samples for the spectrum analysis |
| `Mac` | `*string` | Optional | MAC Address of the AP that ran the spectrum analysis |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*int` | Optional | Timestamp when the spectrum analysis was run in epoch seconds |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "band": "band2",
  "channel_usage": [
    {
      "channel": 192,
      "noise": 76.92,
      "non_wifi": 164.5,
      "wifi": 198.3,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "channel": 192,
      "noise": 76.92,
      "non_wifi": 164.5,
      "wifi": 198.3,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "fft_samples": [
    {
      "frequency": 91.6,
      "rssi": 42.86,
      "signal7": 18.34,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "mac": "mac4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

