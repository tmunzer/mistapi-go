
# Wlan Cisco Cwa

Cisco CWA (central web authentication) required RADIUS with COA in order to work. See CWA: https://www.cisco.com/c/en/us/support/docs/security/identity-services-engine/115732-central-web-auth-00.html

## Structure

`WlanCiscoCwa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedHostnames` | `[]string` | Optional | List of hostnames without http(s):// (matched by substring) |
| `AllowedSubnets` | `[]string` | Optional | CIDR subnets allowed for Cisco CWA client access before authorization |
| `BlockedSubnets` | `[]string` | Optional | List of blocked CIDRs |
| `Enabled` | `*bool` | Optional | Whether Cisco CWA is enabled for this WLAN<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanCiscoCwa := models.WlanCiscoCwa{
        AllowedHostnames:     []string{
            "allowed_hostnames6",
        },
        AllowedSubnets:       []string{
            "allowed_subnets2",
            "allowed_subnets3",
        },
        BlockedSubnets:       []string{
            "blocked_subnets6",
            "blocked_subnets7",
        },
        Enabled:              models.ToPointer(false),
    }

}
```

