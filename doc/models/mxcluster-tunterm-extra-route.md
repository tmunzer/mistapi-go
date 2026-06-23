
# Mxcluster Tunterm Extra Route

Extra route for Mist Tunneled VLAN traffic

## Structure

`MxclusterTuntermExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | Next-hop IP address for this extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxclusterTuntermExtraRoute := models.MxclusterTuntermExtraRoute{
        Via:                  models.ToPointer("via2"),
    }

}
```

