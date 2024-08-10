
# Response Config History Search Item

## Structure

`ResponseConfigHistorySearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel24` | `int` | Required | - |
| `Channel5` | `int` | Required | - |
| `RadioMacs` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Radios` | [`[]models.ResponseConfigHistorySearchItemRadio`](../../doc/models/response-config-history-search-item-radio.md) | Optional | **Constraints**: *Unique Items Required* |
| `SecpolicyViolated` | `bool` | Required | - |
| `Ssids` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Ssids24` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Ssids5` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Timestamp` | `float64` | Required | - |
| `Version` | `string` | Required | - |
| `Wlans` | [`[]models.ResponseConfigHistorySearchItemWlan`](../../doc/models/response-config-history-search-item-wlan.md) | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "channel_24": 76,
  "channel_5": 168,
  "radio_macs": [
    "radio_macs5",
    "radio_macs6"
  ],
  "radios": [
    {
      "band": "band2",
      "channel": 156
    }
  ],
  "secpolicy_violated": false,
  "ssids": [
    "ssids9",
    "ssids8"
  ],
  "ssids_24": [
    "ssids_244",
    "ssids_243"
  ],
  "ssids_5": [
    "ssids_58"
  ],
  "timestamp": 237.76,
  "version": "version4"
}
```

