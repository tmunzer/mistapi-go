
# Insight Rogue Ap

*This model accepts additional fields of type interface{}.*

## Structure

`InsightRogueAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | MAC of the device that had strongest signal strength for ssid/bssid pair |
| `AvgRssi` | `float64` | Required | Average signal strength of ap_mac for ssid/bssid pair |
| `Bssid` | `string` | Required | BSSID of the network detected as threat |
| `Channel` | `string` | Required | Channel over which ap_mac heard ssid/bssid pair |
| `DeltaX` | `*float64` | Optional | X position relative to the reporting AP (`ap_mac`) |
| `DeltaY` | `*float64` | Optional | Y position relative to the reporting AP (`ap_mac`) |
| `NumAps` | `int` | Required | Num of aps that heard the ssid/bssid pair |
| `SeenOnLan` | `*bool` | Optional | Whether the reporting AP see a wireless client (on LAN) connecting to it |
| `Ssid` | `*string` | Optional | SSID of the network detected as threat |
| `TimesHeard` | `*int` | Optional | Represents number of times the pair was heard in the interval. Each count roughly corresponds to a minute. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac4",
  "avg_rssi": 133.82,
  "bssid": "bssid6",
  "channel": "channel2",
  "delta_x": 22.92,
  "delta_y": 101.78,
  "num_aps": 138,
  "seen_on_lan": false,
  "ssid": "ssid0",
  "times_heard": 112,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

