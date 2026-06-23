
# Wlan Hotspot 20

Hotspot 2.0 WLAN settings

## Structure

`WlanHotspot20`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DomainName` | `[]string` | Optional | Domain names advertised to Hotspot 2.0 clients |
| `Enabled` | `*bool` | Optional | Whether to enable hotspot 2.0 config |
| `NaiRealms` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Operators` | [`[]models.WlanHotspot20OperatorsItemEnum`](../../doc/models/wlan-hotspot-20-operators-item-enum.md) | Optional | List of operators to support |
| `Rcoi` | `[]string` | Optional | Roaming Consortium Organization Identifiers advertised for Hotspot 2.0 |
| `VenueName` | `*string` | Optional | Venue name, default is site name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanHotspot20 := models.WlanHotspot20{
        DomainName:           []string{
            "mist.com",
        },
        Enabled:              models.ToPointer(false),
        NaiRealms:            []string{
            "nai_realms1",
            "nai_realms0",
            "nai_realms9",
        },
        Operators:            []models.WlanHotspot20OperatorsItemEnum{
            models.WlanHotspot20OperatorsItemEnum_GOOGLE,
            models.WlanHotspot20OperatorsItemEnum_ATT,
        },
        Rcoi:                 []string{
            "5A03BA0000",
        },
        VenueName:            models.ToPointer("some_name"),
    }

}
```

