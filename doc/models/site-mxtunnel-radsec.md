
# Site Mxtunnel Radsec

## Structure

`SiteMxtunnelRadsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServers` | [`[]models.RadiusAcctServer`](../../doc/models/radius-acct-server.md) | Optional | - |
| `AuthServers` | [`[]models.RadiusAuthServer`](../../doc/models/radius-auth-server.md) | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `UseMxedge` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "acct_servers": [
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek0",
      "keywrap_mack": "keywrap_mack2",
      "port": 254,
      "secret": "secret0"
    },
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek0",
      "keywrap_mack": "keywrap_mack2",
      "port": 254,
      "secret": "secret0"
    },
    {
      "host": "host4",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
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
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6",
      "port": 114,
      "secret": "secret4"
    },
    {
      "host": "host0",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6",
      "port": 114,
      "secret": "secret4"
    },
    {
      "host": "host0",
      "keywrap_enabled": false,
      "keywrap_format": "ascii",
      "keywrap_kek": "keywrap_kek4",
      "keywrap_mack": "keywrap_mack6",
      "port": 114,
      "secret": "secret4"
    }
  ],
  "use_mxedge": false
}
```

