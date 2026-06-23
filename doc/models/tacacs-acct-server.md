
# Tacacs Acct Server

TACACS+ accounting server settings

## Structure

`TacacsAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Address or hostname of the TACACS+ accounting server |
| `Port` | `*string` | Optional | TCP port used by the TACACS+ accounting server |
| `Secret` | `*string` | Optional | Shared secret used with this TACACS+ accounting server |
| `Timeout` | `*int` | Optional | TACACS+ accounting server timeout, in seconds<br><br>**Default**: `10` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tacacsAcctServer := models.TacacsAcctServer{
        Host:                 models.ToPointer("host4"),
        Port:                 models.ToPointer("port2"),
        Secret:               models.ToPointer("secret8"),
        Timeout:              models.ToPointer(10),
    }

}
```

