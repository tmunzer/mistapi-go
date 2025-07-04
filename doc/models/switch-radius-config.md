
# Switch Radius Config

Junos Radius config

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchRadiusConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctImmediateUpdate` | `*bool` | Optional | - |
| `AcctInterimInterval` | `*int` | Optional | How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | **Constraints**: *Unique Items Required* |
| `AuthServerSelection` | [`*models.SwitchRadiusConfigAuthServerSelectionEnum`](../../doc/models/switch-radius-config-auth-server-selection-enum.md) | Optional | enum: `ordered`, `unordered`<br><br>**Default**: `"ordered"` |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AuthServersRetries` | `*int` | Optional | Radius auth session retries<br><br>**Default**: `3` |
| `AuthServersTimeout` | `*int` | Optional | Radius auth session timeout<br><br>**Default**: `5` |
| `CoaEnabled` | `*bool` | Optional | **Default**: `false` |
| `CoaPort` | [`*models.RadiusCoaPort`](../../doc/models/containers/radius-coa-port.md) | Optional | Radius CoA Port, value from 1 to 65535, default is 3799 |
| `FastDot1xTimers` | `*bool` | Optional | **Default**: `false` |
| `Network` | `*string` | Optional | Use `network`or `source_ip`. Which network the RADIUS server resides, if there's static IP for this network, we'd use it as source-ip |
| `SourceIp` | `*string` | Optional | Use `network`or `source_ip` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "acct_interim_interval": 0,
  "auth_server_selection": "ordered",
  "auth_servers_retries": 3,
  "auth_servers_timeout": 5,
  "coa_enabled": false,
  "fast_dot1x_timers": false,
  "acct_immediate_update": false,
  "acct_servers": [
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek0",
      "keywrap_mack": "keywrap_mack2",
      "port": 176,
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
      "port": 36,
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

