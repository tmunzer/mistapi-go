
# Config Switch Radius

by default, `radius_config` will be used. if a different one has to be used set `use_different_radius

## Structure

`ConfigSwitchRadius`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `RadiusConfig` | [`*models.RadiusConfig`](../../doc/models/radius-config.md) | Optional | Junos Radius config |
| `UseDifferentRadius` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "radius_config": {
    "acct_interim_interval": 118,
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
    "auth_servers_retries": 194,
    "auth_servers_timeout": 232
  },
  "use_different_radius": "use_different_radius4"
}
```

