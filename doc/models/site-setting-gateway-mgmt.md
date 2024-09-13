
# Site Setting Gateway Mgmt

Gateway Site settings

## Structure

`SiteSettingGatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminSshkeys` | `[]string` | Optional | for SSR only, as direct root access is not allowed |
| `AppProbing` | [`*models.AppProbing`](../../doc/models/app-probing.md) | Optional | - |
| `AppUsage` | `*bool` | Optional | consumes uplink bandwidth, requires WA license |
| `AutoSignatureUpdate` | [`*models.SiteSettingGatewayMgmtAutoSignatureUpdate`](../../doc/models/site-setting-gateway-mgmt-auto-signature-update.md) | Optional | - |
| `ConfigRevertTimer` | `*int` | Optional | he rollback timer for commit confirmed<br>**Default**: `10`<br>**Constraints**: `>= 1`, `<= 30` |
| `DisableConsole` | `*bool` | Optional | for both SSR and SRX disable console port<br>**Default**: `false` |
| `DisableOob` | `*bool` | Optional | for both SSR and SRX disable management interface<br>**Default**: `false` |
| `ProbeHosts` | `[]string` | Optional | - |
| `RootPassword` | `*string` | Optional | for SRX only |
| `SecurityLogSourceAddress` | `*string` | Optional | - |
| `SecurityLogSourceInterface` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "admin_sshkeys": [
    "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAA...Wxa6p6UW0ZbcP john@host"
  ],
  "config_revert_timer": 10,
  "disable_console": false,
  "disable_oob": false,
  "probe_hosts": [
    "8.8.8.8"
  ],
  "security_log_source_address": "192.168.1.1",
  "security_log_source_interface": "ge-0/0/1.0",
  "app_probing": {
    "apps": [
      "apps8",
      "apps9",
      "apps0"
    ],
    "custom_apps": [
      {
        "address": "address4",
        "app_type": "app_type2",
        "hostnames": [
          "hostnames4",
          "hostnames5"
        ],
        "key": "key8",
        "name": "name8"
      },
      {
        "address": "address4",
        "app_type": "app_type2",
        "hostnames": [
          "hostnames4",
          "hostnames5"
        ],
        "key": "key8",
        "name": "name8"
      }
    ],
    "enabled": false
  },
  "app_usage": false,
  "auto_signature_update": {
    "day_of_week": "any",
    "enable": false,
    "time_of_day": "time_of_day6"
  }
}
```

