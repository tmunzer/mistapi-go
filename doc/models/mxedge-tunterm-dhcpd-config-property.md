
# Mxedge Tunterm Dhcpd Config Property

Per-VLAN DHCP relay settings for a Mist Tunneled VLAN

## Structure

`MxedgeTuntermDhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether DHCP relay is enabled for this tunneled VLAN<br><br>**Default**: `false` |
| `Servers` | `[]string` | Optional | List of DHCP servers; required if `type`==`relay` |
| `Type` | [`*models.MxedgeTuntermDhcpdConfigTypeEnum`](../../doc/models/mxedge-tunterm-dhcpd-config-type-enum.md) | Optional | DHCP handling mode for this tunneled VLAN. enum: `relay`<br><br>**Default**: `"relay"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermDhcpdConfigProperty := models.MxedgeTuntermDhcpdConfigProperty{
        Enabled:              models.ToPointer(false),
        Servers:              []string{
            "servers5",
            "servers6",
            "servers7",
        },
        Type:                 models.ToPointer(models.MxedgeTuntermDhcpdConfigTypeEnum_RELAY),
    }

}
```

