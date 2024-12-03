
# Error Delete Failed Exception

*This model accepts additional fields of type interface{}.*

## Structure

`ErrorDeleteFailedException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `string` | Required | - |
| `OrgId` | `uuid.UUID` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "detail": "inventory not empty",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

