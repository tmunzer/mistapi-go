
# Ap Ip Config

IP AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | If `type`==`static` |
| `DnsSuffix` | `[]string` | Optional | Required if `type`==`static` |
| `Gateway` | `*string` | Optional | Required if `type`==`static` |
| `Gateway6` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | Required if `type`==`static` |
| `Ip6` | `*string` | Optional | - |
| `Mtu` | `*int` | Optional | - |
| `Netmask` | `*string` | Optional | Required if `type`==`static` |
| `Netmask6` | `*string` | Optional | - |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpType6Enum`](../../doc/models/ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"disabled"` |
| `VlanId` | `*int` | Optional | Management VLAN id, default is 1 (untagged)<br><br>**Default**: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "vlan_id": 1,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

