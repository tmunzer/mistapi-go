
# Mxcluster Radsec Acct Server

*This model accepts additional fields of type interface{}.*

## Structure

`MxclusterRadsecAcctServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | ip / hostname of RADIUS server |
| `Port` | `*int` | Optional | Acct port of RADIUS server<br>**Default**: `1813` |
| `Secret` | `*string` | Optional | secret of RADIUS server |
| `Ssids` | `[]string` | Optional | list of ssids that will use this server if match_ssid is true and match is found |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "port": 1813,
  "host": "host8",
  "secret": "secret4",
  "ssids": [
    "ssids9",
    "ssids0",
    "ssids1"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

