
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
| `PppoeAuth` | [`*models.GatewayWanPpoeAuthEnum`](../../doc/models/gateway-wan-ppoe-auth-enum.md) | Optional | if `type`==`pppoe`<br>**Default**: `"none"` |
| `PppoeUsername` | `*string` | Optional | if `type`==`pppoe` |
| `Type` | [`*models.GatewayWanTypeEnum`](../../doc/models/gateway-wan-type-enum.md) | Optional | **Default**: `"dhcp"` |

## Example (as JSON)

```json
{
  "pppoe_auth": "none",
  "type": "dhcp",
  "dns": [
    "dns3",
    "dns2",
    "dns1"
  ],
  "dns_suffix": [
    "dns_suffix7",
    "dns_suffix6"
  ],
  "gateway": "gateway4",
  "ip": "ip8",
  "netmask": "netmask4"
}
```

