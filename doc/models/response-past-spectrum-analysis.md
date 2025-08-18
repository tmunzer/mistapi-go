
# Response Past Spectrum Analysis

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePastSpectrumAnalysis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | End time of the spectrum analysis in epoch seconds |
| `Limit` | `*int` | Optional | Limit of the number of results returned |
| `Page` | `*int` | Optional | Page number of the results returned |
| `Results` | [`[]models.ResponsePastSpectrumAnalysisResult`](../../doc/models/response-past-spectrum-analysis-result.md) | Optional | - |
| `Start` | `*int` | Optional | Start time of the spectrum analysis in epoch seconds |
| `Total` | `*int` | Optional | Total number of results available for the given time range |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 216,
  "limit": 210,
  "page": 68,
  "results": [
    {
      "band": "band8",
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
        },
        {
          "frequency": 91.6,
          "rssi": 42.86,
          "signal7": 18.34,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
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
      "mac": "mac0",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "band": "band8",
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
        },
        {
          "frequency": 91.6,
          "rssi": 42.86,
          "signal7": 18.34,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
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
      "mac": "mac0",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "band": "band8",
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
        },
        {
          "frequency": 91.6,
          "rssi": 42.86,
          "signal7": 18.34,
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
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
      "mac": "mac0",
      "org_id": "00002492-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 174,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

