
# Flow Capture Request

Flow capture request for one or more switches

## Structure

`FlowCaptureRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DstIp` | `*string` | Optional | - |
| `DstPort` | `*int` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `Duration` | `*int` | Optional | **Default**: `600`<br><br>**Constraints**: `>= 60`, `<= 900` |
| `Protocol` | [`*models.FlowCaptureProtocolEnum`](../../doc/models/flow-capture-protocol-enum.md) | Optional | Flow capture protocol filter. enum: `tcp`, `udp`, `icmp`, `icmp6` |
| `SrcIp` | `*string` | Optional | - |
| `SrcPort` | `*int` | Optional | **Constraints**: `>= 1`, `<= 65535` |
| `Switches` | `[]string` | Required | **Constraints**: *Minimum Items*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    flowCaptureRequest := models.FlowCaptureRequest{
        DstIp:                models.ToPointer("dst_ip0"),
        DstPort:              models.ToPointer(30),
        Duration:             models.ToPointer(600),
        Protocol:             models.ToPointer(models.FlowCaptureProtocolEnum_ICMP),
        SrcIp:                models.ToPointer("src_ip4"),
        Switches:             []string{
            "switches9",
            "switches0",
        },
    }

}
```

