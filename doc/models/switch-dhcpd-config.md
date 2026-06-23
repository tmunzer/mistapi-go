
# Switch Dhcpd Config

Switch DHCP server or relay configuration keyed by network name

*This model accepts additional fields of type [models.SwitchDhcpdConfigProperty](../../doc/models/switch-dhcpd-config-property.md).*

## Structure

`SwitchDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether switch DHCP server or relay configuration is enabled<br><br>**Default**: `false` |
| `AdditionalProperties` | [`map[string]models.SwitchDhcpdConfigProperty`](../../doc/models/switch-dhcpd-config-property.md) | Optional | DHCP server or relay configuration for one switch network. The property key is the network name. In case of DHCP relay, it's common for many networks to use the same dhcp relay, comma-separated network names can be used here (e.g. "net1,net2") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchDhcpdConfig := models.SwitchDhcpdConfig{
        Enabled:              models.ToPointer(false),
        AdditionalProperties: map[string]models.SwitchDhcpdConfigProperty{
            "exampleAdditionalProperty": models.SwitchDhcpdConfigProperty{
                DnsServers:           []string{
                    "dns_servers8",
                    "dns_servers9",
                },
                DnsSuffix:            []string{
                    "dns_suffix5",
                    "dns_suffix6",
                },
                FixedBindings:        map[string]models.DhcpdConfigFixedBinding{
                    "key0": nil,
                    "key1": models.DhcpdConfigFixedBinding{
                    },
                },
                Gateway:              models.ToPointer("gateway2"),
                IpEnd:                models.ToPointer("ip_end2"),
            },
        },
    }

}
```

