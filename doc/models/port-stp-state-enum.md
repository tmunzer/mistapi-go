
# Port Stp State Enum

Spanning Tree Protocol state for the port. enum: `""`, `blocking`, `disabled`, `forwarding`, `learning`, `listening`

## Enumeration

`PortStpStateEnum`

## Fields

| Name |
|  --- |
| `blocking` |
| `disabled` |
| `forwarding` |
| `learning` |
| `listening` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    portStpState := models.PortStpStateEnum_FORWARDING

}
```

