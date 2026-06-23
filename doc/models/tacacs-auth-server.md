
# Tacacs Auth Server

TACACS+ authentication server settings

## Structure

`TacacsAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Address or hostname of the TACACS+ authentication server |
| `Port` | `*string` | Optional | TCP port used by the TACACS+ authentication server |
| `Secret` | `*string` | Optional | Shared secret used with this TACACS+ authentication server |
| `Timeout` | `*int` | Optional | TACACS+ authentication server timeout, in seconds<br><br>**Default**: `10` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tacacsAuthServer := models.TacacsAuthServer{
        Host:                 models.ToPointer("host6"),
        Port:                 models.ToPointer("port2"),
        Secret:               models.ToPointer("secret8"),
        Timeout:              models.ToPointer(10),
    }

}
```

