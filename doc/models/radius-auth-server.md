
# Radius Auth Server

Authentication Server

## Structure

`RadiusAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `string` | Required | ip / hostname of RADIUS server |
| `KeywrapEnabled` | `*bool` | Optional | - |
| `KeywrapFormat` | [`*models.RadiusKeywrapFormatEnum`](../../doc/models/radius-keywrap-format-enum.md) | Optional | enum: `ascii`, `hex` |
| `KeywrapKek` | `*string` | Optional | - |
| `KeywrapMack` | `*string` | Optional | - |
| `Port` | `*int` | Optional | Auth port of RADIUS server<br>**Default**: `1812`<br>**Constraints**: `>= 1`, `<= 65535` |
| `Secret` | `string` | Required | secret of RADIUS server |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "keywrap_kek": "1122334455",
  "keywrap_mack": "1122334455",
  "port": 1812,
  "secret": "testing123",
  "keywrap_enabled": false,
  "keywrap_format": "ascii"
}
```

