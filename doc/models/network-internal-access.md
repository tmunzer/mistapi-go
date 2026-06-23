
# Network Internal Access

Internal access settings for an organization network

## Structure

`NetworkInternalAccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether internal access is enabled for this network |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkInternalAccess := models.NetworkInternalAccess{
        Enabled:              models.ToPointer(false),
    }

}
```

