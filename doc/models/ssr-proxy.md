
# Ssr Proxy

SSR proxy configuration to talk to Mist

*This model accepts additional fields of type interface{}.*

## Structure

`SsrProxy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `Url` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": true,
  "url": "https://proxy.corp.com:8080/",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

