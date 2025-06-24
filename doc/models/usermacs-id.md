
# Usermacs Id

*This model accepts additional fields of type interface{}.*

## Structure

`UsermacsId`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `UsermacIds` | `[]uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "usermac_ids": [
    "000017db-0000-0000-0000-000000000000",
    "000017da-0000-0000-0000-000000000000"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

