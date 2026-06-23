
# Vrrp Group

Junos VRRP group authentication and network settings

## Structure

`VrrpGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthKey` | `*string` | Optional | If `auth_type`==`md5`, authentication key used by the VRRP group |
| `AuthPassword` | `*string` | Optional | If `auth_type`==`simple`, password used by the VRRP group |
| `AuthType` | [`*models.VrrpGroupAuthTypeEnum`](../../doc/models/vrrp-group-auth-type-enum.md) | Optional | Authentication method used by the VRRP group. enum: `md5`, `simple`<br><br>**Default**: `"md5"` |
| `Networks` | [`map[string]models.VrrpGroupNetwork`](../../doc/models/vrrp-group-network.md) | Optional | Property key is the network name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vrrpGroup := models.VrrpGroup{
        AuthKey:              models.ToPointer("auth-key-1"),
        AuthPassword:         models.ToPointer("auth_password8"),
        AuthType:             models.ToPointer(models.VrrpGroupAuthTypeEnum_MD5),
        Networks:             map[string]models.VrrpGroupNetwork{
            "data": models.VrrpGroupNetwork{
                Ip:                   models.ToPointer("10.182.96.1"),
            },
            "mgmt": models.VrrpGroupNetwork{
                Ip:                   models.ToPointer("10.182.104.1"),
            },
            "v10": models.VrrpGroupNetwork{
                Ip:                   models.ToPointer("10.182.104.129"),
            },
            "wap": models.VrrpGroupNetwork{
                Ip:                   models.ToPointer("10.182.102.1"),
            },
        },
    }

}
```

