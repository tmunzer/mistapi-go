
# Device Ap Lacp Config

LACP settings for supported AP Ethernet uplinks

## Structure

`DeviceApLacpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable LACP on supported AP Ethernet uplinks<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    deviceApLacpConfig := models.DeviceApLacpConfig{
        Enabled:              models.ToPointer(false),
    }

}
```

