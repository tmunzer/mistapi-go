
# Ap Radio Mac

Access point MAC address and its related radio MAC addresses

## Structure

`ApRadioMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | Access point MAC address for this radio MAC mapping<br><br>**Constraints**: *Minimum Length*: `1` |
| `RadioMacs` | `[]string` | Required | Radio MAC addresses associated with an access point |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apRadioMac := models.ApRadioMac{
        Mac:                  "5c5b350001a0",
        RadioMacs:            []string{
            "5c5b350001a0",
            "5c5b350001a1",
        },
    }

}
```

