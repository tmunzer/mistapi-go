
# Network Source Nat

If `routed`==`false` (usually at Spoke), but some hosts needs to be reachable from Hub

*This model accepts additional fields of type interface{}.*

## Structure

`NetworkSourceNat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExternalIp` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "external_ip": "172.16.0.8/30",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

