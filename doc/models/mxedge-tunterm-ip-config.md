
# Mxedge Tunterm Ip Config

IPconfiguration of the Mist Tunnel interface

## Structure

`MxedgeTuntermIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `string` | Required | - |
| `Gateway6` | `*string` | Optional | - |
| `Ip` | `string` | Required | Untagged VLAN |
| `Ip6` | `*string` | Optional | - |
| `Netmask` | `string` | Required | - |
| `Netmask6` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "gateway": "10.2.1.254",
  "gateway6": "2001:1010:1010:1010::1",
  "ip": "10.2.1.1",
  "ip6": "2001:1010:1010:1010::2",
  "netmask": "255.255.255.0",
  "netmask6": "/64"
}
```

