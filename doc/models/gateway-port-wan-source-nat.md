
# Gateway Port Wan Source Nat

Only if `usage`==`wan`, optional. By default, source-NAT is performed on all WAN Ports using the interface-ip

## Structure

`GatewayPortWanSourceNat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | Or to disable the source-nat<br><br>**Default**: `false` |
| `Nat6Pool` | `*string` | Optional | If alternative nat_pool is desired |
| `NatPool` | `*string` | Optional | If alternative nat_pool is desired |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayPortWanSourceNat := models.GatewayPortWanSourceNat{
        Disabled:             models.ToPointer(false),
        Nat6Pool:             models.ToPointer("2601:1700:43c0:dc0:20c:29ff:fea7:93bc/126"),
        NatPool:              models.ToPointer("64.2.4.0/30"),
    }

}
```

