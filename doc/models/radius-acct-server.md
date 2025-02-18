
# Radius Acct Server

*This model accepts additional fields of type interface{}.*

## Structure

`RadiusAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `string` | Required | IP/ hostname of RADIUS server |
| `KeywrapEnabled` | `*bool` | Optional | - |
| `KeywrapFormat` | [`*models.RadiusKeywrapFormatEnum`](../../doc/models/radius-keywrap-format-enum.md) | Optional | enum: `ascii`, `hex` |
| `KeywrapKek` | `*string` | Optional | - |
| `KeywrapMack` | `*string` | Optional | - |
| `Port` | `*int` | Optional | Acct port of RADIUS server<br>**Default**: `1813`<br>**Constraints**: `>= 1`, `<= 65535` |
| `Secret` | `string` | Required | Secretof RADIUS server |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "keywrap_kek": "1122334455",
  "keywrap_mack": "1122334455",
  "port": 1813,
  "secret": "testing123",
  "keywrap_enabled": false,
  "keywrap_format": "ascii",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

