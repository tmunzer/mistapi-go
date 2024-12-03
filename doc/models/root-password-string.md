
# Root Password String

*This model accepts additional fields of type interface{}.*

## Structure

`RootPasswordString`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RootPassword` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "root_password": "root_password6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

