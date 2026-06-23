
# Wlan Dynamic Vlan Type Enum

standard (using Tunnel-Private-Group-ID, widely supported), airespace-interface-name (Airespace/Cisco). enum: `airespace-interface-name`, `standard`

## Enumeration

`WlanDynamicVlanTypeEnum`

## Fields

| Name |
|  --- |
| `airespace-interface-name` |
| `standard` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanDynamicVlanType := models.WlanDynamicVlanTypeEnum_AIRESPACEINTERFACENAME

}
```

