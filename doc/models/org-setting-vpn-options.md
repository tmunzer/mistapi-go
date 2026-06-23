
# Org Setting Vpn Options

Organization VPN behavior options

## Structure

`OrgSettingVpnOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsBase` | `*int` | Optional | Base BGP autonomous system number used for generated VPN configurations<br><br>**Constraints**: `>= 1`, `<= 2147483647` |
| `EnableIpv6` | `*bool` | Optional | Whether IPv6 is enabled for organization VPN configuration<br><br>**Default**: `false` |
| `StSubnet` | `*string` | Optional | requiring /12 or bigger to support 16 private IPs for 65535 gateways<br><br>**Default**: `"10.224.0.0/12"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingVpnOptions := models.OrgSettingVpnOptions{
        AsBase:               models.ToPointer(70),
        EnableIpv6:           models.ToPointer(false),
        StSubnet:             models.ToPointer("10.224.0.0/12"),
    }

}
```

