
# Proxy

Proxy Configuration to talk to Mist

*This model accepts additional fields of type interface{}.*

## Structure

`Proxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "url": "http://proxy.corp.com:8080/",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

