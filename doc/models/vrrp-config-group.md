
# Vrrp Config Group

VRRP group behavior settings

## Structure

`VrrpConfigGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Preempt` | `*bool` | Optional | If `true`, allow preemption (a backup router can preempt a primary router)<br><br>**Default**: `false` |
| `Priority` | `*int` | Optional | VRRP priority for this router in the group |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vrrpConfigGroup := models.VrrpConfigGroup{
        Preempt:              models.ToPointer(false),
        Priority:             models.ToPointer(64),
    }

}
```

