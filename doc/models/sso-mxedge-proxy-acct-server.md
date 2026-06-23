
# Sso Mxedge Proxy Acct Server

RADIUS accounting server for the Mist Edge SSO proxy

## Structure

`SsoMxedgeProxyAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | RADIUS accounting server hostname or IP address |
| `Port` | `*int` | Optional | UDP port for RADIUS accounting requests<br><br>**Default**: `1813` |
| `Secret` | `*string` | Optional | Shared secret used with the RADIUS accounting server |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoMxedgeProxyAcctServer := models.SsoMxedgeProxyAcctServer{
        Host:                 models.ToPointer("1.2.3.4"),
        Port:                 models.ToPointer(1813),
        Secret:               models.ToPointer("testing123"),
    }

}
```

