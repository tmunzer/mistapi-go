
# Vrf Extra Route

*This model accepts additional fields of type interface{}.*

## Structure

`VrfExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Via` | `*string` | Optional | Next-hop address |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "via": "via8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

