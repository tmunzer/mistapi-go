
# Call Troubleshoot Summar

*This model accepts additional fields of type interface{}.*

## Structure

`CallTroubleshootSummar`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AudioIn` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | - |
| `AudioOut` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | - |
| `Timestamp` | `*int` | Optional | - |
| `VideoIn` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | - |
| `VideoOut` | [`*models.CallTroubleshootSummaryData`](../../doc/models/call-troubleshoot-summary-data.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "timestamp": 1695425115,
  "audio_in": {
    "ap_num_clients": 152.32,
    "ap_rtt": 133.36,
    "client_cpu": 164.78,
    "client_n_streams": 206.36,
    "client_radio_band": 43.4,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "audio_out": {
    "ap_num_clients": 71.16,
    "ap_rtt": 52.2,
    "client_cpu": 245.94,
    "client_n_streams": 125.2,
    "client_radio_band": 218.24,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "video_in": {
    "ap_num_clients": 44.64,
    "ap_rtt": 25.68,
    "client_cpu": 16.46,
    "client_n_streams": 98.68,
    "client_radio_band": 191.72,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "video_out": {
    "ap_num_clients": 6.98,
    "ap_rtt": 244.02,
    "client_cpu": 201.88,
    "client_n_streams": 61.02,
    "client_radio_band": 154.06,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

