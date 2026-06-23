
# Network Source Nat

If `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub

## Structure

`NetworkSourceNat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExternalIp` | `*string` | Optional | External source NAT IP or subnet used when spoke hosts must be reachable from the hub |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkSourceNat := models.NetworkSourceNat{
        ExternalIp:           models.ToPointer("172.16.0.8/30"),
    }

}
```

