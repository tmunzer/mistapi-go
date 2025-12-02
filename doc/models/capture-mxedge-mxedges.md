
# Capture Mxedge Mxedges

Property key is the Mx Edge ID, currently limited to one mxedge per org capture session

## Structure

`CaptureMxedgeMxedges`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | [`map[string]models.CaptureMxedgeMxedgesInterfaces`](../../doc/models/capture-mxedge-mxedges-interfaces.md) | Optional | - |

## Example (as JSON)

```json
{
  "interfaces": {
    "key0": {
      "tcpdump_expression": "tcpdump_expression4"
    },
    "key1": {
      "tcpdump_expression": "tcpdump_expression4"
    }
  }
}
```

