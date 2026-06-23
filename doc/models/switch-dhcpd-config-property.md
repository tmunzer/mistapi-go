
# Switch Dhcpd Config Property

DHCP server or relay configuration for one switch network. The property key is the network name. In case of DHCP relay, it's common for many networks to use the same dhcp relay, comma-separated network names can be used here (e.g. "net1,net2")

## Structure

`SwitchDhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsServers` | `[]string` | Optional | If `type`==`server` or `type6`==`server` - optional, if not defined, system one will be used |
| `DnsSuffix` | `[]string` | Optional | If `type`==`server` or `type6`==`server` - optional, if not defined, system one will be used |
| `FixedBindings` | [`map[string]models.DhcpdConfigFixedBinding`](../../doc/models/dhcpd-config-fixed-binding.md) | Optional | If `type`==`server` or `type6`==`server`. Property key is the MAC address. Format is `[0-9a-f]{12}` (e.g. "5684dae9ac8b") |
| `Gateway` | `*string` | Optional | If `type`==`server` - optional, `ip` will be used if not provided |
| `IpEnd` | `*string` | Optional | If `type`==`server`, ending IPv4 address for the DHCP lease pool |
| `IpEnd6` | `*string` | Optional | If `type6`==`server`, ending IPv6 address for the DHCP lease pool |
| `IpStart` | `*string` | Optional | If `type`==`server`, starting IPv4 address for the DHCP lease pool |
| `IpStart6` | `*string` | Optional | If `type6`==`server`, starting IPv6 address for the DHCP lease pool |
| `LeaseTime` | `*int` | Optional | In seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day]<br><br>**Default**: `86400`<br><br>**Constraints**: `>= 3600`, `<= 604800` |
| `Options` | [`map[string]models.DhcpdConfigOption`](../../doc/models/dhcpd-config-option.md) | Optional | If `type`==`server` or `type6`==`server`. Property key is the DHCP option number |
| `ServerIdOverride` | `*bool` | Optional | `server_id_override`==`true` means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,<br>should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address.<br><br>**Default**: `false` |
| `Servers` | `[]string` | Optional | If `type`==`relay`, upstream IPv4 DHCP servers |
| `Servers6` | `[]string` | Optional | If `type6`==`relay`, upstream IPv6 DHCP servers |
| `Type` | [`*models.SwitchDhcpdConfigTypeEnum`](../../doc/models/switch-dhcpd-config-type-enum.md) | Optional | enum: `none`, `relay` (DHCP Relay), `server` (DHCP Server) |
| `Type6` | [`*models.SwitchDhcpdConfigTypeEnum`](../../doc/models/switch-dhcpd-config-type-enum.md) | Optional | enum: `none`, `relay` (DHCP Relay), `server` (DHCP Server)<br><br>**Default**: `"none"` |
| `VendorEncapsulated` | [`map[string]models.DhcpdConfigVendorOption`](../../doc/models/dhcpd-config-vendor-option.md) | Optional | If `type`==`server` or `type6`==`server`. Property key is <enterprise number>:<sub option code>, with<br><br>* enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers)<br>* sub option code: 1-255, sub-option code' |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchDhcpdConfigProperty := models.SwitchDhcpdConfigProperty{
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
        IpEnd:                models.ToPointer("192.168.70.200"),
        IpEnd6:               models.ToPointer("2607:f8b0:4005:808::ff"),
        IpStart:              models.ToPointer("192.168.70.100"),
        IpStart6:             models.ToPointer("2607:f8b0:4005:808::2"),
        LeaseTime:            models.ToPointer(86400),
        ServerIdOverride:     models.ToPointer(false),
        Servers:              []string{
            "11.2.3.4",
        },
        Servers6:             []string{
            "2607:f8b0:4005:808::64",
        },
        Type6:                models.ToPointer(models.SwitchDhcpdConfigTypeEnum_NONE),
    }

}
```

