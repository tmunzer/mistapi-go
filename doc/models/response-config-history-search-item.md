
# Response Config History Search Item

Device config history entry

## Structure

`ResponseConfigHistorySearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel24` | `int` | Required | 2.4 GHz channel configured in this history entry |
| `Channel5` | `int` | Required | 5 GHz channel configured in this history entry |
| `RadioMacs` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Radios` | [`[]models.ResponseConfigHistorySearchItemRadio`](../../doc/models/response-config-history-search-item-radio.md) | Optional | Radio config history details<br><br>**Constraints**: *Unique Items Required* |
| `SecpolicyViolated` | `bool` | Required | Whether the device configuration violated a security policy |
| `Ssids` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Ssids24` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Ssids5` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Version` | `string` | Required | Configuration version associated with this history entry |
| `Wlans` | [`[]models.ResponseConfigHistorySearchItemWlan`](../../doc/models/response-config-history-search-item-wlan.md) | Optional | WLAN config history details<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseConfigHistorySearchItem := models.ResponseConfigHistorySearchItem{
        Channel24:            100,
        Channel5:             248,
        RadioMacs:            []string{
            "radio_macs9",
            "radio_macs0",
            "radio_macs1",
        },
        Radios:               []models.ResponseConfigHistorySearchItemRadio{
            models.ResponseConfigHistorySearchItemRadio{
                Band:                 "band2",
                Channel:              156,
            },
            models.ResponseConfigHistorySearchItemRadio{
                Band:                 "band2",
                Channel:              156,
            },
        },
        SecpolicyViolated:    false,
        Ssids:                []string{
            "ssids3",
            "ssids2",
        },
        Ssids24:              []string{
            "ssids_248",
        },
        Ssids5:               []string{
            "ssids_52",
            "ssids_53",
        },
        Timestamp:            float64(5.6),
        Version:              "version8",
    }

}
```

