
# Junos Ip Config

Junos IP Config

*This model accepts additional fields of type interface{}.*

## Structure

`JunosIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dns` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `DnsSuffix` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Gateway` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `Netmask` | `*string` | Optional | Used only if `subnet` is not specified in `networks` |
| `Network` | `*string` | Optional | Network where this mgmt IP reside, this will be used as default network for outbound-ssh, dns, ntp, dns, tacplus, radius, syslog, snmp |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | enum: `dhcp`, `static`<br>**Default**: `"dhcp"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "type": "static",
  "dns": [
    "dns5"
  ],
  "dns_suffix": [
    "dns_suffix1"
  ],
  "gateway": "gateway6",
  "ip": "ip0",
  "netmask": "netmask6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

