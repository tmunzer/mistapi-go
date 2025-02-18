
# Tunnel Config Node

Only if `provider`==`zscaler-ipsec`, `provider`==`jse-ipsec` or `provider`==`custom-ipsec`

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelConfigNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hosts` | `[]string` | Required | - |
| `InternalIps` | `[]string` | Optional | Only if `provider`==`zscaler-gre`, `provider`==`jse-ipsec`, `provider`==`custom-ipsec` or `provider`==`custom-gre` |
| `ProbeIps` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `RemoteIds` | `[]string` | Optional | Only if  `provider`==`jse-ipsec` or `provider`==`custom-ipsec` |
| `WanNames` | `[]string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "hosts": [
    "hosts7"
  ],
  "internal_ips": [
    "internal_ips4",
    "internal_ips5"
  ],
  "probe_ips": [
    "probe_ips7"
  ],
  "remote_ids": [
    "remote_ids5"
  ],
  "wan_names": [
    "wan_names8",
    "wan_names9"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

