
# Skyatp List Domain

*This model accepts additional fields of type interface{}.*

## Structure

`SkyatpListDomain`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Comment` | `*string` | Optional | - |
| `Value` | `string` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "comment": "restricted",
  "value": "unsafe.com",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

