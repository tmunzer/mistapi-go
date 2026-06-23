
# Tunterm Dhcpd Config Property

Per-VLAN DHCP relay configuration for tunnel termination

## Structure

`TuntermDhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether DHCP relay is enabled for this tunneled VLAN<br><br>**Default**: `false` |
| `Servers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Type` | [`*models.TuntermDhcpdTypeEnum`](../../doc/models/tunterm-dhcpd-type-enum.md) | Optional | DHCP handling mode for tunnel termination. enum: `relay`<br><br>**Default**: `"relay"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tuntermDhcpdConfigProperty := models.TuntermDhcpdConfigProperty{
        Enabled:              models.ToPointer(false),
        Servers:              []string{
            "servers5",
            "servers6",
        },
        Type:                 models.ToPointer(models.TuntermDhcpdTypeEnum_RELAY),
    }

}
```

