
# Gateway Template Tunnel Node

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
  ]
}
```

