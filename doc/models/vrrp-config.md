
# Vrrp Config

Junos VRRP configuration applied to a switch or switch profile

## Structure

`VrrpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether VRRP configuration is enabled |
| `Groups` | [`map[string]models.VrrpConfigGroup`](../../doc/models/vrrp-config-group.md) | Optional | Property key is the VRRP name |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vrrpConfig := models.VrrpConfig{
        Enabled:              models.ToPointer(false),
        Groups:               map[string]models.VrrpConfigGroup{
            "key0": models.VrrpConfigGroup{
                Preempt:              models.ToPointer(false),
                Priority:             models.ToPointer(102),
            },
        },
    }

}
```

