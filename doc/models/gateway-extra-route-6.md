
# Gateway Extra Route 6

Gateway IPv6 extra route next-hop settings

## Structure

`GatewayExtraRoute6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | Next-hop IPv6 address for the gateway extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayExtraRoute6 := models.GatewayExtraRoute6{
        Via:                  models.ToPointer("via8"),
    }

}
```

