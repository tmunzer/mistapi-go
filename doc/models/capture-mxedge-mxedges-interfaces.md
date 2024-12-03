
# Capture Mxedge Mxedges Interfaces

Property key is the Port name (e.g. "port1", "kni0", "lacp0", "ipsec", "drop", "oobm"), currently limited to specifying one interface per mxedge

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureMxedgeMxedgesInterfaces`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression common for wired,radiotap |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "tcpdump_expression": "tcpdump_expression8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

