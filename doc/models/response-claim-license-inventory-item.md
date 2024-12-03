
# Response Claim License Inventory Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClaimLicenseInventoryItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `string` | Required | - |
| `Magic` | `string` | Required | - |
| `Model` | `string` | Required | - |
| `Serial` | `string` | Required | - |
| `Type` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mac": "mac8",
  "magic": "magic4",
  "model": "model2",
  "serial": "serial4",
  "type": "type6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

