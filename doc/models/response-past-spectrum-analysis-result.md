
# Response Past Spectrum Analysis Result

Result of a past spectrum analysis

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
      "wifi": 198.3
    },
    {
      "channel": 192,
      "noise": 76.92,
      "non_wifi": 164.5,
      "wifi": 198.3
    }
  ],
  "fft_samples": [
    {
      "frequency": 91.6,
      "rssi": 42.86,
      "signal7": 18.34
    }
  ],
  "mac": "mac4"
}
```

