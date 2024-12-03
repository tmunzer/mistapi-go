
# Inventory Search Result Member

*This model accepts additional fields of type interface{}.*

## Structure

`InventorySearchResultMember`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | - |
| `Model` | `*string` | Optional | - |
| `Serial` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "f01c2df166e0",
  "model": "EX4300-48P",
  "serial": "PD3714460200",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

