
# Sso Mxedge Proxy Auth Server

## Structure

`SsoMxedgeProxyAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | - |
| `Port` | `*int` | Optional | **Default**: `1812` |
| `RequireMessageAuthenticator` | `*bool` | Optional | whether to require Message-Authenticator in requests<br>**Default**: `false` |
| `Secret` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "port": 1812,
  "require_message_authenticator": false,
  "secret": "testing123"
}
```

