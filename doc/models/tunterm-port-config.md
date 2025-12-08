
# Tunterm Port Config

Ethernet port configurations

## Structure

`TuntermPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownstreamPorts` | `[]string` | Optional | List of ports to be used for downstream (to AP) purpose |
| `SeparateUpstreamDownstream` | `*bool` | Optional | Whether to separate upstream / downstream ports. default is false where all ports will be used.<br><br>**Default**: `false` |
| `UpstreamPortVlanId` | [`*models.TuntermPortConfigUpstreamPortVlanId`](../../doc/models/containers/tunterm-port-config-upstream-port-vlan-id.md) | Optional | Native VLAN id for upstream ports |
| `UpstreamPorts` | `[]string` | Optional | List of ports to be used for upstream purpose (to LAN) |

## Example (as JSON)

```json
{
  "downstream_ports": [
    "2",
    "3"
  ],
  "separate_upstream_downstream": false,
  "upstream_ports": [
    "0",
    "1"
  ],
  "upstream_port_vlan_id": "String1"
}
```

