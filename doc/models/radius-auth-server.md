
# Radius Auth Server

Authentication Server

*This model accepts additional fields of type interface{}.*

## Structure

`RadiusAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `string` | Required | IP/ hostname of RADIUS server |
| `KeywrapEnabled` | `*bool` | Optional | - |
| `KeywrapFormat` | [`*models.RadiusKeywrapFormatEnum`](../../doc/models/radius-keywrap-format-enum.md) | Optional | enum: `ascii`, `hex` |
| `KeywrapKek` | `*string` | Optional | - |
| `KeywrapMack` | `*string` | Optional | - |
| `Port` | [`*models.RadiusAuthPort`](../../doc/models/containers/radius-auth-port.md) | Optional | Radius Auth Port, value from 1 to 65535, default is 1812 |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br><br>**Default**: `false` |
| `Secret` | `string` | Required | Secret of RADIUS server |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "keywrap_kek": "1122334455",
  "keywrap_mack": "1122334455",
  "require_message_authenticator": false,
  "secret": "testing123",
  "keywrap_enabled": false,
  "keywrap_format": "ascii",
  "port": 150,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

