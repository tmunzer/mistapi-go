
# Mxcluster Radsec Acct Server

## Structure

`MxclusterRadsecAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | ip / hostname of RADIUS server |
| `Port` | `*int` | Optional | Acct port of RADIUS server<br>**Default**: `1813` |
| `Secret` | `*string` | Optional | secret of RADIUS server |
| `Ssids` | `[]string` | Optional | list of ssids that will use this server if match_ssid is true and match is found |

## Example (as JSON)

```json
{
  "port": 1813,
  "host": "host0",
  "secret": "secret6",
  "ssids": [
    "ssids9",
    "ssids8",
    "ssids7"
  ]
}
```

