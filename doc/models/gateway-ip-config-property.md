
# Gateway Ip Config Property

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayIpConfigProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | - |
| `SecondaryIps` | `[]string` | Optional | Optional list of secondary IPs in CIDR format |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "netmask": "/24",
  "secondary_ips": [
    "192.168.50.1/24",
    "192.168.60.1/26"
  ],
  "type": "static",
  "ip": "ip6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

