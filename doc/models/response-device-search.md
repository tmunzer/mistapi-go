
# Response Device Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseDeviceSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseDeviceSearchResultsItems`](../../doc/models/containers/response-device-search-results-items.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 54,
  "limit": 116,
  "results": [
    {
      "band_24_bandwidth": "band_24_bandwidth2",
      "band_24_channel": 200,
      "band_24_power": 154,
      "band_5_bandwidth": "band_5_bandwidth0",
      "band_5_channel": 132,
      "mxtunnel_status": "mxtunnel_status4",
      "power_constrained": false,
      "power_opmode": "power_opmode0",
      "wlans": [
        {
          "id": "00001c56-0000-0000-0000-000000000000",
          "ssid": "ssid8",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "id": "00001c56-0000-0000-0000-000000000000",
          "ssid": "ssid8",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "id": "00001c56-0000-0000-0000-000000000000",
          "ssid": "ssid8",
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
    },
    {
      "band_24_bandwidth": "band_24_bandwidth2",
      "band_24_channel": 200,
      "band_24_power": 154,
      "band_5_bandwidth": "band_5_bandwidth0",
      "band_5_channel": 132,
      "mxtunnel_status": "mxtunnel_status4",
      "power_constrained": false,
      "power_opmode": "power_opmode0",
      "wlans": [
        {
          "id": "00001c56-0000-0000-0000-000000000000",
          "ssid": "ssid8",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "id": "00001c56-0000-0000-0000-000000000000",
          "ssid": "ssid8",
          "exampleAdditionalProperty": {
            "key1": "val1",
            "key2": "val2"
          }
        },
        {
          "id": "00001c56-0000-0000-0000-000000000000",
          "ssid": "ssid8",
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
  ],
  "start": 12,
  "total": 210,
  "next": "next4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

