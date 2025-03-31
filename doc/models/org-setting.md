
# Org Setting

Org Settings

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 240` |
| `ApiPolicy` | [`*models.OrgSettingApiPolicy`](../../doc/models/org-setting-api-policy.md) | Optional | - |
| `AutoDeviceNaming` | [`*models.OrgSettingAutoDeviceNaming`](../../doc/models/org-setting-auto-device-naming.md) | Optional | - |
| `AutoDeviceprofileAssignment` | [`*models.OrgSettingAutoDeviceprofileAssignment`](../../doc/models/org-setting-auto-deviceprofile-assignment.md) | Optional | - |
| `AutoSiteAssignment` | [`*models.OrgSettingAutoSiteAssignment`](../../doc/models/org-setting-auto-site-assignment.md) | Optional | - |
| `BlacklistUrl` | `*string` | Optional | - |
| `Cacerts` | `[]string` | Optional | List of PEM-encoded ca certs |
| `Celona` | [`*models.OrgSettingCelona`](../../doc/models/org-setting-celona.md) | Optional | - |
| `Cloudshark` | [`*models.OrgSettingCloudshark`](../../doc/models/org-setting-cloudshark.md) | Optional | - |
| `Cradlepoint` | [`*models.OrgSettingCradlepoint`](../../doc/models/org-setting-cradlepoint.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `DeviceCert` | [`*models.OrgSettingDeviceCert`](../../doc/models/org-setting-device-cert.md) | Optional | common device cert, optional |
| `DeviceUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery via<br><br>* device-updowns webhooks topic,<br>* Mist Alert Framework; e.g. send AP/SW/GW down event only if AP/SW/GW Up is not seen within the threshold in minutes; 0 - 240, default is 0 (trigger immediate)<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 240` |
| `DisablePcap` | `*bool` | Optional | Whether to disallow Mist to analyze pcap files (this is required for marvis pcap)<br>**Default**: `false` |
| `DisableRemoteShell` | `*bool` | Optional | Whether to disable remote shell access for an entire org<br>**Default**: `false` |
| `ForSite` | `*bool` | Optional | - |
| `GatewayMgmt` | [`*models.OrgSettingGatewayMgmt`](../../doc/models/org-setting-gateway-mgmt.md) | Optional | - |
| `GatewayUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 240` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Installer` | [`*models.OrgSettingInstaller`](../../doc/models/org-setting-installer.md) | Optional | - |
| `Jcloud` | [`*models.OrgSettingJcloud`](../../doc/models/org-setting-jcloud.md) | Optional | - |
| `JcloudRa` | [`*models.OrgSettingJcloudRa`](../../doc/models/org-setting-jcloud-ra.md) | Optional | JCloud Routing Assurance connexion |
| `Juniper` | [`*models.AccountJuniperInfo`](../../doc/models/account-juniper-info.md) | Optional | - |
| `JunosShellAccess` | [`*models.OrgSettingJunosShellAccess`](../../doc/models/org-setting-junos-shell-access.md) | Optional | by default, webshell access is only enabled for Admin user |
| `Mgmt` | [`*models.OrgSettingMgmt`](../../doc/models/org-setting-mgmt.md) | Optional | management-related properties |
| `MistNac` | [`*models.OrgSettingMistNac`](../../doc/models/org-setting-mist-nac.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | - |
| `OpticPortConfig` | [`map[string]models.OpticPortConfigPort`](../../doc/models/optic-port-config-port.md) | Optional | Property key is the interface name or range (e.g. `et-0/0/47`, `et-0/0/48-49`) |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PasswordPolicy` | [`*models.OrgSettingPasswordPolicy`](../../doc/models/org-setting-password-policy.md) | Optional | password policy |
| `Pcap` | [`*models.OrgSettingPcap`](../../doc/models/org-setting-pcap.md) | Optional | - |
| `PcapBucketVerified` | `*bool` | Optional | - |
| `Security` | [`*models.OrgSettingSecurity`](../../doc/models/org-setting-security.md) | Optional | - |
| `SimpleAlert` | [`*models.SimpleAlert`](../../doc/models/simple-alert.md) | Optional | Set of heuristic rules will be enabled when marvis subscription is not available. It triggers when, in a Z minute window, there are more than Y distinct client encountering over X failures |
| `SwitchMgmt` | [`*models.OrgSettingSwitchMgmt`](../../doc/models/org-setting-switch-mgmt.md) | Optional | - |
| `SwitchUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.<br>**Default**: `0` |
| `SyntheticTest` | [`*models.SynthetictestConfig`](../../doc/models/synthetictest-config.md) | Optional | - |
| `Tags` | `[]string` | Optional | List of tags |
| `UiIdleTimeout` | `*int` | Optional | Automatically logout the user when UI session is inactive. `0` means disabled<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 480` |
| `VpnOptions` | [`*models.OrgSettingVpnOptions`](../../doc/models/org-setting-vpn-options.md) | Optional | - |
| `WanPma` | [`*models.OrgSettingWanPma`](../../doc/models/org-setting-wan-pma.md) | Optional | - |
| `WiredPma` | [`*models.OrgSettingWiredPma`](../../doc/models/org-setting-wired-pma.md) | Optional | - |
| `WirelessPma` | [`*models.OrgSettingWirelessPma`](../../doc/models/org-setting-wireless-pma.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_updown_threshold": 0,
  "blacklist_url": "https://papi.s3.amazonaws.com/blacklist/xxx...",
  "device_updown_threshold": 0,
  "disable_pcap": false,
  "disable_remote_shell": false,
  "gateway_updown_threshold": 10,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "msp_id": "b9d42c2e-88ee-41f8-b798-f009ce7fe909",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "switch_updown_threshold": 0,
  "ui_idle_timeout": 10,
  "api_policy": {
    "no_reveal": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "auto_device_naming": {
    "enable": false,
    "rules": [
      {
        "expression": "expression4",
        "match_device": "ap",
        "prefix": "prefix6",
        "src": "lldp_port_desc",
        "suffix": "suffix2",
        "exampleAdditionalProperty": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "expression": "expression4",
        "match_device": "ap",
        "prefix": "prefix6",
        "src": "lldp_port_desc",
        "suffix": "suffix2",
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
  "auto_deviceprofile_assignment": {
    "enable": false,
    "rules": [
      {
        "create_new_site_if_needed": false,
        "expression": "expression4",
        "gatewaytemplate_id": "gatewaytemplate_id0",
        "match_country": "match_country8",
        "match_device_type": "switch",
        "src": "name",
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
  "auto_site_assignment": {
    "enable": false,
    "rules": [
      {
        "create_new_site_if_needed": false,
        "expression": "expression4",
        "gatewaytemplate_id": "gatewaytemplate_id0",
        "match_country": "match_country8",
        "match_device_type": "switch",
        "src": "name",
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

