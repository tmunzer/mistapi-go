
# Mxcluster Radsec Server Selection Enum

When ordered, Mist Edge will prefer and go back to the first RADIUS server if possible. enum: `ordered`, `unordered`

## Enumeration

`MxclusterRadsecServerSelectionEnum`

## Fields

| Name |
|  --- |
| `ordered` |
| `unordered` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxclusterRadsecServerSelection := models.MxclusterRadsecServerSelectionEnum_ORDERED

}
```

