
# Const Device Switch Default

Default switch port ranges for a model

## Structure

`ConstDeviceSwitchDefault`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | `*string` | Optional | Default switch port range list for this model |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceSwitchDefault := models.ConstDeviceSwitchDefault{
        Ports:                models.ToPointer("ge-0/0/0-47, et-0/1/0-3, xe-0/2/0-3, ge-0/2/0-3"),
    }

}
```

