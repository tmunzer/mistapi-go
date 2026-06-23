
# Mxedge Tunterm Extra Route

Extra route for Mist Tunnel traffic on a Mist Edge

## Structure

`MxedgeTuntermExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | Next-hop IP address for this Mist Tunnel extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermExtraRoute := models.MxedgeTuntermExtraRoute{
        Via:                  models.ToPointer("via8"),
    }

}
```

