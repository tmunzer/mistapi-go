
# Stats Mxedge Port Stat Sfp

SFP transceiver details reported for a Mist Edge port

## Structure

`StatsMxedgePortStatSfp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Codes` | `*string` | Optional | Transceiver capability codes reported by the SFP module |
| `Mbps` | `*int` | Optional | Nominal transceiver speed, in Mbps |
| `PartNo` | `*string` | Optional | Manufacturer part number reported by the SFP module |
| `SerialNo` | `*string` | Optional | Manufacturer serial number reported by the SFP module |
| `Type` | `*int` | Optional | Transceiver type code reported by the SFP module |
| `Vendor` | `*string` | Optional | Manufacturer name reported by the SFP module |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgePortStatSfp := models.StatsMxedgePortStatSfp{
        Codes:                models.ToPointer("codes2"),
        Mbps:                 models.ToPointer(158),
        PartNo:               models.ToPointer("part_no6"),
        SerialNo:             models.ToPointer("serial_no6"),
        Type:                 models.ToPointer(14),
    }

}
```

