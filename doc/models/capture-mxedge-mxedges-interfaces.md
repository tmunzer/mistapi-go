
# Capture Mxedge Mxedges Interfaces

Property key is the Port name (e.g. "port1", "kni0", "lacp0", "ipsec", "drop", "oobm"), currently limited to specifying one interface per mxedge

## Structure

`CaptureMxedgeMxedgesInterfaces`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression common for wired,radiotap |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    captureMxedgeMxedgesInterfaces := models.CaptureMxedgeMxedgesInterfaces{
        TcpdumpExpression:    models.ToPointer("tcpdump_expression2"),
    }

}
```

