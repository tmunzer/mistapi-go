
# Dhcpd Config Property

*This model accepts additional fields of type interface{}.*

## Structure

`DhcpdConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsServers` | `[]string` | Optional | If `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used |
| `DnsSuffix` | `[]string` | Optional | If `type`==`local` or `type6`==`local` - optional, if not defined, system one will be used |
| `FixedBindings` | [`map[string]models.DhcpdConfigFixedBinding`](../../doc/models/dhcpd-config-fixed-binding.md) | Optional | If `type`==`local` or `type6`==`local`. Property key is the MAC Address. Format is `[0-9a-f]{12}` (e.g. "5684dae9ac8b") |
| `Gateway` | `*string` | Optional | If `type`==`local` - optional, `ip` will be used if not provided |
| `IpEnd` | `*string` | Optional | If `type`==`local` |
| `IpEnd6` | `*string` | Optional | If `type6`==`local` |
| `IpStart` | `*string` | Optional | If `type`==`local` |
| `IpStart6` | `*string` | Optional | If `type6`==`local` |
| `LeaseTime` | `*int` | Optional | In seconds, lease time has to be between 3600 [1hr] - 604800 [1 week], default is 86400 [1 day]<br>**Default**: `86400`<br>**Constraints**: `>= 3600`, `<= 604800` |
| `Options` | [`map[string]models.DhcpdConfigOption`](../../doc/models/dhcpd-config-option.md) | Optional | If `type`==`local` or `type6`==`local`. Property key is the DHCP option number |
| `ServerIdOverride` | `*bool` | Optional | `server_id_override`==`true` means the device, when acts as DHCP relay and forwards DHCP responses from DHCP server to clients,<br>should overwrite the Sever Identifier option (i.e. DHCP option 54) in DHCP responses with its own IP address.<br>**Default**: `false` |
| `Servers` | `[]string` | Optional | If `type`==`relay` |
| `Servers6` | `[]string` | Optional | If `type6`==`relay` |
| `Type` | [`*models.DhcpdConfigTypeEnum`](../../doc/models/dhcpd-config-type-enum.md) | Optional | enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)<br>**Default**: `"local"` |
| `Type6` | [`*models.DhcpdConfigTypeEnum`](../../doc/models/dhcpd-config-type-enum.md) | Optional | enum: `local` (DHCP Server), `none`, `relay` (DHCP Relay)<br>**Default**: `"none"` |
| `VendorEncapsulated` | [`map[string]models.DhcpdConfigVendorOption`](../../doc/models/dhcpd-config-vendor-option.md) | Optional | If `type`==`local` or `type6`==`local`. Property key is <enterprise number>:<sub option code>, with<br><br>* enterprise number: 1-65535 (https://www.iana.org/assignments/enterprise-numbers/enterprise-numbers)<br>* sub option code: 1-255, sub-option code |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dns_servers": [
    "8.8.8.8",
    "4.4.4.4"
  ],
  "dns_suffix": [
    ".mist.local",
    ".mist.com"
  ],
  "fixed_bindings": {
    "5684dae9ac8b": {
      "ip": "192.168.70.35",
      "name": "John",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "gateway": "192.168.70.1",
  "ip_end": "192.168.70.200",
  "ip_end6": "2607:f8b0:4005:808::ff",
  "ip_start": "192.168.70.100",
  "ip_start6": "2607:f8b0:4005:808::2",
  "lease_time": 86400,
  "server_id_override": false,
  "servers": [
    "11.2.3.4"
  ],
  "servers6": [
    "2607:f8b0:4005:808::64"
  ],
  "type": "local",
  "type6": "none",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

