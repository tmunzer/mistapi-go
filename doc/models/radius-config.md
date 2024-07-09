
# Radius Config

Junos Radius config

## Structure

`RadiusConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctInterimInterval` | `*int` | Optional | how frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from RADIUS Server). Very frequent messages can affect the performance of the radius server, 600 and up is recommended when enabled<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 65535` |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | **Constraints**: *Unique Items Required* |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AuthServersRetries` | `*int` | Optional | radius auth session retries<br>**Default**: `3` |
| `AuthServersTimeout` | `*int` | Optional | radius auth session timeout<br>**Default**: `5` |
| `CoaEnabled` | `*bool` | Optional | **Default**: `false` |
| `CoaPort` | `*int` | Optional | **Default**: `3799` |
| `Network` | `*string` | Optional | use `network`or `source_ip`<br>which network the RADIUS server resides, if there's static IP for this network, we'd use it as source-ip |
| `SourceIp` | `*string` | Optional | use `network`or `source_ip` |

## Example (as JSON)

```json
{
  "acct_interim_interval": 0,
  "auth_servers_retries": 3,
  "auth_servers_timeout": 5,
  "coa_enabled": false,
  "coa_port": 3799,
  "acct_servers": [
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "hex",
      "keywrap_kek": "keywrap_kek0",
      "keywrap_mack": "keywrap_mack2",
      "port": 254,
      "secret": "secret0"
    },
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "hex",
      "keywrap_kek": "keywrap_kek0",
      "keywrap_mack": "keywrap_mack2",
      "port": 254,
      "secret": "secret0"
    }
  ],
  "auth_servers": [
    {
      "host": "host0",
      "keywrap_enabled": false,
      "keywrap_format": "hex",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6",
      "port": 114,
      "secret": "secret4"
    },
    {
      "host": "host0",
      "keywrap_enabled": false,
      "keywrap_format": "hex",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6",
      "port": 114,
      "secret": "secret4"
    }
  ]
}
```

