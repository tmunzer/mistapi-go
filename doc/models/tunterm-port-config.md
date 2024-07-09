
# Tunterm Port Config

ethernet port configurations

## Structure

`TuntermPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownstreamPorts` | `[]string` | Optional | list of ports to be used for downstream (to AP) purpose |
| `SeparateUpstreamDownstream` | `*bool` | Optional | weather to separate upstream / downstream ports. default is false where all ports will be used.<br>**Default**: `false` |
| `UpstreamPortVlanId` | `*int` | Optional | native VLAN id for upstream ports<br>**Default**: `1` |
| `UpstreamPorts` | `[]string` | Optional | list of ports to be used for upstrea purpose (to LAN) |

## Example (as JSON)

```json
{
  "downstream_ports": [
    "2",
    "3"
  ],
  "separate_upstream_downstream": false,
  "upstream_port_vlan_id": 30,
  "upstream_ports": [
    "0",
    "1"
  ]
}
```

