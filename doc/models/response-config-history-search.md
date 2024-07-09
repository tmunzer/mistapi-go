
# Response Config History Search

## Structure

`ResponseConfigHistorySearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ResponseConfigHistorySearchItem`](../../doc/models/response-config-history-search-item.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 74,
  "limit": 96,
  "results": [
    {
      "channel_24": 140,
      "channel_5": 208,
      "radio_macs": [
        "radio_macs3",
        "radio_macs4",
        "radio_macs5"
      ],
      "radios": [
        {
          "band": "band2",
          "channel": 156
        },
        {
          "band": "band2",
          "channel": 156
        },
        {
          "band": "band2",
          "channel": 156
        }
      ],
      "secpolicy_violated": false,
      "ssids": [
        "ssids3",
        "ssids4"
      ],
      "ssids_24": [
        "ssids_248"
      ],
      "ssids_5": [
        "ssids_56",
        "ssids_57"
      ],
      "timestamp": 2.64,
      "version": "version2"
    }
  ],
  "start": 32,
  "total": 190,
  "next": "next8"
}
```
