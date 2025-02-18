
# Capture Mxedge Mxedges

Property key is the Mx Edge ID, currently limited to one mxedge per org capture session

*This model accepts additional fields of type interface{}.*

## Structure

`CaptureMxedgeMxedges`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | [`map[string]models.CaptureMxedgeMxedgesInterfaces`](../../doc/models/capture-mxedge-mxedges-interfaces.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "interfaces": {
    "key0": {
      "tcpdump_expression": "tcpdump_expression4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "tcpdump_expression": "tcpdump_expression4",
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

