
# Evpn Topology Switch Role Enum

use `role`==`none` to remove a switch from the topology. enum: `access`, `border`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none`

## Enumeration

`EvpnTopologySwitchRoleEnum`

## Fields

| Name |
|  --- |
| `access` |
| `border` |
| `collapsed-core` |
| `core` |
| `distribution` |
| `esilag-access` |
| `none` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnTopologySwitchRole := models.EvpnTopologySwitchRoleEnum_ESILAGACCESS

}
```

