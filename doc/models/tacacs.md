
# Tacacs

TACACS+ settings for switch management authentication and accounting

## Structure

`Tacacs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.TacacsAcctServer`](../../doc/models/tacacs-acct-server.md) | Optional | List of TACACS+ accounting servers |
| `DefaultRole` | [`*models.TacacsDefaultRoleEnum`](../../doc/models/tacacs-default-role-enum.md) | Optional | enum: `admin`, `helpdesk`, `none`, `read`<br><br>**Default**: `"none"` |
| `Enabled` | `*bool` | Optional | Whether TACACS+ is enabled for switch management authentication |
| `Network` | `*string` | Optional | Source network used for connectivity to the TACACS+ servers |
| `TacplusServers` | [`[]models.TacacsAuthServer`](../../doc/models/tacacs-auth-server.md) | Optional | List of TACACS+ authentication servers |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tacacs := models.Tacacs{
        AcctServers:          []models.TacacsAcctServer{
            models.TacacsAcctServer{
                Host:                 models.ToPointer("host4"),
                Port:                 models.ToPointer("port4"),
                Secret:               models.ToPointer("secret0"),
                Timeout:              models.ToPointer(254),
            },
            models.TacacsAcctServer{
                Host:                 models.ToPointer("host4"),
                Port:                 models.ToPointer("port4"),
                Secret:               models.ToPointer("secret0"),
                Timeout:              models.ToPointer(254),
            },
        },
        DefaultRole:          models.ToPointer(models.TacacsDefaultRoleEnum_NONE),
        Enabled:              models.ToPointer(false),
        Network:              models.ToPointer("network6"),
        TacplusServers:       []models.TacacsAuthServer{
            models.TacacsAuthServer{
                Host:                 models.ToPointer("host6"),
                Port:                 models.ToPointer("port2"),
                Secret:               models.ToPointer("secret2"),
                Timeout:              models.ToPointer(18),
            },
        },
    }

}
```

