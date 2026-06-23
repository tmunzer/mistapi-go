
# Switch Port Usage Dynamic Reset Default when Enum

Only if `mode`==`dynamic` Control when the DPC port should be changed to the default port usage. enum: `link_down`, `none` (let the DPC port keep at the current port usage)

## Enumeration

`SwitchPortUsageDynamicResetDefaultWhenEnum`

## Fields

| Name |
|  --- |
| `link_down` |
| `none` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchPortUsageDynamicResetDefaultWhen := models.SwitchPortUsageDynamicResetDefaultWhenEnum_LINKDOWN

}
```

