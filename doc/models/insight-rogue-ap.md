
# Insight Rogue Ap

Rogue or neighbor AP observation returned by site insights

## Structure

`InsightRogueAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | Reporting AP MAC address with the strongest signal for the SSID/BSSID pair |
| `AvgRssi` | `float64` | Required | Average RSSI for the SSID/BSSID pair as heard by the reporting AP |
| `Bssid` | `string` | Required | Rogue or neighbor BSSID detected as a threat |
| `Channel` | `string` | Required | Radio channel where the reporting AP heard the SSID/BSSID pair |
| `DeltaX` | `*float64` | Optional | X position relative to the reporting AP (`ap_mac`) |
| `DeltaY` | `*float64` | Optional | Y position relative to the reporting AP (`ap_mac`) |
| `NumAps` | `int` | Required | Number of APs that heard the SSID/BSSID pair |
| `SeenOnLan` | `*bool` | Optional | Whether the reporting AP sees a LAN-side client associated with the BSSID |
| `Ssid` | `*string` | Optional | Wireless SSID detected for the rogue or neighbor BSSID |
| `TimesHeard` | `*int` | Optional | Represents number of times the pair was heard in the interval. Each count roughly corresponds to a minute. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    insightRogueAp := models.InsightRogueAp{
        ApMac:                "ap_mac6",
        AvgRssi:              float64(90.64),
        Bssid:                "bssid8",
        Channel:              "channel0",
        DeltaX:               models.ToPointer(float64(235.74)),
        DeltaY:               models.ToPointer(float64(58.6)),
        NumAps:               104,
        SeenOnLan:            models.ToPointer(false),
        Ssid:                 models.ToPointer("ssid2"),
        TimesHeard:           models.ToPointer(146),
    }

}
```

