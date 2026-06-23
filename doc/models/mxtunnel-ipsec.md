
# Mxtunnel Ipsec

IPsec settings for a Mist Tunnel

## Structure

`MxtunnelIpsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsServers` | `models.Optional[[]string]` | Optional | Name server addresses advertised for IPsec tunnel clients |
| `DnsSuffix` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Enabled` | `*bool` | Optional | Whether IPsec support is enabled for this Mist Tunnel |
| `ExtraRoutes` | [`[]models.MxtunnelIpsecExtraRoute`](../../doc/models/mxtunnel-ipsec-extra-route.md) | Optional | Additional routes advertised for IPsec tunnel clients |
| `SplitTunnel` | `*bool` | Optional | Whether split tunneling is enabled for IPsec clients |
| `UseMxedge` | `*bool` | Optional | Whether IPsec termination uses Mist Edge |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxtunnelIpsec := models.MxtunnelIpsec{
        DnsServers:           models.NewOptional(models.ToPointer([]string{
            "dns_servers6",
            "dns_servers7",
            "dns_servers8",
        })),
        DnsSuffix:            []string{
            "dns_suffix7",
            "dns_suffix6",
            "dns_suffix5",
        },
        Enabled:              models.ToPointer(false),
        ExtraRoutes:          []models.MxtunnelIpsecExtraRoute{
            models.MxtunnelIpsecExtraRoute{
                Dest:                 models.ToPointer("dest4"),
                NextHop:              models.ToPointer("next_hop4"),
            },
            models.MxtunnelIpsecExtraRoute{
                Dest:                 models.ToPointer("dest4"),
                NextHop:              models.ToPointer("next_hop4"),
            },
            models.MxtunnelIpsecExtraRoute{
                Dest:                 models.ToPointer("dest4"),
                NextHop:              models.ToPointer("next_hop4"),
            },
        },
        SplitTunnel:          models.ToPointer(false),
    }

}
```

