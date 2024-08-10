
# Stats Mxedge Tunterm Port Config

## Structure

`StatsMxedgeTuntermPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownstreamPorts` | `[]string` | Optional | - |
| `SeparateUpstreamDownstream` | `*bool` | Optional | - |
| `UpstreamPorts` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "downstream_ports": [
    "0",
    "1"
  ],
  "separate_upstream_downstream": false,
  "upstream_ports": [
    "0",
    "1"
  ]
}
```

