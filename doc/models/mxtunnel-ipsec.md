
# Mxtunnel Ipsec

*This model accepts additional fields of type interface{}.*

## Structure

`MxtunnelIpsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DnsServers` | `models.Optional[[]string]` | Optional | - |
| `DnsSuffix` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Enabled` | `*bool` | Optional | - |
| `ExtraRoutes` | [`[]models.MxtunnelIpsecExtraRoute`](../../doc/models/mxtunnel-ipsec-extra-route.md) | Optional | - |
| `SplitTunnel` | `*bool` | Optional | - |
| `UseMxedge` | `*bool` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dns_servers": [
    "dns_servers8",
    "dns_servers7"
  ],
  "dns_suffix": [
    "dns_suffix1",
    "dns_suffix2"
  ],
  "enabled": false,
  "extra_routes": [
    {
      "dest": "dest4",
      "next_hop": "next_hop4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "dest": "dest4",
      "next_hop": "next_hop4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "dest": "dest4",
      "next_hop": "next_hop4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "split_tunnel": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

