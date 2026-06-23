
# Vrf Extra Route 6

Additional IPv6 static route for a VRF instance

## Structure

`VrfExtraRoute6`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | IPv6 next-hop address for this VRF extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vrfExtraRoute6 := models.VrfExtraRoute6{
        Via:                  models.ToPointer("via8"),
    }

}
```

