
# Bgp Config Type Enum

Required if `via`==`lan`, `via`==`tunnel` or `via`==`wan`. enum: `external`, `internal`

## Enumeration

`BgpConfigTypeEnum`

## Fields

| Name |
|  --- |
| `external` |
| `internal` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    bgpConfigType := models.BgpConfigTypeEnum_EXTERNAL

}
```

