
# Radius Acct Server

## Structure

`RadiusAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `string` | Required | ip / hostname of RADIUS server |
| `KeywrapEnabled` | `*bool` | Optional | - |
| `KeywrapFormat` | [`*models.RadiusKeywrapFormatEnum`](../../doc/models/radius-keywrap-format-enum.md) | Optional | - |
| `KeywrapKek` | `*string` | Optional | - |
| `KeywrapMack` | `*string` | Optional | - |
| `Port` | `*int` | Optional | Acct port of RADIUS server<br>**Default**: `1813` |
| `Secret` | `string` | Required | secret of RADIUS server |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "keywrap_kek": "1122334455",
  "keywrap_mack": "1122334455",
  "port": 1813,
  "secret": "testing123",
  "keywrap_enabled": false,
  "keywrap_format": "hex"
}
```

