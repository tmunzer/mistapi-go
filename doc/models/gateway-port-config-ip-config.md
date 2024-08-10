
# Gateway Port Config Ip Config

Junos IP Config

## Structure

`GatewayPortConfigIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | except for out-of_band interface (vme/em0/fxp0) |
| `DnsSuffix` | `[]string` | Optional | except for out-of_band interface (vme/em0/fxp0) |
| `Gateway` | `*string` | Optional | except for out-of_band interface (vme/em0/fxp0) |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | used only if `subnet` is not specified in `networks` |
| `Network` | `*string` | Optional | optional, the network to be used for mgmt |
| `PoserPassword` | `*string` | Optional | if `type`==`pppoe` |
| `PppoeAuth` | [`*models.GatewayWanPpoeAuthEnum`](../../doc/models/gateway-wan-ppoe-auth-enum.md) | Optional | if `type`==`pppoe`. enum: `chap`, `none`, `pap`<br>**Default**: `"none"` |
| `PppoeUsername` | `*string` | Optional | if `type`==`pppoe` |
| `Type` | [`*models.GatewayWanTypeEnum`](../../doc/models/gateway-wan-type-enum.md) | Optional | enum: `dhcp`, `pppoe`, `static`<br>**Default**: `"dhcp"` |

## Example (as JSON)

```json
{
  "pppoe_auth": "none",
  "type": "dhcp",
  "dns": [
    "dns7",
    "dns8"
  ],
  "dns_suffix": [
    "dns_suffix9",
    "dns_suffix0"
  ],
  "gateway": "gateway0",
  "ip": "ip4",
  "netmask": "netmask0"
}
```

