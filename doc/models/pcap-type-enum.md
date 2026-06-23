
# Pcap Type Enum

enum: `client`, `gateway`, `new_assoc`, `radiotap`, `radiotap,wired`, `wired`, `wireless`

## Enumeration

`PcapTypeEnum`

## Fields

| Name |
|  --- |
| `client` |
| `gateway` |
| `new_assoc` |
| `radiotap` |
| `radiotap,wired` |
| `wired` |
| `wireless` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pcapType := models.PcapTypeEnum_WIRED

}
```

