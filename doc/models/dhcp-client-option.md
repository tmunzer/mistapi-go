
# Dhcp Client Option

DHCP client option observed in a DHCP packet

## Structure

`DhcpClientOption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | DHCP option code and option name |
| `Data` | `*string` | Optional | Decoded value carried by the DHCP option |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpClientOption := models.DhcpClientOption{
        Code:                 models.ToPointer("DHO_DHCP_MESSAGE_TYPE(53)"),
        Data:                 models.ToPointer("DHCPREQUEST"),
    }

}
```

