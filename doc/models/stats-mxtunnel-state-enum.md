
# Stats Mxtunnel State Enum

enum: `established`, `established_with_sessions`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`

## Enumeration

`StatsMxtunnelStateEnum`

## Fields

| Name |
|  --- |
| `established` |
| `established_with_sessions` |
| `idle` |
| `wait-ctrl-conn` |
| `wait-ctrl-reply` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxtunnelState := models.StatsMxtunnelStateEnum_ESTABLISHED

}
```

