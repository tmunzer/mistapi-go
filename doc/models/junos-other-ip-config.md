
# Junos Other Ip Config

optional, if it's required to have switch's L3 presense on a network/vlan

## Structure

`JunosOtherIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EvpnAnycast` | `*bool` | Optional | for EVPN, if anycast is desired<br>**Default**: `false` |
| `Ip` | `*string` | Optional | required if `type`==`static` |
| `Ip6` | `*string` | Optional | required if `type6`==`static` |
| `Netmask` | `*string` | Optional | optional, `subnet` from `network` definition will be used if defined |
| `Netmask6` | `*string` | Optional | optional, `subnet` from `network` definition will be used if defined |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpType6Enum`](../../doc/models/ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br>**Default**: `"disabled"` |

## Example (as JSON)

```json
{
  "evpn_anycast": false,
  "ip": "10.3.3.1",
  "ip6": "fdad:b0bc:f29e::3d16",
  "netmask": "255.255.255.0",
  "netmask6": "/64",
  "type": "static",
  "type6": "static"
}
```

