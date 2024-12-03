
# Gateway Template Tunnel Node

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayTemplateTunnelNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Hosts` | `[]string` | Optional | - |
| `InternalIps` | `[]string` | Optional | Only if:<br><br>* `provider`== `zscaler-gre`<br>* `provider`== `custom-gre` |
| `ProbeIps` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `RemoteIds` | `[]string` | Optional | Only if `provider`== `custom-ipsec` |
| `WanNames` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "hosts": [
    "hosts9",
    "hosts0"
  ],
  "internal_ips": [
    "internal_ips0"
  ],
  "probe_ips": [
    "probe_ips3",
    "probe_ips4",
    "probe_ips5"
  ],
  "remote_ids": [
    "remote_ids9",
    "remote_ids0"
  ],
  "wan_names": [
    "wan_names4"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

