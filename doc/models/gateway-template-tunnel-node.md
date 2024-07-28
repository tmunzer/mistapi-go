
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
    "hosts1",
    "hosts2"
  ],
  "internal_ips": [
    "internal_ips2",
    "internal_ips3",
    "internal_ips4"
  ],
  "probe_ips": [
    "probe_ips5",
    "probe_ips6"
  ],
  "remote_ids": [
    "remote_ids9",
    "remote_ids8"
  ],
  "wan_names": [
    "wan_names6",
    "wan_names7",
    "wan_names8"
  ]
}
```

