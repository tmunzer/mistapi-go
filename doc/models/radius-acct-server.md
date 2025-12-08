
# Radius Acct Server

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
| `Port` | [`*models.RadiusAcctPort`](../../doc/models/containers/radius-acct-port.md) | Optional | Radius Auth Port, value from 1 to 65535, default is 1813 |
| `Secret` | `string` | Required | Secret of RADIUS server |

## Example (as JSON)

```json
{
  "host": "1.2.3.4",
  "keywrap_kek": "1122334455",
  "keywrap_mack": "1122334455",
  "secret": "testing123",
  "keywrap_enabled": false,
  "keywrap_format": "ascii",
  "port": 232
}
```

