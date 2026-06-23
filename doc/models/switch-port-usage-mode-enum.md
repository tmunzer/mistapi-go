
# Switch Port Usage Mode Enum

`mode`==`dynamic` must only be used if the port usage name is `dynamic`. enum: `access`, `dynamic`, `inet`, `trunk`

## Enumeration

`SwitchPortUsageModeEnum`

## Fields

| Name |
|  --- |
| `access` |
| `dynamic` |
| `inet` |
| `trunk` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortUsageMode := models.SwitchPortUsageModeEnum_ACCESS

}
```

