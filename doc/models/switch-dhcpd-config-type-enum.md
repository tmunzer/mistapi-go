
# Switch Dhcpd Config Type Enum

enum: `none`, `relay` (DHCP Relay), `server` (DHCP Server)

## Enumeration

`SwitchDhcpdConfigTypeEnum`

## Fields

| Name |
|  --- |
| `none` |
| `relay` |
| `server` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchDhcpdConfigType := models.SwitchDhcpdConfigTypeEnum_SERVER

}
```

