
# Snmpv 3 Config Notify Type Enum

Delivery mode for this SNMPv3 notification, such as trap or inform. enum: `inform`, `trap`

## Enumeration

`Snmpv3ConfigNotifyTypeEnum`

## Fields

| Name |
|  --- |
| `inform` |
| `trap` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    snmpv3ConfigNotifyType := models.Snmpv3ConfigNotifyTypeEnum_INFORM

}
```

