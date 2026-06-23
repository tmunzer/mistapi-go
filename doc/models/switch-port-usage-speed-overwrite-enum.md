
# Switch Port Usage Speed Overwrite Enum

Port Speed, default is auto to automatically negotiate speed enum: `100m`, `10m`, `1g`, `2.5g`, `5g`, `10g`, `25g`, `40g`, `100g`,`auto`

## Enumeration

`SwitchPortUsageSpeedOverwriteEnum`

## Fields

| Name |
|  --- |
| `10m` |
| `100m` |
| `1g` |
| `2.5g` |
| `5g` |
| `10g` |
| `25g` |
| `40g` |
| `100g` |
| `auto` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortUsageSpeedOverwrite := models.SwitchPortUsageSpeedOverwriteEnum_ENUM10M

}
```

