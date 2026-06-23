
# Vrf Config

VRF enablement settings applied when supported on the device

## Structure

`VrfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable VRF (when supported on the device) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vrfConfig := models.VrfConfig{
        Enabled:              models.ToPointer(false),
    }

}
```

