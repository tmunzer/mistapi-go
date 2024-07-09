
# Org Setting

Org Settings

## Structure

`OrgSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored. |
| `ApiPolicy` | [`*models.OrgSettingApiPolicy`](../../doc/models/org-setting-api-policy.md) | Optional | - |
| `AutoDeviceNaming` | [`*models.OrgSettingAutoDeviceNaming`](../../doc/models/org-setting-auto-device-naming.md) | Optional | - |
| `AutoDeviceprofileAssignment` | [`*models.OrgSettingAutoDeviceprofileAssignment`](../../doc/models/org-setting-auto-deviceprofile-assignment.md) | Optional | - |
| `AutoSiteAssignment` | [`*models.OrgSettingAutoSiteAssignment`](../../doc/models/org-setting-auto-site-assignment.md) | Optional | - |
| `BlacklistUrl` | `*string` | Optional | - |
| `Cacerts` | `[]string` | Optional | list of PEM-encoded ca certs |
| `Celona` | [`*models.OrgSettingCelona`](../../doc/models/org-setting-celona.md) | Optional | - |
| `Cloudshark` | [`*models.OrgSettingCloudshark`](../../doc/models/org-setting-cloudshark.md) | Optional | - |
| `Cradlepoint` | [`*models.AccountCradlepointConfig`](../../doc/models/account-cradlepoint-config.md) | Optional | - |
| `CreatedTime` | `*float64` | Optional | - |
| `DeviceCert` | [`*models.OrgSettingDeviceCert`](../../doc/models/org-setting-device-cert.md) | Optional | common device cert, optional |
| `DeviceUpdownThreshold` | `*int` | Optional | enable threshold-based device down delivery via<br><br>1) device-updowns webhooks topic,<br>2) Mist Alert Framework; e.g. send AP/SW/GW down event only if AP/SW/GW Up is not seen within the threshold in minutes; 0 - 240, default is 0 (trigger immediate)<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 30` |
| `DisablePcap` | `*bool` | Optional | whether to disallow Mist to analyze pcap files (this is required for marvis pcap)<br>**Default**: `false` |
| `DisableRemoteShell` | `*bool` | Optional | whether to disable remote shell access for an entire org<br>**Default**: `false` |
| `ForSite` | `*bool` | Optional | - |
| `GatewayMgmt` | [`*models.OrgSettingGatewayMgmt`](../../doc/models/org-setting-gateway-mgmt.md) | Optional | - |
| `GatewayUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored. |
| `Id` | `*uuid.UUID` | Optional | - |
| `Installer` | [`*models.OrgSettingInstaller`](../../doc/models/org-setting-installer.md) | Optional | - |
| `Jcloud` | [`*models.OrgSettingJcloud`](../../doc/models/org-setting-jcloud.md) | Optional | - |
| `Juniper` | [`*models.AccountJuniperInfo`](../../doc/models/account-juniper-info.md) | Optional | - |
| `Mgmt` | [`*models.OrgSettingMgmt`](../../doc/models/org-setting-mgmt.md) | Optional | management-related properties |
| `MistNac` | [`*models.OrgSettingMistNac`](../../doc/models/org-setting-mist-nac.md) | Optional | - |
| `ModifiedTime` | `*float64` | Optional | - |
| `MspId` | `*uuid.UUID` | Optional | - |
| `MxedgeFipsEnabled` | `*bool` | Optional | **Default**: `false` |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PasswordPolicy` | [`*models.OrgSettingPasswordPolicy`](../../doc/models/org-setting-password-policy.md) | Optional | password policy |
| `Pcap` | [`*models.OrgSettingPcap`](../../doc/models/org-setting-pcap.md) | Optional | - |
| `PcapBucketVerified` | `*bool` | Optional | - |
| `Security` | [`*models.OrgSettingSecurity`](../../doc/models/org-setting-security.md) | Optional | - |
| `SimpleAlert` | [`*models.SimpleAlert`](../../doc/models/simple-alert.md) | Optional | Set of heuristic rules will be enabled when marvis subscription is not available.<br>It triggers when, in a Z minute window, there are more than Y distinct client encountring over X failures |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `SwitchMgmt` | [`*models.OrgSettingSwitchMgmt`](../../doc/models/org-setting-switch-mgmt.md) | Optional | - |
| `SwitchUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored. |
| `Synthetictest` | [`*models.SynthetictestConfig`](../../doc/models/synthetictest-config.md) | Optional | - |
| `Tags` | `[]string` | Optional | list of tags |
| `UiIdleTimeout` | `*int` | Optional | automatically logout the user when UI session is inactive. `0` means disabled<br>**Default**: `0`<br>**Constraints**: `>= 0`, `<= 480` |
| `VpnOptions` | [`*models.OrgSettingVpnOptions`](../../doc/models/org-setting-vpn-options.md) | Optional | - |
| `WanPma` | [`*models.OrgSettingWanPma`](../../doc/models/org-setting-wan-pma.md) | Optional | - |
| `WiredPma` | [`*models.OrgSettingWiredPma`](../../doc/models/org-setting-wired-pma.md) | Optional | - |
| `WirelessPma` | [`*models.OrgSettingWirelessPma`](../../doc/models/org-setting-wireless-pma.md) | Optional | - |

## Example (as JSON)

```json
{
  "blacklist_url": "https://papi.s3.amazonaws.com/blacklist/xxx...",
  "device_updown_threshold": 0,
  "disable_pcap": false,
  "disable_remote_shell": false,
  "gateway_updown_threshold": 10,
  "mxedge_fips_enabled": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "switch_updown_threshold": 0,
  "ui_idle_timeout": 10,
  "ap_updown_threshold": 56,
  "api_policy": {
    "no_reveal": false
  },
  "auto_device_naming": {
    "enable": false,
    "rules": [
      {
        "expression": "expression4",
        "match_device_type": "gateway",
        "match_model": "match_model0",
        "model": "model4",
        "prefix": "prefix6",
        "src": "model"
      },
      {
        "expression": "expression4",
        "match_device_type": "gateway",
        "match_model": "match_model0",
        "model": "model4",
        "prefix": "prefix6",
        "src": "model"
      }
    ]
  },
  "auto_deviceprofile_assignment": {
    "enable": false,
    "rules": [
      {
        "expression": "expression4",
        "match_device_type": "gateway",
        "match_model": "match_model0",
        "model": "model4",
        "prefix": "prefix6",
        "src": "model"
      }
    ]
  },
  "auto_site_assignment": {
    "enable": false,
    "rules": [
      {
        "expression": "expression4",
        "match_device_type": "gateway",
        "match_model": "match_model0",
        "model": "model4",
        "prefix": "prefix6",
        "src": "model"
      }
    ]
  }
}
```

