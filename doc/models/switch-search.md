
# Switch Search

Switch record returned by device search endpoints

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Clustered` | `*bool` | Optional | Whether the switch is part of a virtual chassis or cluster |
| `EvpnMissingLinks` | `*bool` | Optional | Whether EVPN topology links are missing for this switch |
| `EvpntopoId` | `*string` | Optional | EVPN topology ID associated with this switch |
| `ExtIp` | `*string` | Optional | External IP address observed for switch management traffic |
| `Hostname` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Ip` | `*string` | Optional | Management IP address currently reported for the switch |
| `LastConfigStatus` | `*string` | Optional | Most recent configuration status reported for the switch |
| `LastHostname` | `*string` | Optional | Most recent hostname detected for the switch |
| `LastTroubleCode` | `*string` | Optional | Most recent trouble code reported for the switch |
| `LastTroubleTimestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Mac` | `*string` | Optional | Switch MAC address reported in search results |
| `Managed` | `*bool` | Optional | Whether the switch is managed by Mist. Deprecated in favor of `mist_configured` |
| `MistConfigured` | `*bool` | Optional | whether the device can be configured by Mist or not. This deprecates `managed` (for adopted device) and `disable_auto_config` for claimed device) |
| `Model` | `*string` | Optional | Switch model reported for this search result |
| `NumMembers` | `*int` | Optional | Number of members in the switch virtual chassis, when applicable |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `RadiusStats` | [`map[string]models.DeviceSearchRadiusStat`](../../doc/models/device-search-radius-stat.md) | Optional | Property key is the RADIUS server IP address |
| `Role` | `*string` | Optional | Switch role reported for this search result |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TimeDrifted` | `*bool` | Optional | Whether the switch clock has drifted from the expected time |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required, Constant | Device Type. enum: `switch`<br><br>**Value**: `"switch"` |
| `Uptime` | `*int` | Optional | Device uptime for the switch, in seconds |
| `Version` | `*string` | Optional | Software version currently running on the switch |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    switchSearch := models.SwitchSearch{
        Clustered:            models.ToPointer(false),
        EvpnMissingLinks:     models.ToPointer(false),
        EvpntopoId:           models.ToPointer("evpntopo_id6"),
        ExtIp:                models.ToPointer("ext_ip0"),
        Hostname:             []string{
            "hostname0",
            "hostname1",
        },
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 "switch",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

