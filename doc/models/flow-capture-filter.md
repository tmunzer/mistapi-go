
# Flow Capture Filter

Normalized filter applied to a flow capture session

## Structure

`FlowCaptureFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DstIp` | `*string` | Optional | - |
| `DstPort` | `*int` | Optional | - |
| `Protocol` | [`*models.FlowCaptureProtocolEnum`](../../doc/models/flow-capture-protocol-enum.md) | Optional | Flow capture protocol filter. enum: `tcp`, `udp`, `icmp`, `icmp6` |
| `SrcIp` | `*string` | Optional | - |
| `SrcPort` | `*int` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    flowCaptureFilter := models.FlowCaptureFilter{
        DstIp:                models.ToPointer("dst_ip2"),
        DstPort:              models.ToPointer(206),
        Protocol:             models.ToPointer(models.FlowCaptureProtocolEnum_ICMP),
        SrcIp:                models.ToPointer("src_ip8"),
        SrcPort:              models.ToPointer(90),
    }

}
```

