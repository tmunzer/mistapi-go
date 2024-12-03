
# Capture Switch Switches

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureSwitchSwitches`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.CaptureSwitchPortsTcpdumpExpression`](../../doc/models/capture-switch-ports-tcpdump-expression.md) | Optional | Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ports": {
    "key0": {
      "tcpdump_expression": "tcpdump_expression0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "tcpdump_expression": "tcpdump_expression0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key2": {
      "tcpdump_expression": "tcpdump_expression0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

