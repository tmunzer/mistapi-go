
# Setting Ssr

## Structure

`SettingSsr`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.SettingSsrAutoUpgrade`](../../doc/models/setting-ssr-auto-upgrade.md) | Optional | auto_upgrade device first time it is onboarded |
| `ConductorHosts` | `[]string` | Optional | List of Conductor IP Addresses or Hosts to be used by the SSR Devices |
| `ConductorToken` | `*string` | Optional | Token to be used by the SSR Devices to connect to the Conductor |
| `DisableStats` | `*bool` | Optional | Disable stats collection on SSR devices |
| `Proxy` | [`*models.SsrProxy`](../../doc/models/ssr-proxy.md) | Optional | SSR proxy configuration to talk to Mist |

## Example (as JSON)

```json
{
  "auto_upgrade": {
    "channel": "beta",
    "custom_versions": {
      "key0": "custom_versions3",
      "key1": "custom_versions2"
    },
    "enabled": false
  },
  "conductor_hosts": [
    "conductor_hosts4",
    "conductor_hosts5"
  ],
  "conductor_token": "conductor_token6",
  "disable_stats": false,
  "proxy": {
    "disabled": false,
    "url": "url6"
  }
}
```

