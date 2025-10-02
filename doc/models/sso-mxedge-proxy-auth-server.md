
# Sso Mxedge Proxy Auth Server

*This model accepts additional fields of type interface{}.*

## Structure

`SsoMxedgeProxyAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Default**: `1812` |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br><br>**Default**: `false` |
| `Retry` | `*int` | Optional | Authentication request retry<br><br>**Default**: `2` |
| `Secret` | `*string` | Optional | - |
| `Timeout` | `*int` | Optional | Authentication request timeout, in seconds<br><br>**Default**: `5` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "port": 1812,
  "require_message_authenticator": false,
  "retry": 2,
  "secret": "testing123",
  "timeout": 5,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

