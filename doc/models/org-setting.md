
# Org Setting

Organization-wide feature, integration, management, and security settings

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSetting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowMist` | `*bool` | Optional | whether to allow Mist to look at this org<br><br>**Default**: `false` |
| `ApUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for AP devices only. When configured it takes effect for AP devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `ApiPolicy` | [`*models.OrgSettingApiPolicy`](../../doc/models/org-setting-api-policy.md) | Optional | Organization API response policy for hiding secrets and passwords |
| `AutoDeviceNaming` | [`*models.OrgSettingAutoDeviceNaming`](../../doc/models/org-setting-auto-device-naming.md) | Optional | Automatic device naming configuration for claimed devices |
| `AutoDeviceprofileAssignment` | [`*models.OrgSettingAutoDeviceprofileAssignment`](../../doc/models/org-setting-auto-deviceprofile-assignment.md) | Optional | Automatic device profile assignment configuration |
| `AutoSiteAssignment` | [`*models.OrgSettingAutoSiteAssignment`](../../doc/models/org-setting-auto-site-assignment.md) | Optional | Automatic site assignment configuration for claimed devices |
| `AutoUpgrade` | [`*models.OrgSettingAutoUpgrade`](../../doc/models/org-setting-auto-upgrade.md) | Optional | Organization-wide AP automatic firmware upgrade policy |
| `BlacklistUrl` | `*string` | Optional, Read-only | Read-only URL for the organization blacklist file |
| `Cacerts` | `[]string` | Optional | RADSec certificates for AP |
| `Celona` | [`*models.OrgSettingCelona`](../../doc/models/org-setting-celona.md) | Optional | Integration settings for Celona |
| `Cloudshark` | [`*models.OrgSettingCloudshark`](../../doc/models/org-setting-cloudshark.md) | Optional | Packet capture integration settings for CloudShark |
| `Cradlepoint` | [`*models.OrgSettingCradlepoint`](../../doc/models/org-setting-cradlepoint.md) | Optional, Read-only | Read-only Cradlepoint integration settings stored for the organization |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `DeviceCert` | [`*models.OrgSettingDeviceCert`](../../doc/models/org-setting-device-cert.md) | Optional | Optional common device certificate configuration for organization settings |
| `DeviceUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery via<br><br>* device-updowns webhooks topic,<br>* Mist Alert Framework; e.g. send AP/SW/GW down event only if AP/SW/GW Up is not seen within the threshold in minutes; 0 - 240, default is 0 (trigger immediate)<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `DisablePcap` | `*bool` | Optional | Whether to disallow Mist to analyze pcap files (this is required for marvis pcap)<br><br>**Default**: `false` |
| `DisableRemoteShell` | `*bool` | Optional | Whether to disable remote shell access for an entire org<br><br>**Default**: `false` |
| `ForSite` | `*bool` | Optional, Read-only | Read-only indicator that the settings object is scoped to a site |
| `GatewayMgmt` | [`*models.OrgSettingGatewayMgmt`](../../doc/models/org-setting-gateway-mgmt.md) | Optional | Organization-level gateway management settings |
| `GatewayTunnelUpdownThreshold` | `models.Optional[int]` | Optional | enable threshold-based gateway tunnel (secure edge tunnels) up-down delivery.<br><br>**Constraints**: `>= 0` |
| `GatewayUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Gateway devices only. When configured it takes effect for GW devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 240` |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Installer` | [`*models.OrgSettingInstaller`](../../doc/models/org-setting-installer.md) | Optional | Organization-level permissions and grace period for installer workflows |
| `Jcloud` | [`*models.OrgSettingJcloud`](../../doc/models/org-setting-jcloud.md) | Optional | JCloud integration settings for this Mist organization |
| `JcloudRa` | [`*models.OrgSettingJcloudRa`](../../doc/models/org-setting-jcloud-ra.md) | Optional | JCloud Routing Assurance connexion |
| `Juniper` | [`*models.AccountJuniperInfo`](../../doc/models/account-juniper-info.md) | Optional | Linked Juniper account information returned by the integration |
| `JuniperSrx` | [`*models.OrgSettingJuniperSrx`](../../doc/models/org-setting-juniper-srx.md) | Optional | Organization settings for Juniper SRX devices |
| `JunosShellAccess` | [`*models.OrgSettingJunosShellAccess`](../../doc/models/org-setting-junos-shell-access.md) | Optional | junos_shell_access: Manages role-based web-shell access.  <br>When junos_shell access is not defined (Default) - No additional users are configured and web-shell uses default `mist` user to login.  <br>When junos_shell_access is defined - Additional users mist-web-admin (admin permission), mist-web-viewer(viewer permission) are configured on the device and web-shell logs in with the mist-web-admin/mist-web-viewer user depending upon the shell access level. Setting the shell access level to "none", disables web-shell access for that specific role. |
| `Marvis` | [`*models.OrgSettingMarvis`](../../doc/models/org-setting-marvis.md) | Optional | Organization settings for Marvis automation |
| `Mgmt` | [`*models.OrgSettingMgmt`](../../doc/models/org-setting-mgmt.md) | Optional | Organization management connectivity settings |
| `MistNac` | [`*models.OrgSettingMistNac`](../../doc/models/org-setting-mist-nac.md) | Optional | Organization-level Mist NAC configuration |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `MxedgeMgmt` | [`*models.MxedgeMgmt`](../../doc/models/mxedge-mgmt.md) | Optional | Management settings for a Mist Edge appliance |
| `OpticPortConfig` | [`map[string]models.OpticPortConfigPort`](../../doc/models/optic-port-config-port.md) | Optional | Property key is the interface name or range (e.g. `et-0/0/47`, `et-0/0/48-49`) |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PasswordPolicy` | [`*models.OrgSettingPasswordPolicy`](../../doc/models/org-setting-password-policy.md) | Optional | Admin credential policy settings for the organization |
| `Pcap` | [`*models.OrgSettingPcap`](../../doc/models/org-setting-pcap.md) | Optional | Packet capture export settings for the organization |
| `PcapBucketVerified` | `*bool` | Optional, Read-only | Whether the configured packet capture bucket has been verified |
| `Security` | [`*models.OrgSettingSecurity`](../../doc/models/org-setting-security.md) | Optional | Organization security controls for local SSH and FIPS zeroize access |
| `SimpleAlert` | [`*models.SimpleAlert`](../../doc/models/simple-alert.md) | Optional | Heuristic alert thresholds used when a Marvis subscription is unavailable |
| `Ssr` | [`*models.SettingSsr`](../../doc/models/setting-ssr.md) | Optional | SSR management settings for device onboarding and connectivity |
| `Switch` | [`*models.OrgSettingSwitch`](../../doc/models/org-setting-switch.md) | Optional | Configuration defaults for switches in this organization |
| `SwitchMgmt` | [`*models.OrgSettingSwitchMgmt`](../../doc/models/org-setting-switch-mgmt.md) | Optional | Organization-level switch management settings |
| `SwitchUpdownThreshold` | `models.Optional[int]` | Optional | Enable threshold-based device down delivery for Switch devices only. When configured it takes effect for SW devices and `device_updown_threshold` is ignored.<br><br>**Default**: `0` |
| `SyntheticTest` | [`*models.SynthetictestConfig`](../../doc/models/synthetictest-config.md) | Optional | Synthetic test configuration for Marvis Minis |
| `Tags` | `[]string` | Optional | Labels associated with these organization settings |
| `UiIdleTimeout` | `*int` | Optional | Automatically logout the user when UI session is inactive. `0` means disabled<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 480` |
| `UiNoTracking` | `*bool` | Optional | Whether UI usage tracking is disabled for the organization<br><br>**Default**: `false` |
| `VpnOptions` | [`*models.OrgSettingVpnOptions`](../../doc/models/org-setting-vpn-options.md) | Optional | Organization VPN behavior options |
| `WanPma` | [`*models.OrgSettingWanPma`](../../doc/models/org-setting-wan-pma.md) | Optional | PMA feature settings for WAN Assurance |
| `WiredPma` | [`*models.OrgSettingWiredPma`](../../doc/models/org-setting-wired-pma.md) | Optional | PMA feature settings for Wired Assurance |
| `WirelessPma` | [`*models.OrgSettingWirelessPma`](../../doc/models/org-setting-wireless-pma.md) | Optional | PMA feature settings for Wireless Assurance |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSetting := models.OrgSetting{
        AllowMist:                    models.ToPointer(false),
        ApUpdownThreshold:            models.NewOptional(models.ToPointer(0)),
        ApiPolicy:                    models.ToPointer(models.OrgSettingApiPolicy{
            NoReveal:             models.ToPointer(false),
            SrcIps:               []string{
                "src_ips6",
                "src_ips7",
            },
        }),
        AutoDeviceNaming:             models.ToPointer(models.OrgSettingAutoDeviceNaming{
            Enable:               models.ToPointer(false),
            Rules:                models.NewOptional(models.ToPointer([]models.OrgSettingAutoDeviceNamingRule{
                models.OrgSettingAutoDeviceNamingRule{
                    Expression:           models.ToPointer("expression4"),
                    MatchDevice:          models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
                    Prefix:               models.ToPointer("prefix6"),
                    Src:                  models.ToPointer(models.OrgSettingAutoDeviceNamingRuleSrcEnum_LLDPPORTDESC),
                    Suffix:               models.ToPointer("suffix2"),
                },
                models.OrgSettingAutoDeviceNamingRule{
                    Expression:           models.ToPointer("expression4"),
                    MatchDevice:          models.ToPointer(models.DeviceTypeDefaultApEnum_AP),
                    Prefix:               models.ToPointer("prefix6"),
                    Src:                  models.ToPointer(models.OrgSettingAutoDeviceNamingRuleSrcEnum_LLDPPORTDESC),
                    Suffix:               models.ToPointer("suffix2"),
                },
            })),
        }),
        AutoDeviceprofileAssignment:  models.ToPointer(models.OrgSettingAutoDeviceprofileAssignment{
            Enable:               models.ToPointer(false),
            Rules:                models.NewOptional(models.ToPointer([]models.OrgSettingAutoAssignmentRule{
                models.OrgSettingAutoAssignmentRule{
                    CreateNewSiteIfNeeded: models.ToPointer(false),
                    Expression:            models.NewOptional(models.ToPointer("expression4")),
                    GatewaytemplateId:     models.ToPointer("gatewaytemplate_id0"),
                    MatchCountry:          models.ToPointer("match_country8"),
                    MatchDeviceType:       models.ToPointer(models.DeviceTypeDefaultApEnum_ENUMSWITCH),
                    Src:                   models.OrgSettingAutoSiteAssignmentSrcEnum_NAME,
                },
            })),
        }),
        BlacklistUrl:                 models.ToPointer("https://papi.s3.amazonaws.com/blacklist/xxx..."),
        DeviceUpdownThreshold:        models.NewOptional(models.ToPointer(0)),
        DisablePcap:                  models.ToPointer(false),
        DisableRemoteShell:           models.ToPointer(false),
        GatewayUpdownThreshold:       models.NewOptional(models.ToPointer(10)),
        Id:                           models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MspId:                        models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        OrgId:                        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SwitchUpdownThreshold:        models.NewOptional(models.ToPointer(0)),
        UiIdleTimeout:                models.ToPointer(10),
        UiNoTracking:                 models.ToPointer(false),
        AdditionalProperties:         map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

