
# Setting Ssr

*This model accepts additional fields of type interface{}.*

## Structure

`SettingSsr`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.SettingSsrAutoUpgrade`](../../doc/models/setting-ssr-auto-upgrade.md) | Optional | auto_upgrade device first time it is onboarded |
| `ConductorHosts` | `[]string` | Optional | List of Conductor IP Addresses or Hosts to be used by the SSR Devices |
| `ConductorToken` | `*string` | Optional | Token to be used by the SSR Devices to connect to the Conductor |
| `DisableStats` | `*bool` | Optional | Disable stats collection on SSR devices |
| `Proxy` | [`*models.Proxy`](../../doc/models/proxy.md) | Optional | Proxy Configuration to talk to Mist |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auto_upgrade": {
    "channel": "beta",
    "custom_versions": {
      "key0": "custom_versions3",
      "key1": "custom_versions2"
    },
    "enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "conductor_hosts": [
    "conductor_hosts4",
    "conductor_hosts5"
  ],
  "conductor_token": "conductor_token6",
  "disable_stats": false,
  "proxy": {
    "url": "url6",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

