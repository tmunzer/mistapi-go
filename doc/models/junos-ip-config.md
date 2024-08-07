
# Junos Ip Config

Junos IP Config

## Structure

`JunosIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `DnsSuffix` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Gateway` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | used only if `subnet` is not specified in `networks` |
| `Network` | `*string` | Optional | the network where this mgmt IP reside, this will be used as default network for outbound-ssh, dns, ntp, dns, tacplus, radius, syslog, snmp |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |

## Example (as JSON)

```json
{
  "type": "static",
  "dns": [
    "dns7"
  ],
  "dns_suffix": [
    "dns_suffix3"
  ],
  "gateway": "gateway4",
  "ip": "ip8",
  "netmask": "netmask4"
}
```

