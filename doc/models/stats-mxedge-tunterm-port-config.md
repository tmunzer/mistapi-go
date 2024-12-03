
# Stats Mxedge Tunterm Port Config

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeTuntermPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownstreamPorts` | `[]string` | Optional | - |
| `SeparateUpstreamDownstream` | `*bool` | Optional | - |
| `UpstreamPorts` | `[]string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

