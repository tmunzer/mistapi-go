
# Gateway Port Mirroring

Port mirroring settings for a gateway interface

## Structure

`GatewayPortMirroring`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortMirror` | [`*models.GatewayPortMirroringPortMirror`](../../doc/models/gateway-port-mirroring-port-mirror.md) | Optional | Gateway port mirroring rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortMirroring := models.GatewayPortMirroring{
        PortMirror:           models.ToPointer(models.GatewayPortMirroringPortMirror{
            FamilyType:           models.ToPointer("family_type0"),
            IngressPortIds:       []string{
                "ingress_port_ids8",
                "ingress_port_ids7",
                "ingress_port_ids6",
            },
            OutputPortId:         models.ToPointer("output_port_id2"),
            Rate:                 models.ToPointer(150),
            RunLength:            models.ToPointer(214),
        }),
    }

}
```

