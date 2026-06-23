
# Gateway Extra Route

Gateway IPv4 extra route next-hop settings

## Structure

`GatewayExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | Next-hop IPv4 address for the gateway extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayExtraRoute := models.GatewayExtraRoute{
        Via:                  models.ToPointer("via8"),
    }

}
```

