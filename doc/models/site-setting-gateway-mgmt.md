
# Site Setting Gateway Mgmt

Gateway Site settings

*This model accepts additional fields of type interface{}.*

## Structure

`SiteSettingGatewayMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminSshkeys` | `[]string` | Optional | For SSR only, as direct root access is not allowed |
| `AppProbing` | [`*models.AppProbing`](../../doc/models/app-probing.md) | Optional | - |
| `AppUsage` | `*bool` | Optional | Consumes uplink bandwidth, requires WA license |
| `AutoSignatureUpdate` | [`*models.SiteSettingGatewayMgmtAutoSignatureUpdate`](../../doc/models/site-setting-gateway-mgmt-auto-signature-update.md) | Optional | - |
| `ConfigRevertTimer` | `*int` | Optional | Rollback timer for commit confirmed<br>**Default**: `10`<br>**Constraints**: `>= 1`, `<= 30` |
| `DisableConsole` | `*bool` | Optional | For both SSR and SRX disable console port<br>**Default**: `false` |
| `DisableOob` | `*bool` | Optional | For both SSR and SRX disable management interface<br>**Default**: `false` |
| `ProbeHosts` | `[]string` | Optional | - |
| `ProtectRe` | [`*models.ProtectRe`](../../doc/models/protect-re.md) | Optional | Restrict inbound-traffic to host<br>when enabled, all traffic that is not essential to our operation will be dropped<br>e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works |
| `RootPassword` | `*string` | Optional | For SRX only |
| `SecurityLogSourceAddress` | `*string` | Optional | - |
| `SecurityLogSourceInterface` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
        "name": "name8",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "address": "address4",
        "app_type": "app_type2",
        "hostnames": [
          "hostnames4",
          "hostnames5"
        ],
        "key": "key8",
        "name": "name8",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
    "enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "app_usage": false,
  "auto_signature_update": {
    "day_of_week": "any",
    "enable": false,
    "time_of_day": "time_of_day6",
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

