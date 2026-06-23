
# Evpn Options Underlay

EVPN underlay BGP and subnet settings

## Structure

`EvpnOptionsUnderlay`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsBase` | `*int` | Optional | Underlay BGP Base AS Number<br><br>**Default**: `65001`<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `RoutedIdPrefix` | `*string` | Optional | Prefix length used for automatically derived underlay router identifiers |
| `Subnet` | `*string` | Optional | Underlay subnet, by default, `10.255.240.0/20`, or `fd31:5700::/64` for ipv6 |
| `UseIpv6` | `*bool` | Optional | If v6 is desired for underlay<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnOptionsUnderlay := models.EvpnOptionsUnderlay{
        AsBase:               models.ToPointer(65001),
        RoutedIdPrefix:       models.ToPointer("/24"),
        Subnet:               models.ToPointer("10.255.240.0/20"),
        UseIpv6:              models.ToPointer(false),
    }

}
```

