
# Switch Dhcpd Config

*This model accepts additional fields of type [models.SwitchDhcpdConfigProperty](../../doc/models/switch-dhcpd-config-property.md).*

## Structure

`SwitchDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | if set to `true`, enable the DHCP server<br>**Default**: `false` |
| `AdditionalProperties` | [`map[string]models.SwitchDhcpdConfigProperty`](../../doc/models/switch-dhcpd-config-property.md) | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "exampleAdditionalProperty": {
    "dns_servers": [
      "dns_servers8",
      "dns_servers9"
    ],
    "dns_suffix": [
      "dns_suffix5",
      "dns_suffix6"
    ],
    "fixed_bindings": {
      "key0": {
        "ip": "ip0",
        "name": "name6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      "key1": {
        "ip": "ip0",
        "name": "name6",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    },
    "gateway": "gateway2",
    "ip_end": "ip_end2",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

