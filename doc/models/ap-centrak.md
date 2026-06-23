
# Ap Centrak

CenTrak integration settings for an AP

## Structure

`ApCentrak`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable Centrak config<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apCentrak := models.ApCentrak{
        Enabled:              models.ToPointer(false),
    }

}
```

