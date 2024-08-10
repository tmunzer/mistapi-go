
# Capture Mxedge Mxedges Interfaces

Property key is the Port name (e.g. "port1", "kni0", "lacp0", "ipsec", "drop", "oobm"), currently limited to specifying one interface per mxedge

## Structure

`CaptureMxedgeMxedgesInterfaces`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression common for wired,radiotap |

## Example (as JSON)

```json
{
  "tcpdump_expression": "tcpdump_expression8"
}
```

