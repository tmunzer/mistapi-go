
# Response Call Troubleshoot Summary

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseCallTroubleshootSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `MeetingId` | `*uuid.UUID` | Optional | - |
| `Results` | [`[]models.CallTroubleshootSummary`](../../doc/models/call-troubleshoot-summary.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "983a78ea4a44",
  "meeting_id": "b784d744-9a7c-4fad-9af0-f78858a319b1",
  "results": [
    {
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
      "timestamp": 2.64,
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
    },
    {
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
      "timestamp": 2.64,
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
    },
    {
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
      "timestamp": 2.64,
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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

