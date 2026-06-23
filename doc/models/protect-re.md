
# Protect Re

Restrict inbound-traffic to host
when enabled, all traffic that is not essential to our operation will be dropped
e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works

## Structure

`ProtectRe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowedServices` | [`[]models.ProtectReAllowedServiceEnum`](../../doc/models/protect-re-allowed-service-enum.md) | Optional | Built-in services explicitly allowed by the Protect RE policy |
| `Custom` | [`[]models.ProtectReCustom`](../../doc/models/protect-re-custom.md) | Optional | Custom Protect RE ACL entries |
| `Enabled` | `*bool` | Optional | When enabled, all traffic that is not essential to our operation will be dropped<br>e.g. ntp / dns / traffic to mist will be allowed by default<br>if dhcpd is enabled, we'll make sure it works<br><br>**Default**: `false` |
| `HitCount` | `*bool` | Optional | Whether to enable hit count for Protect_RE policy<br><br>**Default**: `false` |
| `TrustedHosts` | `[]string` | Optional | Host or subnet entries allowed by the Protect RE policy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    protectRe := models.ProtectRe{
        AllowedServices:      []models.ProtectReAllowedServiceEnum{
            models.ProtectReAllowedServiceEnum_ICMP,
            models.ProtectReAllowedServiceEnum_SSH,
        },
        Custom:               []models.ProtectReCustom{
            models.ProtectReCustom{
                PortRange:            models.ToPointer("port_range6"),
                Protocol:             models.ToPointer(models.ProtectReCustomProtocolEnum_ANY),
                Subnets:              []string{
                    "subnets5",
                    "subnets6",
                },
            },
            models.ProtectReCustom{
                PortRange:            models.ToPointer("port_range6"),
                Protocol:             models.ToPointer(models.ProtectReCustomProtocolEnum_ANY),
                Subnets:              []string{
                    "subnets5",
                    "subnets6",
                },
            },
        },
        Enabled:              models.ToPointer(false),
        HitCount:             models.ToPointer(false),
        TrustedHosts:         []string{
            "trusted_hosts6",
            "trusted_hosts7",
        },
    }

}
```

