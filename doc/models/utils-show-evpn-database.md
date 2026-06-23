
# Utils Show Evpn Database

EVPN database lookup request for device command output

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowEvpnDatabase`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Refresh duration in seconds; set only when `interval` is nonzero<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | Refresh interval in seconds for repeated command output<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `Mac` | `*string` | Optional | Client MAC address filter for the EVPN database lookup |
| `PortId` | `*string` | Optional | Interface identifier filter for the EVPN database lookup |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowEvpnDatabase := models.UtilsShowEvpnDatabase{
        Duration:             models.ToPointer(0),
        Interval:             models.ToPointer(0),
        Mac:                  models.ToPointer("f8c1165c6400"),
        PortId:               models.ToPointer("ge-0/0/0.0"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

