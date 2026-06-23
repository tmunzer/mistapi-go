
# Ap Zigbee Allow Join Enum

Controls whether new Zigbee devices are allowed to join the network. enum: `always`, `manual`

## Enumeration

`ApZigbeeAllowJoinEnum`

## Fields

| Name |
|  --- |
| `always` |
| `manual` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apZigbeeAllowJoin := models.ApZigbeeAllowJoinEnum_ALWAYS

}
```

