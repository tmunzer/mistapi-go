
# Switch Radius

By default, `radius_config` will be used. if a different one has to be used set `use_different_radius

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchRadius`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `RadiusConfig` | [`*models.SwitchRadiusConfig`](../../doc/models/switch-radius-config.md) | Optional | Junos Radius config |
| `UseDifferentRadius` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "radius_config": {
    "acct_immediate_update": false,
    "acct_interim_interval": 118,
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
      },
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
    "auth_server_selection": "ordered",
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
      },
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
      },
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
  },
  "use_different_radius": "use_different_radius4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

