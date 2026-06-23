
# Mxcluster Radsec Acct Server

RadSec accounting server settings for a Mist Edge cluster

## Structure

`MxclusterRadsecAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | IP / hostname of RADIUS server |
| `Port` | `*int` | Optional | Acct port of RADIUS server<br><br>**Default**: `1813` |
| `Secret` | `*string` | Optional | Shared secret used with this RADIUS accounting server |
| `Ssids` | `[]string` | Optional | List of ssids that will use this server if match_ssid is true and match is found |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxclusterRadsecAcctServer := models.MxclusterRadsecAcctServer{
        Host:                 models.ToPointer("host4"),
        Port:                 models.ToPointer(1813),
        Secret:               models.ToPointer("secret0"),
        Ssids:                []string{
            "ssids5",
            "ssids4",
        },
    }

}
```

