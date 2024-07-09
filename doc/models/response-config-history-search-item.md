
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
  "channel_24": 160,
  "channel_5": 188,
  "radio_macs": [
    "radio_macs1",
    "radio_macs2"
  ],
  "radios": [
    {
      "band": "band2",
      "channel": 156
    }
  ],
  "secpolicy_violated": false,
  "ssids": [
    "ssids5",
    "ssids4",
    "ssids3"
  ],
  "ssids_24": [
    "ssids_240",
    "ssids_241"
  ],
  "ssids_5": [
    "ssids_54"
  ],
  "timestamp": 189.32,
  "version": "version0"
}
```

