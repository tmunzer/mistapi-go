
# Vrrp Group Network

Virtual IP assignment for a VRRP network

## Structure

`VrrpGroupNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | Virtual IP address used by the VRRP group on this network |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vrrpGroupNetwork := models.VrrpGroupNetwork{
        Ip:                   models.ToPointer("ip2"),
    }

}
```

