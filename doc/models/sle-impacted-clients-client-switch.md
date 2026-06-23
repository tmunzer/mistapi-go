
# Sle Impacted Clients Client Switch

Switch association for an impacted client

## Structure

`SleImpactedClientsClientSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `SwitchMac` | `*string` | Optional | MAC address of the switch associated with this impacted client |
| `SwitchName` | `*string` | Optional | Display name of the switch associated with this impacted client |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sleImpactedClientsClientSwitch := models.SleImpactedClientsClientSwitch{
        Interfaces:           []string{
            "interfaces5",
            "interfaces6",
        },
        SwitchMac:            models.ToPointer("switch_mac2"),
        SwitchName:           models.ToPointer("switch_name6"),
    }

}
```

