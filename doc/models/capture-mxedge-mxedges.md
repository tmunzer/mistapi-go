
# Capture Mxedge Mxedges

Property key is the Mx Edge ID, currently limited to one mxedge per org capture session

## Structure

`CaptureMxedgeMxedges`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | [`map[string]models.CaptureMxedgeMxedgesInterfaces`](../../doc/models/capture-mxedge-mxedges-interfaces.md) | Optional | Mx Edge interfaces to include in the capture, keyed by interface name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureMxedgeMxedges := models.CaptureMxedgeMxedges{
        Interfaces:           map[string]models.CaptureMxedgeMxedgesInterfaces{
            "key0": models.CaptureMxedgeMxedgesInterfaces{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression4"),
            },
            "key1": models.CaptureMxedgeMxedgesInterfaces{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression4"),
            },
            "key2": models.CaptureMxedgeMxedgesInterfaces{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression4"),
            },
        },
    }

}
```

