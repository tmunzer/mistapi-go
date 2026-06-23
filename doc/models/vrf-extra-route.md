
# Vrf Extra Route

Additional IPv4 static route for a VRF instance

## Structure

`VrfExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | IPv4 next-hop address for this VRF extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vrfExtraRoute := models.VrfExtraRoute{
        Via:                  models.ToPointer("via8"),
    }

}
```

