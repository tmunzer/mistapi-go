
# Mxedge Tunterm Dhcpd Config

Global and per-VLAN DHCP relay settings for Mist Tunneled VLANs; property key is the VLAN ID

*This model accepts additional fields of type [models.MxedgeTuntermDhcpdConfigProperty](../../doc/models/mxedge-tunterm-dhcpd-config-property.md).*

## Structure

`MxedgeTuntermDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether DHCP relay is enabled for Mist Tunneled VLANs<br><br>**Default**: `false` |
| `Servers` | `[]string` | Optional | List of DHCP servers; required if `type`==`relay` |
| `Type` | [`*models.MxedgeTuntermDhcpdTypeEnum`](../../doc/models/mxedge-tunterm-dhcpd-type-enum.md) | Optional | DHCP handling mode for the global tunnel termination config. enum: `relay`<br><br>**Default**: `"relay"` |
| `AdditionalProperties` | [`map[string]models.MxedgeTuntermDhcpdConfigProperty`](../../doc/models/mxedge-tunterm-dhcpd-config-property.md) | Optional | Per-VLAN DHCP relay settings for a Mist Tunneled VLAN |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermDhcpdConfig := models.MxedgeTuntermDhcpdConfig{
        Enabled:              models.ToPointer(false),
        Servers:              []string{
            "servers9",
            "servers0",
        },
        Type:                 models.ToPointer(models.MxedgeTuntermDhcpdTypeEnum_RELAY),
        AdditionalProperties: map[string]models.MxedgeTuntermDhcpdConfigProperty{
            "exampleAdditionalProperty": models.MxedgeTuntermDhcpdConfigProperty{
                Enabled:              models.ToPointer(false),
                Servers:              []string{
                    "servers9",
                    "servers0",
                    "servers1",
                },
                Type:                 models.ToPointer(models.MxedgeTuntermDhcpdConfigTypeEnum_RELAY),
            },
        },
    }

}
```

