
# Mxcluster Radsec Auth Server

*This model accepts additional fields of type interface{}.*

## Structure

`MxclusterRadsecAuthServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | IP / hostname of RADIUS server |
| `InbandStatusCheck` | `*bool` | Optional | Whether to enable inband status check<br><br>**Default**: `false` |
| `InbandStatusInterval` | `*int` | Optional | Inband status interval, in seconds<br><br>**Default**: `300`<br><br>**Constraints**: `>= 0` |
| `KeywrapEnabled` | `*bool` | Optional | If used for Mist APs, enable keywrap algorithm. Default is false |
| `KeywrapFormat` | [`models.Optional[models.MxclusterRadAuthServerKeywrapFormatEnum]`](../../doc/models/mxcluster-rad-auth-server-keywrap-format-enum.md) | Optional | if used for Mist APs. enum: `ascii`, `hex`<br><br>**Default**: `"ascii"` |
| `KeywrapKek` | `*string` | Optional | If used for Mist APs, encryption key |
| `KeywrapMack` | `*string` | Optional | If used for Mist APs, Message Authentication Code Key |
| `Port` | `*int` | Optional | Auth port of RADIUS server<br><br>**Default**: `1812` |
| `Retry` | `*int` | Optional | Authentication request retry<br><br>**Default**: `2` |
| `Secret` | `*string` | Optional | Secret of RADIUS server |
| `Ssids` | `[]string` | Optional | List of ssids that will use this server if match_ssid is true and match is found |
| `Timeout` | `*int` | Optional | Authentication request timeout, in seconds<br><br>**Default**: `5` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "inband_status_check": false,
  "inband_status_interval": 300,
  "keywrap_format": "ascii",
  "port": 1812,
  "retry": 2,
  "timeout": 5,
  "host": "host6",
  "keywrap_enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

