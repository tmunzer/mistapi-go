
# Response Http 429 Error Exception

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseHttp429ErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "detail": "Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

