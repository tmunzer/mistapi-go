
# Tunnel via Enum

If `via`==`tunnel`, specifies which tunnel (primary/secondary) this neighbor is associated with. enum: `primary`, `secondary`

## Enumeration

`TunnelViaEnum`

## Fields

| Name |
|  --- |
| `primary` |
| `secondary` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelVia := models.TunnelViaEnum_PRIMARY

}
```

