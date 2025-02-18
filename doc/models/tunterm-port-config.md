
# Tunterm Port Config

Ethernet port configurations

*This model accepts additional fields of type interface{}.*

## Structure

`TuntermPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownstreamPorts` | `[]string` | Optional | List of ports to be used for downstream (to AP) purpose |
| `SeparateUpstreamDownstream` | `*bool` | Optional | Whether to separate upstream / downstream ports. default is false where all ports will be used.<br>**Default**: `false` |
| `UpstreamPortVlanId` | `*int` | Optional | Native VLAN id for upstream ports<br>**Default**: `1` |
| `UpstreamPorts` | `[]string` | Optional | List of ports to be used for upstrea purpose (to LAN) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

