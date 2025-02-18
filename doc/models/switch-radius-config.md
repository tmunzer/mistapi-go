
# Switch Radius Config

Junos Radius config

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchRadiusConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctInterimInterval` | `*int` | Optional | How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 65535` |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | **Constraints**: *Unique Items Required* |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AuthServersRetries` | `*int` | Optional | Radius auth session retries<br>**Default**: `3` |
| `AuthServersTimeout` | `*int` | Optional | Radius auth session timeout<br>**Default**: `5` |
| `Network` | `*string` | Optional | Use `network`or `source_ip`. Which network the RADIUS server resides, if there's static IP for this network, we'd use it as source-ip |
| `SourceIp` | `*string` | Optional | Use `network`or `source_ip` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "acct_interim_interval": 0,
  "auth_servers_retries": 3,
  "auth_servers_timeout": 5,
  "acct_servers": [
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek0",
      "keywrap_mack": "keywrap_mack2",
      "port": 254,
      "secret": "secret0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "auth_servers": [
    {
      "host": "host0",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6",
      "port": 114,
      "secret": "secret4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

