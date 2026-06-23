
# Gateway Port Mirroring Port Mirror

Gateway port mirroring rule

## Structure

`GatewayPortMirroringPortMirror`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FamilyType` | `*string` | Optional | Packet family used for this port mirroring rule |
| `IngressPortIds` | `[]string` | Optional | Gateway port IDs used as ingress sources for mirrored traffic |
| `OutputPortId` | `*string` | Optional | Destination gateway port ID that receives mirrored traffic |
| `Rate` | `*int` | Optional | Sampling rate applied to mirrored traffic |
| `RunLength` | `*int` | Optional | Number of bytes copied from each mirrored packet<br><br>**Constraints**: `>= 0` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortMirroringPortMirror := models.GatewayPortMirroringPortMirror{
        FamilyType:           models.ToPointer("family_type0"),
        IngressPortIds:       []string{
            "ingress_port_ids2",
        },
        OutputPortId:         models.ToPointer("ge-0/0/5"),
        Rate:                 models.ToPointer(108),
        RunLength:            models.ToPointer(0),
    }

}
```

