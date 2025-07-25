
# Switch Dhcpd Config

*This model accepts additional fields of type [models.SwitchDhcpdConfigProperty](../../doc/models/switch-dhcpd-config-property.md).*

## Structure

`SwitchDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | If set to `true`, enable the DHCP server<br><br>**Default**: `false` |
| `AdditionalProperties` | [`map[string]models.SwitchDhcpdConfigProperty`](../../doc/models/switch-dhcpd-config-property.md) | Optional | the Property key is the network name. In case of DHCP relay, it's common for many networks to use the same dhcp relay, comma-separated network names can be used here (e.g. "net1,net2") |

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
      "key0": null,
      "key1": {}
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

