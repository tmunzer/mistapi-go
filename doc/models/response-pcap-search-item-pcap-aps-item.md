
# Response Pcap Search Item Pcap Aps Item

AP radio settings captured for a packet capture record

## Structure

`ResponsePcapSearchItemPcapApsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | Radio band used for this AP capture |
| `Bandwidth` | `*string` | Optional | Channel bandwidth used for this AP capture, in MHz |
| `Channel` | `*int` | Optional | Radio channel used for this AP capture |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | Tcpdump filter expression applied to this AP capture, or null when no filter is applied |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responsePcapSearchItemPcapApsItem := models.ResponsePcapSearchItemPcapApsItem{
        Band:                 models.ToPointer("band6"),
        Bandwidth:            models.ToPointer("bandwidth8"),
        Channel:              models.ToPointer(190),
        TcpdumpExpression:    models.NewOptional(models.ToPointer("tcpdump_expression4")),
    }

}
```

