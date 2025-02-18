
# Dhcpd Config

*This model accepts additional fields of type [models.DhcpdConfigProperty](../../doc/models/dhcpd-config-property.md).*

## Structure

`DhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | If set to `false`, disable the DHCP server<br>**Default**: `true` |
| `AdditionalProperties` | [`map[string]models.DhcpdConfigProperty`](../../doc/models/dhcpd-config-property.md) | Optional | - |

## Example (as JSON)

```json
{
  "enabled": true,
  "exampleAdditionalProperty": {
    "dns_servers": [
      "dns_servers2",
      "dns_servers3",
      "dns_servers4"
    ],
    "dns_suffix": [
      "dns_suffix1",
      "dns_suffix0",
      "dns_suffix9"
    ],
    "fixed_bindings": {
      "key0": {
        "ip": "ip0",
        "name": "name6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    },
    "gateway": "gateway8",
    "ip_end": "ip_end4",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

