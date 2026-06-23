
# Site Rogue

Rogue AP detection settings for a site

## Structure

`SiteRogue`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedVlanIds` | `[]int` | Optional | list of VLAN IDs on which rogue APs are ignored<br><br>**Constraints**: `>= 0`, `<= 4096` |
| `Enabled` | `*bool` | Optional | Whether rogue detection is enabled<br><br>**Default**: `false` |
| `HoneypotEnabled` | `*bool` | Optional | Whether honeypot detection is enabled<br><br>**Default**: `false` |
| `MinDuration` | `*int` | Optional | Minimum duration for a bssid to be considered neighbor<br><br>**Default**: `10`<br><br>**Constraints**: `<= 59` |
| `MinRogueDuration` | `*int` | Optional | Minimum duration for a bssid to be considered rogue<br><br>**Default**: `10`<br><br>**Constraints**: `<= 59` |
| `MinRogueRssi` | `*int` | Optional | Minimum RSSI for an AP to be considered rogue<br><br>**Default**: `-80`<br><br>**Constraints**: `>= -85` |
| `MinRssi` | `*int` | Optional | Minimum RSSI for an AP to be considered neighbor (ignoring APs that’s far away)<br><br>**Default**: `-80`<br><br>**Constraints**: `>= -85` |
| `WhitelistedBssids` | `[]string` | Optional | BSSID values or wildcard patterns excluded from rogue detection |
| `WhitelistedSsids` | `[]string` | Optional | SSID names excluded from rogue detection |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteRogue := models.SiteRogue{
        AllowedVlanIds:       []int{
            37,
            38,
        },
        Enabled:              models.ToPointer(false),
        HoneypotEnabled:      models.ToPointer(false),
        MinDuration:          models.ToPointer(10),
        MinRogueDuration:     models.ToPointer(10),
        MinRogueRssi:         models.ToPointer(-80),
        MinRssi:              models.ToPointer(-80),
        WhitelistedBssids:    []string{
            "NeighborSSID",
        },
        WhitelistedSsids:     []string{
            "cc:8e:6f:d4:bf:16",
            "cc-8e-6f-d4-bf-16",
            "cc-73-*",
            "cc:82:*",
        },
    }

}
```

