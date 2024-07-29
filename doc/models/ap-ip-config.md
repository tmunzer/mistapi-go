
# Ap Ip Config

IP AP settings

## Structure

`ApIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | if `type`==`static` |
| `DnsSuffix` | `[]string` | Optional | required if `type`==`static` |
| `Gateway` | `*string` | Optional | required if `type`==`static` |
| `Gateway6` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | required if `type`==`static` |
| `Ip6` | `*string` | Optional | - |
| `Mtu` | `*int` | Optional | - |
| `Netmask` | `*string` | Optional | required if `type`==`static` |
| `Netmask6` | `*string` | Optional | - |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpType6Enum`](../../doc/models/ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br>**Default**: `"disabled"` |
| `VlanId` | `*int` | Optional | management vlan id, default is 1 (untagged)<br>**Default**: `1` |

## Example (as JSON)

```json
{
  "dns": [
    "8.8.8.8",
    "4.4.4.4"
  ],
  "dns_suffix": [
    ".mist.local",
    ".mist.com"
  ],
  "gateway": "10.2.1.254",
  "gateway6": "2607:f8b0:4005:808::1",
  "ip": "10.2.1.1",
  "ip6": "2607:f8b0:4005:808::2004",
  "mtu": 0,
  "netmask": "255.255.255.0",
  "netmask6": "/32",
  "type": "static",
  "type6": "static",
  "vlan_id": 1
}
```

