
# Rrm Event Type Enum

enum: `interference-ap-co-channel`, `interference-ap-non-wifi`, `neighbor-ap-down`, `neighbor-ap-recovered`, `radar-detected`, `rrm-radar`, `scheduled-site_rrm`, `triggered-site_rrm`

## Enumeration

`RrmEventTypeEnum`

## Fields

| Name |
|  --- |
| `interference-ap-co-channel` |
| `interference-ap-non-wifi` |
| `neighbor-ap-down` |
| `neighbor-ap-recovered` |
| `radar-detected` |
| `rrm-radar` |
| `scheduled-site_rrm` |
| `triggered-site_rrm` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    rrmEventType := models.RrmEventTypeEnum_SCHEDULEDSITERRM

}
```

