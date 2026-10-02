
# Switch Mgmt Radius Auth Server

RADIUS authentication server used for switch management authentication

## Structure

`SwitchMgmtRadiusAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `string` | Required | Address or hostname of the RADIUS authentication server |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique identifier for this RADIUS authentication server entry |
| `Port` | [`*models.RadiusAuthPort`](../../doc/models/containers/radius-auth-port.md) | Optional | RADIUS Auth Port, value from 1 to 65535, default is 1812 |
| `Secret` | `string` | Required | Shared secret used with this RADIUS authentication server |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMgmtRadiusAuthServer := models.SwitchMgmtRadiusAuthServer{
        Host:                 "1.2.3.4",
        Port:                 models.ToPointer(models.RadiusAuthPortContainer.FromNumber(232)),
        Secret:               "testing123",
    }

}
```

