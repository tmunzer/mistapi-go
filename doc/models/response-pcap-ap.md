
# Response Pcap Ap

AP radio settings used for a packet capture

## Structure

`ResponsePcapAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*int` | Optional | Radio band used for the AP packet capture |
| `Bandwidth` | `*int` | Optional | Channel bandwidth used for the AP packet capture, in MHz |
| `Channel` | `*int` | Optional | Radio channel used for the AP packet capture |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | Tcpdump filter expression applied to the AP packet capture, or null when no filter is applied |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responsePcapAp := models.ResponsePcapAp{
        Band:                 models.ToPointer(122),
        Bandwidth:            models.ToPointer(56),
        Channel:              models.ToPointer(172),
        TcpdumpExpression:    models.NewOptional(models.ToPointer("tcpdump_expression8")),
    }

}
```

