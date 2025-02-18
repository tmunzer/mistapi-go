
# Response Logout

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseLogout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ForwardUrl` | `*string` | Optional | If configured in SSO as custom_logout_url |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "forward_url": "forward_url0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

