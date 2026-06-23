
# Tunterm Dhcpd Config

DHCP server/relay configuration of Mist Tunneled VLANs. Property key is the VLAN ID

*This model accepts additional fields of type [models.TuntermDhcpdConfigProperty](../../doc/models/tunterm-dhcpd-config-property.md).*

## Structure

`TuntermDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether DHCP relay is enabled for tunnel termination VLANs<br><br>**Default**: `false` |
| `Servers` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Type` | [`*models.TuntermDhcpdTypeEnum`](../../doc/models/tunterm-dhcpd-type-enum.md) | Optional | DHCP handling mode for tunnel termination. enum: `relay`<br><br>**Default**: `"relay"` |
| `AdditionalProperties` | [`map[string]models.TuntermDhcpdConfigProperty`](../../doc/models/tunterm-dhcpd-config-property.md) | Optional | Per-VLAN DHCP relay configuration for tunnel termination |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tuntermDhcpdConfig := models.TuntermDhcpdConfig{
        Enabled:              models.ToPointer(false),
        Servers:              []string{
            "servers1",
        },
        Type:                 models.ToPointer(models.TuntermDhcpdTypeEnum_RELAY),
        AdditionalProperties: map[string]models.TuntermDhcpdConfigProperty{
            "exampleAdditionalProperty": models.TuntermDhcpdConfigProperty{
                Enabled:              models.ToPointer(false),
                Servers:              []string{
                    "servers7",
                },
                Type:                 models.ToPointer(models.TuntermDhcpdTypeEnum_RELAY),
            },
        },
    }

}
```

