
# Switch Mgmt

Switch settings

## Structure

`SwitchMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApAffinityThreshold` | `*int` | Optional | ap_affinity_threshold ap_affinity_threshold can be added as a field under site/setting. By default this value is set to 12. If the field is set in both site/setting and org/setting, the value from site/setting will be used.<br>**Default**: `10` |
| `CliBanner` | `*string` | Optional | Set Banners for switches. Allows markup formatting |
| `CliIdleTimeout` | `*int` | Optional | Sets timeout for switches<br>**Constraints**: `>= 1`, `<= 60` |
| `ConfigRevertTimer` | `*int` | Optional | the rollback timer for commit confirmed<br>**Default**: `10`<br>**Constraints**: `>= 1`, `<= 30` |
| `DhcpOptionFqdn` | `*bool` | Optional | Enable to provide the FQDN with DHCP option 81<br>**Default**: `false` |
| `DisableOobDownAlarm` | `*bool` | Optional | - |
| `LocalAccounts` | [`map[string]models.ConfigSwitchLocalAccountsUser`](../../doc/models/config-switch-local-accounts-user.md) | Optional | Property key is the user name. For Local user authentication |
| `MxedgeProxyHost` | `*string` | Optional | - |
| `MxedgeProxyPort` | `*int` | Optional | **Default**: `2222`<br>**Constraints**: `>= 1`, `<= 65535` |
| `ProtectRe` | [`*models.ProtectRe`](../../doc/models/protect-re.md) | Optional | restrict inbound-traffic to host<br>when enabled, all traffic that is not essential to our operation will be dropped<br>e.g. ntp / dns / traffic to mist will be allowed by default, if dhcpd is enabled, we'll make sure it works |
| `Radius` | [`*models.SwitchRadius`](../../doc/models/switch-radius.md) | Optional | by default, `radius_config` will be used. if a different one has to be used set `use_different_radius |
| `RootPassword` | `*string` | Optional | - |
| `Tacacs` | [`*models.Tacacs`](../../doc/models/tacacs.md) | Optional | - |
| `UseMxedgeProxy` | `*bool` | Optional | to use mxedge as proxy |

## Example (as JSON)

```json
{
  "ap_affinity_threshold": 10,
  "cli_banner": "\\t\\tWELCOME!",
  "config_revert_timer": 10,
  "dhcp_option_fqdn": false,
  "mxedge_proxy_port": 2222,
  "cli_idle_timeout": 142
}
```

