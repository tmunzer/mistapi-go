
# Dhcpd Config Type Enum

enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)

## Enumeration

`DhcpdConfigTypeEnum`

## Fields

| Name |
|  --- |
| `local` |
| `none` |
| `relay` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpdConfigType := models.DhcpdConfigTypeEnum_RELAY

}
```

