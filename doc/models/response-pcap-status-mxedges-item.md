
# Response Pcap Status Mxedges Item

## Structure

`ResponsePcapStatusMxedgesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interfaces` | [`map[string]models.CaptureMxedgeMxedgesInterfaces`](../../doc/models/capture-mxedge-mxedges-interfaces.md) | Optional | Dict of interfaces to capture on, property key is the port name |

## Example (as JSON)

```json
{
  "interfaces": {
    "key0": {
      "tcpdump_expression": "tcpdump_expression4"
    },
    "key1": {
      "tcpdump_expression": "tcpdump_expression4"
    },
    "key2": {
      "tcpdump_expression": "tcpdump_expression4"
    }
  }
}
```

