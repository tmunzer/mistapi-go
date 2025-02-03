
# Sso Mxedge Proxy Auth Server

*This model accepts additional fields of type interface{}.*

## Structure

`SsoMxedgeProxyAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Default**: `1812` |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br>**Default**: `false` |
| `Secret` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "port": 1812,
  "require_message_authenticator": false,
  "secret": "testing123",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

