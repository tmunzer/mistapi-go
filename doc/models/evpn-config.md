
# Evpn Config

EVPN configuration settings applied to a Junos switch

## Structure

`EvpnConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional, Read-only | Whether EVPN configuration is enabled on the switch |
| `Role` | [`*models.EvpnConfigRoleEnum`](../../doc/models/evpn-config-role-enum.md) | Optional, Read-only | enum: `access`, `border`, `collapsed-core`, `core`, `distribution`, `esilag-access`, `none` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnConfig := models.EvpnConfig{
        Enabled:              models.ToPointer(false),
        Role:                 models.ToPointer(models.EvpnConfigRoleEnum_ACCESS),
    }

}
```

