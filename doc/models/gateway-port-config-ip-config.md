
# Gateway Port Config Ip Config

Junos IP Config

## Structure

`GatewayPortConfigIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | Except for out-of_band interface (vme/em0/fxp0) |
| `DnsSuffix` | `[]string` | Optional | Except for out-of_band interface (vme/em0/fxp0) |
| `Gateway` | `*string` | Optional | Except for out-of_band interface (vme/em0/fxp0). Interface Default Gateway IP Address (i.e. "192.168.1.1") or a Variable (i.e. "{{myvar}}") |
| `Gateway6` | `*string` | Optional | Except for out-of_band interface (vme/em0/fxp0). Interface Default Gateway IPv6 Address (i.e. "2001:db8::1") or a Variable (i.e. "{{myvar}}") |
| `Ip` | `*string` | Optional | Interface IP Address (i.e. "192.168.1.8") or a Variable (i.e. "{{myvar}}") |
| `Ip6` | `*string` | Optional | Interface IPv6 Address (i.e. "2001:db8::123") or a Variable (i.e. "{{myvar}}") |
| `Netmask` | `*string` | Optional | Used only if `subnet` is not specified in `networks`. Interface Netmask (i.e. "/24") or a Variable (i.e. "{{myvar}}") |
| `Netmask6` | `*string` | Optional | Used only if `subnet` is not specified in `networks`. Interface IPv6 Netmask (i.e. "/64") or a Variable (i.e. "{{myvar}}") |
| `Network` | `*string` | Optional | Optional, the network to be used for mgmt |
| `PoserPassword` | `*string` | Optional | If `type`==`pppoe` |
| `PppoeAuth` | [`*models.GatewayWanPpoeAuthEnum`](../../doc/models/gateway-wan-ppoe-auth-enum.md) | Optional | if `type`==`pppoe`. enum: `chap`, `none`, `pap`<br><br>**Default**: `"none"` |
| `PppoeUsername` | `*string` | Optional | If `type`==`pppoe` |
| `Type` | [`*models.GatewayWanTypeEnum`](../../doc/models/gateway-wan-type-enum.md) | Optional | enum: `dhcp`, `pppoe`, `static`<br><br>**Default**: `"dhcp"` |
| `Type6` | [`*models.GatewayWanType6Enum`](../../doc/models/gateway-wan-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `static`<br><br>**Default**: `"autoconf"` |

## Example (as JSON)

```json
{
  "gateway": "192.168.1.1",
  "gateway6": "2001:db8::1",
  "ip": "192.168.1.8",
  "ip6": "2001:db8::123",
  "netmask": "/24",
  "netmask6": "/64",
  "pppoe_auth": "none",
  "type": "dhcp",
  "type6": "autoconf",
  "dns": [
    "dns7",
    "dns8"
  ],
  "dns_suffix": [
    "dns_suffix9",
    "dns_suffix0"
  ]
}
```

