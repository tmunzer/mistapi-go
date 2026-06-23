
# Marvis Client Synthetic Test

Synthetic test settings for Marvis Client

## Structure

`MarvisClientSyntheticTest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether synthetic testing is enabled for Marvis Client |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    marvisClientSyntheticTest := models.MarvisClientSyntheticTest{
        Enabled:              models.ToPointer(false),
    }

}
```

