
# Optic Port Config Port

Per-interface optic port override settings

## Structure

`OpticPortConfigPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channelized` | `*bool` | Optional | Whether channelization is enabled on this optic port<br><br>**Default**: `false` |
| `Speed` | `*string` | Optional | Interface speed (e.g. `25g`, `50g`), use the chassis speed by default |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    opticPortConfigPort := models.OpticPortConfigPort{
        Channelized:          models.ToPointer(false),
        Speed:                models.ToPointer("25g"),
    }

}
```

