
# Mxedge Tunterm Igmp Snooping Querier

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeTuntermIgmpSnoopingQuerier`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaxResponseTime` | `*int` | Optional | querier’s query response interval, in tenths-of-seconds |
| `Mtu` | `*int` | Optional | the MTU we use (needed when forming large IGMPv3 Reports) |
| `QueryInterval` | `*int` | Optional | querier’s query interval, in seconds |
| `Robustness` | `*int` | Optional | querier’s robustness<br>**Constraints**: `>= 1`, `<= 7` |
| `Version` | `*int` | Optional | querier’s maximum protocol version |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "max_response_time": 10,
  "mtu": 1500,
  "query_interval": 125,
  "version": 3,
  "robustness": 150,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

