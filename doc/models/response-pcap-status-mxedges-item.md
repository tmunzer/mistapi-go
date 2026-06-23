
# Response Pcap Status Mxedges Item

Mist Edge interface settings for a packet capture

## Structure

`ResponsePcapStatusMxedgesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | [`map[string]models.CaptureMxedgeMxedgesInterfaces`](../../doc/models/capture-mxedge-mxedges-interfaces.md) | Optional | Dict of interfaces to capture on, property key is the port name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responsePcapStatusMxedgesItem := models.ResponsePcapStatusMxedgesItem{
        Interfaces:           map[string]models.CaptureMxedgeMxedgesInterfaces{
            "key0": models.CaptureMxedgeMxedgesInterfaces{
                TcpdumpExpression:    models.ToPointer("tcpdump_expression4"),
            },
        },
    }

}
```

