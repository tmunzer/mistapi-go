
# Switch Radius

RADIUS settings for switch management authentication. `auth_servers`, `auth_servers_retries`, and `auth_servers_timeout` are required when `enabled`==`true` and `use_different_radius`==`true`.

## Structure

`SwitchRadius`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthServers` | [`[]models.SwitchMgmtRadiusAuthServer`](../../doc/models/switch-mgmt-radius-auth-server.md) | Optional | List of RADIUS authentication servers used for switch management authentication<br><br>**Constraints**: *Unique Items Required* |
| `AuthServersRetries` | `*int` | Optional | RADIUS auth session retries. Required when `enabled`==`true` and `use_different_radius`==`true`.<br><br>**Default**: `3` |
| `AuthServersTimeout` | `*int` | Optional | RADIUS auth session timeout, in seconds. Required when `enabled`==`true` and `use_different_radius`==`true`.<br><br>**Default**: `5` |
| `Enabled` | `*bool` | Optional | Whether RADIUS is enabled for switch management authentication |
| `Network` | `*string` | Optional | Source network used for connectivity to the RADIUS servers |
| `UseDifferentRadius` | `*bool` | Optional | Whether to use alternate RADIUS settings instead of the default switch `radius_config` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchRadius := models.SwitchRadius{
        AuthServers:          []models.SwitchMgmtRadiusAuthServer{
            models.SwitchMgmtRadiusAuthServer{
                Host:                 "host0",
                Port:                 models.ToPointer(models.RadiusAuthPortContainer.FromNumber(36)),
                Secret:               "secret4",
            },
        },
        AuthServersRetries:   models.ToPointer(3),
        AuthServersTimeout:   models.ToPointer(5),
        Enabled:              models.ToPointer(false),
        Network:              models.ToPointer("network4"),
    }

}
```

