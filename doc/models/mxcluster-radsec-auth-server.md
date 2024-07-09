
# Mxcluster Radsec Auth Server

## Structure

`MxclusterRadsecAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | ip / hostname of RADIUS server |
| `KeywrapEnabled` | `*bool` | Optional | if used for Mist APs, enable keywrap algorithm. Default is false |
| `KeywrapFormat` | [`models.Optional[models.MxclusterRadAuthServerKeywrapFormatEnum]`](../../doc/models/mxcluster-rad-auth-server-keywrap-format-enum.md) | Optional | if used for Mist APs<br>**Default**: `"ascii"` |
| `KeywrapKek` | `*string` | Optional | if used for Mist APs, encryption key |
| `KeywrapMack` | `*string` | Optional | if used for Mist APs, Message Authentication Code Key |
| `Port` | `*int` | Optional | Auth port of RADIUS server<br>**Default**: `1812` |
| `Secret` | `*string` | Optional | secret of RADIUS server |
| `Ssids` | `[]string` | Optional | list of ssids that will use this server if match_ssid is true and match is found |

## Example (as JSON)

```json
{
  "keywrap_format": "ascii",
  "port": 1812,
  "host": "host0",
  "keywrap_enabled": false,
  "keywrap_kek": "keywrap_kek4",
  "keywrap_mack": "keywrap_mack6"
}
```

