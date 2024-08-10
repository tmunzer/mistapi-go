
# Response Device Search

## Structure

`ResponseDeviceSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseDeviceSearchResultsItems`](../../doc/models/containers/response-device-search-results-items.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 54,
  "limit": 116,
  "results": [
    {
      "band_24_bandwith": "band_24_bandwith8",
      "band_24_channel": 200,
      "band_24_power": 154,
      "band_5_bandwith": "band_5_bandwith6",
      "band_5_channel": 132
    },
    {
      "band_24_bandwith": "band_24_bandwith8",
      "band_24_channel": 200,
      "band_24_power": 154,
      "band_5_bandwith": "band_5_bandwith6",
      "band_5_channel": 132
    }
  ],
  "start": 12,
  "total": 210,
  "next": "next4"
}
```

