
# Dhcpd Config

DHCP server configuration map with a global enable flag

*This model accepts additional fields of type [models.DhcpdConfigProperty](../../doc/models/dhcpd-config-property.md).*

## Structure

`DhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | If set to `false`, disable the DHCP server<br><br>**Default**: `true` |
| `AdditionalProperties` | [`map[string]models.DhcpdConfigProperty`](../../doc/models/dhcpd-config-property.md) | Optional | DHCP server or relay configuration for one network |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpdConfig := models.DhcpdConfig{
        Enabled:              models.ToPointer(true),
        AdditionalProperties: map[string]models.DhcpdConfigProperty{
            "exampleAdditionalProperty": models.DhcpdConfigProperty{
                DnsServers:           []string{
                    "dns_servers2",
                    "dns_servers3",
                    "dns_servers4",
                },
                DnsSuffix:            []string{
                    "dns_suffix1",
                    "dns_suffix0",
                    "dns_suffix9",
                },
                FixedBindings:        map[string]models.DhcpdConfigFixedBinding{
                    "key0": nil,
                },
                Gateway:              models.ToPointer("gateway8"),
                Ip6End:               models.ToPointer("ip6_end8"),
            },
        },
    }

}
```

