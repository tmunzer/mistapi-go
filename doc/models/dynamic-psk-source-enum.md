
# Dynamic Psk Source Enum

Origin used to retrieve per-user PSKs. enum: `cloud_psks`, `radius`

## Enumeration

`DynamicPskSourceEnum`

## Fields

| Name |
|  --- |
| `cloud_psks` |
| `radius` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dynamicPskSource := models.DynamicPskSourceEnum_CLOUDPSKS

}
```

