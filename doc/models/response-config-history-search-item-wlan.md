
# Response Config History Search Item Wlan

WLAN config history detail

## Structure

`ResponseConfigHistorySearchItemWlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | `string` | Required | Configured authentication method for this WLAN |
| `Bands` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Ssid` | `string` | Required | WLAN SSID for this config history detail |
| `VlanIds` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseConfigHistorySearchItemWlan := models.ResponseConfigHistorySearchItemWlan{
        Auth:                 "auth0",
        Bands:                []string{
            "bands4",
            "bands5",
        },
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Ssid:                 "ssid2",
        VlanIds:              []string{
            "vlan_ids0",
            "vlan_ids1",
        },
    }

}
```

