
# Stats Mxedge Lag Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeLagStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ActivePorts` | `[]string` | Optional | List of ports active on the LAG defined by the LACP |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "active_ports": [
    "active_ports9",
    "active_ports8",
    "active_ports7"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

