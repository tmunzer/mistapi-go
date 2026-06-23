
# Marvis Client Location

Location collection settings for Marvis Client

## Structure

`MarvisClientLocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether location collection is enabled for Marvis Client |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisClientLocation := models.MarvisClientLocation{
        Enabled:              models.ToPointer(false),
    }

}
```

