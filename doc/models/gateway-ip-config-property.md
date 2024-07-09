
# Gateway Ip Config Property

## Structure

`GatewayIpConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | - |
| `SecondaryIps` | `[]string` | Optional | optional list of secondary IPs in CIDR format |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | **Default**: `"dhcp"` |

## Example (as JSON)

```json
{
  "netmask": "/24",
  "secondary_ips": [
    "192.168.50.1/24",
    "192.168.60.1/26"
  ],
  "type": "static",
  "ip": "ip8"
}
```

