
# Response Pcap Ap

*This model accepts additional fields of type interface{}.*

## Structure

`ResponsePcapAp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | `*int` | Optional | - |
| `Bandwidth` | `*int` | Optional | - |
| `Channel` | `*int` | Optional | - |
| `TcpdumpExpresssion` | `models.Optional[string]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "band": 104,
  "bandwidth": 38,
  "channel": 190,
  "tcpdump_expresssion": "tcpdump_expresssion8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

