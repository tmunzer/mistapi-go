
# Gateway Mgmt Host Out Policy

Host-out path policy for gateway-originated management traffic

## Structure

`GatewayMgmtHostOutPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PathPreference` | `*string` | Optional | Preferred path name for this gateway-originated service traffic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMgmtHostOutPolicy := models.GatewayMgmtHostOutPolicy{
        PathPreference:       models.ToPointer("path_preference8"),
    }

}
```

