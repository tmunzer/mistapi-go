
# Capture Switch Switches

## Structure

`CaptureSwitchSwitches`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.CaptureSwitchPortsTcpdumpExpression`](../../doc/models/capture-switch-ports-tcpdump-expression.md) | Optional | Property key is the port name. 6 ports max per switch supported, or 5 max with irb port auto-included into capture request |

## Example (as JSON)

```json
{
  "ports": {
    "key0": {
      "tcpdump_expression": "tcpdump_expression0"
    },
    "key1": {
      "tcpdump_expression": "tcpdump_expression0"
    },
    "key2": {
      "tcpdump_expression": "tcpdump_expression0"
    }
  }
}
```

