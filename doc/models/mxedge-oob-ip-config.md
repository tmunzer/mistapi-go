
# Mxedge Oob Ip Config

IPconfiguration of the Mist Edge out-of_band management interface

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Autoconf6` | `*bool` | Optional | **Default**: `true` |
| `Dhcp6` | `*bool` | Optional | **Default**: `true` |
| `Dns` | `[]string` | Optional | IPv4 ignored if `type`!=`static`, IPv6 ignored if `type6`!=`static` |
| `Gateway` | `*string` | Optional | If `type`=`static` |
| `Gateway6` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | If `type`=`static` |
| `Ip6` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | If `type`=`static` |
| `Netmask6` | `*string` | Optional | - |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `Type6` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "autoconf6": true,
  "dhcp6": true,
  "dns": [
    "8.8.8.8",
    "4.4.4.4",
    "2001:4860:4860::8888",
    "2001:4860:4860::8844"
  ],
  "gateway": "10.2.1.254",
  "gateway6": "2601:1700:43c0:dc0::1",
  "ip": "10.2.1.2",
  "ip6": "2601:1700:43c0:dc0:20c:29ff:fea7:93bc",
  "netmask": "255.255.255.0",
  "netmask6": "/64",
  "type": "static",
  "type6": "static",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

