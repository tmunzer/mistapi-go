
# Client Info Source Enum

source from where the client was learned (lldp, mac). enum: `lldp`, `mac`

## Enumeration

`ClientInfoSourceEnum`

## Fields

| Name |
|  --- |
| `lldp` |
| `mac` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    clientInfoSource := models.ClientInfoSourceEnum_LLDP

}
```

