
# Response Http 400 Exception

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp400Exception`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "detail": "JSON parse error - Expecting value: line 5 column 8 (char 56)",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

