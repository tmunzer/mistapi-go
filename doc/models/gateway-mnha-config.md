
# Gateway Mnha Config

Multi-Node High Availability (MNHA) configuration, supported on SRX devices only. When enabled, the device operates in MNHA mode instead of chassis-cluster mode.

## Structure

`GatewayMnhaConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether MNHA mode is enabled<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMnhaConfig := models.GatewayMnhaConfig{
        Enabled:              models.ToPointer(false),
    }

}
```

