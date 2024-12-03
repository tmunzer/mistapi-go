
# Response Pcap Search Item Pcap Aps Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePcapSearchItemPcapApsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*string` | Optional | - |
| `Bandwidth` | `*string` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `TcpdumpExpression` | `models.Optional[string]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": "band4",
  "bandwidth": "bandwidth6",
  "channel": 16,
  "tcpdump_expression": "tcpdump_expression2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

