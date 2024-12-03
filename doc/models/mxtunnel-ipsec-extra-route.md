
# Mxtunnel Ipsec Extra Route

*This model accepts additional fields of type interface{}.*

## Structure

`MxtunnelIpsecExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dest` | `*string` | Optional | - |
| `NextHop` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dest": "dest2",
  "next_hop": "next_hop2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

