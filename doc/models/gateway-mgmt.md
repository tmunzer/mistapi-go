
# Gateway Mgmt

Gateway management-plane and access settings

## Structure

`GatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminSshkeys` | `[]string` | Optional | For SSR only, as direct root access is not allowed |
| `AppProbing` | [`*models.AppProbing`](../../doc/models/app-probing.md) | Optional | Application reachability probing settings for gateway management |
| `AppUsage` | `*bool` | Optional | Consumes uplink bandwidth, requires WA license |
| `AutoSignatureUpdate` | [`*models.GatewayMgmtAutoSignatureUpdate`](../../doc/models/gateway-mgmt-auto-signature-update.md) | Optional | Automatic security signature update schedule |
| `ConfigRevertTimer` | `*int` | Optional | Rollback timer for commit confirmed<br><br>**Default**: `10`<br><br>**Constraints**: `>= 1`, `<= 30` |
| `DisableConsole` | `*bool` | Optional | For SSR and SRX, disable console port<br><br>**Default**: `false` |
| `DisableOob` | `*bool` | Optional | For SSR and SRX, disable management interface<br><br>**Default**: `false` |
| `DisableUsb` | `*bool` | Optional | For SSR and SRX, disable usb interface<br><br>**Default**: `false` |
| `FipsEnabled` | `*bool` | Optional | Whether FIPS mode is enabled on the gateway<br><br>**Default**: `false` |
| `ProbeHosts` | `[]string` | Optional | IPv4 probe hosts used for gateway connectivity checks |
| `ProbeHostsv6` | `[]string` | Optional | IPv6 probe hosts used for gateway connectivity checks |
| `ProtectRe` | [`*models.ProtectRe`](../../doc/models/protect-re.md) | Optional | Restrict inbound-traffic to host<br>when enabled, all traffic that is not essential to our operation will be dropped<br>e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works |
| `RootPassword` | `*string` | Optional | SRX only. Root password for local gateway access |
| `SecurityLogSourceAddress` | `*string` | Optional | IPv4 source address used for gateway security log traffic |
| `SecurityLogSourceInterface` | `*string` | Optional | Source interface used for gateway security log traffic |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayMgmt := models.GatewayMgmt{
        AdminSshkeys:               []string{
            "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host",
        },
        AppProbing:                 models.ToPointer(models.AppProbing{
            Apps:                 []string{
                "apps8",
                "apps9",
                "apps0",
            },
            CustomApps:           []models.AppProbingCustomApp{
                models.AppProbingCustomApp{
                    Address:              models.ToPointer("address4"),
                    AppType:              models.ToPointer("app_type2"),
                    Hostnames:            []string{
                        "hostnames4",
                        "hostnames5",
                    },
                    Key:                  models.ToPointer("key8"),
                    Name:                 models.ToPointer("name8"),
                },
                models.AppProbingCustomApp{
                    Address:              models.ToPointer("address4"),
                    AppType:              models.ToPointer("app_type2"),
                    Hostnames:            []string{
                        "hostnames4",
                        "hostnames5",
                    },
                    Key:                  models.ToPointer("key8"),
                    Name:                 models.ToPointer("name8"),
                },
            },
            Enabled:              models.ToPointer(false),
        }),
        AppUsage:                   models.ToPointer(false),
        AutoSignatureUpdate:        models.ToPointer(models.GatewayMgmtAutoSignatureUpdate{
            DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_ANY),
            Enable:               models.ToPointer(false),
            TimeOfDay:            models.ToPointer("time_of_day6"),
        }),
        ConfigRevertTimer:          models.ToPointer(10),
        DisableConsole:             models.ToPointer(false),
        DisableOob:                 models.ToPointer(false),
        DisableUsb:                 models.ToPointer(false),
        FipsEnabled:                models.ToPointer(false),
        ProbeHosts:                 []string{
            "8.8.8.8",
        },
        ProbeHostsv6:               []string{
            "2001:4860:4860::8888",
        },
        SecurityLogSourceAddress:   models.ToPointer("192.168.1.1"),
        SecurityLogSourceInterface: models.ToPointer("ge-0/0/1.0"),
    }

}
```

