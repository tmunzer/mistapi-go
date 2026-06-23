
# Config Vc Port Member

Virtual Chassis member and port list for a VC port operation

## Structure

`ConfigVcPortMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `float64` | Required | Virtual Chassis member ID whose VC ports are being configured |
| `VcPorts` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    configVcPortMember := models.ConfigVcPortMember{
        Member:               float64(163.06),
        VcPorts:              []string{
            "vc_ports4",
            "vc_ports5",
            "vc_ports6",
        },
    }

}
```

