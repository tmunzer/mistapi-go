
# Network Multicast Group

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkMulticastGroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RpIp` | `*string` | Optional | RP (rendezvous point) IP Address |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "rp_ip": "rp_ip4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

