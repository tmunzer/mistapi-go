
# Dhcpd Config Property

DHCP server or relay configuration for one network

## Structure

`DhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsServers` | `[]string` | Optional | If `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used |
| `DnsSuffix` | `[]string` | Optional | If `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used |
| `FixedBindings` | [`map[string]models.DhcpdConfigFixedBinding`](../../doc/models/dhcpd-config-fixed-binding.md) | Optional | If `type`==`local` or `type6`==`local`. Property key is the client MAC address. Format is `[0-9a-f]{12}` (e.g. "5684dae9ac8b") |
| `Gateway` | `*string` | Optional | If `type`==`local` - optional, `ip` will be used if not provided |
| `Ip6End` | `*string` | Optional | If `type6`==`local`, ending IPv6 address for the DHCP lease pool |
| `Ip6Start` | `*string` | Optional | If `type6`==`local`, starting IPv6 address for the DHCP lease pool |
| `IpEnd` | `*string` | Optional | If `type`==`local`, ending IPv4 address for the DHCP lease pool |
| `IpStart` | `*string` | Optional | If `type`==`local`, starting IPv4 address for the DHCP lease pool |
| `LeaseTime` | `*int` | Optional | In seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day]<br><br>**Default**: `86400`<br><br>**Constraints**: `>= 3600`, `<= 604800` |
| `Options` | [`map[string]models.DhcpdConfigOption`](../../doc/models/dhcpd-config-option.md) | Optional | If `type`==`local` or `type6`==`local`. Property key is the DHCP option number |
| `ServerIdOverride` | `*bool` | Optional | `server_id_override`==`true` means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,<br>should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address.<br><br>**Default**: `false` |
| `Servers` | `[]string` | Optional | If `type`==`relay`, upstream IPv4 DHCP servers |
| `Serversv6` | `[]string` | Optional | If `type6`==`relay`, upstream IPv6 DHCP servers |
| `Type` | [`*models.DhcpdConfigTypeEnum`](../../doc/models/dhcpd-config-type-enum.md) | Optional | enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)<br><br>**Default**: `"local"` |
| `Type6` | [`*models.DhcpdConfigTypeEnum`](../../doc/models/dhcpd-config-type-enum.md) | Optional | enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)<br><br>**Default**: `"none"` |
| `VendorEncapsulated` | [`map[string]models.DhcpdConfigVendorOption`](../../doc/models/dhcpd-config-vendor-option.md) | Optional | If `type`==`local` or `type6`==`local`. Property key is <enterprise number>:<sub option code>, with<br><br>* enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers)<br>* sub option code: 1-255, sub-option code |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpdConfigProperty := models.DhcpdConfigProperty{
        DnsServers:           []string{
            "8.8.8.8",
            "4.4.4.4",
            "2001:4860:4860::8888",
        },
        DnsSuffix:            []string{
            ".mist.local",
            ".mist.com",
        },
        FixedBindings:        map[string]models.DhcpdConfigFixedBinding{
            "5684dae9ac8b": models.DhcpdConfigFixedBinding{
                Ip:                   models.ToPointer("192.168.70.35"),
                Ip6:                  models.ToPointer("ip66"),
                Name:                 models.ToPointer("John"),
            },
        },
        Gateway:              models.ToPointer("192.168.70.1"),
        Ip6End:               models.ToPointer("2607:f8b0:4005:808::ff"),
        Ip6Start:             models.ToPointer("2607:f8b0:4005:808::2"),
        IpEnd:                models.ToPointer("192.168.70.200"),
        IpStart:              models.ToPointer("192.168.70.100"),
        LeaseTime:            models.ToPointer(86400),
        ServerIdOverride:     models.ToPointer(false),
        Servers:              []string{
            "11.2.3.4",
        },
        Serversv6:            []string{
            "2607:f8b0:4005:808::64",
        },
        Type:                 models.ToPointer(models.DhcpdConfigTypeEnum_LOCAL),
        Type6:                models.ToPointer(models.DhcpdConfigTypeEnum_NONE),
    }

}
```

