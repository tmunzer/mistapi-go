
# Response Device Search

## Structure

`ResponseDeviceSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseDeviceSearchResults`](../../doc/models/containers/response-device-search-results.md) | Required | This is Array of a container for any-of cases.<br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 238,
  "limit": 188,
  "results": [
    {
      "band_24_bandwith": "band_24_bandwith2",
      "band_24_channel": 0,
      "band_24_power": 98,
      "band_5_bandwith": "band_5_bandwith2",
      "band_5_channel": 188
    },
    {
      "band_24_bandwith": "band_24_bandwith2",
      "band_24_channel": 0,
      "band_24_power": 98,
      "band_5_bandwith": "band_5_bandwith2",
      "band_5_channel": 188
    },
    {
      "band_24_bandwith": "band_24_bandwith2",
      "band_24_channel": 0,
      "band_24_power": 98,
      "band_5_bandwith": "band_5_bandwith2",
      "band_5_channel": 188
    }
  ],
  "start": 196,
  "total": 26,
  "next": "next2"
}
```

