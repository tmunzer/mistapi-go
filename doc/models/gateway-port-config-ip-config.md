
# Gateway Port Config Ip Config

Junos IP Config

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayPortConfigIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | Except for out-of_band interface (vme/em0/fxp0) |
| `DnsSuffix` | `[]string` | Optional | Except for out-of_band interface (vme/em0/fxp0) |
| `Gateway` | `*string` | Optional | Except for out-of_band interface (vme/em0/fxp0). Interface Default Gateway IP Address (i.e. "192.168.1.1") or a Variable (i.e. "{{myvar}}") |
| `Ip` | `*string` | Optional | Interface IP Address (i.e. "192.168.1.8") or a Variable (i.e. "{{myvar}}") |
| `Netmask` | `*string` | Optional | Used only if `subnet` is not specified in `networks`. Interface Netmask (i.e. "/24") or a Variable (i.e. "{{myvar}}") |
| `Network` | `*string` | Optional | Optional, the network to be used for mgmt |
| `PoserPassword` | `*string` | Optional | If `type`==`pppoe` |
| `PppoeAuth` | [`*models.GatewayWanPpoeAuthEnum`](../../doc/models/gateway-wan-ppoe-auth-enum.md) | Optional | if `type`==`pppoe`. enum: `chap`, `none`, `pap`<br><br>**Default**: `"none"` |
| `PppoeUsername` | `*string` | Optional | If `type`==`pppoe` |
| `Type` | [`*models.GatewayWanTypeEnum`](../../doc/models/gateway-wan-type-enum.md) | Optional | enum: `dhcp`, `pppoe`, `static`<br><br>**Default**: `"dhcp"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "gateway": "192.168.1.1",
  "ip": "192.168.1.8",
  "netmask": "/24",
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

